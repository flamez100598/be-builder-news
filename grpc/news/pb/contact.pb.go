// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: grpc/news/proto/msg/contact.proto

package pb

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

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Subject   string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Msg       string `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	Status    int32  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Id        string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_news_proto_msg_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_news_proto_msg_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_grpc_news_proto_msg_contact_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Contact) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Contact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contact) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Contact) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Contact) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Contact) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Contact) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ContactReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone   string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Subject string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Msg     string `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	Offset  int32  `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit   int32  `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ContactReq) Reset() {
	*x = ContactReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_news_proto_msg_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactReq) ProtoMessage() {}

func (x *ContactReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_news_proto_msg_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactReq.ProtoReflect.Descriptor instead.
func (*ContactReq) Descriptor() ([]byte, []int) {
	return file_grpc_news_proto_msg_contact_proto_rawDescGZIP(), []int{1}
}

func (x *ContactReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ContactReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ContactReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContactReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ContactReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ContactReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ContactReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ContactResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Contact *Contact `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *ContactResp) Reset() {
	*x = ContactResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_news_proto_msg_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactResp) ProtoMessage() {}

func (x *ContactResp) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_news_proto_msg_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactResp.ProtoReflect.Descriptor instead.
func (*ContactResp) Descriptor() ([]byte, []int) {
	return file_grpc_news_proto_msg_contact_proto_rawDescGZIP(), []int{2}
}

func (x *ContactResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ContactResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ContactResp) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

type ContactsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Contacts []*Contact `protobuf:"bytes,3,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *ContactsResp) Reset() {
	*x = ContactsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_news_proto_msg_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactsResp) ProtoMessage() {}

func (x *ContactsResp) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_news_proto_msg_contact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactsResp.ProtoReflect.Descriptor instead.
func (*ContactsResp) Descriptor() ([]byte, []int) {
	return file_grpc_news_proto_msg_contact_proto_rawDescGZIP(), []int{3}
}

func (x *ContactsResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ContactsResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ContactsResp) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

var File_grpc_news_proto_msg_contact_proto protoreflect.FileDescriptor

var file_grpc_news_proto_msg_contact_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa6, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5a,
	0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x5d, 0x0a, 0x0c, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x27, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_grpc_news_proto_msg_contact_proto_rawDescOnce sync.Once
	file_grpc_news_proto_msg_contact_proto_rawDescData = file_grpc_news_proto_msg_contact_proto_rawDesc
)

func file_grpc_news_proto_msg_contact_proto_rawDescGZIP() []byte {
	file_grpc_news_proto_msg_contact_proto_rawDescOnce.Do(func() {
		file_grpc_news_proto_msg_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_news_proto_msg_contact_proto_rawDescData)
	})
	return file_grpc_news_proto_msg_contact_proto_rawDescData
}

var file_grpc_news_proto_msg_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_news_proto_msg_contact_proto_goTypes = []interface{}{
	(*Contact)(nil),      // 0: pb.Contact
	(*ContactReq)(nil),   // 1: pb.ContactReq
	(*ContactResp)(nil),  // 2: pb.ContactResp
	(*ContactsResp)(nil), // 3: pb.ContactsResp
}
var file_grpc_news_proto_msg_contact_proto_depIdxs = []int32{
	0, // 0: pb.ContactResp.contact:type_name -> pb.Contact
	0, // 1: pb.ContactsResp.contacts:type_name -> pb.Contact
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_news_proto_msg_contact_proto_init() }
func file_grpc_news_proto_msg_contact_proto_init() {
	if File_grpc_news_proto_msg_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_news_proto_msg_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_grpc_news_proto_msg_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactReq); i {
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
		file_grpc_news_proto_msg_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactResp); i {
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
		file_grpc_news_proto_msg_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactsResp); i {
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
			RawDescriptor: file_grpc_news_proto_msg_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_news_proto_msg_contact_proto_goTypes,
		DependencyIndexes: file_grpc_news_proto_msg_contact_proto_depIdxs,
		MessageInfos:      file_grpc_news_proto_msg_contact_proto_msgTypes,
	}.Build()
	File_grpc_news_proto_msg_contact_proto = out.File
	file_grpc_news_proto_msg_contact_proto_rawDesc = nil
	file_grpc_news_proto_msg_contact_proto_goTypes = nil
	file_grpc_news_proto_msg_contact_proto_depIdxs = nil
}