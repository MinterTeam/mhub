// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/v1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TxStatusType int32

const (
	TX_STATUS_NOT_FOUND        TxStatusType = 0
	TX_STATUS_DEPOSIT_RECEIVED TxStatusType = 1
	TX_STATUS_BATCH_CREATED    TxStatusType = 2
	TX_STATUS_BATCH_EXECUTED   TxStatusType = 3
	TX_STATUS_REFUNDED         TxStatusType = 4
)

var TxStatusType_name = map[int32]string{
	0: "TX_STATUS_NOT_FOUND",
	1: "TX_STATUS_DEPOSIT_RECEIVED",
	2: "TX_STATUS_BATCH_CREATED",
	3: "TX_STATUS_BATCH_EXECUTED",
	4: "TX_STATUS_REFUNDED",
}

var TxStatusType_value = map[string]int32{
	"TX_STATUS_NOT_FOUND":        0,
	"TX_STATUS_DEPOSIT_RECEIVED": 1,
	"TX_STATUS_BATCH_CREATED":    2,
	"TX_STATUS_BATCH_EXECUTED":   3,
	"TX_STATUS_REFUNDED":         4,
}

func (x TxStatusType) String() string {
	return proto.EnumName(TxStatusType_name, int32(x))
}

func (TxStatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b54af2de77c923e3, []int{0}
}

// It's difficult to serialize and deserialize
// interfaces, instead we can make this struct
// that stores all the data the interface requires
// and use it to store and then re-create a interface
// object with all the same properties.
type GenericClaim struct {
	Epoch        uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	ClaimType    int32  `protobuf:"varint,2,opt,name=claim_type,json=claimType,proto3" json:"claim_type,omitempty"`
	Hash         []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	EventClaimer string `protobuf:"bytes,4,opt,name=event_claimer,json=eventClaimer,proto3" json:"event_claimer,omitempty"`
	// Types that are valid to be assigned to Claim:
	//	*GenericClaim_PriceClaim
	Claim isGenericClaim_Claim `protobuf_oneof:"claim"`
}

func (m *GenericClaim) Reset()         { *m = GenericClaim{} }
func (m *GenericClaim) String() string { return proto.CompactTextString(m) }
func (*GenericClaim) ProtoMessage()    {}
func (*GenericClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_b54af2de77c923e3, []int{0}
}
func (m *GenericClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenericClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenericClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenericClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericClaim.Merge(m, src)
}
func (m *GenericClaim) XXX_Size() int {
	return m.Size()
}
func (m *GenericClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericClaim.DiscardUnknown(m)
}

var xxx_messageInfo_GenericClaim proto.InternalMessageInfo

type isGenericClaim_Claim interface {
	isGenericClaim_Claim()
	MarshalTo([]byte) (int, error)
	Size() int
}

type GenericClaim_PriceClaim struct {
	PriceClaim *MsgPriceClaim `protobuf:"bytes,5,opt,name=price_claim,json=priceClaim,proto3,oneof" json:"price_claim,omitempty"`
}

func (*GenericClaim_PriceClaim) isGenericClaim_Claim() {}

func (m *GenericClaim) GetClaim() isGenericClaim_Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

func (m *GenericClaim) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *GenericClaim) GetClaimType() int32 {
	if m != nil {
		return m.ClaimType
	}
	return 0
}

func (m *GenericClaim) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *GenericClaim) GetEventClaimer() string {
	if m != nil {
		return m.EventClaimer
	}
	return ""
}

func (m *GenericClaim) GetPriceClaim() *MsgPriceClaim {
	if x, ok := m.GetClaim().(*GenericClaim_PriceClaim); ok {
		return x.PriceClaim
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GenericClaim) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GenericClaim_PriceClaim)(nil),
	}
}

type Epoch struct {
	Nonce uint64  `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Votes []*Vote `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes,omitempty"`
}

