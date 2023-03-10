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
// source: autonomic/ext/telemetry/dynamics.proto

package telemetry

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

type ThreeAxisSensorState_SensorType int32

const (
	ThreeAxisSensorState_TYPE_UNSPECIFIED   ThreeAxisSensorState_SensorType = 0
	ThreeAxisSensorState_GYRO_DATA          ThreeAxisSensorState_SensorType = 1
	ThreeAxisSensorState_ROTATION_RATE      ThreeAxisSensorState_SensorType = 2
	ThreeAxisSensorState_ACCELEROMETER_DATA ThreeAxisSensorState_SensorType = 3
	ThreeAxisSensorState_ACCELERATION_RATE  ThreeAxisSensorState_SensorType = 4
	ThreeAxisSensorState_MAGNETOMETER_DATA  ThreeAxisSensorState_SensorType = 5
	ThreeAxisSensorState_MAGNETIC_FIELD     ThreeAxisSensorState_SensorType = 6
)

var ThreeAxisSensorState_SensorType_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "GYRO_DATA",
	2: "ROTATION_RATE",
	3: "ACCELEROMETER_DATA",
	4: "ACCELERATION_RATE",
	5: "MAGNETOMETER_DATA",
	6: "MAGNETIC_FIELD",
}

var ThreeAxisSensorState_SensorType_value = map[string]int32{
	"TYPE_UNSPECIFIED":   0,
	"GYRO_DATA":          1,
	"ROTATION_RATE":      2,
	"ACCELEROMETER_DATA": 3,
	"ACCELERATION_RATE":  4,
	"MAGNETOMETER_DATA":  5,
	"MAGNETIC_FIELD":     6,
}

func (x ThreeAxisSensorState_SensorType) String() string {
	return proto.EnumName(ThreeAxisSensorState_SensorType_name, int32(x))
}

func (ThreeAxisSensorState_SensorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_22e8d1dcc70558dd, []int{4, 0}
}

// https://en.wikipedia.org/wiki/Quaternions_and_spatial_rotation
type Quaternion struct {
	X                    float64  `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64  `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float64  `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
	W                    float64  `protobuf:"fixed64,4,opt,name=w,proto3" json:"w,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Quaternion) Reset()         { *m = Quaternion{} }
func (m *Quaternion) String() string { return proto.CompactTextString(m) }
func (*Quaternion) ProtoMessage()    {}
func (*Quaternion) Descriptor() ([]byte, []int) {
	return fileDescriptor_22e8d1dcc70558dd, []int{0}
}

func (m *Quaternion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quaternion.Unmarshal(m, b)
}
func (m *Quaternion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quaternion.Marshal(b, m, deterministic)
}
func (m *Quaternion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quaternion.Merge(m, src)
}
func (m *Quaternion) XXX_Size() int {
	return xxx_messageInfo_Quaternion.Size(m)
}
func (m *Quaternion) XXX_DiscardUnknown() {
	xxx_messageInfo_Quaternion.DiscardUnknown(m)
}

var xxx_messageInfo_Quaternion proto.InternalMessageInfo

func (m *Quaternion) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Quaternion) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Quaternion) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Quaternion) GetW() float64 {
	if m != nil {
		return m.W
	}
	return 0
}

