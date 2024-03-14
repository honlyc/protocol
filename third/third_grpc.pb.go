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
// source: third/third.proto

package third

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
	Third_PartLimit_FullMethodName               = "/OpenIMServer.third.third/PartLimit"
	Third_PartSize_FullMethodName                = "/OpenIMServer.third.third/PartSize"
	Third_InitiateMultipartUpload_FullMethodName = "/OpenIMServer.third.third/InitiateMultipartUpload"
	Third_AuthSign_FullMethodName                = "/OpenIMServer.third.third/AuthSign"
	Third_CompleteMultipartUpload_FullMethodName = "/OpenIMServer.third.third/CompleteMultipartUpload"
	Third_AccessURL_FullMethodName               = "/OpenIMServer.third.third/AccessURL"
	Third_InitiateFormData_FullMethodName        = "/OpenIMServer.third.third/InitiateFormData"
	Third_CompleteFormData_FullMethodName        = "/OpenIMServer.third.third/CompleteFormData"
	Third_FcmUpdateToken_FullMethodName          = "/OpenIMServer.third.third/FcmUpdateToken"
	Third_SetAppBadge_FullMethodName             = "/OpenIMServer.third.third/SetAppBadge"
	Third_UploadLogs_FullMethodName              = "/OpenIMServer.third.third/UploadLogs"
	Third_DeleteLogs_FullMethodName              = "/OpenIMServer.third.third/DeleteLogs"
	Third_SearchLogs_FullMethodName              = "/OpenIMServer.third.third/SearchLogs"
)

// ThirdClient is the client API for Third service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThirdClient interface {
	PartLimit(ctx context.Context, in *PartLimitReq, opts ...grpc.CallOption) (*PartLimitResp, error)
	PartSize(ctx context.Context, in *PartSizeReq, opts ...grpc.CallOption) (*PartSizeResp, error)
	InitiateMultipartUpload(ctx context.Context, in *InitiateMultipartUploadReq, opts ...grpc.CallOption) (*InitiateMultipartUploadResp, error)
	AuthSign(ctx context.Context, in *AuthSignReq, opts ...grpc.CallOption) (*AuthSignResp, error)
	CompleteMultipartUpload(ctx context.Context, in *CompleteMultipartUploadReq, opts ...grpc.CallOption) (*CompleteMultipartUploadResp, error)
	AccessURL(ctx context.Context, in *AccessURLReq, opts ...grpc.CallOption) (*AccessURLResp, error)
	InitiateFormData(ctx context.Context, in *InitiateFormDataReq, opts ...grpc.CallOption) (*InitiateFormDataResp, error)
	CompleteFormData(ctx context.Context, in *CompleteFormDataReq, opts ...grpc.CallOption) (*CompleteFormDataResp, error)
	FcmUpdateToken(ctx context.Context, in *FcmUpdateTokenReq, opts ...grpc.CallOption) (*FcmUpdateTokenResp, error)
	SetAppBadge(ctx context.Context, in *SetAppBadgeReq, opts ...grpc.CallOption) (*SetAppBadgeResp, error)
	// 日志
	UploadLogs(ctx context.Context, in *UploadLogsReq, opts ...grpc.CallOption) (*UploadLogsResp, error)
	DeleteLogs(ctx context.Context, in *DeleteLogsReq, opts ...grpc.CallOption) (*DeleteLogsResp, error)
	SearchLogs(ctx context.Context, in *SearchLogsReq, opts ...grpc.CallOption) (*SearchLogsResp, error)
}

type thirdClient struct {
	cc grpc.ClientConnInterface
}

func NewThirdClient(cc grpc.ClientConnInterface) ThirdClient {
	return &thirdClient{cc}
}

