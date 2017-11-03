// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkin/protobuf/empty.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	checkin/protobuf/empty.proto

It has these top-level messages:
	Empty
*/
package protobuf

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

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(checkin.protobuf.Empty) returns (checkin.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Empty)(nil), "checkin.protobuf.Empty")
}

func init() { proto.RegisterFile("checkin/protobuf/empty.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 75 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xce, 0x48, 0x4d,
	0xce, 0xce, 0xcc, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0xcd, 0x2d,
	0x28, 0xa9, 0xd4, 0x03, 0x73, 0x85, 0x04, 0xa0, 0xb2, 0x7a, 0x30, 0x59, 0x25, 0x76, 0x2e, 0x56,
	0x57, 0x90, 0x02, 0x27, 0xae, 0x28, 0x0e, 0x98, 0x60, 0x12, 0x1b, 0x98, 0x65, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x21, 0xc1, 0xe3, 0x04, 0x4d, 0x00, 0x00, 0x00,
}
