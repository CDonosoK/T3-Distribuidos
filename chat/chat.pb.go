// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: chat.proto

package chat

import (
	context "context"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta  string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
	Ciudad   string `protobuf:"bytes,2,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
	Valor    string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tipo     string `protobuf:"bytes,4,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Servidor int32  `protobuf:"varint,5,opt,name=servidor,proto3" json:"servidor,omitempty"`
	X        string `protobuf:"bytes,6,opt,name=X,proto3" json:"X,omitempty"`
	Y        string `protobuf:"bytes,7,opt,name=Y,proto3" json:"Y,omitempty"`
	Z        string `protobuf:"bytes,8,opt,name=Z,proto3" json:"Z,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *Message) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *Message) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *Message) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Message) GetServidor() int32 {
	if x != nil {
		return x.Servidor
	}
	return 0
}

func (x *Message) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

func (x *Message) GetY() string {
	if x != nil {
		return x.Y
	}
	return ""
}

func (x *Message) GetZ() string {
	if x != nil {
		return x.Z
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0xab, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x75, 0x64,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x59, 0x12, 0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x5a,
	0x32, 0xd9, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x30, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x43, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chat_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: chat.Message
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: chat.Chat.AddCityMessage:input_type -> chat.Message
	0, // 1: chat.Chat.UpdateNameMessage:input_type -> chat.Message
	0, // 2: chat.Chat.UpdateNumberMessage:input_type -> chat.Message
	0, // 3: chat.Chat.DeleteCityMessage:input_type -> chat.Message
	0, // 4: chat.Chat.AddCityMessage:output_type -> chat.Message
	0, // 5: chat.Chat.UpdateNameMessage:output_type -> chat.Message
	0, // 6: chat.Chat.UpdateNumberMessage:output_type -> chat.Message
	0, // 7: chat.Chat.DeleteCityMessage:output_type -> chat.Message
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	AddCityMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	UpdateNameMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	UpdateNumberMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	DeleteCityMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) AddCityMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.Chat/AddCityMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) UpdateNameMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.Chat/UpdateNameMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) UpdateNumberMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.Chat/UpdateNumberMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) DeleteCityMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.Chat/DeleteCityMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	AddCityMessage(context.Context, *Message) (*Message, error)
	UpdateNameMessage(context.Context, *Message) (*Message, error)
	UpdateNumberMessage(context.Context, *Message) (*Message, error)
	DeleteCityMessage(context.Context, *Message) (*Message, error)
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) AddCityMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCityMessage not implemented")
}
func (*UnimplementedChatServer) UpdateNameMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNameMessage not implemented")
}
func (*UnimplementedChatServer) UpdateNumberMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNumberMessage not implemented")
}
func (*UnimplementedChatServer) DeleteCityMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCityMessage not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_AddCityMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).AddCityMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/AddCityMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).AddCityMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_UpdateNameMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).UpdateNameMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/UpdateNameMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).UpdateNameMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_UpdateNumberMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).UpdateNumberMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/UpdateNumberMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).UpdateNumberMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_DeleteCityMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).DeleteCityMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/DeleteCityMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).DeleteCityMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCityMessage",
			Handler:    _Chat_AddCityMessage_Handler,
		},
		{
			MethodName: "UpdateNameMessage",
			Handler:    _Chat_UpdateNameMessage_Handler,
		},
		{
			MethodName: "UpdateNumberMessage",
			Handler:    _Chat_UpdateNumberMessage_Handler,
		},
		{
			MethodName: "DeleteCityMessage",
			Handler:    _Chat_DeleteCityMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
