// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cross/core/packets.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
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

type PacketDataIBCSignTx struct {
	TxID    TxID        `protobuf:"bytes,1,opt,name=txID,proto3,casttype=TxID" json:"txID,omitempty"`
	Signers []AccountID `protobuf:"bytes,2,rep,name=signers,proto3,casttype=AccountID" json:"signers,omitempty"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,3,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
}

func (m *PacketDataIBCSignTx) Reset()         { *m = PacketDataIBCSignTx{} }
func (m *PacketDataIBCSignTx) String() string { return proto.CompactTextString(m) }
func (*PacketDataIBCSignTx) ProtoMessage()    {}
func (*PacketDataIBCSignTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f827282d1386ebe9, []int{0}
}
func (m *PacketDataIBCSignTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketDataIBCSignTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketDataIBCSignTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketDataIBCSignTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketDataIBCSignTx.Merge(m, src)
}
func (m *PacketDataIBCSignTx) XXX_Size() int {
	return m.Size()
}
func (m *PacketDataIBCSignTx) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketDataIBCSignTx.DiscardUnknown(m)
}

var xxx_messageInfo_PacketDataIBCSignTx proto.InternalMessageInfo

type PacketAcknowledgementIBCSignTx struct {
	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *PacketAcknowledgementIBCSignTx) Reset()         { *m = PacketAcknowledgementIBCSignTx{} }
func (m *PacketAcknowledgementIBCSignTx) String() string { return proto.CompactTextString(m) }
func (*PacketAcknowledgementIBCSignTx) ProtoMessage()    {}
func (*PacketAcknowledgementIBCSignTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f827282d1386ebe9, []int{1}
}
func (m *PacketAcknowledgementIBCSignTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketAcknowledgementIBCSignTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketAcknowledgementIBCSignTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketAcknowledgementIBCSignTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketAcknowledgementIBCSignTx.Merge(m, src)
}
func (m *PacketAcknowledgementIBCSignTx) XXX_Size() int {
	return m.Size()
}
func (m *PacketAcknowledgementIBCSignTx) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketAcknowledgementIBCSignTx.DiscardUnknown(m)
}

var xxx_messageInfo_PacketAcknowledgementIBCSignTx proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PacketDataIBCSignTx)(nil), "cross.core.PacketDataIBCSignTx")
	proto.RegisterType((*PacketAcknowledgementIBCSignTx)(nil), "cross.core.PacketAcknowledgementIBCSignTx")
}

func init() { proto.RegisterFile("cross/core/packets.proto", fileDescriptor_f827282d1386ebe9) }

var fileDescriptor_f827282d1386ebe9 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbf, 0x6e, 0xd4, 0x40,
	0x10, 0xc6, 0xed, 0xc4, 0x0a, 0xb0, 0xe4, 0x10, 0x98, 0x3f, 0xb2, 0x4e, 0x61, 0x6d, 0xb9, 0xc1,
	0xa2, 0xb0, 0x15, 0xe8, 0x52, 0x20, 0xc5, 0xb8, 0x88, 0x3b, 0x64, 0xae, 0xa2, 0x81, 0xf5, 0x66,
	0x65, 0xaf, 0x62, 0x7b, 0x2d, 0xef, 0x38, 0x38, 0x4f, 0x01, 0x8f, 0xc0, 0xe3, 0x5c, 0x99, 0x92,
	0xca, 0x82, 0xbb, 0x86, 0x3a, 0xe5, 0x55, 0xe8, 0x76, 0x7d, 0xa0, 0xab, 0x76, 0xf6, 0x9b, 0xdf,
	0x8c, 0x46, 0xdf, 0x87, 0x1c, 0xda, 0x09, 0x29, 0x23, 0x2a, 0x3a, 0x16, 0xb5, 0x84, 0x5e, 0x31,
	0x90, 0x61, 0xdb, 0x09, 0x10, 0x36, 0x52, 0x9d, 0x70, 0xdb, 0x99, 0x3f, 0x2b, 0x44, 0x21, 0x94,
	0x1c, 0x6d, 0x2b, 0x4d, 0xcc, 0x5d, 0x9e, 0x53, 0x3d, 0x49, 0x2b, 0xce, 0x1a, 0x88, 0xae, 0x4f,
	0xa7, 0x4a, 0x03, 0xfe, 0xb7, 0x03, 0xf4, 0xf4, 0x83, 0x5a, 0x9a, 0x10, 0x20, 0x69, 0xfc, 0xfe,
	0x23, 0x2f, 0x9a, 0xc5, 0x60, 0x9f, 0x20, 0x0b, 0x86, 0x34, 0x71, 0x4c, 0xcf, 0x0c, 0x8e, 0xe3,
	0xfb, 0x9b, 0xd1, 0xb5, 0x16, 0x43, 0x9a, 0x64, 0x4a, 0xb5, 0x5f, 0xa1, 0x7b, 0x92, 0x17, 0x0d,
	0xeb, 0xa4, 0x73, 0xe0, 0x1d, 0x06, 0xc7, 0xf1, 0x6c, 0x33, 0xba, 0x0f, 0xce, 0x29, 0x15, 0x7d,
	0x03, 0x69, 0x92, 0xed, 0xba, 0xf6, 0x17, 0xf4, 0x08, 0x78, 0xcd, 0x44, 0x0f, 0x9f, 0x4b, 0xc6,
	0x8b, 0x12, 0x9c, 0x43, 0xcf, 0x0c, 0x1e, 0xbe, 0x99, 0x87, 0x3c, 0xa7, 0xea, 0xf0, 0x70, 0x3a,
	0xe7, 0xfa, 0x34, 0xbc, 0x50, 0x44, 0xfc, 0x72, 0x39, 0xba, 0xc6, 0xdd, 0xe8, 0x3e, 0xbf, 0x21,
	0x75, 0x75, 0xe6, 0xef, 0xcf, 0xfb, 0xd9, 0x6c, 0x12, 0x34, 0x6d, 0xa7, 0xe8, 0xc9, 0x8e, 0xd8,
	0xbe, 0x12, 0x48, 0xdd, 0x3a, 0x96, 0x67, 0x06, 0x56, 0x7c, 0x72, 0x37, 0xba, 0xce, 0xfe, 0x92,
	0x7f, 0x88, 0x9f, 0x3d, 0x9e, 0xb4, 0xc5, 0x4e, 0x3a, 0xb3, 0xfe, 0xfc, 0x70, 0x0d, 0xff, 0x1d,
	0xc2, 0xda, 0x90, 0x73, 0x7a, 0xd5, 0x88, 0xaf, 0x15, 0xbb, 0x2c, 0x58, 0xcd, 0x1a, 0xf8, 0xef,
	0xcd, 0x0b, 0x74, 0x24, 0x81, 0x40, 0x2f, 0x95, 0x3b, 0xb3, 0x6c, 0xfa, 0xe9, 0xf9, 0xf8, 0x62,
	0xf9, 0x1b, 0x1b, 0xcb, 0x15, 0x36, 0x6f, 0x57, 0xd8, 0xfc, 0xb5, 0xc2, 0xe6, 0xf7, 0x35, 0x36,
	0x6e, 0xd7, 0xd8, 0xf8, 0xb9, 0xc6, 0xc6, 0xa7, 0xd7, 0x05, 0x87, 0xb2, 0xcf, 0x43, 0x2a, 0xea,
	0xe8, 0x92, 0x00, 0xa1, 0x25, 0xe1, 0x4d, 0x45, 0xf2, 0x48, 0x87, 0x3c, 0xe8, 0xb0, 0xe0, 0xa6,
	0x65, 0x32, 0x3f, 0x52, 0x11, 0xbd, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x49, 0x9f, 0x48, 0xfa,
	0x01, 0x02, 0x00, 0x00,
}

func (m *PacketDataIBCSignTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketDataIBCSignTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketDataIBCSignTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintPackets(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintPackets(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.TxID) > 0 {
		i -= len(m.TxID)
		copy(dAtA[i:], m.TxID)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.TxID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PacketAcknowledgementIBCSignTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketAcknowledgementIBCSignTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketAcknowledgementIBCSignTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintPackets(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPackets(dAtA []byte, offset int, v uint64) int {
	offset -= sovPackets(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PacketDataIBCSignTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxID)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, b := range m.Signers {
			l = len(b)
			n += 1 + l + sovPackets(uint64(l))
		}
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovPackets(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovPackets(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *PacketAcknowledgementIBCSignTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovPackets(uint64(m.Status))
	}
	return n
}

func sovPackets(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPackets(x uint64) (n int) {
	return sovPackets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PacketDataIBCSignTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackets
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
			return fmt.Errorf("proto: PacketDataIBCSignTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketDataIBCSignTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
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
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxID = append(m.TxID[:0], dAtA[iNdEx:postIndex]...)
			if m.TxID == nil {
				m.TxID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
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
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, make([]byte, postIndex-iNdEx))
			copy(m.Signers[len(m.Signers)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
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
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPackets
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPackets
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
func (m *PacketAcknowledgementIBCSignTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackets
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
			return fmt.Errorf("proto: PacketAcknowledgementIBCSignTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketAcknowledgementIBCSignTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPackets
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPackets
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
func skipPackets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPackets
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
					return 0, ErrIntOverflowPackets
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
					return 0, ErrIntOverflowPackets
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
				return 0, ErrInvalidLengthPackets
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPackets
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPackets
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPackets        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPackets          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPackets = fmt.Errorf("proto: unexpected end of group")
)
