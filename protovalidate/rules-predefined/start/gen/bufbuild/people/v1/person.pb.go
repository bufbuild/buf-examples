// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: bufbuild/people/v1/person.proto

package peoplev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Person struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required, 50 or fewer characters.
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// Optional, 50 or fewer characters.
	MiddleName string `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	// Required, 50 or fewer characters.
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// Optional, non-standard length
	Title         string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Person) Reset() {
	*x = Person{}
	mi := &file_bufbuild_people_v1_person_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_bufbuild_people_v1_person_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_bufbuild_people_v1_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Person) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *Person) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Person) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_bufbuild_people_v1_person_proto protoreflect.FileDescriptor

var file_bufbuild_people_v1_person_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x65, 0x6f, 0x70,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x18, 0x32, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x18,
	0x40, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0xf2, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x0b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x61, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x2d, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x6f,
	0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x50, 0x58, 0xaa, 0x02, 0x12, 0x42, 0x75,
	0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x12, 0x42, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5c, 0x50, 0x65, 0x6f, 0x70,
	0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5c, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x42, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x3a, 0x3a, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_bufbuild_people_v1_person_proto_rawDescOnce sync.Once
	file_bufbuild_people_v1_person_proto_rawDescData []byte
)

func file_bufbuild_people_v1_person_proto_rawDescGZIP() []byte {
	file_bufbuild_people_v1_person_proto_rawDescOnce.Do(func() {
		file_bufbuild_people_v1_person_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_bufbuild_people_v1_person_proto_rawDesc), len(file_bufbuild_people_v1_person_proto_rawDesc)))
	})
	return file_bufbuild_people_v1_person_proto_rawDescData
}

var file_bufbuild_people_v1_person_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bufbuild_people_v1_person_proto_goTypes = []any{
	(*Person)(nil), // 0: bufbuild.people.v1.Person
}
var file_bufbuild_people_v1_person_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bufbuild_people_v1_person_proto_init() }
func file_bufbuild_people_v1_person_proto_init() {
	if File_bufbuild_people_v1_person_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_bufbuild_people_v1_person_proto_rawDesc), len(file_bufbuild_people_v1_person_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bufbuild_people_v1_person_proto_goTypes,
		DependencyIndexes: file_bufbuild_people_v1_person_proto_depIdxs,
		MessageInfos:      file_bufbuild_people_v1_person_proto_msgTypes,
	}.Build()
	File_bufbuild_people_v1_person_proto = out.File
	file_bufbuild_people_v1_person_proto_goTypes = nil
	file_bufbuild_people_v1_person_proto_depIdxs = nil
}
