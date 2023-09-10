// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/epochstorage/stake_entry.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type StakeEntry struct {
	Stake              types.Coin `protobuf:"bytes,1,opt,name=stake,proto3" json:"stake"`
	Address            string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	StakeAppliedBlock  uint64     `protobuf:"varint,3,opt,name=stake_applied_block,json=stakeAppliedBlock,proto3" json:"stake_applied_block,omitempty"`
	Endpoints          []Endpoint `protobuf:"bytes,4,rep,name=endpoints,proto3" json:"endpoints"`
	Geolocation        uint64     `protobuf:"varint,5,opt,name=geolocation,proto3" json:"geolocation,omitempty"`
	Chain              string     `protobuf:"bytes,6,opt,name=chain,proto3" json:"chain,omitempty"`
	Moniker            string     `protobuf:"bytes,8,opt,name=moniker,proto3" json:"moniker,omitempty"`
	DelegateTotal      types.Coin `protobuf:"bytes,9,opt,name=delegate_total,json=delegateTotal,proto3" json:"delegate_total"`
	DelegateLimit      types.Coin `protobuf:"bytes,10,opt,name=delegate_limit,json=delegateLimit,proto3" json:"delegate_limit"`
	DelegateCommission uint64     `protobuf:"varint,11,opt,name=delegate_commission,json=delegateCommission,proto3" json:"delegate_commission,omitempty"`
}

func (m *StakeEntry) Reset()         { *m = StakeEntry{} }
func (m *StakeEntry) String() string { return proto.CompactTextString(m) }
func (*StakeEntry) ProtoMessage()    {}
func (*StakeEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_df6302d6b53c056e, []int{0}
}
func (m *StakeEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakeEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakeEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakeEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeEntry.Merge(m, src)
}
func (m *StakeEntry) XXX_Size() int {
	return m.Size()
}
func (m *StakeEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeEntry.DiscardUnknown(m)
}

var xxx_messageInfo_StakeEntry proto.InternalMessageInfo

func (m *StakeEntry) GetStake() types.Coin {
	if m != nil {
		return m.Stake
	}
	return types.Coin{}
}

func (m *StakeEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StakeEntry) GetStakeAppliedBlock() uint64 {
	if m != nil {
		return m.StakeAppliedBlock
	}
	return 0
}

func (m *StakeEntry) GetEndpoints() []Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *StakeEntry) GetGeolocation() int32 {
	if m != nil {
		return m.Geolocation
	}
	return 0
}

func (m *StakeEntry) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *StakeEntry) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *StakeEntry) GetDelegateTotal() types.Coin {
	if m != nil {
		return m.DelegateTotal
	}
	return types.Coin{}
}

func (m *StakeEntry) GetDelegateLimit() types.Coin {
	if m != nil {
		return m.DelegateLimit
	}
	return types.Coin{}
}

func (m *StakeEntry) GetDelegateCommission() uint64 {
	if m != nil {
		return m.DelegateCommission
	}
	return 0
}

func init() {
	proto.RegisterType((*StakeEntry)(nil), "lavanet.lava.epochstorage.StakeEntry")
}

func init() {
	proto.RegisterFile("lavanet/lava/epochstorage/stake_entry.proto", fileDescriptor_df6302d6b53c056e)
}

