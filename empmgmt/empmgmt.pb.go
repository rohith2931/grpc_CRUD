// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: empmgmt/empmgmt.proto

package grpc

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

type NewEmp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mid  int32  `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	Did  int32  `protobuf:"varint,3,opt,name=did,proto3" json:"did,omitempty"`
}

func (x *NewEmp) Reset() {
	*x = NewEmp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_empmgmt_empmgmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEmp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEmp) ProtoMessage() {}

func (x *NewEmp) ProtoReflect() protoreflect.Message {
	mi := &file_empmgmt_empmgmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEmp.ProtoReflect.Descriptor instead.
func (*NewEmp) Descriptor() ([]byte, []int) {
	return file_empmgmt_empmgmt_proto_rawDescGZIP(), []int{0}
}

func (x *NewEmp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewEmp) GetMid() int32 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *NewEmp) GetDid() int32 {
	if x != nil {
		return x.Did
	}
	return 0
}

type Emp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mid  int32  `protobuf:"varint,3,opt,name=mid,proto3" json:"mid,omitempty"`
	Did  int32  `protobuf:"varint,4,opt,name=did,proto3" json:"did,omitempty"`
}

func (x *Emp) Reset() {
	*x = Emp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_empmgmt_empmgmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emp) ProtoMessage() {}

func (x *Emp) ProtoReflect() protoreflect.Message {
	mi := &file_empmgmt_empmgmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emp.ProtoReflect.Descriptor instead.
func (*Emp) Descriptor() ([]byte, []int) {
	return file_empmgmt_empmgmt_proto_rawDescGZIP(), []int{1}
}

func (x *Emp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Emp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Emp) GetMid() int32 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *Emp) GetDid() int32 {
	if x != nil {
		return x.Did
	}
	return 0
}

type GetEmpParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEmpParams) Reset() {
	*x = GetEmpParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_empmgmt_empmgmt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmpParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmpParams) ProtoMessage() {}

func (x *GetEmpParams) ProtoReflect() protoreflect.Message {
	mi := &file_empmgmt_empmgmt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmpParams.ProtoReflect.Descriptor instead.
func (*GetEmpParams) Descriptor() ([]byte, []int) {
	return file_empmgmt_empmgmt_proto_rawDescGZIP(), []int{2}
}

type EmpList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employees []*Emp `protobuf:"bytes,1,rep,name=Employees,proto3" json:"Employees,omitempty"`
}

func (x *EmpList) Reset() {
	*x = EmpList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_empmgmt_empmgmt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmpList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmpList) ProtoMessage() {}

func (x *EmpList) ProtoReflect() protoreflect.Message {
	mi := &file_empmgmt_empmgmt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmpList.ProtoReflect.Descriptor instead.
func (*EmpList) Descriptor() ([]byte, []int) {
	return file_empmgmt_empmgmt_proto_rawDescGZIP(), []int{3}
}

func (x *EmpList) GetEmployees() []*Emp {
	if x != nil {
		return x.Employees
	}
	return nil
}

var File_empmgmt_empmgmt_proto protoreflect.FileDescriptor

var file_empmgmt_empmgmt_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x6d, 0x70, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x65, 0x6d, 0x70, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x6d, 0x70, 0x6d, 0x67, 0x6d, 0x74,
	0x22, 0x40, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x45, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64,
	0x69, 0x64, 0x22, 0x4d, 0x0a, 0x03, 0x45, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x69,
	0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x22, 0x35, 0x0a, 0x07, 0x45, 0x6d, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x09,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x65, 0x6d, 0x70, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x52, 0x09, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x32, 0xcc, 0x01, 0x0a, 0x0d, 0x45, 0x6d, 0x70,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x45, 0x6d, 0x70, 0x12, 0x0f, 0x2e, 0x65, 0x6d, 0x70,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x6d, 0x70, 0x1a, 0x0c, 0x2e, 0x65, 0x6d,
	0x70, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x45, 0x6d, 0x70, 0x73, 0x12, 0x15, 0x2e, 0x65, 0x6d, 0x70, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e,
	0x65, 0x6d, 0x70, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x00, 0x12, 0x29, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x12, 0x0c,
	0x2e, 0x65, 0x6d, 0x70, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x1a, 0x0c, 0x2e, 0x65,
	0x6d, 0x70, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x12, 0x0c, 0x2e, 0x65, 0x6d, 0x70, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x1a, 0x0c, 0x2e, 0x65, 0x6d, 0x70, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x45, 0x6d, 0x70, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_empmgmt_empmgmt_proto_rawDescOnce sync.Once
	file_empmgmt_empmgmt_proto_rawDescData = file_empmgmt_empmgmt_proto_rawDesc
)

func file_empmgmt_empmgmt_proto_rawDescGZIP() []byte {
	file_empmgmt_empmgmt_proto_rawDescOnce.Do(func() {
		file_empmgmt_empmgmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_empmgmt_empmgmt_proto_rawDescData)
	})
	return file_empmgmt_empmgmt_proto_rawDescData
}

var file_empmgmt_empmgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_empmgmt_empmgmt_proto_goTypes = []interface{}{
	(*NewEmp)(nil),       // 0: empmgmt.NewEmp
	(*Emp)(nil),          // 1: empmgmt.Emp
	(*GetEmpParams)(nil), // 2: empmgmt.GetEmpParams
	(*EmpList)(nil),      // 3: empmgmt.EmpList
}
var file_empmgmt_empmgmt_proto_depIdxs = []int32{
	1, // 0: empmgmt.EmpList.Employees:type_name -> empmgmt.Emp
	0, // 1: empmgmt.EmpManagement.CreateNewEmp:input_type -> empmgmt.NewEmp
	2, // 2: empmgmt.EmpManagement.GetEmps:input_type -> empmgmt.GetEmpParams
	1, // 3: empmgmt.EmpManagement.UpdateEmp:input_type -> empmgmt.Emp
	1, // 4: empmgmt.EmpManagement.DeleteEmp:input_type -> empmgmt.Emp
	1, // 5: empmgmt.EmpManagement.CreateNewEmp:output_type -> empmgmt.Emp
	3, // 6: empmgmt.EmpManagement.GetEmps:output_type -> empmgmt.EmpList
	1, // 7: empmgmt.EmpManagement.UpdateEmp:output_type -> empmgmt.Emp
	1, // 8: empmgmt.EmpManagement.DeleteEmp:output_type -> empmgmt.Emp
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_empmgmt_empmgmt_proto_init() }
func file_empmgmt_empmgmt_proto_init() {
	if File_empmgmt_empmgmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_empmgmt_empmgmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEmp); i {
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
		file_empmgmt_empmgmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emp); i {
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
		file_empmgmt_empmgmt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmpParams); i {
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
		file_empmgmt_empmgmt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmpList); i {
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
			RawDescriptor: file_empmgmt_empmgmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_empmgmt_empmgmt_proto_goTypes,
		DependencyIndexes: file_empmgmt_empmgmt_proto_depIdxs,
		MessageInfos:      file_empmgmt_empmgmt_proto_msgTypes,
	}.Build()
	File_empmgmt_empmgmt_proto = out.File
	file_empmgmt_empmgmt_proto_rawDesc = nil
	file_empmgmt_empmgmt_proto_goTypes = nil
	file_empmgmt_empmgmt_proto_depIdxs = nil
}
