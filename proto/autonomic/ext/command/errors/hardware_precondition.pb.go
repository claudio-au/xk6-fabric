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
// source: autonomic/ext/command/errors/hardware_precondition.proto

package errors

import (
	fmt "fmt"
	enumerations "xk6-fabric/proto/autonomic/ext/telemetry/enumerations"
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

type MotorControlButtonViolation_MotorControlButtonState int32

const (
	MotorControlButtonViolation_UNKNOWN_MOTOR_CONTROL_BUTTON_STATE MotorControlButtonViolation_MotorControlButtonState = 0
	// Motor control button was depressed when the command was received
	MotorControlButtonViolation_MOTOR_CONTROL_BUTTON_DEPRESSED MotorControlButtonViolation_MotorControlButtonState = 1
	// Motor control button was released when command was received
	MotorControlButtonViolation_MOTOR_CONTROL_BUTTON_RELEASED MotorControlButtonViolation_MotorControlButtonState = 2
)

var MotorControlButtonViolation_MotorControlButtonState_name = map[int32]string{
	0: "UNKNOWN_MOTOR_CONTROL_BUTTON_STATE",
	1: "MOTOR_CONTROL_BUTTON_DEPRESSED",
	2: "MOTOR_CONTROL_BUTTON_RELEASED",
}

var MotorControlButtonViolation_MotorControlButtonState_value = map[string]int32{
	"UNKNOWN_MOTOR_CONTROL_BUTTON_STATE": 0,
	"MOTOR_CONTROL_BUTTON_DEPRESSED":     1,
	"MOTOR_CONTROL_BUTTON_RELEASED":      2,
}

func (x MotorControlButtonViolation_MotorControlButtonState) String() string {
	return proto.EnumName(MotorControlButtonViolation_MotorControlButtonState_name, int32(x))
}

func (MotorControlButtonViolation_MotorControlButtonState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{2, 0}
}

type VehicleMotionViolation_VehicleMotionState int32

const (
	VehicleMotionViolation_UNKNOWN_VEHICLE_MOTION_STATE VehicleMotionViolation_VehicleMotionState = 0
	// Vehicle was in motion when command was received
	VehicleMotionViolation_VEHICLE_IN_MOTION VehicleMotionViolation_VehicleMotionState = 1
	// Vehicle was stationary when command was received
	VehicleMotionViolation_VEHICLE_STATIONARY VehicleMotionViolation_VehicleMotionState = 2
)

var VehicleMotionViolation_VehicleMotionState_name = map[int32]string{
	0: "UNKNOWN_VEHICLE_MOTION_STATE",
	1: "VEHICLE_IN_MOTION",
	2: "VEHICLE_STATIONARY",
}

var VehicleMotionViolation_VehicleMotionState_value = map[string]int32{
	"UNKNOWN_VEHICLE_MOTION_STATE": 0,
	"VEHICLE_IN_MOTION":            1,
	"VEHICLE_STATIONARY":           2,
}

func (x VehicleMotionViolation_VehicleMotionState) String() string {
	return proto.EnumName(VehicleMotionViolation_VehicleMotionState_name, int32(x))
}

func (VehicleMotionViolation_VehicleMotionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{3, 0}
}

type KeyFobViolation_KeyFobProximityState int32

const (
	KeyFobViolation_UNKNOWN_KEYFOB_PROXIMITY_STATE KeyFobViolation_KeyFobProximityState = 0
	// Key fob was in the vehicle when the command was received
	KeyFobViolation_KEY_FOB_IN_VEHICLE KeyFobViolation_KeyFobProximityState = 1
	// Key fob was in the vehicle proximity when the command was received
	KeyFobViolation_KEY_FOB_IN_VEHICLE_PROXIMITY KeyFobViolation_KeyFobProximityState = 2
	// Key fob was not in the vehicle proximity when the command was received
	KeyFobViolation_KEY_FOB_NOT_IN_VEHICLE_PROXIMITY KeyFobViolation_KeyFobProximityState = 3
)

var KeyFobViolation_KeyFobProximityState_name = map[int32]string{
	0: "UNKNOWN_KEYFOB_PROXIMITY_STATE",
	1: "KEY_FOB_IN_VEHICLE",
	2: "KEY_FOB_IN_VEHICLE_PROXIMITY",
	3: "KEY_FOB_NOT_IN_VEHICLE_PROXIMITY",
}

var KeyFobViolation_KeyFobProximityState_value = map[string]int32{
	"UNKNOWN_KEYFOB_PROXIMITY_STATE":   0,
	"KEY_FOB_IN_VEHICLE":               1,
	"KEY_FOB_IN_VEHICLE_PROXIMITY":     2,
	"KEY_FOB_NOT_IN_VEHICLE_PROXIMITY": 3,
}

func (x KeyFobViolation_KeyFobProximityState) String() string {
	return proto.EnumName(KeyFobViolation_KeyFobProximityState_name, int32(x))
}

func (KeyFobViolation_KeyFobProximityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{4, 0}
}

// Note: this file is imported directly by the parent file command_error_data.proto.
type HardwareStateViolation struct {
	// Types that are valid to be assigned to Type:
	//	*HardwareStateViolation_AlarmStateViolation
	//	*HardwareStateViolation_DoorLockStateViolation
	//	*HardwareStateViolation_DoorStateViolation
	//	*HardwareStateViolation_DoorHandleStateViolation
	//	*HardwareStateViolation_EngineStateViolation
	//	*HardwareStateViolation_MotorControlButtonViolation
	//	*HardwareStateViolation_VehicleMotionViolation
	//	*HardwareStateViolation_KeyFobViolation
	//	*HardwareStateViolation_HoodStateViolation
	//	*HardwareStateViolation_GearPositionViolation
	Type                 isHardwareStateViolation_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *HardwareStateViolation) Reset()         { *m = HardwareStateViolation{} }
func (m *HardwareStateViolation) String() string { return proto.CompactTextString(m) }
func (*HardwareStateViolation) ProtoMessage()    {}
func (*HardwareStateViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{0}
}

func (m *HardwareStateViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HardwareStateViolation.Unmarshal(m, b)
}
func (m *HardwareStateViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HardwareStateViolation.Marshal(b, m, deterministic)
}
func (m *HardwareStateViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HardwareStateViolation.Merge(m, src)
}
func (m *HardwareStateViolation) XXX_Size() int {
	return xxx_messageInfo_HardwareStateViolation.Size(m)
}
func (m *HardwareStateViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_HardwareStateViolation.DiscardUnknown(m)
}

var xxx_messageInfo_HardwareStateViolation proto.InternalMessageInfo

type isHardwareStateViolation_Type interface {
	isHardwareStateViolation_Type()
}

type HardwareStateViolation_AlarmStateViolation struct {
	AlarmStateViolation *AlarmStateViolation `protobuf:"bytes,1,opt,name=alarm_state_violation,json=alarmStateViolation,proto3,oneof"`
}

type HardwareStateViolation_DoorLockStateViolation struct {
	DoorLockStateViolation *DoorLockStateViolation `protobuf:"bytes,2,opt,name=door_lock_state_violation,json=doorLockStateViolation,proto3,oneof"`
}

type HardwareStateViolation_DoorStateViolation struct {
	DoorStateViolation *DoorStateViolation `protobuf:"bytes,3,opt,name=door_state_violation,json=doorStateViolation,proto3,oneof"`
}

type HardwareStateViolation_DoorHandleStateViolation struct {
	DoorHandleStateViolation *DoorHandleStateViolation `protobuf:"bytes,4,opt,name=door_handle_state_violation,json=doorHandleStateViolation,proto3,oneof"`
}

type HardwareStateViolation_EngineStateViolation struct {
	EngineStateViolation *EngineStateViolation `protobuf:"bytes,5,opt,name=engine_state_violation,json=engineStateViolation,proto3,oneof"`
}

type HardwareStateViolation_MotorControlButtonViolation struct {
	MotorControlButtonViolation *MotorControlButtonViolation `protobuf:"bytes,6,opt,name=motor_control_button_violation,json=motorControlButtonViolation,proto3,oneof"`
}

type HardwareStateViolation_VehicleMotionViolation struct {
	VehicleMotionViolation *VehicleMotionViolation `protobuf:"bytes,7,opt,name=vehicle_motion_violation,json=vehicleMotionViolation,proto3,oneof"`
}

type HardwareStateViolation_KeyFobViolation struct {
	KeyFobViolation *KeyFobViolation `protobuf:"bytes,8,opt,name=key_fob_violation,json=keyFobViolation,proto3,oneof"`
}

type HardwareStateViolation_HoodStateViolation struct {
	HoodStateViolation *HoodStateViolation `protobuf:"bytes,9,opt,name=hood_state_violation,json=hoodStateViolation,proto3,oneof"`
}

type HardwareStateViolation_GearPositionViolation struct {
	GearPositionViolation *GearPositionViolation `protobuf:"bytes,10,opt,name=gear_position_violation,json=gearPositionViolation,proto3,oneof"`
}

func (*HardwareStateViolation_AlarmStateViolation) isHardwareStateViolation_Type() {}

func (*HardwareStateViolation_DoorLockStateViolation) isHardwareStateViolation_Type() {}

func (*HardwareStateViolation_DoorStateViolation) isHardwareStateViolation_Type() {}

func (*HardwareStateViolation_DoorHandleStateViolation) isHardwareStateViolation_Type() {}

func (*HardwareStateViolation_EngineStateViolation) isHardwareStateViolation_Type() {}

func (*HardwareStateViolation_MotorControlButtonViolation) isHardwareStateViolation_Type() {}

func (*HardwareStateViolation_VehicleMotionViolation) isHardwareStateViolation_Type() {}

func (*HardwareStateViolation_KeyFobViolation) isHardwareStateViolation_Type() {}

func (*HardwareStateViolation_HoodStateViolation) isHardwareStateViolation_Type() {}

func (*HardwareStateViolation_GearPositionViolation) isHardwareStateViolation_Type() {}

func (m *HardwareStateViolation) GetType() isHardwareStateViolation_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *HardwareStateViolation) GetAlarmStateViolation() *AlarmStateViolation {
	if x, ok := m.GetType().(*HardwareStateViolation_AlarmStateViolation); ok {
		return x.AlarmStateViolation
	}
	return nil
}

