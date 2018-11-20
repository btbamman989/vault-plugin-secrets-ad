package plugin

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-errors/errors"
	"github.com/hashicorp/vault-plugin-secrets-ad/plugin/util"
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

const (
	credPrefix = "creds/"
	storageKey = "creds"

	// Since Active Directory offers eventual consistency, in testing we found that sometimes
	// Active Directory returned "password last set" times that were _later_ than our own,
	// even though ours were captured after synchronously completing a password update operation.
	//
	// An example we captured was:
	// 		last_vault_rotation     2018-04-18T22:29:57.385454779Z
	// 		password_last_set       2018-04-18T22:29:57.3902786Z
	//
	// Thus we add a short time buffer when checking whether anyone _else_ updated the AD password
	// since Vault last rotated it.
	passwordLastSetBuffer = time.Second

	// Since password TTL can be set to as low as 1 second,
	// we can't cache passwords for an entire second.
	credCacheCleanup    = time.Second / 3
	credCacheExpiration = time.Second / 2
)

// deleteCred fulfills the DeleteWatcher interface in roles.
// It allows the roleHandler to let us know when a role's been deleted so we can delete its associated creds too.
func (b *backend) deleteCred(ctx context.Context, storage logical.Storage, roleName string) error {
	if err := storage.Delete(ctx, storageKey+"/"+roleName); err != nil {
		return err
	}
	b.credCache.Delete(roleName)
	return nil
}

func (b *backend) invalidateCred(ctx context.Context, key string) {
	if strings.HasPrefix(key, credPrefix) {
		roleName := key[len(credPrefix):]
		b.credCache.Delete(roleName)
	}
}

func (b *backend) pathCreds() *framework.Path {
	return &framework.Path{
		Pattern: credPrefix + framework.GenericNameRegex("name"),
		Fields: map[string]*framework.FieldSchema{
			"name": {
				Type:        framework.TypeString,
				Description: "Name of the role",
			},
		},
		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.ReadOperation: b.credReadOperation,
		},
		HelpSynopsis:    credHelpSynopsis,
		HelpDescription: credHelpDescription,
	}
}

func (b *backend) credReadOperation(ctx context.Context, req *logical.Request, fieldData *framework.FieldData) (*logical.Response, error) {
	cred := make(map[string]interface{})

	roleName := fieldData.Get("name").(string)

	// We act upon quite a few things below that could be racy if not locked:
	// 		- Roles. If a new cred is created, the role is updated to include the new LastVaultRotation time,
	//		  effecting role storage (and the role cache, but that's already thread-safe).
	//		- Creds. New creds involve writing to cred storage and the cred cache (also already thread-safe).
	// Rather than setting read locks of different types, and upgrading them to write locks, let's keep complexity
	// low and use one simple mutex.
	b.credLock.Lock()
	defer b.credLock.Unlock()

	role, err := b.readRole(ctx, req.Storage, roleName)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, nil
	}

	var resp *logical.Response
	var respErr error
	var unset time.Time

	switch {

	case role.LastVaultRotation == unset:
		// We've never managed this cred before.
		// We need to rotate the password so Vault will know it.
		resp, respErr = b.generateAndReturnCreds(ctx, req.Storage, roleName, cred)

	case role.PasswordLastSet.After(role.LastVaultRotation.Add(passwordLastSetBuffer)):
		// Someone has manually rotated the password in Active Directory since we last rolled it.
		// We need to rotate it now so Vault will know it and be able to return it.
		resp, respErr = b.generateAndReturnCreds(ctx, req.Storage, roleName, cred)

	default:
		// Since we should know the last password, let's retrieve it now so we can return it with the new one.
		credIfc, found := b.credCache.Get(roleName)
		if found {
			cred = credIfc.(map[string]interface{})
		} else {
			entry, err := req.Storage.Get(ctx, storageKey+"/"+roleName)
			if err != nil {
				return nil, err
			}
			if entry == nil {
				// If the creds aren't in storage, but roles are and we've created creds before,
				// this is an unexpected state and something has gone wrong.
				// Let's be explicit and error about this.
				return nil, fmt.Errorf("should have the creds for %+v but they're not found", role)
			}
			if err := entry.DecodeJSON(&cred); err != nil {
				return nil, err
			}
			b.credCache.SetDefault(roleName, cred)
		}

		// Is the password too old?
		// If so, time for a new one!
		now := time.Now().UTC()
		shouldBeRolled := role.LastVaultRotation.Add(time.Duration(role.TTL) * time.Second) // already in UTC
		if now.After(shouldBeRolled) {
			resp, respErr = b.generateAndReturnCreds(ctx, req.Storage, roleName, cred)
		} else {
			resp = &logical.Response{
				Data: cred,
			}
		}
	}
	if respErr != nil {
		return nil, respErr
	}
	return resp, nil
}

