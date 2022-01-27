// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: proto/krs.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MataKuliah struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kode     string `protobuf:"bytes,1,opt,name=kode,proto3" json:"kode,omitempty"`
	Nama     string `protobuf:"bytes,2,opt,name=nama,proto3" json:"nama,omitempty"`
	Sks      int32  `protobuf:"varint,3,opt,name=sks,proto3" json:"sks,omitempty"`
	Dosen    string `protobuf:"bytes,4,opt,name=dosen,proto3" json:"dosen,omitempty"`
	Semester string `protobuf:"bytes,5,opt,name=semester,proto3" json:"semester,omitempty"`
}

func (x *MataKuliah) Reset() {
	*x = MataKuliah{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_krs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MataKuliah) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MataKuliah) ProtoMessage() {}

func (x *MataKuliah) ProtoReflect() protoreflect.Message {
	mi := &file_proto_krs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MataKuliah.ProtoReflect.Descriptor instead.
func (*MataKuliah) Descriptor() ([]byte, []int) {
	return file_proto_krs_proto_rawDescGZIP(), []int{0}
}

func (x *MataKuliah) GetKode() string {
	if x != nil {
		return x.Kode
	}
	return ""
}

func (x *MataKuliah) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *MataKuliah) GetSks() int32 {
	if x != nil {
		return x.Sks
	}
	return 0
}

func (x *MataKuliah) GetDosen() string {
	if x != nil {
		return x.Dosen
	}
	return ""
}

func (x *MataKuliah) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

type CreateUpdateKRSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IdMahasiswa int32         `protobuf:"varint,2,opt,name=id_mahasiswa,json=idMahasiswa,proto3" json:"id_mahasiswa,omitempty"`
	MataKuliahs []*MataKuliah `protobuf:"bytes,3,rep,name=mata_kuliahs,json=mataKuliahs,proto3" json:"mata_kuliahs,omitempty"`
}

func (x *CreateUpdateKRSRequest) Reset() {
	*x = CreateUpdateKRSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_krs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUpdateKRSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUpdateKRSRequest) ProtoMessage() {}

func (x *CreateUpdateKRSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_krs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUpdateKRSRequest.ProtoReflect.Descriptor instead.
func (*CreateUpdateKRSRequest) Descriptor() ([]byte, []int) {
	return file_proto_krs_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUpdateKRSRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateUpdateKRSRequest) GetIdMahasiswa() int32 {
	if x != nil {
		return x.IdMahasiswa
	}
	return 0
}

func (x *CreateUpdateKRSRequest) GetMataKuliahs() []*MataKuliah {
	if x != nil {
		return x.MataKuliahs
	}
	return nil
}

type ReadKRSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IdMahasiswa int32  `protobuf:"varint,2,opt,name=id_mahasiswa,json=idMahasiswa,proto3" json:"id_mahasiswa,omitempty"`
}

func (x *ReadKRSRequest) Reset() {
	*x = ReadKRSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_krs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadKRSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadKRSRequest) ProtoMessage() {}

func (x *ReadKRSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_krs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadKRSRequest.ProtoReflect.Descriptor instead.
func (*ReadKRSRequest) Descriptor() ([]byte, []int) {
	return file_proto_krs_proto_rawDescGZIP(), []int{2}
}

func (x *ReadKRSRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ReadKRSRequest) GetIdMahasiswa() int32 {
	if x != nil {
		return x.IdMahasiswa
	}
	return 0
}

type DeleteKRSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IdMahasiswa  int32  `protobuf:"varint,2,opt,name=id_mahasiswa,json=idMahasiswa,proto3" json:"id_mahasiswa,omitempty"`
	IdMataKuliah int32  `protobuf:"varint,3,opt,name=id_mata_kuliah,json=idMataKuliah,proto3" json:"id_mata_kuliah,omitempty"`
}

func (x *DeleteKRSRequest) Reset() {
	*x = DeleteKRSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_krs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKRSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKRSRequest) ProtoMessage() {}

func (x *DeleteKRSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_krs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKRSRequest.ProtoReflect.Descriptor instead.
func (*DeleteKRSRequest) Descriptor() ([]byte, []int) {
	return file_proto_krs_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteKRSRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteKRSRequest) GetIdMahasiswa() int32 {
	if x != nil {
		return x.IdMahasiswa
	}
	return 0
}

func (x *DeleteKRSRequest) GetIdMataKuliah() int32 {
	if x != nil {
		return x.IdMataKuliah
	}
	return 0
}

type KRSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MataKuliahs []*MataKuliah `protobuf:"bytes,1,rep,name=mata_kuliahs,json=mataKuliahs,proto3" json:"mata_kuliahs,omitempty"`
}

func (x *KRSResponse) Reset() {
	*x = KRSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_krs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KRSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KRSResponse) ProtoMessage() {}

