/*-
 * ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
 * Autonomic Proprietary 1.0
 * ——————————————————————————————————————————————————————————————————————————————
 * Copyright (C) 2020 Autonomic, LLC - All rights reserved
 * ——————————————————————————————————————————————————————————————————————————————
 * Proprietary and confidential.
 *
 * NOTICE:  All information contained herein is, and remains the property of
 * Autonomic, LLC and its suppliers, if any.  The intellectual and technical
 * concepts contained herein are proprietary to Autonomic, LLC and its suppliers
 * and may be covered by U.S. and Foreign Patents, patents in process, and are
 * protected by trade secret or copyright law. Dissemination of this information
 * or reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Autonomic, LLC.
 *
 * Unauthorized copy of this file, via any medium is strictly prohibited.
 * ______________________________________________________________________________
 */
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autonomic/ext/common/range.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Range struct {
	// Types that are valid to be assigned to Value:
	//	*Range_IntRange
	//	*Range_DoubleRange
	Value                isRange_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4389dbcf6e302fc, []int{0}
}

func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

type isRange_Value interface {
	isRange_Value()
}

type Range_IntRange struct {
	IntRange *IntRange `protobuf:"bytes,1,opt,name=int_range,json=intRange,proto3,oneof"`
}

type Range_DoubleRange struct {
	DoubleRange *DoubleRange `protobuf:"bytes,2,opt,name=double_range,json=doubleRange,proto3,oneof"`
}

func (*Range_IntRange) isRange_Value() {}

func (*Range_DoubleRange) isRange_Value() {}

func (m *Range) GetValue() isRange_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Range) GetIntRange() *IntRange {
	if x, ok := m.GetValue().(*Range_IntRange); ok {
		return x.IntRange
	}
	return nil
}

func (m *Range) GetDoubleRange() *DoubleRange {
	if x, ok := m.GetValue().(*Range_DoubleRange); ok {
		return x.DoubleRange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Range) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Range_IntRange)(nil),
		(*Range_DoubleRange)(nil),
	}
}