func (c *thirdClient) PartLimit(ctx context.Context, in *PartLimitReq, opts ...grpc.CallOption) (*PartLimitResp, error) {
	out := new(PartLimitResp)
	err := c.cc.Invoke(ctx, Third_PartLimit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) PartSize(ctx context.Context, in *PartSizeReq, opts ...grpc.CallOption) (*PartSizeResp, error) {
	out := new(PartSizeResp)
	err := c.cc.Invoke(ctx, Third_PartSize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) InitiateMultipartUpload(ctx context.Context, in *InitiateMultipartUploadReq, opts ...grpc.CallOption) (*InitiateMultipartUploadResp, error) {
	out := new(InitiateMultipartUploadResp)
	err := c.cc.Invoke(ctx, Third_InitiateMultipartUpload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) AuthSign(ctx context.Context, in *AuthSignReq, opts ...grpc.CallOption) (*AuthSignResp, error) {
	out := new(AuthSignResp)
	err := c.cc.Invoke(ctx, Third_AuthSign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) CompleteMultipartUpload(ctx context.Context, in *CompleteMultipartUploadReq, opts ...grpc.CallOption) (*CompleteMultipartUploadResp, error) {
	out := new(CompleteMultipartUploadResp)
	err := c.cc.Invoke(ctx, Third_CompleteMultipartUpload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) AccessURL(ctx context.Context, in *AccessURLReq, opts ...grpc.CallOption) (*AccessURLResp, error) {
	out := new(AccessURLResp)
	err := c.cc.Invoke(ctx, Third_AccessURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) InitiateFormData(ctx context.Context, in *InitiateFormDataReq, opts ...grpc.CallOption) (*InitiateFormDataResp, error) {
	out := new(InitiateFormDataResp)
	err := c.cc.Invoke(ctx, Third_InitiateFormData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) CompleteFormData(ctx context.Context, in *CompleteFormDataReq, opts ...grpc.CallOption) (*CompleteFormDataResp, error) {
	out := new(CompleteFormDataResp)
	err := c.cc.Invoke(ctx, Third_CompleteFormData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) FcmUpdateToken(ctx context.Context, in *FcmUpdateTokenReq, opts ...grpc.CallOption) (*FcmUpdateTokenResp, error) {
	out := new(FcmUpdateTokenResp)
	err := c.cc.Invoke(ctx, Third_FcmUpdateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) SetAppBadge(ctx context.Context, in *SetAppBadgeReq, opts ...grpc.CallOption) (*SetAppBadgeResp, error) {
	out := new(SetAppBadgeResp)
	err := c.cc.Invoke(ctx, Third_SetAppBadge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) UploadLogs(ctx context.Context, in *UploadLogsReq, opts ...grpc.CallOption) (*UploadLogsResp, error) {
	out := new(UploadLogsResp)
	err := c.cc.Invoke(ctx, Third_UploadLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) DeleteLogs(ctx context.Context, in *DeleteLogsReq, opts ...grpc.CallOption) (*DeleteLogsResp, error) {
	out := new(DeleteLogsResp)
	err := c.cc.Invoke(ctx, Third_DeleteLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdClient) SearchLogs(ctx context.Context, in *SearchLogsReq, opts ...grpc.CallOption) (*SearchLogsResp, error) {
	out := new(SearchLogsResp)
	err := c.cc.Invoke(ctx, Third_SearchLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThirdServer is the server API for Third service.
// All implementations must embed UnimplementedThirdServer
// for forward compatibility
type ThirdServer interface {
	PartLimit(context.Context, *PartLimitReq) (*PartLimitResp, error)
	PartSize(context.Context, *PartSizeReq) (*PartSizeResp, error)
	InitiateMultipartUpload(context.Context, *InitiateMultipartUploadReq) (*InitiateMultipartUploadResp, error)
	AuthSign(context.Context, *AuthSignReq) (*AuthSignResp, error)
	CompleteMultipartUpload(context.Context, *CompleteMultipartUploadReq) (*CompleteMultipartUploadResp, error)
	AccessURL(context.Context, *AccessURLReq) (*AccessURLResp, error)
	InitiateFormData(context.Context, *InitiateFormDataReq) (*InitiateFormDataResp, error)
	CompleteFormData(context.Context, *CompleteFormDataReq) (*CompleteFormDataResp, error)
	FcmUpdateToken(context.Context, *FcmUpdateTokenReq) (*FcmUpdateTokenResp, error)
	SetAppBadge(context.Context, *SetAppBadgeReq) (*SetAppBadgeResp, error)
	// 日志
	UploadLogs(context.Context, *UploadLogsReq) (*UploadLogsResp, error)
	DeleteLogs(context.Context, *DeleteLogsReq) (*DeleteLogsResp, error)
	SearchLogs(context.Context, *SearchLogsReq) (*SearchLogsResp, error)
	mustEmbedUnimplementedThirdServer()
}

// UnimplementedThirdServer must be embedded to have forward compatible implementations.
type UnimplementedThirdServer struct {
}

func (UnimplementedThirdServer) PartLimit(context.Context, *PartLimitReq) (*PartLimitResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PartLimit not implemented")
}
func (UnimplementedThirdServer) PartSize(context.Context, *PartSizeReq) (*PartSizeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PartSize not implemented")
}
func (UnimplementedThirdServer) InitiateMultipartUpload(context.Context, *InitiateMultipartUploadReq) (*InitiateMultipartUploadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateMultipartUpload not implemented")
}
func (UnimplementedThirdServer) AuthSign(context.Context, *AuthSignReq) (*AuthSignResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthSign not implemented")
}
func (UnimplementedThirdServer) CompleteMultipartUpload(context.Context, *CompleteMultipartUploadReq) (*CompleteMultipartUploadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteMultipartUpload not implemented")
}
func (UnimplementedThirdServer) AccessURL(context.Context, *AccessURLReq) (*AccessURLResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessURL not implemented")
}
func (UnimplementedThirdServer) InitiateFormData(context.Context, *InitiateFormDataReq) (*InitiateFormDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateFormData not implemented")
}
func (UnimplementedThirdServer) CompleteFormData(context.Context, *CompleteFormDataReq) (*CompleteFormDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteFormData not implemented")
}
func (UnimplementedThirdServer) FcmUpdateToken(context.Context, *FcmUpdateTokenReq) (*FcmUpdateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FcmUpdateToken not implemented")
}
func (UnimplementedThirdServer) SetAppBadge(context.Context, *SetAppBadgeReq) (*SetAppBadgeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAppBadge not implemented")
}
func (UnimplementedThirdServer) UploadLogs(context.Context, *UploadLogsReq) (*UploadLogsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadLogs not implemented")
}
func (UnimplementedThirdServer) DeleteLogs(context.Context, *DeleteLogsReq) (*DeleteLogsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLogs not implemented")
}
func (UnimplementedThirdServer) SearchLogs(context.Context, *SearchLogsReq) (*SearchLogsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchLogs not implemented")
}
func (UnimplementedThirdServer) mustEmbedUnimplementedThirdServer() {}

// UnsafeThirdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThirdServer will
// result in compilation errors.
type UnsafeThirdServer interface {
	mustEmbedUnimplementedThirdServer()
}

func RegisterThirdServer(s grpc.ServiceRegistrar, srv ThirdServer) {
	s.RegisterService(&Third_ServiceDesc, srv)
}

func _Third_PartLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartLimitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).PartLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_PartLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).PartLimit(ctx, req.(*PartLimitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_PartSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartSizeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).PartSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_PartSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).PartSize(ctx, req.(*PartSizeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_InitiateMultipartUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitiateMultipartUploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).InitiateMultipartUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_InitiateMultipartUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).InitiateMultipartUpload(ctx, req.(*InitiateMultipartUploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_AuthSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthSignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).AuthSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_AuthSign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).AuthSign(ctx, req.(*AuthSignReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_CompleteMultipartUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteMultipartUploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).CompleteMultipartUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_CompleteMultipartUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).CompleteMultipartUpload(ctx, req.(*CompleteMultipartUploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_AccessURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).AccessURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_AccessURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).AccessURL(ctx, req.(*AccessURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_InitiateFormData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitiateFormDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).InitiateFormData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_InitiateFormData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).InitiateFormData(ctx, req.(*InitiateFormDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_CompleteFormData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteFormDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).CompleteFormData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_CompleteFormData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).CompleteFormData(ctx, req.(*CompleteFormDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_FcmUpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FcmUpdateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).FcmUpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_FcmUpdateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).FcmUpdateToken(ctx, req.(*FcmUpdateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_SetAppBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAppBadgeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).SetAppBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_SetAppBadge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).SetAppBadge(ctx, req.(*SetAppBadgeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_UploadLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadLogsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).UploadLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_UploadLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).UploadLogs(ctx, req.(*UploadLogsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_DeleteLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).DeleteLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_DeleteLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).DeleteLogs(ctx, req.(*DeleteLogsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Third_SearchLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchLogsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdServer).SearchLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Third_SearchLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdServer).SearchLogs(ctx, req.(*SearchLogsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Third_ServiceDesc is the grpc.ServiceDesc for Third service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Third_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OpenIMServer.third.third",
	HandlerType: (*ThirdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PartLimit",
			Handler:    _Third_PartLimit_Handler,
		},
		{
			MethodName: "PartSize",
			Handler:    _Third_PartSize_Handler,
		},
		{
			MethodName: "InitiateMultipartUpload",
			Handler:    _Third_InitiateMultipartUpload_Handler,
		},
		{
			MethodName: "AuthSign",
			Handler:    _Third_AuthSign_Handler,
		},
		{
			MethodName: "CompleteMultipartUpload",
			Handler:    _Third_CompleteMultipartUpload_Handler,
		},
		{
			MethodName: "AccessURL",
			Handler:    _Third_AccessURL_Handler,
		},
		{
			MethodName: "InitiateFormData",
			Handler:    _Third_InitiateFormData_Handler,
		},
		{
			MethodName: "CompleteFormData",
			Handler:    _Third_CompleteFormData_Handler,
		},
		{
			MethodName: "FcmUpdateToken",
			Handler:    _Third_FcmUpdateToken_Handler,
		},
		{
			MethodName: "SetAppBadge",
			Handler:    _Third_SetAppBadge_Handler,
		},
		{
			MethodName: "UploadLogs",
			Handler:    _Third_UploadLogs_Handler,
		},
		{
			MethodName: "DeleteLogs",
			Handler:    _Third_DeleteLogs_Handler,
		},
		{
			MethodName: "SearchLogs",
			Handler:    _Third_SearchLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "third/third.proto",
}
