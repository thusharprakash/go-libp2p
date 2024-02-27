// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: pb/identify.proto

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

type Identify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// protocolVersion determines compatibility between peers
	ProtocolVersion *string `protobuf:"bytes,5,opt,name=protocolVersion" json:"protocolVersion,omitempty"` // e.g. ipfs/1.0.0
	// agentVersion is like a UserAgent string in browsers, or client version in bittorrent
	// includes the client name and client.
	AgentVersion *string `protobuf:"bytes,6,opt,name=agentVersion" json:"agentVersion,omitempty"` // e.g. go-ipfs/0.1.0
	// publicKey is this node's public key (which also gives its node.ID)
	// - may not need to be sent, as secure channel implies it has been sent.
	// - then again, if we change / disable secure channel, may still want it.
	PublicKey []byte `protobuf:"bytes,1,opt,name=publicKey" json:"publicKey,omitempty"`
	// listenAddrs are the multiaddrs the sender node listens for open connections on
	ListenAddrs [][]byte `protobuf:"bytes,2,rep,name=listenAddrs" json:"listenAddrs,omitempty"`
	// oservedAddr is the multiaddr of the remote endpoint that the sender node perceives
	// this is useful information to convey to the other side, as it helps the remote endpoint
	// determine whether its connection to the local peer goes through NAT.
	ObservedAddr []byte `protobuf:"bytes,4,opt,name=observedAddr" json:"observedAddr,omitempty"`
	// protocols are the services this node is running
	Protocols []string `protobuf:"bytes,3,rep,name=protocols" json:"protocols,omitempty"`
	// signedPeerRecord contains a serialized SignedEnvelope containing a PeerRecord,
	// signed by the sending node. It contains the same addresses as the listenAddrs field, but
	// in a form that lets us share authenticated addrs with other peers.
	// see github.com/thusharprakash/go-libp2p/core/record/pb/envelope.proto and
	// github.com/thusharprakash/go-libp2p/core/peer/pb/peer_record.proto for message definitions.
	SignedPeerRecord []byte `protobuf:"bytes,8,opt,name=signedPeerRecord" json:"signedPeerRecord,omitempty"`
}

func (x *Identify) Reset() {
	*x = Identify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_identify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identify) ProtoMessage() {}

func (x *Identify) ProtoReflect() protoreflect.Message {
	mi := &file_pb_identify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identify.ProtoReflect.Descriptor instead.
func (*Identify) Descriptor() ([]byte, []int) {
	return file_pb_identify_proto_rawDescGZIP(), []int{0}
}

func (x *Identify) GetProtocolVersion() string {
	if x != nil && x.ProtocolVersion != nil {
		return *x.ProtocolVersion
	}
	return ""
}

func (x *Identify) GetAgentVersion() string {
	if x != nil && x.AgentVersion != nil {
		return *x.AgentVersion
	}
	return ""
}

func (x *Identify) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Identify) GetListenAddrs() [][]byte {
	if x != nil {
		return x.ListenAddrs
	}
	return nil
}

func (x *Identify) GetObservedAddr() []byte {
	if x != nil {
		return x.ObservedAddr
	}
	return nil
}

func (x *Identify) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *Identify) GetSignedPeerRecord() []byte {
	if x != nil {
		return x.SignedPeerRecord
	}
	return nil
}

var File_pb_identify_proto protoreflect.FileDescriptor

var file_pb_identify_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x62,
	0x22, 0x86, 0x02, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79, 0x12, 0x28, 0x0a,
	0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x2a, 0x0a,
	0x10, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
}

var (
	file_pb_identify_proto_rawDescOnce sync.Once
	file_pb_identify_proto_rawDescData = file_pb_identify_proto_rawDesc
)

func file_pb_identify_proto_rawDescGZIP() []byte {
	file_pb_identify_proto_rawDescOnce.Do(func() {
		file_pb_identify_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_identify_proto_rawDescData)
	})
	return file_pb_identify_proto_rawDescData
}

var file_pb_identify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_identify_proto_goTypes = []interface{}{
	(*Identify)(nil), // 0: identify.pb.Identify
}
var file_pb_identify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_identify_proto_init() }
func file_pb_identify_proto_init() {
	if File_pb_identify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_identify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identify); i {
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
			RawDescriptor: file_pb_identify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_identify_proto_goTypes,
		DependencyIndexes: file_pb_identify_proto_depIdxs,
		MessageInfos:      file_pb_identify_proto_msgTypes,
	}.Build()
	File_pb_identify_proto = out.File
	file_pb_identify_proto_rawDesc = nil
	file_pb_identify_proto_goTypes = nil
	file_pb_identify_proto_depIdxs = nil
}
