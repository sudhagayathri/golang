// Code generated by protoc-gen-go. DO NOT EDIT.
// source: personproto.proto

package main

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

type Personproto struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Personproto) Reset()         { *m = Personproto{} }
func (m *Personproto) String() string { return proto.CompactTextString(m) }
func (*Personproto) ProtoMessage()    {}
func (*Personproto) Descriptor() ([]byte, []int) {
	return fileDescriptor_93c3eee4dd21433b, []int{0}
}

func (m *Personproto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Personproto.Unmarshal(m, b)
}
func (m *Personproto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Personproto.Marshal(b, m, deterministic)
}
func (m *Personproto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Personproto.Merge(m, src)
}
func (m *Personproto) XXX_Size() int {
	return xxx_messageInfo_Personproto.Size(m)
}
func (m *Personproto) XXX_DiscardUnknown() {
	xxx_messageInfo_Personproto.DiscardUnknown(m)
}

var xxx_messageInfo_Personproto proto.InternalMessageInfo

func (m *Personproto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Personproto) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*Personproto)(nil), "main.Personproto")
}

func init() { proto.RegisterFile("personproto.proto", fileDescriptor_93c3eee4dd21433b) }

var fileDescriptor_93c3eee4dd21433b = []byte{
	// 90 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x03, 0x93, 0x42, 0x2c, 0xb9, 0x89, 0x99, 0x79,
	0x4a, 0xc6, 0x5c, 0xdc, 0x01, 0x08, 0x29, 0x21, 0x21, 0x2e, 0x16, 0xbf, 0xc4, 0xdc, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x80, 0x8b, 0xd9, 0x31, 0x3d, 0x55, 0x82,
	0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc4, 0x4c, 0x62, 0x03, 0x2b, 0x36, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x90, 0x4b, 0x7e, 0xb5, 0x56, 0x00, 0x00, 0x00,
}