func (b *backend) generateAndReturnCreds(ctx context.Context, storage logical.Storage, roleName string, previousCred map[string]interface{}) (*logical.Response, error) {
	engineConf, err := b.readConfig(ctx, storage)
	if err != nil {
		return nil, err
	}
	if engineConf == nil {
		return nil, errors.New("the config is currently unset")
	}

	// Although role was available to many callers and could be passed in,
	// we check it here to make sure it still exists and isn't nil. This is
	// to prevent the Task Manager from forever rotating the creds of roles
	// that no longer exist.
	role, err := b.readRole(ctx, storage, roleName)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, nil
	}

	newPassword, err := util.GeneratePassword(engineConf.PasswordConf.Formatter, engineConf.PasswordConf.Length)
	if err != nil {
		return nil, err
	}

	if err := b.client.UpdatePassword(engineConf.ADConf, role.ServiceAccountName, newPassword); err != nil {
		return nil, err
	}

	// Time recorded is in UTC for easier user comparison to AD's last rotated time, which is set to UTC by Microsoft.
	role.LastVaultRotation = time.Now().UTC()
	if err := b.writeRole(ctx, storage, roleName, role); err != nil {
		return nil, err
	}

	// Although a service account name is typically my_app@example.com,
	// the username it uses is just my_app, or everything before the @.
	var username string
	fields := strings.Split(role.ServiceAccountName, "@")
	if len(fields) > 0 {
		username = fields[0]
	} else {
		return nil, fmt.Errorf("unable to infer username from service account name: %s", role.ServiceAccountName)
	}

	cred := map[string]interface{}{
		"username":         username,
		"current_password": newPassword,
	}
	if previousCred["current_password"] != nil {
		cred["last_password"] = previousCred["current_password"]
	}

	// Use our TaskTracker to rotate this secret when it expires.
	futureRotation := &Task{
		// There should be only one rotation task per service account, which is why it's used as the identifier here.
		// If other threads later create future rotations, the most recent one will be the only one kept. That's
		// intentional.
		Identifier:   role.ServiceAccountName,
		Type:         TaskTypeRotation,
		ExecuteAfter: role.LastVaultRotation.Add(time.Duration(role.TTL) * time.Second),
		Execute: func(executeCtx context.Context, executeReq *logical.Request) error {
			if _, err := b.generateAndReturnCreds(executeCtx, executeReq.Storage, roleName, previousCred); err != nil {
				return err
			}
			return nil
		},
	}
	b.taskTracker.Upsert(ctx, storage, futureRotation)

	// Cache and save the cred.
	entry, err := logical.StorageEntryJSON(storageKey+"/"+roleName, cred)
	if err != nil {
		return nil, err
	}
	if err := storage.Put(ctx, entry); err != nil {
		return nil, err
	}
	b.credCache.SetDefault(roleName, cred)

	return &logical.Response{
		Data: cred,
	}, nil
}

const (
	credHelpSynopsis = `
Retrieve a role's creds by role name.
`
	credHelpDescription = `
Read creds using a role's name to view the login, current password, and last password.
`
)
