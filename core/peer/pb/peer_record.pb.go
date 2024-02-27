// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: pb/peer_record.proto

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

// PeerRecord messages contain information that is useful to share with other peers.
// Currently, a PeerRecord contains the public listen addresses for a peer, but this
// is expected to expand to include other information in the future.
//
// PeerRecords are designed to be serialized to bytes and placed inside of
// SignedEnvelopes before sharing with other peers.
// See https://github.com/thusharprakash/go-libp2p/core/record/pb/envelope.proto for
// the SignedEnvelope definition.
type PeerRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// peer_id contains a libp2p peer id in its binary representation.
	PeerId []byte `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	// seq contains a monotonically-increasing sequence counter to order PeerRecords in time.
	Seq uint64 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	// addresses is a list of public listen addresses for the peer.
	Addresses []*PeerRecord_AddressInfo `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *PeerRecord) Reset() {
	*x = PeerRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_peer_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerRecord) ProtoMessage() {}

func (x *PeerRecord) ProtoReflect() protoreflect.Message {
	mi := &file_pb_peer_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerRecord.ProtoReflect.Descriptor instead.
func (*PeerRecord) Descriptor() ([]byte, []int) {
	return file_pb_peer_record_proto_rawDescGZIP(), []int{0}
}

func (x *PeerRecord) GetPeerId() []byte {
	if x != nil {
		return x.PeerId
	}
	return nil
}

func (x *PeerRecord) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *PeerRecord) GetAddresses() []*PeerRecord_AddressInfo {
	if x != nil {
		return x.Addresses
	}
	return nil
}

// AddressInfo is a wrapper around a binary multiaddr. It is defined as a
// separate message to allow us to add per-address metadata in the future.
type PeerRecord_AddressInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Multiaddr []byte `protobuf:"bytes,1,opt,name=multiaddr,proto3" json:"multiaddr,omitempty"`
}

func (x *PeerRecord_AddressInfo) Reset() {
	*x = PeerRecord_AddressInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_peer_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerRecord_AddressInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerRecord_AddressInfo) ProtoMessage() {}

func (x *PeerRecord_AddressInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_peer_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerRecord_AddressInfo.ProtoReflect.Descriptor instead.
func (*PeerRecord_AddressInfo) Descriptor() ([]byte, []int) {
	return file_pb_peer_record_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PeerRecord_AddressInfo) GetMultiaddr() []byte {
	if x != nil {
		return x.Multiaddr
	}
	return nil
}

var File_pb_peer_record_proto protoreflect.FileDescriptor

var file_pb_peer_record_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x62, 0x22,
	0xa3, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x3d, 0x0a, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x65, 0x65, 0x72, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x1a, 0x2b, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x61, 0x64, 0x64, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_peer_record_proto_rawDescOnce sync.Once
	file_pb_peer_record_proto_rawDescData = file_pb_peer_record_proto_rawDesc
)

func file_pb_peer_record_proto_rawDescGZIP() []byte {
	file_pb_peer_record_proto_rawDescOnce.Do(func() {
		file_pb_peer_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_peer_record_proto_rawDescData)
	})
	return file_pb_peer_record_proto_rawDescData
}

var file_pb_peer_record_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_peer_record_proto_goTypes = []interface{}{
	(*PeerRecord)(nil),             // 0: peer.pb.PeerRecord
	(*PeerRecord_AddressInfo)(nil), // 1: peer.pb.PeerRecord.AddressInfo
}
var file_pb_peer_record_proto_depIdxs = []int32{
	1, // 0: peer.pb.PeerRecord.addresses:type_name -> peer.pb.PeerRecord.AddressInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_peer_record_proto_init() }
func file_pb_peer_record_proto_init() {
	if File_pb_peer_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_peer_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerRecord); i {
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
		file_pb_peer_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerRecord_AddressInfo); i {
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
			RawDescriptor: file_pb_peer_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_peer_record_proto_goTypes,
		DependencyIndexes: file_pb_peer_record_proto_depIdxs,
		MessageInfos:      file_pb_peer_record_proto_msgTypes,
	}.Build()
	File_pb_peer_record_proto = out.File
	file_pb_peer_record_proto_rawDesc = nil
	file_pb_peer_record_proto_goTypes = nil
	file_pb_peer_record_proto_depIdxs = nil
}