func (m *HardwareStateViolation) GetDoorLockStateViolation() *DoorLockStateViolation {
	if x, ok := m.GetType().(*HardwareStateViolation_DoorLockStateViolation); ok {
		return x.DoorLockStateViolation
	}
	return nil
}

func (m *HardwareStateViolation) GetDoorStateViolation() *DoorStateViolation {
	if x, ok := m.GetType().(*HardwareStateViolation_DoorStateViolation); ok {
		return x.DoorStateViolation
	}
	return nil
}

func (m *HardwareStateViolation) GetDoorHandleStateViolation() *DoorHandleStateViolation {
	if x, ok := m.GetType().(*HardwareStateViolation_DoorHandleStateViolation); ok {
		return x.DoorHandleStateViolation
	}
	return nil
}

func (m *HardwareStateViolation) GetEngineStateViolation() *EngineStateViolation {
	if x, ok := m.GetType().(*HardwareStateViolation_EngineStateViolation); ok {
		return x.EngineStateViolation
	}
	return nil
}

func (m *HardwareStateViolation) GetMotorControlButtonViolation() *MotorControlButtonViolation {
	if x, ok := m.GetType().(*HardwareStateViolation_MotorControlButtonViolation); ok {
		return x.MotorControlButtonViolation
	}
	return nil
}