func (x *KRSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_krs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KRSResponse.ProtoReflect.Descriptor instead.
func (*KRSResponse) Descriptor() ([]byte, []int) {
	return file_proto_krs_proto_rawDescGZIP(), []int{4}
}

func (x *KRSResponse) GetMataKuliahs() []*MataKuliah {
	if x != nil {
		return x.MataKuliahs
	}
	return nil
}

type DeleteKRSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteKRSResponse) Reset() {
	*x = DeleteKRSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_krs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKRSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKRSResponse) ProtoMessage() {}

func (x *DeleteKRSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_krs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKRSResponse.ProtoReflect.Descriptor instead.
func (*DeleteKRSResponse) Descriptor() ([]byte, []int) {
	return file_proto_krs_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteKRSResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_krs_proto protoreflect.FileDescriptor

var file_proto_krs_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x61,
	0x4b, 0x75, 0x6c, 0x69, 0x61, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x6b, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x6f, 0x73, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x64, 0x6f, 0x73, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x22, 0x87, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4b, 0x52, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x68, 0x61, 0x73, 0x69,
	0x73, 0x77, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x64, 0x4d, 0x61, 0x68,
	0x61, 0x73, 0x69, 0x73, 0x77, 0x61, 0x12, 0x34, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x61, 0x5f, 0x6b,
	0x75, 0x6c, 0x69, 0x61, 0x68, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x61, 0x4b, 0x75, 0x6c, 0x69, 0x61, 0x68, 0x52,
	0x0b, 0x6d, 0x61, 0x74, 0x61, 0x4b, 0x75, 0x6c, 0x69, 0x61, 0x68, 0x73, 0x22, 0x49, 0x0a, 0x0e,
	0x52, 0x65, 0x61, 0x64, 0x4b, 0x52, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x68, 0x61, 0x73,
	0x69, 0x73, 0x77, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x64, 0x4d, 0x61,
	0x68, 0x61, 0x73, 0x69, 0x73, 0x77, 0x61, 0x22, 0x71, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x52, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x68, 0x61, 0x73, 0x69, 0x73, 0x77,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x64, 0x4d, 0x61, 0x68, 0x61, 0x73,
	0x69, 0x73, 0x77, 0x61, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x74, 0x61, 0x5f,
	0x6b, 0x75, 0x6c, 0x69, 0x61, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x64,
	0x4d, 0x61, 0x74, 0x61, 0x4b, 0x75, 0x6c, 0x69, 0x61, 0x68, 0x22, 0x43, 0x0a, 0x0b, 0x4b, 0x52,
	0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x6d, 0x61, 0x74,
	0x61, 0x5f, 0x6b, 0x75, 0x6c, 0x69, 0x61, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x61, 0x4b, 0x75, 0x6c, 0x69,
	0x61, 0x68, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x61, 0x4b, 0x75, 0x6c, 0x69, 0x61, 0x68, 0x73, 0x22,
	0x2b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x52, 0x53, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xf6, 0x01, 0x0a,
	0x0a, 0x4b, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x52,
	0x65, 0x61, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x4b, 0x52, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4b, 0x52, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x52, 0x53,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4b, 0x52, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x52, 0x53, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x52, 0x53,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x52, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x52, 0x53, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x50, 0x01, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_krs_proto_rawDescOnce sync.Once
	file_proto_krs_proto_rawDescData = file_proto_krs_proto_rawDesc
)

func file_proto_krs_proto_rawDescGZIP() []byte {
	file_proto_krs_proto_rawDescOnce.Do(func() {
		file_proto_krs_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_krs_proto_rawDescData)
	})
	return file_proto_krs_proto_rawDescData
}

