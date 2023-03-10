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
// source: autonomic/ext/telemetry/configurations/vehicle_control_policy.proto

package configurations

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

type VehicleControlPolicy_ControlPolicyScope int32

const (
	VehicleControlPolicy_UNKNOWN_CONTROL_POLICY_SCOPE VehicleControlPolicy_ControlPolicyScope = 0
	// Scope to reset all settings to factory settings
	VehicleControlPolicy_FACTORY_RESET VehicleControlPolicy_ControlPolicyScope = 1
	// Scope to reset selected module parameters to factory settings. Specified module parameters or components,
	// can be passed through `oem_params` fields
	VehicleControlPolicy_SELECTIVE_RESET VehicleControlPolicy_ControlPolicyScope = 2
)

var VehicleControlPolicy_ControlPolicyScope_name = map[int32]string{
	0: "UNKNOWN_CONTROL_POLICY_SCOPE",
	1: "FACTORY_RESET",
	2: "SELECTIVE_RESET",
}

var VehicleControlPolicy_ControlPolicyScope_value = map[string]int32{
	"UNKNOWN_CONTROL_POLICY_SCOPE": 0,
	"FACTORY_RESET":                1,
	"SELECTIVE_RESET":              2,
}

func (x VehicleControlPolicy_ControlPolicyScope) String() string {
	return proto.EnumName(VehicleControlPolicy_ControlPolicyScope_name, int32(x))
}

func (VehicleControlPolicy_ControlPolicyScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cf5e26f66b45d0ca, []int{0, 0}
}

type VehicleControlPolicy_ControlPolicyPermission int32

const (
	VehicleControlPolicy_UNKNOWN_CONTROL_POLICY_PERMISSION VehicleControlPolicy_ControlPolicyPermission = 0
	VehicleControlPolicy_ALLOWED                           VehicleControlPolicy_ControlPolicyPermission = 1
	VehicleControlPolicy_PROHIBITED                        VehicleControlPolicy_ControlPolicyPermission = 2
	VehicleControlPolicy_REMOTE_ACCESS_ONLY                VehicleControlPolicy_ControlPolicyPermission = 3
)

var VehicleControlPolicy_ControlPolicyPermission_name = map[int32]string{
	0: "UNKNOWN_CONTROL_POLICY_PERMISSION",
	1: "ALLOWED",
	2: "PROHIBITED",
	3: "REMOTE_ACCESS_ONLY",
}

var VehicleControlPolicy_ControlPolicyPermission_value = map[string]int32{
	"UNKNOWN_CONTROL_POLICY_PERMISSION": 0,
	"ALLOWED":                           1,
	"PROHIBITED":                        2,
	"REMOTE_ACCESS_ONLY":                3,
}

func (x VehicleControlPolicy_ControlPolicyPermission) String() string {
	return proto.EnumName(VehicleControlPolicy_ControlPolicyPermission_name, int32(x))
}

func (VehicleControlPolicy_ControlPolicyPermission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cf5e26f66b45d0ca, []int{0, 1}
}

