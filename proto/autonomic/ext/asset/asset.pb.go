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
// source: autonomic/ext/asset/asset.proto

package asset

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

type AssetChangeEvent_EventType int32

const (
	AssetChangeEvent_CREATED AssetChangeEvent_EventType = 0
	AssetChangeEvent_UPDATED AssetChangeEvent_EventType = 1
	AssetChangeEvent_DELETED AssetChangeEvent_EventType = 2
)

var AssetChangeEvent_EventType_name = map[int32]string{
	0: "CREATED",
	1: "UPDATED",
	2: "DELETED",
}

var AssetChangeEvent_EventType_value = map[string]int32{
	"CREATED": 0,
	"UPDATED": 1,
	"DELETED": 2,
}

func (x AssetChangeEvent_EventType) String() string {
	return proto.EnumName(AssetChangeEvent_EventType_name, int32(x))
}

func (AssetChangeEvent_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb875ae67226b88c, []int{0, 0}
}

type AssetDetails_AssetType int32

const (
	AssetDetails_VEHICLE AssetDetails_AssetType = 0
	AssetDetails_DEVICE  AssetDetails_AssetType = 1
)

var AssetDetails_AssetType_name = map[int32]string{
	0: "VEHICLE",
	1: "DEVICE",
}

var AssetDetails_AssetType_value = map[string]int32{
	"VEHICLE": 0,
	"DEVICE":  1,
}

func (x AssetDetails_AssetType) String() string {
	return proto.EnumName(AssetDetails_AssetType_name, int32(x))
}

func (AssetDetails_AssetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb875ae67226b88c, []int{1, 0}
}

type AssetChangeEvent struct {
	Type                 AssetChangeEvent_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=autonomic.ext.asset.AssetChangeEvent_EventType" json:"type,omitempty"`
	AssetDetails         *AssetDetails              `protobuf:"bytes,2,opt,name=asset_details,json=assetDetails,proto3" json:"asset_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AssetChangeEvent) Reset()         { *m = AssetChangeEvent{} }
func (m *AssetChangeEvent) String() string { return proto.CompactTextString(m) }
func (*AssetChangeEvent) ProtoMessage()    {}
func (*AssetChangeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb875ae67226b88c, []int{0}
}

func (m *AssetChangeEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetChangeEvent.Unmarshal(m, b)
}
func (m *AssetChangeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetChangeEvent.Marshal(b, m, deterministic)
}
func (m *AssetChangeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetChangeEvent.Merge(m, src)
}
func (m *AssetChangeEvent) XXX_Size() int {
	return xxx_messageInfo_AssetChangeEvent.Size(m)
}
func (m *AssetChangeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetChangeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AssetChangeEvent proto.InternalMessageInfo

func (m *AssetChangeEvent) GetType() AssetChangeEvent_EventType {
	if m != nil {
		return m.Type
	}
	return AssetChangeEvent_CREATED
}

func (m *AssetChangeEvent) GetAssetDetails() *AssetDetails {
	if m != nil {
		return m.AssetDetails
	}
	return nil
}

type AssetDetails struct {
	// The UUID of the asset.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The broad category of asset type.
	Type                 AssetDetails_AssetType `protobuf:"varint,2,opt,name=type,proto3,enum=autonomic.ext.asset.AssetDetails_AssetType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AssetDetails) Reset()         { *m = AssetDetails{} }
func (m *AssetDetails) String() string { return proto.CompactTextString(m) }
func (*AssetDetails) ProtoMessage()    {}
func (*AssetDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb875ae67226b88c, []int{1}
}

func (m *AssetDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetDetails.Unmarshal(m, b)
}
func (m *AssetDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetDetails.Marshal(b, m, deterministic)
}
func (m *AssetDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetDetails.Merge(m, src)
}
func (m *AssetDetails) XXX_Size() int {
	return xxx_messageInfo_AssetDetails.Size(m)
}
func (m *AssetDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetDetails.DiscardUnknown(m)
}

var xxx_messageInfo_AssetDetails proto.InternalMessageInfo

func (m *AssetDetails) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AssetDetails) GetType() AssetDetails_AssetType {
	if m != nil {
		return m.Type
	}
	return AssetDetails_VEHICLE
}

func init() {
	proto.RegisterEnum("autonomic.ext.asset.AssetChangeEvent_EventType", AssetChangeEvent_EventType_name, AssetChangeEvent_EventType_value)
	proto.RegisterEnum("autonomic.ext.asset.AssetDetails_AssetType", AssetDetails_AssetType_name, AssetDetails_AssetType_value)
	proto.RegisterType((*AssetChangeEvent)(nil), "autonomic.ext.asset.AssetChangeEvent")
	proto.RegisterType((*AssetDetails)(nil), "autonomic.ext.asset.AssetDetails")
}

