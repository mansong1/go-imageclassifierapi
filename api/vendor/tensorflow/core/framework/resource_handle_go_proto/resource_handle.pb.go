// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: tensorflow/core/framework/resource_handle.proto

package tensorflow

import (
	proto "github.com/golang/protobuf/proto"
	tensor_shape_go_proto "tensorflow/core/framework/tensor_shape_go_proto"
	types_go_proto "tensorflow/core/framework/types_go_proto"
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

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandleProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode,proto3" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName string `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName,proto3" json:"maybe_type_name,omitempty"`
	// Data types and shapes for the underlying resource.
	DtypesAndShapes []*ResourceHandleProto_DtypeAndShape `protobuf:"bytes,6,rep,name=dtypes_and_shapes,json=dtypesAndShapes,proto3" json:"dtypes_and_shapes,omitempty"`
	// A set of devices containing the resource. If empty, the resource only
	// exists on `device`.
	AllowedDevices []string `protobuf:"bytes,7,rep,name=allowed_devices,json=allowedDevices,proto3" json:"allowed_devices,omitempty"`
}

func (x *ResourceHandleProto) Reset() {
	*x = ResourceHandleProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_resource_handle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceHandleProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceHandleProto) ProtoMessage() {}

func (x *ResourceHandleProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_resource_handle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceHandleProto.ProtoReflect.Descriptor instead.
func (*ResourceHandleProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_resource_handle_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceHandleProto) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *ResourceHandleProto) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

func (x *ResourceHandleProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceHandleProto) GetHashCode() uint64 {
	if x != nil {
		return x.HashCode
	}
	return 0
}

func (x *ResourceHandleProto) GetMaybeTypeName() string {
	if x != nil {
		return x.MaybeTypeName
	}
	return ""
}

func (x *ResourceHandleProto) GetDtypesAndShapes() []*ResourceHandleProto_DtypeAndShape {
	if x != nil {
		return x.DtypesAndShapes
	}
	return nil
}

func (x *ResourceHandleProto) GetAllowedDevices() []string {
	if x != nil {
		return x.AllowedDevices
	}
	return nil
}

// Protocol buffer representing a pair of (data type, tensor shape).
type ResourceHandleProto_DtypeAndShape struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dtype types_go_proto.DataType                 `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
}

func (x *ResourceHandleProto_DtypeAndShape) Reset() {
	*x = ResourceHandleProto_DtypeAndShape{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_resource_handle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceHandleProto_DtypeAndShape) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceHandleProto_DtypeAndShape) ProtoMessage() {}

func (x *ResourceHandleProto_DtypeAndShape) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_resource_handle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceHandleProto_DtypeAndShape.ProtoReflect.Descriptor instead.
func (*ResourceHandleProto_DtypeAndShape) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_resource_handle_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ResourceHandleProto_DtypeAndShape) GetDtype() types_go_proto.DataType {
	if x != nil {
		return x.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (x *ResourceHandleProto_DtypeAndShape) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if x != nil {
		return x.Shape
	}
	return nil
}

var File_tensorflow_core_framework_resource_handle_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_resource_handle_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2c, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f,
	0x73, 0x68, 0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x99, 0x03, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x79, 0x62, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x79, 0x62,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x11, 0x64, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x74, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x68,
	0x61, 0x70, 0x65, 0x52, 0x0f, 0x64, 0x74, 0x79, 0x70, 0x65, 0x73, 0x41, 0x6e, 0x64, 0x53, 0x68,
	0x61, 0x70, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x6f, 0x0a,
	0x0d, 0x44, 0x74, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61,
	0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x42, 0x87,
	0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x50, 0x01, 0x5a, 0x56, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x67, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_resource_handle_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_resource_handle_proto_rawDescData = file_tensorflow_core_framework_resource_handle_proto_rawDesc
)

func file_tensorflow_core_framework_resource_handle_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_resource_handle_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_resource_handle_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_resource_handle_proto_rawDescData)
	})
	return file_tensorflow_core_framework_resource_handle_proto_rawDescData
}

var file_tensorflow_core_framework_resource_handle_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_framework_resource_handle_proto_goTypes = []interface{}{
	(*ResourceHandleProto)(nil),                    // 0: tensorflow.ResourceHandleProto
	(*ResourceHandleProto_DtypeAndShape)(nil),      // 1: tensorflow.ResourceHandleProto.DtypeAndShape
	(types_go_proto.DataType)(0),                   // 2: tensorflow.DataType
	(*tensor_shape_go_proto.TensorShapeProto)(nil), // 3: tensorflow.TensorShapeProto
}
var file_tensorflow_core_framework_resource_handle_proto_depIdxs = []int32{
	1, // 0: tensorflow.ResourceHandleProto.dtypes_and_shapes:type_name -> tensorflow.ResourceHandleProto.DtypeAndShape
	2, // 1: tensorflow.ResourceHandleProto.DtypeAndShape.dtype:type_name -> tensorflow.DataType
	3, // 2: tensorflow.ResourceHandleProto.DtypeAndShape.shape:type_name -> tensorflow.TensorShapeProto
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_resource_handle_proto_init() }
func file_tensorflow_core_framework_resource_handle_proto_init() {
	if File_tensorflow_core_framework_resource_handle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_resource_handle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceHandleProto); i {
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
		file_tensorflow_core_framework_resource_handle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceHandleProto_DtypeAndShape); i {
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
			RawDescriptor: file_tensorflow_core_framework_resource_handle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_resource_handle_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_resource_handle_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_resource_handle_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_resource_handle_proto = out.File
	file_tensorflow_core_framework_resource_handle_proto_rawDesc = nil
	file_tensorflow_core_framework_resource_handle_proto_goTypes = nil
	file_tensorflow_core_framework_resource_handle_proto_depIdxs = nil
}
