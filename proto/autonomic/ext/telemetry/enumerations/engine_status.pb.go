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
// source: autonomic/ext/telemetry/enumerations/engine_status.proto

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

type EngineStatus int32

const (
	EngineStatus_UNKNOWN   EngineStatus = 0
	EngineStatus_OFF       EngineStatus = 1
	EngineStatus_ON        EngineStatus = 2
	EngineStatus_AUTO_STOP EngineStatus = 3 // Deprecated: Do not use.
)

var EngineStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "OFF",
	2: "ON",
	3: "AUTO_STOP",
}

var EngineStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"OFF":       1,
	"ON":        2,
	"AUTO_STOP": 3,
}

func (x EngineStatus) String() string {
	return proto.EnumName(EngineStatus_name, int32(x))
}

func (EngineStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c360e1cfc9d6dce4, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.engineStatus.EngineStatus", EngineStatus_name, EngineStatus_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/engine_status.proto", fileDescriptor_c360e1cfc9d6dce4)
}

var fileDescriptor_c360e1cfc9d6dce4 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x60, 0xb3, 0x0b, 0xad, 0x46, 0x0f, 0xdb, 0x1c, 0x7d, 0x02, 0x11, 0x9a, 0x41, 0xbc,
	0x78, 0x93, 0x16, 0xec, 0x45, 0xd8, 0x2c, 0x74, 0x17, 0x41, 0x0f, 0x25, 0x5d, 0x86, 0x1a, 0x68,
	0x12, 0x49, 0x26, 0x50, 0x5f, 0x47, 0xf0, 0x1d, 0x3d, 0x4a, 0x53, 0x58, 0xf6, 0xd8, 0xe3, 0x7f,
	0xf8, 0xf2, 0xe7, 0x1f, 0xfe, 0xa4, 0x13, 0x79, 0xe7, 0xad, 0xe9, 0x01, 0x0f, 0x04, 0x84, 0x7b,
	0xb4, 0x48, 0xe1, 0x1b, 0xd0, 0x25, 0x8b, 0x41, 0x93, 0xf1, 0x2e, 0x02, 0xba, 0x9d, 0x71, 0xb8,
	0x89, 0xa4, 0x29, 0x45, 0xf9, 0x15, 0x3c, 0x79, 0xf1, 0x30, 0x48, 0x89, 0x07, 0x92, 0x83, 0x94,
	0x63, 0x29, 0x4f, 0x72, 0x9d, 0xe1, 0xfd, 0x33, 0xbf, 0x79, 0x19, 0x65, 0x71, 0xcd, 0xa7, 0x5d,
	0xfd, 0x5a, 0xab, 0xb7, 0xba, 0xba, 0x10, 0x53, 0x5e, 0xaa, 0xd5, 0xaa, 0x62, 0x62, 0xc2, 0x0b,
	0x55, 0x57, 0x85, 0x98, 0xf1, 0xab, 0x45, 0xd7, 0xaa, 0xcd, 0xba, 0x55, 0x4d, 0x55, 0xde, 0x16,
	0x97, 0x6c, 0xf9, 0xcb, 0xf8, 0x5d, 0xef, 0xad, 0x3c, 0xa7, 0x7a, 0x39, 0x1b, 0x77, 0x35, 0xc7,
	0x3f, 0x37, 0xec, 0xfd, 0x63, 0x67, 0xe8, 0x33, 0x6d, 0x65, 0xef, 0x2d, 0x0c, 0xaf, 0xcc, 0xb5,
	0x39, 0xae, 0xc7, 0xe0, 0xf4, 0x7e, 0x9e, 0xd7, 0x45, 0x88, 0xa1, 0x07, 0xab, 0x8d, 0x83, 0x9c,
	0xe1, 0x9c, 0x33, 0xfd, 0x31, 0xf6, 0x53, 0x94, 0x8b, 0xae, 0xdd, 0x4e, 0x32, 0x7a, 0xfc, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0x31, 0xee, 0x24, 0xcc, 0x5e, 0x01, 0x00, 0x00,
}
