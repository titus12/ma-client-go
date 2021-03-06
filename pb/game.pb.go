// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package proto

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

// Ping 服务器的消息
type PingC2S struct {
	Time                 uint32   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingC2S) Reset()         { *m = PingC2S{} }
func (m *PingC2S) String() string { return proto.CompactTextString(m) }
func (*PingC2S) ProtoMessage()    {}
func (*PingC2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *PingC2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingC2S.Unmarshal(m, b)
}
func (m *PingC2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingC2S.Marshal(b, m, deterministic)
}
func (m *PingC2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingC2S.Merge(m, src)
}
func (m *PingC2S) XXX_Size() int {
	return xxx_messageInfo_PingC2S.Size(m)
}
func (m *PingC2S) XXX_DiscardUnknown() {
	xxx_messageInfo_PingC2S.DiscardUnknown(m)
}

var xxx_messageInfo_PingC2S proto.InternalMessageInfo

func (m *PingC2S) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type PingS2C struct {
	Time                 uint32   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingS2C) Reset()         { *m = PingS2C{} }
func (m *PingS2C) String() string { return proto.CompactTextString(m) }
func (*PingS2C) ProtoMessage()    {}
func (*PingS2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *PingS2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingS2C.Unmarshal(m, b)
}
func (m *PingS2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingS2C.Marshal(b, m, deterministic)
}
func (m *PingS2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingS2C.Merge(m, src)
}
func (m *PingS2C) XXX_Size() int {
	return xxx_messageInfo_PingS2C.Size(m)
}
func (m *PingS2C) XXX_DiscardUnknown() {
	xxx_messageInfo_PingS2C.DiscardUnknown(m)
}

var xxx_messageInfo_PingS2C proto.InternalMessageInfo

func (m *PingS2C) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type LoginC2S struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Device               string   `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	Provider             string   `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	Token                string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginC2S) Reset()         { *m = LoginC2S{} }
func (m *LoginC2S) String() string { return proto.CompactTextString(m) }
func (*LoginC2S) ProtoMessage()    {}
func (*LoginC2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2}
}

func (m *LoginC2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginC2S.Unmarshal(m, b)
}
func (m *LoginC2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginC2S.Marshal(b, m, deterministic)
}
func (m *LoginC2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginC2S.Merge(m, src)
}
func (m *LoginC2S) XXX_Size() int {
	return xxx_messageInfo_LoginC2S.Size(m)
}
func (m *LoginC2S) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginC2S.DiscardUnknown(m)
}

var xxx_messageInfo_LoginC2S proto.InternalMessageInfo

func (m *LoginC2S) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *LoginC2S) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LoginC2S) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *LoginC2S) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *LoginC2S) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LoginS2C struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginS2C) Reset()         { *m = LoginS2C{} }
func (m *LoginS2C) String() string { return proto.CompactTextString(m) }
func (*LoginS2C) ProtoMessage()    {}
func (*LoginS2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{3}
}

func (m *LoginS2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginS2C.Unmarshal(m, b)
}
func (m *LoginS2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginS2C.Marshal(b, m, deterministic)
}
func (m *LoginS2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginS2C.Merge(m, src)
}
func (m *LoginS2C) XXX_Size() int {
	return xxx_messageInfo_LoginS2C.Size(m)
}
func (m *LoginS2C) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginS2C.DiscardUnknown(m)
}

var xxx_messageInfo_LoginS2C proto.InternalMessageInfo

func (m *LoginS2C) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type IntoGameC2S struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntoGameC2S) Reset()         { *m = IntoGameC2S{} }
func (m *IntoGameC2S) String() string { return proto.CompactTextString(m) }
func (*IntoGameC2S) ProtoMessage()    {}
func (*IntoGameC2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4}
}

func (m *IntoGameC2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntoGameC2S.Unmarshal(m, b)
}
func (m *IntoGameC2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntoGameC2S.Marshal(b, m, deterministic)
}
func (m *IntoGameC2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntoGameC2S.Merge(m, src)
}
func (m *IntoGameC2S) XXX_Size() int {
	return xxx_messageInfo_IntoGameC2S.Size(m)
}
func (m *IntoGameC2S) XXX_DiscardUnknown() {
	xxx_messageInfo_IntoGameC2S.DiscardUnknown(m)
}

