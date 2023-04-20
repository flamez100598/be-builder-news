// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: grpc/news/proto/msg/news_view.proto

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

type NewsView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title             string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content           string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Description       string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ImgUrl            string  `protobuf:"bytes,5,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
	MetaKw            string  `protobuf:"bytes,6,opt,name=metaKw,proto3" json:"metaKw,omitempty"`
	MetaDesc          string  `protobuf:"bytes,7,opt,name=metaDesc,proto3" json:"metaDesc,omitempty"`
	Slug              string  `protobuf:"bytes,8,opt,name=slug,proto3" json:"slug,omitempty"`
	Category          int32   `protobuf:"varint,9,opt,name=category,proto3" json:"category,omitempty"`
	SubCategory       int32   `protobuf:"varint,10,opt,name=subCategory,proto3" json:"subCategory,omitempty"`
	CommentNum        int32   `protobuf:"varint,11,opt,name=commentNum,proto3" json:"commentNum,omitempty"`
	VoteNum           int32   `protobuf:"varint,12,opt,name=voteNum,proto3" json:"voteNum,omitempty"`
	ViewNum           int32   `protobuf:"varint,13,opt,name=viewNum,proto3" json:"viewNum,omitempty"`
	Status            int32   `protobuf:"varint,14,opt,name=status,proto3" json:"status,omitempty"`
	PublishBy         string  `protobuf:"bytes,15,opt,name=publishBy,proto3" json:"publishBy,omitempty"`
	Ranking           float64 `protobuf:"fixed64,16,opt,name=ranking,proto3" json:"ranking,omitempty"`
	CreatedAt         string  `protobuf:"bytes,17,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt         string  `protobuf:"bytes,18,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	PublishedAt       string  `protobuf:"bytes,19,opt,name=publishedAt,proto3" json:"publishedAt,omitempty"`
	Order             int32   `protobuf:"varint,20,opt,name=order,proto3" json:"order,omitempty"`
	UsernamePublishBy string  `protobuf:"bytes,21,opt,name=usernamePublishBy,proto3" json:"usernamePublishBy,omitempty"`
	NamePublishBy     string  `protobuf:"bytes,22,opt,name=namePublishBy,proto3" json:"namePublishBy,omitempty"`
	AvatarPublishBy   string  `protobuf:"bytes,23,opt,name=avatarPublishBy,proto3" json:"avatarPublishBy,omitempty"`
	UsernameUpdateBy  string  `protobuf:"bytes,24,opt,name=usernameUpdateBy,proto3" json:"usernameUpdateBy,omitempty"`
	NameUpdateBy      string  `protobuf:"bytes,25,opt,name=nameUpdateBy,proto3" json:"nameUpdateBy,omitempty"`
	AvatarUpdateBy    string  `protobuf:"bytes,26,opt,name=avatarUpdateBy,proto3" json:"avatarUpdateBy,omitempty"`
	ContentJp         string  `protobuf:"bytes,27,opt,name=contentJp,proto3" json:"contentJp,omitempty"`
	DescriptionJp     string  `protobuf:"bytes,28,opt,name=descriptionJp,proto3" json:"descriptionJp,omitempty"`
}

func (x *NewsView) Reset() {
	*x = NewsView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_news_proto_msg_news_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsView) ProtoMessage() {}

func (x *NewsView) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_news_proto_msg_news_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsView.ProtoReflect.Descriptor instead.
func (*NewsView) Descriptor() ([]byte, []int) {
	return file_grpc_news_proto_msg_news_view_proto_rawDescGZIP(), []int{0}
}

func (x *NewsView) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NewsView) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsView) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NewsView) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewsView) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *NewsView) GetMetaKw() string {
	if x != nil {
		return x.MetaKw
	}
	return ""
}

func (x *NewsView) GetMetaDesc() string {
	if x != nil {
		return x.MetaDesc
	}
	return ""
}

func (x *NewsView) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *NewsView) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *NewsView) GetSubCategory() int32 {
	if x != nil {
		return x.SubCategory
	}
	return 0
}

func (x *NewsView) GetCommentNum() int32 {
	if x != nil {
		return x.CommentNum
	}
	return 0
}

func (x *NewsView) GetVoteNum() int32 {
	if x != nil {
		return x.VoteNum
	}
	return 0
}

