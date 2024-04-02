// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: api.proto

package grpcapi

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

// Функции для терминала
type GetCodeForPrintReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tname    string `protobuf:"bytes,1,opt,name=tname,proto3" json:"tname,omitempty"`
	Gtin     string `protobuf:"bytes,2,opt,name=gtin,proto3" json:"gtin,omitempty"`
	Proddate string `protobuf:"bytes,3,opt,name=proddate,proto3" json:"proddate,omitempty"`
}

func (x *GetCodeForPrintReq) Reset() {
	*x = GetCodeForPrintReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodeForPrintReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodeForPrintReq) ProtoMessage() {}

func (x *GetCodeForPrintReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodeForPrintReq.ProtoReflect.Descriptor instead.
func (*GetCodeForPrintReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetCodeForPrintReq) GetTname() string {
	if x != nil {
		return x.Tname
	}
	return ""
}

func (x *GetCodeForPrintReq) GetGtin() string {
	if x != nil {
		return x.Gtin
	}
	return ""
}

func (x *GetCodeForPrintReq) GetProddate() string {
	if x != nil {
		return x.Proddate
	}
	return ""
}

type GetCodeForPrintResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gtin   string `protobuf:"bytes,1,opt,name=gtin,proto3" json:"gtin,omitempty"`
	Serial string `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	Crypto string `protobuf:"bytes,3,opt,name=crypto,proto3" json:"crypto,omitempty"`
	Id     uint32 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCodeForPrintResp) Reset() {
	*x = GetCodeForPrintResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodeForPrintResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodeForPrintResp) ProtoMessage() {}

func (x *GetCodeForPrintResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodeForPrintResp.ProtoReflect.Descriptor instead.
func (*GetCodeForPrintResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetCodeForPrintResp) GetGtin() string {
	if x != nil {
		return x.Gtin
	}
	return ""
}

func (x *GetCodeForPrintResp) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *GetCodeForPrintResp) GetCrypto() string {
	if x != nil {
		return x.Crypto
	}
	return ""
}

func (x *GetCodeForPrintResp) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// ProducePrinted
type ProducePrintedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tname    string `protobuf:"bytes,1,opt,name=tname,proto3" json:"tname,omitempty"`
	Gtin     string `protobuf:"bytes,2,opt,name=gtin,proto3" json:"gtin,omitempty"`
	Serial   string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	Proddate string `protobuf:"bytes,4,opt,name=proddate,proto3" json:"proddate,omitempty"`
}

func (x *ProducePrintedReq) Reset() {
	*x = ProducePrintedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProducePrintedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProducePrintedReq) ProtoMessage() {}

func (x *ProducePrintedReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProducePrintedReq.ProtoReflect.Descriptor instead.
func (*ProducePrintedReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *ProducePrintedReq) GetTname() string {
	if x != nil {
		return x.Tname
	}
	return ""
}

func (x *ProducePrintedReq) GetGtin() string {
	if x != nil {
		return x.Gtin
	}
	return ""
}

func (x *ProducePrintedReq) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *ProducePrintedReq) GetProddate() string {
	if x != nil {
		return x.Proddate
	}
	return ""
}

// Отбраковка
type DiscardBySerialReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tname  string `protobuf:"bytes,1,opt,name=tname,proto3" json:"tname,omitempty"`
	Gtin   string `protobuf:"bytes,2,opt,name=gtin,proto3" json:"gtin,omitempty"`
	Serial string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
}

func (x *DiscardBySerialReq) Reset() {
	*x = DiscardBySerialReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscardBySerialReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscardBySerialReq) ProtoMessage() {}

func (x *DiscardBySerialReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscardBySerialReq.ProtoReflect.Descriptor instead.
func (*DiscardBySerialReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *DiscardBySerialReq) GetTname() string {
	if x != nil {
		return x.Tname
	}
	return ""
}

func (x *DiscardBySerialReq) GetGtin() string {
	if x != nil {
		return x.Gtin
	}
	return ""
}

func (x *DiscardBySerialReq) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

// Возвращает количество произведенных продуктов
type GetProducedCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tname string `protobuf:"bytes,1,opt,name=tname,proto3" json:"tname,omitempty"`
	Gtin  string `protobuf:"bytes,2,opt,name=gtin,proto3" json:"gtin,omitempty"`
	Date  string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetProducedCountReq) Reset() {
	*x = GetProducedCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProducedCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProducedCountReq) ProtoMessage() {}

func (x *GetProducedCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProducedCountReq.ProtoReflect.Descriptor instead.
func (*GetProducedCountReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetProducedCountReq) GetTname() string {
	if x != nil {
		return x.Tname
	}
	return ""
}

func (x *GetProducedCountReq) GetGtin() string {
	if x != nil {
		return x.Gtin
	}
	return ""
}

func (x *GetProducedCountReq) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetProducedCountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThisTerm int64 `protobuf:"varint,1,opt,name=thisTerm,proto3" json:"thisTerm,omitempty"`
}

func (x *GetProducedCountResp) Reset() {
	*x = GetProducedCountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProducedCountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProducedCountResp) ProtoMessage() {}

func (x *GetProducedCountResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProducedCountResp.ProtoReflect.Descriptor instead.
func (*GetProducedCountResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetProducedCountResp) GetThisTerm() int64 {
	if x != nil {
		return x.ThisTerm
	}
	return 0
}

type AddCodeForPrintReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sname  string `protobuf:"bytes,1,opt,name=sname,proto3" json:"sname,omitempty"`
	Gtin   string `protobuf:"bytes,2,opt,name=gtin,proto3" json:"gtin,omitempty"`
	Serial string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	Crypto string `protobuf:"bytes,4,opt,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *AddCodeForPrintReq) Reset() {
	*x = AddCodeForPrintReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCodeForPrintReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCodeForPrintReq) ProtoMessage() {}

