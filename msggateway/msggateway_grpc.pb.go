// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: msggateway/msggateway.proto

package msggateway

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

const (
	MsgGateway_OnlinePushMsg_FullMethodName                   = "/OpenIMServer.msggateway.msgGateway/OnlinePushMsg"
	MsgGateway_GetUsersOnlineStatus_FullMethodName            = "/OpenIMServer.msggateway.msgGateway/GetUsersOnlineStatus"
	MsgGateway_OnlineBatchPushOneMsg_FullMethodName           = "/OpenIMServer.msggateway.msgGateway/OnlineBatchPushOneMsg"
	MsgGateway_SuperGroupOnlineBatchPushOneMsg_FullMethodName = "/OpenIMServer.msggateway.msgGateway/SuperGroupOnlineBatchPushOneMsg"
	MsgGateway_KickUserOffline_FullMethodName                 = "/OpenIMServer.msggateway.msgGateway/KickUserOffline"
	MsgGateway_MultiTerminalLoginCheck_FullMethodName         = "/OpenIMServer.msggateway.msgGateway/MultiTerminalLoginCheck"
)

// MsgGatewayClient is the client API for MsgGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgGatewayClient interface {
	OnlinePushMsg(ctx context.Context, in *OnlinePushMsgReq, opts ...grpc.CallOption) (*OnlinePushMsgResp, error)
	GetUsersOnlineStatus(ctx context.Context, in *GetUsersOnlineStatusReq, opts ...grpc.CallOption) (*GetUsersOnlineStatusResp, error)
	OnlineBatchPushOneMsg(ctx context.Context, in *OnlineBatchPushOneMsgReq, opts ...grpc.CallOption) (*OnlineBatchPushOneMsgResp, error)
	SuperGroupOnlineBatchPushOneMsg(ctx context.Context, in *OnlineBatchPushOneMsgReq, opts ...grpc.CallOption) (*OnlineBatchPushOneMsgResp, error)
	KickUserOffline(ctx context.Context, in *KickUserOfflineReq, opts ...grpc.CallOption) (*KickUserOfflineResp, error)
	MultiTerminalLoginCheck(ctx context.Context, in *MultiTerminalLoginCheckReq, opts ...grpc.CallOption) (*MultiTerminalLoginCheckResp, error)
}

type msgGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgGatewayClient(cc grpc.ClientConnInterface) MsgGatewayClient {
	return &msgGatewayClient{cc}
}

