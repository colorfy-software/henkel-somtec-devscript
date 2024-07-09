// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: colorfy/iot/device/v3/rpc.proto

package iotdevpb

import (
	proto "github.com/golang/protobuf/proto"
	any1 "github.com/golang/protobuf/ptypes/any"
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

type Stream int32

const (
	Stream_STREAM_NONE     Stream = 0
	Stream_STREAM_CONTINUE Stream = 1
	Stream_STREAM_STOP     Stream = 2
	Stream_STREAM_START    Stream = 3
)

// Enum value maps for Stream.
var (
	Stream_name = map[int32]string{
		0: "STREAM_NONE",
		1: "STREAM_CONTINUE",
		2: "STREAM_STOP",
		3: "STREAM_START",
	}
	Stream_value = map[string]int32{
		"STREAM_NONE":     0,
		"STREAM_CONTINUE": 1,
		"STREAM_STOP":     2,
		"STREAM_START":    3,
	}
)

func (x Stream) Enum() *Stream {
	p := new(Stream)
	*p = x
	return p
}

func (x Stream) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Stream) Descriptor() protoreflect.EnumDescriptor {
	return file_colorfy_iot_device_v3_rpc_proto_enumTypes[0].Descriptor()
}

func (Stream) Type() protoreflect.EnumType {
	return &file_colorfy_iot_device_v3_rpc_proto_enumTypes[0]
}

func (x Stream) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Stream.Descriptor instead.
func (Stream) EnumDescriptor() ([]byte, []int) {
	return file_colorfy_iot_device_v3_rpc_proto_rawDescGZIP(), []int{0}
}

type Response_Code int32

const (
	Response_CODE_OK                  Response_Code = 0
	Response_CODE_CANCELLED           Response_Code = 1
	Response_CODE_UNKNOWN             Response_Code = 2
	Response_CODE_INVALID_ARGUMENT    Response_Code = 3
	Response_CODE_DEADLINE_EXCEEDED   Response_Code = 4
	Response_CODE_NOT_FOUND           Response_Code = 5
	Response_CODE_ALREADY_EXISTS      Response_Code = 6
	Response_CODE_PERMISSION_DENIED   Response_Code = 7
	Response_CODE_UNAUTHENTICATED     Response_Code = 16
	Response_CODE_RESOURCE_EXHAUSTED  Response_Code = 8
	Response_CODE_FAILED_PRECONDITION Response_Code = 9
	Response_CODE_ABORTED             Response_Code = 10
	Response_CODE_OUT_OF_RANGE        Response_Code = 11
	Response_CODE_UNIMPLEMENTED       Response_Code = 12
	Response_CODE_INTERNAL            Response_Code = 13
	Response_CODE_UNAVAILABLE         Response_Code = 14
	Response_CODE_DATA_LOSS           Response_Code = 15
)

// Enum value maps for Response_Code.
var (
	Response_Code_name = map[int32]string{
		0:  "CODE_OK",
		1:  "CODE_CANCELLED",
		2:  "CODE_UNKNOWN",
		3:  "CODE_INVALID_ARGUMENT",
		4:  "CODE_DEADLINE_EXCEEDED",
		5:  "CODE_NOT_FOUND",
		6:  "CODE_ALREADY_EXISTS",
		7:  "CODE_PERMISSION_DENIED",
		16: "CODE_UNAUTHENTICATED",
		8:  "CODE_RESOURCE_EXHAUSTED",
		9:  "CODE_FAILED_PRECONDITION",
		10: "CODE_ABORTED",
		11: "CODE_OUT_OF_RANGE",
		12: "CODE_UNIMPLEMENTED",
		13: "CODE_INTERNAL",
		14: "CODE_UNAVAILABLE",
		15: "CODE_DATA_LOSS",
	}
	Response_Code_value = map[string]int32{
		"CODE_OK":                  0,
		"CODE_CANCELLED":           1,
		"CODE_UNKNOWN":             2,
		"CODE_INVALID_ARGUMENT":    3,
		"CODE_DEADLINE_EXCEEDED":   4,
		"CODE_NOT_FOUND":           5,
		"CODE_ALREADY_EXISTS":      6,
		"CODE_PERMISSION_DENIED":   7,
		"CODE_UNAUTHENTICATED":     16,
		"CODE_RESOURCE_EXHAUSTED":  8,
		"CODE_FAILED_PRECONDITION": 9,
		"CODE_ABORTED":             10,
		"CODE_OUT_OF_RANGE":        11,
		"CODE_UNIMPLEMENTED":       12,
		"CODE_INTERNAL":            13,
		"CODE_UNAVAILABLE":         14,
		"CODE_DATA_LOSS":           15,
	}
)

