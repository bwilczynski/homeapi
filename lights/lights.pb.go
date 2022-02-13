// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: lights.proto

package lights

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

type Light struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Light) Reset() {
	*x = Light{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lights_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Light) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Light) ProtoMessage() {}

func (x *Light) ProtoReflect() protoreflect.Message {
	mi := &file_lights_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Light.ProtoReflect.Descriptor instead.
func (*Light) Descriptor() ([]byte, []int) {
	return file_lights_proto_rawDescGZIP(), []int{0}
}

func (x *Light) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Light) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lights_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lights_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_lights_proto_rawDescGZIP(), []int{1}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lights []*Light `protobuf:"bytes,1,rep,name=lights,proto3" json:"lights,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lights_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lights_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_lights_proto_rawDescGZIP(), []int{2}
}

func (x *ListResponse) GetLights() []*Light {
	if x != nil {
		return x.Lights
	}
	return nil
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lights []string `protobuf:"bytes,3,rep,name=lights,proto3" json:"lights,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lights_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_lights_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_lights_proto_rawDescGZIP(), []int{3}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetLights() []string {
	if x != nil {
		return x.Lights
	}
	return nil
}

type ListGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGroupsRequest) Reset() {
	*x = ListGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lights_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupsRequest) ProtoMessage() {}

func (x *ListGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lights_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListGroupsRequest) Descriptor() ([]byte, []int) {
	return file_lights_proto_rawDescGZIP(), []int{4}
}

type ListGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *ListGroupsResponse) Reset() {
	*x = ListGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lights_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupsResponse) ProtoMessage() {}

func (x *ListGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lights_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupsResponse.ProtoReflect.Descriptor instead.
func (*ListGroupsResponse) Descriptor() ([]byte, []int) {
	return file_lights_proto_rawDescGZIP(), []int{5}
}

func (x *ListGroupsResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type ToggleGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *ToggleGroupRequest) Reset() {
	*x = ToggleGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lights_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToggleGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleGroupRequest) ProtoMessage() {}

func (x *ToggleGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lights_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleGroupRequest.ProtoReflect.Descriptor instead.
func (*ToggleGroupRequest) Descriptor() ([]byte, []int) {
	return file_lights_proto_rawDescGZIP(), []int{6}
}

func (x *ToggleGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type ToggleGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ToggleGroupResponse) Reset() {
	*x = ToggleGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lights_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToggleGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToggleGroupResponse) ProtoMessage() {}

func (x *ToggleGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lights_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToggleGroupResponse.ProtoReflect.Descriptor instead.
func (*ToggleGroupResponse) Descriptor() ([]byte, []int) {
	return file_lights_proto_rawDescGZIP(), []int{7}
}

var File_lights_proto protoreflect.FileDescriptor

var file_lights_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b,
	0x0a, 0x05, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x06, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x43, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22,
	0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x2f, 0x0a, 0x12, 0x54, 0x6f,
	0x67, 0x67, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x54,
	0x6f, 0x67, 0x67, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xa4, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x0b, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x13,
	0x2e, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x77, 0x69, 0x6c, 0x63, 0x7a, 0x79, 0x6e,
	0x73, 0x6b, 0x69, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lights_proto_rawDescOnce sync.Once
	file_lights_proto_rawDescData = file_lights_proto_rawDesc
)

func file_lights_proto_rawDescGZIP() []byte {
	file_lights_proto_rawDescOnce.Do(func() {
		file_lights_proto_rawDescData = protoimpl.X.CompressGZIP(file_lights_proto_rawDescData)
	})
	return file_lights_proto_rawDescData
}

var file_lights_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_lights_proto_goTypes = []interface{}{
	(*Light)(nil),               // 0: Light
	(*ListRequest)(nil),         // 1: ListRequest
	(*ListResponse)(nil),        // 2: ListResponse
	(*Group)(nil),               // 3: Group
	(*ListGroupsRequest)(nil),   // 4: ListGroupsRequest
	(*ListGroupsResponse)(nil),  // 5: ListGroupsResponse
	(*ToggleGroupRequest)(nil),  // 6: ToggleGroupRequest
	(*ToggleGroupResponse)(nil), // 7: ToggleGroupResponse
}
var file_lights_proto_depIdxs = []int32{
	0, // 0: ListResponse.lights:type_name -> Light
	3, // 1: ListGroupsResponse.groups:type_name -> Group
	1, // 2: LightService.List:input_type -> ListRequest
	4, // 3: LightService.ListGroups:input_type -> ListGroupsRequest
	6, // 4: LightService.ToggleGroup:input_type -> ToggleGroupRequest
	2, // 5: LightService.List:output_type -> ListResponse
	5, // 6: LightService.ListGroups:output_type -> ListGroupsResponse
	7, // 7: LightService.ToggleGroup:output_type -> ToggleGroupResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lights_proto_init() }
func file_lights_proto_init() {
	if File_lights_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lights_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Light); i {
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
		file_lights_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_lights_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_lights_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_lights_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupsRequest); i {
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
		file_lights_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupsResponse); i {
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
		file_lights_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToggleGroupRequest); i {
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
		file_lights_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToggleGroupResponse); i {
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
			RawDescriptor: file_lights_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lights_proto_goTypes,
		DependencyIndexes: file_lights_proto_depIdxs,
		MessageInfos:      file_lights_proto_msgTypes,
	}.Build()
	File_lights_proto = out.File
	file_lights_proto_rawDesc = nil
	file_lights_proto_goTypes = nil
	file_lights_proto_depIdxs = nil
}
