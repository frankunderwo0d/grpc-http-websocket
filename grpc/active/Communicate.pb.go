// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/Communicate.proto

package proto

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

type CRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtocolId uint32 `protobuf:"varint,1,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"` // 协议ID
	Data       []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`                                // 具体结构体序列化数据
}

func (x *CRequest) Reset() {
	*x = CRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Communicate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CRequest) ProtoMessage() {}

func (x *CRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Communicate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CRequest.ProtoReflect.Descriptor instead.
func (*CRequest) Descriptor() ([]byte, []int) {
	return file_proto_Communicate_proto_rawDescGZIP(), []int{0}
}

func (x *CRequest) GetProtocolId() uint32 {
	if x != nil {
		return x.ProtocolId
	}
	return 0
}

func (x *CRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码
	ProtocolId uint32 `protobuf:"varint,2,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"` // 协议ID
	Data       []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`                                // 具体结构体序列化数据
}

func (x *CResponse) Reset() {
	*x = CResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Communicate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CResponse) ProtoMessage() {}

func (x *CResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Communicate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CResponse.ProtoReflect.Descriptor instead.
func (*CResponse) Descriptor() ([]byte, []int) {
	return file_proto_Communicate_proto_rawDescGZIP(), []int{1}
}

func (x *CResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CResponse) GetProtocolId() uint32 {
	if x != nil {
		return x.ProtocolId
	}
	return 0
}

func (x *CResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_Communicate_proto protoreflect.FileDescriptor

var file_proto_Communicate_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3f, 0x0a, 0x08, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x61, 0x0a, 0x09, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x3c, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_Communicate_proto_rawDescOnce sync.Once
	file_proto_Communicate_proto_rawDescData = file_proto_Communicate_proto_rawDesc
)

func file_proto_Communicate_proto_rawDescGZIP() []byte {
	file_proto_Communicate_proto_rawDescOnce.Do(func() {
		file_proto_Communicate_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Communicate_proto_rawDescData)
	})
	return file_proto_Communicate_proto_rawDescData
}

var file_proto_Communicate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_Communicate_proto_goTypes = []interface{}{
	(*CRequest)(nil),  // 0: proto.CRequest
	(*CResponse)(nil), // 1: proto.CResponse
}
var file_proto_Communicate_proto_depIdxs = []int32{
	0, // 0: proto.Communication.Send:input_type -> proto.CRequest
	1, // 1: proto.Communication.Send:output_type -> proto.CResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_Communicate_proto_init() }
func file_proto_Communicate_proto_init() {
	if File_proto_Communicate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_Communicate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CRequest); i {
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
		file_proto_Communicate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CResponse); i {
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
			RawDescriptor: file_proto_Communicate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_Communicate_proto_goTypes,
		DependencyIndexes: file_proto_Communicate_proto_depIdxs,
		MessageInfos:      file_proto_Communicate_proto_msgTypes,
	}.Build()
	File_proto_Communicate_proto = out.File
	file_proto_Communicate_proto_rawDesc = nil
	file_proto_Communicate_proto_goTypes = nil
	file_proto_Communicate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommunicationClient is the client API for Communication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommunicationClient interface {
	Send(ctx context.Context, in *CRequest, opts ...grpc.CallOption) (*CResponse, error)
}

type communicationClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationClient(cc grpc.ClientConnInterface) CommunicationClient {
	return &communicationClient{cc}
}

func (c *communicationClient) Send(ctx context.Context, in *CRequest, opts ...grpc.CallOption) (*CResponse, error) {
	out := new(CResponse)
	err := c.cc.Invoke(ctx, "/proto.Communication/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationServer is the server API for Communication service.
type CommunicationServer interface {
	Send(context.Context, *CRequest) (*CResponse, error)
}

// UnimplementedCommunicationServer can be embedded to have forward compatible implementations.
type UnimplementedCommunicationServer struct {
}

func (*UnimplementedCommunicationServer) Send(context.Context, *CRequest) (*CResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterCommunicationServer(s *grpc.Server, srv CommunicationServer) {
	s.RegisterService(&_Communication_serviceDesc, srv)
}

func _Communication_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Communication/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).Send(ctx, req.(*CRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Communication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Communication",
	HandlerType: (*CommunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Communication_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Communicate.proto",
}