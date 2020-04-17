// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: tensorflow/core/framework/op_def.proto

package op_def_go_proto

import (
	proto "github.com/golang/protobuf/proto"
	attr_value_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/attr_value_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
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

// Defines an operation. A NodeDef in a GraphDef specifies an Op by
// using the "op" field which should match the name of a OpDef.
// LINT.IfChange
type OpDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Op names starting with an underscore are reserved for internal use.
	// Names should be CamelCase and match the regexp "[A-Z][a-zA-Z0-9>_]*".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the input(s).
	InputArg []*OpDef_ArgDef `protobuf:"bytes,2,rep,name=input_arg,json=inputArg,proto3" json:"input_arg,omitempty"`
	// Description of the output(s).
	OutputArg []*OpDef_ArgDef `protobuf:"bytes,3,rep,name=output_arg,json=outputArg,proto3" json:"output_arg,omitempty"`
	// Named control outputs for this operation. Useful only for composite
	// operations (i.e. functions) which want to name different control outputs.
	ControlOutput []string         `protobuf:"bytes,20,rep,name=control_output,json=controlOutput,proto3" json:"control_output,omitempty"`
	Attr          []*OpDef_AttrDef `protobuf:"bytes,4,rep,name=attr,proto3" json:"attr,omitempty"`
	// Optional deprecation based on GraphDef versions.
	Deprecation *OpDeprecation `protobuf:"bytes,8,opt,name=deprecation,proto3" json:"deprecation,omitempty"`
	// One-line human-readable description of what the Op does.
	Summary string `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	// Additional, longer human-readable description of what the Op does.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// True if the operation is commutative ("op(a,b) == op(b,a)" for all inputs)
	IsCommutative bool `protobuf:"varint,18,opt,name=is_commutative,json=isCommutative,proto3" json:"is_commutative,omitempty"`
	// If is_aggregate is true, then this operation accepts N >= 2
	// inputs and produces 1 output all of the same type.  Should be
	// associative and commutative, and produce output with the same
	// shape as the input.  The optimizer may replace an aggregate op
	// taking input from multiple devices with a tree of aggregate ops
	// that aggregate locally within each device (and possibly within
	// groups of nearby devices) before communicating.
	// TODO(josh11b): Implement that optimization.
	IsAggregate bool `protobuf:"varint,16,opt,name=is_aggregate,json=isAggregate,proto3" json:"is_aggregate,omitempty"` // for things like add
	// Ops are marked as stateful if their behavior depends on some state beyond
	// their input tensors (e.g. variable reading op) or if they have
	// a side-effect (e.g. printing or asserting ops). Equivalently, stateless ops
	// must always produce the same output for the same input and have
	// no side-effects.
	//
	// By default Ops may be moved between devices.  Stateful ops should
	// either not be moved, or should only be moved if that state can also
	// be moved (e.g. via some sort of save / restore).
	// Stateful ops are guaranteed to never be optimized away by Common
	// Subexpression Elimination (CSE).
	IsStateful bool `protobuf:"varint,17,opt,name=is_stateful,json=isStateful,proto3" json:"is_stateful,omitempty"` // for things like variables, queue
	// By default, all inputs to an Op must be initialized Tensors.  Ops
	// that may initialize tensors for the first time should set this
	// field to true, to allow the Op to take an uninitialized Tensor as
	// input.
	AllowsUninitializedInput bool `protobuf:"varint,19,opt,name=allows_uninitialized_input,json=allowsUninitializedInput,proto3" json:"allows_uninitialized_input,omitempty"` // for Assign, etc.
}

func (x *OpDef) Reset() {
	*x = OpDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_op_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpDef) ProtoMessage() {}

func (x *OpDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_op_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpDef.ProtoReflect.Descriptor instead.
func (*OpDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_op_def_proto_rawDescGZIP(), []int{0}
}

func (x *OpDef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OpDef) GetInputArg() []*OpDef_ArgDef {
	if x != nil {
		return x.InputArg
	}
	return nil
}

func (x *OpDef) GetOutputArg() []*OpDef_ArgDef {
	if x != nil {
		return x.OutputArg
	}
	return nil
}

func (x *OpDef) GetControlOutput() []string {
	if x != nil {
		return x.ControlOutput
	}
	return nil
}

func (x *OpDef) GetAttr() []*OpDef_AttrDef {
	if x != nil {
		return x.Attr
	}
	return nil
}

func (x *OpDef) GetDeprecation() *OpDeprecation {
	if x != nil {
		return x.Deprecation
	}
	return nil
}

func (x *OpDef) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *OpDef) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OpDef) GetIsCommutative() bool {
	if x != nil {
		return x.IsCommutative
	}
	return false
}

func (x *OpDef) GetIsAggregate() bool {
	if x != nil {
		return x.IsAggregate
	}
	return false
}

func (x *OpDef) GetIsStateful() bool {
	if x != nil {
		return x.IsStateful
	}
	return false
}

func (x *OpDef) GetAllowsUninitializedInput() bool {
	if x != nil {
		return x.AllowsUninitializedInput
	}
	return false
}

// Information about version-dependent deprecation of an op
type OpDeprecation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// First GraphDef version at which the op is disallowed.
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Explanation of why it was deprecated and what to use instead.
	Explanation string `protobuf:"bytes,2,opt,name=explanation,proto3" json:"explanation,omitempty"`
}

func (x *OpDeprecation) Reset() {
	*x = OpDeprecation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_op_def_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpDeprecation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpDeprecation) ProtoMessage() {}

func (x *OpDeprecation) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_op_def_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpDeprecation.ProtoReflect.Descriptor instead.
func (*OpDeprecation) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_op_def_proto_rawDescGZIP(), []int{1}
}

func (x *OpDeprecation) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *OpDeprecation) GetExplanation() string {
	if x != nil {
		return x.Explanation
	}
	return ""
}

// A collection of OpDefs
type OpList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op []*OpDef `protobuf:"bytes,1,rep,name=op,proto3" json:"op,omitempty"`
}

func (x *OpList) Reset() {
	*x = OpList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_op_def_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpList) ProtoMessage() {}

func (x *OpList) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_op_def_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpList.ProtoReflect.Descriptor instead.
func (*OpList) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_op_def_proto_rawDescGZIP(), []int{2}
}

func (x *OpList) GetOp() []*OpDef {
	if x != nil {
		return x.Op
	}
	return nil
}

// For describing inputs and outputs.
type OpDef_ArgDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name for the input/output.  Should match the regexp "[a-z][a-z0-9_]*".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Human readable description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Describes the type of one or more tensors that are accepted/produced
	// by this input/output arg.  The only legal combinations are:
	// * For a single tensor: either the "type" field is set or the
	//   "type_attr" field is set to the name of an attr with type "type".
	// * For a sequence of tensors with the same type: the "number_attr"
	//   field will be set to the name of an attr with type "int", and
	//   either the "type" or "type_attr" field will be set as for
	//   single tensors.
	// * For a sequence of tensors, the "type_list_attr" field will be set
	//   to the name of an attr with type "list(type)".
	Type       types_go_proto.DataType `protobuf:"varint,3,opt,name=type,proto3,enum=tensorflow.DataType" json:"type,omitempty"`
	TypeAttr   string                  `protobuf:"bytes,4,opt,name=type_attr,json=typeAttr,proto3" json:"type_attr,omitempty"`       // if specified, attr must have type "type"
	NumberAttr string                  `protobuf:"bytes,5,opt,name=number_attr,json=numberAttr,proto3" json:"number_attr,omitempty"` // if specified, attr must have type "int"
	// If specified, attr must have type "list(type)", and none of
	// type, type_attr, and number_attr may be specified.
	TypeListAttr string `protobuf:"bytes,6,opt,name=type_list_attr,json=typeListAttr,proto3" json:"type_list_attr,omitempty"`
	// For inputs: if true, the inputs are required to be refs.
	//   By default, inputs can be either refs or non-refs.
	// For outputs: if true, outputs are refs, otherwise they are not.
	IsRef bool `protobuf:"varint,16,opt,name=is_ref,json=isRef,proto3" json:"is_ref,omitempty"`
}

func (x *OpDef_ArgDef) Reset() {
	*x = OpDef_ArgDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_op_def_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpDef_ArgDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpDef_ArgDef) ProtoMessage() {}

func (x *OpDef_ArgDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_op_def_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpDef_ArgDef.ProtoReflect.Descriptor instead.
func (*OpDef_ArgDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_op_def_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OpDef_ArgDef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OpDef_ArgDef) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OpDef_ArgDef) GetType() types_go_proto.DataType {
	if x != nil {
		return x.Type
	}
	return types_go_proto.DataType_DT_INVALID
}

func (x *OpDef_ArgDef) GetTypeAttr() string {
	if x != nil {
		return x.TypeAttr
	}
	return ""
}

func (x *OpDef_ArgDef) GetNumberAttr() string {
	if x != nil {
		return x.NumberAttr
	}
	return ""
}

func (x *OpDef_ArgDef) GetTypeListAttr() string {
	if x != nil {
		return x.TypeListAttr
	}
	return ""
}

func (x *OpDef_ArgDef) GetIsRef() bool {
	if x != nil {
		return x.IsRef
	}
	return false
}

// Description of the graph-construction-time configuration of this
// Op.  That is to say, this describes the attr fields that will
// be specified in the NodeDef.
type OpDef_AttrDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A descriptive name for the argument.  May be used, e.g. by the
	// Python client, as a keyword argument name, and so should match
	// the regexp "[a-z][a-z0-9_]+".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// One of the type names from attr_value.proto ("string", "list(string)",
	// "int", etc.).
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// A reasonable default for this attribute if the user does not supply
	// a value.  If not specified, the user must supply a value.
	DefaultValue *attr_value_go_proto.AttrValue `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	// Human-readable description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// For type == "int", this is a minimum value.  For "list(___)"
	// types, this is the minimum length.
	HasMinimum bool  `protobuf:"varint,5,opt,name=has_minimum,json=hasMinimum,proto3" json:"has_minimum,omitempty"`
	Minimum    int64 `protobuf:"varint,6,opt,name=minimum,proto3" json:"minimum,omitempty"`
	// The set of allowed values.  Has type that is the "list" version
	// of the "type" field above (uses the "list" field of AttrValue).
	// If type == "type" or "list(type)" above, then the "type" field
	// of "allowed_values.list" has the set of allowed DataTypes.
	// If type == "string" or "list(string)", then the "s" field of
	// "allowed_values.list" has the set of allowed strings.
	AllowedValues *attr_value_go_proto.AttrValue `protobuf:"bytes,7,opt,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
}

func (x *OpDef_AttrDef) Reset() {
	*x = OpDef_AttrDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_op_def_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpDef_AttrDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpDef_AttrDef) ProtoMessage() {}

func (x *OpDef_AttrDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_op_def_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpDef_AttrDef.ProtoReflect.Descriptor instead.
func (*OpDef_AttrDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_op_def_proto_rawDescGZIP(), []int{0, 1}
}

func (x *OpDef_AttrDef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OpDef_AttrDef) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OpDef_AttrDef) GetDefaultValue() *attr_value_go_proto.AttrValue {
	if x != nil {
		return x.DefaultValue
	}
	return nil
}

func (x *OpDef_AttrDef) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OpDef_AttrDef) GetHasMinimum() bool {
	if x != nil {
		return x.HasMinimum
	}
	return false
}

func (x *OpDef_AttrDef) GetMinimum() int64 {
	if x != nil {
		return x.Minimum
	}
	return 0
}

func (x *OpDef_AttrDef) GetAllowedValues() *attr_value_go_proto.AttrValue {
	if x != nil {
		return x.AllowedValues
	}
	return nil
}

var File_tensorflow_core_framework_op_def_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_op_def_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x70, 0x5f, 0x64,
	0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x61, 0x74, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x07, 0x0a, 0x05, 0x4f, 0x70, 0x44, 0x65,
	0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x61,
	0x72, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4f, 0x70, 0x44, 0x65, 0x66, 0x2e, 0x41, 0x72, 0x67, 0x44,
	0x65, 0x66, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x72, 0x67, 0x12, 0x37, 0x0a, 0x0a,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4f, 0x70,
	0x44, 0x65, 0x66, 0x2e, 0x41, 0x72, 0x67, 0x44, 0x65, 0x66, 0x52, 0x09, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x41, 0x72, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2d, 0x0a, 0x04,
	0x61, 0x74, 0x74, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4f, 0x70, 0x44, 0x65, 0x66, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x44, 0x65, 0x66, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x64,
	0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4f, 0x70,
	0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x12,
	0x3c, 0x0a, 0x1a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x18, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x55, 0x6e, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0xe3, 0x01,
	0x0a, 0x06, 0x41, 0x72, 0x67, 0x44, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x15, 0x0a, 0x06,
	0x69, 0x73, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73,
	0x52, 0x65, 0x66, 0x1a, 0x88, 0x02, 0x0a, 0x07, 0x41, 0x74, 0x74, 0x72, 0x44, 0x65, 0x66, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x69, 0x6e,
	0x69, 0x6d, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x4d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d,
	0x12, 0x3c, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x4b,
	0x0a, 0x0d, 0x4f, 0x70, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70,
	0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x06, 0x4f,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4f,
	0x70, 0x44, 0x65, 0x66, 0x52, 0x02, 0x6f, 0x70, 0x42, 0x7b, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0b, 0x4f, 0x70, 0x44, 0x65, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x6f, 0x70, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_op_def_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_op_def_proto_rawDescData = file_tensorflow_core_framework_op_def_proto_rawDesc
)

func file_tensorflow_core_framework_op_def_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_op_def_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_op_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_op_def_proto_rawDescData)
	})
	return file_tensorflow_core_framework_op_def_proto_rawDescData
}

var file_tensorflow_core_framework_op_def_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tensorflow_core_framework_op_def_proto_goTypes = []interface{}{
	(*OpDef)(nil),                         // 0: tensorflow.OpDef
	(*OpDeprecation)(nil),                 // 1: tensorflow.OpDeprecation
	(*OpList)(nil),                        // 2: tensorflow.OpList
	(*OpDef_ArgDef)(nil),                  // 3: tensorflow.OpDef.ArgDef
	(*OpDef_AttrDef)(nil),                 // 4: tensorflow.OpDef.AttrDef
	(types_go_proto.DataType)(0),          // 5: tensorflow.DataType
	(*attr_value_go_proto.AttrValue)(nil), // 6: tensorflow.AttrValue
}
var file_tensorflow_core_framework_op_def_proto_depIdxs = []int32{
	3, // 0: tensorflow.OpDef.input_arg:type_name -> tensorflow.OpDef.ArgDef
	3, // 1: tensorflow.OpDef.output_arg:type_name -> tensorflow.OpDef.ArgDef
	4, // 2: tensorflow.OpDef.attr:type_name -> tensorflow.OpDef.AttrDef
	1, // 3: tensorflow.OpDef.deprecation:type_name -> tensorflow.OpDeprecation
	0, // 4: tensorflow.OpList.op:type_name -> tensorflow.OpDef
	5, // 5: tensorflow.OpDef.ArgDef.type:type_name -> tensorflow.DataType
	6, // 6: tensorflow.OpDef.AttrDef.default_value:type_name -> tensorflow.AttrValue
	6, // 7: tensorflow.OpDef.AttrDef.allowed_values:type_name -> tensorflow.AttrValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_op_def_proto_init() }
func file_tensorflow_core_framework_op_def_proto_init() {
	if File_tensorflow_core_framework_op_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_op_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpDef); i {
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
		file_tensorflow_core_framework_op_def_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpDeprecation); i {
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
		file_tensorflow_core_framework_op_def_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpList); i {
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
		file_tensorflow_core_framework_op_def_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpDef_ArgDef); i {
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
		file_tensorflow_core_framework_op_def_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpDef_AttrDef); i {
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
			RawDescriptor: file_tensorflow_core_framework_op_def_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_op_def_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_op_def_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_op_def_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_op_def_proto = out.File
	file_tensorflow_core_framework_op_def_proto_rawDesc = nil
	file_tensorflow_core_framework_op_def_proto_goTypes = nil
	file_tensorflow_core_framework_op_def_proto_depIdxs = nil
}
