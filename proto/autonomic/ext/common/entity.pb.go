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
// source: autonomic/ext/common/entity.proto

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

type Author struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantId             string   `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2f7bdc3284c07db, []int{0}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Author) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Author) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Author) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Author)(nil), "autonomic.ext.common.entity.Author")
}

func init() { proto.RegisterFile("autonomic/ext/common/entity.proto", fileDescriptor_e2f7bdc3284c07db) }

var fileDescriptor_e2f7bdc3284c07db = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4a, 0x43, 0x31,
	0x14, 0x86, 0xc9, 0x6d, 0x2d, 0x36, 0x82, 0x60, 0xa6, 0x40, 0x07, 0xab, 0x93, 0x4b, 0x93, 0xc1,
	0x27, 0x68, 0xc5, 0x41, 0x5c, 0x8a, 0xe8, 0xe2, 0x22, 0x69, 0x12, 0xec, 0x01, 0x93, 0x73, 0xc9,
	0x3d, 0x17, 0xda, 0xc9, 0x17, 0xf0, 0x29, 0x7c, 0x4a, 0x47, 0xc9, 0xb9, 0x50, 0x1c, 0xa4, 0xdb,
	0xf9, 0xbf, 0x7c, 0x3f, 0xe4, 0x1c, 0x79, 0xe5, 0x7a, 0xc2, 0x8c, 0x09, 0xbc, 0x8d, 0x3b, 0xb2,
	0x1e, 0x53, 0xc2, 0x6c, 0x63, 0x26, 0xa0, 0xbd, 0x69, 0x0b, 0x12, 0xaa, 0xd9, 0x41, 0x31, 0x71,
	0x47, 0x66, 0x50, 0xcc, 0xa0, 0x5c, 0x7f, 0xca, 0xc9, 0xb2, 0xa7, 0x2d, 0x16, 0x75, 0x2e, 0x1b,
	0x08, 0x5a, 0xcc, 0xc5, 0xcd, 0xf4, 0xa9, 0x81, 0xa0, 0x66, 0x72, 0x4a, 0x31, 0xbb, 0x4c, 0x6f,
	0x10, 0x74, 0xc3, 0xf8, 0x74, 0x00, 0x0f, 0x41, 0x29, 0x39, 0xce, 0x2e, 0x45, 0x3d, 0x62, 0xce,
	0xb3, 0x9a, 0xcb, 0xb3, 0x10, 0x3b, 0x5f, 0xa0, 0x25, 0xc0, 0xac, 0xc7, 0xfc, 0xf4, 0x17, 0xd5,
	0x16, 0xed, 0xdb, 0xa8, 0x4f, 0x86, 0x56, 0x9d, 0x57, 0x5f, 0x42, 0x5e, 0x7a, 0x4c, 0xe6, 0xc8,
	0x27, 0x57, 0x17, 0x77, 0x1c, 0xef, 0x39, 0xad, 0xeb, 0x52, 0x6b, 0xf1, 0xfa, 0xf8, 0x0e, 0xb4,
	0xed, 0x37, 0x55, 0xb5, 0x87, 0xf2, 0xc2, 0x41, 0xbd, 0x43, 0x2c, 0xd9, 0x7d, 0x2c, 0x78, 0xfd,
	0xce, 0x76, 0xc5, 0xdb, 0xe4, 0x20, 0x5b, 0xce, 0xf6, 0xbf, 0x83, 0xfd, 0x08, 0xf1, 0xdd, 0x8c,
	0x96, 0x2f, 0xcf, 0x9b, 0x09, 0x4b, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x87, 0x73, 0x52,
	0x62, 0x58, 0x01, 0x00, 0x00,
}
