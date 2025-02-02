// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: event/emspb/event_type.proto

package emspb

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

type EventType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventTypeName string `protobuf:"bytes,2,opt,name=eventTypeName,proto3" json:"eventTypeName,omitempty"`
}

func (x *EventType) Reset() {
	*x = EventType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_emspb_event_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventType) ProtoMessage() {}

func (x *EventType) ProtoReflect() protoreflect.Message {
	mi := &file_event_emspb_event_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventType.ProtoReflect.Descriptor instead.
func (*EventType) Descriptor() ([]byte, []int) {
	return file_event_emspb_event_type_proto_rawDescGZIP(), []int{0}
}

func (x *EventType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventType) GetEventTypeName() string {
	if x != nil {
		return x.EventTypeName
	}
	return ""
}

type CreateEventTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType *EventType `protobuf:"bytes,1,opt,name=eventType,proto3" json:"eventType,omitempty"`
}

func (x *CreateEventTypeRequest) Reset() {
	*x = CreateEventTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_emspb_event_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventTypeRequest) ProtoMessage() {}

func (x *CreateEventTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_emspb_event_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateEventTypeRequest) Descriptor() ([]byte, []int) {
	return file_event_emspb_event_type_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventTypeRequest) GetEventType() *EventType {
	if x != nil {
		return x.EventType
	}
	return nil
}

type CreateEventTypeRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType *EventType `protobuf:"bytes,1,opt,name=eventType,proto3" json:"eventType,omitempty"`
}

func (x *CreateEventTypeRespone) Reset() {
	*x = CreateEventTypeRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_emspb_event_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventTypeRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventTypeRespone) ProtoMessage() {}

func (x *CreateEventTypeRespone) ProtoReflect() protoreflect.Message {
	mi := &file_event_emspb_event_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventTypeRespone.ProtoReflect.Descriptor instead.
func (*CreateEventTypeRespone) Descriptor() ([]byte, []int) {
	return file_event_emspb_event_type_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEventTypeRespone) GetEventType() *EventType {
	if x != nil {
		return x.EventType
	}
	return nil
}

type ReadEventTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType *EventType `protobuf:"bytes,1,opt,name=eventType,proto3" json:"eventType,omitempty"`
}

func (x *ReadEventTypeRequest) Reset() {
	*x = ReadEventTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_emspb_event_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEventTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEventTypeRequest) ProtoMessage() {}

func (x *ReadEventTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_emspb_event_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEventTypeRequest.ProtoReflect.Descriptor instead.
func (*ReadEventTypeRequest) Descriptor() ([]byte, []int) {
	return file_event_emspb_event_type_proto_rawDescGZIP(), []int{3}
}

func (x *ReadEventTypeRequest) GetEventType() *EventType {
	if x != nil {
		return x.EventType
	}
	return nil
}

type ReadEventTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType *EventType `protobuf:"bytes,1,opt,name=eventType,proto3" json:"eventType,omitempty"`
}

func (x *ReadEventTypeResponse) Reset() {
	*x = ReadEventTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_emspb_event_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadEventTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadEventTypeResponse) ProtoMessage() {}

func (x *ReadEventTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_emspb_event_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadEventTypeResponse.ProtoReflect.Descriptor instead.
func (*ReadEventTypeResponse) Descriptor() ([]byte, []int) {
	return file_event_emspb_event_type_proto_rawDescGZIP(), []int{4}
}

func (x *ReadEventTypeResponse) GetEventType() *EventType {
	if x != nil {
		return x.EventType
	}
	return nil
}

var File_event_emspb_event_type_proto protoreflect.FileDescriptor

var file_event_emspb_event_type_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6d, 0x73, 0x70, 0x62, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x2e, 0x0a, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x14,
	0x52, 0x65, 0x61, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x32, 0xaf, 0x01,
	0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6d, 0x73, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_emspb_event_type_proto_rawDescOnce sync.Once
	file_event_emspb_event_type_proto_rawDescData = file_event_emspb_event_type_proto_rawDesc
)

func file_event_emspb_event_type_proto_rawDescGZIP() []byte {
	file_event_emspb_event_type_proto_rawDescOnce.Do(func() {
		file_event_emspb_event_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_emspb_event_type_proto_rawDescData)
	})
	return file_event_emspb_event_type_proto_rawDescData
}

