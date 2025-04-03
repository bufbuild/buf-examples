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
// source: invoice/v1/invoice.proto

package invoicev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Invoice is a collection of goods or services sold to a customer.
type Invoice struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// invoice_id is a unique identifier for this invoice.
	InvoiceId string `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	// account_id is the unique identifier for the customer associated with this
	// invoice.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// invoice_date is the date for an invoice. It should represent a date and
	// have no values for time components.
	InvoiceDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=invoice_date,json=invoiceDate,proto3" json:"invoice_date,omitempty"`
	// line_items represent individual items on this invoice.
	LineItems     []*LineItem `protobuf:"bytes,4,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	mi := &file_invoice_v1_invoice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_v1_invoice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_invoice_v1_invoice_proto_rawDescGZIP(), []int{0}
}

func (x *Invoice) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *Invoice) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Invoice) GetInvoiceDate() *timestamppb.Timestamp {
	if x != nil {
		return x.InvoiceDate
	}
	return nil
}

func (x *Invoice) GetLineItems() []*LineItem {
	if x != nil {
		return x.LineItems
	}
	return nil
}

// LineItem is an individual good or service added to an invoice.
type LineItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// line_item_id is a unique identifier for this LineItem.
	LineItemId string `protobuf:"bytes,1,opt,name=line_item_id,json=lineItemId,proto3" json:"line_item_id,omitempty"`
	// product_id is the unique identifier for the good or service on this line.
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// quantity is the unit count of the good or service provided.
	Quantity uint64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// unit_price is the price per unit. Note that it is an integer:
	// if this system is using U.S. Dollars, it would represent one cent. We are
	// not working towards internationalization or fractions of a cent within
	// this example.
	UnitPrice     uint64 `protobuf:"varint,4,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LineItem) Reset() {
	*x = LineItem{}
	mi := &file_invoice_v1_invoice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItem) ProtoMessage() {}

func (x *LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_v1_invoice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItem.ProtoReflect.Descriptor instead.
func (*LineItem) Descriptor() ([]byte, []int) {
	return file_invoice_v1_invoice_proto_rawDescGZIP(), []int{1}
}

func (x *LineItem) GetLineItemId() string {
	if x != nil {
		return x.LineItemId
	}
	return ""
}

func (x *LineItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *LineItem) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *LineItem) GetUnitPrice() uint64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

var File_invoice_v1_invoice_proto protoreflect.FileDescriptor