func (x *NewsView) GetViewNum() int32 {
	if x != nil {
		return x.ViewNum
	}
	return 0
}

func (x *NewsView) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *NewsView) GetPublishBy() string {
	if x != nil {
		return x.PublishBy
	}
	return ""
}

func (x *NewsView) GetRanking() float64 {
	if x != nil {
		return x.Ranking
	}
	return 0
}

func (x *NewsView) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NewsView) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *NewsView) GetPublishedAt() string {
	if x != nil {
		return x.PublishedAt
	}
	return ""
}

func (x *NewsView) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *NewsView) GetUsernamePublishBy() string {
	if x != nil {
		return x.UsernamePublishBy
	}
	return ""
}

func (x *NewsView) GetNamePublishBy() string {
	if x != nil {
		return x.NamePublishBy
	}
	return ""
}

func (x *NewsView) GetAvatarPublishBy() string {
	if x != nil {
		return x.AvatarPublishBy
	}
	return ""
}

func (x *NewsView) GetUsernameUpdateBy() string {
	if x != nil {
		return x.UsernameUpdateBy
	}
	return ""
}

func (x *NewsView) GetNameUpdateBy() string {
	if x != nil {
		return x.NameUpdateBy
	}
	return ""
}

func (x *NewsView) GetAvatarUpdateBy() string {
	if x != nil {
		return x.AvatarUpdateBy
	}
	return ""
}

func (x *NewsView) GetContentJp() string {
	if x != nil {
		return x.ContentJp
	}
	return ""
}

func (x *NewsView) GetDescriptionJp() string {
	if x != nil {
		return x.DescriptionJp
	}
	return ""
}

var File_grpc_news_proto_msg_news_view_proto protoreflect.FileDescriptor

var file_grpc_news_proto_msg_news_view_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xdc, 0x06, 0x0a, 0x08, 0x4e, 0x65,
	0x77, 0x73, 0x56, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x67, 0x55,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x61, 0x4b, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x61, 0x4b, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x44, 0x65, 0x73, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x6f, 0x74, 0x65, 0x4e, 0x75,
	0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x6f, 0x74, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x65, 0x77, 0x4e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x76, 0x69, 0x65, 0x77, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x42, 0x79, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x42, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x41, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x42, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x42, 0x79, 0x12, 0x24, 0x0a, 0x0d,
	0x6e, 0x61, 0x6d, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x42, 0x79, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x42, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x42, 0x79, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x42, 0x79, 0x12, 0x2a, 0x0a, 0x10,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6e, 0x61, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x26, 0x0a, 0x0e,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4a,
	0x70, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x4a, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4a, 0x70, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x70, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_news_proto_msg_news_view_proto_rawDescOnce sync.Once
	file_grpc_news_proto_msg_news_view_proto_rawDescData = file_grpc_news_proto_msg_news_view_proto_rawDesc
)

func file_grpc_news_proto_msg_news_view_proto_rawDescGZIP() []byte {
	file_grpc_news_proto_msg_news_view_proto_rawDescOnce.Do(func() {
		file_grpc_news_proto_msg_news_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_news_proto_msg_news_view_proto_rawDescData)
	})
	return file_grpc_news_proto_msg_news_view_proto_rawDescData
}

var file_grpc_news_proto_msg_news_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_news_proto_msg_news_view_proto_goTypes = []interface{}{
	(*NewsView)(nil), // 0: pb.NewsView
}
var file_grpc_news_proto_msg_news_view_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_news_proto_msg_news_view_proto_init() }
func file_grpc_news_proto_msg_news_view_proto_init() {
	if File_grpc_news_proto_msg_news_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_news_proto_msg_news_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsView); i {
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
			RawDescriptor: file_grpc_news_proto_msg_news_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_news_proto_msg_news_view_proto_goTypes,
		DependencyIndexes: file_grpc_news_proto_msg_news_view_proto_depIdxs,
		MessageInfos:      file_grpc_news_proto_msg_news_view_proto_msgTypes,
	}.Build()
	File_grpc_news_proto_msg_news_view_proto = out.File
	file_grpc_news_proto_msg_news_view_proto_rawDesc = nil
	file_grpc_news_proto_msg_news_view_proto_goTypes = nil
	file_grpc_news_proto_msg_news_view_proto_depIdxs = nil
}
