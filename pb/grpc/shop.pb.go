// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: shop.proto

package grpc

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

type CreateOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Date      string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Count     int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Id        string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateOrderReq) Reset() {
	*x = CreateOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderReq) ProtoMessage() {}

func (x *CreateOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderReq.ProtoReflect.Descriptor instead.
func (*CreateOrderReq) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderReq) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CreateOrderReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateOrderReq) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *CreateOrderReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CreateOrderReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateOrderRes) Reset() {
	*x = CreateOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRes) ProtoMessage() {}

func (x *CreateOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRes.ProtoReflect.Descriptor instead.
func (*CreateOrderRes) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateOrderRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CancelOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CancelOrderReq) Reset() {
	*x = CancelOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderReq) ProtoMessage() {}

func (x *CancelOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderReq.ProtoReflect.Descriptor instead.
func (*CancelOrderReq) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{2}
}

func (x *CancelOrderReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CancelOrderReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CancelOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error        string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	RemainAmount int64  `protobuf:"varint,3,opt,name=remainAmount,proto3" json:"remainAmount,omitempty"`
}

func (x *CancelOrderRes) Reset() {
	*x = CancelOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRes) ProtoMessage() {}

func (x *CancelOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRes.ProtoReflect.Descriptor instead.
func (*CancelOrderRes) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{3}
}

func (x *CancelOrderRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CancelOrderRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CancelOrderRes) GetRemainAmount() int64 {
	if x != nil {
		return x.RemainAmount
	}
	return 0
}

type GetByUsernameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetByUsernameReq) Reset() {
	*x = GetByUsernameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUsernameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUsernameReq) ProtoMessage() {}

func (x *GetByUsernameReq) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUsernameReq.ProtoReflect.Descriptor instead.
func (*GetByUsernameReq) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{4}
}

func (x *GetByUsernameReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Date      string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Count     int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Id        string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{5}
}

func (x *Order) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Order) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Order) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Order) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByUsernameRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order []*Order `protobuf:"bytes,1,rep,name=order,proto3" json:"order,omitempty"`
}

func (x *GetByUsernameRes) Reset() {
	*x = GetByUsernameRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUsernameRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUsernameRes) ProtoMessage() {}

func (x *GetByUsernameRes) ProtoReflect() protoreflect.Message {
	mi := &file_shop_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUsernameRes.ProtoReflect.Descriptor instead.
func (*GetByUsernameRes) Descriptor() ([]byte, []int) {
	return file_shop_proto_rawDescGZIP(), []int{6}
}

func (x *GetByUsernameRes) GetOrder() []*Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_shop_proto protoreflect.FileDescriptor

var file_shop_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x68,
	0x6f, 0x70, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x46, 0x0a, 0x0e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x62, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7b, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0xc0, 0x01, 0x0a, 0x0b, 0x53, 0x68,
	0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x14, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_proto_rawDescOnce sync.Once
	file_shop_proto_rawDescData = file_shop_proto_rawDesc
)

func file_shop_proto_rawDescGZIP() []byte {
	file_shop_proto_rawDescOnce.Do(func() {
		file_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_proto_rawDescData)
	})
	return file_shop_proto_rawDescData
}

var file_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_shop_proto_goTypes = []interface{}{
	(*CreateOrderReq)(nil),   // 0: shop.CreateOrderReq
	(*CreateOrderRes)(nil),   // 1: shop.CreateOrderRes
	(*CancelOrderReq)(nil),   // 2: shop.CancelOrderReq
	(*CancelOrderRes)(nil),   // 3: shop.CancelOrderRes
	(*GetByUsernameReq)(nil), // 4: shop.GetByUsernameReq
	(*Order)(nil),            // 5: shop.Order
	(*GetByUsernameRes)(nil), // 6: shop.GetByUsernameRes
}
var file_shop_proto_depIdxs = []int32{
	5, // 0: shop.GetByUsernameRes.order:type_name -> shop.Order
	0, // 1: shop.ShopService.Create:input_type -> shop.CreateOrderReq
	2, // 2: shop.ShopService.Cancel:input_type -> shop.CancelOrderReq
	4, // 3: shop.ShopService.GetByUsername:input_type -> shop.GetByUsernameReq
	1, // 4: shop.ShopService.Create:output_type -> shop.CreateOrderRes
	3, // 5: shop.ShopService.Cancel:output_type -> shop.CancelOrderRes
	6, // 6: shop.ShopService.GetByUsername:output_type -> shop.GetByUsernameRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shop_proto_init() }
func file_shop_proto_init() {
	if File_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderReq); i {
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
		file_shop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRes); i {
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
		file_shop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderReq); i {
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
		file_shop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderRes); i {
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
		file_shop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUsernameReq); i {
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
		file_shop_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_shop_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUsernameRes); i {
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
			RawDescriptor: file_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_proto_goTypes,
		DependencyIndexes: file_shop_proto_depIdxs,
		MessageInfos:      file_shop_proto_msgTypes,
	}.Build()
	File_shop_proto = out.File
	file_shop_proto_rawDesc = nil
	file_shop_proto_goTypes = nil
	file_shop_proto_depIdxs = nil
}
