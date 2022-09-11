// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: agent/agent.proto

package agent

import (
	common "github.com/dollarkillerx/go-proxy-manager/proto/common"
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

type AgentInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentID string `protobuf:"bytes,1,opt,name=agentID,proto3" json:"agentID,omitempty"`
}

func (x *AgentInfoReq) Reset() {
	*x = AgentInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfoReq) ProtoMessage() {}

func (x *AgentInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfoReq.ProtoReflect.Descriptor instead.
func (*AgentInfoReq) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{0}
}

func (x *AgentInfoReq) GetAgentID() string {
	if x != nil {
		return x.AgentID
	}
	return ""
}

type AgentInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip                   string  `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	TasksNumber          int32   `protobuf:"varint,2,opt,name=tasksNumber,proto3" json:"tasksNumber,omitempty"`                   // 任务总数
	TasksAbnormalNumber  int32   `protobuf:"varint,3,opt,name=tasksAbnormalNumber,proto3" json:"tasksAbnormalNumber,omitempty"`   // 异常任务数
	TasksSuspendedNumber int32   `protobuf:"varint,4,opt,name=tasksSuspendedNumber,proto3" json:"tasksSuspendedNumber,omitempty"` // 暂停任务数
	Load                 float32 `protobuf:"fixed32,5,opt,name=load,proto3" json:"load,omitempty"`
	CpuInfo              string  `protobuf:"bytes,6,opt,name=cpuInfo,proto3" json:"cpuInfo,omitempty"`
	MemInfo              string  `protobuf:"bytes,7,opt,name=memInfo,proto3" json:"memInfo,omitempty"`
}

func (x *AgentInfoResp) Reset() {
	*x = AgentInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfoResp) ProtoMessage() {}

func (x *AgentInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfoResp.ProtoReflect.Descriptor instead.
func (*AgentInfoResp) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{1}
}

func (x *AgentInfoResp) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *AgentInfoResp) GetTasksNumber() int32 {
	if x != nil {
		return x.TasksNumber
	}
	return 0
}

func (x *AgentInfoResp) GetTasksAbnormalNumber() int32 {
	if x != nil {
		return x.TasksAbnormalNumber
	}
	return 0
}

func (x *AgentInfoResp) GetTasksSuspendedNumber() int32 {
	if x != nil {
		return x.TasksSuspendedNumber
	}
	return 0
}

func (x *AgentInfoResp) GetLoad() float32 {
	if x != nil {
		return x.Load
	}
	return 0
}

func (x *AgentInfoResp) GetCpuInfo() string {
	if x != nil {
		return x.CpuInfo
	}
	return ""
}

func (x *AgentInfoResp) GetMemInfo() string {
	if x != nil {
		return x.MemInfo
	}
	return ""
}

type AddTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain       string               `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	EnableSsl    bool                 `protobuf:"varint,2,opt,name=enableSsl,proto3" json:"enableSsl,omitempty"`
	EnableWaf    bool                 `protobuf:"varint,3,opt,name=enableWaf,proto3" json:"enableWaf,omitempty"`
	ProxyType    common.ProxyType     `protobuf:"varint,4,opt,name=proxyType,proto3,enum=common.ProxyType" json:"proxyType,omitempty"`
	ReverseProxy *common.ReverseProxy `protobuf:"bytes,5,opt,name=reverseProxy,proto3" json:"reverseProxy,omitempty"`
	StreamProxy  *common.StreamProxy  `protobuf:"bytes,6,opt,name=streamProxy,proto3" json:"streamProxy,omitempty"`
	Certificate  *common.Certificate  `protobuf:"bytes,7,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *AddTaskReq) Reset() {
	*x = AddTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskReq) ProtoMessage() {}

func (x *AddTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskReq.ProtoReflect.Descriptor instead.
func (*AddTaskReq) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{2}
}

func (x *AddTaskReq) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *AddTaskReq) GetEnableSsl() bool {
	if x != nil {
		return x.EnableSsl
	}
	return false
}

func (x *AddTaskReq) GetEnableWaf() bool {
	if x != nil {
		return x.EnableWaf
	}
	return false
}

func (x *AddTaskReq) GetProxyType() common.ProxyType {
	if x != nil {
		return x.ProxyType
	}
	return common.ProxyType(0)
}

func (x *AddTaskReq) GetReverseProxy() *common.ReverseProxy {
	if x != nil {
		return x.ReverseProxy
	}
	return nil
}

func (x *AddTaskReq) GetStreamProxy() *common.StreamProxy {
	if x != nil {
		return x.StreamProxy
	}
	return nil
}

func (x *AddTaskReq) GetCertificate() *common.Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

type ModifyTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID       string               `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	EnableProxy  bool                 `protobuf:"varint,2,opt,name=enableProxy,proto3" json:"enableProxy,omitempty"`
	EnableSsl    bool                 `protobuf:"varint,3,opt,name=enableSsl,proto3" json:"enableSsl,omitempty"`
	EnableWaf    bool                 `protobuf:"varint,4,opt,name=enableWaf,proto3" json:"enableWaf,omitempty"`
	DeleteTask   bool                 `protobuf:"varint,5,opt,name=deleteTask,proto3" json:"deleteTask,omitempty"`
	ProxyType    common.ProxyType     `protobuf:"varint,6,opt,name=proxyType,proto3,enum=common.ProxyType" json:"proxyType,omitempty"`
	ReverseProxy *common.ReverseProxy `protobuf:"bytes,7,opt,name=reverseProxy,proto3" json:"reverseProxy,omitempty"`
	StreamProxy  *common.StreamProxy  `protobuf:"bytes,8,opt,name=streamProxy,proto3" json:"streamProxy,omitempty"`
	Certificate  *common.Certificate  `protobuf:"bytes,9,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *ModifyTaskReq) Reset() {
	*x = ModifyTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyTaskReq) ProtoMessage() {}

func (x *ModifyTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyTaskReq.ProtoReflect.Descriptor instead.
func (*ModifyTaskReq) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{3}
}

func (x *ModifyTaskReq) GetTaskID() string {
	if x != nil {
		return x.TaskID
	}
	return ""
}

func (x *ModifyTaskReq) GetEnableProxy() bool {
	if x != nil {
		return x.EnableProxy
	}
	return false
}

func (x *ModifyTaskReq) GetEnableSsl() bool {
	if x != nil {
		return x.EnableSsl
	}
	return false
}

func (x *ModifyTaskReq) GetEnableWaf() bool {
	if x != nil {
		return x.EnableWaf
	}
	return false
}

func (x *ModifyTaskReq) GetDeleteTask() bool {
	if x != nil {
		return x.DeleteTask
	}
	return false
}

func (x *ModifyTaskReq) GetProxyType() common.ProxyType {
	if x != nil {
		return x.ProxyType
	}
	return common.ProxyType(0)
}

func (x *ModifyTaskReq) GetReverseProxy() *common.ReverseProxy {
	if x != nil {
		return x.ReverseProxy
	}
	return nil
}

func (x *ModifyTaskReq) GetStreamProxy() *common.StreamProxy {
	if x != nil {
		return x.StreamProxy
	}
	return nil
}

func (x *ModifyTaskReq) GetCertificate() *common.Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

type ApplyCertificateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *ApplyCertificateReq) Reset() {
	*x = ApplyCertificateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyCertificateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyCertificateReq) ProtoMessage() {}

func (x *ApplyCertificateReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyCertificateReq.ProtoReflect.Descriptor instead.
func (*ApplyCertificateReq) Descriptor() ([]byte, []int) {
	return file_agent_agent_proto_rawDescGZIP(), []int{4}
}

func (x *ApplyCertificateReq) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

var File_agent_agent_proto protoreflect.FileDescriptor

var file_agent_agent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x28, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0xef, 0x01, 0x0a, 0x0d, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a,
	0x13, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x41, 0x62, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x41, 0x62, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x32, 0x0a, 0x14, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x70, 0x75, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x70, 0x75, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xb9, 0x02, 0x0a, 0x0a,
	0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x73, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x73, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x61, 0x66, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x61, 0x66, 0x12, 0x2f,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x38, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x12, 0x35, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0xfe, 0x02, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x73, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x73,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x61, 0x66, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x61, 0x66, 0x12,
	0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x2f, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x0c, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x32, 0xdf, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x31, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x11, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x31, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x14, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x6b, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x78, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2d, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_agent_proto_rawDescOnce sync.Once
	file_agent_agent_proto_rawDescData = file_agent_agent_proto_rawDesc
)

func file_agent_agent_proto_rawDescGZIP() []byte {
	file_agent_agent_proto_rawDescOnce.Do(func() {
		file_agent_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_agent_proto_rawDescData)
	})
	return file_agent_agent_proto_rawDescData
}

var file_agent_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_agent_agent_proto_goTypes = []interface{}{
	(*AgentInfoReq)(nil),        // 0: agent.AgentInfoReq
	(*AgentInfoResp)(nil),       // 1: agent.AgentInfoResp
	(*AddTaskReq)(nil),          // 2: agent.AddTaskReq
	(*ModifyTaskReq)(nil),       // 3: agent.ModifyTaskReq
	(*ApplyCertificateReq)(nil), // 4: agent.ApplyCertificateReq
	(common.ProxyType)(0),       // 5: common.ProxyType
	(*common.ReverseProxy)(nil), // 6: common.ReverseProxy
	(*common.StreamProxy)(nil),  // 7: common.StreamProxy
	(*common.Certificate)(nil),  // 8: common.Certificate
	(*common.Empty)(nil),        // 9: common.Empty
}
var file_agent_agent_proto_depIdxs = []int32{
	5,  // 0: agent.AddTaskReq.proxyType:type_name -> common.ProxyType
	6,  // 1: agent.AddTaskReq.reverseProxy:type_name -> common.ReverseProxy
	7,  // 2: agent.AddTaskReq.streamProxy:type_name -> common.StreamProxy
	8,  // 3: agent.AddTaskReq.certificate:type_name -> common.Certificate
	5,  // 4: agent.ModifyTaskReq.proxyType:type_name -> common.ProxyType
	6,  // 5: agent.ModifyTaskReq.reverseProxy:type_name -> common.ReverseProxy
	7,  // 6: agent.ModifyTaskReq.streamProxy:type_name -> common.StreamProxy
	8,  // 7: agent.ModifyTaskReq.certificate:type_name -> common.Certificate
	0,  // 8: agent.Agent.Info:input_type -> agent.AgentInfoReq
	2,  // 9: agent.Agent.AddTask:input_type -> agent.AddTaskReq
	3,  // 10: agent.Agent.ModifyTask:input_type -> agent.ModifyTaskReq
	4,  // 11: agent.Agent.ApplyCertificate:input_type -> agent.ApplyCertificateReq
	1,  // 12: agent.Agent.Info:output_type -> agent.AgentInfoResp
	9,  // 13: agent.Agent.AddTask:output_type -> common.Empty
	9,  // 14: agent.Agent.ModifyTask:output_type -> common.Empty
	8,  // 15: agent.Agent.ApplyCertificate:output_type -> common.Certificate
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_agent_agent_proto_init() }
func file_agent_agent_proto_init() {
	if File_agent_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentInfoReq); i {
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
		file_agent_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentInfoResp); i {
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
		file_agent_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskReq); i {
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
		file_agent_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyTaskReq); i {
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
		file_agent_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyCertificateReq); i {
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
			RawDescriptor: file_agent_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_agent_proto_goTypes,
		DependencyIndexes: file_agent_agent_proto_depIdxs,
		MessageInfos:      file_agent_agent_proto_msgTypes,
	}.Build()
	File_agent_agent_proto = out.File
	file_agent_agent_proto_rawDesc = nil
	file_agent_agent_proto_goTypes = nil
	file_agent_agent_proto_depIdxs = nil
}
