// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: api/proto/keeper.proto

package api

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

type AskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile     *Profile     `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	ShopSummary *ShopSummary `protobuf:"bytes,2,opt,name=shop_summary,json=shopSummary,proto3" json:"shop_summary,omitempty"`
	PaymentInfo *PaymentInfo `protobuf:"bytes,3,opt,name=payment_info,json=paymentInfo,proto3" json:"payment_info,omitempty"`
}

func (x *AskRequest) Reset() {
	*x = AskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_keeper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskRequest) ProtoMessage() {}

func (x *AskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_keeper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskRequest.ProtoReflect.Descriptor instead.
func (*AskRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_keeper_proto_rawDescGZIP(), []int{0}
}

func (x *AskRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *AskRequest) GetShopSummary() *ShopSummary {
	if x != nil {
		return x.ShopSummary
	}
	return nil
}

func (x *AskRequest) GetPaymentInfo() *PaymentInfo {
	if x != nil {
		return x.PaymentInfo
	}
	return nil
}

type AskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowPass bool  `protobuf:"varint,1,opt,name=allow_pass,json=allowPass,proto3" json:"allow_pass,omitempty"`
	Score     int32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *AskResponse) Reset() {
	*x = AskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_keeper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskResponse) ProtoMessage() {}

func (x *AskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_keeper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskResponse.ProtoReflect.Descriptor instead.
func (*AskResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_keeper_proto_rawDescGZIP(), []int{1}
}

func (x *AskResponse) GetAllowPass() bool {
	if x != nil {
		return x.AllowPass
	}
	return false
}

func (x *AskResponse) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_api_proto_keeper_proto protoreflect.FileDescriptor

var file_api_proto_keeper_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0a, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x73, 0x68, 0x6f,
	0x70, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x70, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x33,
	0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x42, 0x0a, 0x0b, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x32, 0x34, 0x0a, 0x06, 0x4b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x73, 0x6b, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a,
	0x07, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_keeper_proto_rawDescOnce sync.Once
	file_api_proto_keeper_proto_rawDescData = file_api_proto_keeper_proto_rawDesc
)

func file_api_proto_keeper_proto_rawDescGZIP() []byte {
	file_api_proto_keeper_proto_rawDescOnce.Do(func() {
		file_api_proto_keeper_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_keeper_proto_rawDescData)
	})
	return file_api_proto_keeper_proto_rawDescData
}

var file_api_proto_keeper_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_keeper_proto_goTypes = []interface{}{
	(*AskRequest)(nil),  // 0: api.AskRequest
	(*AskResponse)(nil), // 1: api.AskResponse
	(*Profile)(nil),     // 2: api.Profile
	(*ShopSummary)(nil), // 3: api.ShopSummary
	(*PaymentInfo)(nil), // 4: api.PaymentInfo
}
var file_api_proto_keeper_proto_depIdxs = []int32{
	2, // 0: api.AskRequest.profile:type_name -> api.Profile
	3, // 1: api.AskRequest.shop_summary:type_name -> api.ShopSummary
	4, // 2: api.AskRequest.payment_info:type_name -> api.PaymentInfo
	0, // 3: api.Keeper.Ask:input_type -> api.AskRequest
	1, // 4: api.Keeper.Ask:output_type -> api.AskResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_keeper_proto_init() }
func file_api_proto_keeper_proto_init() {
	if File_api_proto_keeper_proto != nil {
		return
	}
	file_api_proto_profile_proto_init()
	file_api_proto_shop_summary_proto_init()
	file_api_proto_payment_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_keeper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskRequest); i {
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
		file_api_proto_keeper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskResponse); i {
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
			RawDescriptor: file_api_proto_keeper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_keeper_proto_goTypes,
		DependencyIndexes: file_api_proto_keeper_proto_depIdxs,
		MessageInfos:      file_api_proto_keeper_proto_msgTypes,
	}.Build()
	File_api_proto_keeper_proto = out.File
	file_api_proto_keeper_proto_rawDesc = nil
	file_api_proto_keeper_proto_goTypes = nil
	file_api_proto_keeper_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeeperClient is the client API for Keeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeeperClient interface {
	Ask(ctx context.Context, in *AskRequest, opts ...grpc.CallOption) (*AskResponse, error)
}

type keeperClient struct {
	cc grpc.ClientConnInterface
}

func NewKeeperClient(cc grpc.ClientConnInterface) KeeperClient {
	return &keeperClient{cc}
}

func (c *keeperClient) Ask(ctx context.Context, in *AskRequest, opts ...grpc.CallOption) (*AskResponse, error) {
	out := new(AskResponse)
	err := c.cc.Invoke(ctx, "/api.Keeper/Ask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeeperServer is the server API for Keeper service.
type KeeperServer interface {
	Ask(context.Context, *AskRequest) (*AskResponse, error)
}

// UnimplementedKeeperServer can be embedded to have forward compatible implementations.
type UnimplementedKeeperServer struct {
}

func (*UnimplementedKeeperServer) Ask(context.Context, *AskRequest) (*AskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ask not implemented")
}

func RegisterKeeperServer(s *grpc.Server, srv KeeperServer) {
	s.RegisterService(&_Keeper_serviceDesc, srv)
}

func _Keeper_Ask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServer).Ask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Keeper/Ask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServer).Ask(ctx, req.(*AskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Keeper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Keeper",
	HandlerType: (*KeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ask",
			Handler:    _Keeper_Ask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/keeper.proto",
}
