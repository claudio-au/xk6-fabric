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
// source: autonomic/ext/telemetry/enumerations/gear_position.proto

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

type GearPosition int32

const (
	GearPosition_UNKNOWN GearPosition = 0
	GearPosition_PARK    GearPosition = 1
	GearPosition_REVERSE GearPosition = 2
	GearPosition_NEUTRAL GearPosition = 3
	GearPosition_DRIVE   GearPosition = 4
	GearPosition_SPORT   GearPosition = 5
	GearPosition_LOW     GearPosition = 6
	// these were in use for a *very* short time,
	// but we want to leave a gap before the numbered gears.
	// Can be removed once all users are using the new values
	// and any existing messages have aged out of the system.
	GearPosition_NOT_GEAR_1 GearPosition = 7  // Deprecated: Do not use.
	GearPosition_NOT_GEAR_2 GearPosition = 8  // Deprecated: Do not use.
	GearPosition_NOT_GEAR_3 GearPosition = 9  // Deprecated: Do not use.
	GearPosition_NOT_GEAR_4 GearPosition = 10 // Deprecated: Do not use.
	GearPosition_NOT_GEAR_5 GearPosition = 11 // Deprecated: Do not use.
	GearPosition_NOT_GEAR_6 GearPosition = 12 // Deprecated: Do not use.
	// Open ended assignments for numbered positions.
	// Big trucks have a *lot* of gears.
	GearPosition_GEAR_1  GearPosition = 21
	GearPosition_GEAR_2  GearPosition = 22
	GearPosition_GEAR_3  GearPosition = 23
	GearPosition_GEAR_4  GearPosition = 24
	GearPosition_GEAR_5  GearPosition = 25
	GearPosition_GEAR_6  GearPosition = 26
	GearPosition_GEAR_7  GearPosition = 27
	GearPosition_GEAR_8  GearPosition = 28
	GearPosition_GEAR_9  GearPosition = 29
	GearPosition_GEAR_10 GearPosition = 30
	GearPosition_GEAR_11 GearPosition = 31
	GearPosition_GEAR_12 GearPosition = 32
	GearPosition_GEAR_13 GearPosition = 33
	GearPosition_GEAR_14 GearPosition = 34
	GearPosition_GEAR_15 GearPosition = 35
	GearPosition_GEAR_16 GearPosition = 36
	GearPosition_GEAR_17 GearPosition = 37
	GearPosition_GEAR_18 GearPosition = 38
	GearPosition_GEAR_19 GearPosition = 39
	GearPosition_GEAR_20 GearPosition = 40
	GearPosition_GEAR_21 GearPosition = 41
	GearPosition_GEAR_22 GearPosition = 42
	GearPosition_GEAR_23 GearPosition = 43
	GearPosition_GEAR_24 GearPosition = 44
)

var GearPosition_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "PARK",
	2:  "REVERSE",
	3:  "NEUTRAL",
	4:  "DRIVE",
	5:  "SPORT",
	6:  "LOW",
	7:  "NOT_GEAR_1",
	8:  "NOT_GEAR_2",
	9:  "NOT_GEAR_3",
	10: "NOT_GEAR_4",
	11: "NOT_GEAR_5",
	12: "NOT_GEAR_6",
	21: "GEAR_1",
	22: "GEAR_2",
	23: "GEAR_3",
	24: "GEAR_4",
	25: "GEAR_5",
	26: "GEAR_6",
	27: "GEAR_7",
	28: "GEAR_8",
	29: "GEAR_9",
	30: "GEAR_10",
	31: "GEAR_11",
	32: "GEAR_12",
	33: "GEAR_13",
	34: "GEAR_14",
	35: "GEAR_15",
	36: "GEAR_16",
	37: "GEAR_17",
	38: "GEAR_18",
	39: "GEAR_19",
	40: "GEAR_20",
	41: "GEAR_21",
	42: "GEAR_22",
	43: "GEAR_23",
	44: "GEAR_24",
}

var GearPosition_value = map[string]int32{
	"UNKNOWN":    0,
	"PARK":       1,
	"REVERSE":    2,
	"NEUTRAL":    3,
	"DRIVE":      4,
	"SPORT":      5,
	"LOW":        6,
	"NOT_GEAR_1": 7,
	"NOT_GEAR_2": 8,
	"NOT_GEAR_3": 9,
	"NOT_GEAR_4": 10,
	"NOT_GEAR_5": 11,
	"NOT_GEAR_6": 12,
	"GEAR_1":     21,
	"GEAR_2":     22,
	"GEAR_3":     23,
	"GEAR_4":     24,
	"GEAR_5":     25,
	"GEAR_6":     26,
	"GEAR_7":     27,
	"GEAR_8":     28,
	"GEAR_9":     29,
	"GEAR_10":    30,
	"GEAR_11":    31,
	"GEAR_12":    32,
	"GEAR_13":    33,
	"GEAR_14":    34,
	"GEAR_15":    35,
	"GEAR_16":    36,
	"GEAR_17":    37,
	"GEAR_18":    38,
	"GEAR_19":    39,
	"GEAR_20":    40,
	"GEAR_21":    41,
	"GEAR_22":    42,
	"GEAR_23":    43,
	"GEAR_24":    44,
}

func (x GearPosition) String() string {
	return proto.EnumName(GearPosition_name, int32(x))
}

func (GearPosition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a20503ffa594f191, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.gearPosition.GearPosition", GearPosition_name, GearPosition_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/gear_position.proto", fileDescriptor_a20503ffa594f191)
}

