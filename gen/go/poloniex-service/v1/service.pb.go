// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: poloniex-service/v1/service.proto

package poloniex_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_poloniex_service_v1_service_proto protoreflect.FileDescriptor

var file_poloniex_service_v1_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6b, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x95, 0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x52, 0x54, 0x12,
	0x19, 0x2e, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x6c,
	0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x4b,
	0x4c, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4b, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f,
	0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x6f, 0x6c, 0x6f, 0x6e, 0x69, 0x65, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_poloniex_service_v1_service_proto_goTypes = []interface{}{
	(*GetRtRequest)(nil),  // 0: poloniex.v1.GetRtRequest
	(*GetKlRequest)(nil),  // 1: poloniex.v1.GetKlRequest
	(*GetRtResponse)(nil), // 2: poloniex.v1.GetRtResponse
	(*GetKlResponse)(nil), // 3: poloniex.v1.GetKlResponse
}
var file_poloniex_service_v1_service_proto_depIdxs = []int32{
	0, // 0: poloniex.v1.PoloniexService.GetRT:input_type -> poloniex.v1.GetRtRequest
	1, // 1: poloniex.v1.PoloniexService.GetKL:input_type -> poloniex.v1.GetKlRequest
	2, // 2: poloniex.v1.PoloniexService.GetRT:output_type -> poloniex.v1.GetRtResponse
	3, // 3: poloniex.v1.PoloniexService.GetKL:output_type -> poloniex.v1.GetKlResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_poloniex_service_v1_service_proto_init() }
func file_poloniex_service_v1_service_proto_init() {
	if File_poloniex_service_v1_service_proto != nil {
		return
	}
	file_poloniex_service_v1_rpc_rt_proto_init()
	file_poloniex_service_v1_rpc_kl_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_poloniex_service_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_poloniex_service_v1_service_proto_goTypes,
		DependencyIndexes: file_poloniex_service_v1_service_proto_depIdxs,
	}.Build()
	File_poloniex_service_v1_service_proto = out.File
	file_poloniex_service_v1_service_proto_rawDesc = nil
	file_poloniex_service_v1_service_proto_goTypes = nil
	file_poloniex_service_v1_service_proto_depIdxs = nil
}
