// Code generated by protoc-gen-go.
// source: common/common.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto
	common/configuration.proto

It has these top-level messages:
	Header
	ChainHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigurationEnvelope
	SignedConfigurationItem
	ConfigurationItem
	ConfigurationSignature
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_UNKNOWN               Status = 0
	Status_SUCCESS               Status = 200
	Status_BAD_REQUEST           Status = 400
	Status_FORBIDDEN             Status = 403
	Status_NOT_FOUND             Status = 404
	Status_INTERNAL_SERVER_ERROR Status = 500
	Status_SERVICE_UNAVAILABLE   Status = 503
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	200: "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	500: "INTERNAL_SERVER_ERROR",
	503: "SERVICE_UNAVAILABLE",
}
var Status_value = map[string]int32{
	"UNKNOWN":               0,
	"SUCCESS":               200,
	"BAD_REQUEST":           400,
	"FORBIDDEN":             403,
	"NOT_FOUND":             404,
	"INTERNAL_SERVER_ERROR": 500,
	"SERVICE_UNAVAILABLE":   503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HeaderType int32

const (
	HeaderType_MESSAGE                   HeaderType = 0
	HeaderType_CONFIGURATION_TRANSACTION HeaderType = 1
	HeaderType_CONFIGURATION_ITEM        HeaderType = 2
	HeaderType_ENDORSER_TRANSACTION      HeaderType = 3
)

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIGURATION_TRANSACTION",
	2: "CONFIGURATION_ITEM",
	3: "ENDORSER_TRANSACTION",
}
var HeaderType_value = map[string]int32{
	"MESSAGE":                   0,
	"CONFIGURATION_TRANSACTION": 1,
	"CONFIGURATION_ITEM":        2,
	"ENDORSER_TRANSACTION":      3,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}
