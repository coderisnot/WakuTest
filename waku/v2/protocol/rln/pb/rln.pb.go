// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: rln.proto

// rfc: https://rfc.vac.dev/spec/17/

package pb

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

type RateLimitProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof         []byte `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
	MerkleRoot    []byte `protobuf:"bytes,2,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	Epoch         []byte `protobuf:"bytes,3,opt,name=epoch,proto3" json:"epoch,omitempty"`
	ShareX        []byte `protobuf:"bytes,4,opt,name=share_x,json=shareX,proto3" json:"share_x,omitempty"`
	ShareY        []byte `protobuf:"bytes,5,opt,name=share_y,json=shareY,proto3" json:"share_y,omitempty"`
	Nullifier     []byte `protobuf:"bytes,6,opt,name=nullifier,proto3" json:"nullifier,omitempty"`
	RlnIdentifier []byte `protobuf:"bytes,7,opt,name=rln_identifier,json=rlnIdentifier,proto3" json:"rln_identifier,omitempty"`
}

func (x *RateLimitProof) Reset() {
	*x = RateLimitProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rln_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitProof) ProtoMessage() {}

func (x *RateLimitProof) ProtoReflect() protoreflect.Message {
	mi := &file_rln_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitProof.ProtoReflect.Descriptor instead.
func (*RateLimitProof) Descriptor() ([]byte, []int) {
	return file_rln_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimitProof) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *RateLimitProof) GetMerkleRoot() []byte {
	if x != nil {
		return x.MerkleRoot
	}
	return nil
}

func (x *RateLimitProof) GetEpoch() []byte {
	if x != nil {
		return x.Epoch
	}
	return nil
}

func (x *RateLimitProof) GetShareX() []byte {
	if x != nil {
		return x.ShareX
	}
	return nil
}

func (x *RateLimitProof) GetShareY() []byte {
	if x != nil {
		return x.ShareY
	}
	return nil
}

func (x *RateLimitProof) GetNullifier() []byte {
	if x != nil {
		return x.Nullifier
	}
	return nil
}

func (x *RateLimitProof) GetRlnIdentifier() []byte {
	if x != nil {
		return x.RlnIdentifier
	}
	return nil
}

var File_rln_proto protoreflect.FileDescriptor

var file_rln_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x6c, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77, 0x61, 0x6b,
	0x75, 0x2e, 0x72, 0x6c, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xd4, 0x01, 0x0a, 0x0e, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x5f, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x58, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x59, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75,
	0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6e,
	0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x6c, 0x6e, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x72, 0x6c, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rln_proto_rawDescOnce sync.Once
	file_rln_proto_rawDescData = file_rln_proto_rawDesc
)

func file_rln_proto_rawDescGZIP() []byte {
	file_rln_proto_rawDescOnce.Do(func() {
		file_rln_proto_rawDescData = protoimpl.X.CompressGZIP(file_rln_proto_rawDescData)
	})
	return file_rln_proto_rawDescData
}

var file_rln_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rln_proto_goTypes = []interface{}{
	(*RateLimitProof)(nil), // 0: waku.rln.v1.RateLimitProof
}
var file_rln_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rln_proto_init() }
func file_rln_proto_init() {
	if File_rln_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rln_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitProof); i {
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
			RawDescriptor: file_rln_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rln_proto_goTypes,
		DependencyIndexes: file_rln_proto_depIdxs,
		MessageInfos:      file_rln_proto_msgTypes,
	}.Build()
	File_rln_proto = out.File
	file_rln_proto_rawDesc = nil
	file_rln_proto_goTypes = nil
	file_rln_proto_depIdxs = nil
}
