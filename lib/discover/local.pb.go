// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: local.proto

/*
	Package discover is a generated protocol buffer package.

	It is generated from these files:
		local.proto

	It has these top-level messages:
		Announce
*/
package discover

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_syncthing_syncthing_lib_protocol "github.com/syncthing/syncthing/lib/protocol"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Announce struct {
	ID         github_com_syncthing_syncthing_lib_protocol.DeviceID `protobuf:"bytes,1,opt,name=id,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.DeviceID" json:"id"`
	Addresses  []string                                             `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty"`
	InstanceID int64                                                `protobuf:"varint,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (m *Announce) Reset()                    { *m = Announce{} }
func (m *Announce) String() string            { return proto.CompactTextString(m) }
func (*Announce) ProtoMessage()               {}
func (*Announce) Descriptor() ([]byte, []int) { return fileDescriptorLocal, []int{0} }

func init() {
	proto.RegisterType((*Announce)(nil), "discover.Announce")
}
func (m *Announce) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Announce) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintLocal(dAtA, i, uint64(m.ID.ProtoSize()))
	n1, err := m.ID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.InstanceID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLocal(dAtA, i, uint64(m.InstanceID))
	}
	return i, nil
}

func encodeFixed64Local(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Local(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLocal(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Announce) ProtoSize() (n int) {
	var l int
	_ = l
	l = m.ID.ProtoSize()
	n += 1 + l + sovLocal(uint64(l))
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovLocal(uint64(l))
		}
	}
	if m.InstanceID != 0 {
		n += 1 + sovLocal(uint64(m.InstanceID))
	}
	return n
}

func sovLocal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLocal(x uint64) (n int) {
	return sovLocal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Announce) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Announce: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Announce: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLocal
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLocal
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstanceID", wireType)
			}
			m.InstanceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InstanceID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLocal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLocal
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
func skipLocal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocal
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
					return 0, ErrIntOverflowLocal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLocal
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLocal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLocal
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLocal(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLocal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocal   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("local.proto", fileDescriptorLocal) }

var fileDescriptorLocal = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x4f, 0x4e, 0x84, 0x30,
	0x14, 0xc6, 0x29, 0x24, 0x66, 0xa6, 0x63, 0x5c, 0x10, 0x17, 0xc4, 0x98, 0x42, 0x5c, 0xb1, 0x11,
	0x16, 0x7a, 0x01, 0x09, 0x9b, 0x6e, 0xb9, 0x80, 0x81, 0xb6, 0x32, 0x2f, 0xc1, 0x3e, 0x43, 0x61,
	0x12, 0x6f, 0xe3, 0x05, 0xbc, 0x07, 0x4b, 0xd7, 0x2e, 0x1a, 0xad, 0x17, 0x31, 0xe9, 0x68, 0x86,
	0xdd, 0xf7, 0xfd, 0xf2, 0x7b, 0x7f, 0xe8, 0x6e, 0x40, 0xd1, 0x0e, 0xc5, 0xcb, 0x88, 0x13, 0xc6,
	0x1b, 0x09, 0x46, 0xe0, 0x41, 0x8d, 0x57, 0xb7, 0x3d, 0x4c, 0xfb, 0xb9, 0x2b, 0x04, 0x3e, 0x97,
	0x3d, 0xf6, 0x58, 0x7a, 0xa1, 0x9b, 0x9f, 0x7c, 0xf3, 0xc5, 0xa7, 0xe3, 0xe0, 0xcd, 0x3b, 0xa1,
	0x9b, 0x07, 0xad, 0x71, 0xd6, 0x42, 0xc5, 0x0d, 0x0d, 0x41, 0x26, 0x24, 0x23, 0xf9, 0x79, 0x55,
	0x2d, 0x36, 0x0d, 0x3e, 0x6d, 0x7a, 0xbf, 0xda, 0x67, 0x5e, 0xb5, 0x98, 0xf6, 0xa0, 0xfb, 0x55,
	0x1a, 0xa0, 0x3b, 0x9e, 0x10, 0x38, 0x14, 0xb5, 0x3a, 0x80, 0x50, 0xbc, 0x76, 0x36, 0x0d, 0x79,
	0xdd, 0x84, 0x20, 0xe3, 0x6b, 0xba, 0x6d, 0xa5, 0x1c, 0x95, 0x31, 0xca, 0x24, 0x61, 0x16, 0xe5,
	0xdb, 0xe6, 0x04, 0xe2, 0x92, 0xee, 0x40, 0x9b, 0xa9, 0xd5, 0x42, 0x3d, 0x82, 0x4c, 0xa2, 0x8c,
	0xe4, 0x51, 0x75, 0xe1, 0x6c, 0x4a, 0xf9, 0x1f, 0xe6, 0x75, 0x43, 0xff, 0x15, 0x2e, 0xab, 0xcb,
	0xe5, 0x9b, 0x05, 0x8b, 0x63, 0xe4, 0xc3, 0x31, 0xf2, 0xe5, 0x58, 0xf0, 0xf6, 0xc3, 0x48, 0x77,
	0xe6, 0x3f, 0xb8, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x46, 0x4f, 0x13, 0x14, 0x01, 0x00,
	0x00,
}
