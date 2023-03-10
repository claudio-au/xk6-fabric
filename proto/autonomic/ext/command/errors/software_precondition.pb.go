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
// source: autonomic/ext/command/errors/software_precondition.proto

package errors

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

type RemoteControlViolation_RemoteControlViolationType int32

const (
	RemoteControlViolation_UNKNOWN_REMOTE_CONTROL_VIOLATION_TYPE RemoteControlViolation_RemoteControlViolationType = 0
	// This device or vehicle is not in a state where remote commands are allowed
	RemoteControlViolation_VEHICLE_REMOTE_CONTROL_DISABLED RemoteControlViolation_RemoteControlViolationType = 1
)

var RemoteControlViolation_RemoteControlViolationType_name = map[int32]string{
	0: "UNKNOWN_REMOTE_CONTROL_VIOLATION_TYPE",
	1: "VEHICLE_REMOTE_CONTROL_DISABLED",
}

var RemoteControlViolation_RemoteControlViolationType_value = map[string]int32{
	"UNKNOWN_REMOTE_CONTROL_VIOLATION_TYPE": 0,
	"VEHICLE_REMOTE_CONTROL_DISABLED":       1,
}

func (x RemoteControlViolation_RemoteControlViolationType) String() string {
	return proto.EnumName(RemoteControlViolation_RemoteControlViolationType_name, int32(x))
}

func (RemoteControlViolation_RemoteControlViolationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_320f11d21832695e, []int{2, 0}
}

type AuthenticationFailure_AuthenticationFailureType int32

const (
	AuthenticationFailure_UNKNOWN_AUTHENTICATION_FAILURE AuthenticationFailure_AuthenticationFailureType = 0
	// Token is invalid or missing
	AuthenticationFailure_INVALID_TOKEN AuthenticationFailure_AuthenticationFailureType = 1
	// Token is expired
	AuthenticationFailure_TOKEN_EXPIRED AuthenticationFailure_AuthenticationFailureType = 2
	// Valid token, but no permission to execute command
	AuthenticationFailure_NOT_AUTHORIZED AuthenticationFailure_AuthenticationFailureType = 3
)

var AuthenticationFailure_AuthenticationFailureType_name = map[int32]string{
	0: "UNKNOWN_AUTHENTICATION_FAILURE",
	1: "INVALID_TOKEN",
	2: "TOKEN_EXPIRED",
	3: "NOT_AUTHORIZED",
}

var AuthenticationFailure_AuthenticationFailureType_value = map[string]int32{
	"UNKNOWN_AUTHENTICATION_FAILURE": 0,
	"INVALID_TOKEN":                  1,
	"TOKEN_EXPIRED":                  2,
	"NOT_AUTHORIZED":                 3,
}

func (x AuthenticationFailure_AuthenticationFailureType) String() string {
	return proto.EnumName(AuthenticationFailure_AuthenticationFailureType_name, int32(x))
}

func (AuthenticationFailure_AuthenticationFailureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_320f11d21832695e, []int{3, 0}
}

type ExecutionSequenceFailure_ExecutionSequenceFailureType int32

const (
	ExecutionSequenceFailure_UNKNOWN_EXECUTION_SEQUENCE_FAILURE ExecutionSequenceFailure_ExecutionSequenceFailureType = 0
	// Command issued was in already progress
	ExecutionSequenceFailure_COMMAND_IN_PROGRESS ExecutionSequenceFailure_ExecutionSequenceFailureType = 1
	// Command was issued out of sequence
	ExecutionSequenceFailure_COMMAND_OUT_OF_SEQUENCE ExecutionSequenceFailure_ExecutionSequenceFailureType = 2
)

var ExecutionSequenceFailure_ExecutionSequenceFailureType_name = map[int32]string{
	0: "UNKNOWN_EXECUTION_SEQUENCE_FAILURE",
	1: "COMMAND_IN_PROGRESS",
	2: "COMMAND_OUT_OF_SEQUENCE",
}

var ExecutionSequenceFailure_ExecutionSequenceFailureType_value = map[string]int32{
	"UNKNOWN_EXECUTION_SEQUENCE_FAILURE": 0,
	"COMMAND_IN_PROGRESS":                1,
	"COMMAND_OUT_OF_SEQUENCE":            2,
}

func (x ExecutionSequenceFailure_ExecutionSequenceFailureType) String() string {
	return proto.EnumName(ExecutionSequenceFailure_ExecutionSequenceFailureType_name, int32(x))
}

