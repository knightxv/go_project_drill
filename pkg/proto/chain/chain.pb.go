// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: chain/chain.proto

package pbChain

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

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_chain_chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_chain_chain_proto_rawDescGZIP(), []int{0}
}

func (x *CommonResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *CommonResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type WithdrawTransferReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationID string `protobuf:"bytes,1,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Amount      int64  `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Decimal     int32  `protobuf:"varint,4,opt,name=Decimal,proto3" json:"Decimal,omitempty"`
	WithdrawId  int64  `protobuf:"varint,5,opt,name=WithdrawId,proto3" json:"WithdrawId,omitempty"`
	Contract    string `protobuf:"bytes,6,opt,name=Contract,proto3" json:"Contract,omitempty"`
	ChainId     int64  `protobuf:"varint,7,opt,name=ChainId,proto3" json:"ChainId,omitempty"`
	Endpoint    string `protobuf:"bytes,8,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
}

func (x *WithdrawTransferReq) Reset() {
	*x = WithdrawTransferReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawTransferReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawTransferReq) ProtoMessage() {}

func (x *WithdrawTransferReq) ProtoReflect() protoreflect.Message {
	mi := &file_chain_chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawTransferReq.ProtoReflect.Descriptor instead.
func (*WithdrawTransferReq) Descriptor() ([]byte, []int) {
	return file_chain_chain_proto_rawDescGZIP(), []int{1}
}

func (x *WithdrawTransferReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *WithdrawTransferReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *WithdrawTransferReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *WithdrawTransferReq) GetDecimal() int32 {
	if x != nil {
		return x.Decimal
	}
	return 0
}

func (x *WithdrawTransferReq) GetWithdrawId() int64 {
	if x != nil {
		return x.WithdrawId
	}
	return 0
}

func (x *WithdrawTransferReq) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *WithdrawTransferReq) GetChainId() int64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *WithdrawTransferReq) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type WithdrawTransferResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp   *CommonResp `protobuf:"bytes,1,opt,name=CommonResp,proto3" json:"CommonResp,omitempty"`
	ChainUpLogId string      `protobuf:"bytes,2,opt,name=ChainUpLogId,proto3" json:"ChainUpLogId,omitempty"`
	TxHash       string      `protobuf:"bytes,3,opt,name=TxHash,proto3" json:"TxHash,omitempty"`
}

func (x *WithdrawTransferResp) Reset() {
	*x = WithdrawTransferResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_chain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawTransferResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawTransferResp) ProtoMessage() {}

func (x *WithdrawTransferResp) ProtoReflect() protoreflect.Message {
	mi := &file_chain_chain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawTransferResp.ProtoReflect.Descriptor instead.
func (*WithdrawTransferResp) Descriptor() ([]byte, []int) {
	return file_chain_chain_proto_rawDescGZIP(), []int{2}
}

func (x *WithdrawTransferResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *WithdrawTransferResp) GetChainUpLogId() string {
	if x != nil {
		return x.ChainUpLogId
	}
	return ""
}

func (x *WithdrawTransferResp) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

var File_chain_chain_proto protoreflect.FileDescriptor

var file_chain_chain_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x3e, 0x0a, 0x0a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x22, 0xf3, 0x01, 0x0a, 0x13, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x1e,
	0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x22, 0x85, 0x01, 0x0a, 0x14, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x0a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x0c,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x55, 0x70, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x55, 0x70, 0x4c, 0x6f, 0x67, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x32, 0x5b, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x3b,
	0x70, 0x62, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chain_chain_proto_rawDescOnce sync.Once
	file_chain_chain_proto_rawDescData = file_chain_chain_proto_rawDesc
)

func file_chain_chain_proto_rawDescGZIP() []byte {
	file_chain_chain_proto_rawDescOnce.Do(func() {
		file_chain_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_chain_chain_proto_rawDescData)
	})
	return file_chain_chain_proto_rawDescData
}

var file_chain_chain_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chain_chain_proto_goTypes = []interface{}{
	(*CommonResp)(nil),           // 0: chain.CommonResp
	(*WithdrawTransferReq)(nil),  // 1: chain.WithdrawTransferReq
	(*WithdrawTransferResp)(nil), // 2: chain.WithdrawTransferResp
}
var file_chain_chain_proto_depIdxs = []int32{
	0, // 0: chain.WithdrawTransferResp.CommonResp:type_name -> chain.CommonResp
	1, // 1: chain.chainService.WithdrawTransfer:input_type -> chain.WithdrawTransferReq
	2, // 2: chain.chainService.WithdrawTransfer:output_type -> chain.WithdrawTransferResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chain_chain_proto_init() }
func file_chain_chain_proto_init() {
	if File_chain_chain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chain_chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
		file_chain_chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawTransferReq); i {
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
		file_chain_chain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawTransferResp); i {
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
			RawDescriptor: file_chain_chain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chain_chain_proto_goTypes,
		DependencyIndexes: file_chain_chain_proto_depIdxs,
		MessageInfos:      file_chain_chain_proto_msgTypes,
	}.Build()
	File_chain_chain_proto = out.File
	file_chain_chain_proto_rawDesc = nil
	file_chain_chain_proto_goTypes = nil
	file_chain_chain_proto_depIdxs = nil
}
