// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imports/test_import_all.proto

package imports

import (
	fmt "fmt"
	proto "github.com/frankee/protobuf/proto"
	fmt1 "github.com/frankee/protobuf/protoc-gen-gogo/testdata/imports/fmt"
	test_a_1 "github.com/frankee/protobuf/protoc-gen-gogo/testdata/imports/test_a_1"
	test_a_2 "github.com/frankee/protobuf/protoc-gen-gogo/testdata/imports/test_a_2"
	test_b_1 "github.com/frankee/protobuf/protoc-gen-gogo/testdata/imports/test_b_1"
	math "math"
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

type All struct {
	Am1                  *test_a_1.M1 `protobuf:"bytes,1,opt,name=am1,proto3" json:"am1,omitempty"`
	Am2                  *test_a_1.M2 `protobuf:"bytes,2,opt,name=am2,proto3" json:"am2,omitempty"`
	Am3                  *test_a_2.M3 `protobuf:"bytes,3,opt,name=am3,proto3" json:"am3,omitempty"`
	Am4                  *test_a_2.M4 `protobuf:"bytes,4,opt,name=am4,proto3" json:"am4,omitempty"`
	Bm1                  *test_b_1.M1 `protobuf:"bytes,5,opt,name=bm1,proto3" json:"bm1,omitempty"`
	Bm2                  *test_b_1.M2 `protobuf:"bytes,6,opt,name=bm2,proto3" json:"bm2,omitempty"`
	Fmt                  *fmt1.M      `protobuf:"bytes,7,opt,name=fmt,proto3" json:"fmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *All) Reset()         { *m = All{} }
func (m *All) String() string { return proto.CompactTextString(m) }
func (*All) ProtoMessage()    {}
func (*All) Descriptor() ([]byte, []int) {
	return fileDescriptor_324466f0afc16f77, []int{0}
}
func (m *All) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_All.Unmarshal(m, b)
}
func (m *All) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_All.Marshal(b, m, deterministic)
}
func (m *All) XXX_Merge(src proto.Message) {
	xxx_messageInfo_All.Merge(m, src)
}
func (m *All) XXX_Size() int {
	return xxx_messageInfo_All.Size(m)
}
func (m *All) XXX_DiscardUnknown() {
	xxx_messageInfo_All.DiscardUnknown(m)
}

var xxx_messageInfo_All proto.InternalMessageInfo

func (m *All) GetAm1() *test_a_1.M1 {
	if m != nil {
		return m.Am1
	}
	return nil
}

func (m *All) GetAm2() *test_a_1.M2 {
	if m != nil {
		return m.Am2
	}
	return nil
}

func (m *All) GetAm3() *test_a_2.M3 {
	if m != nil {
		return m.Am3
	}
	return nil
}

func (m *All) GetAm4() *test_a_2.M4 {
	if m != nil {
		return m.Am4
	}
	return nil
}

func (m *All) GetBm1() *test_b_1.M1 {
	if m != nil {
		return m.Bm1
	}
	return nil
}

func (m *All) GetBm2() *test_b_1.M2 {
	if m != nil {
		return m.Bm2
	}
	return nil
}

func (m *All) GetFmt() *fmt1.M {
	if m != nil {
		return m.Fmt
	}
	return nil
}

func init() {
	proto.RegisterType((*All)(nil), "test.All")
}

func init() { proto.RegisterFile("imports/test_import_all.proto", fileDescriptor_324466f0afc16f77) }

var fileDescriptor_324466f0afc16f77 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x06, 0x60, 0xe5, 0x73, 0xbf, 0x20, 0x99, 0x05, 0x85, 0xc5, 0x20, 0x90, 0x50, 0x27, 0x96,
	0xda, 0xb2, 0x9d, 0x11, 0x21, 0xc1, 0xde, 0xa5, 0x23, 0x4b, 0x64, 0x97, 0x38, 0x54, 0xe4, 0xea,
	0x28, 0xbd, 0xfe, 0x5e, 0xfe, 0x0a, 0xb2, 0x0f, 0x24, 0x08, 0xcd, 0x96, 0xbc, 0xcf, 0x6b, 0x9d,
	0x7d, 0xfc, 0x76, 0x07, 0x43, 0x1c, 0xf1, 0xa0, 0xb0, 0x3d, 0x60, 0x43, 0x3f, 0x8d, 0xeb, 0x7b,
	0x39, 0x8c, 0x11, 0x63, 0xb5, 0x48, 0xf1, 0xf5, 0xd5, 0xaf, 0x92, 0x6b, 0xb4, 0x02, 0x4d, 0x85,
	0x53, 0x64, 0x66, 0xc8, 0x28, 0xb0, 0xf3, 0x54, 0x9f, 0x24, 0x3f, 0x3f, 0xcb, 0xff, 0x9c, 0x75,
	0xf9, 0x4d, 0x01, 0x50, 0x01, 0x85, 0xcb, 0x8f, 0x82, 0xb3, 0xa7, 0xbe, 0xaf, 0x6e, 0x38, 0x73,
	0xa0, 0x45, 0x71, 0x57, 0xdc, 0x9f, 0x1b, 0x2e, 0xd3, 0x69, 0xe9, 0xe4, 0x5a, 0x6f, 0x52, 0x4c,
	0x6a, 0xc4, 0xbf, 0x89, 0x9a, 0xa4, 0x86, 0xd4, 0x0a, 0x36, 0x51, 0x9b, 0xd4, 0x92, 0xd6, 0x62,
	0x31, 0xd1, 0x3a, 0x69, 0x5d, 0x2d, 0x39, 0xf3, 0xa0, 0xc5, 0xff, 0xac, 0x17, 0xa4, 0x5e, 0x0e,
	0x6e, 0x44, 0x9d, 0xa7, 0x7b, 0xd0, 0xd4, 0x31, 0xa2, 0xfc, 0xdb, 0x31, 0xf9, 0x0e, 0x1e, 0x4c,
	0x25, 0x38, 0x0b, 0x80, 0xe2, 0x2c, 0x77, 0x4a, 0x19, 0x00, 0xe5, 0x7a, 0x93, 0xa2, 0xe7, 0xc7,
	0x97, 0x87, 0x6e, 0x87, 0x6f, 0x47, 0x2f, 0xb7, 0x11, 0x54, 0x18, 0xdd, 0xfe, 0xbd, 0x6d, 0x55,
	0x7e, 0xbd, 0x3f, 0x06, 0xfa, 0xd8, 0xae, 0xba, 0x76, 0xbf, 0xea, 0x62, 0x17, 0xf3, 0xde, 0x5e,
	0x1d, 0x3a, 0xf5, 0xb5, 0x2d, 0x5f, 0xe6, 0x86, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x82,
	0x34, 0xe4, 0x06, 0x02, 0x00, 0x00,
}