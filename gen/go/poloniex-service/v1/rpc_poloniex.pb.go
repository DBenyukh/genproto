// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: poloniex-service/v1/rpc_poloniex.proto

package poloniex_service

import (
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

type GetRtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ []string `protobuf:"bytes,1,rep,name=string,proto3" json:"string,omitempty"`
}

func (x *GetRtRequest) Reset() {
	*x = GetRtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRtRequest) ProtoMessage() {}

func (x *GetRtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRtRequest.ProtoReflect.Descriptor instead.
func (*GetRtRequest) Descriptor() ([]byte, []int) {
	return file_poloniex_service_v1_rpc_poloniex_proto_rawDescGZIP(), []int{0}
}

func (x *GetRtRequest) GetString_() []string {
	if x != nil {
		return x.String_
	}
	return nil
}

type GetRtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trade  *RecentTrade       `protobuf:"bytes,1,opt,name=trade,proto3" json:"trade,omitempty"`
	Status SubscriptionStatus `protobuf:"varint,2,opt,name=status,proto3,enum=poloniex.v1.SubscriptionStatus" json:"status,omitempty"`
}

func (x *GetRtResponse) Reset() {
	*x = GetRtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRtResponse) ProtoMessage() {}

func (x *GetRtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRtResponse.ProtoReflect.Descriptor instead.
func (*GetRtResponse) Descriptor() ([]byte, []int) {
	return file_poloniex_service_v1_rpc_poloniex_proto_rawDescGZIP(), []int{1}
}

func (x *GetRtResponse) GetTrade() *RecentTrade {
	if x != nil {
		return x.Trade
	}
	return nil
}

func (x *GetRtResponse) GetStatus() SubscriptionStatus {
	if x != nil {
		return x.Status
	}
	return SubscriptionStatus_SUBSCRIPTION_STATUS_UNSPECIFIED
}

type GetKlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ []string `protobuf:"bytes,1,rep,name=string,proto3" json:"string,omitempty"`
}

func (x *GetKlRequest) Reset() {
	*x = GetKlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKlRequest) ProtoMessage() {}

func (x *GetKlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKlRequest.ProtoReflect.Descriptor instead.
func (*GetKlRequest) Descriptor() ([]byte, []int) {
	return file_poloniex_service_v1_rpc_poloniex_proto_rawDescGZIP(), []int{2}
}

func (x *GetKlRequest) GetString_() []string {
	if x != nil {
		return x.String_
	}
	return nil
}

type GetKlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kline []*Kline `protobuf:"bytes,1,rep,name=kline,proto3" json:"kline,omitempty"`
}

func (x *GetKlResponse) Reset() {
	*x = GetKlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKlResponse) ProtoMessage() {}

func (x *GetKlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKlResponse.ProtoReflect.Descriptor instead.
func (*GetKlResponse) Descriptor() ([]byte, []int) {
	return file_poloniex_service_v1_rpc_poloniex_proto_rawDescGZIP(), []int{3}
}

func (x *GetKlResponse) GetKline() []*Kline {
	if x != nil {
		return x.Kline
	}
	return nil
}

var File_poloniex_service_v1_rpc_poloniex_proto protoreflect.FileDescriptor

var file_poloniex_service_v1_rpc_poloniex_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69,
	0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22,
	0x78, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x05, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1f, 0x2e, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x26, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4b, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x22, 0x39, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4b, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x05, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_poloniex_service_v1_rpc_poloniex_proto_rawDescOnce sync.Once
	file_poloniex_service_v1_rpc_poloniex_proto_rawDescData = file_poloniex_service_v1_rpc_poloniex_proto_rawDesc
)

func file_poloniex_service_v1_rpc_poloniex_proto_rawDescGZIP() []byte {
	file_poloniex_service_v1_rpc_poloniex_proto_rawDescOnce.Do(func() {
		file_poloniex_service_v1_rpc_poloniex_proto_rawDescData = protoimpl.X.CompressGZIP(file_poloniex_service_v1_rpc_poloniex_proto_rawDescData)
	})
	return file_poloniex_service_v1_rpc_poloniex_proto_rawDescData
}

var file_poloniex_service_v1_rpc_poloniex_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_poloniex_service_v1_rpc_poloniex_proto_goTypes = []interface{}{
	(*GetRtRequest)(nil),    // 0: poloniex.v1.GetRtRequest
	(*GetRtResponse)(nil),   // 1: poloniex.v1.GetRtResponse
	(*GetKlRequest)(nil),    // 2: poloniex.v1.GetKlRequest
	(*GetKlResponse)(nil),   // 3: poloniex.v1.GetKlResponse
	(*RecentTrade)(nil),     // 4: poloniex.v1.RecentTrade
	(SubscriptionStatus)(0), // 5: poloniex.v1.SubscriptionStatus
	(*Kline)(nil),           // 6: poloniex.v1.Kline
}
var file_poloniex_service_v1_rpc_poloniex_proto_depIdxs = []int32{
	4, // 0: poloniex.v1.GetRtResponse.trade:type_name -> poloniex.v1.RecentTrade
	5, // 1: poloniex.v1.GetRtResponse.status:type_name -> poloniex.v1.SubscriptionStatus
	6, // 2: poloniex.v1.GetKlResponse.kline:type_name -> poloniex.v1.Kline
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_poloniex_service_v1_rpc_poloniex_proto_init() }
func file_poloniex_service_v1_rpc_poloniex_proto_init() {
	if File_poloniex_service_v1_rpc_poloniex_proto != nil {
		return
	}
	file_poloniex_service_v1_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRtRequest); i {
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
		file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRtResponse); i {
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
		file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKlRequest); i {
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
		file_poloniex_service_v1_rpc_poloniex_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKlResponse); i {
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
			RawDescriptor: file_poloniex_service_v1_rpc_poloniex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_poloniex_service_v1_rpc_poloniex_proto_goTypes,
		DependencyIndexes: file_poloniex_service_v1_rpc_poloniex_proto_depIdxs,
		MessageInfos:      file_poloniex_service_v1_rpc_poloniex_proto_msgTypes,
	}.Build()
	File_poloniex_service_v1_rpc_poloniex_proto = out.File
	file_poloniex_service_v1_rpc_poloniex_proto_rawDesc = nil
	file_poloniex_service_v1_rpc_poloniex_proto_goTypes = nil
	file_poloniex_service_v1_rpc_poloniex_proto_depIdxs = nil
}
