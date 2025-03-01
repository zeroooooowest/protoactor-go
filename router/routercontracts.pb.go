// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: routercontracts.proto

package router

import (
	actor "github.com/asynkron/protoactor-go/actor"
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

type AddRoutee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PID *actor.PID `protobuf:"bytes,1,opt,name=PID,proto3" json:"PID,omitempty"`
}

func (x *AddRoutee) Reset() {
	*x = AddRoutee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routercontracts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoutee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoutee) ProtoMessage() {}

func (x *AddRoutee) ProtoReflect() protoreflect.Message {
	mi := &file_routercontracts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoutee.ProtoReflect.Descriptor instead.
func (*AddRoutee) Descriptor() ([]byte, []int) {
	return file_routercontracts_proto_rawDescGZIP(), []int{0}
}

func (x *AddRoutee) GetPID() *actor.PID {
	if x != nil {
		return x.PID
	}
	return nil
}

type RemoveRoutee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PID *actor.PID `protobuf:"bytes,1,opt,name=PID,proto3" json:"PID,omitempty"`
}

func (x *RemoveRoutee) Reset() {
	*x = RemoveRoutee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routercontracts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRoutee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoutee) ProtoMessage() {}

func (x *RemoveRoutee) ProtoReflect() protoreflect.Message {
	mi := &file_routercontracts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoutee.ProtoReflect.Descriptor instead.
func (*RemoveRoutee) Descriptor() ([]byte, []int) {
	return file_routercontracts_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveRoutee) GetPID() *actor.PID {
	if x != nil {
		return x.PID
	}
	return nil
}

type AdjustPoolSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Change int32 `protobuf:"varint,1,opt,name=change,proto3" json:"change,omitempty"`
}

func (x *AdjustPoolSize) Reset() {
	*x = AdjustPoolSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routercontracts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdjustPoolSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustPoolSize) ProtoMessage() {}

func (x *AdjustPoolSize) ProtoReflect() protoreflect.Message {
	mi := &file_routercontracts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjustPoolSize.ProtoReflect.Descriptor instead.
func (*AdjustPoolSize) Descriptor() ([]byte, []int) {
	return file_routercontracts_proto_rawDescGZIP(), []int{2}
}

func (x *AdjustPoolSize) GetChange() int32 {
	if x != nil {
		return x.Change
	}
	return 0
}

type GetRoutees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRoutees) Reset() {
	*x = GetRoutees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routercontracts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoutees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoutees) ProtoMessage() {}

func (x *GetRoutees) ProtoReflect() protoreflect.Message {
	mi := &file_routercontracts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoutees.ProtoReflect.Descriptor instead.
func (*GetRoutees) Descriptor() ([]byte, []int) {
	return file_routercontracts_proto_rawDescGZIP(), []int{3}
}

type Routees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PIDs []*actor.PID `protobuf:"bytes,1,rep,name=PIDs,proto3" json:"PIDs,omitempty"`
}

func (x *Routees) Reset() {
	*x = Routees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routercontracts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Routees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Routees) ProtoMessage() {}

func (x *Routees) ProtoReflect() protoreflect.Message {
	mi := &file_routercontracts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Routees.ProtoReflect.Descriptor instead.
func (*Routees) Descriptor() ([]byte, []int) {
	return file_routercontracts_proto_rawDescGZIP(), []int{4}
}

func (x *Routees) GetPIDs() []*actor.PID {
	if x != nil {
		return x.PIDs
	}
	return nil
}

var File_routercontracts_proto protoreflect.FileDescriptor

var file_routercontracts_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x1a,
	0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x09,
	0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x50, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50,
	0x49, 0x44, 0x52, 0x03, 0x50, 0x49, 0x44, 0x22, 0x2c, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44,
	0x52, 0x03, 0x50, 0x49, 0x44, 0x22, 0x28, 0x0a, 0x0e, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x50,
	0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22,
	0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x65, 0x73, 0x22, 0x29, 0x0a,
	0x07, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x50, 0x49, 0x44, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50,
	0x49, 0x44, 0x52, 0x04, 0x50, 0x49, 0x44, 0x73, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x6b, 0x72, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routercontracts_proto_rawDescOnce sync.Once
	file_routercontracts_proto_rawDescData = file_routercontracts_proto_rawDesc
)

func file_routercontracts_proto_rawDescGZIP() []byte {
	file_routercontracts_proto_rawDescOnce.Do(func() {
		file_routercontracts_proto_rawDescData = protoimpl.X.CompressGZIP(file_routercontracts_proto_rawDescData)
	})
	return file_routercontracts_proto_rawDescData
}

var file_routercontracts_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_routercontracts_proto_goTypes = []interface{}{
	(*AddRoutee)(nil),      // 0: router.AddRoutee
	(*RemoveRoutee)(nil),   // 1: router.RemoveRoutee
	(*AdjustPoolSize)(nil), // 2: router.AdjustPoolSize
	(*GetRoutees)(nil),     // 3: router.GetRoutees
	(*Routees)(nil),        // 4: router.Routees
	(*actor.PID)(nil),      // 5: actor.PID
}
var file_routercontracts_proto_depIdxs = []int32{
	5, // 0: router.AddRoutee.PID:type_name -> actor.PID
	5, // 1: router.RemoveRoutee.PID:type_name -> actor.PID
	5, // 2: router.Routees.PIDs:type_name -> actor.PID
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_routercontracts_proto_init() }
func file_routercontracts_proto_init() {
	if File_routercontracts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routercontracts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoutee); i {
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
		file_routercontracts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRoutee); i {
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
		file_routercontracts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdjustPoolSize); i {
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
		file_routercontracts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoutees); i {
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
		file_routercontracts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Routees); i {
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
			RawDescriptor: file_routercontracts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_routercontracts_proto_goTypes,
		DependencyIndexes: file_routercontracts_proto_depIdxs,
		MessageInfos:      file_routercontracts_proto_msgTypes,
	}.Build()
	File_routercontracts_proto = out.File
	file_routercontracts_proto_rawDesc = nil
	file_routercontracts_proto_goTypes = nil
	file_routercontracts_proto_depIdxs = nil
}
