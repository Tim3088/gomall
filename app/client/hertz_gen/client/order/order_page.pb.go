// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: order_page.proto

package order

import (
	_ "Go-Mall/app/client/hertz_gen/api"
	common "Go-Mall/app/client/hertz_gen/client/common"
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

type ListOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" query:"user_id"`
}

func (x *ListOrderReq) Reset() {
	*x = ListOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderReq) ProtoMessage() {}

func (x *ListOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderReq.ProtoReflect.Descriptor instead.
func (*ListOrderReq) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{0}
}

func (x *ListOrderReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty" form:"street_address" query:"street_address"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty" form:"city" query:"city"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty" form:"state" query:"state"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty" form:"country" query:"country"`
	ZipCode       int32  `protobuf:"varint,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty" form:"zip_code" query:"zip_code"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty" form:"product_id" query:"product_id"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty" form:"quantity" query:"quantity"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{2}
}

func (x *CartItem) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *CartItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty" form:"item" query:"item"`
	Cost float32   `protobuf:"fixed32,2,opt,name=cost,proto3" json:"cost,omitempty" form:"cost" query:"cost"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{3}
}

func (x *OrderItem) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *OrderItem) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type PlaceOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCurrency string       `protobuf:"bytes,1,opt,name=user_currency,json=userCurrency,proto3" json:"user_currency,omitempty" form:"user_currency" query:"user_currency"`
	Address      *Address     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" form:"address" query:"address"`
	Email        string       `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" form:"email" query:"email"`
	OrderItems   []*OrderItem `protobuf:"bytes,4,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty" form:"order_items" query:"order_items"`
	Firstname    string       `protobuf:"bytes,5,opt,name=firstname,proto3" json:"firstname,omitempty" form:"firstname" query:"firstname"`
	Lastname     string       `protobuf:"bytes,6,opt,name=lastname,proto3" json:"lastname,omitempty" form:"lastname" query:"lastname"`
}

func (x *PlaceOrderReq) Reset() {
	*x = PlaceOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_page_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderReq) ProtoMessage() {}

func (x *PlaceOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_page_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderReq.ProtoReflect.Descriptor instead.
func (*PlaceOrderReq) Descriptor() ([]byte, []int) {
	return file_order_page_proto_rawDescGZIP(), []int{4}
}

func (x *PlaceOrderReq) GetUserCurrency() string {
	if x != nil {
		return x.UserCurrency
	}
	return ""
}

func (x *PlaceOrderReq) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PlaceOrderReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PlaceOrderReq) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

func (x *PlaceOrderReq) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *PlaceOrderReq) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

var File_order_page_proto protoreflect.FileDescriptor

var file_order_page_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x34, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x24, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x0b, 0xb2, 0xbb, 0x18, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x45, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x4b, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2a, 0x0a, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0xef, 0x01, 0x0a,
	0x0d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x23,
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x2f, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x38, 0x0a, 0x0b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xb1,
	0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4e, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0f,
	0xca, 0xc1, 0x18, 0x0b, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x51, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x47, 0x6f, 0x2d, 0x4d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_page_proto_rawDescOnce sync.Once
	file_order_page_proto_rawDescData = file_order_page_proto_rawDesc
)

func file_order_page_proto_rawDescGZIP() []byte {
	file_order_page_proto_rawDescOnce.Do(func() {
		file_order_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_page_proto_rawDescData)
	})
	return file_order_page_proto_rawDescData
}

var file_order_page_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_order_page_proto_goTypes = []interface{}{
	(*ListOrderReq)(nil),  // 0: client.order.ListOrderReq
	(*Address)(nil),       // 1: client.order.Address
	(*CartItem)(nil),      // 2: client.order.CartItem
	(*OrderItem)(nil),     // 3: client.order.OrderItem
	(*PlaceOrderReq)(nil), // 4: client.order.PlaceOrderReq
	(*common.Empty)(nil),  // 5: client.common.Empty
}
var file_order_page_proto_depIdxs = []int32{
	2, // 0: client.order.OrderItem.item:type_name -> client.order.CartItem
	1, // 1: client.order.PlaceOrderReq.address:type_name -> client.order.Address
	3, // 2: client.order.PlaceOrderReq.order_items:type_name -> client.order.OrderItem
	0, // 3: client.order.OrderService.OrderList:input_type -> client.order.ListOrderReq
	4, // 4: client.order.OrderService.PlaceOrder:input_type -> client.order.PlaceOrderReq
	5, // 5: client.order.OrderService.OrderList:output_type -> client.common.Empty
	5, // 6: client.order.OrderService.PlaceOrder:output_type -> client.common.Empty
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_order_page_proto_init() }
func file_order_page_proto_init() {
	if File_order_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderReq); i {
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
		file_order_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_order_page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_order_page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_order_page_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderReq); i {
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
			RawDescriptor: file_order_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_page_proto_goTypes,
		DependencyIndexes: file_order_page_proto_depIdxs,
		MessageInfos:      file_order_page_proto_msgTypes,
	}.Build()
	File_order_page_proto = out.File
	file_order_page_proto_rawDesc = nil
	file_order_page_proto_goTypes = nil
	file_order_page_proto_depIdxs = nil
}
