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
// source: autonomic/ext/telemetry/enumerations/seatbelt_status.proto

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

type SeatbeltStatus int32

const (
	SeatbeltStatus_UNKNOWN   SeatbeltStatus = 0
	SeatbeltStatus_BUCKLED   SeatbeltStatus = 1
	SeatbeltStatus_UNBUCKLED SeatbeltStatus = 2
)

var SeatbeltStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "BUCKLED",
	2: "UNBUCKLED",
}

var SeatbeltStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"BUCKLED":   1,
	"UNBUCKLED": 2,
}

func (x SeatbeltStatus) String() string {
	return proto.EnumName(SeatbeltStatus_name, int32(x))
}

func (SeatbeltStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eab6cba916a16e1f, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.enumerations.seatbeltStatus.SeatbeltStatus", SeatbeltStatus_name, SeatbeltStatus_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/enumerations/seatbelt_status.proto", fileDescriptor_eab6cba916a16e1f)
}

var fileDescriptor_eab6cba916a16e1f = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xcd, 0x0a, 0x8a, 0x11, 0x65, 0xa9, 0x2f, 0x21, 0xc2, 0x66, 0x0e, 0x7b, 0xd2, 0x9b,
	0x55, 0x4f, 0x2b, 0x71, 0x61, 0x0d, 0x82, 0x1e, 0x24, 0x0d, 0x83, 0x06, 0x9a, 0x44, 0x92, 0x09,
	0xd4, 0xd7, 0xf1, 0xe2, 0x2b, 0x7a, 0x94, 0x06, 0x5a, 0xb6, 0xb7, 0x1e, 0xbf, 0x81, 0x6f, 0xfe,
	0xf9, 0x87, 0xdf, 0xe8, 0x4c, 0xc1, 0x07, 0x67, 0x0d, 0x60, 0x47, 0x40, 0xd8, 0xa2, 0x43, 0x8a,
	0xdf, 0x80, 0x3e, 0x3b, 0x8c, 0x9a, 0x6c, 0xf0, 0x09, 0x12, 0x6a, 0x6a, 0xb0, 0xa5, 0xf7, 0x44,
	0x9a, 0x72, 0x12, 0x5f, 0x31, 0x50, 0xa8, 0xd6, 0xa3, 0x2b, 0xb0, 0x23, 0x31, 0xba, 0x62, 0xdf,
	0x15, 0x83, 0xbb, 0x2b, 0xea, 0xd5, 0x35, 0x3f, 0xdf, 0x4d, 0x26, 0xd5, 0x29, 0x3f, 0x56, 0x72,
	0x23, 0x9f, 0x5e, 0xe4, 0xf2, 0xa0, 0x87, 0x5a, 0xdd, 0x6d, 0x1e, 0x1f, 0xee, 0x97, 0xac, 0x3a,
	0xe3, 0x27, 0x4a, 0x0e, 0xb8, 0xa8, 0x7f, 0x19, 0xbf, 0x34, 0xc1, 0x89, 0x39, 0xb1, 0xf5, 0xc5,
	0x34, 0x65, 0xdb, 0x5f, 0xbc, 0x65, 0xaf, 0x6f, 0x1f, 0x96, 0x3e, 0x73, 0x23, 0x4c, 0x70, 0x30,
	0xee, 0x59, 0x69, 0xdb, 0xb7, 0xc7, 0xe8, 0x75, 0xbb, 0x2a, 0xdd, 0x12, 0xa4, 0x68, 0xc0, 0x69,
	0xeb, 0xa1, 0x30, 0xcc, 0x79, 0xd3, 0x1f, 0x63, 0x3f, 0x8b, 0xc3, 0x5b, 0xf5, 0xdc, 0x1c, 0x15,
	0x69, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xe8, 0x42, 0x0c, 0x5e, 0x01, 0x00, 0x00,
}