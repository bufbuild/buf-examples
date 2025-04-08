// Copyright 2020-2025 Buf Technologies, Inc.
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
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: bufstream/demo/v1/demo.proto

// Implements types for the Bufstream demo.

package demov1

import (
	_ "buf.build/gen/go/bufbuild/confluent/protocolbuffers/go/buf/confluent/v1"
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

const file_bufstream_demo_v1_demo_proto_rawDesc = "" +
	"\n" +
	"\x1cbufstream/demo/v1/demo.proto\x12\x11bufstream.demo.v1\x1a!buf/confluent/v1/extensions.proto\x1a\x1bbuf/validate/validate.proto\"\xc5\x01\n" +
	"\fEmailUpdated\x12\x18\n" +
	"\x02id\x18\x01 \x01(\tB\b\xbaH\x05r\x03\xb0\x01\x01R\x02id\x129\n" +
	"\x11old_email_address\x18\x02 \x01(\tB\r\xbaH\a\xd8\x01\x01r\x02`\x01\x80\x01\x01R\x0foldEmailAddress\x126\n" +
	"\x11new_email_address\x18\x03 \x01(\tB\n" +
	"\xbaH\a\xc8\x01\x01r\x02`\x01R\x0fnewEmailAddress:(\xb2H%\n" +
	"\x0ebufstream-demo\x12\x13email-updated-valueB\xe0\x01\n" +
	"\x15com.bufstream.demo.v1B\tDemoProtoP\x01ZVgithub.com/bufbuild/buf-examples/bufstream/iceberg-quickstart/bufstream/demo/v1;demov1\xa2\x02\x03BDX\xaa\x02\x11Bufstream.Demo.V1\xca\x02\x11Bufstream\\Demo\\V1\xe2\x02\x1dBufstream\\Demo\\V1\\GPBMetadata\xea\x02\x13Bufstream::Demo::V1b\x06proto3"

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
