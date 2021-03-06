// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: checkin/user/v1alpha/user.proto

package user

import jspb "github.com/johanbrandhorst/protobuf/jspb"
import checkin_protobuf "github.com/checkinhq/checkin/web/apis/checkin/protobuf"

import (
	context "context"

	grpcweb "github.com/johanbrandhorst/protobuf/grpcweb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

type CreateUserRequest struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

// GetEmail gets the Email of the CreateUserRequest.
func (m *CreateUserRequest) GetEmail() (x string) {
	if m == nil {
		return x
	}
	return m.Email
}

// GetPassword gets the Password of the CreateUserRequest.
func (m *CreateUserRequest) GetPassword() (x string) {
	if m == nil {
		return x
	}
	return m.Password
}

// GetFirstName gets the FirstName of the CreateUserRequest.
func (m *CreateUserRequest) GetFirstName() (x string) {
	if m == nil {
		return x
	}
	return m.FirstName
}

// GetLastName gets the LastName of the CreateUserRequest.
func (m *CreateUserRequest) GetLastName() (x string) {
	if m == nil {
		return x
	}
	return m.LastName
}

// MarshalToWriter marshals CreateUserRequest to the provided writer.
func (m *CreateUserRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Email) > 0 {
		writer.WriteString(1, m.Email)
	}

	if len(m.Password) > 0 {
		writer.WriteString(2, m.Password)
	}

	if len(m.FirstName) > 0 {
		writer.WriteString(3, m.FirstName)
	}

	if len(m.LastName) > 0 {
		writer.WriteString(4, m.LastName)
	}

	return
}

// Marshal marshals CreateUserRequest to a slice of bytes.
func (m *CreateUserRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a CreateUserRequest from the provided reader.
func (m *CreateUserRequest) UnmarshalFromReader(reader jspb.Reader) *CreateUserRequest {
	for reader.Next() {
		if m == nil {
			m = &CreateUserRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Email = reader.ReadString()
		case 2:
			m.Password = reader.ReadString()
		case 3:
			m.FirstName = reader.ReadString()
		case 4:
			m.LastName = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a CreateUserRequest from a slice of bytes.
func (m *CreateUserRequest) Unmarshal(rawBytes []byte) (*CreateUserRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpcweb.Client

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpcweb package it is being compiled against.
const _ = grpcweb.GrpcWebPackageIsVersion2

// Client API for User service

type UserClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpcweb.CallOption) (*checkin_protobuf.Empty, error)
}

type userClient struct {
	client *grpcweb.Client
}

// NewUserClient creates a new gRPC-Web client.
func NewUserClient(hostname string, opts ...grpcweb.DialOption) UserClient {
	return &userClient{
		client: grpcweb.NewClient(hostname, "checkin.user.v1alpha.User", opts...),
	}
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpcweb.CallOption) (*checkin_protobuf.Empty, error) {
	resp, err := c.client.RPCCall(ctx, "CreateUser", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(checkin_protobuf.Empty).Unmarshal(resp)
}