func (ExecutionSequenceFailure_ExecutionSequenceFailureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_320f11d21832695e, []int{4, 0}
}

type SoftwareStateViolation struct {
	// Types that are valid to be assigned to Type:
	//	*SoftwareStateViolation_SequentialExecutionViolation
	//	*SoftwareStateViolation_RemoteControlViolation
	//	*SoftwareStateViolation_AuthenticationFailure
	//	*SoftwareStateViolation_ExecutionSequenceFailure
	Type                 isSoftwareStateViolation_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SoftwareStateViolation) Reset()         { *m = SoftwareStateViolation{} }
func (m *SoftwareStateViolation) String() string { return proto.CompactTextString(m) }
func (*SoftwareStateViolation) ProtoMessage()    {}
func (*SoftwareStateViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_320f11d21832695e, []int{0}
}

func (m *SoftwareStateViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoftwareStateViolation.Unmarshal(m, b)
}
func (m *SoftwareStateViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoftwareStateViolation.Marshal(b, m, deterministic)
}
func (m *SoftwareStateViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoftwareStateViolation.Merge(m, src)
}
func (m *SoftwareStateViolation) XXX_Size() int {
	return xxx_messageInfo_SoftwareStateViolation.Size(m)
}
func (m *SoftwareStateViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_SoftwareStateViolation.DiscardUnknown(m)
}

var xxx_messageInfo_SoftwareStateViolation proto.InternalMessageInfo

type isSoftwareStateViolation_Type interface {
	isSoftwareStateViolation_Type()
}

type SoftwareStateViolation_SequentialExecutionViolation struct {
	SequentialExecutionViolation *SequentialExecutionViolation `protobuf:"bytes,1,opt,name=sequential_execution_violation,json=sequentialExecutionViolation,proto3,oneof"`
}

type SoftwareStateViolation_RemoteControlViolation struct {
	RemoteControlViolation *RemoteControlViolation `protobuf:"bytes,2,opt,name=remote_control_violation,json=remoteControlViolation,proto3,oneof"`
}

type SoftwareStateViolation_AuthenticationFailure struct {
	AuthenticationFailure *AuthenticationFailure `protobuf:"bytes,3,opt,name=authentication_failure,json=authenticationFailure,proto3,oneof"`
}

type SoftwareStateViolation_ExecutionSequenceFailure struct {
	ExecutionSequenceFailure *ExecutionSequenceFailure `protobuf:"bytes,4,opt,name=execution_sequence_failure,json=executionSequenceFailure,proto3,oneof"`
}

func (*SoftwareStateViolation_SequentialExecutionViolation) isSoftwareStateViolation_Type() {}

func (*SoftwareStateViolation_RemoteControlViolation) isSoftwareStateViolation_Type() {}

func (*SoftwareStateViolation_AuthenticationFailure) isSoftwareStateViolation_Type() {}

func (*SoftwareStateViolation_ExecutionSequenceFailure) isSoftwareStateViolation_Type() {}

func (m *SoftwareStateViolation) GetType() isSoftwareStateViolation_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SoftwareStateViolation) GetSequentialExecutionViolation() *SequentialExecutionViolation {
	if x, ok := m.GetType().(*SoftwareStateViolation_SequentialExecutionViolation); ok {
		return x.SequentialExecutionViolation
	}
	return nil
}

func (m *SoftwareStateViolation) GetRemoteControlViolation() *RemoteControlViolation {
	if x, ok := m.GetType().(*SoftwareStateViolation_RemoteControlViolation); ok {
		return x.RemoteControlViolation
	}
	return nil
}

func (m *SoftwareStateViolation) GetAuthenticationFailure() *AuthenticationFailure {
	if x, ok := m.GetType().(*SoftwareStateViolation_AuthenticationFailure); ok {
		return x.AuthenticationFailure
	}
	return nil
}

func (m *SoftwareStateViolation) GetExecutionSequenceFailure() *ExecutionSequenceFailure {
	if x, ok := m.GetType().(*SoftwareStateViolation_ExecutionSequenceFailure); ok {
		return x.ExecutionSequenceFailure
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SoftwareStateViolation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SoftwareStateViolation_SequentialExecutionViolation)(nil),
		(*SoftwareStateViolation_RemoteControlViolation)(nil),
		(*SoftwareStateViolation_AuthenticationFailure)(nil),
		(*SoftwareStateViolation_ExecutionSequenceFailure)(nil),
	}
}

