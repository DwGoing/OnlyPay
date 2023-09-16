// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: proto/funds_service/treasury.proto

package treasury_generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRechargeOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalIdentity string  `protobuf:"bytes,1,opt,name=externalIdentity,proto3" json:"externalIdentity,omitempty"`
	ExternalData     []byte  `protobuf:"bytes,2,opt,name=externalData,proto3,oneof" json:"externalData,omitempty"`
	CallbackUrl      string  `protobuf:"bytes,3,opt,name=callbackUrl,proto3" json:"callbackUrl,omitempty"`
	ChainType        string  `protobuf:"bytes,4,opt,name=chainType,proto3" json:"chainType,omitempty"`
	Amount           float64 `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	WalletIndex      int64   `protobuf:"varint,6,opt,name=walletIndex,proto3" json:"walletIndex,omitempty"`
}

func (x *CreateRechargeOrderRequest) Reset() {
	*x = CreateRechargeOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_funds_service_treasury_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRechargeOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRechargeOrderRequest) ProtoMessage() {}

func (x *CreateRechargeOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_funds_service_treasury_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRechargeOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateRechargeOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_funds_service_treasury_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRechargeOrderRequest) GetExternalIdentity() string {
	if x != nil {
		return x.ExternalIdentity
	}
	return ""
}

func (x *CreateRechargeOrderRequest) GetExternalData() []byte {
	if x != nil {
		return x.ExternalData
	}
	return nil
}

func (x *CreateRechargeOrderRequest) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

func (x *CreateRechargeOrderRequest) GetChainType() string {
	if x != nil {
		return x.ChainType
	}
	return ""
}

func (x *CreateRechargeOrderRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateRechargeOrderRequest) GetWalletIndex() int64 {
	if x != nil {
		return x.WalletIndex
	}
	return 0
}

type CreateRechargeOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	ExpireAt string `protobuf:"bytes,2,opt,name=expireAt,proto3" json:"expireAt,omitempty"`
}

func (x *CreateRechargeOrderResponse) Reset() {
	*x = CreateRechargeOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_funds_service_treasury_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRechargeOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRechargeOrderResponse) ProtoMessage() {}

func (x *CreateRechargeOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_funds_service_treasury_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRechargeOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateRechargeOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_funds_service_treasury_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRechargeOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateRechargeOrderResponse) GetExpireAt() string {
	if x != nil {
		return x.ExpireAt
	}
	return ""
}

type SubmitRechargeOrderTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	TxHash  string `protobuf:"bytes,2,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (x *SubmitRechargeOrderTransactionRequest) Reset() {
	*x = SubmitRechargeOrderTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_funds_service_treasury_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRechargeOrderTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRechargeOrderTransactionRequest) ProtoMessage() {}

func (x *SubmitRechargeOrderTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_funds_service_treasury_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRechargeOrderTransactionRequest.ProtoReflect.Descriptor instead.
func (*SubmitRechargeOrderTransactionRequest) Descriptor() ([]byte, []int) {
	return file_proto_funds_service_treasury_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitRechargeOrderTransactionRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *SubmitRechargeOrderTransactionRequest) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

var File_proto_funds_service_treasury_proto protoreflect.FileDescriptor

var file_proto_funds_service_treasury_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0c,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x53, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x41, 0x74, 0x22, 0x59, 0x0a, 0x25, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x32, 0xc4, 0x01, 0x0a, 0x08, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x12, 0x53, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x70, 0x63, 0x12, 0x1b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x63, 0x0a, 0x21, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x70, 0x63, 0x12, 0x26, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x3b, 0x74, 0x72,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_funds_service_treasury_proto_rawDescOnce sync.Once
	file_proto_funds_service_treasury_proto_rawDescData = file_proto_funds_service_treasury_proto_rawDesc
)

func file_proto_funds_service_treasury_proto_rawDescGZIP() []byte {
	file_proto_funds_service_treasury_proto_rawDescOnce.Do(func() {
		file_proto_funds_service_treasury_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_funds_service_treasury_proto_rawDescData)
	})
	return file_proto_funds_service_treasury_proto_rawDescData
}

var file_proto_funds_service_treasury_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_funds_service_treasury_proto_goTypes = []interface{}{
	(*CreateRechargeOrderRequest)(nil),            // 0: CreateRechargeOrderRequest
	(*CreateRechargeOrderResponse)(nil),           // 1: CreateRechargeOrderResponse
	(*SubmitRechargeOrderTransactionRequest)(nil), // 2: SubmitRechargeOrderTransactionRequest
	(*emptypb.Empty)(nil),                         // 3: google.protobuf.Empty
}
var file_proto_funds_service_treasury_proto_depIdxs = []int32{
	0, // 0: Treasury.CreateRechargeOrderRpc:input_type -> CreateRechargeOrderRequest
	2, // 1: Treasury.SubmitRechargeOrderTransactionRpc:input_type -> SubmitRechargeOrderTransactionRequest
	1, // 2: Treasury.CreateRechargeOrderRpc:output_type -> CreateRechargeOrderResponse
	3, // 3: Treasury.SubmitRechargeOrderTransactionRpc:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_funds_service_treasury_proto_init() }
func file_proto_funds_service_treasury_proto_init() {
	if File_proto_funds_service_treasury_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_funds_service_treasury_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRechargeOrderRequest); i {
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
		file_proto_funds_service_treasury_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRechargeOrderResponse); i {
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
		file_proto_funds_service_treasury_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRechargeOrderTransactionRequest); i {
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
	file_proto_funds_service_treasury_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_funds_service_treasury_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_funds_service_treasury_proto_goTypes,
		DependencyIndexes: file_proto_funds_service_treasury_proto_depIdxs,
		MessageInfos:      file_proto_funds_service_treasury_proto_msgTypes,
	}.Build()
	File_proto_funds_service_treasury_proto = out.File
	file_proto_funds_service_treasury_proto_rawDesc = nil
	file_proto_funds_service_treasury_proto_goTypes = nil
	file_proto_funds_service_treasury_proto_depIdxs = nil
}
