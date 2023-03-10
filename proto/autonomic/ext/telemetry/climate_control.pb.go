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
// source: autonomic/ext/telemetry/climate_control.proto

package telemetry

import (
	fmt "fmt"
	//enumerations "xk6-fabric/proto/autonomic/ext/telemetry/enumerations"
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

type AirFlow_AirVentDistribution int32

const (
	AirFlow_UNKNOWN_AIR_VENT          AirFlow_AirVentDistribution = 0
	AirFlow_FLOOR_AIR_VENT            AirFlow_AirVentDistribution = 1
	AirFlow_INSTRUMENT_PANEL_AIR_VENT AirFlow_AirVentDistribution = 2
	AirFlow_DEFROST_AIR_VENT          AirFlow_AirVentDistribution = 3
)

var AirFlow_AirVentDistribution_name = map[int32]string{
	0: "UNKNOWN_AIR_VENT",
	1: "FLOOR_AIR_VENT",
	2: "INSTRUMENT_PANEL_AIR_VENT",
	3: "DEFROST_AIR_VENT",
}

var AirFlow_AirVentDistribution_value = map[string]int32{
	"UNKNOWN_AIR_VENT":          0,
	"FLOOR_AIR_VENT":            1,
	"INSTRUMENT_PANEL_AIR_VENT": 2,
	"DEFROST_AIR_VENT":          3,
}

func (x AirFlow_AirVentDistribution) String() string {
	return proto.EnumName(AirFlow_AirVentDistribution_name, int32(x))
}

func (AirFlow_AirVentDistribution) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0268bf8a2e0da29f, []int{0, 0}
}

type AirFlow_AirCirculationStatus int32

const (
	AirFlow_UNKNOWN_AIR_CIRCULATION AirFlow_AirCirculationStatus = 0
	AirFlow_FRESH_INTAKE            AirFlow_AirCirculationStatus = 1
	AirFlow_RECIRCULATE             AirFlow_AirCirculationStatus = 2
)

var AirFlow_AirCirculationStatus_name = map[int32]string{
	0: "UNKNOWN_AIR_CIRCULATION",
	1: "FRESH_INTAKE",
	2: "RECIRCULATE",
}

var AirFlow_AirCirculationStatus_value = map[string]int32{
	"UNKNOWN_AIR_CIRCULATION": 0,
	"FRESH_INTAKE":            1,
	"RECIRCULATE":             2,
}

func (x AirFlow_AirCirculationStatus) String() string {
	return proto.EnumName(AirFlow_AirCirculationStatus_name, int32(x))
}

func (AirFlow_AirCirculationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0268bf8a2e0da29f, []int{0, 1}
}

type ClimateControl_AirConditionerStatus int32

const (
	ClimateControl_UNKNOWN_AIR_CONDITIONER_STATUS ClimateControl_AirConditionerStatus = 0
	ClimateControl_OFF                            ClimateControl_AirConditionerStatus = 1
	ClimateControl_ON                             ClimateControl_AirConditionerStatus = 2
)

var ClimateControl_AirConditionerStatus_name = map[int32]string{
	0: "UNKNOWN_AIR_CONDITIONER_STATUS",
	1: "OFF",
	2: "ON",
}

var ClimateControl_AirConditionerStatus_value = map[string]int32{
	"UNKNOWN_AIR_CONDITIONER_STATUS": 0,
	"OFF":                            1,
	"ON":                             2,
}

func (x ClimateControl_AirConditionerStatus) String() string {
	return proto.EnumName(ClimateControl_AirConditionerStatus_name, int32(x))
}

func (ClimateControl_AirConditionerStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0268bf8a2e0da29f, []int{2, 0}
}

type ClimateControl_SystemControlMode int32

const (
	ClimateControl_UNKNOWN_SYSTEM_CONTROL_MODE   ClimateControl_SystemControlMode = 0
	ClimateControl_MANUAL_SYSTEM_CONTROL_MODE    ClimateControl_SystemControlMode = 1
	ClimateControl_AUTOMATIC_SYSTEM_CONTROL_MODE ClimateControl_SystemControlMode = 2
)

var ClimateControl_SystemControlMode_name = map[int32]string{
	0: "UNKNOWN_SYSTEM_CONTROL_MODE",
	1: "MANUAL_SYSTEM_CONTROL_MODE",
	2: "AUTOMATIC_SYSTEM_CONTROL_MODE",
}

