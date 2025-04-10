// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: acme/weather/v1/weather.proto

package weatherv1

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

type Condition int32

const (
	Condition_CONDITION_UNKNOWN Condition = 0
	Condition_CONDITION_RAINY   Condition = 1
	Condition_CONDITION_SUNNY   Condition = 2
	Condition_CONDITION_CLOUDY  Condition = 3
)

// Enum value maps for Condition.
var (
	Condition_name = map[int32]string{
		0: "CONDITION_UNKNOWN",
		1: "CONDITION_RAINY",
		2: "CONDITION_SUNNY",
		3: "CONDITION_CLOUDY",
	}
	Condition_value = map[string]int32{
		"CONDITION_UNKNOWN": 0,
		"CONDITION_RAINY":   1,
		"CONDITION_SUNNY":   2,
		"CONDITION_CLOUDY":  3,
	}
)

func (x Condition) Enum() *Condition {
	p := new(Condition)
	*p = x
	return p
}

func (x Condition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Condition) Descriptor() protoreflect.EnumDescriptor {
	return file_acme_weather_v1_weather_proto_enumTypes[0].Descriptor()
}

func (Condition) Type() protoreflect.EnumType {
	return &file_acme_weather_v1_weather_proto_enumTypes[0]
}

func (x Condition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Condition.Descriptor instead.
func (Condition) EnumDescriptor() ([]byte, []int) {
	return file_acme_weather_v1_weather_proto_rawDescGZIP(), []int{0}
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	mi := &file_acme_weather_v1_weather_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_acme_weather_v1_weather_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_acme_weather_v1_weather_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Weather struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location    *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Temperature float32   `protobuf:"fixed32,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	WindSpeed   float32   `protobuf:"fixed32,3,opt,name=wind_speed,json=windSpeed,proto3" json:"wind_speed,omitempty"`
	Condition   Condition `protobuf:"varint,4,opt,name=condition,proto3,enum=acme.weather.v1.Condition" json:"condition,omitempty"`
}

func (x *Weather) Reset() {
	*x = Weather{}
	mi := &file_acme_weather_v1_weather_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Weather) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weather) ProtoMessage() {}

func (x *Weather) ProtoReflect() protoreflect.Message {
	mi := &file_acme_weather_v1_weather_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weather.ProtoReflect.Descriptor instead.
func (*Weather) Descriptor() ([]byte, []int) {
	return file_acme_weather_v1_weather_proto_rawDescGZIP(), []int{1}
}

func (x *Weather) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Weather) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Weather) GetWindSpeed() float32 {
	if x != nil {
		return x.WindSpeed
	}
	return 0
}

func (x *Weather) GetCondition() Condition {
	if x != nil {
		return x.Condition
	}
	return Condition_CONDITION_UNKNOWN
}

var File_acme_weather_v1_weather_proto protoreflect.FileDescriptor

var file_acme_weather_v1_weather_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77,
	0x69, 0x6e, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x77, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x62, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x44,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x41, 0x49, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x4e, 0x4e, 0x59,
	0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x4c, 0x4f, 0x55, 0x44, 0x59, 0x10, 0x03, 0x42, 0xb6, 0x01, 0x0a, 0x12, 0x69, 0x6f, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x6d, 0x65,
	0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x6d, 0x65,
	0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x57, 0x58, 0xaa, 0x02, 0x0f, 0x41, 0x63,
	0x6d, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f,
	0x41, 0x63, 0x6d, 0x65, 0x5c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1b, 0x41, 0x63, 0x6d, 0x65, 0x5c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11,
	0x41, 0x63, 0x6d, 0x65, 0x3a, 0x3a, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acme_weather_v1_weather_proto_rawDescOnce sync.Once
	file_acme_weather_v1_weather_proto_rawDescData = file_acme_weather_v1_weather_proto_rawDesc
)

func file_acme_weather_v1_weather_proto_rawDescGZIP() []byte {
	file_acme_weather_v1_weather_proto_rawDescOnce.Do(func() {
		file_acme_weather_v1_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_acme_weather_v1_weather_proto_rawDescData)
	})
	return file_acme_weather_v1_weather_proto_rawDescData
}

var file_acme_weather_v1_weather_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_acme_weather_v1_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_acme_weather_v1_weather_proto_goTypes = []any{
	(Condition)(0),   // 0: acme.weather.v1.Condition
	(*Location)(nil), // 1: acme.weather.v1.Location
	(*Weather)(nil),  // 2: acme.weather.v1.Weather
}
var file_acme_weather_v1_weather_proto_depIdxs = []int32{
	1, // 0: acme.weather.v1.Weather.location:type_name -> acme.weather.v1.Location
	0, // 1: acme.weather.v1.Weather.condition:type_name -> acme.weather.v1.Condition
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_acme_weather_v1_weather_proto_init() }
func file_acme_weather_v1_weather_proto_init() {
	if File_acme_weather_v1_weather_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_acme_weather_v1_weather_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acme_weather_v1_weather_proto_goTypes,
		DependencyIndexes: file_acme_weather_v1_weather_proto_depIdxs,
		EnumInfos:         file_acme_weather_v1_weather_proto_enumTypes,
		MessageInfos:      file_acme_weather_v1_weather_proto_msgTypes,
	}.Build()
	File_acme_weather_v1_weather_proto = out.File
	file_acme_weather_v1_weather_proto_rawDesc = nil
	file_acme_weather_v1_weather_proto_goTypes = nil
	file_acme_weather_v1_weather_proto_depIdxs = nil
}
