// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: common/common.proto

package common

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

type ProxyType int32

const (
	ProxyType_ReverseProxyEnum  ProxyType = 0 // 反向代理
	ProxyType_RedirectProxyEnum ProxyType = 1 // 重定向
	ProxyType_StreamProxyEnum   ProxyType = 2 // 流代理
	ProxyType_R404ProxyEnum     ProxyType = 4 // 404
)

// Enum value maps for ProxyType.
var (
	ProxyType_name = map[int32]string{
		0: "ReverseProxyEnum",
		1: "RedirectProxyEnum",
		2: "StreamProxyEnum",
		4: "R404ProxyEnum",
	}
	ProxyType_value = map[string]int32{
		"ReverseProxyEnum":  0,
		"RedirectProxyEnum": 1,
		"StreamProxyEnum":   2,
		"R404ProxyEnum":     4,
	}
)

func (x ProxyType) Enum() *ProxyType {
	p := new(ProxyType)
	*p = x
	return p
}

func (x ProxyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProxyType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[0].Descriptor()
}

func (ProxyType) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[0]
}

func (x ProxyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyType.Descriptor instead.
func (ProxyType) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

type RedirectStateType int32

const (
	RedirectStateType_MultipleChoices300   RedirectStateType = 0
	RedirectStateType_MovedPermanently301  RedirectStateType = 1
	RedirectStateType_Found302             RedirectStateType = 2
	RedirectStateType_SeeOther303          RedirectStateType = 3
	RedirectStateType_TemporaryRedirect307 RedirectStateType = 4
	RedirectStateType_PermanentRedirect308 RedirectStateType = 5
)

// Enum value maps for RedirectStateType.
var (
	RedirectStateType_name = map[int32]string{
		0: "MultipleChoices300",
		1: "MovedPermanently301",
		2: "Found302",
		3: "SeeOther303",
		4: "TemporaryRedirect307",
		5: "PermanentRedirect308",
	}
	RedirectStateType_value = map[string]int32{
		"MultipleChoices300":   0,
		"MovedPermanently301":  1,
		"Found302":             2,
		"SeeOther303":          3,
		"TemporaryRedirect307": 4,
		"PermanentRedirect308": 5,
	}
)

func (x RedirectStateType) Enum() *RedirectStateType {
	p := new(RedirectStateType)
	*p = x
	return p
}

func (x RedirectStateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RedirectStateType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[1].Descriptor()
}

func (RedirectStateType) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[1]
}

func (x RedirectStateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RedirectStateType.Descriptor instead.
func (RedirectStateType) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey        string `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`               // 公钥
	PrivateKey       string `protobuf:"bytes,2,opt,name=privateKey,proto3" json:"privateKey,omitempty"`             // 私钥
	Expiration       int64  `protobuf:"varint,3,opt,name=expiration,proto3" json:"expiration,omitempty"`            // 过期时间
	Domain           string `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`                     // 域名
	CertificateBrand string `protobuf:"bytes,5,opt,name=certificateBrand,proto3" json:"certificateBrand,omitempty"` // 证书品牌
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *Certificate) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Certificate) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *Certificate) GetExpiration() int64 {
	if x != nil {
		return x.Expiration
	}
	return 0
}

func (x *Certificate) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Certificate) GetCertificateBrand() string {
	if x != nil {
		return x.CertificateBrand
	}
	return ""
}

type ReverseProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target            string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	CacheAssets       bool   `protobuf:"varint,2,opt,name=cacheAssets,proto3" json:"cacheAssets,omitempty"`
	WebsocketsSupport bool   `protobuf:"varint,3,opt,name=websocketsSupport,proto3" json:"websocketsSupport,omitempty"`
}

func (x *ReverseProxy) Reset() {
	*x = ReverseProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseProxy) ProtoMessage() {}

func (x *ReverseProxy) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseProxy.ProtoReflect.Descriptor instead.
func (*ReverseProxy) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *ReverseProxy) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *ReverseProxy) GetCacheAssets() bool {
	if x != nil {
		return x.CacheAssets
	}
	return false
}

