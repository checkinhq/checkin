// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkin/user/v1alpha/user.proto

package user

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

type CreateUserRequest struct {
	Email     string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CreateUserRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "checkin.user.v1alpha.CreateUserRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*checkin_protobuf.Empty, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*checkin_protobuf.Empty, error) {
	out := new(checkin_protobuf.Empty)
	err := grpc.Invoke(ctx, "/checkin.user.v1alpha.User/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*checkin_protobuf.Empty, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkin.user.v1alpha.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "checkin.user.v1alpha.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkin/user/v1alpha/user.proto",
}

func init() { proto.RegisterFile("checkin/user/v1alpha/user.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0xa9, 0x56, 0xb1, 0x0f, 0x1c, 0x0c, 0x07, 0x96, 0xaa, 0x28, 0x2e, 0x3a, 0x25, 0xa8,
	0x8b, 0xa3, 0x28, 0xae, 0x37, 0x1c, 0xe8, 0xe0, 0x22, 0xaf, 0xbd, 0x77, 0x97, 0x9e, 0x4d, 0x13,
	0x93, 0x54, 0x71, 0xf4, 0xdf, 0x76, 0x92, 0xa4, 0xd7, 0xeb, 0xa0, 0x4b, 0xc8, 0xf7, 0xc7, 0x87,
	0xbc, 0x3c, 0x38, 0xad, 0x24, 0x55, 0x6f, 0x75, 0x2b, 0x3a, 0x47, 0x56, 0x7c, 0x5c, 0x61, 0x63,
	0x24, 0x46, 0xc1, 0x8d, 0xd5, 0x5e, 0xb3, 0xc9, 0xba, 0xc0, 0xa3, 0xb7, 0x2e, 0x14, 0xc7, 0x03,
	0x16, 0x4b, 0x65, 0xb7, 0x10, 0xa4, 0x8c, 0xff, 0xea, 0x99, 0xe2, 0x76, 0x59, 0x7b, 0xd9, 0x95,
	0xbc, 0xd2, 0x4a, 0xac, 0xb4, 0xc4, 0xb6, 0xb4, 0xd8, 0xce, 0xa5, 0xb6, 0xce, 0x8f, 0x40, 0xbc,
	0x88, 0xa5, 0x36, 0x92, 0xec, 0xca, 0xf5, 0xe4, 0xf9, 0x77, 0x02, 0x07, 0x0f, 0x96, 0xd0, 0xd3,
	0x93, 0x23, 0x3b, 0xa3, 0xf7, 0x8e, 0x9c, 0x67, 0x13, 0xd8, 0x21, 0x85, 0x75, 0x93, 0x27, 0x67,
	0xc9, 0x65, 0x36, 0xeb, 0x05, 0x2b, 0x60, 0xcf, 0xa0, 0x73, 0x9f, 0xda, 0xce, 0xf3, 0xad, 0x18,
	0x6c, 0x34, 0x3b, 0x01, 0x58, 0xd4, 0xd6, 0xf9, 0xd7, 0x16, 0x15, 0xe5, 0xdb, 0x31, 0xcd, 0xa2,
	0x33, 0x45, 0x45, 0xec, 0x08, 0xb2, 0x06, 0x87, 0x34, 0xed, 0xd9, 0x60, 0x84, 0xf0, 0xfa, 0x19,
	0xd2, 0xf0, 0x38, 0x9b, 0x02, 0x8c, 0xa3, 0xb0, 0x0b, 0xfe, 0xdf, 0x22, 0xf8, 0x9f, 0x61, 0x8b,
	0xc3, 0x4d, 0x71, 0xf8, 0x2a, 0x7f, 0x0c, 0xbb, 0xb9, 0xdf, 0xff, 0xb9, 0x4b, 0x03, 0xfb, 0x12,
	0xcf, 0x72, 0x37, 0xc6, 0x37, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xea, 0x0e, 0x40, 0x80, 0x82,
	0x01, 0x00, 0x00,
}
