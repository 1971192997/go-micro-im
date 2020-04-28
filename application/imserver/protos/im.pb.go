// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im.proto

package im

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

type PublishMessageRequest struct {
	FromToken            string   `protobuf:"bytes,1,opt,name=fromToken,proto3" json:"fromToken,omitempty"`
	ToToken              string   `protobuf:"bytes,2,opt,name=toToken,proto3" json:"toToken,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	ServerName           string   `protobuf:"bytes,4,opt,name=serverName,proto3" json:"serverName,omitempty"`
	Topic                string   `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Address              string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishMessageRequest) Reset()         { *m = PublishMessageRequest{} }
func (m *PublishMessageRequest) String() string { return proto.CompactTextString(m) }
func (*PublishMessageRequest) ProtoMessage()    {}
func (*PublishMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f2114a3e4ddb9e, []int{0}
}

func (m *PublishMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishMessageRequest.Unmarshal(m, b)
}
func (m *PublishMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishMessageRequest.Marshal(b, m, deterministic)
}
func (m *PublishMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishMessageRequest.Merge(m, src)
}
func (m *PublishMessageRequest) XXX_Size() int {
	return xxx_messageInfo_PublishMessageRequest.Size(m)
}
func (m *PublishMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishMessageRequest proto.InternalMessageInfo

func (m *PublishMessageRequest) GetFromToken() string {
	if m != nil {
		return m.FromToken
	}
	return ""
}

func (m *PublishMessageRequest) GetToToken() string {
	if m != nil {
		return m.ToToken
	}
	return ""
}

func (m *PublishMessageRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *PublishMessageRequest) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *PublishMessageRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishMessageRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type PublishMessageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishMessageResponse) Reset()         { *m = PublishMessageResponse{} }
func (m *PublishMessageResponse) String() string { return proto.CompactTextString(m) }
func (*PublishMessageResponse) ProtoMessage()    {}
func (*PublishMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f2114a3e4ddb9e, []int{1}
}

func (m *PublishMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishMessageResponse.Unmarshal(m, b)
}
func (m *PublishMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishMessageResponse.Marshal(b, m, deterministic)
}
func (m *PublishMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishMessageResponse.Merge(m, src)
}
func (m *PublishMessageResponse) XXX_Size() int {
	return xxx_messageInfo_PublishMessageResponse.Size(m)
}
func (m *PublishMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishMessageResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PublishMessageRequest)(nil), "PublishMessageRequest")
	proto.RegisterType((*PublishMessageResponse)(nil), "PublishMessageResponse")
}

func init() { proto.RegisterFile("im.proto", fileDescriptor_36f2114a3e4ddb9e) }

var fileDescriptor_36f2114a3e4ddb9e = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0xcc, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0xda, 0xcc, 0xc8, 0x25, 0x1a, 0x50, 0x9a, 0x94, 0x93, 0x59, 0x9c,
	0xe1, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x1a, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24,
	0xc3, 0xc5, 0x99, 0x56, 0x94, 0x9f, 0x1b, 0x92, 0x9f, 0x9d, 0x9a, 0x27, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0x84, 0x10, 0x10, 0x92, 0xe0, 0x62, 0x2f, 0xc9, 0x87, 0xc8, 0x31, 0x81, 0xe5, 0x60,
	0x5c, 0x21, 0x21, 0x2e, 0x96, 0xa4, 0xfc, 0x94, 0x4a, 0x09, 0x66, 0xb0, 0x30, 0x98, 0x2d, 0x24,
	0xc7, 0xc5, 0x55, 0x9c, 0x5a, 0x54, 0x96, 0x5a, 0xe4, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x02, 0x96,
	0x41, 0x12, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2f, 0xc8, 0x4c, 0x96, 0x60, 0x05, 0x4b, 0x41,
	0x38, 0x20, 0x3b, 0x12, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0xd8, 0x20, 0x76, 0x40, 0xb9,
	0x4a, 0x12, 0x5c, 0x62, 0xe8, 0x8e, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x35, 0xf2, 0xe4, 0x62,
	0xf2, 0xcc, 0x15, 0x72, 0xe6, 0xe2, 0x43, 0x95, 0x17, 0x12, 0xd3, 0xc3, 0xea, 0x4b, 0x29, 0x71,
	0x3d, 0xec, 0x06, 0x29, 0x31, 0x24, 0xb1, 0x81, 0x43, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xf9, 0xf1, 0x92, 0x53, 0x2d, 0x01, 0x00, 0x00,
}