var file_proto_krs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_krs_proto_goTypes = []interface{}{
	(*MataKuliah)(nil),             // 0: proto.MataKuliah
	(*CreateUpdateKRSRequest)(nil), // 1: proto.CreateUpdateKRSRequest
	(*ReadKRSRequest)(nil),         // 2: proto.ReadKRSRequest
	(*DeleteKRSRequest)(nil),       // 3: proto.DeleteKRSRequest
	(*KRSResponse)(nil),            // 4: proto.KRSResponse
	(*DeleteKRSResponse)(nil),      // 5: proto.DeleteKRSResponse
}
var file_proto_krs_proto_depIdxs = []int32{
	0, // 0: proto.CreateUpdateKRSRequest.mata_kuliahs:type_name -> proto.MataKuliah
	0, // 1: proto.KRSResponse.mata_kuliahs:type_name -> proto.MataKuliah
	2, // 2: proto.KrsService.Read:input_type -> proto.ReadKRSRequest
	1, // 3: proto.KrsService.Create:input_type -> proto.CreateUpdateKRSRequest
	1, // 4: proto.KrsService.Update:input_type -> proto.CreateUpdateKRSRequest
	3, // 5: proto.KrsService.Delete:input_type -> proto.DeleteKRSRequest
	4, // 6: proto.KrsService.Read:output_type -> proto.KRSResponse
	4, // 7: proto.KrsService.Create:output_type -> proto.KRSResponse
	4, // 8: proto.KrsService.Update:output_type -> proto.KRSResponse
	5, // 9: proto.KrsService.Delete:output_type -> proto.DeleteKRSResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_krs_proto_init() }
func file_proto_krs_proto_init() {
	if File_proto_krs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_krs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MataKuliah); i {
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
		file_proto_krs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUpdateKRSRequest); i {
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
		file_proto_krs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadKRSRequest); i {
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
		file_proto_krs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKRSRequest); i {
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
		file_proto_krs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KRSResponse); i {
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
		file_proto_krs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKRSResponse); i {
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
			RawDescriptor: file_proto_krs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_krs_proto_goTypes,
		DependencyIndexes: file_proto_krs_proto_depIdxs,
		MessageInfos:      file_proto_krs_proto_msgTypes,
	}.Build()
	File_proto_krs_proto = out.File
	file_proto_krs_proto_rawDesc = nil
	file_proto_krs_proto_goTypes = nil
	file_proto_krs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KrsServiceClient is the client API for KrsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KrsServiceClient interface {
	Read(ctx context.Context, in *ReadKRSRequest, opts ...grpc.CallOption) (*KRSResponse, error)
	Create(ctx context.Context, in *CreateUpdateKRSRequest, opts ...grpc.CallOption) (*KRSResponse, error)
	Update(ctx context.Context, in *CreateUpdateKRSRequest, opts ...grpc.CallOption) (*KRSResponse, error)
	Delete(ctx context.Context, in *DeleteKRSRequest, opts ...grpc.CallOption) (*DeleteKRSResponse, error)
}

type krsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKrsServiceClient(cc grpc.ClientConnInterface) KrsServiceClient {
	return &krsServiceClient{cc}
}

func (c *krsServiceClient) Read(ctx context.Context, in *ReadKRSRequest, opts ...grpc.CallOption) (*KRSResponse, error) {
	out := new(KRSResponse)
	err := c.cc.Invoke(ctx, "/proto.KrsService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *krsServiceClient) Create(ctx context.Context, in *CreateUpdateKRSRequest, opts ...grpc.CallOption) (*KRSResponse, error) {
	out := new(KRSResponse)
	err := c.cc.Invoke(ctx, "/proto.KrsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *krsServiceClient) Update(ctx context.Context, in *CreateUpdateKRSRequest, opts ...grpc.CallOption) (*KRSResponse, error) {
	out := new(KRSResponse)
	err := c.cc.Invoke(ctx, "/proto.KrsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *krsServiceClient) Delete(ctx context.Context, in *DeleteKRSRequest, opts ...grpc.CallOption) (*DeleteKRSResponse, error) {
	out := new(DeleteKRSResponse)
	err := c.cc.Invoke(ctx, "/proto.KrsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KrsServiceServer is the server API for KrsService service.
type KrsServiceServer interface {
	Read(context.Context, *ReadKRSRequest) (*KRSResponse, error)
	Create(context.Context, *CreateUpdateKRSRequest) (*KRSResponse, error)
	Update(context.Context, *CreateUpdateKRSRequest) (*KRSResponse, error)
	Delete(context.Context, *DeleteKRSRequest) (*DeleteKRSResponse, error)
}

// UnimplementedKrsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKrsServiceServer struct {
}

func (*UnimplementedKrsServiceServer) Read(context.Context, *ReadKRSRequest) (*KRSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedKrsServiceServer) Create(context.Context, *CreateUpdateKRSRequest) (*KRSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedKrsServiceServer) Update(context.Context, *CreateUpdateKRSRequest) (*KRSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedKrsServiceServer) Delete(context.Context, *DeleteKRSRequest) (*DeleteKRSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterKrsServiceServer(s *grpc.Server, srv KrsServiceServer) {
	s.RegisterService(&_KrsService_serviceDesc, srv)
}

func _KrsService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadKRSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KrsServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KrsService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KrsServiceServer).Read(ctx, req.(*ReadKRSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KrsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUpdateKRSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KrsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KrsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KrsServiceServer).Create(ctx, req.(*CreateUpdateKRSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KrsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUpdateKRSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KrsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KrsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KrsServiceServer).Update(ctx, req.(*CreateUpdateKRSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KrsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKRSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KrsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KrsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KrsServiceServer).Delete(ctx, req.(*DeleteKRSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KrsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.KrsService",
	HandlerType: (*KrsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _KrsService_Read_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _KrsService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _KrsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KrsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/krs.proto",
}
