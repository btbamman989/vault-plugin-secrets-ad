// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sdk/logical/identity.proto

package logical

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Entity struct {
	// ID is the unique identifier for the entity
	ID string `sentinel:"" protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Name is the human-friendly unique identifier for the entity
	Name string `sentinel:"" protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Aliases contains thhe alias mappings for the given entity
	Aliases []*Alias `sentinel:"" protobuf:"bytes,3,rep,name=aliases,proto3" json:"aliases,omitempty"`
	// Metadata represents the custom data tied to this entity
	Metadata             map[string]string `sentinel:"" protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a34d35719c603a1, []int{0}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entity) GetAliases() []*Alias {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *Entity) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Alias struct {
	// MountType is the backend mount's type to which this identity belongs
	MountType string `sentinel:"" protobuf:"bytes,1,opt,name=mount_type,json=mountType,proto3" json:"mount_type,omitempty"`
	// MountAccessor is the identifier of the mount entry to which this
	// identity belongs
	MountAccessor string `sentinel:"" protobuf:"bytes,2,opt,name=mount_accessor,json=mountAccessor,proto3" json:"mount_accessor,omitempty"`
	// Name is the identifier of this identity in its authentication source
	Name string `sentinel:"" protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Metadata represents the custom data tied to this alias
	Metadata             map[string]string `sentinel:"" protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}
func (*Alias) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a34d35719c603a1, []int{1}
}

func (m *Alias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alias.Unmarshal(m, b)
}
func (m *Alias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alias.Marshal(b, m, deterministic)
}
func (m *Alias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alias.Merge(m, src)
}
func (m *Alias) XXX_Size() int {
	return xxx_messageInfo_Alias.Size(m)
}
func (m *Alias) XXX_DiscardUnknown() {
	xxx_messageInfo_Alias.DiscardUnknown(m)
}

var xxx_messageInfo_Alias proto.InternalMessageInfo

func (m *Alias) GetMountType() string {
	if m != nil {
		return m.MountType
	}
	return ""
}

func (m *Alias) GetMountAccessor() string {
	if m != nil {
		return m.MountAccessor
	}
	return ""
}

func (m *Alias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Alias) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Entity)(nil), "logical.Entity")
	proto.RegisterMapType((map[string]string)(nil), "logical.Entity.MetadataEntry")
	proto.RegisterType((*Alias)(nil), "logical.Alias")
	proto.RegisterMapType((map[string]string)(nil), "logical.Alias.MetadataEntry")
}

func init() { proto.RegisterFile("sdk/logical/identity.proto", fileDescriptor_4a34d35719c603a1) }

var fileDescriptor_4a34d35719c603a1 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x4f, 0x6b, 0x83, 0x40,
	0x10, 0xc5, 0x51, 0xf3, 0xa7, 0x99, 0x12, 0x29, 0x4b, 0x0f, 0x12, 0x1a, 0x08, 0x81, 0x16, 0x4f,
	0x0a, 0xed, 0x25, 0x6d, 0x4f, 0x29, 0xc9, 0x21, 0x87, 0x5e, 0xa4, 0xa7, 0x5e, 0xca, 0x44, 0x97,
	0xb8, 0x44, 0x5d, 0x71, 0xc7, 0x80, 0x5f, 0xb2, 0xe7, 0x7e, 0x9c, 0x92, 0x75, 0x23, 0x09, 0x3d,
	0xf7, 0x36, 0xfe, 0xde, 0x38, 0xfb, 0xe6, 0x0d, 0x4c, 0x54, 0xb2, 0x0f, 0x33, 0xb9, 0x13, 0x31,
	0x66, 0xa1, 0x48, 0x78, 0x41, 0x82, 0x9a, 0xa0, 0xac, 0x24, 0x49, 0x36, 0x34, 0x7c, 0xfe, 0x6d,
	0xc1, 0x60, 0xad, 0x15, 0xe6, 0x82, 0xbd, 0x59, 0x79, 0xd6, 0xcc, 0xf2, 0x47, 0x91, 0xbd, 0x59,
	0x31, 0x06, 0xbd, 0x02, 0x73, 0xee, 0xd9, 0x9a, 0xe8, 0x9a, 0xf9, 0x30, 0xc4, 0x4c, 0xa0, 0xe2,
	0xca, 0x73, 0x66, 0x8e, 0x7f, 0xfd, 0xe8, 0x06, 0x66, 0x52, 0xb0, 0x3c, 0xf2, 0xe8, 0x24, 0xb3,
	0x67, 0xb8, 0xca, 0x39, 0x61, 0x82, 0x84, 0x5e, 0x4f, 0xb7, 0x4e, 0xbb, 0xd6, 0xf6, 0xc1, 0xe0,
	0xdd, 0xe8, 0xeb, 0x82, 0xaa, 0x26, 0xea, 0xda, 0x27, 0xaf, 0x30, 0xbe, 0x90, 0xd8, 0x0d, 0x38,
	0x7b, 0xde, 0x18, 0x6b, 0xc7, 0x92, 0xdd, 0x42, 0xff, 0x80, 0x59, 0x7d, 0x32, 0xd7, 0x7e, 0xbc,
	0xd8, 0x0b, 0x6b, 0xfe, 0x63, 0x41, 0x5f, 0x5b, 0x61, 0x53, 0x80, 0x5c, 0xd6, 0x05, 0x7d, 0x51,
	0x53, 0x72, 0xf3, 0xf3, 0x48, 0x93, 0x8f, 0xa6, 0xe4, 0xec, 0x1e, 0xdc, 0x56, 0xc6, 0x38, 0xe6,
	0x4a, 0xc9, 0xca, 0xcc, 0x1a, 0x6b, 0xba, 0x34, 0xb0, 0x4b, 0xc1, 0x39, 0x4b, 0x61, 0xf1, 0x67,
	0xb7, 0xbb, 0xcb, 0x18, 0xfe, 0x65, 0xb5, 0x37, 0xff, 0xf3, 0x61, 0x27, 0x28, 0xad, 0xb7, 0x41,
	0x2c, 0xf3, 0x30, 0x45, 0x95, 0x8a, 0x58, 0x56, 0x65, 0x78, 0xc0, 0x3a, 0xa3, 0xf0, 0xec, 0xda,
	0xdb, 0x81, 0xbe, 0xf2, 0xd3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0xa9, 0x44, 0x63, 0x03,
	0x02, 0x00, 0x00,
}