func (x *ReverseProxy) GetWebsocketsSupport() bool {
	if x != nil {
		return x.WebsocketsSupport
	}
	return false
}

type RedirectProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target            string            `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	ReservedPath      bool              `protobuf:"varint,2,opt,name=reservedPath,proto3" json:"reservedPath,omitempty"` // 保留路径
	RedirectStateType RedirectStateType `protobuf:"varint,3,opt,name=redirectStateType,proto3,enum=common.RedirectStateType" json:"redirectStateType,omitempty"`
}

func (x *RedirectProxy) Reset() {
	*x = RedirectProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectProxy) ProtoMessage() {}

func (x *RedirectProxy) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectProxy.ProtoReflect.Descriptor instead.
func (*RedirectProxy) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *RedirectProxy) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *RedirectProxy) GetReservedPath() bool {
	if x != nil {
		return x.ReservedPath
	}
	return false
}

func (x *RedirectProxy) GetRedirectStateType() RedirectStateType {
	if x != nil {
		return x.RedirectStateType
	}
	return RedirectStateType_MultipleChoices300
}

type StreamProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target      string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	ExposedPort int64  `protobuf:"varint,2,opt,name=exposedPort,proto3" json:"exposedPort,omitempty"` // 暴露端口
	EnableTcp   bool   `protobuf:"varint,3,opt,name=enableTcp,proto3" json:"enableTcp,omitempty"`
	EnableUdp   bool   `protobuf:"varint,4,opt,name=enableUdp,proto3" json:"enableUdp,omitempty"`
}

func (x *StreamProxy) Reset() {
	*x = StreamProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamProxy) ProtoMessage() {}

func (x *StreamProxy) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamProxy.ProtoReflect.Descriptor instead.
func (*StreamProxy) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *StreamProxy) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *StreamProxy) GetExposedPort() int64 {
	if x != nil {
		return x.ExposedPort
	}
	return 0
}

func (x *StreamProxy) GetEnableTcp() bool {
	if x != nil {
		return x.EnableTcp
	}
	return false
}

func (x *StreamProxy) GetEnableUdp() bool {
	if x != nil {
		return x.EnableUdp
	}
	return false
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xaf, 0x01, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2a, 0x0a, 0x10,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x76, 0x0a, 0x0c, 0x52, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x77,
	0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x94, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x47,
	0x0a, 0x11, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x63, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x63, 0x70, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x64, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x64, 0x70, 0x2a, 0x60, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x45, 0x6e, 0x75, 0x6d, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d,
	0x52, 0x34, 0x30, 0x34, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x10, 0x04, 0x2a,
	0x97, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x33, 0x30, 0x30, 0x10, 0x00, 0x12, 0x17, 0x0a,
	0x13, 0x4d, 0x6f, 0x76, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x6c,
	0x79, 0x33, 0x30, 0x31, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x33,
	0x30, 0x32, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x65, 0x65, 0x4f, 0x74, 0x68, 0x65, 0x72,
	0x33, 0x30, 0x33, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x33, 0x30, 0x37, 0x10, 0x04, 0x12,
	0x18, 0x0a, 0x14, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x33, 0x30, 0x38, 0x10, 0x05, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x6b, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x78, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2d, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_common_proto_goTypes = []interface{}{
	(ProxyType)(0),         // 0: common.ProxyType
	(RedirectStateType)(0), // 1: common.RedirectStateType
	(*Empty)(nil),          // 2: common.Empty
	(*Certificate)(nil),    // 3: common.Certificate
	(*ReverseProxy)(nil),   // 4: common.ReverseProxy
	(*RedirectProxy)(nil),  // 5: common.RedirectProxy
	(*StreamProxy)(nil),    // 6: common.StreamProxy
}
var file_common_common_proto_depIdxs = []int32{
	1, // 0: common.RedirectProxy.redirectStateType:type_name -> common.RedirectStateType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
		file_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseProxy); i {
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
		file_common_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectProxy); i {
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
		file_common_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamProxy); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		EnumInfos:         file_common_common_proto_enumTypes,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
