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
// source: autonomic/ext/telemetry/enumerations/tire_pressure_status.proto

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

type TirePressureStatus int32

const (
	TirePressureStatus_UNKNOWN     TirePressureStatus = 0
	TirePressureStatus_NORMAL      TirePressureStatus = 1
	TirePressureStatus_LOW         TirePressureStatus = 2
	TirePressureStatus_ALERT       TirePressureStatus = 3
	TirePressureStatus_UNSUPPORTED TirePressureStatus = 4
)

var TirePressureStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "NORMAL",
	2: "LOW",
	3: "ALERT",
	4: "UNSUPPORTED",
}

var TirePressureStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"NORMAL":      1,
	"LOW":         2,
	"ALERT":       3,
	"UNSUPPORTED": 4,
}

func (x TirePressureStatus) String() string {
	return proto.EnumName(TirePressureStatus_name, int32(x))
}

func (TirePressureStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb8e231fed8bbb90, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.tirePressureStatus.TirePressureStatus", TirePressureStatus_name, TirePressureStatus_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/tire_pressure_status.proto", fileDescriptor_cb8e231fed8bbb90)
}

var fileDescriptor_cb8e231fed8bbb90 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0xbf, 0xed, 0x7e, 0xb6, 0x98, 0x1e, 0x0c, 0xb9, 0xf8, 0x17, 0x44, 0x68, 0x72, 0xf0,
	0xe0, 0x51, 0xb6, 0xd8, 0x93, 0x6b, 0x36, 0x6c, 0x77, 0x29, 0xe8, 0xa1, 0xa4, 0xcb, 0xa0, 0x81,
	0x26, 0x29, 0xc9, 0x04, 0xea, 0xdf, 0xf1, 0xea, 0x1f, 0xf4, 0x28, 0xbb, 0xc8, 0x22, 0xf4, 0xd2,
	0xe3, 0x0c, 0x3c, 0xcf, 0xcc, 0xfb, 0x92, 0x07, 0x9d, 0xd0, 0x3b, 0x6f, 0x4d, 0x27, 0xe0, 0x88,
	0x02, 0x61, 0x0f, 0x16, 0x30, 0x7c, 0x08, 0x70, 0xc9, 0x42, 0xd0, 0x68, 0xbc, 0x8b, 0x02, 0x4d,
	0x80, 0xed, 0x21, 0x40, 0x8c, 0x29, 0xc0, 0x36, 0xa2, 0xc6, 0x14, 0xf9, 0x21, 0x78, 0xf4, 0xec,
	0x7e, 0x14, 0x70, 0x38, 0x22, 0x1f, 0x05, 0xfc, 0xaf, 0x80, 0xf7, 0x02, 0xf5, 0xcb, 0xaf, 0x07,
	0xfc, 0xb6, 0x26, 0xac, 0x39, 0xd9, 0xb2, 0x39, 0x99, 0xb5, 0xf2, 0x49, 0x56, 0x1b, 0x49, 0xff,
	0x31, 0x42, 0xa6, 0xb2, 0xaa, 0x9f, 0x8b, 0x92, 0x66, 0x6c, 0x46, 0xf2, 0xb2, 0xda, 0xd0, 0x09,
	0xbb, 0x24, 0x17, 0x45, 0xb9, 0xaa, 0x1b, 0x9a, 0xb3, 0x2b, 0x32, 0x6f, 0xe5, 0xba, 0x55, 0xaa,
	0xaa, 0x9b, 0xd5, 0x23, 0xfd, 0xbf, 0xfc, 0xca, 0xc8, 0x4d, 0xe7, 0x2d, 0x3f, 0xe7, 0xa7, 0xe5,
	0xf5, 0xe9, 0x79, 0xd5, 0x47, 0x52, 0xd9, 0xcb, 0xeb, 0x9b, 0xc1, 0xf7, 0xb4, 0xe3, 0x9d, 0xb7,
	0x62, 0x74, 0x2d, 0xb4, 0xe9, 0x3b, 0x82, 0xe0, 0xf4, 0x7e, 0x31, 0x84, 0x8f, 0x22, 0x86, 0x4e,
	0x58, 0x6d, 0x9c, 0x18, 0x66, 0x71, 0x4e, 0x99, 0xdf, 0x59, 0xf6, 0x39, 0xc9, 0x8b, 0xb6, 0xd9,
	0x4d, 0x07, 0xe8, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x55, 0xf1, 0xcc, 0x84, 0x01, 0x00,
	0x00,
}
