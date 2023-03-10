// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: leaderboard/player_info.proto

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

type PlayerInfo struct {
	PlayerID  string `protobuf:"bytes,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	Score     uint64 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	DateAdded string `protobuf:"bytes,3,opt,name=dateAdded,proto3" json:"dateAdded,omitempty"`
	GameId    string `protobuf:"bytes,4,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

func (m *PlayerInfo) Reset()         { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()    {}
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2746532f25366801, []int{0}
}
func (m *PlayerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlayerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlayerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlayerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerInfo.Merge(m, src)
}
func (m *PlayerInfo) XXX_Size() int {
	return m.Size()
}
func (m *PlayerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerInfo proto.InternalMessageInfo

func (m *PlayerInfo) GetPlayerID() string {
	if m != nil {
		return m.PlayerID
	}
	return ""
}

func (m *PlayerInfo) GetScore() uint64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *PlayerInfo) GetDateAdded() string {
	if m != nil {
		return m.DateAdded
	}
	return ""
}

func (m *PlayerInfo) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func init() {
	proto.RegisterType((*PlayerInfo)(nil), "cosmonaut.leaderboard.leaderboard.PlayerInfo")
}

func init() { proto.RegisterFile("leaderboard/player_info.proto", fileDescriptor_2746532f25366801) }

var fileDescriptor_2746532f25366801 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x49, 0x4d, 0x4c,
	0x49, 0x2d, 0x4a, 0xca, 0x4f, 0x2c, 0x4a, 0xd1, 0x2f, 0xc8, 0x49, 0xac, 0x4c, 0x2d, 0x8a, 0xcf,
	0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4c, 0xce, 0x2f, 0xce, 0xcd,
	0xcf, 0x4b, 0x2c, 0x2d, 0xd1, 0x43, 0x52, 0x88, 0xcc, 0x96, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0xab, 0xd6, 0x07, 0xb1, 0x20, 0x1a, 0x95, 0x4a, 0xb8, 0xb8, 0x02, 0xc0, 0xa6, 0x79, 0xe6, 0xa5,
	0xe5, 0x0b, 0x49, 0x71, 0x71, 0x40, 0xcc, 0xf6, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x82, 0xf3, 0x85, 0x44, 0xb8, 0x58, 0x8b, 0x93, 0xf3, 0x8b, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35,
	0x58, 0x82, 0x20, 0x1c, 0x21, 0x19, 0x2e, 0xce, 0x94, 0xc4, 0x92, 0x54, 0xc7, 0x94, 0x94, 0xd4,
	0x14, 0x09, 0x66, 0xb0, 0x16, 0x84, 0x80, 0x90, 0x18, 0x17, 0x5b, 0x7a, 0x62, 0x6e, 0xaa, 0x67,
	0x8a, 0x04, 0x0b, 0x58, 0x0a, 0xca, 0x73, 0xf2, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0x93, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xb8,
	0x9f, 0xf4, 0x91, 0x3d, 0x5f, 0x81, 0xc2, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x7b,
	0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xdc, 0xd3, 0x66, 0x26, 0x01, 0x00, 0x00,
}

func (m *PlayerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlayerInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlayerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GameId) > 0 {
		i -= len(m.GameId)
		copy(dAtA[i:], m.GameId)
		i = encodeVarintPlayerInfo(dAtA, i, uint64(len(m.GameId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DateAdded) > 0 {
		i -= len(m.DateAdded)
		copy(dAtA[i:], m.DateAdded)
		i = encodeVarintPlayerInfo(dAtA, i, uint64(len(m.DateAdded)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Score != 0 {
		i = encodeVarintPlayerInfo(dAtA, i, uint64(m.Score))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PlayerID) > 0 {
		i -= len(m.PlayerID)
		copy(dAtA[i:], m.PlayerID)
		i = encodeVarintPlayerInfo(dAtA, i, uint64(len(m.PlayerID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlayerInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlayerInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PlayerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PlayerID)
	if l > 0 {
		n += 1 + l + sovPlayerInfo(uint64(l))
	}
	if m.Score != 0 {
		n += 1 + sovPlayerInfo(uint64(m.Score))
	}
	l = len(m.DateAdded)
	if l > 0 {
		n += 1 + l + sovPlayerInfo(uint64(l))
	}
	l = len(m.GameId)
	if l > 0 {
		n += 1 + l + sovPlayerInfo(uint64(l))
	}
	return n
}

func sovPlayerInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlayerInfo(x uint64) (n int) {
	return sovPlayerInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlayerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerInfo
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
			return fmt.Errorf("proto: PlayerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlayerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerInfo
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
				return ErrInvalidLengthPlayerInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateAdded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerInfo
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
				return ErrInvalidLengthPlayerInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DateAdded = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerInfo
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
				return ErrInvalidLengthPlayerInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GameId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlayerInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlayerInfo
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
func skipPlayerInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlayerInfo
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
					return 0, ErrIntOverflowPlayerInfo
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
					return 0, ErrIntOverflowPlayerInfo
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
				return 0, ErrInvalidLengthPlayerInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlayerInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlayerInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlayerInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlayerInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlayerInfo = fmt.Errorf("proto: unexpected end of group")
)
