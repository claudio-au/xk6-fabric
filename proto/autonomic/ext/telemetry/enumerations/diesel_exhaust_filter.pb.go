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
// source: autonomic/ext/telemetry/enumerations/diesel_exhaust_filter.proto

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

type DieselExhaustFilter int32

const (
	DieselExhaustFilter_UNKNOWN              DieselExhaustFilter = 0
	DieselExhaustFilter_NORMAL_OPERATION     DieselExhaustFilter = 1
	DieselExhaustFilter_OVERLOADED           DieselExhaustFilter = 2
	DieselExhaustFilter_CLEANING_IN_PROGRESS DieselExhaustFilter = 3 // Deprecated: Do not use.
	DieselExhaustFilter_AT_LIMIT             DieselExhaustFilter = 4
	DieselExhaustFilter_OVER_LIMIT           DieselExhaustFilter = 5
)

var DieselExhaustFilter_name = map[int32]string{
	0: "UNKNOWN",
	1: "NORMAL_OPERATION",
	2: "OVERLOADED",
	3: "CLEANING_IN_PROGRESS",
	4: "AT_LIMIT",
	5: "OVER_LIMIT",
}

var DieselExhaustFilter_value = map[string]int32{
	"UNKNOWN":              0,
	"NORMAL_OPERATION":     1,
	"OVERLOADED":           2,
	"CLEANING_IN_PROGRESS": 3,
	"AT_LIMIT":             4,
	"OVER_LIMIT":           5,
}

func (x DieselExhaustFilter) String() string {
	return proto.EnumName(DieselExhaustFilter_name, int32(x))
}

func (DieselExhaustFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b71d207f15cc1ac4, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.dieselExhaustFilter.DieselExhaustFilter", DieselExhaustFilter_name, DieselExhaustFilter_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/diesel_exhaust_filter.proto", fileDescriptor_b71d207f15cc1ac4)
}

var fileDescriptor_b71d207f15cc1ac4 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xbf, 0x6d, 0x3f, 0xb5, 0xac, 0x22, 0x21, 0xf6, 0x20, 0xe2, 0x0f, 0x10, 0xa1, 0xbb,
	0x07, 0x2f, 0x1e, 0x4d, 0x6d, 0x2c, 0xc1, 0x74, 0x37, 0xa4, 0xa9, 0x82, 0x1e, 0x96, 0x6d, 0x1c,
	0xed, 0x42, 0x36, 0x2b, 0x9b, 0x0d, 0xd4, 0xbb, 0xbf, 0xc4, 0xb3, 0x3f, 0xd0, 0xa3, 0x34, 0xd5,
	0x20, 0xe8, 0xa1, 0xc7, 0x77, 0xe0, 0x7d, 0x98, 0x79, 0x06, 0x5f, 0xc8, 0xda, 0x99, 0xd2, 0x68,
	0x95, 0x53, 0x58, 0x3a, 0xea, 0xa0, 0x00, 0x0d, 0xce, 0xbe, 0x50, 0x28, 0x6b, 0x0d, 0x56, 0x3a,
	0x65, 0xca, 0x8a, 0x3e, 0x28, 0xa8, 0xa0, 0x10, 0xb0, 0x5c, 0xc8, 0xba, 0x72, 0xe2, 0x51, 0x15,
	0x0e, 0x2c, 0x79, 0xb6, 0xc6, 0x19, 0xff, 0xbc, 0x25, 0x10, 0x58, 0x3a, 0xd2, 0x12, 0xc8, 0x4f,
	0x02, 0x59, 0x13, 0xc2, 0x35, 0xe0, 0xaa, 0xe9, 0x9f, 0xbe, 0x22, 0x7c, 0x30, 0xfa, 0x3d, 0xf7,
	0x77, 0xf1, 0xce, 0x8c, 0x5d, 0x33, 0x7e, 0xcb, 0xbc, 0x7f, 0x7e, 0x1f, 0x7b, 0x8c, 0xa7, 0x93,
	0x20, 0x16, 0x3c, 0x09, 0xd3, 0x20, 0x8b, 0x38, 0xf3, 0x90, 0xbf, 0x8f, 0x31, 0xbf, 0x09, 0xd3,
	0x98, 0x07, 0xa3, 0x70, 0xe4, 0x75, 0xfc, 0x63, 0xdc, 0xbf, 0x8c, 0xc3, 0x80, 0x45, 0x6c, 0x2c,
	0x22, 0x26, 0x92, 0x94, 0x8f, 0xd3, 0x70, 0x3a, 0xf5, 0xba, 0x47, 0x9d, 0x1e, 0xf2, 0xf7, 0x70,
	0x2f, 0xc8, 0x44, 0x1c, 0x4d, 0xa2, 0xcc, 0xfb, 0xff, 0xdd, 0xfd, 0xca, 0x5b, 0xc3, 0x77, 0x84,
	0x4f, 0x72, 0xa3, 0xc9, 0x26, 0x77, 0x0c, 0x0f, 0xff, 0x58, 0x38, 0x59, 0x79, 0x48, 0xd0, 0xdd,
	0xfd, 0x93, 0x72, 0x8b, 0x7a, 0x4e, 0x72, 0xa3, 0x69, 0x0b, 0x1b, 0x48, 0xb5, 0x32, 0x0b, 0xb6,
	0x94, 0xc5, 0xa0, 0x31, 0x56, 0xd1, 0xca, 0xe6, 0x54, 0x4b, 0x55, 0xd2, 0x26, 0xd3, 0x4d, 0x5e,
	0xf0, 0x81, 0xd0, 0x5b, 0xa7, 0x1b, 0xcc, 0xb2, 0xf9, 0x76, 0x53, 0x3a, 0xfb, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0xec, 0x7e, 0xaf, 0xcd, 0xba, 0x01, 0x00, 0x00,
}
