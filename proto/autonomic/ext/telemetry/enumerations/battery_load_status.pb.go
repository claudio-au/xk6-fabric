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
// source: autonomic/ext/telemetry/enumerations/battery_load_status.proto

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

type BatteryLoadStatus int32

const (
	BatteryLoadStatus_UNKNOWN                       BatteryLoadStatus = 0
	BatteryLoadStatus_ACCESSORY_MODE                BatteryLoadStatus = 1
	BatteryLoadStatus_NORMAL_OPERATION              BatteryLoadStatus = 2
	BatteryLoadStatus_LOW_BATTERY                   BatteryLoadStatus = 3
	BatteryLoadStatus_BATTERY_FAULT                 BatteryLoadStatus = 4
	BatteryLoadStatus_ERROR_STATE                   BatteryLoadStatus = 5
	BatteryLoadStatus_OVERLOAD_SHEDDING_IN_PROGRESS BatteryLoadStatus = 6
)

var BatteryLoadStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "ACCESSORY_MODE",
	2: "NORMAL_OPERATION",
	3: "LOW_BATTERY",
	4: "BATTERY_FAULT",
	5: "ERROR_STATE",
	6: "OVERLOAD_SHEDDING_IN_PROGRESS",
}

var BatteryLoadStatus_value = map[string]int32{
	"UNKNOWN":                       0,
	"ACCESSORY_MODE":                1,
	"NORMAL_OPERATION":              2,
	"LOW_BATTERY":                   3,
	"BATTERY_FAULT":                 4,
	"ERROR_STATE":                   5,
	"OVERLOAD_SHEDDING_IN_PROGRESS": 6,
}

func (x BatteryLoadStatus) String() string {
	return proto.EnumName(BatteryLoadStatus_name, int32(x))
}

func (BatteryLoadStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5db6130cf452a94d, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.batteryLoadStatus.BatteryLoadStatus", BatteryLoadStatus_name, BatteryLoadStatus_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/battery_load_status.proto", fileDescriptor_5db6130cf452a94d)
}

var fileDescriptor_5db6130cf452a94d = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4b, 0xfb, 0x30,
	0x14, 0xc7, 0x7f, 0xdd, 0x7e, 0x4e, 0xc8, 0x50, 0xb3, 0x20, 0xde, 0x3c, 0x78, 0x13, 0x61, 0xcd,
	0x41, 0xf0, 0x28, 0xa4, 0x6b, 0x9c, 0xc3, 0xae, 0x29, 0x49, 0xe6, 0x98, 0x1e, 0x42, 0xd6, 0x05,
	0x2d, 0xac, 0x8d, 0xb4, 0x29, 0x6c, 0xff, 0xce, 0x8e, 0xfe, 0x85, 0x1e, 0x65, 0x65, 0x0c, 0xd1,
	0xcb, 0x8e, 0xdf, 0xc7, 0xfb, 0x3c, 0xf8, 0x7e, 0x1e, 0xb8, 0xd7, 0xb5, 0xb3, 0x85, 0xcd, 0xb3,
	0x14, 0x9b, 0x95, 0xc3, 0xce, 0x2c, 0x4d, 0x6e, 0x5c, 0xb9, 0xc6, 0xa6, 0xa8, 0x73, 0x53, 0x6a,
	0x97, 0xd9, 0xa2, 0xc2, 0x73, 0xed, 0x9c, 0x29, 0xd7, 0x6a, 0x69, 0xf5, 0x42, 0x55, 0x4e, 0xbb,
	0xba, 0xf2, 0x3f, 0x4a, 0xeb, 0x2c, 0xba, 0xdb, 0xf3, 0xbe, 0x59, 0x39, 0x7f, 0xcf, 0xfb, 0x3f,
	0x79, 0x7f, 0xc7, 0x47, 0x56, 0x2f, 0x44, 0x43, 0xdf, 0x6c, 0x3c, 0xd0, 0x0b, 0x7e, 0x4f, 0x51,
	0x17, 0x1c, 0x4f, 0xe2, 0xa7, 0x98, 0x4d, 0x63, 0xf8, 0x0f, 0x21, 0x70, 0x4a, 0x06, 0x03, 0x2a,
	0x04, 0xe3, 0x33, 0x35, 0x66, 0x21, 0x85, 0x1e, 0x3a, 0x07, 0x30, 0x66, 0x7c, 0x4c, 0x22, 0xc5,
	0x12, 0xca, 0x89, 0x1c, 0xb1, 0x18, 0xb6, 0xd0, 0x19, 0xe8, 0x46, 0x6c, 0xaa, 0x02, 0x22, 0x25,
	0xe5, 0x33, 0xd8, 0x46, 0x3d, 0x70, 0xb2, 0x0b, 0xea, 0x81, 0x4c, 0x22, 0x09, 0xff, 0x6f, 0x77,
	0x28, 0xe7, 0x8c, 0x2b, 0x21, 0x89, 0xa4, 0xf0, 0x08, 0x5d, 0x81, 0x4b, 0xf6, 0x4c, 0x79, 0xc4,
	0x48, 0xa8, 0xc4, 0x23, 0x0d, 0xc3, 0x51, 0x3c, 0x54, 0xa3, 0x58, 0x25, 0x9c, 0x0d, 0x39, 0x15,
	0x02, 0x76, 0x82, 0x4f, 0x0f, 0x5c, 0xa7, 0x36, 0xf7, 0x0f, 0xe9, 0x18, 0x5c, 0xfc, 0xa9, 0x93,
	0x6c, 0x0d, 0x25, 0xde, 0xcb, 0xeb, 0x5b, 0xe6, 0xde, 0xeb, 0xb9, 0x9f, 0xda, 0x1c, 0xef, 0x4f,
	0xf5, 0x75, 0xb6, 0x35, 0x6e, 0xca, 0x42, 0x2f, 0xfb, 0x8d, 0xcb, 0x0a, 0x57, 0x65, 0x8a, 0x73,
	0x9d, 0x15, 0xb8, 0xc9, 0xf8, 0x90, 0xd7, 0x7c, 0x79, 0xde, 0xa6, 0xd5, 0x26, 0x13, 0x39, 0xef,
	0x34, 0xd0, 0xed, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x43, 0x06, 0xf3, 0xd2, 0x01, 0x00,
	0x00,
}
