// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: proto/LogicService/LogicService.proto

package LogicService

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

type ProtoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId       string  `protobuf:"bytes,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	IsBeat      int32   `protobuf:"varint,2,opt,name=isBeat,proto3" json:"isBeat,omitempty"`
	Token       string  `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	ServiceName *string `protobuf:"bytes,4,opt,name=serviceName,proto3,oneof" json:"serviceName,omitempty"`
	ActionName  *string `protobuf:"bytes,5,opt,name=actionName,proto3,oneof" json:"actionName,omitempty"`
	Data        []byte  `protobuf:"bytes,6,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *ProtoMessage) Reset() {
	*x = ProtoMessage{}
	mi := &file_proto_LogicService_LogicService_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoMessage) ProtoMessage() {}

func (x *ProtoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_LogicService_LogicService_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoMessage.ProtoReflect.Descriptor instead.
func (*ProtoMessage) Descriptor() ([]byte, []int) {
	return file_proto_LogicService_LogicService_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoMessage) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *ProtoMessage) GetIsBeat() int32 {
	if x != nil {
		return x.IsBeat
	}
	return 0
}

func (x *ProtoMessage) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ProtoMessage) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

func (x *ProtoMessage) GetActionName() string {
	if x != nil && x.ActionName != nil {
		return *x.ActionName
	}
	return ""
}

func (x *ProtoMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProtoMessageResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId       string  `protobuf:"bytes,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	IsAck       int32   `protobuf:"varint,2,opt,name=isAck,proto3" json:"isAck,omitempty"`
	AccountId   int64   `protobuf:"varint,3,opt,name=accountId,proto3" json:"accountId,omitempty"`
	ServiceName *string `protobuf:"bytes,4,opt,name=serviceName,proto3,oneof" json:"serviceName,omitempty"`
	ActionName  *string `protobuf:"bytes,5,opt,name=actionName,proto3,oneof" json:"actionName,omitempty"`
	Data        []byte  `protobuf:"bytes,6,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *ProtoMessageResult) Reset() {
	*x = ProtoMessageResult{}
	mi := &file_proto_LogicService_LogicService_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoMessageResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoMessageResult) ProtoMessage() {}

func (x *ProtoMessageResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_LogicService_LogicService_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoMessageResult.ProtoReflect.Descriptor instead.
func (*ProtoMessageResult) Descriptor() ([]byte, []int) {
	return file_proto_LogicService_LogicService_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoMessageResult) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *ProtoMessageResult) GetIsAck() int32 {
	if x != nil {
		return x.IsAck
	}
	return 0
}

func (x *ProtoMessageResult) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *ProtoMessageResult) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

func (x *ProtoMessageResult) GetActionName() string {
	if x != nil && x.ActionName != nil {
		return *x.ActionName
	}
	return ""
}

func (x *ProtoMessageResult) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_LogicService_LogicService_proto protoreflect.FileDescriptor

var file_proto_LogicService_LogicService_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x73, 0x42, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x69, 0x73, 0x42, 0x65, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x02, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01,
	0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0xeb, 0x01, 0x0a, 0x12, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x41, 0x63, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x73, 0x41, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x02, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x32, 0x41, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f,
	0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_LogicService_LogicService_proto_rawDescOnce sync.Once
	file_proto_LogicService_LogicService_proto_rawDescData = file_proto_LogicService_LogicService_proto_rawDesc
)

func file_proto_LogicService_LogicService_proto_rawDescGZIP() []byte {
	file_proto_LogicService_LogicService_proto_rawDescOnce.Do(func() {
		file_proto_LogicService_LogicService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_LogicService_LogicService_proto_rawDescData)
	})
	return file_proto_LogicService_LogicService_proto_rawDescData
}

var file_proto_LogicService_LogicService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_LogicService_LogicService_proto_goTypes = []any{
	(*ProtoMessage)(nil),       // 0: ProtoMessage
	(*ProtoMessageResult)(nil), // 1: ProtoMessageResult
}
var file_proto_LogicService_LogicService_proto_depIdxs = []int32{
	0, // 0: LogicService.SendMessage:input_type -> ProtoMessage
	1, // 1: LogicService.SendMessage:output_type -> ProtoMessageResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_LogicService_LogicService_proto_init() }
func file_proto_LogicService_LogicService_proto_init() {
	if File_proto_LogicService_LogicService_proto != nil {
		return
	}
	file_proto_LogicService_LogicService_proto_msgTypes[0].OneofWrappers = []any{}
	file_proto_LogicService_LogicService_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_LogicService_LogicService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_LogicService_LogicService_proto_goTypes,
		DependencyIndexes: file_proto_LogicService_LogicService_proto_depIdxs,
		MessageInfos:      file_proto_LogicService_LogicService_proto_msgTypes,
	}.Build()
	File_proto_LogicService_LogicService_proto = out.File
	file_proto_LogicService_LogicService_proto_rawDesc = nil
	file_proto_LogicService_LogicService_proto_goTypes = nil
	file_proto_LogicService_LogicService_proto_depIdxs = nil
}