type VehicleControlPolicy struct {
	PolicyScope      VehicleControlPolicy_ControlPolicyScope      `protobuf:"varint,1,opt,name=policy_scope,json=policyScope,proto3,enum=autonomic.ext.telemetry.configurations.VehicleControlPolicy_ControlPolicyScope" json:"policy_scope,omitempty"`
	PolicyPermission VehicleControlPolicy_ControlPolicyPermission `protobuf:"varint,2,opt,name=policy_permission,json=policyPermission,proto3,enum=autonomic.ext.telemetry.configurations.VehicleControlPolicy_ControlPolicyPermission" json:"policy_permission,omitempty"`
	// The amount of time in seconds the policy change should be in effect. Unset or 0 makes it an indefinite change.
	RemainingDurationSeconds int32    `protobuf:"varint,3,opt,name=remaining_duration_seconds,json=remainingDurationSeconds,proto3" json:"remaining_duration_seconds,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *VehicleControlPolicy) Reset()         { *m = VehicleControlPolicy{} }
func (m *VehicleControlPolicy) String() string { return proto.CompactTextString(m) }
func (*VehicleControlPolicy) ProtoMessage()    {}
func (*VehicleControlPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5e26f66b45d0ca, []int{0}
}

func (m *VehicleControlPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehicleControlPolicy.Unmarshal(m, b)
}
func (m *VehicleControlPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehicleControlPolicy.Marshal(b, m, deterministic)
}
func (m *VehicleControlPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehicleControlPolicy.Merge(m, src)
}
func (m *VehicleControlPolicy) XXX_Size() int {
	return xxx_messageInfo_VehicleControlPolicy.Size(m)
}
func (m *VehicleControlPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_VehicleControlPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_VehicleControlPolicy proto.InternalMessageInfo

func (m *VehicleControlPolicy) GetPolicyScope() VehicleControlPolicy_ControlPolicyScope {
	if m != nil {
		return m.PolicyScope
	}
	return VehicleControlPolicy_UNKNOWN_CONTROL_POLICY_SCOPE
}

func (m *VehicleControlPolicy) GetPolicyPermission() VehicleControlPolicy_ControlPolicyPermission {
	if m != nil {
		return m.PolicyPermission
	}
	return VehicleControlPolicy_UNKNOWN_CONTROL_POLICY_PERMISSION
}

func (m *VehicleControlPolicy) GetRemainingDurationSeconds() int32 {
	if m != nil {
		return m.RemainingDurationSeconds
	}
	return 0
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.configurations.VehicleControlPolicy_ControlPolicyScope", VehicleControlPolicy_ControlPolicyScope_name, VehicleControlPolicy_ControlPolicyScope_value)
	proto.RegisterEnum("autonomic.ext.telemetry.configurations.VehicleControlPolicy_ControlPolicyPermission", VehicleControlPolicy_ControlPolicyPermission_name, VehicleControlPolicy_ControlPolicyPermission_value)
	proto.RegisterType((*VehicleControlPolicy)(nil), "autonomic.ext.telemetry.configurations.VehicleControlPolicy")
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/configurations/vehicle_control_policy.proto", fileDescriptor_cf5e26f66b45d0ca)
}

var fileDescriptor_cf5e26f66b45d0ca = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe7, 0x96, 0x3f, 0x92, 0x07, 0x23, 0x33, 0x08, 0x0a, 0xe2, 0x50, 0x2a, 0x81, 0x26,
	0xa4, 0x25, 0x12, 0x5c, 0xb9, 0xb4, 0xae, 0x11, 0x11, 0x59, 0x1c, 0xd9, 0xd9, 0xa6, 0x22, 0x81,
	0x95, 0x79, 0xa6, 0xb3, 0x94, 0xd8, 0x91, 0xe3, 0xa2, 0xed, 0xc8, 0x57, 0x81, 0x4f, 0xc0, 0xb7,
	0xe3, 0x88, 0x9a, 0x56, 0x41, 0x83, 0x4d, 0xea, 0x61, 0xc7, 0xe7, 0xf5, 0xeb, 0xdf, 0xa3, 0xd7,
	0xef, 0x63, 0x88, 0x8b, 0x85, 0xb7, 0xc6, 0x56, 0x5a, 0x46, 0xea, 0xdc, 0x47, 0x5e, 0x95, 0xaa,
	0x52, 0xde, 0x5d, 0x44, 0xd2, 0x9a, 0xaf, 0x7a, 0xbe, 0x70, 0x85, 0xd7, 0xd6, 0x34, 0xd1, 0x37,
	0x75, 0xa6, 0x65, 0xa9, 0x84, 0xb4, 0xc6, 0x3b, 0x5b, 0x8a, 0xda, 0x96, 0x5a, 0x5e, 0x84, 0xb5,
	0xb3, 0xde, 0xa2, 0x57, 0x1d, 0x24, 0x54, 0xe7, 0x3e, 0xec, 0x20, 0xe1, 0x65, 0xc8, 0xe8, 0xe7,
	0x2d, 0xf8, 0xe8, 0x68, 0x05, 0xc2, 0x2b, 0x4e, 0xd6, 0x62, 0x90, 0x83, 0xf7, 0x56, 0x40, 0xd1,
	0x48, 0x5b, 0xab, 0x01, 0x18, 0x82, 0xbd, 0x9d, 0x37, 0x34, 0xdc, 0x8c, 0x1b, 0x5e, 0xc5, 0x0c,
	0x2f, 0x29, 0xbe, 0xc4, 0xb2, 0xed, 0xfa, 0xaf, 0x40, 0xdf, 0x01, 0xdc, 0x5d, 0x9b, 0xd6, 0xca,
	0x55, 0xba, 0x69, 0xb4, 0x35, 0x83, 0x5e, 0xeb, 0x9c, 0xdf, 0x9c, 0x73, 0xd6, 0xb1, 0x59, 0x50,
	0xff, 0x53, 0x41, 0xef, 0xe0, 0x33, 0xa7, 0xaa, 0x42, 0x1b, 0x6d, 0xe6, 0xe2, 0x74, 0x4d, 0x15,
	0x8d, 0x92, 0xd6, 0x9c, 0x36, 0x83, 0xfe, 0x10, 0xec, 0xdd, 0x66, 0x83, 0xae, 0x63, 0xba, 0x6e,
	0xe0, 0xab, 0xf3, 0xd1, 0x17, 0x88, 0xfe, 0x1f, 0x12, 0x0d, 0xe1, 0xf3, 0xc3, 0xf4, 0x63, 0x4a,
	0x8f, 0x53, 0x81, 0x69, 0x9a, 0x33, 0x9a, 0x88, 0x8c, 0x26, 0x31, 0x9e, 0x09, 0x8e, 0x69, 0x46,
	0x82, 0x2d, 0xb4, 0x0b, 0xef, 0xbf, 0x1f, 0xe3, 0x9c, 0xb2, 0x99, 0x60, 0x84, 0x93, 0x3c, 0x00,
	0xe8, 0x21, 0x7c, 0xc0, 0x49, 0x42, 0x70, 0x1e, 0x1f, 0x91, 0x75, 0xb1, 0x37, 0x5a, 0xc0, 0x27,
	0xd7, 0x8c, 0x82, 0x5e, 0xc2, 0x17, 0xd7, 0x98, 0x64, 0x84, 0x1d, 0xc4, 0x9c, 0xc7, 0x34, 0x0d,
	0xb6, 0xd0, 0x36, 0xbc, 0x3b, 0x4e, 0x12, 0x7a, 0x4c, 0xa6, 0x01, 0x40, 0x3b, 0x10, 0x66, 0x8c,
	0x7e, 0x88, 0x27, 0x71, 0x4e, 0xa6, 0x41, 0x0f, 0x3d, 0x86, 0x88, 0x91, 0x03, 0x9a, 0x13, 0x31,
	0xc6, 0x98, 0x70, 0x2e, 0x68, 0x9a, 0xcc, 0x82, 0xfe, 0xe4, 0x17, 0x80, 0xaf, 0xa5, 0xad, 0x36,
	0x5c, 0xc1, 0xe4, 0xe9, 0x55, 0x3b, 0xc8, 0x96, 0xb9, 0xcc, 0xc0, 0xa7, 0xcf, 0x73, 0xed, 0xcf,
	0x16, 0x27, 0xa1, 0xb4, 0x55, 0xd4, 0xf1, 0xf6, 0x0b, 0xbd, 0x0c, 0xbb, 0x72, 0xa6, 0x28, 0xf7,
	0xdb, 0x04, 0x37, 0x51, 0xe3, 0x64, 0xb4, 0x7c, 0xe6, 0xa8, 0xd5, 0xd1, 0x66, 0xbf, 0xe2, 0x37,
	0x00, 0x3f, 0x7a, 0xfd, 0xf1, 0x61, 0x7e, 0x72, 0xa7, 0xbd, 0xf6, 0xf6, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x23, 0x4a, 0x13, 0x79, 0x4f, 0x03, 0x00, 0x00,
}