func (c *msgGatewayClient) OnlinePushMsg(ctx context.Context, in *OnlinePushMsgReq, opts ...grpc.CallOption) (*OnlinePushMsgResp, error) {
	out := new(OnlinePushMsgResp)
	err := c.cc.Invoke(ctx, MsgGateway_OnlinePushMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgGatewayClient) GetUsersOnlineStatus(ctx context.Context, in *GetUsersOnlineStatusReq, opts ...grpc.CallOption) (*GetUsersOnlineStatusResp, error) {
	out := new(GetUsersOnlineStatusResp)
	err := c.cc.Invoke(ctx, MsgGateway_GetUsersOnlineStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgGatewayClient) OnlineBatchPushOneMsg(ctx context.Context, in *OnlineBatchPushOneMsgReq, opts ...grpc.CallOption) (*OnlineBatchPushOneMsgResp, error) {
	out := new(OnlineBatchPushOneMsgResp)
	err := c.cc.Invoke(ctx, MsgGateway_OnlineBatchPushOneMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgGatewayClient) SuperGroupOnlineBatchPushOneMsg(ctx context.Context, in *OnlineBatchPushOneMsgReq, opts ...grpc.CallOption) (*OnlineBatchPushOneMsgResp, error) {
	out := new(OnlineBatchPushOneMsgResp)
	err := c.cc.Invoke(ctx, MsgGateway_SuperGroupOnlineBatchPushOneMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgGatewayClient) KickUserOffline(ctx context.Context, in *KickUserOfflineReq, opts ...grpc.CallOption) (*KickUserOfflineResp, error) {
	out := new(KickUserOfflineResp)
	err := c.cc.Invoke(ctx, MsgGateway_KickUserOffline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgGatewayClient) MultiTerminalLoginCheck(ctx context.Context, in *MultiTerminalLoginCheckReq, opts ...grpc.CallOption) (*MultiTerminalLoginCheckResp, error) {
	out := new(MultiTerminalLoginCheckResp)
	err := c.cc.Invoke(ctx, MsgGateway_MultiTerminalLoginCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgGatewayServer is the server API for MsgGateway service.
// All implementations must embed UnimplementedMsgGatewayServer
// for forward compatibility
type MsgGatewayServer interface {
	OnlinePushMsg(context.Context, *OnlinePushMsgReq) (*OnlinePushMsgResp, error)
	GetUsersOnlineStatus(context.Context, *GetUsersOnlineStatusReq) (*GetUsersOnlineStatusResp, error)
	OnlineBatchPushOneMsg(context.Context, *OnlineBatchPushOneMsgReq) (*OnlineBatchPushOneMsgResp, error)
	SuperGroupOnlineBatchPushOneMsg(context.Context, *OnlineBatchPushOneMsgReq) (*OnlineBatchPushOneMsgResp, error)
	KickUserOffline(context.Context, *KickUserOfflineReq) (*KickUserOfflineResp, error)
	MultiTerminalLoginCheck(context.Context, *MultiTerminalLoginCheckReq) (*MultiTerminalLoginCheckResp, error)
	mustEmbedUnimplementedMsgGatewayServer()
}

// UnimplementedMsgGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedMsgGatewayServer struct {
}

func (UnimplementedMsgGatewayServer) OnlinePushMsg(context.Context, *OnlinePushMsgReq) (*OnlinePushMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlinePushMsg not implemented")
}
func (UnimplementedMsgGatewayServer) GetUsersOnlineStatus(context.Context, *GetUsersOnlineStatusReq) (*GetUsersOnlineStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersOnlineStatus not implemented")
}
func (UnimplementedMsgGatewayServer) OnlineBatchPushOneMsg(context.Context, *OnlineBatchPushOneMsgReq) (*OnlineBatchPushOneMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineBatchPushOneMsg not implemented")
}
func (UnimplementedMsgGatewayServer) SuperGroupOnlineBatchPushOneMsg(context.Context, *OnlineBatchPushOneMsgReq) (*OnlineBatchPushOneMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuperGroupOnlineBatchPushOneMsg not implemented")
}
func (UnimplementedMsgGatewayServer) KickUserOffline(context.Context, *KickUserOfflineReq) (*KickUserOfflineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickUserOffline not implemented")
}
func (UnimplementedMsgGatewayServer) MultiTerminalLoginCheck(context.Context, *MultiTerminalLoginCheckReq) (*MultiTerminalLoginCheckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiTerminalLoginCheck not implemented")
}
func (UnimplementedMsgGatewayServer) mustEmbedUnimplementedMsgGatewayServer() {}

// UnsafeMsgGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgGatewayServer will
// result in compilation errors.
type UnsafeMsgGatewayServer interface {
	mustEmbedUnimplementedMsgGatewayServer()
}

func RegisterMsgGatewayServer(s grpc.ServiceRegistrar, srv MsgGatewayServer) {
	s.RegisterService(&MsgGateway_ServiceDesc, srv)
}

func _MsgGateway_OnlinePushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlinePushMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgGatewayServer).OnlinePushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgGateway_OnlinePushMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgGatewayServer).OnlinePushMsg(ctx, req.(*OnlinePushMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgGateway_GetUsersOnlineStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersOnlineStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgGatewayServer).GetUsersOnlineStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgGateway_GetUsersOnlineStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgGatewayServer).GetUsersOnlineStatus(ctx, req.(*GetUsersOnlineStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgGateway_OnlineBatchPushOneMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineBatchPushOneMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgGatewayServer).OnlineBatchPushOneMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgGateway_OnlineBatchPushOneMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgGatewayServer).OnlineBatchPushOneMsg(ctx, req.(*OnlineBatchPushOneMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgGateway_SuperGroupOnlineBatchPushOneMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineBatchPushOneMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgGatewayServer).SuperGroupOnlineBatchPushOneMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgGateway_SuperGroupOnlineBatchPushOneMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgGatewayServer).SuperGroupOnlineBatchPushOneMsg(ctx, req.(*OnlineBatchPushOneMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgGateway_KickUserOffline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickUserOfflineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgGatewayServer).KickUserOffline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgGateway_KickUserOffline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgGatewayServer).KickUserOffline(ctx, req.(*KickUserOfflineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgGateway_MultiTerminalLoginCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiTerminalLoginCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgGatewayServer).MultiTerminalLoginCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgGateway_MultiTerminalLoginCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgGatewayServer).MultiTerminalLoginCheck(ctx, req.(*MultiTerminalLoginCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MsgGateway_ServiceDesc is the grpc.ServiceDesc for MsgGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MsgGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OpenIMServer.msggateway.msgGateway",
	HandlerType: (*MsgGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnlinePushMsg",
			Handler:    _MsgGateway_OnlinePushMsg_Handler,
		},
		{
			MethodName: "GetUsersOnlineStatus",
			Handler:    _MsgGateway_GetUsersOnlineStatus_Handler,
		},
		{
			MethodName: "OnlineBatchPushOneMsg",
			Handler:    _MsgGateway_OnlineBatchPushOneMsg_Handler,
		},
		{
			MethodName: "SuperGroupOnlineBatchPushOneMsg",
			Handler:    _MsgGateway_SuperGroupOnlineBatchPushOneMsg_Handler,
		},
		{
			MethodName: "KickUserOffline",
			Handler:    _MsgGateway_KickUserOffline_Handler,
		},
		{
			MethodName: "MultiTerminalLoginCheck",
			Handler:    _MsgGateway_MultiTerminalLoginCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msggateway/msggateway.proto",
}
