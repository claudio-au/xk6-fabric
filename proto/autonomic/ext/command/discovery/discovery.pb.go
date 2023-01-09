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
// source: autonomic/ext/command/discovery/discovery.proto

package discovery

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

// A request to perform a discovery operation
type DiscoveryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a2aed7a1665296, []int{0}
}

func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(m, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DiscoveryRequest)(nil), "autonomic.ext.command.discovery.DiscoveryRequest")
}

func init() {
	proto.RegisterFile("autonomic/ext/command/discovery/discovery.proto", fileDescriptor_d7a2aed7a1665296)
}

var fileDescriptor_d7a2aed7a1665296 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0x2c, 0x2d, 0xc9,
	0xcf, 0xcb, 0xcf, 0xcd, 0x4c, 0xd6, 0x4f, 0xad, 0x28, 0xd1, 0x4f, 0xce, 0xcf, 0xcd, 0x4d, 0xcc,
	0x4b, 0xd1, 0x4f, 0xc9, 0x2c, 0x4e, 0xce, 0x2f, 0x4b, 0x2d, 0xaa, 0x44, 0xb0, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0xe1, 0x1a, 0xf4, 0x52, 0x2b, 0x4a, 0xf4, 0xa0, 0x1a, 0xf4, 0xe0,
	0xca, 0x94, 0x84, 0xb8, 0x04, 0x5c, 0x60, 0x9c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xa7,
	0x19, 0x8c, 0x5c, 0xca, 0xc9, 0xf9, 0xb9, 0x7a, 0x04, 0xf4, 0x3a, 0xf1, 0xc1, 0x75, 0x06, 0x80,
	0x2c, 0x0b, 0x60, 0x8c, 0x0a, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x02, 0xa9, 0x45, 0x38, 0x55,
	0x37, 0x31, 0x13, 0xe4, 0xda, 0xd4, 0xa2, 0xbc, 0xc4, 0x1c, 0x5d, 0xb0, 0xb3, 0x8a, 0xf5, 0x8b,
	0x8b, 0x92, 0xf5, 0x73, 0x13, 0x33, 0xf3, 0xf4, 0xc1, 0x7c, 0x42, 0xde, 0xfa, 0xc1, 0xc8, 0xb8,
	0x88, 0x89, 0xd9, 0x31, 0x34, 0x24, 0x89, 0x0d, 0xac, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xd9, 0xbe, 0x9d, 0xc8, 0x09, 0x01, 0x00, 0x00,
}
