// Code generated by protoc-gen-go. DO NOT EDIT.
// source: AttrsChange.proto

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

type NumberChange_S struct {
	CfgId                *string  `protobuf:"bytes,1,opt,name=cfgId" json:"cfgId,omitempty"`
	Value                *int64   `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberChange_S) Reset()         { *m = NumberChange_S{} }
func (m *NumberChange_S) String() string { return proto.CompactTextString(m) }
func (*NumberChange_S) ProtoMessage()    {}
func (*NumberChange_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_AttrsChange_e40cb90330404d63, []int{0}
}
func (m *NumberChange_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberChange_S.Unmarshal(m, b)
}
func (m *NumberChange_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberChange_S.Marshal(b, m, deterministic)
}
func (dst *NumberChange_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberChange_S.Merge(dst, src)
}
func (m *NumberChange_S) XXX_Size() int {
	return xxx_messageInfo_NumberChange_S.Size(m)
}
func (m *NumberChange_S) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberChange_S.DiscardUnknown(m)
}

var xxx_messageInfo_NumberChange_S proto.InternalMessageInfo

func (m *NumberChange_S) GetCfgId() string {
	if m != nil && m.CfgId != nil {
		return *m.CfgId
	}
	return ""
}

func (m *NumberChange_S) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type HPChange_S struct {
	Guid                 *int64   `protobuf:"varint,1,opt,name=guid" json:"guid,omitempty"`
	CurHp                *int64   `protobuf:"varint,2,opt,name=curHp" json:"curHp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HPChange_S) Reset()         { *m = HPChange_S{} }
func (m *HPChange_S) String() string { return proto.CompactTextString(m) }
func (*HPChange_S) ProtoMessage()    {}
func (*HPChange_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_AttrsChange_e40cb90330404d63, []int{1}
}
func (m *HPChange_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HPChange_S.Unmarshal(m, b)
}
func (m *HPChange_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HPChange_S.Marshal(b, m, deterministic)
}
func (dst *HPChange_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HPChange_S.Merge(dst, src)
}
func (m *HPChange_S) XXX_Size() int {
	return xxx_messageInfo_HPChange_S.Size(m)
}
func (m *HPChange_S) XXX_DiscardUnknown() {
	xxx_messageInfo_HPChange_S.DiscardUnknown(m)
}

var xxx_messageInfo_HPChange_S proto.InternalMessageInfo

func (m *HPChange_S) GetGuid() int64 {
	if m != nil && m.Guid != nil {
		return *m.Guid
	}
	return 0
}

func (m *HPChange_S) GetCurHp() int64 {
	if m != nil && m.CurHp != nil {
		return *m.CurHp
	}
	return 0
}

type AttrsChange_C struct {
	Guid                 *int64   `protobuf:"varint,1,opt,name=guid" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttrsChange_C) Reset()         { *m = AttrsChange_C{} }
func (m *AttrsChange_C) String() string { return proto.CompactTextString(m) }
func (*AttrsChange_C) ProtoMessage()    {}
func (*AttrsChange_C) Descriptor() ([]byte, []int) {
	return fileDescriptor_AttrsChange_e40cb90330404d63, []int{2}
}
func (m *AttrsChange_C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttrsChange_C.Unmarshal(m, b)
}
func (m *AttrsChange_C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttrsChange_C.Marshal(b, m, deterministic)
}
func (dst *AttrsChange_C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttrsChange_C.Merge(dst, src)
}
func (m *AttrsChange_C) XXX_Size() int {
	return xxx_messageInfo_AttrsChange_C.Size(m)
}
func (m *AttrsChange_C) XXX_DiscardUnknown() {
	xxx_messageInfo_AttrsChange_C.DiscardUnknown(m)
}

var xxx_messageInfo_AttrsChange_C proto.InternalMessageInfo

func (m *AttrsChange_C) GetGuid() int64 {
	if m != nil && m.Guid != nil {
		return *m.Guid
	}
	return 0
}

type AttrsChange_S struct {
	AttrInfos            []*NetAttrInfos `protobuf:"bytes,1,rep,name=attrInfos" json:"attrInfos,omitempty"`
	Guid                 *int64          `protobuf:"varint,2,opt,name=guid" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AttrsChange_S) Reset()         { *m = AttrsChange_S{} }
func (m *AttrsChange_S) String() string { return proto.CompactTextString(m) }
func (*AttrsChange_S) ProtoMessage()    {}
func (*AttrsChange_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_AttrsChange_e40cb90330404d63, []int{3}
}
func (m *AttrsChange_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttrsChange_S.Unmarshal(m, b)
}
func (m *AttrsChange_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttrsChange_S.Marshal(b, m, deterministic)
}
func (dst *AttrsChange_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttrsChange_S.Merge(dst, src)
}
func (m *AttrsChange_S) XXX_Size() int {
	return xxx_messageInfo_AttrsChange_S.Size(m)
}
func (m *AttrsChange_S) XXX_DiscardUnknown() {
	xxx_messageInfo_AttrsChange_S.DiscardUnknown(m)
}

var xxx_messageInfo_AttrsChange_S proto.InternalMessageInfo

func (m *AttrsChange_S) GetAttrInfos() []*NetAttrInfos {
	if m != nil {
		return m.AttrInfos
	}
	return nil
}

func (m *AttrsChange_S) GetGuid() int64 {
	if m != nil && m.Guid != nil {
		return *m.Guid
	}
	return 0
}

type ExpChange_S struct {
	Type                 *int32   `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Guid                 *int64   `protobuf:"varint,2,opt,name=guid" json:"guid,omitempty"`
	Level                *int32   `protobuf:"varint,3,opt,name=level" json:"level,omitempty"`
	Exp                  *int64   `protobuf:"varint,4,opt,name=exp" json:"exp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpChange_S) Reset()         { *m = ExpChange_S{} }
func (m *ExpChange_S) String() string { return proto.CompactTextString(m) }
func (*ExpChange_S) ProtoMessage()    {}
func (*ExpChange_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_AttrsChange_e40cb90330404d63, []int{4}
}
func (m *ExpChange_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpChange_S.Unmarshal(m, b)
}
func (m *ExpChange_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpChange_S.Marshal(b, m, deterministic)
}
func (dst *ExpChange_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpChange_S.Merge(dst, src)
}
func (m *ExpChange_S) XXX_Size() int {
	return xxx_messageInfo_ExpChange_S.Size(m)
}
func (m *ExpChange_S) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpChange_S.DiscardUnknown(m)
}

var xxx_messageInfo_ExpChange_S proto.InternalMessageInfo

func (m *ExpChange_S) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *ExpChange_S) GetGuid() int64 {
	if m != nil && m.Guid != nil {
		return *m.Guid
	}
	return 0
}

