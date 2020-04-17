// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: tensorflow_serving/config/monitoring_config.proto

package tensorflow_serving

import (
	proto "github.com/golang/protobuf/proto"
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

// Configuration for Prometheus monitoring.
type PrometheusConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether to expose Prometheus metrics.
	Enable bool `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	// The endpoint to expose Prometheus metrics.
	// If not specified, PrometheusExporter::kPrometheusPath value is used.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PrometheusConfig) Reset() {
	*x = PrometheusConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_monitoring_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusConfig) ProtoMessage() {}

func (x *PrometheusConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_monitoring_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusConfig.ProtoReflect.Descriptor instead.
func (*PrometheusConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_monitoring_config_proto_rawDescGZIP(), []int{0}
}

func (x *PrometheusConfig) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *PrometheusConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Configuration for monitoring.
type MonitoringConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrometheusConfig *PrometheusConfig `protobuf:"bytes,1,opt,name=prometheus_config,json=prometheusConfig,proto3" json:"prometheus_config,omitempty"`
}

func (x *MonitoringConfig) Reset() {
	*x = MonitoringConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_monitoring_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitoringConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitoringConfig) ProtoMessage() {}

func (x *MonitoringConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_monitoring_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitoringConfig.ProtoReflect.Descriptor instead.
func (*MonitoringConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_monitoring_config_proto_rawDescGZIP(), []int{1}
}

func (x *MonitoringConfig) GetPrometheusConfig() *PrometheusConfig {
	if x != nil {
		return x.PrometheusConfig
	}
	return nil
}

var File_tensorflow_serving_config_monitoring_config_proto protoreflect.FileDescriptor

var file_tensorflow_serving_config_monitoring_config_proto_rawDesc = []byte{
	0x0a, 0x31, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x22, 0x3e, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x65, 0x0a, 0x10, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a, 0x11, 0x70,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x70, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03,
	0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_config_monitoring_config_proto_rawDescOnce sync.Once
	file_tensorflow_serving_config_monitoring_config_proto_rawDescData = file_tensorflow_serving_config_monitoring_config_proto_rawDesc
)

func file_tensorflow_serving_config_monitoring_config_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_config_monitoring_config_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_config_monitoring_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_config_monitoring_config_proto_rawDescData)
	})
	return file_tensorflow_serving_config_monitoring_config_proto_rawDescData
}

var file_tensorflow_serving_config_monitoring_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_serving_config_monitoring_config_proto_goTypes = []interface{}{
	(*PrometheusConfig)(nil), // 0: tensorflow.serving.PrometheusConfig
	(*MonitoringConfig)(nil), // 1: tensorflow.serving.MonitoringConfig
}
var file_tensorflow_serving_config_monitoring_config_proto_depIdxs = []int32{
	0, // 0: tensorflow.serving.MonitoringConfig.prometheus_config:type_name -> tensorflow.serving.PrometheusConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_config_monitoring_config_proto_init() }
func file_tensorflow_serving_config_monitoring_config_proto_init() {
	if File_tensorflow_serving_config_monitoring_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_config_monitoring_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusConfig); i {
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
		file_tensorflow_serving_config_monitoring_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitoringConfig); i {
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
			RawDescriptor: file_tensorflow_serving_config_monitoring_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_config_monitoring_config_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_config_monitoring_config_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_config_monitoring_config_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_config_monitoring_config_proto = out.File
	file_tensorflow_serving_config_monitoring_config_proto_rawDesc = nil
	file_tensorflow_serving_config_monitoring_config_proto_goTypes = nil
	file_tensorflow_serving_config_monitoring_config_proto_depIdxs = nil
}
