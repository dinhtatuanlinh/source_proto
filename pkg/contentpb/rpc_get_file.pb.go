// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: rpc_get_file.proto

package source_proto

import (
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

type GetFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        int64                  `protobuf:"varint,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	mi := &file_rpc_get_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_file_proto_rawDescGZIP(), []int{0}
}

func (x *GetFileRequest) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

type GetFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	File          *File                  `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFileResponse) Reset() {
	*x = GetFileResponse{}
	mi := &file_rpc_get_file_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResponse) ProtoMessage() {}

func (x *GetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_file_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResponse.ProtoReflect.Descriptor instead.
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_file_proto_rawDescGZIP(), []int{1}
}

func (x *GetFileResponse) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

var File_rpc_get_file_proto protoreflect.FileDescriptor

var file_rpc_get_file_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x2f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x6e, 0x68, 0x74, 0x61, 0x74, 0x75, 0x61, 0x6e, 0x6c, 0x69, 0x6e, 0x68, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_rpc_get_file_proto_rawDescOnce sync.Once
	file_rpc_get_file_proto_rawDescData []byte
)

func file_rpc_get_file_proto_rawDescGZIP() []byte {
	file_rpc_get_file_proto_rawDescOnce.Do(func() {
		file_rpc_get_file_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_get_file_proto_rawDesc), len(file_rpc_get_file_proto_rawDesc)))
	})
	return file_rpc_get_file_proto_rawDescData
}

var file_rpc_get_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_get_file_proto_goTypes = []any{
	(*GetFileRequest)(nil),  // 0: pb.GetFileRequest
	(*GetFileResponse)(nil), // 1: pb.GetFileResponse
	(*File)(nil),            // 2: pb.File
}
var file_rpc_get_file_proto_depIdxs = []int32{
	2, // 0: pb.GetFileResponse.file:type_name -> pb.File
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_get_file_proto_init() }
func file_rpc_get_file_proto_init() {
	if File_rpc_get_file_proto != nil {
		return
	}
	file_file_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_get_file_proto_rawDesc), len(file_rpc_get_file_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_get_file_proto_goTypes,
		DependencyIndexes: file_rpc_get_file_proto_depIdxs,
		MessageInfos:      file_rpc_get_file_proto_msgTypes,
	}.Build()
	File_rpc_get_file_proto = out.File
	file_rpc_get_file_proto_goTypes = nil
	file_rpc_get_file_proto_depIdxs = nil
}