func (x *AddCodeForPrintReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCodeForPrintReq.ProtoReflect.Descriptor instead.
func (*AddCodeForPrintReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *AddCodeForPrintReq) GetSname() string {
	if x != nil {
		return x.Sname
	}
	return ""
}

func (x *AddCodeForPrintReq) GetGtin() string {
	if x != nil {
		return x.Gtin
	}
	return ""
}

func (x *AddCodeForPrintReq) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *AddCodeForPrintReq) GetCrypto() string {
	if x != nil {
		return x.Crypto
	}
	return ""
}

// Админские функции
type AddGoodReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sname string `protobuf:"bytes,1,opt,name=sname,proto3" json:"sname,omitempty"`
	Gtin  string `protobuf:"bytes,2,opt,name=gtin,proto3" json:"gtin,omitempty"`
	Desc  string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *AddGoodReq) Reset() {
	*x = AddGoodReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGoodReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGoodReq) ProtoMessage() {}

func (x *AddGoodReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGoodReq.ProtoReflect.Descriptor instead.
func (*AddGoodReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *AddGoodReq) GetSname() string {
	if x != nil {
		return x.Sname
	}
	return ""
}

func (x *AddGoodReq) GetGtin() string {
	if x != nil {
		return x.Gtin
	}
	return ""
}

func (x *AddGoodReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type GetGoodsCodeGood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gtin  string `protobuf:"bytes,1,opt,name=gtin,proto3" json:"gtin,omitempty"`
	Desc  string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Count int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetGoodsCodeGood) Reset() {
	*x = GetGoodsCodeGood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsCodeGood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsCodeGood) ProtoMessage() {}

func (x *GetGoodsCodeGood) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsCodeGood.ProtoReflect.Descriptor instead.
func (*GetGoodsCodeGood) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *GetGoodsCodeGood) GetGtin() string {
	if x != nil {
		return x.Gtin
	}
	return ""
}