func (m *ExpChange_S) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *ExpChange_S) GetExp() int64 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func init() {
	proto.RegisterType((*NumberChange_S)(nil), "message.NumberChange_S")
	proto.RegisterType((*HPChange_S)(nil), "message.HPChange_S")
	proto.RegisterType((*AttrsChange_C)(nil), "message.AttrsChange_C")
	proto.RegisterType((*AttrsChange_S)(nil), "message.AttrsChange_S")
	proto.RegisterType((*ExpChange_S)(nil), "message.ExpChange_S")
}

func init() { proto.RegisterFile("AttrsChange.proto", fileDescriptor_AttrsChange_e40cb90330404d63) }

var fileDescriptor_AttrsChange_e40cb90330404d63 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x4f, 0xcf, 0x4b, 0xc3, 0x30,
	0x18, 0xa5, 0xcb, 0x86, 0xec, 0xab, 0x8a, 0x06, 0x85, 0xb0, 0x53, 0xa9, 0x97, 0x9e, 0x7a, 0x50,
	0xf0, 0xe4, 0x65, 0x1d, 0xc2, 0x76, 0x99, 0x92, 0x81, 0x78, 0x91, 0x11, 0xe7, 0xb7, 0x2a, 0xb4,
	0x6b, 0xc8, 0x8f, 0x31, 0xff, 0x7b, 0x69, 0x62, 0x67, 0x84, 0xde, 0xde, 0xfb, 0xf2, 0xde, 0xcb,
	0x7b, 0x70, 0x39, 0x35, 0x46, 0xe9, 0xd9, 0xa7, 0xd8, 0x95, 0x98, 0x4b, 0xd5, 0x98, 0x86, 0x9e,
	0xd4, 0xa8, 0xb5, 0x28, 0x71, 0x72, 0x5a, 0x08, 0x8d, 0x2f, 0x4f, 0xfe, 0x9c, 0x3e, 0xc0, 0xf9,
	0xd2, 0xd6, 0xef, 0xa8, 0xbc, 0x78, 0xbd, 0xa2, 0x57, 0x30, 0xda, 0x6c, 0xcb, 0xc5, 0x07, 0x8b,
	0x92, 0x28, 0x1b, 0x73, 0x4f, 0xda, 0xeb, 0x5e, 0x54, 0x16, 0xd9, 0x20, 0x89, 0x32, 0xc2, 0x3d,
	0x49, 0xef, 0x01, 0xe6, 0xcf, 0x47, 0x27, 0x85, 0x61, 0x69, 0xbf, 0xbc, 0x91, 0x70, 0x87, 0x5d,
	0x9a, 0x55, 0x73, 0xd9, 0xf9, 0x1c, 0x49, 0x6f, 0xe0, 0x2c, 0x68, 0xb8, 0x9e, 0xf5, 0x59, 0xd3,
	0xd7, 0xff, 0xa2, 0x15, 0xbd, 0x83, 0xb1, 0x30, 0x46, 0x2d, 0x76, 0xdb, 0x46, 0xb3, 0x28, 0x21,
	0x59, 0x7c, 0x7b, 0x9d, 0xff, 0xce, 0xca, 0x97, 0x68, 0xa6, 0xdd, 0x23, 0xff, 0xd3, 0x1d, 0x93,
	0x07, 0x41, 0xf2, 0x1b, 0xc4, 0x8f, 0x07, 0x19, 0xf6, 0x36, 0xdf, 0x12, 0xdd, 0xe7, 0x23, 0xee,
	0x70, 0x9f, 0xad, 0xdd, 0x52, 0xe1, 0x1e, 0x2b, 0x46, 0x9c, 0xd0, 0x13, 0x7a, 0x01, 0x04, 0x0f,
	0x92, 0x0d, 0x9d, 0xb0, 0x85, 0xc5, 0x04, 0xe2, 0x4d, 0x53, 0x77, 0xcd, 0x8a, 0x38, 0x58, 0xf1,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x4c, 0x85, 0x1d, 0x9a, 0x01, 0x00, 0x00,
}