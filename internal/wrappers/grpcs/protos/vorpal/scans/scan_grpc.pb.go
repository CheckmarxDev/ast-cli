// Code generated by protoc-gen-go-grpcs. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpcs v1.3.0
// - protoc             v4.25.3
// source: scans/scan.vorpal

package scans

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpcs package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ScanService_Scan_FullMethodName = "/cx.microsast.service.v1.scan.ScanService/Scan"
)

// ScanServiceClient is the client API for ScanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScanServiceClient interface {
	// Performs a scan based on the provided request.
	Scan(ctx context.Context, in *SingleScanRequest, opts ...grpc.CallOption) (*ScanResult, error)
}

type scanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScanServiceClient(cc grpc.ClientConnInterface) ScanServiceClient {
	return &scanServiceClient{cc}
}

func (c *scanServiceClient) Scan(ctx context.Context, in *SingleScanRequest, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := c.cc.Invoke(ctx, ScanService_Scan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScanServiceServer is the server API for ScanService service.
// All implementations must embed UnimplementedScanServiceServer
// for forward compatibility
type ScanServiceServer interface {
	// Performs a scan based on the provided request.
	Scan(context.Context, *SingleScanRequest) (*ScanResult, error)
	mustEmbedUnimplementedScanServiceServer()
}

// UnimplementedScanServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScanServiceServer struct {
}

func (UnimplementedScanServiceServer) Scan(context.Context, *SingleScanRequest) (*ScanResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (UnimplementedScanServiceServer) mustEmbedUnimplementedScanServiceServer() {}

// UnsafeScanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScanServiceServer will
// result in compilation errors.
type UnsafeScanServiceServer interface {
	mustEmbedUnimplementedScanServiceServer()
}

func RegisterScanServiceServer(s grpc.ServiceRegistrar, srv ScanServiceServer) {
	s.RegisterService(&ScanService_ServiceDesc, srv)
}

func _ScanService_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanServiceServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScanService_Scan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanServiceServer).Scan(ctx, req.(*SingleScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScanService_ServiceDesc is the grpc.ServiceDesc for ScanService service.
// It's only intended for direct use with grpcs.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cx.microsast.service.v1.scan.ScanService",
	HandlerType: (*ScanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Scan",
			Handler:    _ScanService_Scan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scans/scan.vorpal",
}