var xxx_messageInfo_IntoGameC2S proto.InternalMessageInfo

func (m *IntoGameC2S) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type IntoGameS2C struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntoGameS2C) Reset()         { *m = IntoGameS2C{} }
func (m *IntoGameS2C) String() string { return proto.CompactTextString(m) }
func (*IntoGameS2C) ProtoMessage()    {}
func (*IntoGameS2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{5}
}

func (m *IntoGameS2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntoGameS2C.Unmarshal(m, b)
}
func (m *IntoGameS2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntoGameS2C.Marshal(b, m, deterministic)
}
func (m *IntoGameS2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntoGameS2C.Merge(m, src)
}
func (m *IntoGameS2C) XXX_Size() int {
	return xxx_messageInfo_IntoGameS2C.Size(m)
}
func (m *IntoGameS2C) XXX_DiscardUnknown() {
	xxx_messageInfo_IntoGameS2C.DiscardUnknown(m)
}

var xxx_messageInfo_IntoGameS2C proto.InternalMessageInfo

func (m *IntoGameS2C) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*PingC2S)(nil), "proto.PingC2S")
	proto.RegisterType((*PingS2C)(nil), "proto.PingS2C")
	proto.RegisterType((*LoginC2S)(nil), "proto.LoginC2S")
	proto.RegisterType((*LoginS2C)(nil), "proto.LoginS2C")
	proto.RegisterType((*IntoGameC2S)(nil), "proto.IntoGameC2S")
	proto.RegisterType((*IntoGameS2C)(nil), "proto.IntoGameS2C")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769) }

var fileDescriptor_38fc58335341d769 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0x87, 0x30,
	0x14, 0x80, 0x59, 0x3f, 0xff, 0x75, 0xaa, 0x9b, 0x11, 0x35, 0x84, 0x40, 0x0c, 0xc2, 0xab, 0x2e,
	0xec, 0x11, 0xbc, 0x08, 0xa1, 0x8b, 0xd0, 0x07, 0x08, 0x73, 0x07, 0x19, 0xb1, 0x1d, 0x99, 0xcb,
	0x57, 0xe8, 0xb5, 0xc3, 0xa9, 0xd1, 0x45, 0x79, 0xb5, 0x7d, 0xfb, 0xce, 0xc6, 0xc7, 0x00, 0x86,
	0x4e, 0xe3, 0xe3, 0x68, 0xc9, 0x11, 0x0f, 0xfd, 0x92, 0x5e, 0xf6, 0xa4, 0x35, 0x99, 0xf5, 0x30,
	0xbf, 0x83, 0xf8, 0x55, 0x99, 0xa1, 0x2a, 0x5b, 0xce, 0x21, 0x70, 0x4a, 0xa3, 0x60, 0x19, 0x2b,
	0xae, 0x1a, 0xbf, 0xdf, 0x75, 0x5b, 0x56, 0x7f, 0xea, 0x2f, 0x06, 0xc9, 0x0b, 0x0d, 0xca, 0x2c,
	0xf7, 0x6f, 0x21, 0xfe, 0x9c, 0xd0, 0xbe, 0x29, 0xe9, 0x67, 0x4e, 0x4d, 0xb4, 0x60, 0x2d, 0xb9,
	0x80, 0x78, 0x46, 0x3b, 0x29, 0x32, 0xe2, 0x2c, 0x63, 0xc5, 0x79, 0xb3, 0x23, 0xbf, 0x81, 0x48,
	0xe2, 0xac, 0x7a, 0x14, 0x27, 0x2f, 0x36, 0xe2, 0x29, 0x24, 0xa3, 0xa5, 0x59, 0x49, 0xb4, 0x22,
	0xf0, 0xe6, 0x87, 0xf9, 0x35, 0x84, 0x8e, 0x3e, 0xd0, 0x88, 0xd0, 0x8b, 0x15, 0xf2, 0xfb, 0x2d,
	0x64, 0x29, 0xfd, 0x2f, 0x24, 0x7f, 0x80, 0x8b, 0xda, 0x38, 0x7a, 0xee, 0x34, 0x1e, 0x05, 0xff,
	0x9e, 0x3b, 0x7a, 0xef, 0x3d, 0xf2, 0x7f, 0xf8, 0xf4, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x06,
	0x33, 0x10, 0x66, 0x01, 0x00, 0x00,
}
