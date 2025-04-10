// Copyright 2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: bufstream/demo/v1/demo.proto

// Implements types for the Bufstream demo.

package demov1

import (
	_ "buf.build/gen/go/bufbuild/confluent/protocolbuffers/go/buf/confluent/v1"
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

// An event where an email address was updated for a given user.
//
// This represents the schema of data sent to the `email-updated-value` subject.
type EmailUpdated struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the user associated with this email address update.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The old email address.
	OldEmailAddress string `protobuf:"bytes,2,opt,name=old_email_address,json=oldEmailAddress,proto3" json:"old_email_address,omitempty"`
	// The new email address.
	NewEmailAddress string `protobuf:"bytes,3,opt,name=new_email_address,json=newEmailAddress,proto3" json:"new_email_address,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *EmailUpdated) Reset() {
	*x = EmailUpdated{}
	mi := &file_bufstream_demo_v1_demo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailUpdated) ProtoMessage() {}

func (x *EmailUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_bufstream_demo_v1_demo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailUpdated.ProtoReflect.Descriptor instead.
func (*EmailUpdated) Descriptor() ([]byte, []int) {
	return file_bufstream_demo_v1_demo_proto_rawDescGZIP(), []int{0}
}

func (x *EmailUpdated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EmailUpdated) GetOldEmailAddress() string {
	if x != nil {
		return x.OldEmailAddress
	}
	return ""
}

func (x *EmailUpdated) GetNewEmailAddress() string {
	if x != nil {
		return x.NewEmailAddress
	}
	return ""
}

var File_bufstream_demo_v1_demo_proto protoreflect.FileDescriptor

var file_bufstream_demo_v1_demo_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x62, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x62, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76,
	0x31, 0x1a, 0x21, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x11, 0x6f, 0x6c, 0x64, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0x80, 0x01, 0x01, 0x52, 0x0f, 0x6f, 0x6c, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6e, 0x65, 0x77, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x3a, 0x28, 0xb2, 0x48, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x12, 0x13, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xea, 0x01, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x44, 0x65, 0x6d, 0x6f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x60, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f,
	0x62, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x2f, 0x62, 0x75, 0x66,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x64,
	0x65, 0x6d, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x44, 0x58, 0xaa, 0x02, 0x11, 0x42, 0x75,
	0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x11, 0x42, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5c, 0x44, 0x65, 0x6d, 0x6f,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x42, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5c,
	0x44, 0x65, 0x6d, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x42, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x3a,
	0x3a, 0x44, 0x65, 0x6d, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_bufstream_demo_v1_demo_proto_rawDescOnce sync.Once
	file_bufstream_demo_v1_demo_proto_rawDescData []byte
)

func file_bufstream_demo_v1_demo_proto_rawDescGZIP() []byte {
	file_bufstream_demo_v1_demo_proto_rawDescOnce.Do(func() {
		file_bufstream_demo_v1_demo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_bufstream_demo_v1_demo_proto_rawDesc), len(file_bufstream_demo_v1_demo_proto_rawDesc)))
	})
	return file_bufstream_demo_v1_demo_proto_rawDescData
}

var file_bufstream_demo_v1_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bufstream_demo_v1_demo_proto_goTypes = []any{
	(*EmailUpdated)(nil), // 0: bufstream.demo.v1.EmailUpdated
}
var file_bufstream_demo_v1_demo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bufstream_demo_v1_demo_proto_init() }
func file_bufstream_demo_v1_demo_proto_init() {
	if File_bufstream_demo_v1_demo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_bufstream_demo_v1_demo_proto_rawDesc), len(file_bufstream_demo_v1_demo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bufstream_demo_v1_demo_proto_goTypes,
		DependencyIndexes: file_bufstream_demo_v1_demo_proto_depIdxs,
		MessageInfos:      file_bufstream_demo_v1_demo_proto_msgTypes,
	}.Build()
	File_bufstream_demo_v1_demo_proto = out.File
	file_bufstream_demo_v1_demo_proto_goTypes = nil
	file_bufstream_demo_v1_demo_proto_depIdxs = nil
}
