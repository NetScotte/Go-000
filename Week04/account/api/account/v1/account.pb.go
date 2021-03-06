// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package account

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

type Id struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*Id)(nil), "account.Id")
	proto.RegisterType((*UserInfo)(nil), "account.UserInfo")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x44, 0xb8,
	0x98, 0x3c, 0x53, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83,
	0x98, 0x32, 0x53, 0x94, 0x1c, 0xb8, 0x38, 0x42, 0x8b, 0x53, 0x8b, 0x3c, 0xf3, 0xd2, 0xf2, 0xd1,
	0xe5, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83,
	0xc0, 0x6c, 0x21, 0x01, 0x2e, 0xe6, 0xc4, 0xf4, 0x54, 0x09, 0x66, 0xb0, 0x22, 0x10, 0xd3, 0xc8,
	0x8a, 0x0b, 0x66, 0x85, 0x90, 0x3e, 0x17, 0xb7, 0x7b, 0x6a, 0x09, 0xc8, 0x3c, 0xa7, 0x4a, 0xcf,
	0x14, 0x21, 0x6e, 0x3d, 0x98, 0x53, 0x3c, 0x53, 0xa4, 0x04, 0xe1, 0x1c, 0x98, 0x7d, 0x4a, 0x0c,
	0x4e, 0xdc, 0x51, 0x9c, 0x7a, 0xd6, 0x50, 0xf1, 0x24, 0x36, 0xb0, 0x83, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x03, 0x93, 0x1c, 0x6e, 0xc1, 0x00, 0x00, 0x00,
}
