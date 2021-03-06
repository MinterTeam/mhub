// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: minter/v1/attestation.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// ClaimType is the cosmos type of an event from the counterpart chain that can
// be handled
type ClaimType int32

const (
	CLAIM_TYPE_UNKNOWN     ClaimType = 0
	CLAIM_TYPE_DEPOSIT     ClaimType = 1
	CLAIM_TYPE_WITHDRAW    ClaimType = 2
	CLAIM_TYPE_VALSET      ClaimType = 3
	CLAIM_TYPE_SEND_TO_ETH ClaimType = 4
)

var ClaimType_name = map[int32]string{
	0: "CLAIM_TYPE_UNKNOWN",
	1: "CLAIM_TYPE_DEPOSIT",
	2: "CLAIM_TYPE_WITHDRAW",
	3: "CLAIM_TYPE_VALSET",
	4: "CLAIM_TYPE_SEND_TO_ETH",
}

var ClaimType_value = map[string]int32{
	"CLAIM_TYPE_UNKNOWN":     0,
	"CLAIM_TYPE_DEPOSIT":     1,
	"CLAIM_TYPE_WITHDRAW":    2,
	"CLAIM_TYPE_VALSET":      3,
	"CLAIM_TYPE_SEND_TO_ETH": 4,
}

func (x ClaimType) String() string {
	return proto.EnumName(ClaimType_name, int32(x))
}

func (ClaimType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dd64cf431ff49b2a, []int{0}
}

// Attestation is an aggregate of `claims` that eventually becomes `observed` by
// all orchestrators
// EVENT_NONCE:
// EventNonce a nonce provided by the peggy contract that is unique per event fired
// These event nonces must be relayed in order. This is a correctness issue,
// if relaying out of order transaction replay attacks become possible
// OBSERVED:
// Observed indicates that >67% of validators have attested to the event,
// and that the event should be executed by the peggy state machine
//
// The actual content of the claims is passed in with the transaction making the claim
// and then passed through the call stack alongside the attestation while it is processed
// the key in which the attestation is stored is keyed on the exact details of the claim
// but there is no reason to store those exact details becuause the next message sender
// will kindly provide you with them.
type Attestation struct {
	EventNonce uint64   `protobuf:"varint,1,opt,name=event_nonce,json=eventNonce,proto3" json:"event_nonce,omitempty"`
	Observed   bool     `protobuf:"varint,2,opt,name=observed,proto3" json:"observed,omitempty"`
	Votes      []string `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes,omitempty"`
}

func (m *Attestation) Reset()         { *m = Attestation{} }
func (m *Attestation) String() string { return proto.CompactTextString(m) }
func (*Attestation) ProtoMessage()    {}
func (*Attestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd64cf431ff49b2a, []int{0}
}
func (m *Attestation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Attestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Attestation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attestation.Merge(m, src)
}
func (m *Attestation) XXX_Size() int {
	return m.Size()
}
func (m *Attestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Attestation.DiscardUnknown(m)
}

var xxx_messageInfo_Attestation proto.InternalMessageInfo

func (m *Attestation) GetEventNonce() uint64 {
	if m != nil {
		return m.EventNonce
	}
	return 0
}

func (m *Attestation) GetObserved() bool {
	if m != nil {
		return m.Observed
	}
	return false
}

func (m *Attestation) GetVotes() []string {
	if m != nil {
		return m.Votes
	}
	return nil
}

// ERC20Token unique identifier for an Ethereum ERC20 token.
// CONTRACT:
// The contract address on ETH of the token (note: developers should look up
// the token symbol using the address on ETH to display for UI)
type MinterCoin struct {
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	CoinId uint64                                 `protobuf:"varint,2,opt,name=coinId,proto3" json:"coinId,omitempty"`
}

func (m *MinterCoin) Reset()         { *m = MinterCoin{} }
func (m *MinterCoin) String() string { return proto.CompactTextString(m) }
func (*MinterCoin) ProtoMessage()    {}
func (*MinterCoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd64cf431ff49b2a, []int{1}
}
func (m *MinterCoin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MinterCoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MinterCoin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MinterCoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinterCoin.Merge(m, src)
}
func (m *MinterCoin) XXX_Size() int {
	return m.Size()
}
func (m *MinterCoin) XXX_DiscardUnknown() {
	xxx_messageInfo_MinterCoin.DiscardUnknown(m)
}

