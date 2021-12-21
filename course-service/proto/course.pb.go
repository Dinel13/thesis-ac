// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/course.proto

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

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Teacher     int32  `protobuf:"varint,4,opt,name=teacher,proto3" json:"teacher,omitempty"`
	Capacity    int32  `protobuf:"varint,5,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Remain      int32  `protobuf:"varint,6,opt,name=remain,proto3" json:"remain,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_proto_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Course) GetTeacher() int32 {
	if x != nil {
		return x.Teacher
	}
	return 0
}

func (x *Course) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Course) GetRemain() int32 {
	if x != nil {
		return x.Remain
	}
	return 0
}

func (x *Course) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Course) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses *Course `protobuf:"bytes,1,opt,name=courses,proto3" json:"courses,omitempty"`
}

func (x *CourseResponse) Reset() {
	*x = CourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseResponse) ProtoMessage() {}

func (x *CourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseResponse.ProtoReflect.Descriptor instead.
func (*CourseResponse) Descriptor() ([]byte, []int) {
	return file_proto_course_proto_rawDescGZIP(), []int{1}
}

func (x *CourseResponse) GetCourses() *Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

type CourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CourseRequest) Reset() {
	*x = CourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseRequest) ProtoMessage() {}

func (x *CourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseRequest.ProtoReflect.Descriptor instead.
func (*CourseRequest) Descriptor() ([]byte, []int) {
	return file_proto_course_proto_rawDescGZIP(), []int{2}
}

func (x *CourseRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_course_proto protoreflect.FileDescriptor

var file_proto_course_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x39, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x44, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_course_proto_rawDescOnce sync.Once
	file_proto_course_proto_rawDescData = file_proto_course_proto_rawDesc
)

func file_proto_course_proto_rawDescGZIP() []byte {
	file_proto_course_proto_rawDescOnce.Do(func() {
		file_proto_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_course_proto_rawDescData)
	})
	return file_proto_course_proto_rawDescData
}

var file_proto_course_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_course_proto_goTypes = []interface{}{
	(*Course)(nil),         // 0: proto.Course
	(*CourseResponse)(nil), // 1: proto.CourseResponse
	(*CourseRequest)(nil),  // 2: proto.CourseRequest
}
var file_proto_course_proto_depIdxs = []int32{
	0, // 0: proto.CourseResponse.courses:type_name -> proto.Course
	2, // 1: proto.CourseService.Read:input_type -> proto.CourseRequest
	1, // 2: proto.CourseService.Read:output_type -> proto.CourseResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_course_proto_init() }
func file_proto_course_proto_init() {
	if File_proto_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
		file_proto_course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseResponse); i {
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
		file_proto_course_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseRequest); i {
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
			RawDescriptor: file_proto_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_course_proto_goTypes,
		DependencyIndexes: file_proto_course_proto_depIdxs,
		MessageInfos:      file_proto_course_proto_msgTypes,
	}.Build()
	File_proto_course_proto = out.File
	file_proto_course_proto_rawDesc = nil
	file_proto_course_proto_goTypes = nil
	file_proto_course_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CourseServiceClient is the client API for CourseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CourseServiceClient interface {
	Read(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CourseResponse, error)
}

type courseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseServiceClient(cc grpc.ClientConnInterface) CourseServiceClient {
	return &courseServiceClient{cc}
}

func (c *courseServiceClient) Read(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CourseResponse, error) {
	out := new(CourseResponse)
	err := c.cc.Invoke(ctx, "/proto.CourseService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServiceServer is the server API for CourseService service.
type CourseServiceServer interface {
	Read(context.Context, *CourseRequest) (*CourseResponse, error)
}

// UnimplementedCourseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCourseServiceServer struct {
}

func (*UnimplementedCourseServiceServer) Read(context.Context, *CourseRequest) (*CourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func RegisterCourseServiceServer(s *grpc.Server, srv CourseServiceServer) {
	s.RegisterService(&_CourseService_serviceDesc, srv)
}

func _CourseService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CourseService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).Read(ctx, req.(*CourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CourseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CourseService",
	HandlerType: (*CourseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _CourseService_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/course.proto",
}