var ClimateControl_SystemControlMode_value = map[string]int32{
	"UNKNOWN_SYSTEM_CONTROL_MODE":   0,
	"MANUAL_SYSTEM_CONTROL_MODE":    1,
	"AUTOMATIC_SYSTEM_CONTROL_MODE": 2,
}

func (x ClimateControl_SystemControlMode) String() string {
	return proto.EnumName(ClimateControl_SystemControlMode_name, int32(x))
}

func (ClimateControl_SystemControlMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0268bf8a2e0da29f, []int{2, 1}
}

type AirFlow struct {
	FanSpeed             float64                       `protobuf:"fixed64,1,opt,name=fan_speed,json=fanSpeed,proto3" json:"fan_speed,omitempty"`
	ActiveVents          []AirFlow_AirVentDistribution `protobuf:"varint,2,rep,packed,name=active_vents,json=activeVents,proto3,enum=autonomic.ext.telemetry.AirFlow_AirVentDistribution" json:"active_vents,omitempty"`
	AirCirculationStatus AirFlow_AirCirculationStatus  `protobuf:"varint,3,opt,name=air_circulation_status,json=airCirculationStatus,proto3,enum=autonomic.ext.telemetry.AirFlow_AirCirculationStatus" json:"air_circulation_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AirFlow) Reset()         { *m = AirFlow{} }
func (m *AirFlow) String() string { return proto.CompactTextString(m) }
func (*AirFlow) ProtoMessage()    {}
func (*AirFlow) Descriptor() ([]byte, []int) {
	return fileDescriptor_0268bf8a2e0da29f, []int{0}
}

func (m *AirFlow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AirFlow.Unmarshal(m, b)
}
func (m *AirFlow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AirFlow.Marshal(b, m, deterministic)
}
func (m *AirFlow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AirFlow.Merge(m, src)
}
func (m *AirFlow) XXX_Size() int {
	return xxx_messageInfo_AirFlow.Size(m)
}
func (m *AirFlow) XXX_DiscardUnknown() {
	xxx_messageInfo_AirFlow.DiscardUnknown(m)
}

var xxx_messageInfo_AirFlow proto.InternalMessageInfo

func (m *AirFlow) GetFanSpeed() float64 {
	if m != nil {
		return m.FanSpeed
	}
	return 0
}

func (m *AirFlow) GetActiveVents() []AirFlow_AirVentDistribution {
	if m != nil {
		return m.ActiveVents
	}
	return nil
}

func (m *AirFlow) GetAirCirculationStatus() AirFlow_AirCirculationStatus {
	if m != nil {
		return m.AirCirculationStatus
	}
	return AirFlow_UNKNOWN_AIR_CIRCULATION
}

type CabinTemperatureSetting struct {
	// Types that are valid to be assigned to TemperatureControl:
	//	*CabinTemperatureSetting_MeasuredTemperatureSetting
	//	*CabinTemperatureSetting_RelativeTemperatureSetting
	TemperatureControl   isCabinTemperatureSetting_TemperatureControl `protobuf_oneof:"temperature_control"`
	UnitTag              *Tag                                         `protobuf:"bytes,3,opt,name=unit_tag,json=unitTag,proto3" json:"unit_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *CabinTemperatureSetting) Reset()         { *m = CabinTemperatureSetting{} }
func (m *CabinTemperatureSetting) String() string { return proto.CompactTextString(m) }
func (*CabinTemperatureSetting) ProtoMessage()    {}
func (*CabinTemperatureSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_0268bf8a2e0da29f, []int{1}
}

func (m *CabinTemperatureSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CabinTemperatureSetting.Unmarshal(m, b)
}
func (m *CabinTemperatureSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CabinTemperatureSetting.Marshal(b, m, deterministic)
}
func (m *CabinTemperatureSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CabinTemperatureSetting.Merge(m, src)
}
func (m *CabinTemperatureSetting) XXX_Size() int {
	return xxx_messageInfo_CabinTemperatureSetting.Size(m)
}
func (m *CabinTemperatureSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CabinTemperatureSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CabinTemperatureSetting proto.InternalMessageInfo

type isCabinTemperatureSetting_TemperatureControl interface {
	isCabinTemperatureSetting_TemperatureControl()
}

type CabinTemperatureSetting_MeasuredTemperatureSetting struct {
	MeasuredTemperatureSetting float64 `protobuf:"fixed64,1,opt,name=measured_temperature_setting,json=measuredTemperatureSetting,proto3,oneof"`
}

type CabinTemperatureSetting_RelativeTemperatureSetting struct {
	RelativeTemperatureSetting float64 `protobuf:"fixed64,2,opt,name=relative_temperature_setting,json=relativeTemperatureSetting,proto3,oneof"`
}

func (*CabinTemperatureSetting_MeasuredTemperatureSetting) isCabinTemperatureSetting_TemperatureControl() {
}

func (*CabinTemperatureSetting_RelativeTemperatureSetting) isCabinTemperatureSetting_TemperatureControl() {
}

func (m *CabinTemperatureSetting) GetTemperatureControl() isCabinTemperatureSetting_TemperatureControl {
	if m != nil {
		return m.TemperatureControl
	}
	return nil
}

func (m *CabinTemperatureSetting) GetMeasuredTemperatureSetting() float64 {
	if x, ok := m.GetTemperatureControl().(*CabinTemperatureSetting_MeasuredTemperatureSetting); ok {
		return x.MeasuredTemperatureSetting
	}
	return 0
}

func (m *CabinTemperatureSetting) GetRelativeTemperatureSetting() float64 {
	if x, ok := m.GetTemperatureControl().(*CabinTemperatureSetting_RelativeTemperatureSetting); ok {
		return x.RelativeTemperatureSetting
	}
	return 0
}

func (m *CabinTemperatureSetting) GetUnitTag() *Tag {
	if m != nil {
		return m.UnitTag
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CabinTemperatureSetting) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CabinTemperatureSetting_MeasuredTemperatureSetting)(nil),
		(*CabinTemperatureSetting_RelativeTemperatureSetting)(nil),
	}
}