func (x *GetGoodsCodeGood) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *GetGoodsCodeGood) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetGoodsCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Good []*GetGoodsCodeGood `protobuf:"bytes,1,rep,name=good,proto3" json:"good,omitempty"`
}

func (x *GetGoodsCodeResp) Reset() {
	*x = GetGoodsCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsCodeResp) ProtoMessage() {}

func (x *GetGoodsCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsCodeResp.ProtoReflect.Descriptor instead.
func (*GetGoodsCodeResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{10}
}

func (x *GetGoodsCodeResp) GetGood() []*GetGoodsCodeGood {
	if x != nil {
		return x.Good
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x5a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x46,
	0x6f, 0x72, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x67, 0x74, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x64, 0x61, 0x74, 0x65,
	0x22, 0x69, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x72,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x11, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x64, 0x61, 0x74, 0x65, 0x22, 0x56,
	0x0a, 0x12, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x42, 0x79, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x74,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x32, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x69, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x68, 0x69, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x22,
	0x6e, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67,
	0x74, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x22,
	0x4a, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x50, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x67, 0x74, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67,
	0x74, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x39, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x25, 0x0a, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x52, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x32, 0xe0, 0x02, 0x0a, 0x03, 0x48, 0x75, 0x62,
	0x12, 0x3c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x72,
	0x69, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x64,
	0x12, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x0f,
	0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x42, 0x79, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12,
	0x13, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x42, 0x79, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a,
	0x0f, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x12, 0x13, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1e, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x0b, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x6f,
	0x6f, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2c, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_proto_goTypes = []interface{}{
	(*Empty)(nil),                // 0: Empty
	(*GetCodeForPrintReq)(nil),   // 1: GetCodeForPrintReq
	(*GetCodeForPrintResp)(nil),  // 2: GetCodeForPrintResp
	(*ProducePrintedReq)(nil),    // 3: ProducePrintedReq
	(*DiscardBySerialReq)(nil),   // 4: DiscardBySerialReq
	(*GetProducedCountReq)(nil),  // 5: GetProducedCountReq
	(*GetProducedCountResp)(nil), // 6: GetProducedCountResp
	(*AddCodeForPrintReq)(nil),   // 7: AddCodeForPrintReq
	(*AddGoodReq)(nil),           // 8: AddGoodReq
	(*GetGoodsCodeGood)(nil),     // 9: GetGoodsCodeGood
	(*GetGoodsCodeResp)(nil),     // 10: GetGoodsCodeResp
}
var file_api_proto_depIdxs = []int32{
	9,  // 0: GetGoodsCodeResp.good:type_name -> GetGoodsCodeGood
	1,  // 1: Hub.GetCodeForPrint:input_type -> GetCodeForPrintReq
	3,  // 2: Hub.ProducePrinted:input_type -> ProducePrintedReq
	4,  // 3: Hub.DiscardBySerial:input_type -> DiscardBySerialReq
	5,  // 4: Hub.GetProducedCount:input_type -> GetProducedCountReq
	7,  // 5: Hub.AddCodeForPrint:input_type -> AddCodeForPrintReq
	8,  // 6: Hub.AddGood:input_type -> AddGoodReq
	0,  // 7: Hub.GetGoodsCodeReq:input_type -> Empty
	2,  // 8: Hub.GetCodeForPrint:output_type -> GetCodeForPrintResp
	0,  // 9: Hub.ProducePrinted:output_type -> Empty
	0,  // 10: Hub.DiscardBySerial:output_type -> Empty
	6,  // 11: Hub.GetProducedCount:output_type -> GetProducedCountResp
	0,  // 12: Hub.AddCodeForPrint:output_type -> Empty
	0,  // 13: Hub.AddGood:output_type -> Empty
	10, // 14: Hub.GetGoodsCodeReq:output_type -> GetGoodsCodeResp
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodeForPrintReq); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodeForPrintResp); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProducePrintedReq); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscardBySerialReq); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProducedCountReq); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProducedCountResp); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCodeForPrintReq); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGoodReq); i {
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
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsCodeGood); i {
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
		file_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsCodeResp); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
