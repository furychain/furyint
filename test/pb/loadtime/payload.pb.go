// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test/loadtime/payload.proto

package loadtime

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Payload is the structure of the loadtime transaction. Proto has a compact
// encoded representation, making it ideal for the loadtime usecase which aims to
// keep the generated transactions small.
type Payload struct {
	Connections uint64    `protobuf:"varint,1,opt,name=connections,proto3" json:"connections,omitempty"`
	Rate        uint64    `protobuf:"varint,2,opt,name=rate,proto3" json:"rate,omitempty"`
	Size_       uint64    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Time        time.Time `protobuf:"bytes,4,opt,name=time,proto3,stdtime" json:"time"`
	Id          []byte    `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	Padding     []byte    `protobuf:"bytes,6,opt,name=padding,proto3" json:"padding,omitempty"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb0d1cb6159a167, []int{0}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return m.Size()
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetConnections() uint64 {
	if m != nil {
		return m.Connections
	}
	return 0
}

func (m *Payload) GetRate() uint64 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *Payload) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *Payload) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *Payload) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Payload) GetPadding() []byte {
	if m != nil {
		return m.Padding
	}
	return nil
}

func init() {
	proto.RegisterType((*Payload)(nil), "loadtime.payload.Payload")
}

func init() { proto.RegisterFile("test/loadtime/payload.proto", fileDescriptor_ecb0d1cb6159a167) }

var fileDescriptor_ecb0d1cb6159a167 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0xc7, 0xe3, 0x7c, 0xf9, 0x5a, 0xe4, 0x22, 0x84, 0x2c, 0x06, 0x2b, 0x48, 0x4e, 0xc4, 0x94,
	0x05, 0x5b, 0x82, 0x85, 0xb9, 0x0b, 0x2b, 0x8a, 0x98, 0xd8, 0x9c, 0xc4, 0x75, 0x2d, 0x35, 0x76,
	0x94, 0x38, 0x43, 0x79, 0x8a, 0xbe, 0x0d, 0xaf, 0xd0, 0xb1, 0x23, 0x13, 0xa0, 0xe4, 0x45, 0x90,
	0x1d, 0x82, 0xd8, 0x7e, 0xf7, 0xbb, 0x3b, 0xfb, 0xaf, 0x83, 0xd7, 0x56, 0x74, 0x96, 0xed, 0x0c,
	0xaf, 0xac, 0xaa, 0x05, 0x6b, 0xf8, 0xde, 0x31, 0x6d, 0x5a, 0x63, 0x0d, 0xba, 0x9c, 0x3d, 0xfd,
	0xf1, 0xf1, 0x95, 0x34, 0xd2, 0xf8, 0x26, 0x73, 0x34, 0xcd, 0xc5, 0x89, 0x34, 0x46, 0xee, 0x04,
	0xf3, 0x55, 0xd1, 0x6f, 0x98, 0xdb, 0xe9, 0x2c, 0xaf, 0x9b, 0x69, 0xe0, 0xe6, 0x0d, 0xc0, 0xe5,
	0xd3, 0xf4, 0x04, 0x4a, 0xe1, 0xaa, 0x34, 0x5a, 0x8b, 0xd2, 0x2a, 0xa3, 0x3b, 0x0c, 0x52, 0x90,
	0x45, 0xf9, 0x5f, 0x85, 0x10, 0x8c, 0x5a, 0x6e, 0x05, 0x0e, 0x7d, 0xcb, 0xb3, 0x73, 0x9d, 0x7a,
	0x15, 0xf8, 0xdf, 0xe4, 0x1c, 0xa3, 0x07, 0x18, 0xb9, 0x8f, 0x70, 0x94, 0x82, 0x6c, 0x75, 0x17,
	0xd3, 0x29, 0x05, 0x9d, 0x53, 0xd0, 0xe7, 0x39, 0xc5, 0xfa, 0xec, 0xf8, 0x91, 0x04, 0x87, 0xcf,
	0x04, 0xe4, 0x7e, 0x03, 0x5d, 0xc0, 0x50, 0x55, 0xf8, 0x7f, 0x0a, 0xb2, 0xf3, 0x3c, 0x54, 0x15,
	0xc2, 0x70, 0xd9, 0xf0, 0xaa, 0x52, 0x5a, 0xe2, 0x85, 0x97, 0x73, 0xb9, 0x7e, 0x3c, 0x0e, 0x04,
	0x9c, 0x06, 0x02, 0xbe, 0x06, 0x02, 0x0e, 0x23, 0x09, 0x4e, 0x23, 0x09, 0xde, 0x47, 0x12, 0xbc,
	0xdc, 0x4a, 0x65, 0xb7, 0x7d, 0x41, 0x4b, 0x53, 0xb3, 0x4d, 0xdf, 0xee, 0xcb, 0x2d, 0x57, 0xda,
	0x93, 0xd2, 0x96, 0xf9, 0xb3, 0x36, 0xc5, 0xef, 0x65, 0x8b, 0x85, 0x8f, 0x75, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x45, 0x7d, 0x77, 0xa9, 0x71, 0x01, 0x00, 0x00,
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Padding) > 0 {
		i -= len(m.Padding)
		copy(dAtA[i:], m.Padding)
		i = encodeVarintPayload(dAtA, i, uint64(len(m.Padding)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPayload(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x2a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPayload(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.Size_ != 0 {
		i = encodeVarintPayload(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x18
	}
	if m.Rate != 0 {
		i = encodeVarintPayload(dAtA, i, uint64(m.Rate))
		i--
		dAtA[i] = 0x10
	}
	if m.Connections != 0 {
		i = encodeVarintPayload(dAtA, i, uint64(m.Connections))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPayload(dAtA []byte, offset int, v uint64) int {
	offset -= sovPayload(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Payload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Connections != 0 {
		n += 1 + sovPayload(uint64(m.Connections))
	}
	if m.Rate != 0 {
		n += 1 + sovPayload(uint64(m.Rate))
	}
	if m.Size_ != 0 {
		n += 1 + sovPayload(uint64(m.Size_))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovPayload(uint64(l))
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	l = len(m.Padding)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	return n
}

func sovPayload(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPayload(x uint64) (n int) {
	return sovPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
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
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Connections", wireType)
			}
			m.Connections = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Connections |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			m.Rate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Padding", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Padding = append(m.Padding[:0], dAtA[iNdEx:postIndex]...)
			if m.Padding == nil {
				m.Padding = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPayload
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
func skipPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayload
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
					return 0, ErrIntOverflowPayload
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
					return 0, ErrIntOverflowPayload
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
				return 0, ErrInvalidLengthPayload
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPayload
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPayload
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPayload        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayload          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPayload = fmt.Errorf("proto: unexpected end of group")
)
