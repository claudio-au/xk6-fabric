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
// source: autonomic/ext/telemetry/enumerations/electric_machine_status.proto

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

type ElectricMachineStatus int32 // Deprecated: Do not use.
const (
	ElectricMachineStatus_UNKNOWN            ElectricMachineStatus = 0
	ElectricMachineStatus_POWER_CONSUMPTION  ElectricMachineStatus = 1
	ElectricMachineStatus_POWER_GENERATION   ElectricMachineStatus = 2
	ElectricMachineStatus_DEACTIVATION_STATE ElectricMachineStatus = 3
	ElectricMachineStatus_READINESS_STATE    ElectricMachineStatus = 4
	ElectricMachineStatus_NO_DATA_EXISTS     ElectricMachineStatus = 5
)

var ElectricMachineStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "POWER_CONSUMPTION",
	2: "POWER_GENERATION",
	3: "DEACTIVATION_STATE",
	4: "READINESS_STATE",
	5: "NO_DATA_EXISTS",
}

var ElectricMachineStatus_value = map[string]int32{
	"UNKNOWN":            0,
	"POWER_CONSUMPTION":  1,
	"POWER_GENERATION":   2,
	"DEACTIVATION_STATE": 3,
	"READINESS_STATE":    4,
	"NO_DATA_EXISTS":     5,
}

func (x ElectricMachineStatus) String() string {
	return proto.EnumName(ElectricMachineStatus_name, int32(x))
}

func (ElectricMachineStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5419dacf2722fdb4, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.electricMachineStatus.ElectricMachineStatus", ElectricMachineStatus_name, ElectricMachineStatus_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/electric_machine_status.proto", fileDescriptor_5419dacf2722fdb4)
}

var fileDescriptor_5419dacf2722fdb4 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4a, 0x03, 0x31,
	0x14, 0xc6, 0x4d, 0xeb, 0x1f, 0x88, 0xa0, 0x63, 0xb4, 0x22, 0x3d, 0x81, 0x08, 0x4d, 0x16, 0xee,
	0xdc, 0xa5, 0x6d, 0x90, 0x41, 0x9a, 0x19, 0x9a, 0xd4, 0x8a, 0x2e, 0x86, 0x34, 0x04, 0x1b, 0x68,
	0x26, 0x32, 0x93, 0x81, 0x7a, 0x11, 0x0f, 0xe0, 0x01, 0x3c, 0x9f, 0x4b, 0x69, 0xaa, 0xc5, 0x45,
	0x17, 0x5d, 0xbe, 0xdf, 0xe3, 0xfd, 0xe0, 0xfb, 0x1e, 0xec, 0xab, 0x26, 0xf8, 0xd2, 0x3b, 0xab,
	0x89, 0x59, 0x06, 0x12, 0xcc, 0xc2, 0x38, 0x13, 0xaa, 0x77, 0x62, 0xca, 0xc6, 0x99, 0x4a, 0x05,
	0xeb, 0xcb, 0x9a, 0x98, 0x85, 0xd1, 0xa1, 0xb2, 0xba, 0x70, 0x4a, 0xcf, 0x6d, 0x69, 0x8a, 0x3a,
	0xa8, 0xd0, 0xd4, 0xf8, 0xad, 0xf2, 0xc1, 0xa3, 0xbb, 0x8d, 0x03, 0x9b, 0x65, 0xc0, 0x1b, 0x07,
	0xfe, 0xef, 0xc0, 0x7f, 0x8e, 0xd1, 0x5a, 0x21, 0xa2, 0xe1, 0xe6, 0x03, 0xc0, 0x0e, 0xdb, 0xb6,
	0x41, 0xc7, 0xf0, 0x68, 0xc2, 0x1f, 0x78, 0x36, 0xe5, 0xc9, 0x1e, 0xea, 0xc0, 0xb3, 0x3c, 0x9b,
	0xb2, 0x71, 0x31, 0xc8, 0xb8, 0x98, 0x8c, 0x72, 0x99, 0x66, 0x3c, 0x01, 0xe8, 0x02, 0x26, 0x6b,
	0x7c, 0xcf, 0x38, 0x1b, 0xd3, 0x48, 0x5b, 0xe8, 0x12, 0xa2, 0x21, 0xa3, 0x03, 0x99, 0x3e, 0x46,
	0x52, 0x08, 0x49, 0x25, 0x4b, 0xda, 0xe8, 0x1c, 0x9e, 0x8e, 0x19, 0x1d, 0xa6, 0x9c, 0x09, 0xf1,
	0x0b, 0xf7, 0x11, 0x82, 0x27, 0x3c, 0x2b, 0x86, 0x54, 0xd2, 0x82, 0x3d, 0xa5, 0x42, 0x8a, 0xe4,
	0xa0, 0xdb, 0xba, 0x02, 0xfd, 0x2f, 0x00, 0xaf, 0xb5, 0x77, 0x78, 0x97, 0x6c, 0xfd, 0xee, 0xd6,
	0x08, 0xf9, 0xaa, 0x9d, 0x1c, 0x3c, 0xbf, 0xbc, 0xda, 0x30, 0x6f, 0x66, 0x58, 0x7b, 0x47, 0x36,
	0xba, 0x9e, 0xb2, 0xab, 0xc6, 0x4d, 0x55, 0xaa, 0x45, 0x2f, 0xf6, 0x58, 0x93, 0xba, 0xd2, 0xc4,
	0x29, 0x5b, 0x92, 0x38, 0x93, 0x5d, 0x5e, 0xf3, 0x0d, 0xc0, 0x67, 0xab, 0x4d, 0x27, 0x72, 0x76,
	0x18, 0x8f, 0x6e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x87, 0x4e, 0x52, 0xd2, 0x01, 0x00,
	0x00,
}