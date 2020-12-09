// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c.proto

package importing

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/frankee/protobuf/gogoproto"
	proto "github.com/frankee/protobuf/proto"
	_ "github.com/frankee/protobuf/test/importcustom-issue389/imported"
	github_com_gogo_protobuf_test_importcustom_issue389_imported "github.com/frankee/protobuf/test/importcustom-issue389/imported"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type C struct {
	F2                   *github_com_gogo_protobuf_test_importcustom_issue389_imported.B `protobuf:"bytes,1,opt,name=f2,proto3,customtype=github.com/frankee/protobuf/test/importcustom-issue389/imported.B" json:"f2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *C) Reset()         { *m = C{} }
func (m *C) String() string { return proto.CompactTextString(m) }
func (*C) ProtoMessage()    {}
func (*C) Descriptor() ([]byte, []int) {
	return fileDescriptor_7086c139101e99ef, []int{0}
}
func (m *C) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_C.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C.Merge(m, src)
}
func (m *C) XXX_Size() int {
	return m.Size()
}
func (m *C) XXX_DiscardUnknown() {
	xxx_messageInfo_C.DiscardUnknown(m)
}

var xxx_messageInfo_C proto.InternalMessageInfo

func init() {
	proto.RegisterType((*C)(nil), "importing.C")
}

func init() { proto.RegisterFile("c.proto", fileDescriptor_7086c139101e99ef) }

var fileDescriptor_7086c139101e99ef = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x62, 0x4f, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0xcc, 0x2d, 0xc8, 0x2f, 0x2a, 0xc9, 0xcc, 0x4b, 0x97, 0xd2,
	0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7,
	0x07, 0xab, 0x48, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x53, 0xca, 0x05, 0xa7,
	0xf2, 0x92, 0xd4, 0xe2, 0x12, 0x7d, 0x88, 0xb9, 0xc9, 0xa5, 0xc5, 0x25, 0xf9, 0xb9, 0xba, 0x99,
	0xc5, 0xc5, 0xa5, 0xa9, 0xc6, 0x16, 0x96, 0x50, 0xd1, 0xd4, 0x14, 0xfd, 0x44, 0x88, 0x29, 0x4a,
	0x29, 0x5c, 0x8c, 0xce, 0x42, 0xf1, 0x5c, 0x4c, 0x69, 0x46, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0xdc, 0x7a, 0x30, 0x35, 0x7a, 0x8e, 0x4e, 0x4e, 0xb7, 0xee, 0xc9, 0xdb, 0x51, 0x62, 0x8f,
	0x9e, 0x53, 0x10, 0x53, 0x9a, 0x91, 0x93, 0xc4, 0x8f, 0x87, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31,
	0xee, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x26, 0xb1, 0x81, 0x4d, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x06, 0x70, 0x86, 0x94,
	0x11, 0x01, 0x00, 0x00,
}

func (this *C) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*C)
	if !ok {
		that2, ok := that.(C)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.F2 == nil {
		if this.F2 != nil {
			return false
		}
	} else if !this.F2.Equal(*that1.F2) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *C) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *C) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.F2 != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintC(dAtA, i, uint64(m.F2.Size()))
		n1, err1 := m.F2.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintC(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedC(r randyC, easy bool) *C {
	this := &C{}
	if r.Intn(10) != 0 {
		this.F2 = github_com_gogo_protobuf_test_importcustom_issue389_imported.NewPopulatedB(r)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedC(r, 2)
	}
	return this
}

type randyC interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneC(r randyC) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringC(r randyC) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneC(r)
	}
	return string(tmps)
}
func randUnrecognizedC(r randyC, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldC(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldC(dAtA []byte, r randyC, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateC(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateC(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateC(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateC(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateC(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateC(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateC(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *C) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.F2 != nil {
		l = m.F2.Size()
		n += 1 + l + sovC(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovC(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozC(x uint64) (n int) {
	return sovC(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *C) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowC
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
			return fmt.Errorf("proto: C: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: C: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field F2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowC
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
				return ErrInvalidLengthC
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthC
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.F2 == nil {
				m.F2 = &github_com_gogo_protobuf_test_importcustom_issue389_imported.B{}
			}
			if err := m.F2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipC(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthC
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthC
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipC(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowC
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
					return 0, ErrIntOverflowC
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
					return 0, ErrIntOverflowC
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
				return 0, ErrInvalidLengthC
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthC
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowC
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
				next, err := skipC(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthC
				}
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
	ErrInvalidLengthC = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowC   = fmt.Errorf("proto: integer overflow")
)
