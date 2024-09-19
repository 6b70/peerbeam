// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: control.proto

package controlpb

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

type FileMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId  string `protobuf:"bytes,1,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"`
	FileName    string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileSize    int64  `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	IsDirectory bool   `protobuf:"varint,4,opt,name=is_directory,json=isDirectory,proto3" json:"is_directory,omitempty"`
}

func (x *FileMetadata) Reset() {
	*x = FileMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadata) ProtoMessage() {}

func (x *FileMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadata.ProtoReflect.Descriptor instead.
func (*FileMetadata) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{0}
}

func (x *FileMetadata) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *FileMetadata) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileMetadata) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FileMetadata) GetIsDirectory() bool {
	if x != nil {
		return x.IsDirectory
	}
	return false
}

type FileMetadataList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId string          `protobuf:"bytes,1,opt,name=list_id,json=listId,proto3" json:"list_id,omitempty"`
	Files  []*FileMetadata `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *FileMetadataList) Reset() {
	*x = FileMetadataList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetadataList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadataList) ProtoMessage() {}

func (x *FileMetadataList) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadataList.ProtoReflect.Descriptor instead.
func (*FileMetadataList) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{1}
}

func (x *FileMetadataList) GetListId() string {
	if x != nil {
		return x.ListId
	}
	return ""
}

func (x *FileMetadataList) GetFiles() []*FileMetadata {
	if x != nil {
		return x.Files
	}
	return nil
}

type TransferConsent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consent bool   `protobuf:"varint,1,opt,name=consent,proto3" json:"consent,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *TransferConsent) Reset() {
	*x = TransferConsent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferConsent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferConsent) ProtoMessage() {}

func (x *TransferConsent) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferConsent.ProtoReflect.Descriptor instead.
func (*TransferConsent) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{2}
}

func (x *TransferConsent) GetConsent() bool {
	if x != nil {
		return x.Consent
	}
	return false
}

func (x *TransferConsent) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_control_proto protoreflect.FileDescriptor

var file_control_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x70, 0x65, 0x65, 0x72, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x22, 0x61, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70,
	0x65, 0x65, 0x72, 0x62, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x36, 0x62, 0x37, 0x30, 0x2f, 0x70, 0x65, 0x65, 0x72,
	0x62, 0x65, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_control_proto_rawDescOnce sync.Once
	file_control_proto_rawDescData = file_control_proto_rawDesc
)

func file_control_proto_rawDescGZIP() []byte {
	file_control_proto_rawDescOnce.Do(func() {
		file_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_control_proto_rawDescData)
	})
	return file_control_proto_rawDescData
}

var file_control_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_control_proto_goTypes = []interface{}{
	(*FileMetadata)(nil),     // 0: peerbeam.control.FileMetadata
	(*FileMetadataList)(nil), // 1: peerbeam.control.FileMetadataList
	(*TransferConsent)(nil),  // 2: peerbeam.control.TransferConsent
}
var file_control_proto_depIdxs = []int32{
	0, // 0: peerbeam.control.FileMetadataList.files:type_name -> peerbeam.control.FileMetadata
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_control_proto_init() }
func file_control_proto_init() {
	if File_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMetadata); i {
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
		file_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMetadataList); i {
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
		file_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferConsent); i {
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
			RawDescriptor: file_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_control_proto_goTypes,
		DependencyIndexes: file_control_proto_depIdxs,
		MessageInfos:      file_control_proto_msgTypes,
	}.Build()
	File_control_proto = out.File
	file_control_proto_rawDesc = nil
	file_control_proto_goTypes = nil
	file_control_proto_depIdxs = nil
}