var xxx_messageInfo_MinterCoin proto.InternalMessageInfo

func (m *MinterCoin) GetCoinId() uint64 {
	if m != nil {
		return m.CoinId
	}
	return 0
}

func init() {
	proto.RegisterEnum("minter.v1.ClaimType", ClaimType_name, ClaimType_value)
	proto.RegisterType((*Attestation)(nil), "minter.v1.Attestation")
	proto.RegisterType((*MinterCoin)(nil), "minter.v1.MinterCoin")
}

func init() { proto.RegisterFile("minter/v1/attestation.proto", fileDescriptor_dd64cf431ff49b2a) }

var fileDescriptor_dd64cf431ff49b2a = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x80, 0x93, 0x36, 0x54, 0x8b, 0x77, 0x09, 0x66, 0x94, 0xca, 0x48, 0x69, 0xb4, 0x03, 0xaa,
	0x10, 0xc4, 0x4c, 0x48, 0xdc, 0xb3, 0x36, 0x68, 0x11, 0x5b, 0x3a, 0xa5, 0x86, 0x0a, 0x2e, 0x21,
	0xc9, 0xac, 0x36, 0x62, 0xb1, 0xab, 0xc6, 0x8d, 0xd8, 0x3f, 0x80, 0x9c, 0xf8, 0x03, 0x39, 0xf1,
	0x67, 0x76, 0xdc, 0x11, 0x71, 0x98, 0x50, 0xfb, 0x47, 0xd0, 0x92, 0x08, 0x22, 0x5a, 0x4e, 0xf6,
	0xfb, 0xbe, 0xf7, 0xac, 0xa7, 0xe7, 0x07, 0x1e, 0x27, 0x31, 0x13, 0x74, 0x89, 0xb3, 0x23, 0x1c,
	0x08, 0x41, 0x53, 0x11, 0x88, 0x98, 0x33, 0x73, 0xb1, 0xe4, 0x82, 0x43, 0xb5, 0x92, 0x66, 0x76,
	0x84, 0x0e, 0x66, 0x7c, 0xc6, 0x4b, 0x8a, 0xef, 0x6e, 0x55, 0xc2, 0xe1, 0x47, 0xb0, 0x6f, 0xfd,
	0xad, 0x82, 0x7d, 0xb0, 0x4f, 0x33, 0xca, 0x84, 0xcf, 0x38, 0x8b, 0x68, 0x4f, 0x36, 0xe4, 0x81,
	0xe2, 0x81, 0x12, 0xb9, 0x77, 0x04, 0x22, 0xb0, 0xc7, 0xc3, 0x94, 0x2e, 0x33, 0x7a, 0xd1, 0x6b,
	0x19, 0xf2, 0x60, 0xcf, 0xfb, 0x13, 0xc3, 0x03, 0x70, 0x2f, 0xe3, 0x82, 0xa6, 0xbd, 0xb6, 0xd1,
	0x1e, 0xa8, 0x5e, 0x15, 0x1c, 0x5e, 0x02, 0x70, 0x56, 0x36, 0x31, 0xe4, 0x31, 0x83, 0xaf, 0x41,
	0x27, 0x48, 0xf8, 0x8a, 0x89, 0xf2, 0x6d, 0xf5, 0xd8, 0xbc, 0xbe, 0xed, 0x4b, 0x3f, 0x6f, 0xfb,
	0x4f, 0x66, 0xb1, 0x98, 0xaf, 0x42, 0x33, 0xe2, 0x09, 0x8e, 0x78, 0x9a, 0xf0, 0xb4, 0x3e, 0x9e,
	0xa7, 0x17, 0x9f, 0xb0, 0xb8, 0x5a, 0xd0, 0xd4, 0x74, 0x98, 0xf0, 0xea, 0x6a, 0xd8, 0x05, 0x9d,
	0x88, 0xc7, 0xcc, 0xa9, 0xba, 0x50, 0xbc, 0x3a, 0x7a, 0xfa, 0xb5, 0x05, 0xd4, 0xe1, 0x65, 0x10,
	0x27, 0xe4, 0x6a, 0x41, 0xa1, 0x09, 0xe0, 0xf0, 0xd4, 0x72, 0xce, 0x7c, 0xf2, 0xfe, 0xdc, 0xf6,
	0xdf, 0xba, 0x6f, 0xdc, 0xf1, 0xd4, 0xd5, 0x24, 0xd4, 0xcd, 0x0b, 0x63, 0x87, 0xf9, 0x27, 0x7f,
	0x64, 0x9f, 0x8f, 0x27, 0x0e, 0xd1, 0xe4, 0xad, 0xfc, 0xda, 0xc0, 0x17, 0xe0, 0x41, 0x83, 0x4e,
	0x1d, 0x72, 0x32, 0xf2, 0xac, 0xa9, 0xd6, 0x42, 0x8f, 0xf2, 0xc2, 0xd8, 0xa5, 0xe0, 0x33, 0x70,
	0xbf, 0x81, 0xdf, 0x59, 0xa7, 0x13, 0x9b, 0x68, 0x6d, 0xf4, 0x30, 0x2f, 0x8c, 0x6d, 0x01, 0x5f,
	0x81, 0x6e, 0x03, 0x4e, 0x6c, 0x77, 0xe4, 0x93, 0xb1, 0x6f, 0x93, 0x13, 0x4d, 0x41, 0x28, 0x2f,
	0x8c, 0xff, 0x58, 0xa4, 0x7c, 0xf9, 0xae, 0x4b, 0xc7, 0xce, 0xf5, 0x5a, 0x97, 0x6f, 0xd6, 0xba,
	0xfc, 0x6b, 0xad, 0xcb, 0xdf, 0x36, 0xba, 0x74, 0xb3, 0xd1, 0xa5, 0x1f, 0x1b, 0x5d, 0xfa, 0x80,
	0x1b, 0xd3, 0xae, 0x3e, 0x87, 0xd0, 0x20, 0xc1, 0xc9, 0x7c, 0x15, 0xe2, 0x68, 0x1e, 0xc4, 0x0c,
	0x7f, 0xc6, 0xf5, 0x5a, 0x95, 0xa3, 0x0f, 0x3b, 0xe5, 0xb6, 0xbc, 0xfc, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x43, 0x53, 0x02, 0x76, 0x6d, 0x02, 0x00, 0x00,
}