var fileDescriptor_df6302d6b53c056e = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0xcd, 0x92, 0xa4, 0x4d, 0x1c, 0x81, 0xc0, 0xed, 0xc1, 0xed, 0x61, 0x59, 0xc1, 0x65, 0x25,
	0x90, 0xad, 0x16, 0xf1, 0x01, 0xa4, 0x6a, 0x91, 0x10, 0xa7, 0xc0, 0x89, 0x4b, 0xe4, 0xf5, 0x8e,
	0x36, 0x56, 0x76, 0x3d, 0xab, 0xb5, 0xa9, 0xe8, 0x5f, 0xf0, 0x59, 0x3d, 0xf6, 0xc8, 0x09, 0xa1,
	0xe4, 0x03, 0xf8, 0x05, 0xe4, 0xf5, 0x6e, 0x69, 0x0e, 0x91, 0xe0, 0x64, 0x8f, 0xdf, 0xbc, 0xe7,
	0xf7, 0x46, 0x43, 0x5e, 0x95, 0xf2, 0x5a, 0x1a, 0x70, 0xc2, 0x9f, 0x02, 0x6a, 0x54, 0x2b, 0xeb,
	0xb0, 0x91, 0x05, 0x08, 0xeb, 0xe4, 0x1a, 0x96, 0x60, 0x5c, 0x73, 0xc3, 0xeb, 0x06, 0x1d, 0xd2,
	0x93, 0xae, 0x99, 0xfb, 0x93, 0x3f, 0x6c, 0x3e, 0x4d, 0xf7, 0xeb, 0x80, 0xc9, 0x6b, 0xd4, 0xc6,
	0x05, 0x91, 0xd3, 0xe3, 0x02, 0x0b, 0x6c, 0xaf, 0xc2, 0xdf, 0xba, 0xd7, 0x58, 0xa1, 0xad, 0xd0,
	0x8a, 0x4c, 0x5a, 0x10, 0xd7, 0x67, 0x19, 0x38, 0x79, 0x26, 0x14, 0x6a, 0x13, 0xf0, 0x17, 0xbf,
	0x87, 0x84, 0x7c, 0xf2, 0x86, 0x2e, 0xbd, 0x1f, 0xfa, 0x96, 0x8c, 0x5b, 0x7b, 0x2c, 0x4a, 0xa2,
	0x74, 0x76, 0x7e, 0xc2, 0x03, 0x9d, 0x7b, 0x3a, 0xef, 0xe8, 0xfc, 0x02, 0xb5, 0x99, 0x8f, 0x6e,
	0x7f, 0x3e, 0x1f, 0x2c, 0x42, 0x37, 0x65, 0xe4, 0x50, 0xe6, 0x79, 0x03, 0xd6, 0xb2, 0x47, 0x49,
	0x94, 0x4e, 0x17, 0x7d, 0x49, 0x39, 0x39, 0x0a, 0x79, 0x65, 0x5d, 0x97, 0x1a, 0xf2, 0x65, 0x56,
	0xa2, 0x5a, 0xb3, 0x61, 0x12, 0xa5, 0xa3, 0xc5, 0xb3, 0x16, 0x7a, 0x17, 0x90, 0xb9, 0x07, 0xe8,
	0x7b, 0x32, 0xed, 0x73, 0x59, 0x36, 0x4a, 0x86, 0xe9, 0xec, 0xfc, 0x25, 0xdf, 0x3b, 0x1e, 0x7e,
	0xd9, 0xf5, 0x76, 0x76, 0xfe, 0x72, 0x69, 0x42, 0x66, 0x05, 0x60, 0x89, 0x4a, 0x3a, 0x8d, 0x86,
	0x8d, 0xdb, 0x0f, 0x1f, 0x3e, 0xd1, 0x63, 0x32, 0x56, 0x2b, 0xa9, 0x0d, 0x3b, 0x68, 0x2d, 0x87,
	0xc2, 0x47, 0xa9, 0xd0, 0xe8, 0x35, 0x34, 0x6c, 0x12, 0xa2, 0x74, 0x25, 0xbd, 0x22, 0x4f, 0x72,
	0x28, 0xa1, 0x90, 0x0e, 0x96, 0x0e, 0x9d, 0x2c, 0xd9, 0xf4, 0xdf, 0x86, 0xf4, 0xb8, 0xa7, 0x7d,
	0xf6, 0xac, 0x1d, 0x9d, 0x52, 0x57, 0xda, 0x31, 0xf2, 0x9f, 0x3a, 0x1f, 0x3d, 0x8b, 0x0a, 0x72,
	0x74, 0xaf, 0xa3, 0xb0, 0xaa, 0xb4, 0xb5, 0x3e, 0xe9, 0xac, 0x4d, 0x4a, 0x7b, 0xe8, 0xe2, 0x1e,
	0xf9, 0x30, 0x9a, 0x1c, 0x3e, 0x9d, 0xcc, 0xaf, 0x6e, 0x37, 0x71, 0x74, 0xb7, 0x89, 0xa3, 0x5f,
	0x9b, 0x38, 0xfa, 0xbe, 0x8d, 0x07, 0x77, 0xdb, 0x78, 0xf0, 0x63, 0x1b, 0x0f, 0xbe, 0xbc, 0x2e,
	0xb4, 0x5b, 0x7d, 0xcd, 0xb8, 0xc2, 0x4a, 0xec, 0xac, 0xdd, 0xb7, 0xdd, 0xc5, 0x73, 0x37, 0x35,
	0xd8, 0xec, 0xa0, 0x5d, 0xa0, 0x37, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x69, 0xb8, 0xeb,
	0xea, 0x02, 0x00, 0x00,
}

func (m *StakeEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakeEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DelegateCommission != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.DelegateCommission))
		i--
		dAtA[i] = 0x58
	}
	{
		size, err := m.DelegateLimit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakeEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size, err := m.DelegateTotal.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakeEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x32
	}
	if m.Geolocation != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.Geolocation))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Endpoints) > 0 {
		for iNdEx := len(m.Endpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Endpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStakeEntry(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.StakeAppliedBlock != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.StakeAppliedBlock))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Stake.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakeEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintStakeEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovStakeEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakeEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Stake.Size()
	n += 1 + l + sovStakeEntry(uint64(l))
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	if m.StakeAppliedBlock != 0 {
		n += 1 + sovStakeEntry(uint64(m.StakeAppliedBlock))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovStakeEntry(uint64(l))
		}
	}
	if m.Geolocation != 0 {
		n += 1 + sovStakeEntry(uint64(m.Geolocation))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	l = m.DelegateTotal.Size()
	n += 1 + l + sovStakeEntry(uint64(l))
	l = m.DelegateLimit.Size()
	n += 1 + l + sovStakeEntry(uint64(l))
	if m.DelegateCommission != 0 {
		n += 1 + sovStakeEntry(uint64(m.DelegateCommission))
	}
	return n
}

func sovStakeEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStakeEntry(x uint64) (n int) {
	return sovStakeEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakeEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStakeEntry
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
			return fmt.Errorf("proto: StakeEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakeEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeAppliedBlock", wireType)
			}
			m.StakeAppliedBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StakeAppliedBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, Endpoint{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Geolocation", wireType)
			}
			m.Geolocation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Geolocation |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateTotal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegateTotal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegateLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateCommission", wireType)
			}
			m.DelegateCommission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegateCommission |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStakeEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStakeEntry
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
func skipStakeEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStakeEntry
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
					return 0, ErrIntOverflowStakeEntry
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
					return 0, ErrIntOverflowStakeEntry
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
				return 0, ErrInvalidLengthStakeEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStakeEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStakeEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStakeEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStakeEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStakeEntry = fmt.Errorf("proto: unexpected end of group")
)
