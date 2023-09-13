// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: shopping/v2/product_list_service.proto

package shopping

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

type TransactionType_Value int32

const (
	TransactionType_APP_LOVIN_REWARDED_VIDEO TransactionType_Value = 0
	TransactionType_GOOGLE_IAP               TransactionType_Value = 1
	TransactionType_APPLE_IAP                TransactionType_Value = 2
)

// Enum value maps for TransactionType_Value.
var (
	TransactionType_Value_name = map[int32]string{
		0: "APP_LOVIN_REWARDED_VIDEO",
		1: "GOOGLE_IAP",
		2: "APPLE_IAP",
	}
	TransactionType_Value_value = map[string]int32{
		"APP_LOVIN_REWARDED_VIDEO": 0,
		"GOOGLE_IAP":               1,
		"APPLE_IAP":                2,
	}
)

func (x TransactionType_Value) Enum() *TransactionType_Value {
	p := new(TransactionType_Value)
	*p = x
	return p
}

func (x TransactionType_Value) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType_Value) Descriptor() protoreflect.EnumDescriptor {
	return file_shopping_v2_product_list_service_proto_enumTypes[0].Descriptor()
}

func (TransactionType_Value) Type() protoreflect.EnumType {
	return &file_shopping_v2_product_list_service_proto_enumTypes[0]
}

func (x TransactionType_Value) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType_Value.Descriptor instead.
func (TransactionType_Value) EnumDescriptor() ([]byte, []int) {
	return file_shopping_v2_product_list_service_proto_rawDescGZIP(), []int{0, 0}
}

type ProductsByCategoryResponse_Result int32

const (
	ProductsByCategoryResponse_OK    ProductsByCategoryResponse_Result = 0
	ProductsByCategoryResponse_ERROR ProductsByCategoryResponse_Result = 1
)

// Enum value maps for ProductsByCategoryResponse_Result.
var (
	ProductsByCategoryResponse_Result_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	ProductsByCategoryResponse_Result_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x ProductsByCategoryResponse_Result) Enum() *ProductsByCategoryResponse_Result {
	p := new(ProductsByCategoryResponse_Result)
	*p = x
	return p
}

func (x ProductsByCategoryResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductsByCategoryResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_shopping_v2_product_list_service_proto_enumTypes[1].Descriptor()
}

func (ProductsByCategoryResponse_Result) Type() protoreflect.EnumType {
	return &file_shopping_v2_product_list_service_proto_enumTypes[1]
}

func (x ProductsByCategoryResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductsByCategoryResponse_Result.Descriptor instead.
func (ProductsByCategoryResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_shopping_v2_product_list_service_proto_rawDescGZIP(), []int{2, 0}
}

type TransactionType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value TransactionType_Value `protobuf:"varint,1,opt,name=value,proto3,enum=mobile.shopping.v2.TransactionType_Value" json:"value,omitempty"`
}

func (x *TransactionType) Reset() {
	*x = TransactionType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_v2_product_list_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionType) ProtoMessage() {}

func (x *TransactionType) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_v2_product_list_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionType.ProtoReflect.Descriptor instead.
func (*TransactionType) Descriptor() ([]byte, []int) {
	return file_shopping_v2_product_list_service_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionType) GetValue() TransactionType_Value {
	if x != nil {
		return x.Value
	}
	return TransactionType_APP_LOVIN_REWARDED_VIDEO
}

type ProductsByCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCategoryId string           `protobuf:"bytes,1,opt,name=product_category_id,json=productCategoryId,proto3" json:"product_category_id,omitempty"`
	ProductType       string           `protobuf:"bytes,2,opt,name=product_type,json=productType,proto3" json:"product_type,omitempty"`
	TransactionType   *TransactionType `protobuf:"bytes,3,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"`
}

func (x *ProductsByCategoryRequest) Reset() {
	*x = ProductsByCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_v2_product_list_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsByCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsByCategoryRequest) ProtoMessage() {}

func (x *ProductsByCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_v2_product_list_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsByCategoryRequest.ProtoReflect.Descriptor instead.
func (*ProductsByCategoryRequest) Descriptor() ([]byte, []int) {
	return file_shopping_v2_product_list_service_proto_rawDescGZIP(), []int{1}
}

func (x *ProductsByCategoryRequest) GetProductCategoryId() string {
	if x != nil {
		return x.ProductCategoryId
	}
	return ""
}

func (x *ProductsByCategoryRequest) GetProductType() string {
	if x != nil {
		return x.ProductType
	}
	return ""
}

func (x *ProductsByCategoryRequest) GetTransactionType() *TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return nil
}

type ProductsByCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       ProductsByCategoryResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mobile.shopping.v2.ProductsByCategoryResponse_Result" json:"result,omitempty"`
	Products     []*Product                        `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	ErrorMessage string                            `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *ProductsByCategoryResponse) Reset() {
	*x = ProductsByCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_v2_product_list_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsByCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsByCategoryResponse) ProtoMessage() {}

