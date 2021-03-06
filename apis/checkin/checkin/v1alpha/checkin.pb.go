// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkin/checkin/v1alpha/checkin.proto

/*
Package checkin is a generated protocol buffer package.

It is generated from these files:
	checkin/checkin/v1alpha/checkin.proto

It has these top-level messages:
	CheckinRequest
*/
package checkin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import checkin_protobuf "github.com/checkinhq/checkin/apis/checkin/protobuf"
import _ "github.com/johanbrandhorst/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CheckinRequest struct {
	Date         string `protobuf:"bytes,1,opt,name=date" json:"date,omitempty"`
	Previous     string `protobuf:"bytes,2,opt,name=previous" json:"previous,omitempty"`
	GoalsReached bool   `protobuf:"varint,3,opt,name=goals_reached,json=goalsReached" json:"goals_reached,omitempty"`
	Next         string `protobuf:"bytes,4,opt,name=next" json:"next,omitempty"`
	Blockers     string `protobuf:"bytes,5,opt,name=blockers" json:"blockers,omitempty"`
}

func (m *CheckinRequest) Reset()                    { *m = CheckinRequest{} }
func (m *CheckinRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckinRequest) ProtoMessage()               {}
func (*CheckinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CheckinRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *CheckinRequest) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *CheckinRequest) GetGoalsReached() bool {
	if m != nil {
		return m.GoalsReached
	}
	return false
}

func (m *CheckinRequest) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func (m *CheckinRequest) GetBlockers() string {
	if m != nil {
		return m.Blockers
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckinRequest)(nil), "checkin.checkin.v1alpha.CheckinRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Checkin service

type CheckinClient interface {
	Checkin(ctx context.Context, in *CheckinRequest, opts ...grpc.CallOption) (*checkin_protobuf.Empty, error)
}

type checkinClient struct {
	cc *grpc.ClientConn
}

func NewCheckinClient(cc *grpc.ClientConn) CheckinClient {
	return &checkinClient{cc}
}

func (c *checkinClient) Checkin(ctx context.Context, in *CheckinRequest, opts ...grpc.CallOption) (*checkin_protobuf.Empty, error) {
	out := new(checkin_protobuf.Empty)
	err := grpc.Invoke(ctx, "/checkin.checkin.v1alpha.Checkin/Checkin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Checkin service

type CheckinServer interface {
	Checkin(context.Context, *CheckinRequest) (*checkin_protobuf.Empty, error)
}

func RegisterCheckinServer(s *grpc.Server, srv CheckinServer) {
	s.RegisterService(&_Checkin_serviceDesc, srv)
}

func _Checkin_Checkin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckinServer).Checkin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkin.checkin.v1alpha.Checkin/Checkin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckinServer).Checkin(ctx, req.(*CheckinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Checkin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "checkin.checkin.v1alpha.Checkin",
	HandlerType: (*CheckinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Checkin",
			Handler:    _Checkin_Checkin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkin/checkin/v1alpha/checkin.proto",
}

func init() { proto.RegisterFile("checkin/checkin/v1alpha/checkin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xbd, 0x4e, 0xc3, 0x30,
	0x18, 0x54, 0xa0, 0xd0, 0x62, 0x01, 0x83, 0x19, 0x6a, 0x45, 0x0c, 0x15, 0x08, 0xd1, 0xc9, 0x11,
	0xb0, 0x30, 0x22, 0x10, 0x13, 0x5b, 0x06, 0x06, 0x16, 0x64, 0x3b, 0x5f, 0xe3, 0xb4, 0x69, 0x6c,
	0xfc, 0x53, 0xc1, 0x83, 0xf0, 0xb2, 0x4c, 0x28, 0x76, 0x4c, 0xd5, 0x81, 0xe9, 0xbb, 0xbb, 0xdc,
	0x9d, 0xe2, 0x43, 0x57, 0x42, 0x82, 0x58, 0x35, 0x5d, 0x91, 0xee, 0xe6, 0x86, 0xb5, 0x5a, 0xb2,
	0xc4, 0xa9, 0x36, 0xca, 0x29, 0x3c, 0x4d, 0x34, 0xdd, 0xc1, 0x96, 0x9f, 0xa7, 0x5c, 0xf0, 0x71,
	0xbf, 0x28, 0x60, 0xad, 0xdd, 0x57, 0x8c, 0xe5, 0xf7, 0x75, 0xe3, 0xa4, 0xe7, 0x54, 0xa8, 0x75,
	0xb1, 0x54, 0x92, 0x75, 0xdc, 0xb0, 0xae, 0x92, 0xca, 0x58, 0xb7, 0x0d, 0x04, 0x50, 0xd4, 0x4a,
	0x4b, 0x30, 0x4b, 0x1b, 0x93, 0x17, 0xdf, 0x19, 0x3a, 0x7d, 0x8a, 0xd5, 0x25, 0x7c, 0x78, 0xb0,
	0x0e, 0x63, 0x34, 0xaa, 0x98, 0x03, 0x92, 0xcd, 0xb2, 0xf9, 0x51, 0x19, 0x30, 0xce, 0xd1, 0x44,
	0x1b, 0xd8, 0x34, 0xca, 0x5b, 0xb2, 0x17, 0xf4, 0x3f, 0x8e, 0x2f, 0xd1, 0x49, 0xad, 0x58, 0x6b,
	0xdf, 0x0d, 0x30, 0x21, 0xa1, 0x22, 0xfb, 0xb3, 0x6c, 0x3e, 0x29, 0x8f, 0x83, 0x58, 0x46, 0xad,
	0x2f, 0xed, 0xe0, 0xd3, 0x91, 0x51, 0x2c, 0xed, 0x71, 0x5f, 0xca, 0x5b, 0x25, 0x56, 0x60, 0x2c,
	0x39, 0x88, 0xa5, 0x89, 0xdf, 0xbe, 0xa2, 0xf1, 0xf0, 0x5b, 0xf8, 0x65, 0x0b, 0xaf, 0xe9, 0x3f,
	0xfb, 0xd0, 0xdd, 0x37, 0xe4, 0x53, 0xba, 0xb3, 0x2b, 0xf7, 0x0b, 0xfa, 0xdc, 0xef, 0xf5, 0x78,
	0xf6, 0xf3, 0x30, 0x1e, 0x3e, 0xbd, 0x25, 0xc0, 0x0f, 0x83, 0xe9, 0xee, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x4c, 0xa5, 0x4b, 0xec, 0xa5, 0x01, 0x00, 0x00,
}
