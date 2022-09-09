// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: job_postings.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// JobPostingsClient is the client API for JobPostings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobPostingsClient interface {
	CreateJobPostings(ctx context.Context, opts ...grpc.CallOption) (JobPostings_CreateJobPostingsClient, error)
	StreamJobPostings(ctx context.Context, in *StreamJobPostingsRequest, opts ...grpc.CallOption) (JobPostings_StreamJobPostingsClient, error)
}

type jobPostingsClient struct {
	cc grpc.ClientConnInterface
}

func NewJobPostingsClient(cc grpc.ClientConnInterface) JobPostingsClient {
	return &jobPostingsClient{cc}
}

func (c *jobPostingsClient) CreateJobPostings(ctx context.Context, opts ...grpc.CallOption) (JobPostings_CreateJobPostingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &JobPostings_ServiceDesc.Streams[0], "/job_postings.JobPostings/CreateJobPostings", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobPostingsCreateJobPostingsClient{stream}
	return x, nil
}

type JobPostings_CreateJobPostingsClient interface {
	Send(*CreateJobPostingsRequest) error
	CloseAndRecv() (*CreateJobPostingsResponse, error)
	grpc.ClientStream
}

type jobPostingsCreateJobPostingsClient struct {
	grpc.ClientStream
}

func (x *jobPostingsCreateJobPostingsClient) Send(m *CreateJobPostingsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *jobPostingsCreateJobPostingsClient) CloseAndRecv() (*CreateJobPostingsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateJobPostingsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jobPostingsClient) StreamJobPostings(ctx context.Context, in *StreamJobPostingsRequest, opts ...grpc.CallOption) (JobPostings_StreamJobPostingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &JobPostings_ServiceDesc.Streams[1], "/job_postings.JobPostings/StreamJobPostings", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobPostingsStreamJobPostingsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JobPostings_StreamJobPostingsClient interface {
	Recv() (*JobPosting, error)
	grpc.ClientStream
}

type jobPostingsStreamJobPostingsClient struct {
	grpc.ClientStream
}

func (x *jobPostingsStreamJobPostingsClient) Recv() (*JobPosting, error) {
	m := new(JobPosting)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// JobPostingsServer is the server API for JobPostings service.
// All implementations must embed UnimplementedJobPostingsServer
// for forward compatibility
type JobPostingsServer interface {
	CreateJobPostings(JobPostings_CreateJobPostingsServer) error
	StreamJobPostings(*StreamJobPostingsRequest, JobPostings_StreamJobPostingsServer) error
	mustEmbedUnimplementedJobPostingsServer()
}

// UnimplementedJobPostingsServer must be embedded to have forward compatible implementations.
type UnimplementedJobPostingsServer struct {
}

func (UnimplementedJobPostingsServer) CreateJobPostings(JobPostings_CreateJobPostingsServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateJobPostings not implemented")
}
func (UnimplementedJobPostingsServer) StreamJobPostings(*StreamJobPostingsRequest, JobPostings_StreamJobPostingsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamJobPostings not implemented")
}
func (UnimplementedJobPostingsServer) mustEmbedUnimplementedJobPostingsServer() {}

// UnsafeJobPostingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobPostingsServer will
// result in compilation errors.
type UnsafeJobPostingsServer interface {
	mustEmbedUnimplementedJobPostingsServer()
}

func RegisterJobPostingsServer(s grpc.ServiceRegistrar, srv JobPostingsServer) {
	s.RegisterService(&JobPostings_ServiceDesc, srv)
}

func _JobPostings_CreateJobPostings_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JobPostingsServer).CreateJobPostings(&jobPostingsCreateJobPostingsServer{stream})
}

type JobPostings_CreateJobPostingsServer interface {
	SendAndClose(*CreateJobPostingsResponse) error
	Recv() (*CreateJobPostingsRequest, error)
	grpc.ServerStream
}

type jobPostingsCreateJobPostingsServer struct {
	grpc.ServerStream
}

func (x *jobPostingsCreateJobPostingsServer) SendAndClose(m *CreateJobPostingsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *jobPostingsCreateJobPostingsServer) Recv() (*CreateJobPostingsRequest, error) {
	m := new(CreateJobPostingsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _JobPostings_StreamJobPostings_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamJobPostingsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobPostingsServer).StreamJobPostings(m, &jobPostingsStreamJobPostingsServer{stream})
}

type JobPostings_StreamJobPostingsServer interface {
	Send(*JobPosting) error
	grpc.ServerStream
}

type jobPostingsStreamJobPostingsServer struct {
	grpc.ServerStream
}

func (x *jobPostingsStreamJobPostingsServer) Send(m *JobPosting) error {
	return x.ServerStream.SendMsg(m)
}

// JobPostings_ServiceDesc is the grpc.ServiceDesc for JobPostings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobPostings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job_postings.JobPostings",
	HandlerType: (*JobPostingsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateJobPostings",
			Handler:       _JobPostings_CreateJobPostings_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamJobPostings",
			Handler:       _JobPostings_StreamJobPostings_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "job_postings.proto",
}
