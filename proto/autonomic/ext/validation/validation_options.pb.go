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
// source: autonomic/ext/validation/validation_options.proto

package validation

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

// TimestampBoundType can be used as field option for Timestamp fields to specify to edge-processor how to validate the timestamp.
type TimestampBoundType int32

const (
	TimestampBoundType_NO_BOUND   TimestampBoundType = 0
	TimestampBoundType_CREATION   TimestampBoundType = 1
	TimestampBoundType_EXPIRATION TimestampBoundType = 2
)

var TimestampBoundType_name = map[int32]string{
	0: "NO_BOUND",
	1: "CREATION",
	2: "EXPIRATION",
}

var TimestampBoundType_value = map[string]int32{
	"NO_BOUND":   0,
	"CREATION":   1,
	"EXPIRATION": 2,
}

func (x TimestampBoundType) String() string {
	return proto.EnumName(TimestampBoundType_name, int32(x))
}

func (TimestampBoundType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4f4f45fd67270ed, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.validation.TimestampBoundType", TimestampBoundType_name, TimestampBoundType_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/validation/validation_options.proto", fileDescriptor_a4f4f45fd67270ed)
}

var fileDescriptor_a4f4f45fd67270ed = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0x2c, 0x2d, 0xc9,
	0xcf, 0xcb, 0xcf, 0xcd, 0x4c, 0xd6, 0x4f, 0xad, 0x28, 0xd1, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49,
	0x2c, 0xc9, 0xcc, 0xcf, 0x43, 0x62, 0xc6, 0xe7, 0x17, 0x80, 0xa8, 0x62, 0xbd, 0x82, 0xa2, 0xfc,
	0x92, 0x7c, 0x21, 0x09, 0xb8, 0x16, 0xbd, 0xd4, 0x8a, 0x12, 0x3d, 0x84, 0x3a, 0x2d, 0x07, 0x2e,
	0xa1, 0x90, 0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0xa7, 0xfc, 0xd2, 0xbc, 0x94, 0x90,
	0xca, 0x82, 0x54, 0x21, 0x1e, 0x2e, 0x0e, 0x3f, 0xff, 0x78, 0x27, 0xff, 0x50, 0x3f, 0x17, 0x01,
	0x06, 0x10, 0xcf, 0x39, 0xc8, 0xd5, 0x31, 0xc4, 0xd3, 0xdf, 0x4f, 0x80, 0x51, 0x88, 0x8f, 0x8b,
	0xcb, 0x35, 0x22, 0xc0, 0x33, 0x08, 0xc2, 0x67, 0x72, 0x9a, 0xc4, 0xc8, 0x25, 0x93, 0x9c, 0x9f,
	0xab, 0x87, 0xcb, 0x0a, 0x27, 0xb1, 0x30, 0x38, 0xdb, 0x1f, 0xe2, 0xaa, 0x00, 0x90, 0xa3, 0x02,
	0x18, 0xa3, 0xfc, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xe1, 0xda,
	0x75, 0x13, 0x33, 0x41, 0xfe, 0x4a, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x05, 0x3b, 0xbf, 0x58, 0xbf,
	0xb8, 0x28, 0x59, 0x3f, 0x37, 0x31, 0x33, 0x4f, 0x1f, 0xcc, 0xd7, 0xc7, 0x15, 0x00, 0x3f, 0x18,
	0x19, 0x17, 0x31, 0x31, 0x3b, 0x86, 0x86, 0x24, 0xb1, 0x81, 0x15, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x28, 0x73, 0xe7, 0xe5, 0x2c, 0x01, 0x00, 0x00,
}
