// Starter Proto file, feel free to change name of the file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: helloworld_2.proto

package fifth_proto

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

// The request message containing the user's name.
type GetNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileContainer string `protobuf:"bytes,1,opt,name=file_container,json=fileContainer,proto3" json:"file_container,omitempty"`
	CodeContainer string `protobuf:"bytes,2,opt,name=code_container,json=codeContainer,proto3" json:"code_container,omitempty"`
	Nmap          string `protobuf:"bytes,3,opt,name=nmap,proto3" json:"nmap,omitempty"`
	Nreduce       string `protobuf:"bytes,4,opt,name=nreduce,proto3" json:"nreduce,omitempty"`
}

func (x *GetNames) Reset() {
	*x = GetNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNames) ProtoMessage() {}

func (x *GetNames) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNames.ProtoReflect.Descriptor instead.
func (*GetNames) Descriptor() ([]byte, []int) {
	return file_helloworld_2_proto_rawDescGZIP(), []int{0}
}

func (x *GetNames) GetFileContainer() string {
	if x != nil {
		return x.FileContainer
	}
	return ""
}

func (x *GetNames) GetCodeContainer() string {
	if x != nil {
		return x.CodeContainer
	}
	return ""
}

func (x *GetNames) GetNmap() string {
	if x != nil {
		return x.Nmap
	}
	return ""
}

func (x *GetNames) GetNreduce() string {
	if x != nil {
		return x.Nreduce
	}
	return ""
}

// The request message containing the user's name.
type PrintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *PrintRequest) Reset() {
	*x = PrintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintRequest) ProtoMessage() {}

func (x *PrintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintRequest.ProtoReflect.Descriptor instead.
func (*PrintRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_2_proto_rawDescGZIP(), []int{1}
}

func (x *PrintRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

// The response message containing the greetings
type PrintReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PrintReply) Reset() {
	*x = PrintReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintReply) ProtoMessage() {}

func (x *PrintReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintReply.ProtoReflect.Descriptor instead.
func (*PrintReply) Descriptor() ([]byte, []int) {
	return file_helloworld_2_proto_rawDescGZIP(), []int{2}
}

func (x *PrintReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The request message containing the user's name.
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapFinished    []string          `protobuf:"bytes,1,rep,name=map_finished,json=mapFinished,proto3" json:"map_finished,omitempty"`
	ReduceFinished map[string]string `protobuf:"bytes,2,rep,name=reduce_finished,json=reduceFinished,proto3" json:"reduce_finished,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WorkerStatus   map[string]string `protobuf:"bytes,3,rep,name=worker_status,json=workerStatus,proto3" json:"worker_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_2_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetMapFinished() []string {
	if x != nil {
		return x.MapFinished
	}
	return nil
}

func (x *GetRequest) GetReduceFinished() map[string]string {
	if x != nil {
		return x.ReduceFinished
	}
	return nil
}

func (x *GetRequest) GetWorkerStatus() map[string]string {
	if x != nil {
		return x.WorkerStatus
	}
	return nil
}

// The response message containing the greetings
type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_helloworld_2_proto_rawDescGZIP(), []int{4}
}

func (x *GetReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_helloworld_2_proto protoreflect.FileDescriptor

var file_helloworld_2_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x32, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6d, 0x61, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x72, 0x65,
	0x64, 0x75, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x72, 0x65, 0x64,
	0x75, 0x63, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x26, 0x0a, 0x0a, 0x50, 0x72, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xcb, 0x02, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x70, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x0f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x5f, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x64, 0x75, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0e, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x12, 0x47, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x41, 0x0a, 0x13, 0x52,
	0x65, 0x64, 0x75, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f,
	0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x24, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xe4, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x12, 0x35, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x47, 0x61, 0x74, 0x65, 0x63, 0x68,
	0x12, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x68, 0x68, 0x32, 0x31,
	0x75, 0x2f, 0x66, 0x69, 0x66, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_2_proto_rawDescOnce sync.Once
	file_helloworld_2_proto_rawDescData = file_helloworld_2_proto_rawDesc
)

func file_helloworld_2_proto_rawDescGZIP() []byte {
	file_helloworld_2_proto_rawDescOnce.Do(func() {
		file_helloworld_2_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_2_proto_rawDescData)
	})
	return file_helloworld_2_proto_rawDescData
}

var file_helloworld_2_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_helloworld_2_proto_goTypes = []interface{}{
	(*GetNames)(nil),     // 0: main.GetNames
	(*PrintRequest)(nil), // 1: main.PrintRequest
	(*PrintReply)(nil),   // 2: main.PrintReply
	(*GetRequest)(nil),   // 3: main.GetRequest
	(*GetReply)(nil),     // 4: main.GetReply
	nil,                  // 5: main.GetRequest.ReduceFinishedEntry
	nil,                  // 6: main.GetRequest.WorkerStatusEntry
}
var file_helloworld_2_proto_depIdxs = []int32{
	5, // 0: main.GetRequest.reduce_finished:type_name -> main.GetRequest.ReduceFinishedEntry
	6, // 1: main.GetRequest.worker_status:type_name -> main.GetRequest.WorkerStatusEntry
	1, // 2: main.Greeter.PrintGatech:input_type -> main.PrintRequest
	3, // 3: main.Greeter.UpdateMaster:input_type -> main.GetRequest
	1, // 4: main.Greeter.ReplyFromWorker:input_type -> main.PrintRequest
	0, // 5: main.Greeter.GetContainers:input_type -> main.GetNames
	2, // 6: main.Greeter.PrintGatech:output_type -> main.PrintReply
	4, // 7: main.Greeter.UpdateMaster:output_type -> main.GetReply
	2, // 8: main.Greeter.ReplyFromWorker:output_type -> main.PrintReply
	2, // 9: main.Greeter.GetContainers:output_type -> main.PrintReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_helloworld_2_proto_init() }
func file_helloworld_2_proto_init() {
	if File_helloworld_2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNames); i {
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
		file_helloworld_2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintRequest); i {
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
		file_helloworld_2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintReply); i {
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
		file_helloworld_2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_helloworld_2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
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
			RawDescriptor: file_helloworld_2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_2_proto_goTypes,
		DependencyIndexes: file_helloworld_2_proto_depIdxs,
		MessageInfos:      file_helloworld_2_proto_msgTypes,
	}.Build()
	File_helloworld_2_proto = out.File
	file_helloworld_2_proto_rawDesc = nil
	file_helloworld_2_proto_goTypes = nil
	file_helloworld_2_proto_depIdxs = nil
}
