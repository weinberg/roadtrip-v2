// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/gameServer/grpc/game.proto

package game_server_grpc

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

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location   *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Mph        float32   `protobuf:"fixed32,3,opt,name=mph,proto3" json:"mph,omitempty"`
	Odometer   float32   `protobuf:"fixed32,4,opt,name=odometer,proto3" json:"odometer,omitempty"`
	Tripometer float32   `protobuf:"fixed32,5,opt,name=tripometer,proto3" json:"tripometer,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gameServer_grpc_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gameServer_grpc_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_internal_gameServer_grpc_game_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Car) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Car) GetMph() float32 {
	if x != nil {
		return x.Mph
	}
	return 0
}

func (x *Car) GetOdometer() float32 {
	if x != nil {
		return x.Odometer
	}
	return 0
}

func (x *Car) GetTripometer() float32 {
	if x != nil {
		return x.Tripometer
	}
	return 0
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId string  `protobuf:"bytes,1,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	Index   int32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Miles   float32 `protobuf:"fixed32,3,opt,name=miles,proto3" json:"miles,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gameServer_grpc_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gameServer_grpc_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_internal_gameServer_grpc_game_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

func (x *Location) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Location) GetMiles() float32 {
	if x != nil {
		return x.Miles
	}
	return 0
}

type Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mph        float32 `protobuf:"fixed32,1,opt,name=mph,proto3" json:"mph,omitempty"`
	Odometer   float32 `protobuf:"fixed32,2,opt,name=odometer,proto3" json:"odometer,omitempty"`
	Tripometer float32 `protobuf:"fixed32,3,opt,name=tripometer,proto3" json:"tripometer,omitempty"`
	Index      int32   `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	Miles      float32 `protobuf:"fixed32,5,opt,name=miles,proto3" json:"miles,omitempty"`
}

func (x *Update) Reset() {
	*x = Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gameServer_grpc_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update) ProtoMessage() {}

func (x *Update) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gameServer_grpc_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update.ProtoReflect.Descriptor instead.
func (*Update) Descriptor() ([]byte, []int) {
	return file_internal_gameServer_grpc_game_proto_rawDescGZIP(), []int{2}
}

func (x *Update) GetMph() float32 {
	if x != nil {
		return x.Mph
	}
	return 0
}

func (x *Update) GetOdometer() float32 {
	if x != nil {
		return x.Odometer
	}
	return 0
}

func (x *Update) GetTripometer() float32 {
	if x != nil {
		return x.Tripometer
	}
	return 0
}

func (x *Update) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Update) GetMiles() float32 {
	if x != nil {
		return x.Miles
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gameServer_grpc_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gameServer_grpc_game_proto_msgTypes[3]
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
	return file_internal_gameServer_grpc_game_proto_rawDescGZIP(), []int{3}
}

type UpsertCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *UpsertCarRequest) Reset() {
	*x = UpsertCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gameServer_grpc_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCarRequest) ProtoMessage() {}

func (x *UpsertCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gameServer_grpc_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCarRequest.ProtoReflect.Descriptor instead.
func (*UpsertCarRequest) Descriptor() ([]byte, []int) {
	return file_internal_gameServer_grpc_game_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertCarRequest) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type GetUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *GetUpdateRequest) Reset() {
	*x = GetUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gameServer_grpc_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdateRequest) ProtoMessage() {}

func (x *GetUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gameServer_grpc_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdateRequest.ProtoReflect.Descriptor instead.
func (*GetUpdateRequest) Descriptor() ([]byte, []int) {
	return file_internal_gameServer_grpc_game_proto_rawDescGZIP(), []int{5}
}

func (x *GetUpdateRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

var File_internal_gameServer_grpc_game_proto protoreflect.FileDescriptor

var file_internal_gameServer_grpc_game_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x22,
	0x93, 0x01, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x6f, 0x61, 0x64,
	0x74, 0x72, 0x69, 0x70, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x70, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x70, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x64, 0x6f,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6f, 0x64, 0x6f,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x6f, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x6f,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x6d, 0x70, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x64, 0x6f, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6f, 0x64, 0x6f, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x6f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x6f, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x33, 0x0a, 0x10, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x63, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x74, 0x72,
	0x69, 0x70, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22, 0x29, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x32, 0x84, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x61, 0x64, 0x54,
	0x72, 0x69, 0x70, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x12, 0x1a, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e,
	0x72, 0x6f, 0x61, 0x64, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x6f, 0x61, 0x64,
	0x74, 0x72, 0x69, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a,
	0x25, 0x69, 0x6e, 0x73, 0x6f, 0x66, 0x61, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x61,
	0x64, 0x54, 0x72, 0x69, 0x70, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_gameServer_grpc_game_proto_rawDescOnce sync.Once
	file_internal_gameServer_grpc_game_proto_rawDescData = file_internal_gameServer_grpc_game_proto_rawDesc
)

func file_internal_gameServer_grpc_game_proto_rawDescGZIP() []byte {
	file_internal_gameServer_grpc_game_proto_rawDescOnce.Do(func() {
		file_internal_gameServer_grpc_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_gameServer_grpc_game_proto_rawDescData)
	})
	return file_internal_gameServer_grpc_game_proto_rawDescData
}

var file_internal_gameServer_grpc_game_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_gameServer_grpc_game_proto_goTypes = []interface{}{
	(*Car)(nil),              // 0: roadtrip.Car
	(*Location)(nil),         // 1: roadtrip.Location
	(*Update)(nil),           // 2: roadtrip.Update
	(*Empty)(nil),            // 3: roadtrip.Empty
	(*UpsertCarRequest)(nil), // 4: roadtrip.UpsertCarRequest
	(*GetUpdateRequest)(nil), // 5: roadtrip.GetUpdateRequest
}
var file_internal_gameServer_grpc_game_proto_depIdxs = []int32{
	1, // 0: roadtrip.Car.location:type_name -> roadtrip.Location
	0, // 1: roadtrip.UpsertCarRequest.car:type_name -> roadtrip.Car
	4, // 2: roadtrip.RoadTripGame.Upsert:input_type -> roadtrip.UpsertCarRequest
	5, // 3: roadtrip.RoadTripGame.GetUpdate:input_type -> roadtrip.GetUpdateRequest
	3, // 4: roadtrip.RoadTripGame.Upsert:output_type -> roadtrip.Empty
	2, // 5: roadtrip.RoadTripGame.GetUpdate:output_type -> roadtrip.Update
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_gameServer_grpc_game_proto_init() }
func file_internal_gameServer_grpc_game_proto_init() {
	if File_internal_gameServer_grpc_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_gameServer_grpc_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_internal_gameServer_grpc_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_internal_gameServer_grpc_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Update); i {
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
		file_internal_gameServer_grpc_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_gameServer_grpc_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertCarRequest); i {
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
		file_internal_gameServer_grpc_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpdateRequest); i {
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
			RawDescriptor: file_internal_gameServer_grpc_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_gameServer_grpc_game_proto_goTypes,
		DependencyIndexes: file_internal_gameServer_grpc_game_proto_depIdxs,
		MessageInfos:      file_internal_gameServer_grpc_game_proto_msgTypes,
	}.Build()
	File_internal_gameServer_grpc_game_proto = out.File
	file_internal_gameServer_grpc_game_proto_rawDesc = nil
	file_internal_gameServer_grpc_game_proto_goTypes = nil
	file_internal_gameServer_grpc_game_proto_depIdxs = nil
}
