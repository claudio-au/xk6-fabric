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
// source: autonomic/ext/telemetry/enumerations/hybrid_vehicle_mode_status.proto

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

type HybridVehicleModeStatus int32

const (
	HybridVehicleModeStatus_UNKNOWN                      HybridVehicleModeStatus = 0
	HybridVehicleModeStatus_AUTO_CHARGE_DEPLETE          HybridVehicleModeStatus = 1
	HybridVehicleModeStatus_AUTO_CHARGE_SUSTAIN          HybridVehicleModeStatus = 2
	HybridVehicleModeStatus_FORCED_CHARGE_SUSTAIN        HybridVehicleModeStatus = 3
	HybridVehicleModeStatus_FORCED_ELECTRIC              HybridVehicleModeStatus = 4
	HybridVehicleModeStatus_FORCED_NON_ELECTRIC          HybridVehicleModeStatus = 5
	HybridVehicleModeStatus_TEMPORARY_CHARGE_SUSTAIN     HybridVehicleModeStatus = 6
	HybridVehicleModeStatus_PRIORITIZE_CHARGE_GENERATION HybridVehicleModeStatus = 7
)

var HybridVehicleModeStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "AUTO_CHARGE_DEPLETE",
	2: "AUTO_CHARGE_SUSTAIN",
	3: "FORCED_CHARGE_SUSTAIN",
	4: "FORCED_ELECTRIC",
	5: "FORCED_NON_ELECTRIC",
	6: "TEMPORARY_CHARGE_SUSTAIN",
	7: "PRIORITIZE_CHARGE_GENERATION",
}

var HybridVehicleModeStatus_value = map[string]int32{
	"UNKNOWN":                      0,
	"AUTO_CHARGE_DEPLETE":          1,
	"AUTO_CHARGE_SUSTAIN":          2,
	"FORCED_CHARGE_SUSTAIN":        3,
	"FORCED_ELECTRIC":              4,
	"FORCED_NON_ELECTRIC":          5,
	"TEMPORARY_CHARGE_SUSTAIN":     6,
	"PRIORITIZE_CHARGE_GENERATION": 7,
}

func (x HybridVehicleModeStatus) String() string {
	return proto.EnumName(HybridVehicleModeStatus_name, int32(x))
}

func (HybridVehicleModeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7f18778138faf622, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.hybridVehicleModeStatus.HybridVehicleModeStatus", HybridVehicleModeStatus_name, HybridVehicleModeStatus_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/hybrid_vehicle_mode_status.proto", fileDescriptor_7f18778138faf622)
}

var fileDescriptor_7f18778138faf622 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xc1, 0x4e, 0xc2, 0x30,
	0x1c, 0xc7, 0x71, 0x07, 0x0a, 0x49, 0x3d, 0xb8, 0x8c, 0x18, 0x34, 0xe1, 0xe0, 0xd1, 0x98, 0xd0,
	0x1e, 0xbc, 0x7a, 0x19, 0xa3, 0xc2, 0x22, 0xb4, 0x4b, 0xe9, 0x34, 0xe2, 0x61, 0x19, 0xa3, 0x91,
	0x26, 0x74, 0x35, 0x5b, 0x67, 0xe0, 0x75, 0x7c, 0x02, 0x5f, 0xc9, 0xb7, 0xf0, 0x68, 0x18, 0x88,
	0x86, 0x68, 0xc2, 0x71, 0xfb, 0xa6, 0x9f, 0xc3, 0xff, 0x07, 0x70, 0x5c, 0x18, 0x9d, 0x6a, 0x25,
	0x13, 0x24, 0x16, 0x06, 0x19, 0x31, 0x17, 0x4a, 0x98, 0x6c, 0x89, 0x44, 0x5a, 0x28, 0x91, 0xc5,
	0x46, 0xea, 0x34, 0x47, 0xb3, 0xe5, 0x24, 0x93, 0xd3, 0xe8, 0x55, 0xcc, 0x64, 0x32, 0x17, 0x91,
	0xd2, 0x53, 0x11, 0xe5, 0x26, 0x36, 0x45, 0x0e, 0x5f, 0x32, 0x6d, 0xb4, 0x73, 0xb3, 0x65, 0xa0,
	0x58, 0x18, 0xb8, 0x65, 0xe0, 0x6f, 0x06, 0xae, 0x99, 0xfb, 0xb5, 0x32, 0xd4, 0x53, 0x31, 0x2a,
	0x8d, 0xab, 0x0f, 0x0b, 0x34, 0xfb, 0x7f, 0x37, 0xe7, 0x18, 0xd4, 0x43, 0x72, 0x47, 0xe8, 0x03,
	0xb1, 0x0f, 0x9c, 0x26, 0x68, 0xb8, 0x21, 0xa7, 0x91, 0xd7, 0x77, 0x59, 0x0f, 0x47, 0x5d, 0x1c,
	0x0c, 0x30, 0xc7, 0xb6, 0xb5, 0x1b, 0x46, 0xe1, 0x88, 0xbb, 0x3e, 0xb1, 0x2b, 0xce, 0x39, 0x38,
	0xbd, 0xa5, 0xcc, 0xc3, 0xdd, 0xdd, 0x54, 0x75, 0x1a, 0xe0, 0x64, 0x93, 0xf0, 0x00, 0x7b, 0x9c,
	0xf9, 0x9e, 0x7d, 0xb8, 0x82, 0x36, 0x3f, 0x09, 0x25, 0x3f, 0xe1, 0xc8, 0x69, 0x81, 0x33, 0x8e,
	0x87, 0x01, 0x65, 0x2e, 0x7b, 0xdc, 0xb5, 0x6a, 0xce, 0x05, 0x68, 0x05, 0xcc, 0xa7, 0xcc, 0xe7,
	0xfe, 0x18, 0x7f, 0xe7, 0x1e, 0x26, 0x98, 0xb9, 0xdc, 0xa7, 0xc4, 0xae, 0x77, 0xde, 0x2d, 0x70,
	0x99, 0x68, 0x05, 0xf7, 0x39, 0x54, 0xa7, 0xf5, 0xcf, 0x35, 0x82, 0xd5, 0xb1, 0x03, 0x6b, 0xfc,
	0xf4, 0x2c, 0xcd, 0xac, 0x98, 0xc0, 0x44, 0x2b, 0xb4, 0x05, 0xdb, 0xb1, 0x5c, 0x6d, 0x28, 0xb2,
	0x34, 0x9e, 0xb7, 0xcb, 0x59, 0x72, 0x94, 0x67, 0x09, 0x52, 0xb1, 0x4c, 0x51, 0xf9, 0x8d, 0xf6,
	0x19, 0xfb, 0xd3, 0xb2, 0xde, 0x2a, 0x55, 0x37, 0xe4, 0x93, 0x5a, 0xf9, 0xe8, 0xfa, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0xae, 0x9a, 0xea, 0xc7, 0x24, 0x02, 0x00, 0x00,
}