func (m *Epoch) Reset()         { *m = Epoch{} }
func (m *Epoch) String() string { return proto.CompactTextString(m) }
func (*Epoch) ProtoMessage()    {}
func (*Epoch) Descriptor() ([]byte, []int) {
	return fileDescriptor_b54af2de77c923e3, []int{1}
}
func (m *Epoch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Epoch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Epoch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Epoch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Epoch.Merge(m, src)
}
func (m *Epoch) XXX_Size() int {
	return m.Size()
}
func (m *Epoch) XXX_DiscardUnknown() {
	xxx_messageInfo_Epoch.DiscardUnknown(m)
}

var xxx_messageInfo_Epoch proto.InternalMessageInfo

func (m *Epoch) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Epoch) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

type Vote struct {
	Oracle string         `protobuf:"bytes,1,opt,name=oracle,proto3" json:"oracle,omitempty"`
	Claim  *MsgPriceClaim `protobuf:"bytes,2,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_b54af2de77c923e3, []int{2}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetOracle() string {
	if m != nil {
		return m.Oracle
	}
	return ""
}

func (m *Vote) GetClaim() *MsgPriceClaim {
	if m != nil {
		return m.Claim
	}
	return nil
}

type Coin struct {
	Denom       string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	EthAddr     string `protobuf:"bytes,2,opt,name=eth_addr,json=ethAddr,proto3" json:"eth_addr,omitempty"`
	MinterId    uint64 `protobuf:"varint,3,opt,name=minter_id,json=minterId,proto3" json:"minter_id"`
	EthDecimals uint64 `protobuf:"varint,4,opt,name=eth_decimals,json=ethDecimals,proto3" json:"eth_decimals,omitempty"`
}

func (m *Coin) Reset()         { *m = Coin{} }
func (m *Coin) String() string { return proto.CompactTextString(m) }
func (*Coin) ProtoMessage()    {}
func (*Coin) Descriptor() ([]byte, []int) {
	return fileDescriptor_b54af2de77c923e3, []int{3}
}
func (m *Coin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Coin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Coin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Coin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coin.Merge(m, src)
}
func (m *Coin) XXX_Size() int {
	return m.Size()
}
func (m *Coin) XXX_DiscardUnknown() {
	xxx_messageInfo_Coin.DiscardUnknown(m)
}

var xxx_messageInfo_Coin proto.InternalMessageInfo

func (m *Coin) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Coin) GetEthAddr() string {
	if m != nil {
		return m.EthAddr
	}
	return ""
}

func (m *Coin) GetMinterId() uint64 {
	if m != nil {
		return m.MinterId
	}
	return 0
}

func (m *Coin) GetEthDecimals() uint64 {
	if m != nil {
		return m.EthDecimals
	}
	return 0
}

type TxStatus struct {
	InTxHash  string       `protobuf:"bytes,1,opt,name=in_tx_hash,json=inTxHash,proto3" json:"in_tx_hash,omitempty"`
	OutTxHash string       `protobuf:"bytes,2,opt,name=out_tx_hash,json=outTxHash,proto3" json:"out_tx_hash,omitempty"`
	Status    TxStatusType `protobuf:"varint,3,opt,name=status,proto3,enum=oracle.v1.TxStatusType" json:"status,omitempty"`
}

func (m *TxStatus) Reset()         { *m = TxStatus{} }
func (m *TxStatus) String() string { return proto.CompactTextString(m) }
func (*TxStatus) ProtoMessage()    {}
func (*TxStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b54af2de77c923e3, []int{4}
}
func (m *TxStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxStatus.Merge(m, src)
}
func (m *TxStatus) XXX_Size() int {
	return m.Size()
}
func (m *TxStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TxStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TxStatus proto.InternalMessageInfo

func (m *TxStatus) GetInTxHash() string {
	if m != nil {
		return m.InTxHash
	}
	return ""
}

func (m *TxStatus) GetOutTxHash() string {
	if m != nil {
		return m.OutTxHash
	}
	return ""
}

func (m *TxStatus) GetStatus() TxStatusType {
	if m != nil {
		return m.Status
	}
	return TX_STATUS_NOT_FOUND
}

func init() {
	proto.RegisterEnum("oracle.v1.TxStatusType", TxStatusType_name, TxStatusType_value)
	proto.RegisterType((*GenericClaim)(nil), "oracle.v1.GenericClaim")
	proto.RegisterType((*Epoch)(nil), "oracle.v1.Epoch")
	proto.RegisterType((*Vote)(nil), "oracle.v1.Vote")
	proto.RegisterType((*Coin)(nil), "oracle.v1.Coin")
	proto.RegisterType((*TxStatus)(nil), "oracle.v1.TxStatus")
}

func init() { proto.RegisterFile("oracle/v1/types.proto", fileDescriptor_b54af2de77c923e3) }

var fileDescriptor_b54af2de77c923e3 = []byte{
	// 647 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xc1, 0x4e, 0xdb, 0x4a,
	0x14, 0x8d, 0x83, 0x03, 0xc9, 0x4d, 0x78, 0x2f, 0x9a, 0x97, 0x07, 0xae, 0x4b, 0x5d, 0x37, 0x55,
	0xa5, 0x88, 0x45, 0x5c, 0xe8, 0xa6, 0x6a, 0xa5, 0x4a, 0x89, 0x6d, 0x4a, 0x16, 0x04, 0x34, 0x71,
	0x10, 0xea, 0xc6, 0x32, 0xf6, 0x28, 0xb6, 0x84, 0x3d, 0x91, 0x3d, 0x89, 0xc2, 0x0f, 0x54, 0x55,
	0x56, 0xfd, 0x81, 0xac, 0xba, 0xed, 0x57, 0x74, 0xd5, 0x25, 0xcb, 0xae, 0xaa, 0x0a, 0x76, 0xfd,
	0x8a, 0xca, 0x63, 0x43, 0x50, 0x05, 0xec, 0xee, 0x3d, 0xe7, 0xcc, 0x99, 0x33, 0x73, 0x3d, 0x86,
	0xff, 0x69, 0xec, 0xb8, 0x67, 0x44, 0x9b, 0xee, 0x68, 0xec, 0x7c, 0x4c, 0x92, 0xf6, 0x38, 0xa6,
	0x8c, 0xa2, 0x4a, 0x06, 0xb7, 0xa7, 0x3b, 0x72, 0x63, 0xa9, 0x08, 0x93, 0x51, 0x2e, 0x90, 0x1b,
	0x23, 0x3a, 0xa2, 0xbc, 0xd4, 0xd2, 0x2a, 0x43, 0x9b, 0xdf, 0x04, 0xa8, 0xbd, 0x27, 0x11, 0x89,
	0x03, 0x57, 0x3f, 0x73, 0x82, 0x10, 0x35, 0xa0, 0x44, 0xc6, 0xd4, 0xf5, 0x25, 0x41, 0x15, 0x5a,
	0x22, 0xce, 0x1a, 0xf4, 0x04, 0xc0, 0x4d, 0x69, 0x3b, 0xdd, 0x52, 0x2a, 0xaa, 0x42, 0xab, 0x84,
	0x2b, 0x1c, 0xb1, 0xce, 0xc7, 0x04, 0x21, 0x10, 0x7d, 0x27, 0xf1, 0xa5, 0x15, 0x55, 0x68, 0xd5,
	0x30, 0xaf, 0xd1, 0x73, 0x58, 0x27, 0x53, 0x12, 0x31, 0x9b, 0xcb, 0x48, 0x2c, 0x89, 0xaa, 0xd0,
	0xaa, 0xe0, 0x1a, 0x07, 0xf5, 0x0c, 0x43, 0x6f, 0xa1, 0x3a, 0x8e, 0x03, 0x97, 0x64, 0x22, 0xa9,
	0xa4, 0x0a, 0xad, 0xea, 0xae, 0xd4, 0xbe, 0x39, 0x4b, 0xfb, 0x20, 0x19, 0x1d, 0xa5, 0x02, 0xbe,
	0x60, 0xbf, 0x80, 0x61, 0x7c, 0xd3, 0x75, 0xd7, 0xa0, 0xc4, 0x97, 0x35, 0x0d, 0x28, 0x99, 0x3c,
	0x66, 0x03, 0x4a, 0x11, 0x8d, 0x5c, 0x72, 0x1d, 0x9e, 0x37, 0xe8, 0x05, 0x94, 0xa6, 0x94, 0x91,
	0x44, 0x2a, 0xaa, 0x2b, 0xad, 0xea, 0xee, 0xbf, 0xb7, 0xec, 0x8f, 0x29, 0x23, 0x38, 0x63, 0x9b,
	0x7d, 0x10, 0xd3, 0x16, 0x6d, 0xc0, 0x6a, 0x26, 0xe0, 0x2e, 0x15, 0x9c, 0x77, 0xa8, 0x9d, 0x6f,
	0xc7, 0x8f, 0xff, 0x40, 0x4a, 0x9c, 0xa7, 0xfa, 0x28, 0x80, 0xa8, 0xd3, 0x20, 0x4a, 0x53, 0x79,
	0x24, 0xa2, 0x61, 0xee, 0x97, 0x35, 0xe8, 0x11, 0x94, 0x09, 0xf3, 0x6d, 0xc7, 0xf3, 0x62, 0xee,
	0x58, 0xc1, 0x6b, 0x84, 0xf9, 0x1d, 0xcf, 0x8b, 0xd1, 0x36, 0x54, 0xc2, 0x20, 0x62, 0x24, 0xb6,
	0x03, 0x8f, 0xdf, 0xa9, 0xd8, 0x5d, 0xff, 0xfd, 0xf3, 0xe9, 0x12, 0xc4, 0xe5, 0xac, 0xec, 0x79,
	0xe8, 0x19, 0xd4, 0x52, 0x1b, 0x8f, 0xb8, 0x41, 0xe8, 0x9c, 0x25, 0xfc, 0x96, 0x45, 0x5c, 0x25,
	0xcc, 0x37, 0x72, 0xa8, 0x79, 0x0e, 0x65, 0x6b, 0x36, 0x60, 0x0e, 0x9b, 0x24, 0x68, 0x0b, 0x20,
	0x88, 0x6c, 0x36, 0xb3, 0xf9, 0xbc, 0xb2, 0x40, 0xe5, 0x20, 0xb2, 0x66, 0xfb, 0xe9, 0xcc, 0x14,
	0xa8, 0xd2, 0x09, 0xbb, 0xa1, 0xb3, 0x58, 0x15, 0x3a, 0x61, 0x39, 0xaf, 0xc1, 0x6a, 0xc2, 0x7d,
	0x78, 0xaa, 0x7f, 0x76, 0x37, 0x6f, 0xdd, 0xc1, 0xf5, 0x16, 0xe9, 0x07, 0x81, 0x73, 0xd9, 0xf6,
	0xd7, 0x22, 0xd4, 0x6e, 0x13, 0xe8, 0x25, 0xfc, 0x67, 0x9d, 0xd8, 0x03, 0xab, 0x63, 0x0d, 0x07,
	0x76, 0xff, 0xd0, 0xb2, 0xf7, 0x0e, 0x87, 0x7d, 0xa3, 0x5e, 0x90, 0x37, 0xe7, 0x0b, 0xf5, 0x2e,
	0x0a, 0xbd, 0x03, 0x79, 0x09, 0x1b, 0xe6, 0xd1, 0xe1, 0xa0, 0x67, 0xd9, 0xd8, 0xd4, 0xcd, 0xde,
	0xb1, 0x69, 0xd4, 0x05, 0x59, 0x99, 0x2f, 0xd4, 0x07, 0x14, 0xe8, 0x35, 0x6c, 0x2e, 0xd9, 0x6e,
	0xc7, 0xd2, 0xf7, 0x6d, 0x1d, 0x9b, 0x1d, 0xcb, 0x34, 0xea, 0x45, 0xf9, 0xf1, 0x7c, 0xa1, 0xde,
	0x47, 0xa3, 0x37, 0x20, 0xfd, 0x4d, 0x99, 0x27, 0xa6, 0x3e, 0x4c, 0x97, 0xae, 0xc8, 0x5b, 0xf3,
	0x85, 0x7a, 0x2f, 0x8f, 0xda, 0x80, 0x96, 0x1c, 0x36, 0xf7, 0x86, 0x7d, 0xc3, 0x34, 0xea, 0xa2,
	0xbc, 0x31, 0x5f, 0xa8, 0x77, 0x30, 0xb2, 0xf8, 0xe9, 0x8b, 0x52, 0xe8, 0xf6, 0xbe, 0x5f, 0x2a,
	0xc2, 0xc5, 0xa5, 0x22, 0xfc, 0xba, 0x54, 0x84, 0xcf, 0x57, 0x4a, 0xe1, 0xe2, 0x4a, 0x29, 0xfc,
	0xb8, 0x52, 0x0a, 0x1f, 0xb4, 0x51, 0xc0, 0xfc, 0xc9, 0x69, 0xdb, 0xa5, 0xa1, 0x76, 0xc0, 0x67,
	0x6f, 0x11, 0x27, 0xd4, 0x42, 0x7f, 0x72, 0xaa, 0xb9, 0xbe, 0x13, 0x44, 0xda, 0x4c, 0xcb, 0x9f,
	0x3d, 0xff, 0x2b, 0x9c, 0xae, 0xf2, 0xf7, 0xfd, 0xea, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37,
	0xdd, 0xc8, 0x4f, 0x2f, 0x04, 0x00, 0x00,
}

func (m *GenericClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenericClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenericClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claim != nil {
		{
			size := m.Claim.Size()
			i -= size
			if _, err := m.Claim.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.EventClaimer) > 0 {
		i -= len(m.EventClaimer)
		copy(dAtA[i:], m.EventClaimer)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EventClaimer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ClaimType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ClaimType))
		i--
		dAtA[i] = 0x10
	}
	if m.Epoch != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GenericClaim_PriceClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenericClaim_PriceClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PriceClaim != nil {
		{
			size, err := m.PriceClaim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Epoch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Epoch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Epoch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Nonce != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Oracle) > 0 {
		i -= len(m.Oracle)
		copy(dAtA[i:], m.Oracle)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Oracle)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Coin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Coin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Coin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EthDecimals != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.EthDecimals))
		i--
		dAtA[i] = 0x20
	}
	if m.MinterId != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MinterId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.EthAddr) > 0 {
		i -= len(m.EthAddr)
		copy(dAtA[i:], m.EthAddr)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EthAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.OutTxHash) > 0 {
		i -= len(m.OutTxHash)
		copy(dAtA[i:], m.OutTxHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.OutTxHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.InTxHash) > 0 {
		i -= len(m.InTxHash)
		copy(dAtA[i:], m.InTxHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.InTxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenericClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovTypes(uint64(m.Epoch))
	}
	if m.ClaimType != 0 {
		n += 1 + sovTypes(uint64(m.ClaimType))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EventClaimer)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Claim != nil {
		n += m.Claim.Size()
	}
	return n
}

func (m *GenericClaim_PriceClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PriceClaim != nil {
		l = m.PriceClaim.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Epoch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovTypes(uint64(m.Nonce))
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Oracle)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Coin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EthAddr)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.MinterId != 0 {
		n += 1 + sovTypes(uint64(m.MinterId))
	}
	if m.EthDecimals != 0 {
		n += 1 + sovTypes(uint64(m.EthDecimals))
	}
	return n
}

func (m *TxStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InTxHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.OutTxHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovTypes(uint64(m.Status))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenericClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenericClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenericClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			m.ClaimType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventClaimer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventClaimer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceClaim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MsgPriceClaim{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Claim = &GenericClaim_PriceClaim{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Epoch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Epoch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Epoch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, &Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Oracle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &MsgPriceClaim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Coin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Coin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Coin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinterId", wireType)
			}
			m.MinterId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinterId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthDecimals", wireType)
			}
			m.EthDecimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthDecimals |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TxStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TxStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= TxStatusType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