// https://en.wikipedia.org/wiki/Euler_angles
type EulerAngles struct {
	Roll                 float64  `protobuf:"fixed64,1,opt,name=roll,proto3" json:"roll,omitempty"`
	Pitch                float64  `protobuf:"fixed64,2,opt,name=pitch,proto3" json:"pitch,omitempty"`
	Yaw                  float64  `protobuf:"fixed64,3,opt,name=yaw,proto3" json:"yaw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EulerAngles) Reset()         { *m = EulerAngles{} }
func (m *EulerAngles) String() string { return proto.CompactTextString(m) }
func (*EulerAngles) ProtoMessage()    {}
func (*EulerAngles) Descriptor() ([]byte, []int) {
	return fileDescriptor_22e8d1dcc70558dd, []int{1}
}

func (m *EulerAngles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EulerAngles.Unmarshal(m, b)
}
func (m *EulerAngles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EulerAngles.Marshal(b, m, deterministic)
}
func (m *EulerAngles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EulerAngles.Merge(m, src)
}
func (m *EulerAngles) XXX_Size() int {
	return xxx_messageInfo_EulerAngles.Size(m)
}
func (m *EulerAngles) XXX_DiscardUnknown() {
	xxx_messageInfo_EulerAngles.DiscardUnknown(m)
}

var xxx_messageInfo_EulerAngles proto.InternalMessageInfo

func (m *EulerAngles) GetRoll() float64 {
	if m != nil {
		return m.Roll
	}
	return 0
}

func (m *EulerAngles) GetPitch() float64 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

func (m *EulerAngles) GetYaw() float64 {
	if m != nil {
		return m.Yaw
	}
	return 0
}

// https://en.wikipedia.org/wiki/Rotation_matrix#Quaternion
type RotationMatrix struct {
	M11                  float64  `protobuf:"fixed64,1,opt,name=m11,proto3" json:"m11,omitempty"`
	M12                  float64  `protobuf:"fixed64,2,opt,name=m12,proto3" json:"m12,omitempty"`
	M13                  float64  `protobuf:"fixed64,3,opt,name=m13,proto3" json:"m13,omitempty"`
	M21                  float64  `protobuf:"fixed64,4,opt,name=m21,proto3" json:"m21,omitempty"`
	M22                  float64  `protobuf:"fixed64,5,opt,name=m22,proto3" json:"m22,omitempty"`
	M23                  float64  `protobuf:"fixed64,6,opt,name=m23,proto3" json:"m23,omitempty"`
	M31                  float64  `protobuf:"fixed64,7,opt,name=m31,proto3" json:"m31,omitempty"`
	M32                  float64  `protobuf:"fixed64,8,opt,name=m32,proto3" json:"m32,omitempty"`
	M33                  float64  `protobuf:"fixed64,9,opt,name=m33,proto3" json:"m33,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RotationMatrix) Reset()         { *m = RotationMatrix{} }
func (m *RotationMatrix) String() string { return proto.CompactTextString(m) }
func (*RotationMatrix) ProtoMessage()    {}
func (*RotationMatrix) Descriptor() ([]byte, []int) {
	return fileDescriptor_22e8d1dcc70558dd, []int{2}
}

func (m *RotationMatrix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RotationMatrix.Unmarshal(m, b)
}
func (m *RotationMatrix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RotationMatrix.Marshal(b, m, deterministic)
}
func (m *RotationMatrix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RotationMatrix.Merge(m, src)
}
func (m *RotationMatrix) XXX_Size() int {
	return xxx_messageInfo_RotationMatrix.Size(m)
}
func (m *RotationMatrix) XXX_DiscardUnknown() {
	xxx_messageInfo_RotationMatrix.DiscardUnknown(m)
}

var xxx_messageInfo_RotationMatrix proto.InternalMessageInfo

func (m *RotationMatrix) GetM11() float64 {
	if m != nil {
		return m.M11
	}
	return 0
}

func (m *RotationMatrix) GetM12() float64 {
	if m != nil {
		return m.M12
	}
	return 0
}

func (m *RotationMatrix) GetM13() float64 {
	if m != nil {
		return m.M13
	}
	return 0
}

func (m *RotationMatrix) GetM21() float64 {
	if m != nil {
		return m.M21
	}
	return 0
}

func (m *RotationMatrix) GetM22() float64 {
	if m != nil {
		return m.M22
	}
	return 0
}

func (m *RotationMatrix) GetM23() float64 {
	if m != nil {
		return m.M23
	}
	return 0
}

func (m *RotationMatrix) GetM31() float64 {
	if m != nil {
		return m.M31
	}
	return 0
}

func (m *RotationMatrix) GetM32() float64 {
	if m != nil {
		return m.M32
	}
	return 0
}

func (m *RotationMatrix) GetM33() float64 {
	if m != nil {
		return m.M33
	}
	return 0
}

type Orientation struct {
	// Types that are valid to be assigned to Orientation:
	//	*Orientation_Quaternion
	//	*Orientation_EulerAngles
	//	*Orientation_RotationMatrix
	Orientation          isOrientation_Orientation `protobuf_oneof:"orientation"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Orientation) Reset()         { *m = Orientation{} }
func (m *Orientation) String() string { return proto.CompactTextString(m) }
func (*Orientation) ProtoMessage()    {}
func (*Orientation) Descriptor() ([]byte, []int) {
	return fileDescriptor_22e8d1dcc70558dd, []int{3}
}

func (m *Orientation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Orientation.Unmarshal(m, b)
}
func (m *Orientation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Orientation.Marshal(b, m, deterministic)
}
func (m *Orientation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Orientation.Merge(m, src)
}
func (m *Orientation) XXX_Size() int {
	return xxx_messageInfo_Orientation.Size(m)
}
func (m *Orientation) XXX_DiscardUnknown() {
	xxx_messageInfo_Orientation.DiscardUnknown(m)
}

var xxx_messageInfo_Orientation proto.InternalMessageInfo

type isOrientation_Orientation interface {
	isOrientation_Orientation()
}

type Orientation_Quaternion struct {
	Quaternion *Quaternion `protobuf:"bytes,1,opt,name=quaternion,proto3,oneof"`
}

type Orientation_EulerAngles struct {
	EulerAngles *EulerAngles `protobuf:"bytes,2,opt,name=euler_angles,json=eulerAngles,proto3,oneof"`
}

type Orientation_RotationMatrix struct {
	RotationMatrix *RotationMatrix `protobuf:"bytes,3,opt,name=rotation_matrix,json=rotationMatrix,proto3,oneof"`
}

func (*Orientation_Quaternion) isOrientation_Orientation() {}

func (*Orientation_EulerAngles) isOrientation_Orientation() {}

func (*Orientation_RotationMatrix) isOrientation_Orientation() {}

func (m *Orientation) GetOrientation() isOrientation_Orientation {
	if m != nil {
		return m.Orientation
	}
	return nil
}

func (m *Orientation) GetQuaternion() *Quaternion {
	if x, ok := m.GetOrientation().(*Orientation_Quaternion); ok {
		return x.Quaternion
	}
	return nil
}

func (m *Orientation) GetEulerAngles() *EulerAngles {
	if x, ok := m.GetOrientation().(*Orientation_EulerAngles); ok {
		return x.EulerAngles
	}
	return nil
}

func (m *Orientation) GetRotationMatrix() *RotationMatrix {
	if x, ok := m.GetOrientation().(*Orientation_RotationMatrix); ok {
		return x.RotationMatrix
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Orientation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Orientation_Quaternion)(nil),
		(*Orientation_EulerAngles)(nil),
		(*Orientation_RotationMatrix)(nil),
	}
}

// Rotation, measured by gyros, is expressed in rad/s
// Acceleration, measured by accelerometers, is expressed in m/s^2
// Magnetic field, measured by magnetometers, is expressed in uTesla
//
// Raw sensor output, meaning the data has not beet filtered, smoothed,
// or otherwise processed to remove noise, is provided by SensorType
// values ending in _DATA. Processed data is provided by corresponding
// SensorType (e.g., ROTATION_RATE corresponds to GYRO_DATA and
// MAGNETIC_FIELD corresponds to MAGNETOMETER_DATA).
type ThreeAxisSensorState struct {
	SensorType           ThreeAxisSensorState_SensorType `protobuf:"varint,2,opt,name=sensor_type,json=sensorType,proto3,enum=autonomic.ext.telemetry.ThreeAxisSensorState_SensorType" json:"sensor_type,omitempty"`
	X                    float64                         `protobuf:"fixed64,3,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64                         `protobuf:"fixed64,4,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float64                         `protobuf:"fixed64,5,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ThreeAxisSensorState) Reset()         { *m = ThreeAxisSensorState{} }
func (m *ThreeAxisSensorState) String() string { return proto.CompactTextString(m) }
func (*ThreeAxisSensorState) ProtoMessage()    {}
func (*ThreeAxisSensorState) Descriptor() ([]byte, []int) {
	return fileDescriptor_22e8d1dcc70558dd, []int{4}
}

func (m *ThreeAxisSensorState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreeAxisSensorState.Unmarshal(m, b)
}
func (m *ThreeAxisSensorState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreeAxisSensorState.Marshal(b, m, deterministic)
}
func (m *ThreeAxisSensorState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreeAxisSensorState.Merge(m, src)
}
func (m *ThreeAxisSensorState) XXX_Size() int {
	return xxx_messageInfo_ThreeAxisSensorState.Size(m)
}
func (m *ThreeAxisSensorState) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreeAxisSensorState.DiscardUnknown(m)
}

var xxx_messageInfo_ThreeAxisSensorState proto.InternalMessageInfo

func (m *ThreeAxisSensorState) GetSensorType() ThreeAxisSensorState_SensorType {
	if m != nil {
		return m.SensorType
	}
	return ThreeAxisSensorState_TYPE_UNSPECIFIED
}

func (m *ThreeAxisSensorState) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ThreeAxisSensorState) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *ThreeAxisSensorState) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.ThreeAxisSensorState_SensorType", ThreeAxisSensorState_SensorType_name, ThreeAxisSensorState_SensorType_value)
	proto.RegisterType((*Quaternion)(nil), "autonomic.ext.telemetry.Quaternion")
	proto.RegisterType((*EulerAngles)(nil), "autonomic.ext.telemetry.EulerAngles")
	proto.RegisterType((*RotationMatrix)(nil), "autonomic.ext.telemetry.RotationMatrix")
	proto.RegisterType((*Orientation)(nil), "autonomic.ext.telemetry.Orientation")
	proto.RegisterType((*ThreeAxisSensorState)(nil), "autonomic.ext.telemetry.ThreeAxisSensorState")
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/dynamics.proto", fileDescriptor_22e8d1dcc70558dd)
}

var fileDescriptor_22e8d1dcc70558dd = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xbd, 0x6e, 0xdb, 0x3c,
	0x14, 0x35, 0xfd, 0x93, 0xef, 0xcb, 0x55, 0xec, 0x3a, 0x44, 0xda, 0x0a, 0xe8, 0x52, 0xb8, 0x45,
	0xdb, 0x25, 0x12, 0x2c, 0x2d, 0x5d, 0x15, 0x9b, 0x49, 0x04, 0x24, 0xb6, 0xcb, 0x28, 0x43, 0xba,
	0x08, 0x8a, 0x4b, 0x24, 0x02, 0x2c, 0xd1, 0xa5, 0x68, 0xc4, 0xca, 0xd8, 0xa5, 0x6f, 0xd1, 0xa1,
	0x8f, 0xd0, 0xb9, 0x0f, 0xd6, 0xb1, 0x20, 0x29, 0x3b, 0x36, 0x50, 0x6f, 0xf7, 0x1c, 0x9c, 0x7b,
	0x44, 0x9e, 0x7b, 0x29, 0x78, 0x97, 0x2c, 0x24, 0xcf, 0x79, 0x96, 0x4e, 0x5d, 0xb6, 0x94, 0xae,
	0x64, 0x33, 0x96, 0x31, 0x29, 0x4a, 0xf7, 0x4b, 0x99, 0x27, 0x59, 0x3a, 0x2d, 0x9c, 0xb9, 0xe0,
	0x92, 0xe3, 0x97, 0x6b, 0x9d, 0xc3, 0x96, 0xd2, 0x59, 0xeb, 0x7a, 0x43, 0x80, 0x4f, 0x8b, 0x44,
	0x32, 0x91, 0xa7, 0x3c, 0xc7, 0x07, 0x80, 0x96, 0x36, 0x7a, 0x8d, 0x3e, 0x20, 0x8a, 0x96, 0x0a,
	0x95, 0x76, 0xdd, 0xa0, 0x52, 0xa1, 0x47, 0xbb, 0x61, 0xd0, 0xa3, 0x42, 0x0f, 0x76, 0xd3, 0xa0,
	0x87, 0x5e, 0x08, 0x16, 0x59, 0xcc, 0x98, 0x08, 0xf2, 0xbb, 0x19, 0x2b, 0x30, 0x86, 0xa6, 0xe0,
	0xb3, 0x59, 0xe5, 0xa4, 0x6b, 0x7c, 0x04, 0xad, 0x79, 0x2a, 0xa7, 0xf7, 0x95, 0xa1, 0x01, 0xb8,
	0x0b, 0x8d, 0x32, 0x79, 0xa8, 0x6c, 0x55, 0xd9, 0xfb, 0x85, 0xa0, 0x43, 0xb9, 0x4c, 0x64, 0xca,
	0xf3, 0xcb, 0x44, 0x8a, 0x74, 0xa9, 0x44, 0x59, 0xbf, 0x5f, 0xb9, 0xa9, 0xd2, 0x30, 0x5e, 0x65,
	0xa5, 0x4a, 0xc3, 0xf8, 0x2b, 0xa3, 0xac, 0xef, 0x6b, 0xc6, 0xeb, 0x57, 0x67, 0x54, 0xa5, 0x61,
	0x3c, 0xbb, 0xb5, 0x62, 0x4c, 0x97, 0xe7, 0xdb, 0x7b, 0x2b, 0xc6, 0x74, 0xf9, 0x7d, 0xfb, 0xbf,
	0x8a, 0xf1, 0x4d, 0x97, 0xef, 0xd9, 0xff, 0xaf, 0x18, 0xd3, 0xe5, 0xfb, 0xf6, 0xfe, 0x8a, 0xf1,
	0x7b, 0xdf, 0xea, 0x60, 0x8d, 0x45, 0xca, 0x72, 0x73, 0x6e, 0x4c, 0x00, 0xbe, 0xae, 0x53, 0xd5,
	0x07, 0xb7, 0xbc, 0x37, 0xce, 0x8e, 0x19, 0x38, 0x4f, 0x03, 0x38, 0xaf, 0xd1, 0x8d, 0x46, 0x1c,
	0xc2, 0x01, 0x53, 0xb1, 0xc6, 0x89, 0xce, 0x55, 0xdf, 0xd7, 0xf2, 0xde, 0xee, 0x34, 0xda, 0x98,
	0xc1, 0x79, 0x8d, 0x5a, 0x6c, 0x63, 0x24, 0x14, 0x9e, 0x89, 0x2a, 0xd5, 0x38, 0xd3, 0xb1, 0xea,
	0xac, 0x2c, 0xef, 0xfd, 0x4e, 0xb7, 0xed, 0x29, 0x9c, 0xd7, 0x68, 0x47, 0x6c, 0x31, 0x27, 0x6d,
	0xb0, 0xf8, 0xd3, 0xa5, 0x7b, 0xbf, 0xeb, 0x70, 0x14, 0xdd, 0x0b, 0xc6, 0x82, 0x65, 0x5a, 0x5c,
	0xb1, 0xbc, 0xe0, 0xe2, 0x4a, 0x26, 0x92, 0xe1, 0x1b, 0xb0, 0x0a, 0x0d, 0x63, 0x59, 0xce, 0x99,
	0xbe, 0x45, 0xc7, 0xfb, 0xb8, 0xf3, 0xbb, 0xff, 0xf2, 0x70, 0x4c, 0x1d, 0x95, 0x73, 0x46, 0xa1,
	0x58, 0xd7, 0x66, 0x61, 0x1b, 0x5b, 0x0b, 0xdb, 0xdc, 0x5a, 0xd8, 0x56, 0xb5, 0xb0, 0xbd, 0x1f,
	0x08, 0xe0, 0xc9, 0x04, 0x1f, 0x41, 0x37, 0xba, 0x99, 0x90, 0xf8, 0x7a, 0x74, 0x35, 0x21, 0x83,
	0xf0, 0x34, 0x24, 0xc3, 0x6e, 0x0d, 0xb7, 0x61, 0xff, 0xec, 0x86, 0x8e, 0xe3, 0x61, 0x10, 0x05,
	0x5d, 0x84, 0x0f, 0xa1, 0x4d, 0xc7, 0x51, 0x10, 0x85, 0xe3, 0x51, 0x4c, 0x83, 0x88, 0x74, 0xeb,
	0xf8, 0x05, 0xe0, 0x60, 0x30, 0x20, 0x17, 0x84, 0x8e, 0x2f, 0x49, 0x44, 0xa8, 0x91, 0x36, 0xf0,
	0x73, 0x38, 0xac, 0xf8, 0x0d, 0x79, 0x53, 0xd1, 0x97, 0xc1, 0xd9, 0x88, 0x44, 0x9b, 0xea, 0x16,
	0xc6, 0xd0, 0x31, 0x74, 0x38, 0x88, 0x4f, 0x43, 0x72, 0x31, 0xec, 0xee, 0x9d, 0x7c, 0x47, 0xf0,
	0x6a, 0xca, 0xb3, 0x5d, 0xb1, 0x9c, 0xb4, 0x87, 0xd5, 0x93, 0x9e, 0xa8, 0x17, 0x3d, 0x41, 0x9f,
	0x47, 0x77, 0xa9, 0xbc, 0x5f, 0xdc, 0x3a, 0x53, 0x9e, 0xb9, 0xeb, 0xa6, 0xe3, 0x24, 0x55, 0x7f,
	0x02, 0x26, 0xf2, 0x64, 0x76, 0xac, 0xdf, 0x7e, 0xe1, 0x16, 0x62, 0xea, 0x66, 0x49, 0x9a, 0xbb,
	0x1a, 0xbb, 0x3b, 0x7e, 0x19, 0x7f, 0x10, 0xfa, 0x59, 0x6f, 0x04, 0xd7, 0xd1, 0xed, 0x9e, 0xd6,
	0xf9, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x34, 0xe3, 0x1d, 0x9f, 0x5d, 0x04, 0x00, 0x00,
}