func (m *HardwareStateViolation) GetVehicleMotionViolation() *VehicleMotionViolation {
	if x, ok := m.GetType().(*HardwareStateViolation_VehicleMotionViolation); ok {
		return x.VehicleMotionViolation
	}
	return nil
}

func (m *HardwareStateViolation) GetKeyFobViolation() *KeyFobViolation {
	if x, ok := m.GetType().(*HardwareStateViolation_KeyFobViolation); ok {
		return x.KeyFobViolation
	}
	return nil
}

func (m *HardwareStateViolation) GetHoodStateViolation() *HoodStateViolation {
	if x, ok := m.GetType().(*HardwareStateViolation_HoodStateViolation); ok {
		return x.HoodStateViolation
	}
	return nil
}

func (m *HardwareStateViolation) GetGearPositionViolation() *GearPositionViolation {
	if x, ok := m.GetType().(*HardwareStateViolation_GearPositionViolation); ok {
		return x.GearPositionViolation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HardwareStateViolation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HardwareStateViolation_AlarmStateViolation)(nil),
		(*HardwareStateViolation_DoorLockStateViolation)(nil),
		(*HardwareStateViolation_DoorStateViolation)(nil),
		(*HardwareStateViolation_DoorHandleStateViolation)(nil),
		(*HardwareStateViolation_EngineStateViolation)(nil),
		(*HardwareStateViolation_MotorControlButtonViolation)(nil),
		(*HardwareStateViolation_VehicleMotionViolation)(nil),
		(*HardwareStateViolation_KeyFobViolation)(nil),
		(*HardwareStateViolation_HoodStateViolation)(nil),
		(*HardwareStateViolation_GearPositionViolation)(nil),
	}
}

// Contextual information about the door handle which caused a precondition failure for the command.
type DoorHandleStateViolation struct {
	DoorHandleState      DoorHandleState `protobuf:"varint,1,opt,name=door_handle_state,json=doorHandleState,proto3,enum=autonomic.ext.command.errors.DoorHandleState" json:"door_handle_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DoorHandleStateViolation) Reset()         { *m = DoorHandleStateViolation{} }
func (m *DoorHandleStateViolation) String() string { return proto.CompactTextString(m) }
func (*DoorHandleStateViolation) ProtoMessage()    {}
func (*DoorHandleStateViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{1}
}

func (m *DoorHandleStateViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoorHandleStateViolation.Unmarshal(m, b)
}
func (m *DoorHandleStateViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoorHandleStateViolation.Marshal(b, m, deterministic)
}
func (m *DoorHandleStateViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoorHandleStateViolation.Merge(m, src)
}
func (m *DoorHandleStateViolation) XXX_Size() int {
	return xxx_messageInfo_DoorHandleStateViolation.Size(m)
}
func (m *DoorHandleStateViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_DoorHandleStateViolation.DiscardUnknown(m)
}

var xxx_messageInfo_DoorHandleStateViolation proto.InternalMessageInfo

func (m *DoorHandleStateViolation) GetDoorHandleState() DoorHandleState {
	if m != nil {
		return m.DoorHandleState
	}
	return DoorHandleState_UNKNOWN_DOOR_HANDLE_STATE
}

// A precondition violation which is the result of the state of the motor control button
// (i.e. vehicle start / stop button), at the time the command was received
type MotorControlButtonViolation struct {
	MotorControlButtonState MotorControlButtonViolation_MotorControlButtonState `protobuf:"varint,1,opt,name=motor_control_button_state,json=motorControlButtonState,proto3,enum=autonomic.ext.command.errors.MotorControlButtonViolation_MotorControlButtonState" json:"motor_control_button_state,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                                            `json:"-"`
	XXX_unrecognized        []byte                                              `json:"-"`
	XXX_sizecache           int32                                               `json:"-"`
}

func (m *MotorControlButtonViolation) Reset()         { *m = MotorControlButtonViolation{} }
func (m *MotorControlButtonViolation) String() string { return proto.CompactTextString(m) }
func (*MotorControlButtonViolation) ProtoMessage()    {}
func (*MotorControlButtonViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{2}
}

func (m *MotorControlButtonViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MotorControlButtonViolation.Unmarshal(m, b)
}
func (m *MotorControlButtonViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MotorControlButtonViolation.Marshal(b, m, deterministic)
}
func (m *MotorControlButtonViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MotorControlButtonViolation.Merge(m, src)
}
func (m *MotorControlButtonViolation) XXX_Size() int {
	return xxx_messageInfo_MotorControlButtonViolation.Size(m)
}
func (m *MotorControlButtonViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_MotorControlButtonViolation.DiscardUnknown(m)
}

var xxx_messageInfo_MotorControlButtonViolation proto.InternalMessageInfo

func (m *MotorControlButtonViolation) GetMotorControlButtonState() MotorControlButtonViolation_MotorControlButtonState {
	if m != nil {
		return m.MotorControlButtonState
	}
	return MotorControlButtonViolation_UNKNOWN_MOTOR_CONTROL_BUTTON_STATE
}

