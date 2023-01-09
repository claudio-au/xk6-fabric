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
// source: autonomic/ext/telemetry/enumerations/indicator_state.proto

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

type IndicatorState int32

const (
	IndicatorState_UNKNOWN_INDICATOR_STATE IndicatorState = 0
	IndicatorState_OFF                     IndicatorState = 1
	IndicatorState_ON                      IndicatorState = 2
)

var IndicatorState_name = map[int32]string{
	0: "UNKNOWN_INDICATOR_STATE",
	1: "OFF",
	2: "ON",
}

var IndicatorState_value = map[string]int32{
	"UNKNOWN_INDICATOR_STATE": 0,
	"OFF":                     1,
	"ON":                      2,
}

func (x IndicatorState) String() string {
	return proto.EnumName(IndicatorState_name, int32(x))
}

func (IndicatorState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1fde1d3d24aa46e, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.indicatorState.IndicatorState", IndicatorState_name, IndicatorState_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/indicator_state.proto", fileDescriptor_e1fde1d3d24aa46e)
}

var fileDescriptor_e1fde1d3d24aa46e = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xcf, 0xcb, 0xcf, 0xcd, 0x4c, 0xd6, 0x4f, 0xad, 0x28, 0xd1, 0x2f, 0x49, 0xcd, 0x49, 0xcd, 0x4d,
	0x2d, 0x29, 0xaa, 0xd4, 0x4f, 0xcd, 0x2b, 0xcd, 0x4d, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b,
	0xd6, 0xcf, 0xcc, 0x4b, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0x2f, 0x8a, 0x2f, 0x2e, 0x49, 0x2c, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x86, 0xeb, 0xd5, 0x4b, 0xad, 0x28, 0xd1, 0x83,
	0xeb, 0xd5, 0x43, 0xd6, 0xab, 0x07, 0xd7, 0x1b, 0x0c, 0xd2, 0xaa, 0x65, 0xc7, 0xc5, 0xe7, 0x89,
	0x22, 0x22, 0x24, 0xcd, 0x25, 0x1e, 0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x17, 0xef, 0xe9, 0xe7,
	0xe2, 0xe9, 0xec, 0x18, 0xe2, 0x1f, 0x14, 0x1f, 0x1c, 0xe2, 0x18, 0xe2, 0x2a, 0xc0, 0x20, 0xc4,
	0xce, 0xc5, 0xec, 0xef, 0xe6, 0x26, 0xc0, 0x28, 0xc4, 0xc6, 0xc5, 0xe4, 0xef, 0x27, 0xc0, 0xe4,
	0xb4, 0x9c, 0x91, 0x4b, 0x23, 0x39, 0x3f, 0x57, 0x8f, 0x18, 0xbb, 0x9d, 0x84, 0x51, 0xad, 0x0a,
	0x00, 0x39, 0x3b, 0x80, 0x31, 0x2a, 0x3a, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x57, 0x1f, 0x6e, 0x8e, 0x6e, 0x62, 0x26, 0x28, 0x08, 0x52, 0x8b, 0xf2, 0x12, 0x73, 0x74, 0xc1,
	0x1e, 0x2c, 0xd6, 0x2f, 0x2e, 0x4a, 0xd6, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x07, 0xf3, 0xf5, 0x89,
	0x09, 0xab, 0x1f, 0x8c, 0x8c, 0x8b, 0x98, 0x98, 0x1d, 0x43, 0x43, 0x92, 0xd8, 0xc0, 0x9a, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xa6, 0x1c, 0x71, 0x63, 0x01, 0x00, 0x00,
}
