// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: reg_server.proto

package main

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Nodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr  string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Weigh int32  `protobuf:"varint,2,opt,name=weigh,proto3" json:"weigh,omitempty"`
}

func (x *Nodes) Reset() {
	*x = Nodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nodes) ProtoMessage() {}

func (x *Nodes) ProtoReflect() protoreflect.Message {
	mi := &file_reg_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nodes.ProtoReflect.Descriptor instead.
func (*Nodes) Descriptor() ([]byte, []int) {
	return file_reg_server_proto_rawDescGZIP(), []int{0}
}

func (x *Nodes) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Nodes) GetWeigh() int32 {
	if x != nil {
		return x.Weigh
	}
	return 0
}

type Router struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Router) Reset() {
	*x = Router{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Router) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Router) ProtoMessage() {}

func (x *Router) ProtoReflect() protoreflect.Message {
	mi := &file_reg_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Router.ProtoReflect.Descriptor instead.
func (*Router) Descriptor() ([]byte, []int) {
	return file_reg_server_proto_rawDescGZIP(), []int{1}
}

func (x *Router) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc   string   `protobuf:"bytes,1,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Name   string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Nodes  []*Nodes `protobuf:"bytes,2,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
	Router *Router  `protobuf:"bytes,3,opt,name=Router,proto3" json:"Router,omitempty"`
}

func (x *ServerRequest) Reset() {
	*x = ServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerRequest) ProtoMessage() {}

func (x *ServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reg_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerRequest.ProtoReflect.Descriptor instead.
func (*ServerRequest) Descriptor() ([]byte, []int) {
	return file_reg_server_proto_rawDescGZIP(), []int{2}
}

func (x *ServerRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ServerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServerRequest) GetNodes() []*Nodes {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *ServerRequest) GetRouter() *Router {
	if x != nil {
		return x.Router
	}
	return nil
}

type ServerResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *ServerResult) Reset() {
	*x = ServerResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerResult) ProtoMessage() {}

func (x *ServerResult) ProtoReflect() protoreflect.Message {
	mi := &file_reg_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerResult.ProtoReflect.Descriptor instead.
func (*ServerResult) Descriptor() ([]byte, []int) {
	return file_reg_server_proto_rawDescGZIP(), []int{3}
}

func (x *ServerResult) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_reg_server_proto protoreflect.FileDescriptor

var file_reg_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x31, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x65, 0x69, 0x67, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x65, 0x69, 0x67, 0x68, 0x22, 0x1a, 0x0a, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x80, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x05, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0x51, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x47, 0x72, 0x70, 0x63, 0x12, 0x3b, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reg_server_proto_rawDescOnce sync.Once
	file_reg_server_proto_rawDescData = file_reg_server_proto_rawDesc
)

func file_reg_server_proto_rawDescGZIP() []byte {
	file_reg_server_proto_rawDescOnce.Do(func() {
		file_reg_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_reg_server_proto_rawDescData)
	})
	return file_reg_server_proto_rawDescData
}

var file_reg_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_reg_server_proto_goTypes = []interface{}{
	(*Nodes)(nil),         // 0: main.nodes
	(*Router)(nil),        // 1: main.router
	(*ServerRequest)(nil), // 2: main.ServerRequest
	(*ServerResult)(nil),  // 3: main.ServerResult
}
var file_reg_server_proto_depIdxs = []int32{
	0, // 0: main.ServerRequest.Nodes:type_name -> main.nodes
	1, // 1: main.ServerRequest.Router:type_name -> main.router
	2, // 2: main.ServerRegistryGrpc.ServerRegistry:input_type -> main.ServerRequest
	3, // 3: main.ServerRegistryGrpc.ServerRegistry:output_type -> main.ServerResult
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_reg_server_proto_init() }
func file_reg_server_proto_init() {
	if File_reg_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reg_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nodes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reg_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Router); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reg_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reg_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reg_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reg_server_proto_goTypes,
		DependencyIndexes: file_reg_server_proto_depIdxs,
		MessageInfos:      file_reg_server_proto_msgTypes,
	}.Build()
	File_reg_server_proto = out.File
	file_reg_server_proto_rawDesc = nil
	file_reg_server_proto_goTypes = nil
	file_reg_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerRegistryGrpcClient is the client API for ServerRegistryGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerRegistryGrpcClient interface {
	ServerRegistry(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*ServerResult, error)
}

type serverRegistryGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewServerRegistryGrpcClient(cc grpc.ClientConnInterface) ServerRegistryGrpcClient {
	return &serverRegistryGrpcClient{cc}
}

func (c *serverRegistryGrpcClient) ServerRegistry(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*ServerResult, error) {
	out := new(ServerResult)
	err := c.cc.Invoke(ctx, "/main.ServerRegistryGrpc/ServerRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerRegistryGrpcServer is the server API for ServerRegistryGrpc service.
type ServerRegistryGrpcServer interface {
	ServerRegistry(context.Context, *ServerRequest) (*ServerResult, error)
}

// UnimplementedServerRegistryGrpcServer can be embedded to have forward compatible implementations.
type UnimplementedServerRegistryGrpcServer struct {
}

func (*UnimplementedServerRegistryGrpcServer) ServerRegistry(context.Context, *ServerRequest) (*ServerResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerRegistry not implemented")
}

func RegisterServerRegistryGrpcServer(s *grpc.Server, srv ServerRegistryGrpcServer) {
	s.RegisterService(&_ServerRegistryGrpc_serviceDesc, srv)
}

func _ServerRegistryGrpc_ServerRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerRegistryGrpcServer).ServerRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ServerRegistryGrpc/ServerRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerRegistryGrpcServer).ServerRegistry(ctx, req.(*ServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerRegistryGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.ServerRegistryGrpc",
	HandlerType: (*ServerRegistryGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerRegistry",
			Handler:    _ServerRegistryGrpc_ServerRegistry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reg_server.proto",
}