var file_event_emspb_event_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_event_emspb_event_type_proto_goTypes = []interface{}{
	(*EventType)(nil),              // 0: event.EventType
	(*CreateEventTypeRequest)(nil), // 1: event.CreateEventTypeRequest
	(*CreateEventTypeRespone)(nil), // 2: event.CreateEventTypeRespone
	(*ReadEventTypeRequest)(nil),   // 3: event.ReadEventTypeRequest
	(*ReadEventTypeResponse)(nil),  // 4: event.ReadEventTypeResponse
}
var file_event_emspb_event_type_proto_depIdxs = []int32{
	0, // 0: event.CreateEventTypeRequest.eventType:type_name -> event.EventType
	0, // 1: event.CreateEventTypeRespone.eventType:type_name -> event.EventType
	0, // 2: event.ReadEventTypeRequest.eventType:type_name -> event.EventType
	0, // 3: event.ReadEventTypeResponse.eventType:type_name -> event.EventType
	1, // 4: event.EventTypeService.CreateEventType:input_type -> event.CreateEventTypeRequest
	3, // 5: event.EventTypeService.ReadEventType:input_type -> event.ReadEventTypeRequest
	2, // 6: event.EventTypeService.CreateEventType:output_type -> event.CreateEventTypeRespone
	4, // 7: event.EventTypeService.ReadEventType:output_type -> event.ReadEventTypeResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_event_emspb_event_type_proto_init() }
func file_event_emspb_event_type_proto_init() {
	if File_event_emspb_event_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_emspb_event_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventType); i {
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
		file_event_emspb_event_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventTypeRequest); i {
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
		file_event_emspb_event_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventTypeRespone); i {
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
		file_event_emspb_event_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEventTypeRequest); i {
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
		file_event_emspb_event_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadEventTypeResponse); i {
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
			RawDescriptor: file_event_emspb_event_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_emspb_event_type_proto_goTypes,
		DependencyIndexes: file_event_emspb_event_type_proto_depIdxs,
		MessageInfos:      file_event_emspb_event_type_proto_msgTypes,
	}.Build()
	File_event_emspb_event_type_proto = out.File
	file_event_emspb_event_type_proto_rawDesc = nil
	file_event_emspb_event_type_proto_goTypes = nil
	file_event_emspb_event_type_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventTypeServiceClient is the client API for EventTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventTypeServiceClient interface {
	CreateEventType(ctx context.Context, in *CreateEventTypeRequest, opts ...grpc.CallOption) (*CreateEventTypeRespone, error)
	ReadEventType(ctx context.Context, in *ReadEventTypeRequest, opts ...grpc.CallOption) (*ReadEventTypeResponse, error)
}

type eventTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventTypeServiceClient(cc grpc.ClientConnInterface) EventTypeServiceClient {
	return &eventTypeServiceClient{cc}
}

func (c *eventTypeServiceClient) CreateEventType(ctx context.Context, in *CreateEventTypeRequest, opts ...grpc.CallOption) (*CreateEventTypeRespone, error) {
	out := new(CreateEventTypeRespone)
	err := c.cc.Invoke(ctx, "/event.EventTypeService/CreateEventType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventTypeServiceClient) ReadEventType(ctx context.Context, in *ReadEventTypeRequest, opts ...grpc.CallOption) (*ReadEventTypeResponse, error) {
	out := new(ReadEventTypeResponse)
	err := c.cc.Invoke(ctx, "/event.EventTypeService/ReadEventType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventTypeServiceServer is the server API for EventTypeService service.
type EventTypeServiceServer interface {
	CreateEventType(context.Context, *CreateEventTypeRequest) (*CreateEventTypeRespone, error)
	ReadEventType(context.Context, *ReadEventTypeRequest) (*ReadEventTypeResponse, error)
}

// UnimplementedEventTypeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventTypeServiceServer struct {
}

func (*UnimplementedEventTypeServiceServer) CreateEventType(context.Context, *CreateEventTypeRequest) (*CreateEventTypeRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEventType not implemented")
}
func (*UnimplementedEventTypeServiceServer) ReadEventType(context.Context, *ReadEventTypeRequest) (*ReadEventTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEventType not implemented")
}

func RegisterEventTypeServiceServer(s *grpc.Server, srv EventTypeServiceServer) {
	s.RegisterService(&_EventTypeService_serviceDesc, srv)
}

func _EventTypeService_CreateEventType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTypeServiceServer).CreateEventType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventTypeService/CreateEventType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTypeServiceServer).CreateEventType(ctx, req.(*CreateEventTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventTypeService_ReadEventType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadEventTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTypeServiceServer).ReadEventType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventTypeService/ReadEventType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTypeServiceServer).ReadEventType(ctx, req.(*ReadEventTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventTypeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event.EventTypeService",
	HandlerType: (*EventTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEventType",
			Handler:    _EventTypeService_CreateEventType_Handler,
		},
		{
			MethodName: "ReadEventType",
			Handler:    _EventTypeService_ReadEventType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event/emspb/event_type.proto",
}