func (x Response_Code) Enum() *Response_Code {
	p := new(Response_Code)
	*p = x
	return p
}

func (x Response_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_colorfy_iot_device_v3_rpc_proto_enumTypes[1].Descriptor()
}

func (Response_Code) Type() protoreflect.EnumType {
	return &file_colorfy_iot_device_v3_rpc_proto_enumTypes[1]
}

func (x Response_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response_Code.Descriptor instead.
func (Response_Code) EnumDescriptor() ([]byte, []int) {
	return file_colorfy_iot_device_v3_rpc_proto_rawDescGZIP(), []int{1, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers map[string]string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Method  string            `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// Request ID
	//
	// Responses to this request need to carry the same Request ID.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Request body. Type is inferred from the service definition
	Body         []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	ClientStream Stream `protobuf:"varint,5,opt,name=client_stream,json=clientStream,proto3,enum=colorfy.iot.device.v3.Stream" json:"client_stream,omitempty"`
	ServerStream Stream `protobuf:"varint,6,opt,name=server_stream,json=serverStream,proto3,enum=colorfy.iot.device.v3.Stream" json:"server_stream,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colorfy_iot_device_v3_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_colorfy_iot_device_v3_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_colorfy_iot_device_v3_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Request) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Request) GetClientStream() Stream {
	if x != nil {
		return x.ClientStream
	}
	return Stream_STREAM_NONE
}

func (x *Request) GetServerStream() Stream {
	if x != nil {
		return x.ServerStream
	}
	return Stream_STREAM_NONE
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers map[string]string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Response Status Code
	//
	// OK if response successful, or other code if response failed
	// In case of failure the body may contain an Error object
	Code Response_Code `protobuf:"varint,2,opt,name=code,proto3,enum=colorfy.iot.device.v3.Response_Code" json:"code,omitempty"`
	// Request ID
	//
	// Should carry the same id as the corresponding request
	RequestId string `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Response body. Type is inferred from the service definition. Only applicable if code = OK
	Body []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	// Error information. Only applicable if code != OK
	Error        *Response_Error `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	ClientStream Stream          `protobuf:"varint,6,opt,name=client_stream,json=clientStream,proto3,enum=colorfy.iot.device.v3.Stream" json:"client_stream,omitempty"`
	ServerStream Stream          `protobuf:"varint,7,opt,name=server_stream,json=serverStream,proto3,enum=colorfy.iot.device.v3.Stream" json:"server_stream,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colorfy_iot_device_v3_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_colorfy_iot_device_v3_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_colorfy_iot_device_v3_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Response) GetCode() Response_Code {
	if x != nil {
		return x.Code
	}
	return Response_CODE_OK
}

func (x *Response) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *Response) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Response) GetError() *Response_Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Response) GetClientStream() Stream {
	if x != nil {
		return x.ClientStream
	}
	return Stream_STREAM_NONE
}

func (x *Response) GetServerStream() Stream {
	if x != nil {
		return x.ServerStream
	}
	return Stream_STREAM_NONE
}

type Response_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Details []*any1.Any `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Response_Error) Reset() {
	*x = Response_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colorfy_iot_device_v3_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Error) ProtoMessage() {}

func (x *Response_Error) ProtoReflect() protoreflect.Message {
	mi := &file_colorfy_iot_device_v3_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Error.ProtoReflect.Descriptor instead.
func (*Response_Error) Descriptor() ([]byte, []int) {
	return file_colorfy_iot_device_v3_rpc_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Response_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response_Error) GetDetails() []*any1.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_colorfy_iot_device_v3_rpc_proto protoreflect.FileDescriptor

var file_colorfy_iot_device_v3_rpc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2f, 0x69, 0x6f, 0x74, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x45, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x42, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x66, 0x79, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x42, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa2, 0x07, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2e, 0x69,
	0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x66, 0x79, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66,
	0x79, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x0c, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x42, 0x0a, 0x0d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x51, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8c, 0x03, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x41, 0x44,
	0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x10, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x45, 0x58, 0x48, 0x41, 0x55, 0x53, 0x54, 0x45, 0x44, 0x10, 0x08, 0x12,
	0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x50,
	0x52, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12,
	0x15, 0x0a, 0x11, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55,
	0x4e, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10,
	0x0d, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49,
	0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x4c, 0x4f, 0x53, 0x53, 0x10, 0x0f, 0x2a, 0x51, 0x0a, 0x06, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x03, 0x42, 0x3c,
	0x5a, 0x3a, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x66, 0x79, 0x2d,
	0x69, 0x6f, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x3b, 0x69, 0x6f, 0x74, 0x64, 0x65, 0x76, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_colorfy_iot_device_v3_rpc_proto_rawDescOnce sync.Once
	file_colorfy_iot_device_v3_rpc_proto_rawDescData = file_colorfy_iot_device_v3_rpc_proto_rawDesc
)

func file_colorfy_iot_device_v3_rpc_proto_rawDescGZIP() []byte {
	file_colorfy_iot_device_v3_rpc_proto_rawDescOnce.Do(func() {
		file_colorfy_iot_device_v3_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_colorfy_iot_device_v3_rpc_proto_rawDescData)
	})
	return file_colorfy_iot_device_v3_rpc_proto_rawDescData
}

var file_colorfy_iot_device_v3_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_colorfy_iot_device_v3_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_colorfy_iot_device_v3_rpc_proto_goTypes = []interface{}{
	(Stream)(0),            // 0: colorfy.iot.device.v3.Stream
	(Response_Code)(0),     // 1: colorfy.iot.device.v3.Response.Code
	(*Request)(nil),        // 2: colorfy.iot.device.v3.Request
	(*Response)(nil),       // 3: colorfy.iot.device.v3.Response
	nil,                    // 4: colorfy.iot.device.v3.Request.HeadersEntry
	(*Response_Error)(nil), // 5: colorfy.iot.device.v3.Response.Error
	nil,                    // 6: colorfy.iot.device.v3.Response.HeadersEntry
	(*any1.Any)(nil),       // 7: google.protobuf.Any
}
var file_colorfy_iot_device_v3_rpc_proto_depIdxs = []int32{
	4, // 0: colorfy.iot.device.v3.Request.headers:type_name -> colorfy.iot.device.v3.Request.HeadersEntry
	0, // 1: colorfy.iot.device.v3.Request.client_stream:type_name -> colorfy.iot.device.v3.Stream
	0, // 2: colorfy.iot.device.v3.Request.server_stream:type_name -> colorfy.iot.device.v3.Stream
	6, // 3: colorfy.iot.device.v3.Response.headers:type_name -> colorfy.iot.device.v3.Response.HeadersEntry
	1, // 4: colorfy.iot.device.v3.Response.code:type_name -> colorfy.iot.device.v3.Response.Code
	5, // 5: colorfy.iot.device.v3.Response.error:type_name -> colorfy.iot.device.v3.Response.Error
	0, // 6: colorfy.iot.device.v3.Response.client_stream:type_name -> colorfy.iot.device.v3.Stream
	0, // 7: colorfy.iot.device.v3.Response.server_stream:type_name -> colorfy.iot.device.v3.Stream
	7, // 8: colorfy.iot.device.v3.Response.Error.details:type_name -> google.protobuf.Any
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_colorfy_iot_device_v3_rpc_proto_init() }
func file_colorfy_iot_device_v3_rpc_proto_init() {
	if File_colorfy_iot_device_v3_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_colorfy_iot_device_v3_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_colorfy_iot_device_v3_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_colorfy_iot_device_v3_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Error); i {
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
			RawDescriptor: file_colorfy_iot_device_v3_rpc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_colorfy_iot_device_v3_rpc_proto_goTypes,
		DependencyIndexes: file_colorfy_iot_device_v3_rpc_proto_depIdxs,
		EnumInfos:         file_colorfy_iot_device_v3_rpc_proto_enumTypes,
		MessageInfos:      file_colorfy_iot_device_v3_rpc_proto_msgTypes,
	}.Build()
	File_colorfy_iot_device_v3_rpc_proto = out.File
	file_colorfy_iot_device_v3_rpc_proto_rawDesc = nil
	file_colorfy_iot_device_v3_rpc_proto_goTypes = nil
	file_colorfy_iot_device_v3_rpc_proto_depIdxs = nil
}