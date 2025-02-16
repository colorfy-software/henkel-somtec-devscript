// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: henkel/device/v3/somat/events.proto

package iotsomatpb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type WashCycleStartedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Event UUID. Should be randomly generated in firmware when event is created
	EventId string `protobuf:"bytes,7,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Device UUID
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Time this event occurred (when wash cycle started)
	Time *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	// Wash Cycle UUID should be randomly generated by firmware whenever a new wash cycle starts
	WashCycleId string `protobuf:"bytes,1,opt,name=wash_cycle_id,json=washCycleId,proto3" json:"wash_cycle_id,omitempty"`
	// Program setting this wash cycle was started with
	Program string `protobuf:"bytes,4,opt,name=program,proto3" json:"program,omitempty"`
	// Detergent fill level when wash cycle started
	DetergentFillLevel int32 `protobuf:"varint,5,opt,name=detergent_fill_level,json=detergentFillLevel,proto3" json:"detergent_fill_level,omitempty"`
	// Battery level when wash cycle started
	BatteryLevel string `protobuf:"bytes,6,opt,name=battery_level,json=batteryLevel,proto3" json:"battery_level,omitempty"`
	// Refill UUID that is used during this wash cycle
	RefillId string `protobuf:"bytes,8,opt,name=refill_id,json=refillId,proto3" json:"refill_id,omitempty"`
}

func (x *WashCycleStartedEvent) Reset() {
	*x = WashCycleStartedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_henkel_device_v3_somat_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WashCycleStartedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WashCycleStartedEvent) ProtoMessage() {}

func (x *WashCycleStartedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_henkel_device_v3_somat_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WashCycleStartedEvent.ProtoReflect.Descriptor instead.
func (*WashCycleStartedEvent) Descriptor() ([]byte, []int) {
	return file_henkel_device_v3_somat_events_proto_rawDescGZIP(), []int{0}
}

func (x *WashCycleStartedEvent) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *WashCycleStartedEvent) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *WashCycleStartedEvent) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *WashCycleStartedEvent) GetWashCycleId() string {
	if x != nil {
		return x.WashCycleId
	}
	return ""
}

func (x *WashCycleStartedEvent) GetProgram() string {
	if x != nil {
		return x.Program
	}
	return ""
}

func (x *WashCycleStartedEvent) GetDetergentFillLevel() int32 {
	if x != nil {
		return x.DetergentFillLevel
	}
	return 0
}

func (x *WashCycleStartedEvent) GetBatteryLevel() string {
	if x != nil {
		return x.BatteryLevel
	}
	return ""
}

func (x *WashCycleStartedEvent) GetRefillId() string {
	if x != nil {
		return x.RefillId
	}
	return ""
}

type WashCycleCompletedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Event UUID. Should be randomly generated in firmware when event is created
	EventId string `protobuf:"bytes,20,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Device UUID
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Time this event occurred (when wash cycle completed)
	Time *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	// Wash Cycle UUID should be randomly generated by firmware whenever a new wash cycle starts
	WashCycleId string `protobuf:"bytes,1,opt,name=wash_cycle_id,json=washCycleId,proto3" json:"wash_cycle_id,omitempty"`
	// Wash cycle duration in milliseconds
	Duration int32 `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	// Program setting this wash cycle was run with
	Program                   string `protobuf:"bytes,5,opt,name=program,proto3" json:"program,omitempty"`
	InitialDetergentFillLevel int32  `protobuf:"varint,6,opt,name=initial_detergent_fill_level,json=initialDetergentFillLevel,proto3" json:"initial_detergent_fill_level,omitempty"`
	FinalDetergentFillLevel   int32  `protobuf:"varint,7,opt,name=final_detergent_fill_level,json=finalDetergentFillLevel,proto3" json:"final_detergent_fill_level,omitempty"`
	InitialEnzymeFillLevel    int32  `protobuf:"varint,14,opt,name=initial_enzyme_fill_level,json=initialEnzymeFillLevel,proto3" json:"initial_enzyme_fill_level,omitempty"`
	FinalEnzymeFillLevel      int32  `protobuf:"varint,15,opt,name=final_enzyme_fill_level,json=finalEnzymeFillLevel,proto3" json:"final_enzyme_fill_level,omitempty"`
	InitialAlkaliFillLevel    int32  `protobuf:"varint,16,opt,name=initial_alkali_fill_level,json=initialAlkaliFillLevel,proto3" json:"initial_alkali_fill_level,omitempty"`
	FinalAlkaliFillLevel      int32  `protobuf:"varint,17,opt,name=final_alkali_fill_level,json=finalAlkaliFillLevel,proto3" json:"final_alkali_fill_level,omitempty"`
	InitialRinserFillLevel    int32  `protobuf:"varint,18,opt,name=initial_rinser_fill_level,json=initialRinserFillLevel,proto3" json:"initial_rinser_fill_level,omitempty"`
	FinalRinserFillLevel      int32  `protobuf:"varint,19,opt,name=final_rinser_fill_level,json=finalRinserFillLevel,proto3" json:"final_rinser_fill_level,omitempty"`
	InitialBatteryLevel       string `protobuf:"bytes,11,opt,name=initial_battery_level,json=initialBatteryLevel,proto3" json:"initial_battery_level,omitempty"`
	FinalBatteryLevel         string `protobuf:"bytes,12,opt,name=final_battery_level,json=finalBatteryLevel,proto3" json:"final_battery_level,omitempty"`
	// Maximum temperature in degree celsius reached during wash cycle
	MaxTemp float32 `protobuf:"fixed32,21,opt,name=max_temp,json=maxTemp,proto3" json:"max_temp,omitempty"`
	// Refill UUID that was used during this wash cycle
	RefillId string `protobuf:"bytes,13,opt,name=refill_id,json=refillId,proto3" json:"refill_id,omitempty"`
	// Deprecated. Serves no purpose
	TimeEstimated bool `protobuf:"varint,10,opt,name=time_estimated,json=timeEstimated,proto3" json:"time_estimated,omitempty"`
}

func (x *WashCycleCompletedEvent) Reset() {
	*x = WashCycleCompletedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_henkel_device_v3_somat_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WashCycleCompletedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WashCycleCompletedEvent) ProtoMessage() {}

func (x *WashCycleCompletedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_henkel_device_v3_somat_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WashCycleCompletedEvent.ProtoReflect.Descriptor instead.
func (*WashCycleCompletedEvent) Descriptor() ([]byte, []int) {
	return file_henkel_device_v3_somat_events_proto_rawDescGZIP(), []int{1}
}

func (x *WashCycleCompletedEvent) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *WashCycleCompletedEvent) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *WashCycleCompletedEvent) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *WashCycleCompletedEvent) GetWashCycleId() string {
	if x != nil {
		return x.WashCycleId
	}
	return ""
}

func (x *WashCycleCompletedEvent) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *WashCycleCompletedEvent) GetProgram() string {
	if x != nil {
		return x.Program
	}
	return ""
}

func (x *WashCycleCompletedEvent) GetInitialDetergentFillLevel() int32 {
	if x != nil {
		return x.InitialDetergentFillLevel
	}
	return 0
}

func (x *WashCycleCompletedEvent) GetFinalDetergentFillLevel() int32 {
	if x != nil {
		return x.FinalDetergentFillLevel
	}
	return 0
}

func (x *WashCycleCompletedEvent) GetInitialEnzymeFillLevel() int32 {
	if x != nil {
		return x.InitialEnzymeFillLevel
	}
	return 0
}

func (x *WashCycleCompletedEvent) GetFinalEnzymeFillLevel() int32 {
	if x != nil {
		return x.FinalEnzymeFillLevel
	}
	return 0
}

func (x *WashCycleCompletedEvent) GetInitialAlkaliFillLevel() int32 {
	if x != nil {
		return x.InitialAlkaliFillLevel
	}
	return 0
}

func (x *WashCycleCompletedEvent) GetFinalAlkaliFillLevel() int32 {
	if x != nil {
		return x.FinalAlkaliFillLevel
	}
	return 0
}

func (x *WashCycleCompletedEvent) GetInitialRinserFillLevel() int32 {
	if x != nil {
		return x.InitialRinserFillLevel
	}
	return 0
}

func (x *WashCycleCompletedEvent) GetFinalRinserFillLevel() int32 {
	if x != nil {
		return x.FinalRinserFillLevel
	}
	return 0
}

func (x *WashCycleCompletedEvent) GetInitialBatteryLevel() string {
	if x != nil {
		return x.InitialBatteryLevel
	}
	return ""
}

func (x *WashCycleCompletedEvent) GetFinalBatteryLevel() string {
	if x != nil {
		return x.FinalBatteryLevel
	}
	return ""
}

func (x *WashCycleCompletedEvent) GetMaxTemp() float32 {
	if x != nil {
		return x.MaxTemp
	}
	return 0
}

func (x *WashCycleCompletedEvent) GetRefillId() string {
	if x != nil {
		return x.RefillId
	}
	return ""
}

func (x *WashCycleCompletedEvent) GetTimeEstimated() bool {
	if x != nil {
		return x.TimeEstimated
	}
	return false
}

type RefillChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Event UUID. Should be randomly generated in firmware when event is created
	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Device UUID
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Time this event occurred (refill change detected)
	Time *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	// New Refill UUID. Should be randomly generated when ever a new refill is inserted
	RefillId string `protobuf:"bytes,4,opt,name=refill_id,json=refillId,proto3" json:"refill_id,omitempty"`
}

func (x *RefillChangedEvent) Reset() {
	*x = RefillChangedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_henkel_device_v3_somat_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefillChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefillChangedEvent) ProtoMessage() {}

func (x *RefillChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_henkel_device_v3_somat_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefillChangedEvent.ProtoReflect.Descriptor instead.
func (*RefillChangedEvent) Descriptor() ([]byte, []int) {
	return file_henkel_device_v3_somat_events_proto_rawDescGZIP(), []int{2}
}

func (x *RefillChangedEvent) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *RefillChangedEvent) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *RefillChangedEvent) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *RefillChangedEvent) GetRefillId() string {
	if x != nil {
		return x.RefillId
	}
	return ""
}

// Deprecated, use WashCycleCompletedEvent instead
type WashCycleEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Wash Cycle ID (UUID)
	Id       string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string               `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Time     *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	// Wash cycle duration in milliseconds
	Duration                  int32  `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Program                   string `protobuf:"bytes,5,opt,name=program,proto3" json:"program,omitempty"`
	InitialDetergentFillLevel int32  `protobuf:"varint,6,opt,name=initial_detergent_fill_level,json=initialDetergentFillLevel,proto3" json:"initial_detergent_fill_level,omitempty"`
	FinalDetergentFillLevel   int32  `protobuf:"varint,7,opt,name=final_detergent_fill_level,json=finalDetergentFillLevel,proto3" json:"final_detergent_fill_level,omitempty"`
	InitialBatteryLevel       string `protobuf:"bytes,11,opt,name=initial_battery_level,json=initialBatteryLevel,proto3" json:"initial_battery_level,omitempty"`
	FinalBatteryLevel         string `protobuf:"bytes,12,opt,name=final_battery_level,json=finalBatteryLevel,proto3" json:"final_battery_level,omitempty"`
	// set to true if timestamp has been estimated by the firmware
	TimeEstimated bool `protobuf:"varint,10,opt,name=time_estimated,json=timeEstimated,proto3" json:"time_estimated,omitempty"`
}

func (x *WashCycleEvent) Reset() {
	*x = WashCycleEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_henkel_device_v3_somat_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WashCycleEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WashCycleEvent) ProtoMessage() {}

func (x *WashCycleEvent) ProtoReflect() protoreflect.Message {
	mi := &file_henkel_device_v3_somat_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WashCycleEvent.ProtoReflect.Descriptor instead.
func (*WashCycleEvent) Descriptor() ([]byte, []int) {
	return file_henkel_device_v3_somat_events_proto_rawDescGZIP(), []int{3}
}

func (x *WashCycleEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WashCycleEvent) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *WashCycleEvent) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *WashCycleEvent) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *WashCycleEvent) GetProgram() string {
	if x != nil {
		return x.Program
	}
	return ""
}

func (x *WashCycleEvent) GetInitialDetergentFillLevel() int32 {
	if x != nil {
		return x.InitialDetergentFillLevel
	}
	return 0
}

func (x *WashCycleEvent) GetFinalDetergentFillLevel() int32 {
	if x != nil {
		return x.FinalDetergentFillLevel
	}
	return 0
}

func (x *WashCycleEvent) GetInitialBatteryLevel() string {
	if x != nil {
		return x.InitialBatteryLevel
	}
	return ""
}

func (x *WashCycleEvent) GetFinalBatteryLevel() string {
	if x != nil {
		return x.FinalBatteryLevel
	}
	return ""
}

func (x *WashCycleEvent) GetTimeEstimated() bool {
	if x != nil {
		return x.TimeEstimated
	}
	return false
}

var File_henkel_device_v3_somat_events_proto protoreflect.FileDescriptor

var file_henkel_device_v3_somat_events_proto_rawDesc = []byte{
	0x0a, 0x23, 0x68, 0x65, 0x6e, 0x6b, 0x65, 0x6c, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x73, 0x6f, 0x6d, 0x61, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x68, 0x65, 0x6e, 0x6b, 0x65, 0x6c, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x73, 0x6f, 0x6d, 0x61, 0x74, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1,
	0x02, 0x0a, 0x15, 0x57, 0x61, 0x73, 0x68, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x77, 0x61, 0x73, 0x68, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x61, 0x73, 0x68, 0x43, 0x79, 0x63,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x30,
	0x0a, 0x14, 0x64, 0x65, 0x74, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x6c,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x64, 0x65,
	0x74, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c,
	0x49, 0x64, 0x22, 0xf2, 0x06, 0x0a, 0x17, 0x57, 0x61, 0x73, 0x68, 0x43, 0x79, 0x63, 0x6c, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x61, 0x73, 0x68, 0x5f, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77,
	0x61, 0x73, 0x68, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x12, 0x3f, 0x0a, 0x1c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x44,
	0x65, 0x74, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x3b, 0x0a, 0x1a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x72,
	0x67, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x39,
	0x0a, 0x19, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x7a, 0x79, 0x6d, 0x65,
	0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x16, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x7a, 0x79, 0x6d, 0x65,
	0x46, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x35, 0x0a, 0x17, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x7a, 0x79, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x45, 0x6e, 0x7a, 0x79, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x39, 0x0a, 0x19, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x6c, 0x6b, 0x61,
	0x6c, 0x69, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x16, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x41, 0x6c, 0x6b, 0x61,
	0x6c, 0x69, 0x46, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x35, 0x0a, 0x17, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x6c, 0x6b, 0x61, 0x6c, 0x69, 0x5f, 0x66, 0x69, 0x6c, 0x6c,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x41, 0x6c, 0x6b, 0x61, 0x6c, 0x69, 0x46, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x39, 0x0a, 0x19, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x69,
	0x6e, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x69,
	0x6e, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x35, 0x0a,
	0x17, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x69,
	0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x6c, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x18, 0x15, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x54,
	0x65, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x22, 0x99, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x69,
	0x6c, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x69, 0x6c,
	0x6c, 0x49, 0x64, 0x22, 0xac, 0x03, 0x0a, 0x0e, 0x57, 0x61, 0x73, 0x68, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x3f, 0x0a, 0x1c, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x66,
	0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x19, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x72, 0x67, 0x65, 0x6e,
	0x74, 0x46, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x1a, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69,
	0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x46, 0x69,
	0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x42,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x34, 0x5a, 0x32, 0x68, 0x65, 0x6e, 0x6b, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x68, 0x65, 0x6e, 0x6b, 0x65, 0x6c, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x6f, 0x6d, 0x61, 0x74, 0x3b, 0x69, 0x6f,
	0x74, 0x73, 0x6f, 0x6d, 0x61, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_henkel_device_v3_somat_events_proto_rawDescOnce sync.Once
	file_henkel_device_v3_somat_events_proto_rawDescData = file_henkel_device_v3_somat_events_proto_rawDesc
)

func file_henkel_device_v3_somat_events_proto_rawDescGZIP() []byte {
	file_henkel_device_v3_somat_events_proto_rawDescOnce.Do(func() {
		file_henkel_device_v3_somat_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_henkel_device_v3_somat_events_proto_rawDescData)
	})
	return file_henkel_device_v3_somat_events_proto_rawDescData
}

var file_henkel_device_v3_somat_events_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_henkel_device_v3_somat_events_proto_goTypes = []interface{}{
	(*WashCycleStartedEvent)(nil),   // 0: henkel.device.v3.somat.WashCycleStartedEvent
	(*WashCycleCompletedEvent)(nil), // 1: henkel.device.v3.somat.WashCycleCompletedEvent
	(*RefillChangedEvent)(nil),      // 2: henkel.device.v3.somat.RefillChangedEvent
	(*WashCycleEvent)(nil),          // 3: henkel.device.v3.somat.WashCycleEvent
	(*timestamp.Timestamp)(nil),     // 4: google.protobuf.Timestamp
}
var file_henkel_device_v3_somat_events_proto_depIdxs = []int32{
	4, // 0: henkel.device.v3.somat.WashCycleStartedEvent.time:type_name -> google.protobuf.Timestamp
	4, // 1: henkel.device.v3.somat.WashCycleCompletedEvent.time:type_name -> google.protobuf.Timestamp
	4, // 2: henkel.device.v3.somat.RefillChangedEvent.time:type_name -> google.protobuf.Timestamp
	4, // 3: henkel.device.v3.somat.WashCycleEvent.time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_henkel_device_v3_somat_events_proto_init() }
func file_henkel_device_v3_somat_events_proto_init() {
	if File_henkel_device_v3_somat_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_henkel_device_v3_somat_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WashCycleStartedEvent); i {
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
		file_henkel_device_v3_somat_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WashCycleCompletedEvent); i {
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
		file_henkel_device_v3_somat_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefillChangedEvent); i {
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
		file_henkel_device_v3_somat_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WashCycleEvent); i {
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
			RawDescriptor: file_henkel_device_v3_somat_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_henkel_device_v3_somat_events_proto_goTypes,
		DependencyIndexes: file_henkel_device_v3_somat_events_proto_depIdxs,
		MessageInfos:      file_henkel_device_v3_somat_events_proto_msgTypes,
	}.Build()
	File_henkel_device_v3_somat_events_proto = out.File
	file_henkel_device_v3_somat_events_proto_rawDesc = nil
	file_henkel_device_v3_somat_events_proto_goTypes = nil
	file_henkel_device_v3_somat_events_proto_depIdxs = nil
}
