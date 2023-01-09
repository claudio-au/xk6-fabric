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
// source: autonomic/ext/telemetry/enumerations/error.proto

package enumerations

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

// Error representation for signals.
type Error int32

const (
	Error_UNKNOWN      Error = 0
	Error_FAULT        Error = 1 // Deprecated: Do not use.
	Error_INVALID      Error = 3
	Error_OUT_OF_RANGE Error = 4
	Error_FAILURE      Error = 5
)

var Error_name = map[int32]string{
	0: "UNKNOWN",
	1: "FAULT",
	3: "INVALID",
	4: "OUT_OF_RANGE",
	5: "FAILURE",
}

var Error_value = map[string]int32{
	"UNKNOWN":      0,
	"FAULT":        1,
	"INVALID":      3,
	"OUT_OF_RANGE": 4,
	"FAILURE":      5,
}

func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}

func (Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc2bfe46513ffd53, []int{0}
}

type ErrorSource int32

const (
	ErrorSource_ERROR_SOURCE_UNKNOWN ErrorSource = 0
	ErrorSource_DEVICE               ErrorSource = 1
)

var ErrorSource_name = map[int32]string{
	0: "ERROR_SOURCE_UNKNOWN",
	1: "DEVICE",
}

var ErrorSource_value = map[string]int32{
	"ERROR_SOURCE_UNKNOWN": 0,
	"DEVICE":               1,
}

func (x ErrorSource) String() string {
	return proto.EnumName(ErrorSource_name, int32(x))
}

func (ErrorSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc2bfe46513ffd53, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.error.Error", Error_name, Error_value)
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.error.ErrorSource", ErrorSource_name, ErrorSource_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/error.proto", fileDescriptor_bc2bfe46513ffd53)
}

var fileDescriptor_bc2bfe46513ffd53 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x7f, 0xe9, 0xfe, 0xfc, 0xc6, 0x33, 0x85, 0x50, 0x3c, 0x88, 0x47, 0x4f, 0x52, 0x58,
	0x22, 0xec, 0x15, 0x74, 0x5b, 0x2a, 0xc5, 0xd2, 0x8e, 0xac, 0x99, 0xa0, 0x87, 0xd2, 0x95, 0xa0,
	0x85, 0xa5, 0x91, 0x34, 0x85, 0xf9, 0x66, 0x3c, 0xf8, 0x2a, 0x3d, 0x4a, 0x73, 0xa8, 0x1e, 0x77,
	0x4b, 0xf8, 0xf2, 0xf9, 0x3e, 0xcf, 0xf3, 0x81, 0xfb, 0xb2, 0xb3, 0xba, 0xd1, 0xaa, 0xae, 0xa8,
	0x3c, 0x59, 0x6a, 0xe5, 0x51, 0x2a, 0x69, 0xcd, 0x07, 0x95, 0x4d, 0xa7, 0xa4, 0x29, 0x6d, 0xad,
	0x9b, 0x96, 0x4a, 0x63, 0xb4, 0x21, 0xef, 0x46, 0x5b, 0xed, 0x07, 0x03, 0x41, 0xe4, 0xc9, 0x92,
	0x81, 0x20, 0x7f, 0x09, 0xe2, 0x88, 0x40, 0xc0, 0x84, 0xf5, 0x0f, 0x7f, 0x0e, 0xff, 0x45, 0xfa,
	0x98, 0x66, 0x4f, 0x29, 0xfe, 0xe7, 0x5f, 0xc2, 0x24, 0x0a, 0x45, 0x92, 0x63, 0x74, 0xe3, 0xcd,
	0x50, 0x9f, 0xc5, 0xe9, 0x3e, 0x4c, 0xe2, 0x0d, 0x1e, 0xf9, 0x18, 0x2e, 0x32, 0x91, 0x17, 0x59,
	0x54, 0xf0, 0x30, 0x7d, 0x60, 0x78, 0xdc, 0xc7, 0x51, 0x18, 0x27, 0x82, 0x33, 0x3c, 0xb9, 0x1d,
	0xcf, 0x3c, 0xec, 0x05, 0x4b, 0x98, 0xbb, 0xda, 0x9d, 0xee, 0x4c, 0x25, 0xfd, 0x6b, 0xb8, 0x62,
	0x9c, 0x67, 0xbc, 0xd8, 0x65, 0x82, 0xaf, 0x59, 0xf1, 0x3b, 0x09, 0x60, 0xba, 0x61, 0xfb, 0x78,
	0xcd, 0x30, 0x5a, 0x7d, 0x22, 0xb8, 0xab, 0xb4, 0x22, 0xe7, 0xac, 0xbf, 0x02, 0xd7, 0xbf, 0xed,
	0x0f, 0xde, 0xa2, 0xe7, 0x97, 0xd7, 0xda, 0xbe, 0x75, 0x07, 0x52, 0x69, 0x45, 0x07, 0x7c, 0x51,
	0xd6, 0xbd, 0x32, 0x69, 0x9a, 0xf2, 0xb8, 0x70, 0x6a, 0x5a, 0xda, 0x9a, 0x8a, 0xaa, 0xb2, 0x6e,
	0xa8, 0xfb, 0xd3, 0x73, 0xdc, 0x7e, 0x23, 0xf4, 0xe5, 0x8d, 0x42, 0x91, 0x1f, 0xa6, 0x0e, 0x5a,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x12, 0x1d, 0x69, 0x93, 0x01, 0x00, 0x00,
}