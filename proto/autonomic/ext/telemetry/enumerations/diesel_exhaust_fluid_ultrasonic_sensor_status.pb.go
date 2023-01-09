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
// autonomic/ext/telemetry/enumerations/diesel_exhaust_fluid_ultrasonic_sensor_status.proto is a deprecated file.

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

type DieselExhaustFluidUltrasonicSensorStatus int32

const (
	DieselExhaustFluidUltrasonicSensorStatus_UNKNOWN     DieselExhaustFluidUltrasonicSensorStatus = 0
	DieselExhaustFluidUltrasonicSensorStatus_PRESENT     DieselExhaustFluidUltrasonicSensorStatus = 1
	DieselExhaustFluidUltrasonicSensorStatus_NOT_PRESENT DieselExhaustFluidUltrasonicSensorStatus = 2
)

var DieselExhaustFluidUltrasonicSensorStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "PRESENT",
	2: "NOT_PRESENT",
}

var DieselExhaustFluidUltrasonicSensorStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"PRESENT":     1,
	"NOT_PRESENT": 2,
}

func (x DieselExhaustFluidUltrasonicSensorStatus) String() string {
	return proto.EnumName(DieselExhaustFluidUltrasonicSensorStatus_name, int32(x))
}

func (DieselExhaustFluidUltrasonicSensorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3cb7be976f25a3c, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.dieselExhaustFluidUltrasonicSensorStatus.DieselExhaustFluidUltrasonicSensorStatus", DieselExhaustFluidUltrasonicSensorStatus_name, DieselExhaustFluidUltrasonicSensorStatus_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/diesel_exhaust_fluid_ultrasonic_sensor_status.proto", fileDescriptor_e3cb7be976f25a3c)
}

var fileDescriptor_e3cb7be976f25a3c = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0x4f, 0x4a, 0xc4, 0x30,
	0x14, 0x06, 0x70, 0x33, 0x82, 0x42, 0x67, 0xe1, 0xd0, 0x53, 0x0c, 0x42, 0x93, 0x85, 0x27, 0x70,
	0xb0, 0x6e, 0xc4, 0x4e, 0xb1, 0x2d, 0x8a, 0x2e, 0x42, 0xa6, 0xf3, 0x74, 0x02, 0xf9, 0x23, 0x79,
	0x2f, 0x50, 0xaf, 0xe3, 0x29, 0x3c, 0x80, 0x87, 0x72, 0x29, 0xcd, 0x60, 0x71, 0xd9, 0xe5, 0x0b,
	0xc9, 0x2f, 0x1f, 0xdf, 0xcb, 0x9e, 0x54, 0x24, 0xef, 0xbc, 0xd5, 0xbd, 0x80, 0x81, 0x04, 0x81,
	0x01, 0x0b, 0x14, 0x3e, 0x04, 0xb8, 0x68, 0x21, 0x28, 0xd2, 0xde, 0xa1, 0xd8, 0x6b, 0x40, 0x30,
	0x12, 0x86, 0x83, 0x8a, 0x48, 0xf2, 0xd5, 0x44, 0xbd, 0x97, 0xd1, 0x50, 0x50, 0xe8, 0x9d, 0xee,
	0x25, 0x82, 0x43, 0x1f, 0x24, 0x92, 0xa2, 0x88, 0xfc, 0x3d, 0x78, 0xf2, 0xf9, 0xfd, 0x24, 0x73,
	0x18, 0x88, 0x4f, 0x32, 0xff, 0x2f, 0xf3, 0xa3, 0x5c, 0x1e, 0xe1, 0xdb, 0xd1, 0xed, 0x26, 0xb6,
	0x49, 0x6a, 0x93, 0xd0, 0xcb, 0x2e, 0x5b, 0xdf, 0xcc, 0xbc, 0x9b, 0x2f, 0xb3, 0xf3, 0xae, 0xba,
	0xab, 0xb6, 0x8f, 0xd5, 0xea, 0x64, 0x1c, 0xea, 0x87, 0xb2, 0x29, 0xab, 0x76, 0xc5, 0xf2, 0x8b,
	0x6c, 0x59, 0x6d, 0x5b, 0xf9, 0x77, 0xb0, 0xd8, 0x7c, 0xb3, 0x6c, 0xdd, 0x7b, 0xcb, 0xe7, 0x84,
	0xdd, 0x14, 0x73, 0x13, 0xd4, 0x63, 0x03, 0x35, 0x7b, 0x7e, 0x79, 0xd3, 0x74, 0x88, 0x3b, 0xde,
	0x7b, 0x2b, 0xa6, 0x1f, 0x0a, 0xa5, 0xc7, 0xae, 0x21, 0x38, 0x65, 0x8a, 0xd4, 0x15, 0x0a, 0x0c,
	0xbd, 0xb0, 0x4a, 0x3b, 0x91, 0x66, 0x31, 0x67, 0x29, 0x5f, 0x8c, 0xfd, 0x30, 0xf6, 0xb9, 0x38,
	0xbd, 0xee, 0xda, 0xdd, 0x59, 0x7a, 0x77, 0xf5, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xd2, 0x23,
	0x85, 0xcf, 0x01, 0x00, 0x00,
}