// A precondition violation which is the result of the movement of the vehicle, at the time the
// command was received.
type VehicleMotionViolation struct {
	VehicleMotionState   VehicleMotionViolation_VehicleMotionState `protobuf:"varint,1,opt,name=vehicle_motion_state,json=vehicleMotionState,proto3,enum=autonomic.ext.command.errors.VehicleMotionViolation_VehicleMotionState" json:"vehicle_motion_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *VehicleMotionViolation) Reset()         { *m = VehicleMotionViolation{} }
func (m *VehicleMotionViolation) String() string { return proto.CompactTextString(m) }
func (*VehicleMotionViolation) ProtoMessage()    {}
func (*VehicleMotionViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{3}
}

func (m *VehicleMotionViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehicleMotionViolation.Unmarshal(m, b)
}
func (m *VehicleMotionViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehicleMotionViolation.Marshal(b, m, deterministic)
}
func (m *VehicleMotionViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehicleMotionViolation.Merge(m, src)
}
func (m *VehicleMotionViolation) XXX_Size() int {
	return xxx_messageInfo_VehicleMotionViolation.Size(m)
}
func (m *VehicleMotionViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_VehicleMotionViolation.DiscardUnknown(m)
}

var xxx_messageInfo_VehicleMotionViolation proto.InternalMessageInfo

func (m *VehicleMotionViolation) GetVehicleMotionState() VehicleMotionViolation_VehicleMotionState {
	if m != nil {
		return m.VehicleMotionState
	}
	return VehicleMotionViolation_UNKNOWN_VEHICLE_MOTION_STATE
}

// A precondition violation which is the result of the location of the key fob, relative to the
// vehicle, at the time the command was received.
type KeyFobViolation struct {
	ProximityState       KeyFobViolation_KeyFobProximityState `protobuf:"varint,1,opt,name=proximity_state,json=proximityState,proto3,enum=autonomic.ext.command.errors.KeyFobViolation_KeyFobProximityState" json:"proximity_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *KeyFobViolation) Reset()         { *m = KeyFobViolation{} }
func (m *KeyFobViolation) String() string { return proto.CompactTextString(m) }
func (*KeyFobViolation) ProtoMessage()    {}
func (*KeyFobViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{4}
}

func (m *KeyFobViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyFobViolation.Unmarshal(m, b)
}
func (m *KeyFobViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyFobViolation.Marshal(b, m, deterministic)
}
func (m *KeyFobViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyFobViolation.Merge(m, src)
}
func (m *KeyFobViolation) XXX_Size() int {
	return xxx_messageInfo_KeyFobViolation.Size(m)
}
func (m *KeyFobViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyFobViolation.DiscardUnknown(m)
}

var xxx_messageInfo_KeyFobViolation proto.InternalMessageInfo

func (m *KeyFobViolation) GetProximityState() KeyFobViolation_KeyFobProximityState {
	if m != nil {
		return m.ProximityState
	}
	return KeyFobViolation_UNKNOWN_KEYFOB_PROXIMITY_STATE
}

// A precondition violation which is the result of the hood state at the time the command
// was received
type HoodStateViolation struct {
	HoodPositionState    HoodPositionState `protobuf:"varint,1,opt,name=hood_position_state,json=hoodPositionState,proto3,enum=autonomic.ext.command.errors.HoodPositionState" json:"hood_position_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HoodStateViolation) Reset()         { *m = HoodStateViolation{} }
func (m *HoodStateViolation) String() string { return proto.CompactTextString(m) }
func (*HoodStateViolation) ProtoMessage()    {}
func (*HoodStateViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{5}
}

func (m *HoodStateViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HoodStateViolation.Unmarshal(m, b)
}
func (m *HoodStateViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HoodStateViolation.Marshal(b, m, deterministic)
}
func (m *HoodStateViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HoodStateViolation.Merge(m, src)
}
func (m *HoodStateViolation) XXX_Size() int {
	return xxx_messageInfo_HoodStateViolation.Size(m)
}
func (m *HoodStateViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_HoodStateViolation.DiscardUnknown(m)
}

var xxx_messageInfo_HoodStateViolation proto.InternalMessageInfo

func (m *HoodStateViolation) GetHoodPositionState() HoodPositionState {
	if m != nil {
		return m.HoodPositionState
	}
	return HoodPositionState_UNKNOWN_HOOD_POSITION_STATE
}

type EngineStateViolation struct {
	// The engine state at the time the command was received,
	// which represents a violating precondition to the execution of the received command
	EngineStatus         enumerations.EngineStatus `protobuf:"varint,1,opt,name=engine_status,json=engineStatus,proto3,enum=autonomic.ext.telemetry.enumerations.engineStatus.EngineStatus" json:"engine_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EngineStateViolation) Reset()         { *m = EngineStateViolation{} }
func (m *EngineStateViolation) String() string { return proto.CompactTextString(m) }
func (*EngineStateViolation) ProtoMessage()    {}
func (*EngineStateViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{6}
}

func (m *EngineStateViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EngineStateViolation.Unmarshal(m, b)
}
func (m *EngineStateViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EngineStateViolation.Marshal(b, m, deterministic)
}
func (m *EngineStateViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EngineStateViolation.Merge(m, src)
}
func (m *EngineStateViolation) XXX_Size() int {
	return xxx_messageInfo_EngineStateViolation.Size(m)
}
func (m *EngineStateViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_EngineStateViolation.DiscardUnknown(m)
}

var xxx_messageInfo_EngineStateViolation proto.InternalMessageInfo

func (m *EngineStateViolation) GetEngineStatus() enumerations.EngineStatus {
	if m != nil {
		return m.EngineStatus
	}
	return enumerations.EngineStatus_UNKNOWN
}

type AlarmStateViolation struct {
	// The current alarm state, which represents a violating precondition to this command's execution
	AlarmStatus          enumerations.AlarmStatus `protobuf:"varint,1,opt,name=alarm_status,json=alarmStatus,proto3,enum=autonomic.ext.telemetry.enumerations.alarmStatus.AlarmStatus" json:"alarm_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AlarmStateViolation) Reset()         { *m = AlarmStateViolation{} }
func (m *AlarmStateViolation) String() string { return proto.CompactTextString(m) }
func (*AlarmStateViolation) ProtoMessage()    {}
func (*AlarmStateViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{7}
}

func (m *AlarmStateViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmStateViolation.Unmarshal(m, b)
}
func (m *AlarmStateViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmStateViolation.Marshal(b, m, deterministic)
}
func (m *AlarmStateViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmStateViolation.Merge(m, src)
}
func (m *AlarmStateViolation) XXX_Size() int {
	return xxx_messageInfo_AlarmStateViolation.Size(m)
}
func (m *AlarmStateViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmStateViolation.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmStateViolation proto.InternalMessageInfo

func (m *AlarmStateViolation) GetAlarmStatus() enumerations.AlarmStatus {
	if m != nil {
		return m.AlarmStatus
	}
	return enumerations.AlarmStatus_UNKNOWN
}

type DoorStateViolation struct {
	// The current door state, which represents a violating precondition to this command's execution
	DoorStatus           enumerations.DoorStatus `protobuf:"varint,1,opt,name=door_status,json=doorStatus,proto3,enum=autonomic.ext.telemetry.enumerations.doorStatus.DoorStatus" json:"door_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DoorStateViolation) Reset()         { *m = DoorStateViolation{} }
func (m *DoorStateViolation) String() string { return proto.CompactTextString(m) }
func (*DoorStateViolation) ProtoMessage()    {}
func (*DoorStateViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{8}
}

func (m *DoorStateViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoorStateViolation.Unmarshal(m, b)
}
func (m *DoorStateViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoorStateViolation.Marshal(b, m, deterministic)
}
func (m *DoorStateViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoorStateViolation.Merge(m, src)
}
func (m *DoorStateViolation) XXX_Size() int {
	return xxx_messageInfo_DoorStateViolation.Size(m)
}
func (m *DoorStateViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_DoorStateViolation.DiscardUnknown(m)
}

var xxx_messageInfo_DoorStateViolation proto.InternalMessageInfo

func (m *DoorStateViolation) GetDoorStatus() enumerations.DoorStatus {
	if m != nil {
		return m.DoorStatus
	}
	return enumerations.DoorStatus_UNKNOWN
}

type DoorLockStateViolation struct {
	// The current door lock state, which represents a violating precondition to this command's execution
	DoorLockStatus       enumerations.DoorLockStatus `protobuf:"varint,1,opt,name=door_lock_status,json=doorLockStatus,proto3,enum=autonomic.ext.telemetry.enumerations.doorLockStatus.DoorLockStatus" json:"door_lock_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *DoorLockStateViolation) Reset()         { *m = DoorLockStateViolation{} }
func (m *DoorLockStateViolation) String() string { return proto.CompactTextString(m) }
func (*DoorLockStateViolation) ProtoMessage()    {}
func (*DoorLockStateViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{9}
}

func (m *DoorLockStateViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoorLockStateViolation.Unmarshal(m, b)
}
func (m *DoorLockStateViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoorLockStateViolation.Marshal(b, m, deterministic)
}
func (m *DoorLockStateViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoorLockStateViolation.Merge(m, src)
}
func (m *DoorLockStateViolation) XXX_Size() int {
	return xxx_messageInfo_DoorLockStateViolation.Size(m)
}
func (m *DoorLockStateViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_DoorLockStateViolation.DiscardUnknown(m)
}

var xxx_messageInfo_DoorLockStateViolation proto.InternalMessageInfo

func (m *DoorLockStateViolation) GetDoorLockStatus() enumerations.DoorLockStatus {
	if m != nil {
		return m.DoorLockStatus
	}
	return enumerations.DoorLockStatus_UNKNOWN
}

type GearPositionViolation struct {
	// The gear position at the time the command was received,
	// which represents a violating precondition to the execution of the received command
	GearPosition         enumerations.GearPosition `protobuf:"varint,1,opt,name=gear_position,json=gearPosition,proto3,enum=autonomic.ext.telemetry.enumerations.gearPosition.GearPosition" json:"gear_position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GearPositionViolation) Reset()         { *m = GearPositionViolation{} }
func (m *GearPositionViolation) String() string { return proto.CompactTextString(m) }
func (*GearPositionViolation) ProtoMessage()    {}
func (*GearPositionViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f51d9503d4e405bc, []int{10}
}

func (m *GearPositionViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GearPositionViolation.Unmarshal(m, b)
}
func (m *GearPositionViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GearPositionViolation.Marshal(b, m, deterministic)
}
func (m *GearPositionViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GearPositionViolation.Merge(m, src)
}
func (m *GearPositionViolation) XXX_Size() int {
	return xxx_messageInfo_GearPositionViolation.Size(m)
}
func (m *GearPositionViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_GearPositionViolation.DiscardUnknown(m)
}

var xxx_messageInfo_GearPositionViolation proto.InternalMessageInfo

func (m *GearPositionViolation) GetGearPosition() enumerations.GearPosition {
	if m != nil {
		return m.GearPosition
	}
	return enumerations.GearPosition_UNKNOWN
}

func init() {
	proto.RegisterEnum("autonomic.ext.command.errors.MotorControlButtonViolation_MotorControlButtonState", MotorControlButtonViolation_MotorControlButtonState_name, MotorControlButtonViolation_MotorControlButtonState_value)
	proto.RegisterEnum("autonomic.ext.command.errors.VehicleMotionViolation_VehicleMotionState", VehicleMotionViolation_VehicleMotionState_name, VehicleMotionViolation_VehicleMotionState_value)
	proto.RegisterEnum("autonomic.ext.command.errors.KeyFobViolation_KeyFobProximityState", KeyFobViolation_KeyFobProximityState_name, KeyFobViolation_KeyFobProximityState_value)
	proto.RegisterType((*HardwareStateViolation)(nil), "autonomic.ext.command.errors.HardwareStateViolation")
	proto.RegisterType((*DoorHandleStateViolation)(nil), "autonomic.ext.command.errors.DoorHandleStateViolation")
	proto.RegisterType((*MotorControlButtonViolation)(nil), "autonomic.ext.command.errors.MotorControlButtonViolation")
	proto.RegisterType((*VehicleMotionViolation)(nil), "autonomic.ext.command.errors.VehicleMotionViolation")
	proto.RegisterType((*KeyFobViolation)(nil), "autonomic.ext.command.errors.KeyFobViolation")
	proto.RegisterType((*HoodStateViolation)(nil), "autonomic.ext.command.errors.HoodStateViolation")
	proto.RegisterType((*EngineStateViolation)(nil), "autonomic.ext.command.errors.EngineStateViolation")
	proto.RegisterType((*AlarmStateViolation)(nil), "autonomic.ext.command.errors.AlarmStateViolation")
	proto.RegisterType((*DoorStateViolation)(nil), "autonomic.ext.command.errors.DoorStateViolation")
	proto.RegisterType((*DoorLockStateViolation)(nil), "autonomic.ext.command.errors.DoorLockStateViolation")
	proto.RegisterType((*GearPositionViolation)(nil), "autonomic.ext.command.errors.GearPositionViolation")
}

func init() {
	proto.RegisterFile("autonomic/ext/command/errors/hardware_precondition.proto", fileDescriptor_f51d9503d4e405bc)
}

var fileDescriptor_f51d9503d4e405bc = []byte{
	// 1035 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x97, 0x51, 0x6f, 0xe3, 0x44,
	0x10, 0xc7, 0xcf, 0xb9, 0xa3, 0x1c, 0xd3, 0xd2, 0xa6, 0xdb, 0x34, 0x2d, 0xed, 0x71, 0x14, 0x0b,
	0x21, 0x5e, 0x1a, 0x73, 0x77, 0xe8, 0x38, 0x74, 0x42, 0x28, 0x49, 0x7d, 0x97, 0xa8, 0x6d, 0x1c,
	0x9c, 0xb4, 0x10, 0x40, 0x32, 0x8e, 0xbd, 0x24, 0x26, 0xb1, 0x37, 0xac, 0xed, 0xb6, 0x91, 0x40,
	0xe2, 0x0d, 0x9e, 0xf8, 0x00, 0xf0, 0x06, 0x9f, 0x82, 0x2f, 0xc5, 0x33, 0xe2, 0x09, 0x79, 0x63,
	0x27, 0xeb, 0xd8, 0x75, 0x7d, 0x7d, 0xf4, 0xcc, 0xfc, 0xe7, 0x97, 0x9d, 0x9d, 0xdd, 0x9d, 0xc0,
	0x33, 0xdd, 0xf7, 0x88, 0x43, 0x6c, 0xcb, 0x90, 0xf0, 0x95, 0x27, 0x19, 0xc4, 0xb6, 0x75, 0xc7,
	0x94, 0x30, 0xa5, 0x84, 0xba, 0xd2, 0x50, 0xa7, 0xe6, 0xa5, 0x4e, 0xb1, 0x36, 0xa1, 0xd8, 0x20,
	0x8e, 0x69, 0x79, 0x16, 0x71, 0x2a, 0x13, 0x4a, 0x3c, 0x82, 0x1e, 0xcc, 0x95, 0x15, 0x7c, 0xe5,
	0x55, 0x42, 0x65, 0x65, 0xa6, 0xdc, 0x7b, 0x94, 0x2f, 0xaf, 0xeb, 0xe9, 0x1e, 0x9e, 0x25, 0xdc,
	0xfb, 0x38, 0x2e, 0xf1, 0xf0, 0x18, 0xdb, 0xd8, 0xa3, 0x53, 0x09, 0x3b, 0xbe, 0x8d, 0xa9, 0x1e,
	0xb0, 0x5d, 0x49, 0x1f, 0xeb, 0xd4, 0x66, 0x3a, 0xdf, 0x0d, 0x85, 0x4f, 0x73, 0x09, 0x4d, 0x42,
	0x68, 0x5c, 0xf7, 0x3c, 0xbf, 0x6e, 0x4c, 0x8c, 0x51, 0x5c, 0xfc, 0x2c, 0x97, 0x18, 0x3b, 0x03,
	0xcb, 0xc1, 0xb7, 0x51, 0x0e, 0xb0, 0x4e, 0xb5, 0x09, 0x71, 0xb9, 0x92, 0x8b, 0xff, 0xdc, 0x87,
	0x72, 0x23, 0x2c, 0x5d, 0x27, 0xa8, 0xdc, 0xb9, 0x45, 0xc6, 0x2c, 0x1e, 0x0d, 0x60, 0x7b, 0x51,
	0x19, 0xac, 0x5d, 0x44, 0x8e, 0x5d, 0xe1, 0x40, 0xf8, 0x60, 0xf5, 0xf1, 0xa3, 0x4a, 0xd6, 0x6e,
	0x55, 0xaa, 0x81, 0x34, 0x9e, 0xb1, 0x71, 0x47, 0xdd, 0xd2, 0x93, 0x66, 0xf4, 0x03, 0xbc, 0x15,
	0xaf, 0x08, 0x0f, 0x2b, 0x30, 0xd8, 0x47, 0xd9, 0xb0, 0x23, 0x42, 0xe8, 0x09, 0x31, 0x46, 0x09,
	0x5e, 0xd9, 0x4c, 0xf5, 0x20, 0x13, 0x4a, 0xf3, 0xcd, 0xe3, 0x69, 0x77, 0x19, 0xed, 0xc3, 0x9b,
	0x69, 0x09, 0x12, 0x32, 0x13, 0x56, 0x74, 0x09, 0xfb, 0x8c, 0x32, 0xd4, 0x1d, 0x73, 0x8c, 0x13,
	0xb0, 0x7b, 0x0c, 0xf6, 0xf4, 0x66, 0x58, 0x83, 0xe9, 0x13, 0xc8, 0x5d, 0xf3, 0x1a, 0x1f, 0xfa,
	0x1e, 0xca, 0x5c, 0x9b, 0xf0, 0xcc, 0xd7, 0x18, 0xf3, 0x71, 0x36, 0x53, 0x66, 0xda, 0x04, 0xaf,
	0x84, 0x53, 0xec, 0xe8, 0x67, 0x01, 0x1e, 0xda, 0xc4, 0x23, 0x54, 0x33, 0x88, 0xe3, 0x51, 0x32,
	0xd6, 0xfa, 0xbe, 0xe7, 0x11, 0x87, 0x83, 0xae, 0x30, 0xe8, 0x27, 0xd9, 0xd0, 0xd3, 0x20, 0x47,
	0x7d, 0x96, 0xa2, 0xc6, 0x32, 0xf0, 0xec, 0x7d, 0xfb, 0x7a, 0x37, 0x9a, 0xc0, 0xee, 0x05, 0x1e,
	0x5a, 0xc6, 0x18, 0x6b, 0x36, 0x09, 0x2c, 0x1c, 0xfb, 0xf5, 0x3c, 0xfd, 0x73, 0x3e, 0x53, 0x9f,
	0x32, 0x71, 0xac, 0x7f, 0x2e, 0x52, 0x3d, 0xe8, 0x6b, 0xd8, 0x1c, 0xe1, 0xa9, 0xf6, 0x1d, 0xe9,
	0x73, 0xa8, 0xfb, 0x0c, 0x75, 0x98, 0x8d, 0x3a, 0xc6, 0xd3, 0x17, 0xa4, 0xcf, 0x33, 0x36, 0x46,
	0x71, 0x53, 0xd0, 0x9c, 0x43, 0x42, 0xcc, 0xc4, 0xde, 0xbd, 0x91, 0xa7, 0x39, 0x1b, 0x84, 0x98,
	0xc9, 0xe6, 0x1c, 0x26, 0xac, 0xc8, 0x86, 0x9d, 0xd8, 0x85, 0xc0, 0x81, 0x80, 0x81, 0x9e, 0x64,
	0x83, 0x5e, 0x62, 0x9d, 0xb6, 0x43, 0x2d, 0xcf, 0xda, 0x1e, 0xa4, 0x39, 0x6a, 0x2b, 0x70, 0xcf,
	0x9b, 0x4e, 0xb0, 0xe8, 0xc3, 0xee, 0x75, 0x2d, 0x8d, 0x7a, 0xb0, 0x99, 0x38, 0x2f, 0xec, 0xb6,
	0x59, 0xbf, 0xa9, 0xaa, 0x4b, 0x29, 0xd5, 0x8d, 0xa5, 0xa3, 0x21, 0xfe, 0x5d, 0x80, 0xfd, 0x8c,
	0x0e, 0x43, 0xbf, 0x09, 0xb0, 0x97, 0xda, 0xc5, 0xfc, 0x8f, 0xf8, 0xfc, 0xd6, 0x1d, 0x9c, 0xe2,
	0x9b, 0xfd, 0xd0, 0x1d, 0x3b, 0xdd, 0x21, 0xfe, 0x2a, 0xc0, 0xce, 0x35, 0x22, 0xf4, 0x3e, 0x88,
	0x67, 0xad, 0xe3, 0x96, 0xf2, 0x45, 0x4b, 0x3b, 0x55, 0xba, 0x8a, 0xaa, 0xd5, 0x95, 0x56, 0x57,
	0x55, 0x4e, 0xb4, 0xda, 0x59, 0xb7, 0xab, 0xb4, 0xb4, 0x4e, 0xb7, 0xda, 0x95, 0x8b, 0x77, 0x90,
	0x08, 0x0f, 0x53, 0xfd, 0x47, 0x72, 0x5b, 0x95, 0x3b, 0x1d, 0xf9, 0xa8, 0x28, 0xa0, 0x77, 0xe1,
	0xed, 0xd4, 0x18, 0x55, 0x3e, 0x91, 0xab, 0x41, 0x48, 0x41, 0xfc, 0x4f, 0x80, 0x72, 0xfa, 0x09,
	0x41, 0x53, 0x28, 0x2d, 0x9d, 0x3c, 0xbe, 0x5e, 0x2f, 0x6f, 0x73, 0xea, 0xe2, 0xe6, 0x59, 0x95,
	0xd0, 0x45, 0xc2, 0x26, 0x62, 0x40, 0xc9, 0x48, 0x74, 0x00, 0x0f, 0xa2, 0xd2, 0x9c, 0xcb, 0x8d,
	0x66, 0xfd, 0x44, 0x0e, 0x4a, 0xd4, 0xe4, 0x8a, 0xb2, 0x0d, 0x9b, 0x91, 0xa7, 0xd9, 0x0a, 0x9d,
	0x45, 0x01, 0x95, 0x01, 0x45, 0xe6, 0x20, 0xb2, 0xa9, 0xb4, 0xaa, 0x6a, 0xaf, 0x58, 0x10, 0xff,
	0x28, 0xc0, 0xc6, 0xd2, 0x99, 0x45, 0x23, 0xd8, 0x98, 0x50, 0x72, 0x65, 0xd9, 0x96, 0x37, 0x8d,
	0x2d, 0xb8, 0xf6, 0x4a, 0x67, 0x3f, 0xfc, 0x6e, 0x47, 0xa9, 0x66, 0x6b, 0x5d, 0x9f, 0xc4, 0xbe,
	0xc5, 0xdf, 0x05, 0x28, 0xa5, 0x05, 0x06, 0xbb, 0x1b, 0x2d, 0xf5, 0x58, 0xee, 0xbd, 0x50, 0x6a,
	0x5a, 0x5b, 0x55, 0xbe, 0x6c, 0x9e, 0x36, 0xbb, 0xbd, 0xf9, 0x62, 0xcb, 0x80, 0x8e, 0xe5, 0x9e,
	0x16, 0x38, 0x9b, 0xf3, 0x8a, 0x14, 0x85, 0xa0, 0x4c, 0x49, 0xfb, 0x42, 0x5f, 0x2c, 0xa0, 0xf7,
	0xe0, 0x20, 0x8a, 0x68, 0x29, 0xdd, 0xf4, 0xa8, 0xbb, 0xa2, 0x0f, 0x28, 0x79, 0xe1, 0x20, 0x0d,
	0xb6, 0xd8, 0x05, 0x36, 0xbf, 0x5a, 0xf8, 0x1a, 0x49, 0x37, 0xdf, 0x5f, 0xd1, 0xed, 0x31, 0x2b,
	0xc8, 0xe6, 0x70, 0xd9, 0x24, 0xfe, 0x08, 0xa5, 0xb4, 0x37, 0x0a, 0x99, 0xf0, 0x66, 0x6c, 0x3c,
	0x0a, 0x91, 0x9f, 0x2d, 0x21, 0xe7, 0xf3, 0x51, 0x85, 0x9f, 0x8f, 0x2a, 0x8b, 0xe7, 0xcd, 0xe7,
	0xdf, 0x40, 0xdf, 0x55, 0xd7, 0x78, 0x97, 0x78, 0x09, 0x5b, 0x29, 0xd3, 0x0d, 0xfa, 0x16, 0xd6,
	0xf8, 0x49, 0x32, 0x64, 0x7f, 0x9a, 0x8f, 0x3d, 0x9f, 0x8b, 0x7c, 0x6e, 0x74, 0xf2, 0x5d, 0x75,
	0x95, 0x73, 0x88, 0x14, 0x50, 0x72, 0xf6, 0x40, 0xdf, 0xc0, 0x2a, 0x37, 0x88, 0x86, 0xd8, 0xe7,
	0xf9, 0xb0, 0xd1, 0xd0, 0xe2, 0x2f, 0xa6, 0x1a, 0xdf, 0x55, 0x61, 0x61, 0x16, 0x7f, 0x11, 0xa0,
	0x9c, 0x3e, 0x5e, 0x21, 0x1b, 0x8a, 0xcb, 0x93, 0x6c, 0x48, 0xaf, 0xe7, 0xa7, 0x47, 0x79, 0xfd,
	0xf8, 0x14, 0xe7, 0xbb, 0xea, 0x7a, 0xdc, 0x2d, 0xfe, 0x04, 0xdb, 0xa9, 0x6f, 0x4e, 0xb0, 0xeb,
	0xb1, 0x97, 0xec, 0xd5, 0x76, 0x9d, 0x7f, 0xae, 0x62, 0x8f, 0x9a, 0xba, 0xc6, 0xbb, 0x6a, 0x7f,
	0x09, 0x70, 0x60, 0x10, 0x3b, 0xb3, 0x7b, 0x6b, 0xef, 0xd4, 0x89, 0x3d, 0x21, 0x0e, 0x76, 0xbc,
	0x36, 0xf7, 0xf7, 0x46, 0x0e, 0x9c, 0xed, 0x60, 0xde, 0x6e, 0x0b, 0x5f, 0x75, 0x06, 0x96, 0x37,
	0xf4, 0xfb, 0x81, 0x52, 0x9a, 0xe7, 0x3a, 0xd4, 0xad, 0x60, 0x72, 0xc7, 0xd4, 0xd1, 0xc7, 0x87,
	0x6c, 0x32, 0x77, 0x25, 0x97, 0x1a, 0x92, 0xad, 0x5b, 0x8e, 0xc4, 0xbe, 0xa5, 0xac, 0x7f, 0x3f,
	0xff, 0x0a, 0xc2, 0x9f, 0x85, 0xbb, 0xd5, 0xb3, 0x6e, 0x7f, 0x85, 0x05, 0x3f, 0xf9, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0xef, 0x0f, 0x0b, 0x41, 0x85, 0x0d, 0x00, 0x00,
}