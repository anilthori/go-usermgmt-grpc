// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: usermgmt.proto

package usermgmt

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

type NewUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age    int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *NewUserRequest) Reset() {
	*x = NewUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUserRequest) ProtoMessage() {}

func (x *NewUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUserRequest.ProtoReflect.Descriptor instead.
func (*NewUserRequest) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{0}
}

func (x *NewUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *NewUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewUserRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type NewUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *NewUserResponse) Reset() {
	*x = NewUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUserResponse) ProtoMessage() {}

func (x *NewUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUserResponse.ProtoReflect.Descriptor instead.
func (*NewUserResponse) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{1}
}

func (x *NewUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserDetailsRequest) Reset() {
	*x = UserDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailsRequest) ProtoMessage() {}

func (x *UserDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailsRequest.ProtoReflect.Descriptor instead.
func (*UserDetailsRequest) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{2}
}

func (x *UserDetailsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserDetailsResponse) Reset() {
	*x = UserDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailsResponse) ProtoMessage() {}

func (x *UserDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailsResponse.ProtoReflect.Descriptor instead.
func (*UserDetailsResponse) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{3}
}

func (x *UserDetailsResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_usermgmt_proto protoreflect.FileDescriptor

var file_usermgmt_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x22, 0x4f, 0x0a, 0x0e, 0x4e, 0x65,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x4e,
	0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xa2, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_usermgmt_proto_rawDescOnce sync.Once
	file_usermgmt_proto_rawDescData = file_usermgmt_proto_rawDesc
)

func file_usermgmt_proto_rawDescGZIP() []byte {
	file_usermgmt_proto_rawDescOnce.Do(func() {
		file_usermgmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_usermgmt_proto_rawDescData)
	})
	return file_usermgmt_proto_rawDescData
}

var file_usermgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_usermgmt_proto_goTypes = []interface{}{
	(*NewUserRequest)(nil),      // 0: usermgmt.NewUserRequest
	(*NewUserResponse)(nil),     // 1: usermgmt.NewUserResponse
	(*UserDetailsRequest)(nil),  // 2: usermgmt.UserDetailsRequest
	(*UserDetailsResponse)(nil), // 3: usermgmt.UserDetailsResponse
}
var file_usermgmt_proto_depIdxs = []int32{
	0, // 0: usermgmt.UserManagement.CreateNewUser:input_type -> usermgmt.NewUserRequest
	2, // 1: usermgmt.UserManagement.GetUser:input_type -> usermgmt.UserDetailsRequest
	1, // 2: usermgmt.UserManagement.CreateNewUser:output_type -> usermgmt.NewUserResponse
	3, // 3: usermgmt.UserManagement.GetUser:output_type -> usermgmt.UserDetailsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_usermgmt_proto_init() }
func file_usermgmt_proto_init() {
	if File_usermgmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usermgmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUserRequest); i {
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
		file_usermgmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUserResponse); i {
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
		file_usermgmt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailsRequest); i {
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
		file_usermgmt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetailsResponse); i {
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
			RawDescriptor: file_usermgmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usermgmt_proto_goTypes,
		DependencyIndexes: file_usermgmt_proto_depIdxs,
		MessageInfos:      file_usermgmt_proto_msgTypes,
	}.Build()
	File_usermgmt_proto = out.File
	file_usermgmt_proto_rawDesc = nil
	file_usermgmt_proto_goTypes = nil
	file_usermgmt_proto_depIdxs = nil
}
