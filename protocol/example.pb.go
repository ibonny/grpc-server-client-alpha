// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: protocol/example.proto

package protocol

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type FileRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Bucket    string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Prefix    string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Filename  string   `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	Chunksize int32    `protobuf:"varint,5,opt,name=chunksize,proto3" json:"chunksize,omitempty"`
	Filesize  int64    `protobuf:"varint,6,opt,name=filesize,proto3" json:"filesize,omitempty"`
	Chunklist []string `protobuf:"bytes,7,rep,name=chunklist,proto3" json:"chunklist,omitempty"`
	Writetime int64    `protobuf:"varint,8,opt,name=writetime,proto3" json:"writetime,omitempty"`
}

func (x *FileRecord) Reset() {
	*x = FileRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRecord) ProtoMessage() {}

func (x *FileRecord) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRecord.ProtoReflect.Descriptor instead.
func (*FileRecord) Descriptor() ([]byte, []int) {
	return file_protocol_example_proto_rawDescGZIP(), []int{0}
}

func (x *FileRecord) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *FileRecord) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *FileRecord) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *FileRecord) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileRecord) GetChunksize() int32 {
	if x != nil {
		return x.Chunksize
	}
	return 0
}

func (x *FileRecord) GetFilesize() int64 {
	if x != nil {
		return x.Filesize
	}
	return 0
}

func (x *FileRecord) GetChunklist() []string {
	if x != nil {
		return x.Chunklist
	}
	return nil
}

func (x *FileRecord) GetWritetime() int64 {
	if x != nil {
		return x.Writetime
	}
	return 0
}

type ChunkRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Length int64  `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *ChunkRecord) Reset() {
	*x = ChunkRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkRecord) ProtoMessage() {}

func (x *ChunkRecord) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkRecord.ProtoReflect.Descriptor instead.
func (*ChunkRecord) Descriptor() ([]byte, []int) {
	return file_protocol_example_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChunkRecord) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChunkRecord) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChunkRequest) Reset() {
	*x = ChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkRequest) ProtoMessage() {}

func (x *ChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkRequest.ProtoReflect.Descriptor instead.
func (*ChunkRequest) Descriptor() ([]byte, []int) {
	return file_protocol_example_proto_rawDescGZIP(), []int{2}
}

func (x *ChunkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful bool           `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
	Cr         []*ChunkRecord `protobuf:"bytes,2,rep,name=cr,proto3" json:"cr,omitempty"`
}

func (x *ChunkResponse) Reset() {
	*x = ChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkResponse) ProtoMessage() {}

func (x *ChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkResponse.ProtoReflect.Descriptor instead.
func (*ChunkResponse) Descriptor() ([]byte, []int) {
	return file_protocol_example_proto_rawDescGZIP(), []int{3}
}

func (x *ChunkResponse) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

func (x *ChunkResponse) GetCr() []*ChunkRecord {
	if x != nil {
		return x.Cr
	}
	return nil
}

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fr *FileRecord `protobuf:"bytes,1,opt,name=fr,proto3" json:"fr,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_protocol_example_proto_rawDescGZIP(), []int{4}
}

func (x *FileRequest) GetFr() *FileRecord {
	if x != nil {
		return x.Fr
	}
	return nil
}

type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Fr      []*FileRecord `protobuf:"bytes,3,rep,name=fr,proto3" json:"fr,omitempty"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_example_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_example_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_protocol_example_proto_rawDescGZIP(), []int{5}
}

func (x *FileResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *FileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FileResponse) GetFr() []*FileRecord {
	if x != nil {
		return x.Fr
	}
	return nil
}

var File_protocol_example_proto protoreflect.FileDescriptor

var file_protocol_example_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe2, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x0b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22,
	0x1e, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x56, 0x0a, 0x0d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x12, 0x25, 0x0a, 0x02, 0x63, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x02, 0x63, 0x72, 0x22, 0x33, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x66, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x02, 0x66, 0x72, 0x22, 0x68, 0x0a, 0x0c,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x24, 0x0a, 0x02, 0x66, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x02, 0x66, 0x72, 0x32, 0xca, 0x03, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_example_proto_rawDescOnce sync.Once
	file_protocol_example_proto_rawDescData = file_protocol_example_proto_rawDesc
)

func file_protocol_example_proto_rawDescGZIP() []byte {
	file_protocol_example_proto_rawDescOnce.Do(func() {
		file_protocol_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_example_proto_rawDescData)
	})
	return file_protocol_example_proto_rawDescData
}

var file_protocol_example_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protocol_example_proto_goTypes = []interface{}{
	(*FileRecord)(nil),    // 0: protocol.FileRecord
	(*ChunkRecord)(nil),   // 1: protocol.ChunkRecord
	(*ChunkRequest)(nil),  // 2: protocol.ChunkRequest
	(*ChunkResponse)(nil), // 3: protocol.ChunkResponse
	(*FileRequest)(nil),   // 4: protocol.FileRequest
	(*FileResponse)(nil),  // 5: protocol.FileResponse
	(*empty.Empty)(nil),   // 6: google.protobuf.Empty
}
var file_protocol_example_proto_depIdxs = []int32{
	1,  // 0: protocol.ChunkResponse.cr:type_name -> protocol.ChunkRecord
	0,  // 1: protocol.FileRequest.fr:type_name -> protocol.FileRecord
	0,  // 2: protocol.FileResponse.fr:type_name -> protocol.FileRecord
	4,  // 3: protocol.DataInterface.GetByFilename:input_type -> protocol.FileRequest
	4,  // 4: protocol.DataInterface.GetById:input_type -> protocol.FileRequest
	6,  // 5: protocol.DataInterface.GetAllFiles:input_type -> google.protobuf.Empty
	4,  // 6: protocol.DataInterface.CreateFile:input_type -> protocol.FileRequest
	4,  // 7: protocol.DataInterface.UpdateFile:input_type -> protocol.FileRequest
	4,  // 8: protocol.DataInterface.DeleteFile:input_type -> protocol.FileRequest
	2,  // 9: protocol.DataInterface.GetChunk:input_type -> protocol.ChunkRequest
	5,  // 10: protocol.DataInterface.GetByFilename:output_type -> protocol.FileResponse
	5,  // 11: protocol.DataInterface.GetById:output_type -> protocol.FileResponse
	5,  // 12: protocol.DataInterface.GetAllFiles:output_type -> protocol.FileResponse
	5,  // 13: protocol.DataInterface.CreateFile:output_type -> protocol.FileResponse
	5,  // 14: protocol.DataInterface.UpdateFile:output_type -> protocol.FileResponse
	5,  // 15: protocol.DataInterface.DeleteFile:output_type -> protocol.FileResponse
	3,  // 16: protocol.DataInterface.GetChunk:output_type -> protocol.ChunkResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_protocol_example_proto_init() }
func file_protocol_example_proto_init() {
	if File_protocol_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRecord); i {
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
		file_protocol_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkRecord); i {
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
		file_protocol_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkRequest); i {
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
		file_protocol_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkResponse); i {
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
		file_protocol_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
		file_protocol_example_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponse); i {
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
			RawDescriptor: file_protocol_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_example_proto_goTypes,
		DependencyIndexes: file_protocol_example_proto_depIdxs,
		MessageInfos:      file_protocol_example_proto_msgTypes,
	}.Build()
	File_protocol_example_proto = out.File
	file_protocol_example_proto_rawDesc = nil
	file_protocol_example_proto_goTypes = nil
	file_protocol_example_proto_depIdxs = nil
}
