// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GM.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateItem_C struct {
	CfgId                *string  `protobuf:"bytes,1,opt,name=cfgId" json:"cfgId,omitempty"`
	Count                *int32   `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateItem_C) Reset()         { *m = CreateItem_C{} }
func (m *CreateItem_C) String() string { return proto.CompactTextString(m) }
func (*CreateItem_C) ProtoMessage()    {}
func (*CreateItem_C) Descriptor() ([]byte, []int) {
	return fileDescriptor_GM_5ddcf347e252f830, []int{0}
}
func (m *CreateItem_C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateItem_C.Unmarshal(m, b)
}
func (m *CreateItem_C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateItem_C.Marshal(b, m, deterministic)
}
func (dst *CreateItem_C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateItem_C.Merge(dst, src)
}
func (m *CreateItem_C) XXX_Size() int {
	return xxx_messageInfo_CreateItem_C.Size(m)
}
func (m *CreateItem_C) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateItem_C.DiscardUnknown(m)
}

var xxx_messageInfo_CreateItem_C proto.InternalMessageInfo

func (m *CreateItem_C) GetCfgId() string {
	if m != nil && m.CfgId != nil {
		return *m.CfgId
	}
	return ""
}

func (m *CreateItem_C) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type ChangeLevel_C struct {
	Level                *int32   `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeLevel_C) Reset()         { *m = ChangeLevel_C{} }
func (m *ChangeLevel_C) String() string { return proto.CompactTextString(m) }
func (*ChangeLevel_C) ProtoMessage()    {}
func (*ChangeLevel_C) Descriptor() ([]byte, []int) {
	return fileDescriptor_GM_5ddcf347e252f830, []int{1}
}
func (m *ChangeLevel_C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeLevel_C.Unmarshal(m, b)
}
func (m *ChangeLevel_C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeLevel_C.Marshal(b, m, deterministic)
}
func (dst *ChangeLevel_C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeLevel_C.Merge(dst, src)
}
func (m *ChangeLevel_C) XXX_Size() int {
	return xxx_messageInfo_ChangeLevel_C.Size(m)
}
func (m *ChangeLevel_C) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeLevel_C.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeLevel_C proto.InternalMessageInfo

func (m *ChangeLevel_C) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateItem_C)(nil), "message.CreateItem_C")
	proto.RegisterType((*ChangeLevel_C)(nil), "message.ChangeLevel_C")
}

func init() { proto.RegisterFile("GM.proto", fileDescriptor_GM_5ddcf347e252f830) }

var fileDescriptor_GM_5ddcf347e252f830 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x70, 0xf7, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0xe2, 0x71,
	0x4a, 0x2c, 0x4e, 0x0d, 0xf3, 0x87, 0x08, 0x2b, 0x59, 0x71, 0xf1, 0x38, 0x17, 0xa5, 0x26, 0x96,
	0xa4, 0x7a, 0x96, 0xa4, 0xe6, 0xc6, 0x3b, 0x0b, 0x89, 0x70, 0xb1, 0x26, 0xa7, 0xa5, 0x7b, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x60, 0xd1, 0xfc, 0xd2, 0xbc, 0x12, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0x49, 0x95, 0x8b, 0xd7, 0x39, 0x23, 0x31, 0x2f,
	0x3d, 0xd5, 0x27, 0xb5, 0x2c, 0x35, 0x07, 0xa2, 0x39, 0x07, 0xc4, 0x04, 0x6b, 0x66, 0x0d, 0x82,
	0x70, 0x9c, 0x04, 0xb9, 0xb8, 0x93, 0xf3, 0x73, 0xf5, 0xa0, 0xf6, 0x3b, 0x31, 0xb9, 0xfb, 0x02,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x27, 0x6d, 0x0a, 0x97, 0x00, 0x00, 0x00,
}