var file_invoice_v1_invoice_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x08, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x12, 0x27, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0xe6, 0x01, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0xa6, 0x01, 0xba, 0x48, 0xa2, 0x01, 0xba, 0x01, 0x9b, 0x01,
	0x0a, 0x14, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x6e,
	0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65,
	0x20, 0x7a, 0x65, 0x72, 0x6f, 0x1a, 0x64, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x48,
	0x6f, 0x75, 0x72, 0x73, 0x28, 0x29, 0x20, 0x3d, 0x3d, 0x20, 0x30, 0x26, 0x26, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x28, 0x29, 0x20,
	0x3d, 0x3d, 0x20, 0x30, 0x26, 0x26, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x28, 0x29, 0x20, 0x3d, 0x3d, 0x20, 0x30, 0x26, 0x26, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x28, 0x29, 0x20, 0x3d, 0x3d, 0x20, 0x30, 0xc8, 0x01, 0x01, 0x52, 0x0b,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x8c, 0x03, 0x0a, 0x0a,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0xd6, 0x02, 0xba, 0x48, 0xd2, 0x02, 0xba, 0x01, 0x9a,
	0x01, 0x0a, 0x1e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x12, 0x26, 0x61, 0x6c, 0x6c, 0x20, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x1a, 0x50, 0x74, 0x68, 0x69, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x28, 0x20, 0x69, 0x74, 0x2c, 0x20, 0x69, 0x74, 0x2e, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x2e, 0x73, 0x69, 0x7a, 0x65,
	0x28, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x29, 0x2e, 0x6d, 0x61, 0x70, 0x28, 0x20, 0x69, 0x74, 0x2c,
	0x20, 0x69, 0x74, 0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x20, 0x29, 0x2e, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x28, 0x29, 0xba, 0x01, 0xa8, 0x01, 0x0a,
	0x1b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x63, 0x61, 0x6c, 0x6c, 0x79, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x43, 0x6c, 0x69,
	0x6e, 0x65, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65,
	0x20, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x1a, 0x44, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x28, 0x20, 0x69, 0x74, 0x2c,
	0x20, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x20, 0x2b,
	0x20, 0x27, 0x2d, 0x27, 0x20, 0x2b, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x28, 0x69, 0x74,
	0x2e, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x29, 0x20, 0x29, 0x2e, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x28, 0x29, 0x92, 0x01, 0x05, 0x08, 0x01, 0x10, 0xe8, 0x07, 0x52,
	0x09, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0xff, 0x02, 0xba, 0x48, 0xfb,
	0x02, 0x1a, 0x71, 0x0a, 0x19, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x2e,
	0x6e, 0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x30,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x61, 0x6d, 0x65,
	0x1a, 0x22, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x20, 0x21, 0x3d, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x1a, 0x81, 0x01, 0x0a, 0x19, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x12, 0x25, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x20, 0x63,
	0x61, 0x6e, 0x27, 0x74, 0x20, 0x61, 0x6c, 0x73, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x1a, 0x3d, 0x21, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x65, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x28, 0x69, 0x74, 0x2c, 0x20, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x20, 0x3d, 0x3d, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x29, 0x1a, 0x81, 0x01, 0x0a, 0x19, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x25, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x20, 0x63, 0x61, 0x6e, 0x27, 0x74, 0x20, 0x61, 0x6c, 0x73, 0x6f, 0x20, 0x62, 0x65,
	0x20, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x1a, 0x3d, 0x21,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x28, 0x69, 0x74, 0x2c, 0x20, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x20, 0x3d, 0x3d, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x29, 0x22, 0x9f, 0x02, 0x0a,
	0x08, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2a, 0x0a, 0x0c, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03,
	0xb0, 0x01, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xba, 0x48, 0x04, 0x32, 0x02, 0x28, 0x00,
	0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x3a, 0x71, 0xba, 0x48, 0x6e,
	0x1a, 0x6c, 0x0a, 0x1b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x2e, 0x6e, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x12,
	0x27, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x20, 0x63, 0x61,
	0x6e, 0x27, 0x74, 0x20, 0x61, 0x6c, 0x73, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x1a, 0x24, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x20, 0x21, 0x3d, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x42, 0xc7,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x72, 0x70,
	0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67,
	0x6f, 0x2f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x16, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_invoice_v1_invoice_proto_rawDescOnce sync.Once
	file_invoice_v1_invoice_proto_rawDescData []byte
)

func file_invoice_v1_invoice_proto_rawDescGZIP() []byte {
	file_invoice_v1_invoice_proto_rawDescOnce.Do(func() {
		file_invoice_v1_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_invoice_v1_invoice_proto_rawDesc), len(file_invoice_v1_invoice_proto_rawDesc)))
	})
	return file_invoice_v1_invoice_proto_rawDescData
}

var file_invoice_v1_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_invoice_v1_invoice_proto_goTypes = []any{
	(*Invoice)(nil),               // 0: invoice.v1.Invoice
	(*LineItem)(nil),              // 1: invoice.v1.LineItem
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_invoice_v1_invoice_proto_depIdxs = []int32{
	2, // 0: invoice.v1.Invoice.invoice_date:type_name -> google.protobuf.Timestamp
	1, // 1: invoice.v1.Invoice.line_items:type_name -> invoice.v1.LineItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_invoice_v1_invoice_proto_init() }
func file_invoice_v1_invoice_proto_init() {
	if File_invoice_v1_invoice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_invoice_v1_invoice_proto_rawDesc), len(file_invoice_v1_invoice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_invoice_v1_invoice_proto_goTypes,
		DependencyIndexes: file_invoice_v1_invoice_proto_depIdxs,
		MessageInfos:      file_invoice_v1_invoice_proto_msgTypes,
	}.Build()
	File_invoice_v1_invoice_proto = out.File
	file_invoice_v1_invoice_proto_goTypes = nil
	file_invoice_v1_invoice_proto_depIdxs = nil
}