// The number of times a command may be executed sequentially within a period of time was exhausted
// (e.g. You may refresh a remote vehicle start N number of times before a manual start is required)
type SequentialExecutionViolation struct {
	// delay until this command may be executed again
	ResetDelay int32 `protobuf:"varint,1,opt,name=reset_delay,json=resetDelay,proto3" json:"reset_delay,omitempty"`
	// the maximum number of times this command is allowed to be executed sequentially before
	// some other action must occur
	MaximumExecutionCount int32    `protobuf:"varint,2,opt,name=maximum_execution_count,json=maximumExecutionCount,proto3" json:"maximum_execution_count,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *SequentialExecutionViolation) Reset()         { *m = SequentialExecutionViolation{} }
func (m *SequentialExecutionViolation) String() string { return proto.CompactTextString(m) }
func (*SequentialExecutionViolation) ProtoMessage()    {}
func (*SequentialExecutionViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_320f11d21832695e, []int{1}
}

func (m *SequentialExecutionViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SequentialExecutionViolation.Unmarshal(m, b)
}
func (m *SequentialExecutionViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SequentialExecutionViolation.Marshal(b, m, deterministic)
}
func (m *SequentialExecutionViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SequentialExecutionViolation.Merge(m, src)
}
func (m *SequentialExecutionViolation) XXX_Size() int {
	return xxx_messageInfo_SequentialExecutionViolation.Size(m)
}
func (m *SequentialExecutionViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_SequentialExecutionViolation.DiscardUnknown(m)
}

var xxx_messageInfo_SequentialExecutionViolation proto.InternalMessageInfo

func (m *SequentialExecutionViolation) GetResetDelay() int32 {
	if m != nil {
		return m.ResetDelay
	}
	return 0
}

func (m *SequentialExecutionViolation) GetMaximumExecutionCount() int32 {
	if m != nil {
		return m.MaximumExecutionCount
	}
	return 0
}

type RemoteControlViolation struct {
	RemoteControlViolation RemoteControlViolation_RemoteControlViolationType `protobuf:"varint,1,opt,name=remote_control_violation,json=remoteControlViolation,proto3,enum=autonomic.ext.command.errors.RemoteControlViolation_RemoteControlViolationType" json:"remote_control_violation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                                          `json:"-"`
	XXX_unrecognized       []byte                                            `json:"-"`
	XXX_sizecache          int32                                             `json:"-"`
}

func (m *RemoteControlViolation) Reset()         { *m = RemoteControlViolation{} }
func (m *RemoteControlViolation) String() string { return proto.CompactTextString(m) }
func (*RemoteControlViolation) ProtoMessage()    {}
func (*RemoteControlViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_320f11d21832695e, []int{2}
}

