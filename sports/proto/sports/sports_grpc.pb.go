// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sports

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

// SportsClient is the client API for Sports service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SportsClient interface {
	ListSports(ctx context.Context, in *ListSportsRequest, opts ...grpc.CallOption) (*ListSportsResponse, error)
	GetSport(ctx context.Context, in *GetSportRequest, opts ...grpc.CallOption) (*Sport, error)
}

type sportsClient struct {
	cc grpc.ClientConnInterface
}

func NewSportsClient(cc grpc.ClientConnInterface) SportsClient {
	return &sportsClient{cc}
}

func (c *sportsClient) ListSports(ctx context.Context, in *ListSportsRequest, opts ...grpc.CallOption) (*ListSportsResponse, error) {
	out := new(ListSportsResponse)
	err := c.cc.Invoke(ctx, "/sports.Sports/ListSports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sportsClient) GetSport(ctx context.Context, in *GetSportRequest, opts ...grpc.CallOption) (*Sport, error) {
	out := new(Sport)
	err := c.cc.Invoke(ctx, "/sports.Sports/GetSport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SportsServer is the server API for Sports service.
// All implementations should embed UnimplementedSportsServer
// for forward compatibility
type SportsServer interface {
	ListSports(context.Context, *ListSportsRequest) (*ListSportsResponse, error)
	GetSport(context.Context, *GetSportRequest) (*Sport, error)
}

// UnimplementedSportsServer should be embedded to have forward compatible implementations.
type UnimplementedSportsServer struct {
}

func (UnimplementedSportsServer) ListSports(context.Context, *ListSportsRequest) (*ListSportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSports not implemented")
}
func (UnimplementedSportsServer) GetSport(context.Context, *GetSportRequest) (*Sport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSport not implemented")
}

// UnsafeSportsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SportsServer will
// result in compilation errors.
type UnsafeSportsServer interface {
	mustEmbedUnimplementedSportsServer()
}

func RegisterSportsServer(s grpc.ServiceRegistrar, srv SportsServer) {
	s.RegisterService(&Sports_ServiceDesc, srv)
}

func _Sports_ListSports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportsServer).ListSports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sports.Sports/ListSports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportsServer).ListSports(ctx, req.(*ListSportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sports_GetSport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportsServer).GetSport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sports.Sports/GetSport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportsServer).GetSport(ctx, req.(*GetSportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sports_ServiceDesc is the grpc.ServiceDesc for Sports service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sports_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sports.Sports",
	HandlerType: (*SportsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSports",
			Handler:    _Sports_ListSports_Handler,
		},
		{
			MethodName: "GetSport",
			Handler:    _Sports_GetSport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sports/sports.proto",
}
