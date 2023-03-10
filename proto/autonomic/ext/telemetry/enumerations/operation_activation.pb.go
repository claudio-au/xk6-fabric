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
// source: autonomic/ext/telemetry/enumerations/operation_activation.proto

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

type OperationActivation int32

const (
	OperationActivation_UNKNOWN   OperationActivation = 0
	OperationActivation_AUTOMATIC OperationActivation = 1
	OperationActivation_MANUAL    OperationActivation = 2
)

var OperationActivation_name = map[int32]string{
	0: "UNKNOWN",
	1: "AUTOMATIC",
	2: "MANUAL",
}

var OperationActivation_value = map[string]int32{
	"UNKNOWN":   0,
	"AUTOMATIC": 1,
	"MANUAL":    2,
}

func (x OperationActivation) String() string {
	return proto.EnumName(OperationActivation_name, int32(x))
}

func (OperationActivation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_99dfafc61341ef59, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.operationActivation.OperationActivation", OperationActivation_name, OperationActivation_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/operation_activation.proto", fileDescriptor_99dfafc61341ef59)
}

var fileDescriptor_99dfafc61341ef59 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0x2c, 0x2d, 0xc9,
	0xcf, 0xcb, 0xcf, 0xcd, 0x4c, 0xd6, 0x4f, 0xad, 0x28, 0xd1, 0x2f, 0x49, 0xcd, 0x49, 0xcd, 0x4d,
	0x2d, 0x29, 0xaa, 0xd4, 0x4f, 0xcd, 0x2b, 0xcd, 0x4d, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b,
	0xd6, 0xcf, 0x2f, 0x80, 0x32, 0xe3, 0x13, 0x93, 0x4b, 0x32, 0xcb, 0xc0, 0x4c, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x0b, 0xb8, 0x01, 0x7a, 0xa9, 0x15, 0x25, 0x7a, 0x70, 0x03, 0xf4, 0x90,
	0x0d, 0xd0, 0x83, 0x1b, 0xe0, 0x08, 0xd7, 0xaf, 0x65, 0xcb, 0x25, 0xec, 0x8f, 0x29, 0x2c, 0xc4,
	0xcd, 0xc5, 0x1e, 0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27, 0xc0, 0x20, 0xc4, 0xcb, 0xc5, 0xe9,
	0x18, 0x1a, 0xe2, 0xef, 0xeb, 0x18, 0xe2, 0xe9, 0x2c, 0xc0, 0x28, 0xc4, 0xc5, 0xc5, 0xe6, 0xeb,
	0xe8, 0x17, 0xea, 0xe8, 0x23, 0xc0, 0xe4, 0xb4, 0x86, 0x91, 0x4b, 0x23, 0x39, 0x3f, 0x57, 0x8f,
	0x18, 0xfb, 0x9d, 0x24, 0xb0, 0xd8, 0x14, 0x00, 0x72, 0x7f, 0x00, 0x63, 0x54, 0x74, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xdc, 0x30, 0xdd, 0xc4, 0x4c, 0x50, 0x80,
	0xa4, 0x16, 0xe5, 0x25, 0xe6, 0xe8, 0x82, 0x7d, 0x5a, 0xac, 0x5f, 0x5c, 0x94, 0xac, 0x9f, 0x9b,
	0x98, 0x99, 0xa7, 0x0f, 0xe6, 0xeb, 0x13, 0x13, 0x72, 0x3f, 0x18, 0x19, 0x17, 0x31, 0x31, 0x3b,
	0x86, 0x86, 0x24, 0xb1, 0x81, 0x35, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x9c, 0x78,
	0x80, 0x71, 0x01, 0x00, 0x00,
}