func init() { proto.RegisterFile("autonomic/ext/asset/asset.proto", fileDescriptor_bb875ae67226b88c) }

var fileDescriptor_bb875ae67226b88c = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x6b, 0xc2, 0x30,
	0x1c, 0xc5, 0x4d, 0x36, 0x1c, 0xfe, 0x75, 0x52, 0xb2, 0xc3, 0xbc, 0xcd, 0x95, 0x1d, 0x84, 0x61,
	0x02, 0xee, 0x03, 0x0c, 0x6d, 0x33, 0xe6, 0xf0, 0x20, 0x45, 0x3d, 0xec, 0x32, 0x62, 0x0d, 0x1a,
	0xb0, 0x8d, 0xd8, 0x38, 0x14, 0x76, 0xdd, 0x17, 0xd9, 0x57, 0xda, 0x97, 0xd9, 0x71, 0x24, 0x4a,
	0x27, 0xa3, 0x63, 0x97, 0xd0, 0x5f, 0xfb, 0xde, 0x6b, 0x1e, 0x0f, 0xae, 0xc4, 0xc6, 0xe8, 0x54,
	0x27, 0x2a, 0x66, 0x72, 0x6b, 0x98, 0xc8, 0x32, 0x79, 0x38, 0xe9, 0x6a, 0xad, 0x8d, 0x26, 0x17,
	0xb9, 0x80, 0xca, 0xad, 0xa1, 0xee, 0x93, 0xff, 0x89, 0xc0, 0xeb, 0xda, 0xa7, 0x60, 0x21, 0xd2,
	0xb9, 0xe4, 0xaf, 0x32, 0x35, 0x24, 0x80, 0x53, 0xb3, 0x5b, 0xc9, 0x06, 0x6a, 0xa2, 0x56, 0xbd,
	0xc3, 0x68, 0x81, 0x91, 0xfe, 0x36, 0x51, 0x77, 0x8e, 0x76, 0x2b, 0x19, 0x39, 0x33, 0x79, 0x80,
	0x73, 0xa7, 0x7c, 0x99, 0x49, 0x23, 0xd4, 0x32, 0x6b, 0xe0, 0x26, 0x6a, 0x55, 0x3b, 0xd7, 0x7f,
	0xa7, 0x85, 0x7b, 0x61, 0x54, 0x13, 0x47, 0xe4, 0x77, 0xa0, 0x92, 0x47, 0x93, 0x2a, 0x9c, 0x05,
	0x11, 0xef, 0x8e, 0x78, 0xe8, 0x95, 0x2c, 0x8c, 0x87, 0xa1, 0x03, 0x64, 0x21, 0xe4, 0x03, 0x6e,
	0x01, 0xfb, 0xef, 0x08, 0x6a, 0xc7, 0x91, 0xa4, 0x0e, 0x58, 0xcd, 0x5c, 0x9f, 0x4a, 0x84, 0xd5,
	0x8c, 0xdc, 0x1f, 0x1a, 0x62, 0xd7, 0xf0, 0xf6, 0xdf, 0x3b, 0xed, 0xe1, 0xa7, 0x9d, 0x7f, 0x03,
	0x95, 0xfc, 0x95, 0xfd, 0xf7, 0x84, 0x3f, 0xf6, 0x83, 0x01, 0xf7, 0x4a, 0x04, 0xa0, 0x1c, 0xf2,
	0x49, 0x3f, 0xe0, 0x1e, 0xea, 0xbd, 0xc1, 0x65, 0xac, 0x93, 0xa2, 0xf4, 0x1e, 0x38, 0xfb, 0xd0,
	0x2e, 0x33, 0x44, 0xcf, 0x4f, 0x73, 0x65, 0x16, 0x9b, 0x29, 0x8d, 0x75, 0xc2, 0x72, 0x75, 0x5b,
	0x28, 0x3b, 0xa5, 0x5c, 0xa7, 0x62, 0xd9, 0x76, 0x1b, 0x66, 0x2c, 0x5b, 0xc7, 0x2c, 0x11, 0x2a,
	0x65, 0x8e, 0x59, 0xc1, 0xe6, 0x5f, 0x08, 0x7d, 0xe0, 0x93, 0xee, 0x78, 0x34, 0x2d, 0x3b, 0xcd,
	0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xd9, 0xf6, 0x12, 0x1a, 0x02, 0x00, 0x00,
}