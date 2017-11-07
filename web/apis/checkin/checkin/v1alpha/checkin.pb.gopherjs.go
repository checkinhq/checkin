// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: checkin/checkin/v1alpha/checkin.proto

/*
	Package checkin is a generated protocol buffer package.

	It is generated from these files:
		checkin/checkin/v1alpha/checkin.proto

	It has these top-level messages:
		CheckinRequest
*/
package checkin

import jspb "github.com/johanbrandhorst/protobuf/jspb"
import checkin_protobuf "github.com/checkinhq/checkin/web/apis/checkin/protobuf"

import (
	context "context"

	grpcweb "github.com/johanbrandhorst/protobuf/grpcweb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

type CheckinRequest struct {
	Date         string
	Previous     string
	GoalsReached bool
	Next         string
	Blockers     string
}

// GetDate gets the Date of the CheckinRequest.
func (m *CheckinRequest) GetDate() (x string) {
	if m == nil {
		return x
	}
	return m.Date
}

// GetPrevious gets the Previous of the CheckinRequest.
func (m *CheckinRequest) GetPrevious() (x string) {
	if m == nil {
		return x
	}
	return m.Previous
}

// GetGoalsReached gets the GoalsReached of the CheckinRequest.
func (m *CheckinRequest) GetGoalsReached() (x bool) {
	if m == nil {
		return x
	}
	return m.GoalsReached
}

// GetNext gets the Next of the CheckinRequest.
func (m *CheckinRequest) GetNext() (x string) {
	if m == nil {
		return x
	}
	return m.Next
}

// GetBlockers gets the Blockers of the CheckinRequest.
func (m *CheckinRequest) GetBlockers() (x string) {
	if m == nil {
		return x
	}
	return m.Blockers
}

// MarshalToWriter marshals CheckinRequest to the provided writer.
func (m *CheckinRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Date) > 0 {
		writer.WriteString(1, m.Date)
	}

	if len(m.Previous) > 0 {
		writer.WriteString(2, m.Previous)
	}

	if m.GoalsReached {
		writer.WriteBool(3, m.GoalsReached)
	}

	if len(m.Next) > 0 {
		writer.WriteString(4, m.Next)
	}

	if len(m.Blockers) > 0 {
		writer.WriteString(5, m.Blockers)
	}

	return
}

// Marshal marshals CheckinRequest to a slice of bytes.
func (m *CheckinRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a CheckinRequest from the provided reader.
func (m *CheckinRequest) UnmarshalFromReader(reader jspb.Reader) *CheckinRequest {
	for reader.Next() {
		if m == nil {
			m = &CheckinRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Date = reader.ReadString()
		case 2:
			m.Previous = reader.ReadString()
		case 3:
			m.GoalsReached = reader.ReadBool()
		case 4:
			m.Next = reader.ReadString()
		case 5:
			m.Blockers = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a CheckinRequest from a slice of bytes.
func (m *CheckinRequest) Unmarshal(rawBytes []byte) (*CheckinRequest, error) {
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

// Client API for Checkin service

type CheckinClient interface {
	Checkin(ctx context.Context, in *CheckinRequest, opts ...grpcweb.CallOption) (*checkin_protobuf.Empty, error)
}

type checkinClient struct {
	client *grpcweb.Client
}

// NewCheckinClient creates a new gRPC-Web client.
func NewCheckinClient(hostname string, opts ...grpcweb.DialOption) CheckinClient {
	return &checkinClient{
		client: grpcweb.NewClient(hostname, "checkin.checkin.v1alpha.Checkin", opts...),
	}
}

func (c *checkinClient) Checkin(ctx context.Context, in *CheckinRequest, opts ...grpcweb.CallOption) (*checkin_protobuf.Empty, error) {
	resp, err := c.client.RPCCall(ctx, "Checkin", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(checkin_protobuf.Empty).Unmarshal(resp)
}