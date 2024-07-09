// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: colorfy/iot/device/v3/updates.proto

package iotdevpb

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type PropertyUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *_struct.Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PropertyUpdate) Reset() {
	*x = PropertyUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colorfy_iot_device_v3_updates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyUpdate) ProtoMessage() {}

func (x *PropertyUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_colorfy_iot_device_v3_updates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyUpdate.ProtoReflect.Descriptor instead.
func (*PropertyUpdate) Descriptor() ([]byte, []int) {
	return file_colorfy_iot_device_v3_updates_proto_rawDescGZIP(), []int{0}
}

func (x *PropertyUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PropertyUpdate) GetValue() *_struct.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type FirmwareUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firmware *Firmware `protobuf:"bytes,1,opt,name=firmware,proto3" json:"firmware,omitempty"`
	Forced   bool      `protobuf:"varint,2,opt,name=forced,proto3" json:"forced,omitempty"`
}

func (x *FirmwareUpdate) Reset() {
	*x = FirmwareUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colorfy_iot_device_v3_updates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwareUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwareUpdate) ProtoMessage() {}

func (x *FirmwareUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_colorfy_iot_device_v3_updates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwareUpdate.ProtoReflect.Descriptor instead.
func (*FirmwareUpdate) Descriptor() ([]byte, []int) {
	return file_colorfy_iot_device_v3_updates_proto_rawDescGZIP(), []int{1}
}

func (x *FirmwareUpdate) GetFirmware() *Firmware {
	if x != nil {
		return x.Firmware
	}
	return nil
}

func (x *FirmwareUpdate) GetForced() bool {
	if x != nil {
		return x.Forced
	}
	return false
}

type ResetDeviceUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetDeviceUpdate) Reset() {
	*x = ResetDeviceUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colorfy_iot_device_v3_updates_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetDeviceUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetDeviceUpdate) ProtoMessage() {}

func (x *ResetDeviceUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_colorfy_iot_device_v3_updates_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetDeviceUpdate.ProtoReflect.Descriptor instead.
func (*ResetDeviceUpdate) Descriptor() ([]byte, []int) {
	return file_colorfy_iot_device_v3_updates_proto_rawDescGZIP(), []int{2}
}

type RebootDeviceUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RebootDeviceUpdate) Reset() {
	*x = RebootDeviceUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colorfy_iot_device_v3_updates_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebootDeviceUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebootDeviceUpdate) ProtoMessage() {}

func (x *RebootDeviceUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_colorfy_iot_device_v3_updates_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebootDeviceUpdate.ProtoReflect.Descriptor instead.
func (*RebootDeviceUpdate) Descriptor() ([]byte, []int) {
	return file_colorfy_iot_device_v3_updates_proto_rawDescGZIP(), []int{3}
}

type DebugCommandUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string          `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Args    *_struct.Struct `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
}

func (x *DebugCommandUpdate) Reset() {
	*x = DebugCommandUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colorfy_iot_device_v3_updates_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugCommandUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugCommandUpdate) ProtoMessage() {}

func (x *DebugCommandUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_colorfy_iot_device_v3_updates_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugCommandUpdate.ProtoReflect.Descriptor instead.
func (*DebugCommandUpdate) Descriptor() ([]byte, []int) {
	return file_colorfy_iot_device_v3_updates_proto_rawDescGZIP(), []int{4}
}

func (x *DebugCommandUpdate) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *DebugCommandUpdate) GetArgs() *_struct.Struct {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_colorfy_iot_device_v3_updates_proto protoreflect.FileDescriptor

var file_colorfy_iot_device_v3_updates_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2f, 0x69, 0x6f, 0x74, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2e, 0x69,
	0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2f, 0x69, 0x6f, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x65, 0x0a, 0x0e, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79,
	0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x46,
	0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x52, 0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x14,
	0x0a, 0x12, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x5b, 0x0a, 0x12, 0x44, 0x65, 0x62, 0x75, 0x67, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x42, 0x3c, 0x5a, 0x3a, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2f, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x66, 0x79, 0x2d, 0x69, 0x6f, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x69, 0x6f, 0x74, 0x64, 0x65, 0x76, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_colorfy_iot_device_v3_updates_proto_rawDescOnce sync.Once
	file_colorfy_iot_device_v3_updates_proto_rawDescData = file_colorfy_iot_device_v3_updates_proto_rawDesc
)

func file_colorfy_iot_device_v3_updates_proto_rawDescGZIP() []byte {
	file_colorfy_iot_device_v3_updates_proto_rawDescOnce.Do(func() {
		file_colorfy_iot_device_v3_updates_proto_rawDescData = protoimpl.X.CompressGZIP(file_colorfy_iot_device_v3_updates_proto_rawDescData)
	})
	return file_colorfy_iot_device_v3_updates_proto_rawDescData
}

var file_colorfy_iot_device_v3_updates_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_colorfy_iot_device_v3_updates_proto_goTypes = []interface{}{
	(*PropertyUpdate)(nil),     // 0: colorfy.iot.device.v3.PropertyUpdate
	(*FirmwareUpdate)(nil),     // 1: colorfy.iot.device.v3.FirmwareUpdate
	(*ResetDeviceUpdate)(nil),  // 2: colorfy.iot.device.v3.ResetDeviceUpdate
	(*RebootDeviceUpdate)(nil), // 3: colorfy.iot.device.v3.RebootDeviceUpdate
	(*DebugCommandUpdate)(nil), // 4: colorfy.iot.device.v3.DebugCommandUpdate
	(*_struct.Value)(nil),      // 5: google.protobuf.Value
	(*Firmware)(nil),           // 6: colorfy.iot.device.v3.Firmware
	(*_struct.Struct)(nil),     // 7: google.protobuf.Struct
}
var file_colorfy_iot_device_v3_updates_proto_depIdxs = []int32{
	5, // 0: colorfy.iot.device.v3.PropertyUpdate.value:type_name -> google.protobuf.Value
	6, // 1: colorfy.iot.device.v3.FirmwareUpdate.firmware:type_name -> colorfy.iot.device.v3.Firmware
	7, // 2: colorfy.iot.device.v3.DebugCommandUpdate.args:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_colorfy_iot_device_v3_updates_proto_init() }
func file_colorfy_iot_device_v3_updates_proto_init() {
	if File_colorfy_iot_device_v3_updates_proto != nil {
		return
	}
	file_colorfy_iot_device_v3_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_colorfy_iot_device_v3_updates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyUpdate); i {
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
		file_colorfy_iot_device_v3_updates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwareUpdate); i {
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
		file_colorfy_iot_device_v3_updates_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetDeviceUpdate); i {
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
		file_colorfy_iot_device_v3_updates_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RebootDeviceUpdate); i {
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
		file_colorfy_iot_device_v3_updates_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugCommandUpdate); i {
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
			RawDescriptor: file_colorfy_iot_device_v3_updates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_colorfy_iot_device_v3_updates_proto_goTypes,
		DependencyIndexes: file_colorfy_iot_device_v3_updates_proto_depIdxs,
		MessageInfos:      file_colorfy_iot_device_v3_updates_proto_msgTypes,
	}.Build()
	File_colorfy_iot_device_v3_updates_proto = out.File
	file_colorfy_iot_device_v3_updates_proto_rawDesc = nil
	file_colorfy_iot_device_v3_updates_proto_goTypes = nil
	file_colorfy_iot_device_v3_updates_proto_depIdxs = nil
}