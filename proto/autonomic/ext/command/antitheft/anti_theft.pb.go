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
// source: autonomic/ext/command/antitheft/anti_theft.proto

package antitheft

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

type AntiTheftRequest_AntiTheftRequestType int32

const (
	AntiTheftRequest_UNKNOWN_ANTI_THEFT_REQUEST_TYPE    AntiTheftRequest_AntiTheftRequestType = 0
	AntiTheftRequest_DISABLE_EARLY_ANTI_THEFT_REPORTING AntiTheftRequest_AntiTheftRequestType = 2
	AntiTheftRequest_ENABLE_ANTI_THEFT_MODE             AntiTheftRequest_AntiTheftRequestType = 3
	AntiTheftRequest_DISABLE_ANTI_THEFT_MODE            AntiTheftRequest_AntiTheftRequestType = 4
	AntiTheftRequest_REPORT_CURRENT_LOCATION            AntiTheftRequest_AntiTheftRequestType = 5
)

var AntiTheftRequest_AntiTheftRequestType_name = map[int32]string{
	0: "UNKNOWN_ANTI_THEFT_REQUEST_TYPE",
	2: "DISABLE_EARLY_ANTI_THEFT_REPORTING",
	3: "ENABLE_ANTI_THEFT_MODE",
	4: "DISABLE_ANTI_THEFT_MODE",
	5: "REPORT_CURRENT_LOCATION",
}

var AntiTheftRequest_AntiTheftRequestType_value = map[string]int32{
	"UNKNOWN_ANTI_THEFT_REQUEST_TYPE":    0,
	"DISABLE_EARLY_ANTI_THEFT_REPORTING": 2,
	"ENABLE_ANTI_THEFT_MODE":             3,
	"DISABLE_ANTI_THEFT_MODE":            4,
	"REPORT_CURRENT_LOCATION":            5,
}

func (x AntiTheftRequest_AntiTheftRequestType) String() string {
	return proto.EnumName(AntiTheftRequest_AntiTheftRequestType_name, int32(x))
}

func (AntiTheftRequest_AntiTheftRequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4439587992f8d8a7, []int{0, 0}
}