func (m *Attestation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attestation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Attestation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Votes[iNdEx])
			copy(dAtA[i:], m.Votes[iNdEx])
			i = encodeVarintAttestation(dAtA, i, uint64(len(m.Votes[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Observed {
		i--
		if m.Observed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.EventNonce != 0 {
		i = encodeVarintAttestation(dAtA, i, uint64(m.EventNonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MinterCoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MinterCoin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MinterCoin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CoinId != 0 {
		i = encodeVarintAttestation(dAtA, i, uint64(m.CoinId))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAttestation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAttestation(dAtA []byte, offset int, v uint64) int {
	offset -= sovAttestation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Attestation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EventNonce != 0 {
		n += 1 + sovAttestation(uint64(m.EventNonce))
	}
	if m.Observed {
		n += 2
	}
	if len(m.Votes) > 0 {
		for _, s := range m.Votes {
			l = len(s)
			n += 1 + l + sovAttestation(uint64(l))
		}
	}
	return n
}

func (m *MinterCoin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovAttestation(uint64(l))
	if m.CoinId != 0 {
		n += 1 + sovAttestation(uint64(m.CoinId))
	}
	return n
}

func sovAttestation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttestation(x uint64) (n int) {
	return sovAttestation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Attestation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: Attestation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attestation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventNonce", wireType)
			}
			m.EventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Observed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Observed = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAttestation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAttestation
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
func (m *MinterCoin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: MinterCoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MinterCoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinId", wireType)
			}
			m.CoinId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoinId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAttestation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAttestation
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
func skipAttestation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
				return 0, ErrInvalidLengthAttestation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAttestation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAttestation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAttestation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttestation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAttestation = fmt.Errorf("proto: unexpected end of group")
)