var fileDescriptor_a20503ffa594f191 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcb, 0x6a, 0xdb, 0x40,
	0x18, 0x85, 0x2b, 0xdb, 0xb1, 0x9d, 0x49, 0x16, 0xa7, 0x82, 0xde, 0xd2, 0xfb, 0x3d, 0x4d, 0x1b,
	0x4d, 0x2c, 0xc9, 0x8e, 0xb3, 0x54, 0xa8, 0x08, 0x25, 0x41, 0x12, 0x8a, 0x9c, 0x40, 0xbb, 0x30,
	0x8a, 0x18, 0x52, 0x41, 0xa4, 0x09, 0xd2, 0x18, 0xd2, 0x07, 0xe8, 0x8b, 0x14, 0xfa, 0x8e, 0x5d,
	0x96, 0x51, 0xc3, 0x8c, 0x8b, 0x36, 0xde, 0x7d, 0xdf, 0x81, 0x33, 0x3f, 0x0c, 0x87, 0x4c, 0xd3,
	0x85, 0xe0, 0x25, 0x2f, 0xf2, 0x8c, 0xb2, 0x1b, 0x41, 0x05, 0xbb, 0x62, 0x05, 0x13, 0xd5, 0x0f,
	0xca, 0xca, 0x45, 0xc1, 0xaa, 0x54, 0xe4, 0xbc, 0xac, 0xe9, 0x25, 0x4b, 0xab, 0xf9, 0x35, 0xaf,
	0x73, 0xa9, 0xd6, 0x75, 0xc5, 0x05, 0x37, 0x47, 0xaa, 0x69, 0xb1, 0x1b, 0x61, 0xa9, 0xa6, 0xb5,
	0xdc, 0xb4, 0x64, 0x33, 0xba, 0x2d, 0xee, 0xfc, 0xec, 0x91, 0xcd, 0xa3, 0xa5, 0xc0, 0xdc, 0x20,
	0x83, 0x59, 0x70, 0x1c, 0x84, 0xe7, 0x01, 0xee, 0x98, 0x43, 0xd2, 0x8b, 0xbc, 0xf8, 0x18, 0x86,
	0x8c, 0x63, 0xff, 0xcc, 0x8f, 0x4f, 0x7d, 0x74, 0xa4, 0x04, 0xfe, 0x2c, 0x89, 0xbd, 0x13, 0x74,
	0xcd, 0x75, 0xb2, 0xf6, 0x39, 0xfe, 0x72, 0xe6, 0xa3, 0x27, 0xf1, 0x34, 0x0a, 0xe3, 0x04, 0x6b,
	0xe6, 0x80, 0x74, 0x4f, 0xc2, 0x73, 0xf4, 0x4d, 0x93, 0x90, 0x20, 0x4c, 0xe6, 0x47, 0xbe, 0x17,
	0xcf, 0x47, 0x18, 0x6c, 0x75, 0x86, 0xc6, 0x7f, 0x99, 0x8d, 0x61, 0x2b, 0x73, 0xb0, 0xde, 0xca,
	0x5c, 0x90, 0x56, 0x36, 0xc6, 0x46, 0x2b, 0x9b, 0x60, 0xb3, 0xc9, 0x08, 0xe9, 0xdf, 0xde, 0xbc,
	0xa7, 0xd8, 0xc6, 0x7d, 0xc5, 0x0e, 0x1e, 0x28, 0x76, 0xf1, 0x50, 0xf1, 0x18, 0x8f, 0x14, 0x4f,
	0xb0, 0xa5, 0x78, 0x1f, 0x8f, 0x15, 0x4f, 0xf1, 0x44, 0xf1, 0x01, 0x9e, 0xca, 0xff, 0xf8, 0x77,
	0x6b, 0x0f, 0xcf, 0xb4, 0x8c, 0xf0, 0x5c, 0x8b, 0x8d, 0x17, 0x5a, 0x1c, 0xbc, 0xd4, 0xe2, 0xe2,
	0x95, 0x96, 0x31, 0x5e, 0x6b, 0x99, 0xe0, 0x8d, 0x96, 0x7d, 0xbc, 0xd5, 0x32, 0xc5, 0x3b, 0x2d,
	0x07, 0x78, 0xaf, 0xc4, 0xde, 0xc3, 0xb6, 0x96, 0x11, 0x3e, 0x68, 0xb1, 0xb1, 0xa3, 0xc5, 0xc1,
	0x47, 0x2d, 0x2e, 0x3e, 0x1d, 0xfe, 0x36, 0xc8, 0x76, 0xc6, 0x0b, 0x6b, 0x95, 0x05, 0x1d, 0xde,
	0x5d, 0x5e, 0x4c, 0x24, 0xa7, 0x17, 0x19, 0x5f, 0xbf, 0x5d, 0xe6, 0xe2, 0xfb, 0xe2, 0xc2, 0xca,
	0x78, 0x41, 0xd5, 0x2b, 0xbb, 0x69, 0x2e, 0x47, 0xcc, 0xaa, 0x32, 0xbd, 0xda, 0x6d, 0x46, 0x5a,
	0xd3, 0xba, 0xca, 0x68, 0x91, 0xe6, 0x25, 0x6d, 0x9c, 0xae, 0xb2, 0xf6, 0x3f, 0x86, 0xf1, 0xab,
	0xd3, 0xf5, 0x66, 0xc9, 0x45, 0xbf, 0x29, 0x39, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x49, 0x17,
	0x76, 0x8d, 0x25, 0x03, 0x00, 0x00,
}
