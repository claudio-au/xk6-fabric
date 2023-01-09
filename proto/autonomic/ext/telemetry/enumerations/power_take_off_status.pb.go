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
// source: autonomic/ext/telemetry/enumerations/power_take_off_status.proto

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

type PowerTakeOffStatus int32

const (
	PowerTakeOffStatus_UNKNOWN                PowerTakeOffStatus = 0
	PowerTakeOffStatus_OFF                    PowerTakeOffStatus = 1
	PowerTakeOffStatus_BATTERY_CHARGE_PROTECT PowerTakeOffStatus = 2
	PowerTakeOffStatus_STATIONARY             PowerTakeOffStatus = 3
	PowerTakeOffStatus_MOBILE                 PowerTakeOffStatus = 4
	PowerTakeOffStatus_SPLIT_SHAFT_STATIONARY PowerTakeOffStatus = 5
)

var PowerTakeOffStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "OFF",
	2: "BATTERY_CHARGE_PROTECT",
	3: "STATIONARY",
	4: "MOBILE",
	5: "SPLIT_SHAFT_STATIONARY",
}

var PowerTakeOffStatus_value = map[string]int32{
	"UNKNOWN":                0,
	"OFF":                    1,
	"BATTERY_CHARGE_PROTECT": 2,
	"STATIONARY":             3,
	"MOBILE":                 4,
	"SPLIT_SHAFT_STATIONARY": 5,
}

func (x PowerTakeOffStatus) String() string {
	return proto.EnumName(PowerTakeOffStatus_name, int32(x))
}

func (PowerTakeOffStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8074a0e7d3d02bc6, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.powerTakeOffStatus.PowerTakeOffStatus", PowerTakeOffStatus_name, PowerTakeOffStatus_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/power_take_off_status.proto", fileDescriptor_8074a0e7d3d02bc6)
}

var fileDescriptor_8074a0e7d3d02bc6 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xed, 0xa6, 0x1b, 0x44, 0x90, 0x92, 0x83, 0x82, 0x9f, 0x40, 0x84, 0x25, 0x07, 0x0f,
	0x5e, 0x4d, 0x47, 0xe7, 0x86, 0xb3, 0x29, 0x6d, 0x86, 0x4c, 0x0f, 0x21, 0x2b, 0xa9, 0x96, 0x2d,
	0xcd, 0x48, 0x5f, 0x71, 0x5e, 0xfc, 0x30, 0x5e, 0xfd, 0x82, 0x1e, 0xa5, 0x3d, 0x94, 0xc1, 0x2e,
	0x3b, 0xbe, 0xc7, 0xfb, 0xff, 0xf8, 0xbf, 0x1f, 0x7a, 0x50, 0x35, 0xd8, 0xd2, 0x9a, 0x22, 0xa3,
	0x7a, 0x07, 0x14, 0xf4, 0x46, 0x1b, 0x0d, 0xee, 0x8b, 0xea, 0xb2, 0x36, 0xda, 0x29, 0x28, 0x6c,
	0x59, 0xd1, 0xad, 0xfd, 0xd4, 0x4e, 0x82, 0x5a, 0x6b, 0x69, 0xf3, 0x5c, 0x56, 0xa0, 0xa0, 0xae,
	0xc8, 0xd6, 0x59, 0xb0, 0xf8, 0xbe, 0x23, 0x10, 0xbd, 0x03, 0xd2, 0x11, 0xc8, 0x3e, 0x81, 0xb4,
	0x04, 0xa1, 0xd6, 0x9a, 0xe7, 0x79, 0xda, 0xc6, 0x6f, 0xbf, 0x11, 0x8e, 0x0f, 0xb6, 0xf8, 0x1c,
	0x0d, 0x17, 0xd1, 0x53, 0xc4, 0x5f, 0x22, 0xff, 0x04, 0x0f, 0x51, 0x9f, 0x4f, 0x26, 0xbe, 0x87,
	0xaf, 0xd1, 0x65, 0xc0, 0x84, 0x08, 0x93, 0xa5, 0x1c, 0x4f, 0x59, 0xf2, 0x18, 0xca, 0x38, 0xe1,
	0x22, 0x1c, 0x0b, 0xbf, 0x87, 0x2f, 0x10, 0x4a, 0x05, 0x13, 0x33, 0x1e, 0xb1, 0x64, 0xe9, 0xf7,
	0x31, 0x42, 0x83, 0x67, 0x1e, 0xcc, 0xe6, 0xa1, 0x7f, 0xda, 0xe4, 0xd2, 0x78, 0x3e, 0x13, 0x32,
	0x9d, 0xb2, 0x89, 0x90, 0x7b, 0x77, 0x67, 0xc1, 0xaf, 0x87, 0x6e, 0x32, 0x6b, 0xc8, 0x31, 0xfd,
	0x83, 0xab, 0xc3, 0xaa, 0x71, 0xf3, 0x7e, 0xec, 0xbd, 0xbe, 0xbd, 0x17, 0xf0, 0x51, 0xaf, 0x48,
	0x66, 0x0d, 0xed, 0x58, 0x23, 0x55, 0x34, 0x42, 0xb5, 0x2b, 0xd5, 0x66, 0xd4, 0x8a, 0xaa, 0x68,
	0xe5, 0x32, 0x6a, 0x54, 0x51, 0xd2, 0x76, 0xa6, 0xc7, 0x98, 0xff, 0xf3, 0xbc, 0x9f, 0x5e, 0x9f,
	0x2d, 0xc4, 0x6a, 0xd0, 0x86, 0xee, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xa0, 0xce, 0xeb,
	0xb1, 0x01, 0x00, 0x00,
}
