// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: activity.proto

package services

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[0]
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
	return file_activity_proto_rawDescGZIP(), []int{0}
}

type ActivityId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ActivityId) Reset() {
	*x = ActivityId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityId) ProtoMessage() {}

func (x *ActivityId) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityId.ProtoReflect.Descriptor instead.
func (*ActivityId) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string         `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description    string         `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	ImageProfile   string         `protobuf:"bytes,3,opt,name=ImageProfile,proto3" json:"ImageProfile,omitempty"`
	Type           *ListOfStrings `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	OwnerId        string         `protobuf:"bytes,5,opt,name=OwnerId,proto3" json:"OwnerId,omitempty"`
	Location       string         `protobuf:"bytes,6,opt,name=Location,proto3" json:"Location,omitempty"`
	MaxParticipant int64          `protobuf:"varint,7,opt,name=MaxParticipant,proto3" json:"MaxParticipant,omitempty"`
	Participant    *ListOfStrings `protobuf:"bytes,8,opt,name=Participant,proto3" json:"Participant,omitempty"`
	// google.protobuf.Timestamp Date = 9;
	Duration int32  `protobuf:"varint,10,opt,name=Duration,proto3" json:"Duration,omitempty"`
	ChatId   string `protobuf:"bytes,11,opt,name=ChatId,proto3" json:"ChatId,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{2}
}

func (x *Activity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Activity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Activity) GetImageProfile() string {
	if x != nil {
		return x.ImageProfile
	}
	return ""
}

func (x *Activity) GetType() *ListOfStrings {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Activity) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Activity) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Activity) GetMaxParticipant() int64 {
	if x != nil {
		return x.MaxParticipant
	}
	return 0
}

func (x *Activity) GetParticipant() *ListOfStrings {
	if x != nil {
		return x.Participant
	}
	return nil
}

func (x *Activity) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Activity) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type ListOfStrings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strings []string `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty"`
}

func (x *ListOfStrings) Reset() {
	*x = ListOfStrings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfStrings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfStrings) ProtoMessage() {}

func (x *ListOfStrings) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfStrings.ProtoReflect.Descriptor instead.
func (*ListOfStrings) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{3}
}

func (x *ListOfStrings) GetStrings() []string {
	if x != nil {
		return x.Strings
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int64  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_activity_proto protoreflect.FileDescriptor

var file_activity_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1c, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xde, 0x02, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x4d,
	0x61, 0x78, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x4d, 0x61, 0x78, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x68,
	0x61, 0x74, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x68, 0x61, 0x74,
	0x49, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x3c, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa8, 0x02, 0x0a, 0x08,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x12, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x1a, 0x12,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x73, 0x12, 0x0f, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x14, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x1a, 0x12,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x36, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x14, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x1a, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activity_proto_rawDescOnce sync.Once
	file_activity_proto_rawDescData = file_activity_proto_rawDesc
)

func file_activity_proto_rawDescGZIP() []byte {
	file_activity_proto_rawDescOnce.Do(func() {
		file_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_proto_rawDescData)
	})
	return file_activity_proto_rawDescData
}

var file_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_activity_proto_goTypes = []interface{}{
	(*Empty)(nil),         // 0: activity.Empty
	(*ActivityId)(nil),    // 1: activity.ActivityId
	(*Activity)(nil),      // 2: activity.Activity
	(*ListOfStrings)(nil), // 3: activity.ListOfStrings
	(*Response)(nil),      // 4: activity.Response
}
var file_activity_proto_depIdxs = []int32{
	3, // 0: activity.Activity.Type:type_name -> activity.ListOfStrings
	3, // 1: activity.Activity.Participant:type_name -> activity.ListOfStrings
	2, // 2: activity.activity.CreateActivity:input_type -> activity.Activity
	0, // 3: activity.activity.GetActivitys:input_type -> activity.Empty
	1, // 4: activity.activity.GetActivity:input_type -> activity.ActivityId
	2, // 5: activity.activity.EditActivity:input_type -> activity.Activity
	1, // 6: activity.activity.DeleteActivity:input_type -> activity.ActivityId
	4, // 7: activity.activity.CreateActivity:output_type -> activity.Response
	2, // 8: activity.activity.GetActivitys:output_type -> activity.Activity
	2, // 9: activity.activity.GetActivity:output_type -> activity.Activity
	4, // 10: activity.activity.EditActivity:output_type -> activity.Response
	4, // 11: activity.activity.DeleteActivity:output_type -> activity.Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_activity_proto_init() }
func file_activity_proto_init() {
	if File_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityId); i {
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
		file_activity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_activity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfStrings); i {
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
		file_activity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activity_proto_goTypes,
		DependencyIndexes: file_activity_proto_depIdxs,
		MessageInfos:      file_activity_proto_msgTypes,
	}.Build()
	File_activity_proto = out.File
	file_activity_proto_rawDesc = nil
	file_activity_proto_goTypes = nil
	file_activity_proto_depIdxs = nil
}
