// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: trisads/admin/v1alpha1/admin.proto

package admin

import (
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "github.com/trisacrypto/testnet/pkg/trisads/pb/api/v1alpha1"
	v1alpha11 "github.com/trisacrypto/testnet/pkg/trisads/pb/models/v1alpha1"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Registration review requests are sent via email to the TRISA admin email address with
// a lightweight token for review. This endpoint allows administrators to submit a review
// determination back to the directory server.
type ReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the VASP to perform the review for.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The verification token sent in the review email.
	// This token provides lightweight authentication but should be replaced with a more
	// robust authentication and authorization scheme.
	AdminVerificationToken string `protobuf:"bytes,2,opt,name=admin_verification_token,json=adminVerificationToken,proto3" json:"admin_verification_token,omitempty"`
	// If accept is false then the request will be rejected and a reject reason must be
	// specified. If it is true, then the certificate issuance process will begin.
	Accept       bool   `protobuf:"varint,3,opt,name=accept,proto3" json:"accept,omitempty"`
	RejectReason string `protobuf:"bytes,4,opt,name=reject_reason,json=rejectReason,proto3" json:"reject_reason,omitempty"`
}

func (x *ReviewRequest) Reset() {
	*x = ReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trisads_admin_v1alpha1_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewRequest) ProtoMessage() {}

func (x *ReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trisads_admin_v1alpha1_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewRequest.ProtoReflect.Descriptor instead.
func (*ReviewRequest) Descriptor() ([]byte, []int) {
	return file_trisads_admin_v1alpha1_admin_proto_rawDescGZIP(), []int{0}
}

func (x *ReviewRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReviewRequest) GetAdminVerificationToken() string {
	if x != nil {
		return x.AdminVerificationToken
	}
	return ""
}

func (x *ReviewRequest) GetAccept() bool {
	if x != nil {
		return x.Accept
	}
	return false
}

func (x *ReviewRequest) GetRejectReason() string {
	if x != nil {
		return x.RejectReason
	}
	return ""
}

type ReviewReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If no error is specified, the verify email request was successful
	Error *v1alpha1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// The verification status of the VASP entity.
	Status  v1alpha11.VerificationState `protobuf:"varint,2,opt,name=status,proto3,enum=trisads.models.v1alpha1.VerificationState" json:"status,omitempty"`
	Message string                      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReviewReply) Reset() {
	*x = ReviewReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trisads_admin_v1alpha1_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewReply) ProtoMessage() {}

func (x *ReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_trisads_admin_v1alpha1_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewReply.ProtoReflect.Descriptor instead.
func (*ReviewReply) Descriptor() ([]byte, []int) {
	return file_trisads_admin_v1alpha1_admin_proto_rawDescGZIP(), []int{1}
}

func (x *ReviewReply) GetError() *v1alpha1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ReviewReply) GetStatus() v1alpha11.VerificationState {
	if x != nil {
		return x.Status
	}
	return v1alpha11.VerificationState_NO_VERIFICATION
}

func (x *ReviewReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_trisads_admin_v1alpha1_admin_proto protoreflect.FileDescriptor

var file_trisads_admin_v1alpha1_admin_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x72, 0x69, 0x73, 0x61, 0x64, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x72, 0x69, 0x73, 0x61, 0x64, 0x73, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x74, 0x72,
	0x69, 0x73, 0x61, 0x64, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x74, 0x72,
	0x69, 0x73, 0x61, 0x64, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x9e, 0x01, 0x0a, 0x0b,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x69,
	0x73, 0x61, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x42,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a,
	0x2e, 0x74, 0x72, 0x69, 0x73, 0x61, 0x64, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x71, 0x0a, 0x17,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x25, 0x2e, 0x74, 0x72, 0x69, 0x73, 0x61, 0x64, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74, 0x72, 0x69, 0x73, 0x61,
	0x64, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72,
	0x69, 0x73, 0x61, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x6e, 0x65,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x69, 0x73, 0x61, 0x64, 0x73, 0x2f, 0x70, 0x62,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trisads_admin_v1alpha1_admin_proto_rawDescOnce sync.Once
	file_trisads_admin_v1alpha1_admin_proto_rawDescData = file_trisads_admin_v1alpha1_admin_proto_rawDesc
)

func file_trisads_admin_v1alpha1_admin_proto_rawDescGZIP() []byte {
	file_trisads_admin_v1alpha1_admin_proto_rawDescOnce.Do(func() {
		file_trisads_admin_v1alpha1_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_trisads_admin_v1alpha1_admin_proto_rawDescData)
	})
	return file_trisads_admin_v1alpha1_admin_proto_rawDescData
}

var file_trisads_admin_v1alpha1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_trisads_admin_v1alpha1_admin_proto_goTypes = []interface{}{
	(*ReviewRequest)(nil),            // 0: trisads.admin.v1alpha1.ReviewRequest
	(*ReviewReply)(nil),              // 1: trisads.admin.v1alpha1.ReviewReply
	(*v1alpha1.Error)(nil),           // 2: trisads.api.v1alpha1.Error
	(v1alpha11.VerificationState)(0), // 3: trisads.models.v1alpha1.VerificationState
}
var file_trisads_admin_v1alpha1_admin_proto_depIdxs = []int32{
	2, // 0: trisads.admin.v1alpha1.ReviewReply.error:type_name -> trisads.api.v1alpha1.Error
	3, // 1: trisads.admin.v1alpha1.ReviewReply.status:type_name -> trisads.models.v1alpha1.VerificationState
	0, // 2: trisads.admin.v1alpha1.DirectoryAdministration.Review:input_type -> trisads.admin.v1alpha1.ReviewRequest
	1, // 3: trisads.admin.v1alpha1.DirectoryAdministration.Review:output_type -> trisads.admin.v1alpha1.ReviewReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_trisads_admin_v1alpha1_admin_proto_init() }
func file_trisads_admin_v1alpha1_admin_proto_init() {
	if File_trisads_admin_v1alpha1_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trisads_admin_v1alpha1_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewRequest); i {
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
		file_trisads_admin_v1alpha1_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewReply); i {
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
			RawDescriptor: file_trisads_admin_v1alpha1_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trisads_admin_v1alpha1_admin_proto_goTypes,
		DependencyIndexes: file_trisads_admin_v1alpha1_admin_proto_depIdxs,
		MessageInfos:      file_trisads_admin_v1alpha1_admin_proto_msgTypes,
	}.Build()
	File_trisads_admin_v1alpha1_admin_proto = out.File
	file_trisads_admin_v1alpha1_admin_proto_rawDesc = nil
	file_trisads_admin_v1alpha1_admin_proto_goTypes = nil
	file_trisads_admin_v1alpha1_admin_proto_depIdxs = nil
}
