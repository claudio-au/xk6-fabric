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
// source: autonomic/ext/telemetry/enumerations/mitigation_brake_status.proto

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

type MitigationBrakeStatus int32

const (
	MitigationBrakeStatus_UNKNOWN               MitigationBrakeStatus = 0
	MitigationBrakeStatus_BRAKE_NOT_PRE_CHARGED MitigationBrakeStatus = 1
	MitigationBrakeStatus_BRAKE_PRE_CHARGED     MitigationBrakeStatus = 2
	MitigationBrakeStatus_BRAKE_APPLIED         MitigationBrakeStatus = 3
)

var MitigationBrakeStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "BRAKE_NOT_PRE_CHARGED",
	2: "BRAKE_PRE_CHARGED",
	3: "BRAKE_APPLIED",
}

var MitigationBrakeStatus_value = map[string]int32{
	"UNKNOWN":               0,
	"BRAKE_NOT_PRE_CHARGED": 1,
	"BRAKE_PRE_CHARGED":     2,
	"BRAKE_APPLIED":         3,
}

func (x MitigationBrakeStatus) String() string {
	return proto.EnumName(MitigationBrakeStatus_name, int32(x))
}

func (MitigationBrakeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2115a4df57cb357c, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.mitigationBrakeStatus.MitigationBrakeStatus", MitigationBrakeStatus_name, MitigationBrakeStatus_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/mitigation_brake_status.proto", fileDescriptor_2115a4df57cb357c)
}

var fileDescriptor_2115a4df57cb357c = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x4d, 0x0b, 0x0a, 0x11, 0xa1, 0x0d, 0xec, 0x41, 0x9f, 0x40, 0x84, 0x26, 0x07, 0x6f,
	0xde, 0x36, 0x76, 0x51, 0xa9, 0x6e, 0xc3, 0xda, 0x22, 0xe8, 0x21, 0x64, 0x97, 0x50, 0x83, 0x4d,
	0x22, 0xc9, 0x2c, 0xd4, 0xd7, 0xf1, 0x01, 0x7c, 0x3e, 0x8f, 0xb2, 0x29, 0xac, 0x3d, 0xf4, 0xd0,
	0xdb, 0xcc, 0xff, 0x33, 0x1f, 0xff, 0xfc, 0x98, 0xab, 0x16, 0xbc, 0xf3, 0xd6, 0x34, 0x4c, 0x6f,
	0x80, 0x81, 0x5e, 0x6b, 0xab, 0x21, 0x7c, 0x31, 0xed, 0x5a, 0xab, 0x83, 0x02, 0xe3, 0x5d, 0x64,
	0xd6, 0x80, 0x59, 0xa5, 0x59, 0xd6, 0x41, 0x7d, 0x68, 0x19, 0x41, 0x41, 0x1b, 0xe9, 0x67, 0xf0,
	0xe0, 0xc9, 0x4d, 0xcf, 0xa0, 0x7a, 0x03, 0xb4, 0x67, 0xd0, 0x5d, 0x06, 0xfd, 0x67, 0xf0, 0x0e,
	0xf1, 0x9c, 0x08, 0x57, 0x06, 0x67, 0x4f, 0xfb, 0x0c, 0x72, 0x8a, 0x4f, 0x96, 0xe5, 0xac, 0x9c,
	0xbf, 0x94, 0xa3, 0x23, 0x72, 0x8e, 0x33, 0x5e, 0xe5, 0xb3, 0x42, 0x96, 0xf3, 0x85, 0x14, 0x55,
	0x21, 0x6f, 0xef, 0xf3, 0xea, 0xae, 0x98, 0x8e, 0x10, 0xc9, 0xf0, 0x78, 0x6b, 0xed, 0xca, 0x03,
	0x32, 0xc6, 0x67, 0x5b, 0x39, 0x17, 0xe2, 0xf1, 0xa1, 0x98, 0x8e, 0x86, 0xfc, 0x07, 0xe1, 0xcb,
	0xc6, 0x5b, 0x7a, 0x48, 0x5a, 0x7e, 0xb1, 0x37, 0x95, 0xe8, 0xfe, 0x15, 0xe8, 0xf5, 0x6d, 0x65,
	0xe0, 0xbd, 0xad, 0x69, 0xe3, 0x2d, 0xeb, 0x71, 0x13, 0x65, 0xba, 0x0e, 0x75, 0x70, 0x6a, 0x3d,
	0x49, 0xcd, 0x44, 0x16, 0x43, 0xc3, 0xac, 0x32, 0x8e, 0xa5, 0x9d, 0x1d, 0x52, 0xf6, 0x2f, 0x42,
	0xdf, 0x83, 0x61, 0xbe, 0x5c, 0xd4, 0xc7, 0xe9, 0xe8, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x2d,
	0x02, 0xf2, 0xca, 0xa4, 0x01, 0x00, 0x00,
}