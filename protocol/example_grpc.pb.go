// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protocol/example.proto

package protocol

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

// DataInterfaceClient is the client API for DataInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataInterfaceClient interface {
	GetChunk(ctx context.Context, in *ChunkRequest, opts ...grpc.CallOption) (*ChunkResponse, error)
}

type dataInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataInterfaceClient(cc grpc.ClientConnInterface) DataInterfaceClient {
	return &dataInterfaceClient{cc}
}

func (c *dataInterfaceClient) GetChunk(ctx context.Context, in *ChunkRequest, opts ...grpc.CallOption) (*ChunkResponse, error) {
	out := new(ChunkResponse)
	err := c.cc.Invoke(ctx, "/protocol.DataInterface/GetChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataInterfaceServer is the server API for DataInterface service.
// All implementations must embed UnimplementedDataInterfaceServer
// for forward compatibility
type DataInterfaceServer interface {
	GetChunk(context.Context, *ChunkRequest) (*ChunkResponse, error)
	mustEmbedUnimplementedDataInterfaceServer()
}

// UnimplementedDataInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedDataInterfaceServer struct {
}

func (UnimplementedDataInterfaceServer) GetChunk(context.Context, *ChunkRequest) (*ChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChunk not implemented")
}
func (UnimplementedDataInterfaceServer) mustEmbedUnimplementedDataInterfaceServer() {}

// UnsafeDataInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataInterfaceServer will
// result in compilation errors.
type UnsafeDataInterfaceServer interface {
	mustEmbedUnimplementedDataInterfaceServer()
}

func RegisterDataInterfaceServer(s grpc.ServiceRegistrar, srv DataInterfaceServer) {
	s.RegisterService(&DataInterface_ServiceDesc, srv)
}

func _DataInterface_GetChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataInterfaceServer).GetChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.DataInterface/GetChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataInterfaceServer).GetChunk(ctx, req.(*ChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataInterface_ServiceDesc is the grpc.ServiceDesc for DataInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.DataInterface",
	HandlerType: (*DataInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChunk",
			Handler:    _DataInterface_GetChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/example.proto",
}
