// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: oci/descriptor.proto

package gen

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Descriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaType    string            `protobuf:"bytes,1,opt,name=MediaType,proto3" json:"MediaType,omitempty"`
	Digest       string            `protobuf:"bytes,2,opt,name=Digest,proto3" json:"Digest,omitempty"`
	Size         int64             `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	URLs         []string          `protobuf:"bytes,4,rep,name=URLs,proto3" json:"URLs,omitempty"`
	Data         []byte            `protobuf:"bytes,5,opt,name=Data,proto3" json:"Data,omitempty"`
	Platform     *Platform         `protobuf:"bytes,6,opt,name=Platform,proto3,oneof" json:"Platform,omitempty"`
	ArtifactType *string           `protobuf:"bytes,7,opt,name=ArtifactType,proto3,oneof" json:"ArtifactType,omitempty"`
	Annotations  map[string]string `protobuf:"bytes,8,rep,name=Annotations,proto3" json:"Annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Descriptor) Reset() {
	*x = Descriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oci_descriptor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Descriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Descriptor) ProtoMessage() {}

func (x *Descriptor) ProtoReflect() protoreflect.Message {
	mi := &file_oci_descriptor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Descriptor.ProtoReflect.Descriptor instead.
func (*Descriptor) Descriptor() ([]byte, []int) {
	return file_oci_descriptor_proto_rawDescGZIP(), []int{0}
}

func (x *Descriptor) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *Descriptor) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Descriptor) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Descriptor) GetURLs() []string {
	if x != nil {
		return x.URLs
	}
	return nil
}

func (x *Descriptor) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Descriptor) GetPlatform() *Platform {
	if x != nil {
		return x.Platform
	}
	return nil
}

func (x *Descriptor) GetArtifactType() string {
	if x != nil && x.ArtifactType != nil {
		return *x.ArtifactType
	}
	return ""
}

func (x *Descriptor) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type Platform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Architecture string   `protobuf:"bytes,1,opt,name=Architecture,proto3" json:"Architecture,omitempty"`
	OS           string   `protobuf:"bytes,2,opt,name=OS,proto3" json:"OS,omitempty"`
	OSVersion    *string  `protobuf:"bytes,3,opt,name=OSVersion,proto3,oneof" json:"OSVersion,omitempty"`
	OSFeatures   []string `protobuf:"bytes,4,rep,name=OSFeatures,proto3" json:"OSFeatures,omitempty"`
	Variant      *string  `protobuf:"bytes,5,opt,name=Variant,proto3,oneof" json:"Variant,omitempty"`
}

func (x *Platform) Reset() {
	*x = Platform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oci_descriptor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform) ProtoMessage() {}

func (x *Platform) ProtoReflect() protoreflect.Message {
	mi := &file_oci_descriptor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform.ProtoReflect.Descriptor instead.
func (*Platform) Descriptor() ([]byte, []int) {
	return file_oci_descriptor_proto_rawDescGZIP(), []int{1}
}

func (x *Platform) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *Platform) GetOS() string {
	if x != nil {
		return x.OS
	}
	return ""
}

func (x *Platform) GetOSVersion() string {
	if x != nil && x.OSVersion != nil {
		return *x.OSVersion
	}
	return ""
}

func (x *Platform) GetOSFeatures() []string {
	if x != nil {
		return x.OSFeatures
	}
	return nil
}

func (x *Platform) GetVariant() string {
	if x != nil && x.Variant != nil {
		return *x.Variant
	}
	return ""
}

var File_oci_descriptor_proto protoreflect.FileDescriptor

var file_oci_descriptor_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x63, 0x69, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x63, 0x69, 0x5f, 0x76, 0x31, 0x22, 0xff,
	0x02, 0x0a, 0x0a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x44,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x52, 0x4c, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x31, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x63, 0x69, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x48, 0x00, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x88,
	0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x0b, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x6f, 0x63, 0x69, 0x5f, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xba, 0x01, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x22, 0x0a,
	0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x53, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f,
	0x53, 0x12, 0x21, 0x0a, 0x09, 0x4f, 0x53, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x4f, 0x53, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x53, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x4f, 0x53, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x4f, 0x53, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oci_descriptor_proto_rawDescOnce sync.Once
	file_oci_descriptor_proto_rawDescData = file_oci_descriptor_proto_rawDesc
)

func file_oci_descriptor_proto_rawDescGZIP() []byte {
	file_oci_descriptor_proto_rawDescOnce.Do(func() {
		file_oci_descriptor_proto_rawDescData = protoimpl.X.CompressGZIP(file_oci_descriptor_proto_rawDescData)
	})
	return file_oci_descriptor_proto_rawDescData
}

var file_oci_descriptor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_oci_descriptor_proto_goTypes = []interface{}{
	(*Descriptor)(nil), // 0: oci_v1.Descriptor
	(*Platform)(nil),   // 1: oci_v1.Platform
	nil,                // 2: oci_v1.Descriptor.AnnotationsEntry
}
var file_oci_descriptor_proto_depIdxs = []int32{
	1, // 0: oci_v1.Descriptor.Platform:type_name -> oci_v1.Platform
	2, // 1: oci_v1.Descriptor.Annotations:type_name -> oci_v1.Descriptor.AnnotationsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_oci_descriptor_proto_init() }
func file_oci_descriptor_proto_init() {
	if File_oci_descriptor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oci_descriptor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Descriptor); i {
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
		file_oci_descriptor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform); i {
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
	file_oci_descriptor_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_oci_descriptor_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oci_descriptor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oci_descriptor_proto_goTypes,
		DependencyIndexes: file_oci_descriptor_proto_depIdxs,
		MessageInfos:      file_oci_descriptor_proto_msgTypes,
	}.Build()
	File_oci_descriptor_proto = out.File
	file_oci_descriptor_proto_rawDesc = nil
	file_oci_descriptor_proto_goTypes = nil
	file_oci_descriptor_proto_depIdxs = nil
}