type ClimateControl struct {
	// Climate Control System Status
	SystemStatus enumerations.OffOnStatus `protobuf:"varint,1,opt,name=system_status,json=systemStatus,proto3,enum=autonomic.ext.telemetry.enumerations.offOnStatus.OffOnStatus" json:"system_status,omitempty"`
	// Flow of air from climate control systems
	AirFlow *AirFlow `protobuf:"bytes,2,opt,name=air_flow,json=airFlow,proto3" json:"air_flow,omitempty"`
	// Active cabin temperature setting
	CabinTemperatureSetting *CabinTemperatureSetting `protobuf:"bytes,3,opt,name=cabin_temperature_setting,json=cabinTemperatureSetting,proto3" json:"cabin_temperature_setting,omitempty"`
	// Air conditioner status
	AirConditionerStatus ClimateControl_AirConditionerStatus `protobuf:"varint,4,opt,name=air_conditioner_status,json=airConditionerStatus,proto3,enum=autonomic.ext.telemetry.ClimateControl_AirConditionerStatus" json:"air_conditioner_status,omitempty"`
	SystemControlMode    ClimateControl_SystemControlMode    `protobuf:"varint,5,opt,name=system_control_mode,json=systemControlMode,proto3,enum=autonomic.ext.telemetry.ClimateControl_SystemControlMode" json:"system_control_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ClimateControl) Reset()         { *m = ClimateControl{} }
func (m *ClimateControl) String() string { return proto.CompactTextString(m) }
func (*ClimateControl) ProtoMessage()    {}
func (*ClimateControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_0268bf8a2e0da29f, []int{2}
}

func (m *ClimateControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClimateControl.Unmarshal(m, b)
}
func (m *ClimateControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClimateControl.Marshal(b, m, deterministic)
}
func (m *ClimateControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClimateControl.Merge(m, src)
}
func (m *ClimateControl) XXX_Size() int {
	return xxx_messageInfo_ClimateControl.Size(m)
}
func (m *ClimateControl) XXX_DiscardUnknown() {
	xxx_messageInfo_ClimateControl.DiscardUnknown(m)
}

var xxx_messageInfo_ClimateControl proto.InternalMessageInfo

func (m *ClimateControl) GetSystemStatus() enumerations.OffOnStatus {
	if m != nil {
		return m.SystemStatus
	}
	return enumerations.OffOnStatus_UNKNOWN
}

func (m *ClimateControl) GetAirFlow() *AirFlow {
	if m != nil {
		return m.AirFlow
	}
	return nil
}

func (m *ClimateControl) GetCabinTemperatureSetting() *CabinTemperatureSetting {
	if m != nil {
		return m.CabinTemperatureSetting
	}
	return nil
}

func (m *ClimateControl) GetAirConditionerStatus() ClimateControl_AirConditionerStatus {
	if m != nil {
		return m.AirConditionerStatus
	}
	return ClimateControl_UNKNOWN_AIR_CONDITIONER_STATUS
}

func (m *ClimateControl) GetSystemControlMode() ClimateControl_SystemControlMode {
	if m != nil {
		return m.SystemControlMode
	}
	return ClimateControl_UNKNOWN_SYSTEM_CONTROL_MODE
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.AirFlow_AirVentDistribution", AirFlow_AirVentDistribution_name, AirFlow_AirVentDistribution_value)
	proto.RegisterEnum("autonomic.ext.telemetry.AirFlow_AirCirculationStatus", AirFlow_AirCirculationStatus_name, AirFlow_AirCirculationStatus_value)
	proto.RegisterEnum("autonomic.ext.telemetry.ClimateControl_AirConditionerStatus", ClimateControl_AirConditionerStatus_name, ClimateControl_AirConditionerStatus_value)
	proto.RegisterEnum("autonomic.ext.telemetry.ClimateControl_SystemControlMode", ClimateControl_SystemControlMode_name, ClimateControl_SystemControlMode_value)
	proto.RegisterType((*AirFlow)(nil), "autonomic.ext.telemetry.AirFlow")
	proto.RegisterType((*CabinTemperatureSetting)(nil), "autonomic.ext.telemetry.CabinTemperatureSetting")
	proto.RegisterType((*ClimateControl)(nil), "autonomic.ext.telemetry.ClimateControl")
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/climate_control.proto", fileDescriptor_0268bf8a2e0da29f)
}

var fileDescriptor_0268bf8a2e0da29f = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0x1d, 0x76, 0x53, 0xa6, 0x25, 0x78, 0x27, 0x85, 0x64, 0x13, 0x76, 0xc9, 0xfa, 0xaa,
	0x37, 0xb5, 0x51, 0x01, 0x21, 0xb4, 0x70, 0xe1, 0x24, 0x8e, 0x36, 0x6a, 0x62, 0x57, 0xe3, 0x49,
	0x57, 0x70, 0x63, 0x4d, 0x9c, 0x49, 0x18, 0x61, 0x7b, 0x2a, 0x7b, 0xdc, 0xee, 0x3e, 0x04, 0x2f,
	0xc1, 0x63, 0x70, 0xcd, 0x2b, 0x21, 0x71, 0x89, 0xc6, 0x76, 0x9a, 0x74, 0xd7, 0x46, 0xbd, 0x4a,
	0xe6, 0xcc, 0x39, 0xdf, 0xf9, 0xf9, 0xce, 0x37, 0x06, 0xe7, 0x24, 0x13, 0x3c, 0xe6, 0x11, 0x0b,
	0x4c, 0xfa, 0x4e, 0x98, 0x82, 0x86, 0x34, 0xa2, 0x22, 0x79, 0x6f, 0x06, 0x21, 0x8b, 0x88, 0xa0,
	0x7e, 0xc0, 0x63, 0x91, 0xf0, 0xd0, 0xb8, 0x49, 0xb8, 0xe0, 0xb0, 0x7b, 0xef, 0x6e, 0xd0, 0x77,
	0xc2, 0xb8, 0x77, 0xef, 0xff, 0x50, 0x87, 0x43, 0xe3, 0x2c, 0xa2, 0x09, 0x11, 0x8c, 0xc7, 0xa9,
	0xc9, 0x37, 0x1b, 0x1e, 0xfb, 0xa9, 0x20, 0x22, 0x4b, 0x0b, 0xc4, 0xfe, 0xab, 0xba, 0x40, 0x41,
	0xb6, 0x85, 0x8b, 0xfe, 0x57, 0x13, 0xb4, 0x2c, 0x96, 0x4c, 0x43, 0x7e, 0x07, 0x07, 0xe0, 0xd3,
	0x0d, 0x89, 0xfd, 0xf4, 0x86, 0xd2, 0x75, 0x4f, 0x19, 0x2a, 0x67, 0x0a, 0x3a, 0xda, 0x90, 0xd8,
	0x93, 0x67, 0xf8, 0x16, 0x9c, 0x90, 0x40, 0xb0, 0x5b, 0xea, 0xdf, 0xd2, 0x58, 0xa4, 0x3d, 0x75,
	0xd8, 0x3c, 0x6b, 0x5f, 0x7c, 0x67, 0xd4, 0x14, 0x6d, 0x94, 0xa0, 0xf2, 0xf7, 0x9a, 0xc6, 0x62,
	0xc2, 0x52, 0x91, 0xb0, 0x55, 0x26, 0x6b, 0x45, 0xc7, 0x05, 0x92, 0xb4, 0xa7, 0xf0, 0x77, 0xf0,
	0x25, 0x61, 0x89, 0x1f, 0xb0, 0x24, 0xc8, 0xc2, 0xbc, 0x97, 0xb2, 0x89, 0x5e, 0x73, 0xa8, 0x9c,
	0xb5, 0x2f, 0xbe, 0x7f, 0x4c, 0x8a, 0xf1, 0x3e, 0xda, 0xcb, 0x83, 0xd1, 0x29, 0xa9, 0xb0, 0xea,
	0x02, 0x74, 0x2a, 0x0a, 0x82, 0xa7, 0x40, 0x5b, 0x3a, 0x97, 0x8e, 0xfb, 0xd6, 0xf1, 0xad, 0x19,
	0xf2, 0xaf, 0x6d, 0x07, 0x6b, 0x0d, 0x08, 0x41, 0x7b, 0x3a, 0x77, 0x5d, 0xb4, 0xb7, 0x29, 0xf0,
	0x05, 0x78, 0x3e, 0x73, 0x3c, 0x8c, 0x96, 0x0b, 0xdb, 0xc1, 0xfe, 0x95, 0xe5, 0xd8, 0xf3, 0xfd,
	0xb5, 0x2a, 0x81, 0x26, 0xf6, 0x14, 0xb9, 0x1e, 0xde, 0x5b, 0x9b, 0xfa, 0x35, 0x38, 0xad, 0xaa,
	0x11, 0x0e, 0x40, 0xf7, 0x30, 0xed, 0x78, 0x86, 0xc6, 0xcb, 0xb9, 0x85, 0x67, 0xae, 0xa3, 0x35,
	0xa0, 0x06, 0x4e, 0xa6, 0xc8, 0xf6, 0xde, 0xf8, 0x33, 0x07, 0x5b, 0x97, 0xb6, 0xa6, 0xc0, 0xcf,
	0xc1, 0x31, 0xb2, 0x77, 0x4e, 0xb6, 0xa6, 0xea, 0xff, 0x28, 0xa0, 0x3b, 0x26, 0x2b, 0x16, 0x63,
	0x1a, 0xdd, 0xc8, 0x45, 0xc8, 0x12, 0xea, 0x51, 0x21, 0x58, 0xbc, 0x85, 0x23, 0xf0, 0x55, 0x44,
	0x49, 0x9a, 0x25, 0x74, 0xed, 0x8b, 0xfd, 0xb5, 0x9f, 0x16, 0xf7, 0x05, 0xbf, 0x6f, 0x1a, 0xa8,
	0xbf, 0xf3, 0xaa, 0xc6, 0x48, 0xa8, 0xac, 0xf8, 0x96, 0x56, 0x62, 0xa8, 0x3b, 0x8c, 0x9d, 0x57,
	0x05, 0xc6, 0x6b, 0x70, 0x94, 0xc5, 0x4c, 0xf8, 0x82, 0x6c, 0x73, 0x42, 0x8f, 0x2f, 0x86, 0xb5,
	0x84, 0xca, 0xb5, 0xc4, 0x64, 0x8b, 0x5a, 0x32, 0x02, 0x93, 0xed, 0xe8, 0x0b, 0xd0, 0x39, 0xcc,
	0x5b, 0xea, 0x45, 0xff, 0xfb, 0x09, 0x68, 0x8f, 0x0b, 0x0d, 0x8d, 0x0b, 0x13, 0x5c, 0x81, 0xcf,
	0xd2, 0xf7, 0xa9, 0xa0, 0xd1, 0x6e, 0x79, 0x94, 0x7c, 0x79, 0x7e, 0xae, 0xcd, 0x75, 0xa8, 0x1d,
	0x83, 0x6f, 0x36, 0x6e, 0x49, 0x8a, 0xe1, 0xee, 0xff, 0xa3, 0x93, 0x02, 0xb3, 0xa4, 0xeb, 0x35,
	0x38, 0x92, 0x9b, 0xba, 0x09, 0xf9, 0x5d, 0xde, 0xfa, 0xff, 0xb5, 0x52, 0xee, 0x26, 0x6a, 0x91,
	0x52, 0x5c, 0x21, 0x78, 0x1e, 0x48, 0xaa, 0x2a, 0x07, 0x59, 0x0c, 0xe6, 0x9b, 0x5a, 0xb4, 0x1a,
	0x92, 0x51, 0x37, 0xa8, 0x61, 0x3f, 0x29, 0x45, 0xc5, 0xe3, 0x35, 0x93, 0x2d, 0xd2, 0x64, 0x37,
	0x97, 0x4f, 0xf2, 0xb9, 0xfc, 0x54, 0x9f, 0xea, 0xc1, 0x5c, 0x73, 0x6d, 0xed, 0x41, 0x0e, 0xb5,
	0xf5, 0xa1, 0x15, 0x32, 0xd0, 0x29, 0x29, 0x28, 0x79, 0xf2, 0x23, 0xbe, 0xa6, 0xbd, 0x27, 0x79,
	0xc2, 0x1f, 0x1f, 0x9b, 0xd0, 0xcb, 0x21, 0xca, 0xd3, 0x82, 0xaf, 0x29, 0x7a, 0x96, 0x7e, 0x68,
	0xd2, 0x2f, 0x0b, 0x41, 0x7d, 0x54, 0x82, 0x0e, 0x5e, 0x3e, 0x10, 0x94, 0xeb, 0x4c, 0x66, 0x52,
	0x4e, 0x36, 0xf2, 0x3d, 0x6c, 0xe1, 0xa5, 0xa7, 0x35, 0x60, 0x0b, 0x34, 0xdd, 0xe9, 0x54, 0x53,
	0xe0, 0x53, 0xa0, 0xba, 0x8e, 0xa6, 0xea, 0x77, 0xe0, 0xd9, 0x47, 0x49, 0xe1, 0xd7, 0x60, 0xb0,
	0x43, 0xf2, 0x7e, 0xf1, 0xb0, 0xbd, 0x90, 0x60, 0x18, 0xb9, 0x73, 0x7f, 0xe1, 0x4e, 0x6c, 0xad,
	0x01, 0x5f, 0x82, 0xfe, 0xc2, 0x72, 0x96, 0xd6, 0xbc, 0xf2, 0x5e, 0x81, 0xaf, 0xc0, 0x0b, 0x6b,
	0x89, 0xdd, 0x85, 0x85, 0x67, 0xe3, 0x4a, 0x17, 0x75, 0xf4, 0x87, 0x02, 0x06, 0x01, 0x8f, 0xea,
	0x26, 0x33, 0xea, 0x3c, 0x1c, 0xcd, 0x95, 0x7c, 0xb0, 0xaf, 0x94, 0x5f, 0x9d, 0x2d, 0x13, 0xbf,
	0x65, 0x2b, 0x23, 0xe0, 0x91, 0x79, 0x1f, 0x7a, 0x4e, 0x98, 0x7c, 0xe3, 0x69, 0x12, 0x93, 0xf0,
	0x3c, 0x7f, 0xda, 0x53, 0x33, 0x4d, 0x02, 0x33, 0x22, 0x2c, 0x36, 0xf3, 0xb3, 0x59, 0xf3, 0x31,
	0xf8, 0x57, 0x51, 0xfe, 0x54, 0x9b, 0xd6, 0x12, 0xaf, 0x9e, 0xe6, 0x7e, 0xdf, 0xfe, 0x17, 0x00,
	0x00, 0xff, 0xff, 0xc6, 0xd2, 0x49, 0x64, 0xb8, 0x06, 0x00, 0x00,
}