func (HeaderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Header struct {
	ChainHeader     *ChainHeader     `protobuf:"bytes,1,opt,name=chainHeader" json:"chainHeader,omitempty"`
	SignatureHeader *SignatureHeader `protobuf:"bytes,2,opt,name=signatureHeader" json:"signatureHeader,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Header) GetChainHeader() *ChainHeader {
	if m != nil {
		return m.ChainHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() *SignatureHeader {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChainHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// Identifier of the chain this message is bound for
	ChainID []byte `protobuf:"bytes,4,opt,name=chainID,proto3" json:"chainID,omitempty"`
	// An unique identifier that is used end-to-end.
	//  -  set by higher layers such as end user or SDK
	//  -  passed to the endorser (which will check for uniqueness)
	//  -  as the header is passed along unchanged, it will be
	//     be retrieved by the committer (uniqueness check here as well)
	//  -  to be stored in the ledger
	TxID string `protobuf:"bytes,5,opt,name=txID" json:"txID,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch uint64 `protobuf:"varint,6,opt,name=epoch" json:"epoch,omitempty"`
	// Extension that may be attached based on the header type
	Extension []byte `protobuf:"bytes,7,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *ChainHeader) Reset()                    { *m = ChainHeader{} }
func (m *ChainHeader) String() string            { return proto.CompactTextString(m) }
func (*ChainHeader) ProtoMessage()               {}
func (*ChainHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChainHeader) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type SignatureHeader struct {
	// Creator of the message, specified as a certificate chain
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *SignatureHeader) Reset()                    { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string            { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()               {}
func (*SignatureHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Payload is the message contents (and header to allow for signing)
type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Payload) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

// Envelope wraps a Payload with a signature so that the message may be authenticated
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// This is finalized block structure to be shared among the orderer and peer
// Note that the BlockHeader chains to the previous BlockHeader, and the BlockData hash is embedded
// in the BlockHeader.  This makes it natural and obvious that the Data is included in the hash, but
// the Metadata is not.
type Block struct {
	Header   *BlockHeader   `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"`
	Data     *BlockData     `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Metadata *BlockMetadata `protobuf:"bytes,3,opt,name=Metadata" json:"Metadata,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetMetadata() *BlockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type BlockHeader struct {
	Number       uint64 `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	PreviousHash []byte `protobuf:"bytes,2,opt,name=PreviousHash,proto3" json:"PreviousHash,omitempty"`
	DataHash     []byte `protobuf:"bytes,3,opt,name=DataHash,proto3" json:"DataHash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type BlockData struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *BlockData) Reset()                    { *m = BlockData{} }
func (m *BlockData) String() string            { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()               {}
func (*BlockData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type BlockMetadata struct {
	Metadata [][]byte `protobuf:"bytes,1,rep,name=Metadata,proto3" json:"Metadata,omitempty"`
}

func (m *BlockMetadata) Reset()                    { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string            { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()               {}
func (*BlockMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*Header)(nil), "common.Header")
	proto.RegisterType((*ChainHeader)(nil), "common.ChainHeader")
	proto.RegisterType((*SignatureHeader)(nil), "common.SignatureHeader")
	proto.RegisterType((*Payload)(nil), "common.Payload")
	proto.RegisterType((*Envelope)(nil), "common.Envelope")
	proto.RegisterType((*Block)(nil), "common.Block")
	proto.RegisterType((*BlockHeader)(nil), "common.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "common.BlockData")
	proto.RegisterType((*BlockMetadata)(nil), "common.BlockMetadata")
	proto.RegisterEnum("common.Status", Status_name, Status_value)
	proto.RegisterEnum("common.HeaderType", HeaderType_name, HeaderType_value)
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x54, 0x4b, 0x6f, 0xd3, 0x4a,
	0x18, 0xbd, 0xae, 0xf3, 0x68, 0xbe, 0xe4, 0xf6, 0xfa, 0x4e, 0x5b, 0x08, 0x15, 0xa8, 0x95, 0x25,
	0x10, 0x6a, 0x45, 0x22, 0x8a, 0x90, 0xd8, 0x3a, 0xc9, 0xb4, 0xb5, 0x68, 0xc7, 0x65, 0xec, 0x14,
	0x89, 0x4d, 0xe4, 0x24, 0xd3, 0x24, 0x90, 0x78, 0x22, 0xc7, 0xa9, 0xda, 0x2d, 0x5b, 0x24, 0x84,
	0x04, 0x3f, 0x8a, 0x7f, 0xc0, 0x1f, 0x41, 0x62, 0xcb, 0x3c, 0xec, 0x3c, 0xba, 0xf2, 0x9c, 0xef,
	0x79, 0xce, 0x19, 0xdb, 0xb0, 0xdd, 0xe3, 0x93, 0x09, 0x8f, 0xea, 0xfa, 0x51, 0x9b, 0xc6, 0x3c,
	0xe1, 0xa8, 0xa0, 0xd1, 0xde, 0xfe, 0x80, 0xf3, 0xc1, 0x98, 0xd5, 0x55, 0xb4, 0x3b, 0xbf, 0xae,
	0x27, 0xa3, 0x09, 0x9b, 0x25, 0xe1, 0x64, 0xaa, 0x0b, 0xed, 0xcf, 0x06, 0x14, 0xce, 0x58, 0xd8,
	0x67, 0x31, 0x7a, 0x0d, 0xe5, 0xde, 0x30, 0x1c, 0x45, 0x1a, 0x56, 0x8d, 0x03, 0xe3, 0x79, 0xf9,
	0x78, 0xbb, 0x96, 0xce, 0x6d, 0x2e, 0x53, 0x74, 0xb5, 0x0e, 0x39, 0xf0, 0xdf, 0x6c, 0x34, 0x88,
	0xc2, 0x64, 0x1e, 0xb3, 0xb4, 0x75, 0x43, 0xb5, 0x3e, 0xcc, 0x5a, 0xfd, 0xf5, 0x34, 0xbd, 0x5f,
	0x6f, 0xff, 0x32, 0xa0, 0xbc, 0x32, 0x1f, 0x21, 0xc8, 0x25, 0x77, 0x53, 0xa6, 0x28, 0xe4, 0xa9,
	0x3a, 0xa3, 0x2a, 0x14, 0x6f, 0x58, 0x3c, 0x1b, 0xf1, 0x48, 0x8d, 0xcf, 0xd3, 0x0c, 0xa2, 0x37,
	0x50, 0x5a, 0xa8, 0xaa, 0x9a, 0x6a, 0xf5, 0x5e, 0x4d, 0xeb, 0xae, 0x65, 0xba, 0x6b, 0x41, 0x56,
	0x41, 0x97, 0xc5, 0x72, 0xa6, 0x52, 0xe2, 0xb6, 0xaa, 0x39, 0xd1, 0x57, 0xa1, 0x19, 0x54, 0x0c,
	0x6e, 0x45, 0x38, 0x2f, 0xc2, 0x25, 0xaa, 0xce, 0x68, 0x07, 0xf2, 0x6c, 0xca, 0x7b, 0xc3, 0x6a,
	0x41, 0x04, 0x73, 0x54, 0x03, 0xf4, 0x18, 0x4a, 0xec, 0x36, 0x61, 0x91, 0x62, 0x56, 0x54, 0x53,
	0x96, 0x01, 0x5b, 0x98, 0x73, 0x4f, 0xbd, 0x5a, 0x1a, 0xb3, 0x30, 0xe1, 0xda, 0x62, 0xb9, 0x54,
	0x43, 0xb9, 0x20, 0xe2, 0x51, 0x8f, 0x29, 0x81, 0x15, 0xaa, 0x81, 0x8d, 0xa1, 0x78, 0x19, 0xde,
	0x8d, 0x79, 0xd8, 0x47, 0xcf, 0xa0, 0x30, 0x5c, 0xbd, 0x9c, 0xad, 0xcc, 0xe1, 0xd4, 0xd8, 0x34,
	0x2b, 0xd9, 0xf7, 0xc3, 0x24, 0x4c, 0xe7, 0xa8, 0xb3, 0xdd, 0x80, 0x4d, 0x1c, 0xdd, 0xb0, 0x31,
	0xd7, 0x5e, 0x4e, 0xf5, 0xc8, 0x8c, 0x42, 0x0a, 0xa5, 0x9a, 0xc5, 0xe5, 0xa4, 0xed, 0xcb, 0x80,
	0xfd, 0xd5, 0x80, 0x7c, 0x63, 0xcc, 0x7b, 0x9f, 0xd0, 0x51, 0xf6, 0xd6, 0xdc, 0x7f, 0x4d, 0x54,
	0x3a, 0xa3, 0x93, 0x2a, 0x7e, 0x0a, 0xb9, 0x56, 0x46, 0xa7, 0x7c, 0xfc, 0xff, 0x5a, 0xa9, 0x4c,
	0x50, 0x95, 0x46, 0x2f, 0x61, 0xf3, 0x82, 0x25, 0xa1, 0x62, 0xae, 0xaf, 0x71, 0x77, 0xad, 0x34,
	0x4b, 0xd2, 0x45, 0x99, 0xcd, 0xa0, 0xbc, 0xb2, 0x10, 0x3d, 0x80, 0x02, 0x99, 0x4f, 0xba, 0x29,
	0xab, 0x1c, 0x4d, 0x11, 0xb2, 0xa1, 0x72, 0x19, 0xb3, 0x9b, 0x11, 0x9f, 0xcf, 0xce, 0xc2, 0xd9,
	0x30, 0x15, 0xb6, 0x16, 0x43, 0x7b, 0xb0, 0x29, 0x59, 0xa8, 0xbc, 0xa9, 0xf2, 0x0b, 0x6c, 0xef,
	0x43, 0x69, 0x41, 0x56, 0x9a, 0xab, 0xd4, 0x18, 0x07, 0xa6, 0x34, 0x57, 0x9e, 0xed, 0x23, 0xf8,
	0x77, 0x8d, 0xa2, 0x9c, 0xb6, 0xd0, 0xa2, 0x0b, 0x17, 0xf8, 0xf0, 0x8b, 0xf8, 0xe4, 0xfc, 0x44,
	0x38, 0x3a, 0x43, 0x65, 0x28, 0xb6, 0xc9, 0x5b, 0xe2, 0xbd, 0x27, 0xd6, 0x3f, 0xa8, 0x02, 0x45,
	0xbf, 0xdd, 0x6c, 0x62, 0xdf, 0xb7, 0x7e, 0x1a, 0xc8, 0x12, 0xd2, 0x9c, 0x56, 0x87, 0xe2, 0x77,
	0x6d, 0xec, 0x07, 0xd6, 0x37, 0x13, 0x6d, 0x41, 0xe9, 0xc4, 0xa3, 0x0d, 0xb7, 0xd5, 0xc2, 0xc4,
	0xfa, 0xae, 0x30, 0xf1, 0x82, 0xce, 0x89, 0xd7, 0x26, 0x2d, 0xeb, 0x87, 0x29, 0x76, 0xee, 0xba,
	0x24, 0xc0, 0x94, 0x38, 0xe7, 0x1d, 0x1f, 0xd3, 0x2b, 0x4c, 0x3b, 0x98, 0x52, 0x8f, 0x5a, 0xbf,
	0x4d, 0x71, 0xe3, 0xdb, 0x32, 0xe4, 0x36, 0x71, 0xa7, 0x4d, 0x9c, 0x2b, 0xc7, 0x3d, 0x77, 0x1a,
	0xe7, 0xd8, 0xfa, 0x63, 0x1e, 0x7e, 0x04, 0xd0, 0xee, 0x05, 0xf2, 0x2b, 0x13, 0x84, 0x2e, 0x04,
	0x01, 0xe7, 0x14, 0x0b, 0x42, 0x4f, 0xe0, 0x51, 0xd3, 0x23, 0x27, 0xee, 0x69, 0x9b, 0x3a, 0x81,
	0xeb, 0x91, 0x4e, 0x40, 0x1d, 0xe2, 0x3b, 0x4d, 0x79, 0xb6, 0x0c, 0xe1, 0x36, 0x5a, 0x4f, 0xbb,
	0x01, 0xbe, 0xb0, 0x36, 0xc4, 0xae, 0x1d, 0x4c, 0x5a, 0x1e, 0x15, 0x0b, 0xd7, 0x3a, 0xcc, 0xc6,
	0x8b, 0x0f, 0x47, 0x83, 0x51, 0x32, 0x9c, 0x77, 0xe5, 0xbd, 0xd6, 0x87, 0x62, 0x61, 0x3c, 0x66,
	0xfd, 0x01, 0x8b, 0xeb, 0xd7, 0x61, 0x37, 0x1e, 0xf5, 0xf4, 0x6f, 0x6a, 0x96, 0xfe, 0xca, 0xba,
	0x05, 0x05, 0x5f, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x46, 0x37, 0x79, 0xe2, 0x04, 0x00,
	0x00,
}