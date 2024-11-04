// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: shortener/shortener.proto

package shortenerv1

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

type UrlShortenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	UserId uint64 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UrlShortenRequest) Reset() {
	*x = UrlShortenRequest{}
	mi := &file_shortener_shortener_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UrlShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlShortenRequest) ProtoMessage() {}

func (x *UrlShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_shortener_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlShortenRequest.ProtoReflect.Descriptor instead.
func (*UrlShortenRequest) Descriptor() ([]byte, []int) {
	return file_shortener_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *UrlShortenRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UrlShortenRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UrlShortenRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UrlShortenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UrlShortenResponse) Reset() {
	*x = UrlShortenResponse{}
	mi := &file_shortener_shortener_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UrlShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlShortenResponse) ProtoMessage() {}

func (x *UrlShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_shortener_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlShortenResponse.ProtoReflect.Descriptor instead.
func (*UrlShortenResponse) Descriptor() ([]byte, []int) {
	return file_shortener_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *UrlShortenResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UrlUnshortenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string                  `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Headers map[string]*StringArray `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ip      string                  `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *UrlUnshortenRequest) Reset() {
	*x = UrlUnshortenRequest{}
	mi := &file_shortener_shortener_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UrlUnshortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlUnshortenRequest) ProtoMessage() {}

func (x *UrlUnshortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_shortener_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlUnshortenRequest.ProtoReflect.Descriptor instead.
func (*UrlUnshortenRequest) Descriptor() ([]byte, []int) {
	return file_shortener_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *UrlUnshortenRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UrlUnshortenRequest) GetHeaders() map[string]*StringArray {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *UrlUnshortenRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type UrlUnshortenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UrlUnshortenResponse) Reset() {
	*x = UrlUnshortenResponse{}
	mi := &file_shortener_shortener_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UrlUnshortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlUnshortenResponse) ProtoMessage() {}

func (x *UrlUnshortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_shortener_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlUnshortenResponse.ProtoReflect.Descriptor instead.
func (*UrlUnshortenResponse) Descriptor() ([]byte, []int) {
	return file_shortener_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *UrlUnshortenResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type StringArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *StringArray) Reset() {
	*x = StringArray{}
	mi := &file_shortener_shortener_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StringArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringArray) ProtoMessage() {}

func (x *StringArray) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_shortener_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringArray.ProtoReflect.Descriptor instead.
func (*StringArray) Descriptor() ([]byte, []int) {
	return file_shortener_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *StringArray) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_shortener_shortener_proto protoreflect.FileDescriptor

var file_shortener_shortener_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x22, 0x53, 0x0a, 0x11, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x12, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xc8,
	0x01, 0x0a, 0x13, 0x55, 0x72, 0x6c, 0x55, 0x6e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x40, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x55, 0x72, 0x6c, 0x55, 0x6e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x1a, 0x4d, 0x0a, 0x0c, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x28, 0x0a, 0x14, 0x55, 0x72, 0x6c,
	0x55, 0x6e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0x8d, 0x01, 0x0a, 0x09, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x72, 0x6c, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x55, 0x6e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x72, 0x6c, 0x55, 0x6e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x72, 0x6c, 0x55, 0x6e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x65, 0x77,
	0x6d, 0x77, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x3b,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_shortener_shortener_proto_rawDescOnce sync.Once
	file_shortener_shortener_proto_rawDescData = file_shortener_shortener_proto_rawDesc
)

func file_shortener_shortener_proto_rawDescGZIP() []byte {
	file_shortener_shortener_proto_rawDescOnce.Do(func() {
		file_shortener_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortener_shortener_proto_rawDescData)
	})
	return file_shortener_shortener_proto_rawDescData
}

var file_shortener_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_shortener_shortener_proto_goTypes = []any{
	(*UrlShortenRequest)(nil),    // 0: auth.UrlShortenRequest
	(*UrlShortenResponse)(nil),   // 1: auth.UrlShortenResponse
	(*UrlUnshortenRequest)(nil),  // 2: auth.UrlUnshortenRequest
	(*UrlUnshortenResponse)(nil), // 3: auth.UrlUnshortenResponse
	(*StringArray)(nil),          // 4: auth.StringArray
	nil,                          // 5: auth.UrlUnshortenRequest.HeadersEntry
}
var file_shortener_shortener_proto_depIdxs = []int32{
	5, // 0: auth.UrlUnshortenRequest.headers:type_name -> auth.UrlUnshortenRequest.HeadersEntry
	4, // 1: auth.UrlUnshortenRequest.HeadersEntry.value:type_name -> auth.StringArray
	0, // 2: auth.Shortener.Shorten:input_type -> auth.UrlShortenRequest
	2, // 3: auth.Shortener.Unshorten:input_type -> auth.UrlUnshortenRequest
	1, // 4: auth.Shortener.Shorten:output_type -> auth.UrlShortenResponse
	3, // 5: auth.Shortener.Unshorten:output_type -> auth.UrlUnshortenResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shortener_shortener_proto_init() }
func file_shortener_shortener_proto_init() {
	if File_shortener_shortener_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shortener_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortener_shortener_proto_goTypes,
		DependencyIndexes: file_shortener_shortener_proto_depIdxs,
		MessageInfos:      file_shortener_shortener_proto_msgTypes,
	}.Build()
	File_shortener_shortener_proto = out.File
	file_shortener_shortener_proto_rawDesc = nil
	file_shortener_shortener_proto_goTypes = nil
	file_shortener_shortener_proto_depIdxs = nil
}