type IntRange struct {
	LowerBound           int64    `protobuf:"varint,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	UpperBound           int64    `protobuf:"varint,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntRange) Reset()         { *m = IntRange{} }
func (m *IntRange) String() string { return proto.CompactTextString(m) }
func (*IntRange) ProtoMessage()    {}
func (*IntRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4389dbcf6e302fc, []int{1}
}

func (m *IntRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntRange.Unmarshal(m, b)
}
func (m *IntRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntRange.Marshal(b, m, deterministic)
}
func (m *IntRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntRange.Merge(m, src)
}
func (m *IntRange) XXX_Size() int {
	return xxx_messageInfo_IntRange.Size(m)
}
func (m *IntRange) XXX_DiscardUnknown() {
	xxx_messageInfo_IntRange.DiscardUnknown(m)
}

var xxx_messageInfo_IntRange proto.InternalMessageInfo

func (m *IntRange) GetLowerBound() int64 {
	if m != nil {
		return m.LowerBound
	}
	return 0
}

func (m *IntRange) GetUpperBound() int64 {
	if m != nil {
		return m.UpperBound
	}
	return 0
}

type DoubleRange struct {
	LowerBound           float64  `protobuf:"fixed64,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	UpperBound           float64  `protobuf:"fixed64,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleRange) Reset()         { *m = DoubleRange{} }
func (m *DoubleRange) String() string { return proto.CompactTextString(m) }
func (*DoubleRange) ProtoMessage()    {}
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4389dbcf6e302fc, []int{2}
}

func (m *DoubleRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleRange.Unmarshal(m, b)
}
func (m *DoubleRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleRange.Marshal(b, m, deterministic)
}
func (m *DoubleRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRange.Merge(m, src)
}
func (m *DoubleRange) XXX_Size() int {
	return xxx_messageInfo_DoubleRange.Size(m)
}
func (m *DoubleRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRange.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRange proto.InternalMessageInfo

func (m *DoubleRange) GetLowerBound() float64 {
	if m != nil {
		return m.LowerBound
	}
	return 0
}

func (m *DoubleRange) GetUpperBound() float64 {
	if m != nil {
		return m.UpperBound
	}
	return 0
}

func init() {
	proto.RegisterType((*Range)(nil), "autonomic.ext.common.entity.Range")
	proto.RegisterType((*IntRange)(nil), "autonomic.ext.common.entity.IntRange")
	proto.RegisterType((*DoubleRange)(nil), "autonomic.ext.common.entity.DoubleRange")
}

func init() { proto.RegisterFile("autonomic/ext/common/range.proto", fileDescriptor_a4389dbcf6e302fc) }

var fileDescriptor_a4389dbcf6e302fc = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xbf, 0x4d, 0xe9, 0x67, 0x9d, 0x78, 0x90, 0x9c, 0x04, 0x41, 0x4b, 0x40, 0xe8, 0xa5,
	0xbb, 0xa0, 0x4f, 0x60, 0xec, 0x41, 0x51, 0xb1, 0x04, 0xbd, 0x78, 0x29, 0x9b, 0x64, 0xa9, 0x0b,
	0xc9, 0x4e, 0x48, 0x77, 0xb5, 0xbe, 0x82, 0x8f, 0x21, 0xf8, 0x8e, 0x1e, 0x65, 0xa7, 0x6d, 0xec,
	0x21, 0xd4, 0x5b, 0x66, 0xf8, 0xfd, 0x7f, 0x93, 0xe4, 0x0f, 0x43, 0xe9, 0x2c, 0x1a, 0xac, 0x74,
	0x2e, 0xd4, 0xd2, 0x8a, 0x1c, 0xab, 0x0a, 0x8d, 0x68, 0xa4, 0x99, 0x2b, 0x5e, 0x37, 0x68, 0x31,
	0x3a, 0x6e, 0x09, 0xae, 0x96, 0x96, 0xaf, 0x08, 0xae, 0x8c, 0xd5, 0xf6, 0x3d, 0xfe, 0x62, 0xd0,
	0x4f, 0x3d, 0x1c, 0x4d, 0x60, 0x5f, 0x1b, 0x3b, 0xa3, 0xe4, 0x11, 0x1b, 0xb2, 0x51, 0x78, 0x7e,
	0xc6, 0x77, 0x44, 0xf9, 0x8d, 0xb1, 0x94, 0xbc, 0xfe, 0x97, 0x0e, 0xf4, 0xfa, 0x39, 0xba, 0x87,
	0x83, 0x02, 0x5d, 0x56, 0xaa, 0xb5, 0x28, 0x20, 0xd1, 0x68, 0xa7, 0x68, 0x42, 0x81, 0x8d, 0x2b,
	0x2c, 0x7e, 0xc7, 0x64, 0x0f, 0xfa, 0xaf, 0xb2, 0x74, 0x2a, 0xbe, 0x83, 0xc1, 0xe6, 0x5e, 0x74,
	0x0a, 0x61, 0x89, 0x6f, 0xaa, 0x99, 0x65, 0xe8, 0x4c, 0x41, 0xef, 0xda, 0x4b, 0x81, 0x56, 0x89,
	0xdf, 0x78, 0xc0, 0xd5, 0x75, 0x0b, 0x04, 0x2b, 0x80, 0x56, 0x04, 0xc4, 0x0f, 0x10, 0x6e, 0x1d,
	0xed, 0x12, 0xb2, 0xbf, 0x84, 0x6c, 0x5b, 0x98, 0x7c, 0x30, 0x38, 0xc9, 0xb1, 0xea, 0xfe, 0x4c,
	0xfa, 0x13, 0xc9, 0xe1, 0x15, 0x4d, 0x74, 0x71, 0xea, 0x8b, 0x99, 0xb2, 0xe7, 0xdb, 0xb9, 0xb6,
	0x2f, 0x2e, 0xf3, 0xa0, 0x68, 0xa3, 0x63, 0xa9, 0x7d, 0x95, 0xaa, 0x31, 0xb2, 0x1c, 0x53, 0x85,
	0x0b, 0xb1, 0x68, 0x72, 0x51, 0x49, 0x6d, 0x04, 0xcd, 0xa2, 0xab, 0xf3, 0x6f, 0xc6, 0x3e, 0x83,
	0xde, 0xe5, 0xd3, 0x63, 0xf6, 0x9f, 0xa0, 0x8b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xa0,
	0x40, 0x74, 0x1b, 0x02, 0x00, 0x00,
}