func (x *ProductsByCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_v2_product_list_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsByCategoryResponse.ProtoReflect.Descriptor instead.
func (*ProductsByCategoryResponse) Descriptor() ([]byte, []int) {
	return file_shopping_v2_product_list_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProductsByCategoryResponse) GetResult() ProductsByCategoryResponse_Result {
	if x != nil {
		return x.Result
	}
	return ProductsByCategoryResponse_OK
}

func (x *ProductsByCategoryResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductsByCategoryResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCategoryId string `protobuf:"bytes,1,opt,name=product_category_id,json=productCategoryId,proto3" json:"product_category_id,omitempty"`
	ProductId         string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductName       string `protobuf:"bytes,3,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ContextData       string `protobuf:"bytes,4,opt,name=context_data,json=contextData,proto3" json:"context_data,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_v2_product_list_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_v2_product_list_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_shopping_v2_product_list_service_proto_rawDescGZIP(), []int{3}
}

func (x *Product) GetProductCategoryId() string {
	if x != nil {
		return x.ProductCategoryId
	}
	return ""
}

func (x *Product) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Product) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Product) GetContextData() string {
	if x != nil {
		return x.ContextData
	}
	return ""
}

var File_shopping_v2_product_list_service_proto protoreflect.FileDescriptor

var file_shopping_v2_product_list_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x22, 0x98, 0x01, 0x0a,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x44, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x50,
	0x50, 0x5f, 0x4c, 0x4f, 0x56, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x45, 0x44,
	0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x4f, 0x4f, 0x47,
	0x4c, 0x45, 0x5f, 0x49, 0x41, 0x50, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x50, 0x50, 0x4c,
	0x45, 0x5f, 0x49, 0x41, 0x50, 0x10, 0x02, 0x22, 0xbe, 0x01, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xe6, 0x01, 0x0a, 0x1a, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x1b, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06,
	0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x01, 0x22, 0x9e, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x2e, 0x0a,
	0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x32, 0x88, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x79, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2d, 0x2e, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x64, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x32, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x78,
	0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shopping_v2_product_list_service_proto_rawDescOnce sync.Once
	file_shopping_v2_product_list_service_proto_rawDescData = file_shopping_v2_product_list_service_proto_rawDesc
)

func file_shopping_v2_product_list_service_proto_rawDescGZIP() []byte {
	file_shopping_v2_product_list_service_proto_rawDescOnce.Do(func() {
		file_shopping_v2_product_list_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_shopping_v2_product_list_service_proto_rawDescData)
	})
	return file_shopping_v2_product_list_service_proto_rawDescData
}

var file_shopping_v2_product_list_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_shopping_v2_product_list_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shopping_v2_product_list_service_proto_goTypes = []interface{}{
	(TransactionType_Value)(0),             // 0: mobile.shopping.v2.TransactionType.Value
	(ProductsByCategoryResponse_Result)(0), // 1: mobile.shopping.v2.ProductsByCategoryResponse.Result
	(*TransactionType)(nil),                // 2: mobile.shopping.v2.TransactionType
	(*ProductsByCategoryRequest)(nil),      // 3: mobile.shopping.v2.ProductsByCategoryRequest
	(*ProductsByCategoryResponse)(nil),     // 4: mobile.shopping.v2.ProductsByCategoryResponse
	(*Product)(nil),                        // 5: mobile.shopping.v2.Product
}
var file_shopping_v2_product_list_service_proto_depIdxs = []int32{
	0, // 0: mobile.shopping.v2.TransactionType.value:type_name -> mobile.shopping.v2.TransactionType.Value
	2, // 1: mobile.shopping.v2.ProductsByCategoryRequest.transaction_type:type_name -> mobile.shopping.v2.TransactionType
	1, // 2: mobile.shopping.v2.ProductsByCategoryResponse.result:type_name -> mobile.shopping.v2.ProductsByCategoryResponse.Result
	5, // 3: mobile.shopping.v2.ProductsByCategoryResponse.products:type_name -> mobile.shopping.v2.Product
	3, // 4: mobile.shopping.v2.ProductList.ListProductsByCategory:input_type -> mobile.shopping.v2.ProductsByCategoryRequest
	4, // 5: mobile.shopping.v2.ProductList.ListProductsByCategory:output_type -> mobile.shopping.v2.ProductsByCategoryResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_shopping_v2_product_list_service_proto_init() }
func file_shopping_v2_product_list_service_proto_init() {
	if File_shopping_v2_product_list_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shopping_v2_product_list_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionType); i {
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
		file_shopping_v2_product_list_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsByCategoryRequest); i {
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
		file_shopping_v2_product_list_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsByCategoryResponse); i {
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
		file_shopping_v2_product_list_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_shopping_v2_product_list_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shopping_v2_product_list_service_proto_goTypes,
		DependencyIndexes: file_shopping_v2_product_list_service_proto_depIdxs,
		EnumInfos:         file_shopping_v2_product_list_service_proto_enumTypes,
		MessageInfos:      file_shopping_v2_product_list_service_proto_msgTypes,
	}.Build()
	File_shopping_v2_product_list_service_proto = out.File
	file_shopping_v2_product_list_service_proto_rawDesc = nil
	file_shopping_v2_product_list_service_proto_goTypes = nil
	file_shopping_v2_product_list_service_proto_depIdxs = nil
}