func (m *RemoteControlViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteControlViolation.Unmarshal(m, b)
}
func (m *RemoteControlViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteControlViolation.Marshal(b, m, deterministic)
}
func (m *RemoteControlViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteControlViolation.Merge(m, src)
}
func (m *RemoteControlViolation) XXX_Size() int {
	return xxx_messageInfo_RemoteControlViolation.Size(m)
}
func (m *RemoteControlViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteControlViolation.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteControlViolation proto.InternalMessageInfo

func (m *RemoteControlViolation) GetRemoteControlViolation() RemoteControlViolation_RemoteControlViolationType {
	if m != nil {
		return m.RemoteControlViolation
	}
	return RemoteControlViolation_UNKNOWN_REMOTE_CONTROL_VIOLATION_TYPE
}

type AuthenticationFailure struct {
	AuthFailureType      AuthenticationFailure_AuthenticationFailureType `protobuf:"varint,1,opt,name=auth_failure_type,json=authFailureType,proto3,enum=autonomic.ext.command.errors.AuthenticationFailure_AuthenticationFailureType" json:"auth_failure_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *AuthenticationFailure) Reset()         { *m = AuthenticationFailure{} }
func (m *AuthenticationFailure) String() string { return proto.CompactTextString(m) }
func (*AuthenticationFailure) ProtoMessage()    {}
func (*AuthenticationFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_320f11d21832695e, []int{3}
}

func (m *AuthenticationFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationFailure.Unmarshal(m, b)
}
func (m *AuthenticationFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationFailure.Marshal(b, m, deterministic)
}
func (m *AuthenticationFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationFailure.Merge(m, src)
}
func (m *AuthenticationFailure) XXX_Size() int {
	return xxx_messageInfo_AuthenticationFailure.Size(m)
}
func (m *AuthenticationFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationFailure.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationFailure proto.InternalMessageInfo

func (m *AuthenticationFailure) GetAuthFailureType() AuthenticationFailure_AuthenticationFailureType {
	if m != nil {
		return m.AuthFailureType
	}
	return AuthenticationFailure_UNKNOWN_AUTHENTICATION_FAILURE
}

type ExecutionSequenceFailure struct {
	ExecutionSequenceFailureType ExecutionSequenceFailure_ExecutionSequenceFailureType `protobuf:"varint,1,opt,name=execution_sequence_failure_type,json=executionSequenceFailureType,proto3,enum=autonomic.ext.command.errors.ExecutionSequenceFailure_ExecutionSequenceFailureType" json:"execution_sequence_failure_type,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                                              `json:"-"`
	XXX_unrecognized             []byte                                                `json:"-"`
	XXX_sizecache                int32                                                 `json:"-"`
}

func (m *ExecutionSequenceFailure) Reset()         { *m = ExecutionSequenceFailure{} }
func (m *ExecutionSequenceFailure) String() string { return proto.CompactTextString(m) }
func (*ExecutionSequenceFailure) ProtoMessage()    {}
func (*ExecutionSequenceFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_320f11d21832695e, []int{4}
}

func (m *ExecutionSequenceFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionSequenceFailure.Unmarshal(m, b)
}
func (m *ExecutionSequenceFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionSequenceFailure.Marshal(b, m, deterministic)
}
func (m *ExecutionSequenceFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionSequenceFailure.Merge(m, src)
}
func (m *ExecutionSequenceFailure) XXX_Size() int {
	return xxx_messageInfo_ExecutionSequenceFailure.Size(m)
}
func (m *ExecutionSequenceFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionSequenceFailure.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionSequenceFailure proto.InternalMessageInfo

func (m *ExecutionSequenceFailure) GetExecutionSequenceFailureType() ExecutionSequenceFailure_ExecutionSequenceFailureType {
	if m != nil {
		return m.ExecutionSequenceFailureType
	}
	return ExecutionSequenceFailure_UNKNOWN_EXECUTION_SEQUENCE_FAILURE
}

func init() {
	proto.RegisterEnum("autonomic.ext.command.errors.RemoteControlViolation_RemoteControlViolationType", RemoteControlViolation_RemoteControlViolationType_name, RemoteControlViolation_RemoteControlViolationType_value)
	proto.RegisterEnum("autonomic.ext.command.errors.AuthenticationFailure_AuthenticationFailureType", AuthenticationFailure_AuthenticationFailureType_name, AuthenticationFailure_AuthenticationFailureType_value)
	proto.RegisterEnum("autonomic.ext.command.errors.ExecutionSequenceFailure_ExecutionSequenceFailureType", ExecutionSequenceFailure_ExecutionSequenceFailureType_name, ExecutionSequenceFailure_ExecutionSequenceFailureType_value)
	proto.RegisterType((*SoftwareStateViolation)(nil), "autonomic.ext.command.errors.SoftwareStateViolation")
	proto.RegisterType((*SequentialExecutionViolation)(nil), "autonomic.ext.command.errors.SequentialExecutionViolation")
	proto.RegisterType((*RemoteControlViolation)(nil), "autonomic.ext.command.errors.RemoteControlViolation")
	proto.RegisterType((*AuthenticationFailure)(nil), "autonomic.ext.command.errors.AuthenticationFailure")
	proto.RegisterType((*ExecutionSequenceFailure)(nil), "autonomic.ext.command.errors.ExecutionSequenceFailure")
}

func init() {
	proto.RegisterFile("autonomic/ext/command/errors/software_precondition.proto", fileDescriptor_320f11d21832695e)
}

var fileDescriptor_320f11d21832695e = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xae, 0x93, 0xb6, 0xc3, 0x55, 0x94, 0xf4, 0x50, 0xd3, 0x50, 0x42, 0x5b, 0x19, 0x81, 0x60,
	0x68, 0x2c, 0xb5, 0xa8, 0x42, 0x6c, 0x8e, 0x73, 0x25, 0x56, 0x13, 0x3b, 0xd8, 0x4e, 0x28, 0x5d,
	0x4e, 0xae, 0x7b, 0xa5, 0x96, 0x6c, 0x5f, 0x70, 0xce, 0x6d, 0x22, 0x31, 0xb1, 0x31, 0xb0, 0x30,
	0xb2, 0xc1, 0x7f, 0xe0, 0xbf, 0xb1, 0x20, 0xa1, 0xbb, 0x24, 0xae, 0x41, 0xb1, 0x87, 0x6e, 0xf6,
	0xfb, 0xde, 0xf7, 0xbe, 0xbb, 0xef, 0xbd, 0xd3, 0x03, 0xaf, 0xdc, 0x84, 0xd1, 0x88, 0x86, 0xbe,
	0xa7, 0x90, 0x31, 0x53, 0x3c, 0x1a, 0x86, 0x6e, 0x74, 0xa1, 0x90, 0x38, 0xa6, 0xf1, 0x48, 0x19,
	0xd1, 0x4b, 0x76, 0xe3, 0xc6, 0x04, 0x0f, 0x63, 0xe2, 0xd1, 0xe8, 0xc2, 0x67, 0x3e, 0x8d, 0x1a,
	0xc3, 0x98, 0x32, 0x0a, 0xeb, 0x29, 0xb3, 0x41, 0xc6, 0xac, 0x31, 0x63, 0x36, 0xa6, 0x4c, 0xf9,
	0x4f, 0x19, 0x54, 0xed, 0x19, 0xdb, 0x66, 0x2e, 0x23, 0x03, 0x9f, 0x06, 0x2e, 0xa7, 0xc3, 0xcf,
	0x12, 0xd8, 0x19, 0x91, 0x8f, 0x09, 0x89, 0x98, 0xef, 0x06, 0x98, 0x8c, 0x89, 0x97, 0x70, 0x00,
	0x5f, 0xcf, 0x53, 0x6a, 0xd2, 0x9e, 0xf4, 0x7c, 0xed, 0xe0, 0x75, 0xa3, 0x48, 0xa2, 0x61, 0xa7,
	0x35, 0xd0, 0xbc, 0x44, 0x2a, 0xd2, 0x5e, 0xb2, 0xea, 0xa3, 0x02, 0x1c, 0x0e, 0x41, 0x2d, 0x26,
	0x21, 0x65, 0x04, 0x7b, 0x34, 0x62, 0x31, 0x0d, 0x32, 0xea, 0x25, 0xa1, 0xfe, 0xb2, 0x58, 0xdd,
	0x12, 0x6c, 0x6d, 0x4a, 0xce, 0xea, 0x56, 0xe3, 0x85, 0x08, 0x0c, 0x40, 0xd5, 0x4d, 0xd8, 0x15,
	0x3f, 0x91, 0x27, 0x22, 0xf8, 0xd2, 0xf5, 0x83, 0x24, 0x26, 0xb5, 0xb2, 0xd0, 0x3b, 0x2c, 0xd6,
	0x53, 0xff, 0xe1, 0x1e, 0x4f, 0xa9, 0xed, 0x25, 0x6b, 0xd3, 0x5d, 0x04, 0xc0, 0x6b, 0xb0, 0x7d,
	0x6b, 0xec, 0xd4, 0x09, 0x8f, 0xa4, 0x8a, 0xcb, 0x42, 0xf1, 0xa8, 0x58, 0x31, 0x75, 0xcd, 0x9e,
	0xd1, 0x6f, 0x45, 0x6b, 0x24, 0x07, 0x6b, 0xae, 0x82, 0x65, 0x36, 0x19, 0x12, 0xf9, 0x06, 0xd4,
	0x8b, 0xfa, 0x03, 0x77, 0xc1, 0x5a, 0x4c, 0x46, 0x84, 0xe1, 0x0b, 0x12, 0xb8, 0x13, 0xd1, 0xf0,
	0x15, 0x0b, 0x88, 0x50, 0x8b, 0x47, 0xe0, 0x11, 0xd8, 0x0a, 0xdd, 0xb1, 0x1f, 0x26, 0x61, 0x66,
	0x42, 0x3c, 0x9a, 0x44, 0x4c, 0xf4, 0x67, 0xc5, 0xda, 0x9c, 0xc1, 0x69, 0x71, 0x8d, 0x83, 0xf2,
	0xb7, 0x12, 0xa8, 0x2e, 0xee, 0x0d, 0xfc, 0x22, 0x15, 0x34, 0x9d, 0x9f, 0x60, 0xfd, 0xc0, 0xbc,
	0x4b, 0xd3, 0x73, 0xc2, 0xce, 0x64, 0x48, 0xf2, 0xa6, 0x41, 0x0e, 0xc0, 0x76, 0x3e, 0x0b, 0xbe,
	0x00, 0x4f, 0xfb, 0xc6, 0x89, 0x61, 0xbe, 0x33, 0xb0, 0x85, 0xba, 0xa6, 0x83, 0xb0, 0x66, 0x1a,
	0x8e, 0x65, 0x76, 0xf0, 0x40, 0x37, 0x3b, 0xaa, 0xa3, 0x9b, 0x06, 0x76, 0xde, 0xf7, 0x50, 0x65,
	0x09, 0x3e, 0x01, 0xbb, 0x03, 0xd4, 0xd6, 0xb5, 0x0e, 0xfa, 0x3f, 0xb5, 0xa5, 0xdb, 0x6a, 0xb3,
	0x83, 0x5a, 0x15, 0x49, 0xfe, 0x5a, 0x02, 0x9b, 0x0b, 0x07, 0x08, 0x4e, 0xc0, 0x06, 0x1f, 0xa0,
	0xf9, 0x64, 0x60, 0xde, 0xbc, 0x99, 0x17, 0xdd, 0x3b, 0x0c, 0xe4, 0xe2, 0xa8, 0x70, 0xe2, 0x3e,
	0xd7, 0xc9, 0x04, 0xe4, 0x09, 0x78, 0x98, 0x9b, 0x0d, 0x65, 0xb0, 0x33, 0x77, 0x40, 0xed, 0x3b,
	0x6d, 0x64, 0x38, 0xba, 0x36, 0xbd, 0xf7, 0xb1, 0xaa, 0x77, 0xfa, 0x16, 0xbf, 0xfa, 0x06, 0xb8,
	0xa7, 0x1b, 0x03, 0xb5, 0xa3, 0xb7, 0xb0, 0x63, 0x9e, 0x20, 0xa3, 0x22, 0xf1, 0x90, 0xf8, 0xc4,
	0xe8, 0xb4, 0xa7, 0x5b, 0xa8, 0x55, 0x29, 0x41, 0x08, 0xd6, 0x0d, 0xd3, 0x11, 0x55, 0x4c, 0x4b,
	0x3f, 0x43, 0xad, 0x4a, 0x59, 0xfe, 0x55, 0x02, 0xb5, 0xbc, 0xf1, 0x86, 0xdf, 0x25, 0xb0, 0x9b,
	0xff, 0x76, 0xb2, 0x0e, 0xd9, 0x77, 0x7b, 0x40, 0xb9, 0x80, 0xf0, 0xa9, 0x4e, 0x0a, 0x50, 0xf9,
	0x13, 0xa8, 0x17, 0xb1, 0xe1, 0x33, 0x20, 0xcf, 0x7d, 0x43, 0xa7, 0x48, 0xeb, 0x0b, 0xcb, 0x6c,
	0xf4, 0xb6, 0x8f, 0x0c, 0x0d, 0x65, 0xbc, 0xdb, 0x02, 0x0f, 0x34, 0xb3, 0xdb, 0x55, 0x8d, 0x16,
	0xd6, 0x0d, 0xdc, 0xb3, 0xcc, 0x37, 0x16, 0xb2, 0xed, 0x8a, 0x04, 0x1f, 0x81, 0xad, 0x39, 0x60,
	0xf6, 0x1d, 0x6c, 0x1e, 0xa7, 0xec, 0x4a, 0xa9, 0xf9, 0x43, 0x02, 0x7b, 0x1e, 0x0d, 0x0b, 0xaf,
	0xdd, 0x7c, 0xac, 0x4d, 0xff, 0x7b, 0x99, 0x9d, 0x81, 0x38, 0xd4, 0xe3, 0x7b, 0xa3, 0x27, 0x9d,
	0xd9, 0x1f, 0x7c, 0x76, 0x95, 0x9c, 0x73, 0x9e, 0x92, 0x56, 0xda, 0x77, 0x7d, 0xbe, 0x81, 0x48,
	0x1c, 0xb9, 0xc1, 0xbe, 0xd8, 0x30, 0x23, 0x65, 0x14, 0x7b, 0x4a, 0xe8, 0xfa, 0x91, 0x22, 0xfe,
	0x95, 0xa2, 0x55, 0xf5, 0x5b, 0x92, 0x7e, 0x96, 0xca, 0x6a, 0xdf, 0x39, 0x5f, 0x15, 0xc9, 0x87,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x80, 0x30, 0x5a, 0xf9, 0xda, 0x06, 0x00, 0x00,
}