// A message to request actions for a possibly stolen asset
type AntiTheftRequest struct {
	AntiTheftRequestType AntiTheftRequest_AntiTheftRequestType `protobuf:"varint,1,opt,name=anti_theft_request_type,json=antiTheftRequestType,proto3,enum=autonomic.ext.command.antitheft.AntiTheftRequest_AntiTheftRequestType" json:"anti_theft_request_type,omitempty"`
	// frequency of reporting if reporting/mode is enabled
	FrequencySeconds     int32    `protobuf:"varint,3,opt,name=frequency_seconds,json=frequencySeconds,proto3" json:"frequency_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AntiTheftRequest) Reset()         { *m = AntiTheftRequest{} }
func (m *AntiTheftRequest) String() string { return proto.CompactTextString(m) }
func (*AntiTheftRequest) ProtoMessage()    {}
func (*AntiTheftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4439587992f8d8a7, []int{0}
}

func (m *AntiTheftRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AntiTheftRequest.Unmarshal(m, b)
}
func (m *AntiTheftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AntiTheftRequest.Marshal(b, m, deterministic)
}
func (m *AntiTheftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AntiTheftRequest.Merge(m, src)
}
func (m *AntiTheftRequest) XXX_Size() int {
	return xxx_messageInfo_AntiTheftRequest.Size(m)
}
func (m *AntiTheftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AntiTheftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AntiTheftRequest proto.InternalMessageInfo

func (m *AntiTheftRequest) GetAntiTheftRequestType() AntiTheftRequest_AntiTheftRequestType {
	if m != nil {
		return m.AntiTheftRequestType
	}
	return AntiTheftRequest_UNKNOWN_ANTI_THEFT_REQUEST_TYPE
}

func (m *AntiTheftRequest) GetFrequencySeconds() int32 {
	if m != nil {
		return m.FrequencySeconds
	}
	return 0
}

func init() {
	proto.RegisterEnum("autonomic.ext.command.antitheft.AntiTheftRequest_AntiTheftRequestType", AntiTheftRequest_AntiTheftRequestType_name, AntiTheftRequest_AntiTheftRequestType_value)
	proto.RegisterType((*AntiTheftRequest)(nil), "autonomic.ext.command.antitheft.AntiTheftRequest")
}

func init() {
	proto.RegisterFile("autonomic/ext/command/antitheft/anti_theft.proto", fileDescriptor_4439587992f8d8a7)
}

var fileDescriptor_4439587992f8d8a7 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0xcf, 0xd2, 0x30,
	0x18, 0xc7, 0x2d, 0x03, 0x43, 0x7a, 0x20, 0x75, 0x21, 0x42, 0xf4, 0x00, 0x19, 0x89, 0x21, 0x31,
	0xac, 0x46, 0x3f, 0xc1, 0x80, 0xa2, 0x53, 0xec, 0x66, 0xe9, 0x34, 0x78, 0x69, 0xc6, 0x28, 0xb2,
	0xc4, 0x75, 0xb8, 0x95, 0x04, 0x0e, 0x5e, 0xfd, 0x0c, 0x9e, 0xfd, 0x22, 0x7e, 0x2d, 0x8f, 0x66,
	0x43, 0xf7, 0xbe, 0x2f, 0x2f, 0x09, 0xb7, 0xe7, 0xf9, 0x3f, 0xff, 0xff, 0xaf, 0x69, 0xfb, 0xc0,
	0x17, 0xe1, 0x5e, 0xa7, 0x2a, 0x4d, 0xe2, 0x08, 0xcb, 0x83, 0xc6, 0x51, 0x9a, 0x24, 0xa1, 0x5a,
	0xe3, 0x50, 0xe9, 0x58, 0x6f, 0xe5, 0x46, 0x97, 0x95, 0x28, 0x4b, 0x7b, 0x97, 0xa5, 0x3a, 0x35,
	0x7b, 0x55, 0xc2, 0x96, 0x07, 0x6d, 0xff, 0x4b, 0xd8, 0x55, 0xc2, 0xfa, 0x61, 0x40, 0xe4, 0x28,
	0x1d, 0xf3, 0xa2, 0x63, 0xf2, 0xdb, 0x5e, 0xe6, 0xda, 0xfc, 0x0e, 0x3b, 0x37, 0x24, 0x91, 0x9d,
	0x54, 0xa1, 0x8f, 0x3b, 0xd9, 0x05, 0x7d, 0x30, 0x6c, 0xbd, 0x9c, 0xd9, 0x57, 0xb8, 0xf6, 0x39,
	0xf3, 0x9e, 0xc0, 0x8f, 0x3b, 0xc9, 0xda, 0xe1, 0x05, 0xd5, 0x7c, 0x0e, 0x1f, 0x6d, 0xca, 0x43,
	0x55, 0x74, 0x14, 0xb9, 0x8c, 0x52, 0xb5, 0xce, 0xbb, 0x46, 0x1f, 0x0c, 0x1b, 0x0c, 0x55, 0x83,
	0xc5, 0x49, 0xb7, 0x7e, 0x03, 0xd8, 0xbe, 0xc4, 0x36, 0x07, 0xb0, 0x17, 0xd0, 0x77, 0xd4, 0xfb,
	0x44, 0x85, 0x43, 0xb9, 0x2b, 0xf8, 0x1b, 0x32, 0xe3, 0x82, 0x91, 0x0f, 0x01, 0x59, 0x70, 0xc1,
	0x97, 0x3e, 0x41, 0x0f, 0xcc, 0x67, 0xd0, 0x9a, 0xba, 0x0b, 0x67, 0x3c, 0x27, 0x82, 0x38, 0x6c,
	0xbe, 0xbc, 0x6b, 0xf5, 0x3d, 0xc6, 0x5d, 0xfa, 0x1a, 0xd5, 0xcc, 0x27, 0xf0, 0x31, 0xa1, 0xa5,
	0xed, 0x96, 0xe1, 0xbd, 0x37, 0x25, 0xc8, 0x30, 0x9f, 0xc2, 0xce, 0x7f, 0xc6, 0xf9, 0xb0, 0x5e,
	0x0c, 0x4f, 0x1c, 0x31, 0x09, 0x18, 0x23, 0x94, 0x8b, 0xb9, 0x37, 0x71, 0xb8, 0xeb, 0x51, 0xd4,
	0xb0, 0xea, 0x4d, 0x80, 0xc0, 0xdb, 0x7a, 0xb3, 0x86, 0x8c, 0xf1, 0x4f, 0x00, 0x07, 0x51, 0x9a,
	0x5c, 0x7b, 0xd8, 0x71, 0xab, 0xba, 0xac, 0x5f, 0xfc, 0xb0, 0x0f, 0x3e, 0x7f, 0xfc, 0x12, 0xeb,
	0xed, 0x7e, 0x55, 0x78, 0x71, 0x95, 0x1e, 0x85, 0x71, 0xb1, 0x23, 0x32, 0x53, 0xe1, 0xd7, 0x51,
	0xb9, 0x0b, 0x39, 0xce, 0xb3, 0x08, 0x27, 0x61, 0xac, 0x70, 0xd9, 0xe3, 0x2b, 0xcb, 0xf4, 0x07,
	0x80, 0x5f, 0x35, 0xc3, 0x09, 0xf8, 0xea, 0x61, 0xe9, 0x7f, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x56, 0x25, 0x1c, 0x8b, 0x7f, 0x02, 0x00, 0x00,
}
