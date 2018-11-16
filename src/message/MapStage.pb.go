// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MapStage.proto

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

type MapPreEnter_C struct {
	Id                   *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapPreEnter_C) Reset()         { *m = MapPreEnter_C{} }
func (m *MapPreEnter_C) String() string { return proto.CompactTextString(m) }
func (*MapPreEnter_C) ProtoMessage()    {}
func (*MapPreEnter_C) Descriptor() ([]byte, []int) {
	return fileDescriptor_MapStage_25e3601dd2b1b942, []int{0}
}
func (m *MapPreEnter_C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapPreEnter_C.Unmarshal(m, b)
}
func (m *MapPreEnter_C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapPreEnter_C.Marshal(b, m, deterministic)
}
func (dst *MapPreEnter_C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapPreEnter_C.Merge(dst, src)
}
func (m *MapPreEnter_C) XXX_Size() int {
	return xxx_messageInfo_MapPreEnter_C.Size(m)
}
func (m *MapPreEnter_C) XXX_DiscardUnknown() {
	xxx_messageInfo_MapPreEnter_C.DiscardUnknown(m)
}

var xxx_messageInfo_MapPreEnter_C proto.InternalMessageInfo

func (m *MapPreEnter_C) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type MapPreEnter_S struct {
	Id                   *string  `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapPreEnter_S) Reset()         { *m = MapPreEnter_S{} }
func (m *MapPreEnter_S) String() string { return proto.CompactTextString(m) }
func (*MapPreEnter_S) ProtoMessage()    {}
func (*MapPreEnter_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_MapStage_25e3601dd2b1b942, []int{1}
}
func (m *MapPreEnter_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapPreEnter_S.Unmarshal(m, b)
}
func (m *MapPreEnter_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapPreEnter_S.Marshal(b, m, deterministic)
}
func (dst *MapPreEnter_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapPreEnter_S.Merge(dst, src)
}
func (m *MapPreEnter_S) XXX_Size() int {
	return xxx_messageInfo_MapPreEnter_S.Size(m)
}
func (m *MapPreEnter_S) XXX_DiscardUnknown() {
	xxx_messageInfo_MapPreEnter_S.DiscardUnknown(m)
}

var xxx_messageInfo_MapPreEnter_S proto.InternalMessageInfo

func (m *MapPreEnter_S) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type MapEnter_C struct {
	Id                   *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapEnter_C) Reset()         { *m = MapEnter_C{} }
func (m *MapEnter_C) String() string { return proto.CompactTextString(m) }
func (*MapEnter_C) ProtoMessage()    {}
func (*MapEnter_C) Descriptor() ([]byte, []int) {
	return fileDescriptor_MapStage_25e3601dd2b1b942, []int{2}
}
func (m *MapEnter_C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapEnter_C.Unmarshal(m, b)
}
func (m *MapEnter_C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapEnter_C.Marshal(b, m, deterministic)
}
func (dst *MapEnter_C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapEnter_C.Merge(dst, src)
}
func (m *MapEnter_C) XXX_Size() int {
	return xxx_messageInfo_MapEnter_C.Size(m)
}
func (m *MapEnter_C) XXX_DiscardUnknown() {
	xxx_messageInfo_MapEnter_C.DiscardUnknown(m)
}

var xxx_messageInfo_MapEnter_C proto.InternalMessageInfo

func (m *MapEnter_C) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type MapEnter_S struct {
	Id                   *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapEnter_S) Reset()         { *m = MapEnter_S{} }
func (m *MapEnter_S) String() string { return proto.CompactTextString(m) }
func (*MapEnter_S) ProtoMessage()    {}
func (*MapEnter_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_MapStage_25e3601dd2b1b942, []int{3}
}
func (m *MapEnter_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapEnter_S.Unmarshal(m, b)
}
func (m *MapEnter_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapEnter_S.Marshal(b, m, deterministic)
}
func (dst *MapEnter_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapEnter_S.Merge(dst, src)
}
func (m *MapEnter_S) XXX_Size() int {
	return xxx_messageInfo_MapEnter_S.Size(m)
}
func (m *MapEnter_S) XXX_DiscardUnknown() {
	xxx_messageInfo_MapEnter_S.DiscardUnknown(m)
}

var xxx_messageInfo_MapEnter_S proto.InternalMessageInfo

func (m *MapEnter_S) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type ElementEnter_S struct {
	Elements             []*NetStageElementVO `protobuf:"bytes,1,rep,name=elements" json:"elements,omitempty"`
	IsSelf               *bool                `protobuf:"varint,2,opt,name=isSelf" json:"isSelf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ElementEnter_S) Reset()         { *m = ElementEnter_S{} }
func (m *ElementEnter_S) String() string { return proto.CompactTextString(m) }
func (*ElementEnter_S) ProtoMessage()    {}
func (*ElementEnter_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_MapStage_25e3601dd2b1b942, []int{4}
}
func (m *ElementEnter_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElementEnter_S.Unmarshal(m, b)
}
func (m *ElementEnter_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElementEnter_S.Marshal(b, m, deterministic)
}
func (dst *ElementEnter_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElementEnter_S.Merge(dst, src)
}
func (m *ElementEnter_S) XXX_Size() int {
	return xxx_messageInfo_ElementEnter_S.Size(m)
}
func (m *ElementEnter_S) XXX_DiscardUnknown() {
	xxx_messageInfo_ElementEnter_S.DiscardUnknown(m)
}

var xxx_messageInfo_ElementEnter_S proto.InternalMessageInfo

func (m *ElementEnter_S) GetElements() []*NetStageElementVO {
	if m != nil {
		return m.Elements
	}
	return nil
}

func (m *ElementEnter_S) GetIsSelf() bool {
	if m != nil && m.IsSelf != nil {
		return *m.IsSelf
	}
	return false
}

type SceneRoleEnter_S struct {
	Elements             []*NetStageElementVO `protobuf:"bytes,1,rep,name=elements" json:"elements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SceneRoleEnter_S) Reset()         { *m = SceneRoleEnter_S{} }
func (m *SceneRoleEnter_S) String() string { return proto.CompactTextString(m) }
func (*SceneRoleEnter_S) ProtoMessage()    {}
func (*SceneRoleEnter_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_MapStage_25e3601dd2b1b942, []int{5}
}
func (m *SceneRoleEnter_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SceneRoleEnter_S.Unmarshal(m, b)
}
func (m *SceneRoleEnter_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SceneRoleEnter_S.Marshal(b, m, deterministic)
}
func (dst *SceneRoleEnter_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SceneRoleEnter_S.Merge(dst, src)
}
func (m *SceneRoleEnter_S) XXX_Size() int {
	return xxx_messageInfo_SceneRoleEnter_S.Size(m)
}
func (m *SceneRoleEnter_S) XXX_DiscardUnknown() {
	xxx_messageInfo_SceneRoleEnter_S.DiscardUnknown(m)
}

var xxx_messageInfo_SceneRoleEnter_S proto.InternalMessageInfo

func (m *SceneRoleEnter_S) GetElements() []*NetStageElementVO {
	if m != nil {
		return m.Elements
	}
	return nil
}

type ElementRemove_S struct {
	Guid                 *int64   `protobuf:"varint,1,opt,name=guid" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElementRemove_S) Reset()         { *m = ElementRemove_S{} }
func (m *ElementRemove_S) String() string { return proto.CompactTextString(m) }
func (*ElementRemove_S) ProtoMessage()    {}
func (*ElementRemove_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_MapStage_25e3601dd2b1b942, []int{6}
}
func (m *ElementRemove_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElementRemove_S.Unmarshal(m, b)
}
func (m *ElementRemove_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElementRemove_S.Marshal(b, m, deterministic)
}
func (dst *ElementRemove_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElementRemove_S.Merge(dst, src)
}
func (m *ElementRemove_S) XXX_Size() int {
	return xxx_messageInfo_ElementRemove_S.Size(m)
}
func (m *ElementRemove_S) XXX_DiscardUnknown() {
	xxx_messageInfo_ElementRemove_S.DiscardUnknown(m)
}

var xxx_messageInfo_ElementRemove_S proto.InternalMessageInfo

func (m *ElementRemove_S) GetGuid() int64 {
	if m != nil && m.Guid != nil {
		return *m.Guid
	}
	return 0
}

type FightDead_S struct {
	Type                 *int32   `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	DeadTargetId         *int64   `protobuf:"varint,2,opt,name=deadTargetId" json:"deadTargetId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FightDead_S) Reset()         { *m = FightDead_S{} }
func (m *FightDead_S) String() string { return proto.CompactTextString(m) }
func (*FightDead_S) ProtoMessage()    {}
func (*FightDead_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_MapStage_25e3601dd2b1b942, []int{7}
}
func (m *FightDead_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FightDead_S.Unmarshal(m, b)
}
func (m *FightDead_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FightDead_S.Marshal(b, m, deterministic)
}
func (dst *FightDead_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FightDead_S.Merge(dst, src)
}
func (m *FightDead_S) XXX_Size() int {
	return xxx_messageInfo_FightDead_S.Size(m)
}
func (m *FightDead_S) XXX_DiscardUnknown() {
	xxx_messageInfo_FightDead_S.DiscardUnknown(m)
}

var xxx_messageInfo_FightDead_S proto.InternalMessageInfo

func (m *FightDead_S) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *FightDead_S) GetDeadTargetId() int64 {
	if m != nil && m.DeadTargetId != nil {
		return *m.DeadTargetId
	}
	return 0
}

type LeaderChange_S struct {
	LeaderGuid           *int64   `protobuf:"varint,1,opt,name=leaderGuid" json:"leaderGuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaderChange_S) Reset()         { *m = LeaderChange_S{} }
func (m *LeaderChange_S) String() string { return proto.CompactTextString(m) }
func (*LeaderChange_S) ProtoMessage()    {}
func (*LeaderChange_S) Descriptor() ([]byte, []int) {
	return fileDescriptor_MapStage_25e3601dd2b1b942, []int{8}
}
func (m *LeaderChange_S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaderChange_S.Unmarshal(m, b)
}
func (m *LeaderChange_S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaderChange_S.Marshal(b, m, deterministic)
}
func (dst *LeaderChange_S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaderChange_S.Merge(dst, src)
}
func (m *LeaderChange_S) XXX_Size() int {
	return xxx_messageInfo_LeaderChange_S.Size(m)
}
func (m *LeaderChange_S) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaderChange_S.DiscardUnknown(m)
}

var xxx_messageInfo_LeaderChange_S proto.InternalMessageInfo

func (m *LeaderChange_S) GetLeaderGuid() int64 {
	if m != nil && m.LeaderGuid != nil {
		return *m.LeaderGuid
	}
	return 0
}

func init() {
	proto.RegisterType((*MapPreEnter_C)(nil), "message.MapPreEnter_C")
	proto.RegisterType((*MapPreEnter_S)(nil), "message.MapPreEnter_S")
	proto.RegisterType((*MapEnter_C)(nil), "message.MapEnter_C")
	proto.RegisterType((*MapEnter_S)(nil), "message.MapEnter_S")
	proto.RegisterType((*ElementEnter_S)(nil), "message.ElementEnter_S")
	proto.RegisterType((*SceneRoleEnter_S)(nil), "message.SceneRoleEnter_S")
	proto.RegisterType((*ElementRemove_S)(nil), "message.ElementRemove_S")
	proto.RegisterType((*FightDead_S)(nil), "message.FightDead_S")
	proto.RegisterType((*LeaderChange_S)(nil), "message.LeaderChange_S")
}

func init() { proto.RegisterFile("MapStage.proto", fileDescriptor_MapStage_25e3601dd2b1b942) }

var fileDescriptor_MapStage_25e3601dd2b1b942 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0x87, 0x49, 0xf2, 0xbe, 0x1a, 0x27, 0x35, 0xca, 0x1e, 0xb4, 0x14, 0xd1, 0xb2, 0x20, 0xf4,
	0x14, 0xc4, 0x83, 0x1f, 0x20, 0x35, 0x8a, 0x62, 0xac, 0xec, 0x4a, 0xaf, 0x75, 0xe9, 0x8e, 0xdb,
	0x40, 0xfe, 0x91, 0xac, 0x82, 0xdf, 0x5e, 0xdc, 0x6c, 0x6a, 0x55, 0x3c, 0x79, 0x9b, 0x3c, 0xf3,
	0x0c, 0x33, 0xbf, 0xb0, 0x10, 0xa6, 0xa2, 0xe6, 0x5a, 0x28, 0x8c, 0xea, 0xa6, 0xd2, 0x15, 0xd9,
	0x2e, 0xb0, 0x6d, 0x85, 0xc2, 0xd1, 0x20, 0x16, 0x2d, 0xce, 0x67, 0x1d, 0xa6, 0x27, 0xb0, 0x9b,
	0x8a, 0xfa, 0xa1, 0xc1, 0xa4, 0xd4, 0xd8, 0x2c, 0xa6, 0x24, 0x04, 0x37, 0x93, 0x43, 0x67, 0xec,
	0x4c, 0x76, 0x98, 0x9b, 0xc9, 0xef, 0x02, 0xb7, 0x82, 0xbb, 0x16, 0x8e, 0x00, 0x52, 0x51, 0xff,
	0x36, 0xbe, 0xd9, 0xe5, 0x3f, 0xba, 0x4f, 0x10, 0x26, 0x39, 0x16, 0x58, 0xea, 0xde, 0xb8, 0x00,
	0x1f, 0x3b, 0xd2, 0x0e, 0x9d, 0xb1, 0x37, 0x09, 0xce, 0x47, 0x91, 0xbd, 0x3c, 0xba, 0x47, 0x6d,
	0x12, 0xd9, 0x91, 0xf9, 0x8c, 0xad, 0x5d, 0x72, 0x00, 0x5b, 0x59, 0xcb, 0x31, 0x7f, 0x36, 0x97,
	0xf9, 0xcc, 0x7e, 0xd1, 0x5b, 0xd8, 0xe7, 0x4b, 0x2c, 0x91, 0x55, 0x39, 0xfe, 0x71, 0x07, 0x3d,
	0x85, 0x3d, 0x8b, 0x19, 0x16, 0xd5, 0x2b, 0x2e, 0x38, 0x21, 0xf0, 0x4f, 0xbd, 0xd8, 0x48, 0x1e,
	0x33, 0x35, 0x4d, 0x20, 0xb8, 0xca, 0xd4, 0x4a, 0x5f, 0xa2, 0x90, 0x9d, 0xa2, 0xdf, 0x6a, 0x34,
	0xca, 0x7f, 0x66, 0x6a, 0x42, 0x61, 0x20, 0x51, 0xc8, 0x47, 0xd1, 0x28, 0xd4, 0x37, 0xdd, 0xdf,
	0xf4, 0xd8, 0x17, 0x46, 0xcf, 0x20, 0xbc, 0x43, 0x21, 0xb1, 0x99, 0xae, 0x44, 0xa9, 0x3e, 0x96,
	0x1d, 0x03, 0xe4, 0x86, 0x5c, 0x7f, 0xae, 0xdc, 0x20, 0xf1, 0x21, 0x04, 0xcb, 0xaa, 0xe8, 0xa3,
	0xc4, 0x7e, 0xff, 0x02, 0xde, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x62, 0xff, 0xa0, 0x0c, 0x02,
	0x00, 0x00,
}
