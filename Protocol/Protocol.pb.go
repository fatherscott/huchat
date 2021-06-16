// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Protocol.proto

package Protocol

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

type HeaderType int32

const (
	HeaderType_T_LoginRequest    HeaderType = 0
	HeaderType_T_LoginResponse   HeaderType = 1
	HeaderType_T_LogoutRequest   HeaderType = 2
	HeaderType_T_LogoutResponse  HeaderType = 3
	HeaderType_T_MessageRequest  HeaderType = 4
	HeaderType_T_MessageResponse HeaderType = 5
)

var HeaderType_name = map[int32]string{
	0: "T_LoginRequest",
	1: "T_LoginResponse",
	2: "T_LogoutRequest",
	3: "T_LogoutResponse",
	4: "T_MessageRequest",
	5: "T_MessageResponse",
}

var HeaderType_value = map[string]int32{
	"T_LoginRequest":    0,
	"T_LoginResponse":   1,
	"T_LogoutRequest":   2,
	"T_LogoutResponse":  3,
	"T_MessageRequest":  4,
	"T_MessageResponse": 5,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}

func (HeaderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_071a0530b1819269, []int{0}
}

type Header struct {
	Type                 HeaderType `protobuf:"varint,1,opt,name=Type,proto3,enum=HeaderType" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_071a0530b1819269, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetType() HeaderType {
	if m != nil {
		return m.Type
	}
	return HeaderType_T_LoginRequest
}

type LoginRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	RoomId               string   `protobuf:"bytes,2,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	Level                int32    `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	NickName             string   `protobuf:"bytes,4,opt,name=NickName,proto3" json:"NickName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_071a0530b1819269, []int{1}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *LoginRequest) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *LoginRequest) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *LoginRequest) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

type LoginResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_071a0530b1819269, []int{2}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *LoginResponse) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type LogoutRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_071a0530b1819269, []int{3}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type LogoutResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	RoomId               string   `protobuf:"bytes,3,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_071a0530b1819269, []int{4}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

func (m *LogoutResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *LogoutResponse) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *LogoutResponse) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type MessageRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_071a0530b1819269, []int{5}
}

func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *MessageRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type MessageResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	RoomId               string   `protobuf:"bytes,3,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_071a0530b1819269, []int{6}
}

func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *MessageResponse) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *MessageResponse) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *MessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("HeaderType", HeaderType_name, HeaderType_value)
	proto.RegisterType((*Header)(nil), "Header")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "LogoutResponse")
	proto.RegisterType((*MessageRequest)(nil), "MessageRequest")
	proto.RegisterType((*MessageResponse)(nil), "MessageResponse")
}

func init() { proto.RegisterFile("Protocol.proto", fileDescriptor_071a0530b1819269) }

var fileDescriptor_071a0530b1819269 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xdd, 0x4a, 0xf3, 0x40,
	0x14, 0xfc, 0x92, 0x26, 0xfd, 0xec, 0x51, 0xd3, 0x78, 0x6c, 0x25, 0x88, 0x60, 0xd9, 0xab, 0x2a,
	0xd8, 0x0b, 0x7d, 0x02, 0x15, 0xc1, 0x42, 0x2d, 0xb2, 0xf4, 0x52, 0x94, 0x36, 0x3d, 0x94, 0x62,
	0x93, 0xad, 0xdd, 0x4d, 0xa1, 0x0f, 0xe1, 0x8b, 0xf8, 0x94, 0xd2, 0xcd, 0xe6, 0xcf, 0xab, 0x82,
	0x78, 0xb7, 0x33, 0x7b, 0x76, 0x66, 0x38, 0x3b, 0xe0, 0x3d, 0xaf, 0x84, 0x12, 0xa1, 0x58, 0xf4,
	0x96, 0xdb, 0x03, 0xbb, 0x80, 0xfa, 0x23, 0x8d, 0xa7, 0xb4, 0xc2, 0x73, 0x70, 0x46, 0x9b, 0x25,
	0x05, 0x56, 0xc7, 0xea, 0x7a, 0xd7, 0xfb, 0xbd, 0x94, 0xde, 0x52, 0x5c, 0x5f, 0xb0, 0x35, 0x1c,
	0x0c, 0xc4, 0x6c, 0x1e, 0x73, 0xfa, 0x48, 0x48, 0x2a, 0x3c, 0x83, 0xc6, 0x6d, 0x18, 0x8a, 0x24,
	0x56, 0xfd, 0xa9, 0x7e, 0xd5, 0xe0, 0x05, 0x81, 0x27, 0x50, 0xe7, 0x42, 0x44, 0xfd, 0x69, 0x60,
	0xeb, 0x2b, 0x83, 0xb0, 0x05, 0xee, 0x80, 0xd6, 0xb4, 0x08, 0x6a, 0x1d, 0xab, 0xeb, 0xf2, 0x14,
	0xe0, 0x29, 0xec, 0x0d, 0xe7, 0xe1, 0xfb, 0x70, 0x1c, 0x51, 0xe0, 0xe8, 0xf9, 0x1c, 0xb3, 0x07,
	0x38, 0x34, 0xbe, 0x72, 0x29, 0x62, 0x49, 0x5a, 0x9a, 0x64, 0xb2, 0x50, 0xda, 0xd5, 0xe5, 0x06,
	0x55, 0x03, 0xd9, 0x3f, 0x02, 0xb1, 0x2b, 0x2d, 0x23, 0x12, 0xb5, 0x53, 0x7e, 0xf6, 0x0a, 0x5e,
	0x36, 0xfe, 0x1b, 0xdb, 0xd2, 0x1e, 0x6a, 0xe5, 0x3d, 0xb0, 0x17, 0xf0, 0x9e, 0x48, 0xca, 0xf1,
	0x8c, 0x76, 0xdb, 0x27, 0x82, 0xa3, 0xb6, 0xdf, 0x63, 0x6b, 0x6f, 0x7d, 0xc6, 0x00, 0xfe, 0x1b,
	0x0d, 0x23, 0x9e, 0x41, 0xb6, 0x81, 0x66, 0xae, 0xfe, 0x17, 0xf1, 0xcb, 0xd6, 0x4e, 0xc5, 0xfa,
	0xf2, 0xd3, 0x02, 0x28, 0xba, 0x83, 0x08, 0xde, 0xe8, 0xad, 0xdc, 0x1b, 0xff, 0x1f, 0x1e, 0x43,
	0x33, 0xe7, 0xd2, 0x74, 0xbe, 0x95, 0x93, 0xc5, 0x0f, 0xf9, 0x36, 0xb6, 0xc0, 0x2f, 0x48, 0x33,
	0x5a, 0x4b, 0xd9, 0xea, 0xf6, 0x7c, 0x07, 0xdb, 0x70, 0x54, 0x62, 0xcd, 0xb0, 0x7b, 0xd7, 0xfe,
	0xb2, 0xf1, 0x5e, 0x44, 0x91, 0x88, 0x07, 0xf3, 0x49, 0x2f, 0xab, 0xff, 0xa4, 0xae, 0xfb, 0x7f,
	0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x43, 0x25, 0xa9, 0x94, 0x11, 0x03, 0x00, 0x00,
}