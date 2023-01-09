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
// source: autonomic/ext/telemetry/well_known_indicators.proto

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

type WellKnownIndicator int32

const (
	WellKnownIndicator_UNKNOWN_INDICATOR                     WellKnownIndicator = 0
	WellKnownIndicator_ADAPTIVE_CRUISE_CONTROL               WellKnownIndicator = 1
	WellKnownIndicator_AIR_FILTER_MINDER                     WellKnownIndicator = 2
	WellKnownIndicator_AIR_SUSPENSION_RIDE_CONTROL_FAULT     WellKnownIndicator = 3
	WellKnownIndicator_ALL_WHEEL_DRIVE_DISABLED              WellKnownIndicator = 4
	WellKnownIndicator_ANTI_THEFT                            WellKnownIndicator = 5
	WellKnownIndicator_ANTILOCK_BRAKE                        WellKnownIndicator = 6
	WellKnownIndicator_BLIND_SPOT_DETECTION                  WellKnownIndicator = 7
	WellKnownIndicator_BRAKE_WARNING                         WellKnownIndicator = 8
	WellKnownIndicator_CHARGE_SYSTEM_FAULT                   WellKnownIndicator = 9
	WellKnownIndicator_CHECK_FUEL_CAP                        WellKnownIndicator = 10
	WellKnownIndicator_CHECK_FUEL_FILL_INLET                 WellKnownIndicator = 11
	WellKnownIndicator_DIESEL_ENGINE_IDLE_SHUTDOWN           WellKnownIndicator = 12
	WellKnownIndicator_DIESEL_ENGINE_WARNING                 WellKnownIndicator = 13
	WellKnownIndicator_DIESEL_EXHAUST_FLUID_LOW              WellKnownIndicator = 14
	WellKnownIndicator_DIESEL_EXHAUST_FLUID_SYSTEM_FAULT     WellKnownIndicator = 15
	WellKnownIndicator_DIESEL_EXHAUST_OVER_TEMP              WellKnownIndicator = 16
	WellKnownIndicator_DIESEL_PARTICULATE_FILTER             WellKnownIndicator = 17
	WellKnownIndicator_DIESEL_PRE_HEAT                       WellKnownIndicator = 18
	WellKnownIndicator_ELECTRIC_TRAILER_BRAKE_CONNECTION     WellKnownIndicator = 19
	WellKnownIndicator_ENGINE_COOLANT_OVER_TEMP              WellKnownIndicator = 20
	WellKnownIndicator_FASTEN_SEAT_BELT_WARNING              WellKnownIndicator = 21
	WellKnownIndicator_FORWARD_COLLISION_WARNING             WellKnownIndicator = 22
	WellKnownIndicator_FUEL_DOOR_OPEN                        WellKnownIndicator = 23
	WellKnownIndicator_HEV_HAZARD                            WellKnownIndicator = 24
	WellKnownIndicator_HILL_DESCENT_CONTROL_FAULT            WellKnownIndicator = 25
	WellKnownIndicator_HILL_START_ASSIST_WARNING             WellKnownIndicator = 26
	WellKnownIndicator_LANE_KEEPING_AID                      WellKnownIndicator = 27
	WellKnownIndicator_LIGHTING_SYSTEM_FAILURE               WellKnownIndicator = 28
	WellKnownIndicator_LOW_ENGINE_OIL_PRESSURE               WellKnownIndicator = 29
	WellKnownIndicator_LOW_FUEL                              WellKnownIndicator = 30
	WellKnownIndicator_LOW_WASHER_FLUID                      WellKnownIndicator = 31
	WellKnownIndicator_MALFUNCTION_INDICATOR                 WellKnownIndicator = 32
	WellKnownIndicator_PARK_AID_MALFUNCTION                  WellKnownIndicator = 33
	WellKnownIndicator_PASSIVE_ENTRY_PASSIVE_START           WellKnownIndicator = 34
	WellKnownIndicator_POWERTRAIN_MALFUNCTION                WellKnownIndicator = 35
	WellKnownIndicator_RESTRAINTS_INDICATOR_WARNING          WellKnownIndicator = 36
	WellKnownIndicator_SERVICE_STEERING                      WellKnownIndicator = 37
	WellKnownIndicator_START_STOP_ENGINE_WARNING             WellKnownIndicator = 38
	WellKnownIndicator_TRACTION_CONTROL_DISABLED             WellKnownIndicator = 39
	WellKnownIndicator_TRACTION_CONTROL_EVENT                WellKnownIndicator = 40
	WellKnownIndicator_TIRE_PRESSURE_MONITOR_SYSTEM_WARNING  WellKnownIndicator = 41
	WellKnownIndicator_UREA_WARNING                          WellKnownIndicator = 42
	WellKnownIndicator_WATER_IN_FUEL                         WellKnownIndicator = 43
	WellKnownIndicator_DIESEL_FILTER_REGENERATION            WellKnownIndicator = 44
	WellKnownIndicator_DIESEL_EXHAUST_FLUID_QUALITY          WellKnownIndicator = 45
	WellKnownIndicator_EV_BATT_HIGH_TEMP_WARNING             WellKnownIndicator = 46 // Deprecated: Do not use.
	WellKnownIndicator_EV_BATT_INSULATION_RESIST_WARNING     WellKnownIndicator = 47 // Deprecated: Do not use.
	WellKnownIndicator_EV_BATT_POOR_CELL_WARNING             WellKnownIndicator = 48 // Deprecated: Do not use.
	WellKnownIndicator_EV_BATT_CHARGE_ENERGY_STORAGE_WARNING WellKnownIndicator = 49 // Deprecated: Do not use.
	WellKnownIndicator_DC_TEMP_WARNING                       WellKnownIndicator = 50
	WellKnownIndicator_EV_BATT_HIGH_LEVEL_WARNING            WellKnownIndicator = 51 // Deprecated: Do not use.
	WellKnownIndicator_HV_INTERLOCKING_STATUS_WARNING        WellKnownIndicator = 52 // Deprecated: Do not use.
	WellKnownIndicator_EV_BATT_LOW_LEVEL_WARNING             WellKnownIndicator = 53 // Deprecated: Do not use.
	WellKnownIndicator_EV_BATT_CELL_MAX_VOLT_WARNING         WellKnownIndicator = 54 // Deprecated: Do not use.
	WellKnownIndicator_EV_BATT_CELL_MIN_VOLT_WARNING         WellKnownIndicator = 55 // Deprecated: Do not use.
	WellKnownIndicator_EV_BATT_JUMP_LEVEL_WARNING            WellKnownIndicator = 56 // Deprecated: Do not use.
	WellKnownIndicator_EV_BATT_TEMP_DIFF_WARNING             WellKnownIndicator = 57 // Deprecated: Do not use.
	WellKnownIndicator_EV_BATT_OVER_CHARGE_WARNING           WellKnownIndicator = 58 // Deprecated: Do not use.
	WellKnownIndicator_DC_WARNING_STATUS                     WellKnownIndicator = 59
	WellKnownIndicator_EV_BATT_MAX_VOLT_VEH_ENERGY_WARNING   WellKnownIndicator = 60 // Deprecated: Do not use.
	WellKnownIndicator_EV_BATT_MIN_VOLT_VEH_ENERGY_WARNING   WellKnownIndicator = 61 // Deprecated: Do not use.
	WellKnownIndicator_MOTOR_CONTROLLER_TEMP_WARNING         WellKnownIndicator = 62
	WellKnownIndicator_TRACTION_MOTOR_TEMP_WARNING           WellKnownIndicator = 63
	WellKnownIndicator_CHECK_FUEL_FILTER                     WellKnownIndicator = 64
	WellKnownIndicator_ELECTRONIC_STABILITY_CONTROL          WellKnownIndicator = 65 // Deprecated: Do not use.
	WellKnownIndicator_LOW_SPEED_COLLISION_MITIGATION        WellKnownIndicator = 66 // Deprecated: Do not use.
	WellKnownIndicator_FORWARD_COLLISION_WARNING_SYSTEM      WellKnownIndicator = 67 // Deprecated: Do not use.
	WellKnownIndicator_LOW_SPEED_COLLISION_MITIGATION_SYSTEM WellKnownIndicator = 68 // Deprecated: Do not use.
	WellKnownIndicator_REAR_CROSS_WARNING                    WellKnownIndicator = 69 // Deprecated: Do not use.
)

var WellKnownIndicator_name = map[int32]string{
	0:  "UNKNOWN_INDICATOR",
	1:  "ADAPTIVE_CRUISE_CONTROL",
	2:  "AIR_FILTER_MINDER",
	3:  "AIR_SUSPENSION_RIDE_CONTROL_FAULT",
	4:  "ALL_WHEEL_DRIVE_DISABLED",
	5:  "ANTI_THEFT",
	6:  "ANTILOCK_BRAKE",
	7:  "BLIND_SPOT_DETECTION",
	8:  "BRAKE_WARNING",
	9:  "CHARGE_SYSTEM_FAULT",
	10: "CHECK_FUEL_CAP",
	11: "CHECK_FUEL_FILL_INLET",
	12: "DIESEL_ENGINE_IDLE_SHUTDOWN",
	13: "DIESEL_ENGINE_WARNING",
	14: "DIESEL_EXHAUST_FLUID_LOW",
	15: "DIESEL_EXHAUST_FLUID_SYSTEM_FAULT",
	16: "DIESEL_EXHAUST_OVER_TEMP",
	17: "DIESEL_PARTICULATE_FILTER",
	18: "DIESEL_PRE_HEAT",
	19: "ELECTRIC_TRAILER_BRAKE_CONNECTION",
	20: "ENGINE_COOLANT_OVER_TEMP",
	21: "FASTEN_SEAT_BELT_WARNING",
	22: "FORWARD_COLLISION_WARNING",
	23: "FUEL_DOOR_OPEN",
	24: "HEV_HAZARD",
	25: "HILL_DESCENT_CONTROL_FAULT",
	26: "HILL_START_ASSIST_WARNING",
	27: "LANE_KEEPING_AID",
	28: "LIGHTING_SYSTEM_FAILURE",
	29: "LOW_ENGINE_OIL_PRESSURE",
	30: "LOW_FUEL",
	31: "LOW_WASHER_FLUID",
	32: "MALFUNCTION_INDICATOR",
	33: "PARK_AID_MALFUNCTION",
	34: "PASSIVE_ENTRY_PASSIVE_START",
	35: "POWERTRAIN_MALFUNCTION",
	36: "RESTRAINTS_INDICATOR_WARNING",
	37: "SERVICE_STEERING",
	38: "START_STOP_ENGINE_WARNING",
	39: "TRACTION_CONTROL_DISABLED",
	40: "TRACTION_CONTROL_EVENT",
	41: "TIRE_PRESSURE_MONITOR_SYSTEM_WARNING",
	42: "UREA_WARNING",
	43: "WATER_IN_FUEL",
	44: "DIESEL_FILTER_REGENERATION",
	45: "DIESEL_EXHAUST_FLUID_QUALITY",
	46: "EV_BATT_HIGH_TEMP_WARNING",
	47: "EV_BATT_INSULATION_RESIST_WARNING",
	48: "EV_BATT_POOR_CELL_WARNING",
	49: "EV_BATT_CHARGE_ENERGY_STORAGE_WARNING",
	50: "DC_TEMP_WARNING",
	51: "EV_BATT_HIGH_LEVEL_WARNING",
	52: "HV_INTERLOCKING_STATUS_WARNING",
	53: "EV_BATT_LOW_LEVEL_WARNING",
	54: "EV_BATT_CELL_MAX_VOLT_WARNING",
	55: "EV_BATT_CELL_MIN_VOLT_WARNING",
	56: "EV_BATT_JUMP_LEVEL_WARNING",
	57: "EV_BATT_TEMP_DIFF_WARNING",
	58: "EV_BATT_OVER_CHARGE_WARNING",
	59: "DC_WARNING_STATUS",
	60: "EV_BATT_MAX_VOLT_VEH_ENERGY_WARNING",
	61: "EV_BATT_MIN_VOLT_VEH_ENERGY_WARNING",
	62: "MOTOR_CONTROLLER_TEMP_WARNING",
	63: "TRACTION_MOTOR_TEMP_WARNING",
	64: "CHECK_FUEL_FILTER",
	65: "ELECTRONIC_STABILITY_CONTROL",
	66: "LOW_SPEED_COLLISION_MITIGATION",
	67: "FORWARD_COLLISION_WARNING_SYSTEM",
	68: "LOW_SPEED_COLLISION_MITIGATION_SYSTEM",
	69: "REAR_CROSS_WARNING",
}

var WellKnownIndicator_value = map[string]int32{
	"UNKNOWN_INDICATOR":                     0,
	"ADAPTIVE_CRUISE_CONTROL":               1,
	"AIR_FILTER_MINDER":                     2,
	"AIR_SUSPENSION_RIDE_CONTROL_FAULT":     3,
	"ALL_WHEEL_DRIVE_DISABLED":              4,
	"ANTI_THEFT":                            5,
	"ANTILOCK_BRAKE":                        6,
	"BLIND_SPOT_DETECTION":                  7,
	"BRAKE_WARNING":                         8,
	"CHARGE_SYSTEM_FAULT":                   9,
	"CHECK_FUEL_CAP":                        10,
	"CHECK_FUEL_FILL_INLET":                 11,
	"DIESEL_ENGINE_IDLE_SHUTDOWN":           12,
	"DIESEL_ENGINE_WARNING":                 13,
	"DIESEL_EXHAUST_FLUID_LOW":              14,
	"DIESEL_EXHAUST_FLUID_SYSTEM_FAULT":     15,
	"DIESEL_EXHAUST_OVER_TEMP":              16,
	"DIESEL_PARTICULATE_FILTER":             17,
	"DIESEL_PRE_HEAT":                       18,
	"ELECTRIC_TRAILER_BRAKE_CONNECTION":     19,
	"ENGINE_COOLANT_OVER_TEMP":              20,
	"FASTEN_SEAT_BELT_WARNING":              21,
	"FORWARD_COLLISION_WARNING":             22,
	"FUEL_DOOR_OPEN":                        23,
	"HEV_HAZARD":                            24,
	"HILL_DESCENT_CONTROL_FAULT":            25,
	"HILL_START_ASSIST_WARNING":             26,
	"LANE_KEEPING_AID":                      27,
	"LIGHTING_SYSTEM_FAILURE":               28,
	"LOW_ENGINE_OIL_PRESSURE":               29,
	"LOW_FUEL":                              30,
	"LOW_WASHER_FLUID":                      31,
	"MALFUNCTION_INDICATOR":                 32,
	"PARK_AID_MALFUNCTION":                  33,
	"PASSIVE_ENTRY_PASSIVE_START":           34,
	"POWERTRAIN_MALFUNCTION":                35,
	"RESTRAINTS_INDICATOR_WARNING":          36,
	"SERVICE_STEERING":                      37,
	"START_STOP_ENGINE_WARNING":             38,
	"TRACTION_CONTROL_DISABLED":             39,
	"TRACTION_CONTROL_EVENT":                40,
	"TIRE_PRESSURE_MONITOR_SYSTEM_WARNING":  41,
	"UREA_WARNING":                          42,
	"WATER_IN_FUEL":                         43,
	"DIESEL_FILTER_REGENERATION":            44,
	"DIESEL_EXHAUST_FLUID_QUALITY":          45,
	"EV_BATT_HIGH_TEMP_WARNING":             46,
	"EV_BATT_INSULATION_RESIST_WARNING":     47,
	"EV_BATT_POOR_CELL_WARNING":             48,
	"EV_BATT_CHARGE_ENERGY_STORAGE_WARNING": 49,
	"DC_TEMP_WARNING":                       50,
	"EV_BATT_HIGH_LEVEL_WARNING":            51,
	"HV_INTERLOCKING_STATUS_WARNING":        52,
	"EV_BATT_LOW_LEVEL_WARNING":             53,
	"EV_BATT_CELL_MAX_VOLT_WARNING":         54,
	"EV_BATT_CELL_MIN_VOLT_WARNING":         55,
	"EV_BATT_JUMP_LEVEL_WARNING":            56,
	"EV_BATT_TEMP_DIFF_WARNING":             57,
	"EV_BATT_OVER_CHARGE_WARNING":           58,
	"DC_WARNING_STATUS":                     59,
	"EV_BATT_MAX_VOLT_VEH_ENERGY_WARNING":   60,
	"EV_BATT_MIN_VOLT_VEH_ENERGY_WARNING":   61,
	"MOTOR_CONTROLLER_TEMP_WARNING":         62,
	"TRACTION_MOTOR_TEMP_WARNING":           63,
	"CHECK_FUEL_FILTER":                     64,
	"ELECTRONIC_STABILITY_CONTROL":          65,
	"LOW_SPEED_COLLISION_MITIGATION":        66,
	"FORWARD_COLLISION_WARNING_SYSTEM":      67,
	"LOW_SPEED_COLLISION_MITIGATION_SYSTEM": 68,
	"REAR_CROSS_WARNING":                    69,
}

func (x WellKnownIndicator) String() string {
	return proto.EnumName(WellKnownIndicator_name, int32(x))
}

func (WellKnownIndicator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f727f5c0e8ffd501, []int{0}
}

func init() {
	proto.RegisterEnum("autonomic.ext.telemetry.indicator.WellKnownIndicator", WellKnownIndicator_name, WellKnownIndicator_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/well_known_indicators.proto", fileDescriptor_f727f5c0e8ffd501)
}

var fileDescriptor_f727f5c0e8ffd501 = []byte{
	// 1166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0x5d, 0x77, 0x13, 0x37,
	0x13, 0x7e, 0x0d, 0x6f, 0x29, 0x55, 0xf9, 0x50, 0x04, 0x21, 0xc4, 0x49, 0x80, 0x40, 0x42, 0x03,
	0x34, 0x76, 0xdb, 0xf4, 0xfb, 0x5b, 0xde, 0x1d, 0x7b, 0x55, 0xcb, 0xd2, 0x56, 0xd2, 0x7a, 0x09,
	0x37, 0x3a, 0x21, 0xf5, 0x69, 0x7d, 0xea, 0xd8, 0x3d, 0x89, 0x39, 0xd0, 0xdf, 0xd2, 0xbb, 0xfe,
	0xbd, 0xfe, 0x81, 0x5e, 0xf6, 0x8c, 0xbc, 0xbb, 0xb1, 0x43, 0x68, 0x2f, 0x57, 0xf3, 0x68, 0xe6,
	0x99, 0x99, 0x67, 0x66, 0x45, 0xf6, 0x0e, 0x5e, 0x4c, 0x27, 0xe3, 0xc9, 0xd1, 0xf0, 0xb0, 0x39,
	0x78, 0x35, 0x6d, 0x4e, 0x07, 0xa3, 0xc1, 0xd1, 0x60, 0x7a, 0xfc, 0x7b, 0xf3, 0xe5, 0x60, 0x34,
	0xf2, 0xbf, 0x8e, 0x27, 0x2f, 0xc7, 0x7e, 0x38, 0xfe, 0x69, 0x78, 0x78, 0x30, 0x9d, 0x1c, 0x9f,
	0x34, 0x7e, 0x3b, 0x9e, 0x4c, 0x27, 0x6c, 0xb3, 0xba, 0xd4, 0x18, 0xbc, 0x9a, 0x36, 0xaa, 0x4b,
	0x8d, 0x0a, 0xf9, 0xf8, 0x2f, 0x4a, 0x58, 0x3e, 0x18, 0x8d, 0xba, 0xe8, 0x41, 0x94, 0xc7, 0x6c,
	0x99, 0x2c, 0x65, 0xaa, 0xab, 0x74, 0xae, 0xbc, 0x50, 0xb1, 0x88, 0xb8, 0xd3, 0x86, 0xfe, 0x8f,
	0xad, 0x91, 0x15, 0x1e, 0xf3, 0xd4, 0x89, 0x3e, 0xf8, 0xc8, 0x64, 0xc2, 0x82, 0x8f, 0xb4, 0x72,
	0x46, 0x4b, 0x5a, 0xc3, 0x3b, 0x5c, 0x18, 0xdf, 0x16, 0xd2, 0x81, 0xf1, 0x3d, 0xa1, 0x62, 0x30,
	0xf4, 0x02, 0xdb, 0x26, 0x9b, 0x78, 0x6c, 0x33, 0x9b, 0x82, 0xb2, 0x42, 0x2b, 0x6f, 0x44, 0x5c,
	0xdd, 0xf3, 0x6d, 0x9e, 0x49, 0x47, 0x2f, 0xb2, 0x75, 0x72, 0x9b, 0x4b, 0xe9, 0xf3, 0x04, 0x40,
	0xfa, 0xd8, 0x60, 0x84, 0x58, 0x58, 0xde, 0x92, 0x10, 0xd3, 0xff, 0xb3, 0x6b, 0x84, 0x70, 0xe5,
	0x84, 0x77, 0x09, 0xb4, 0x1d, 0x7d, 0x8b, 0x31, 0x72, 0x0d, 0xbf, 0xa5, 0x8e, 0xba, 0xbe, 0x65,
	0x78, 0x17, 0xe8, 0x25, 0x76, 0x9b, 0xdc, 0x6c, 0x49, 0xa1, 0x62, 0x6f, 0x53, 0xed, 0x7c, 0x0c,
	0x0e, 0x22, 0x27, 0xb4, 0xa2, 0x6f, 0xb3, 0x25, 0x72, 0x35, 0x80, 0x7c, 0xce, 0x8d, 0x12, 0xaa,
	0x43, 0x2f, 0xb3, 0x15, 0x72, 0x23, 0x4a, 0xb8, 0xe9, 0x80, 0xb7, 0xfb, 0xd6, 0x41, 0xaf, 0xe0,
	0xf1, 0x0e, 0x7a, 0x8e, 0x12, 0x88, 0xba, 0xbe, 0x9d, 0x81, 0xf4, 0x11, 0x4f, 0x29, 0x61, 0xab,
	0x64, 0x79, 0xee, 0xac, 0x2d, 0xa4, 0xf4, 0x42, 0x49, 0x70, 0xf4, 0x5d, 0x76, 0x97, 0xac, 0xc5,
	0x02, 0x2c, 0x48, 0x0f, 0xaa, 0x23, 0x14, 0x78, 0x11, 0x4b, 0xf0, 0x36, 0xc9, 0x5c, 0xac, 0x73,
	0x45, 0xaf, 0xe0, 0xdd, 0x45, 0x40, 0xc9, 0xe1, 0x2a, 0xa6, 0x5c, 0x9a, 0x9e, 0x26, 0x3c, 0xb3,
	0xce, 0xb7, 0x65, 0x26, 0x62, 0x2f, 0x75, 0x4e, 0xaf, 0x61, 0xdd, 0xce, 0xb5, 0x2e, 0xf0, 0xbd,
	0x7e, 0x8e, 0x13, 0xdd, 0x07, 0xe3, 0x1d, 0xf4, 0x52, 0x4a, 0xd9, 0x06, 0x59, 0x2d, 0xac, 0x29,
	0x37, 0x4e, 0x44, 0x99, 0xe4, 0x0e, 0x8a, 0x16, 0xd1, 0x25, 0x76, 0x83, 0x5c, 0x2f, 0xcd, 0x06,
	0x7c, 0x02, 0xdc, 0x51, 0x86, 0x81, 0x41, 0x42, 0xe4, 0x8c, 0x88, 0xbc, 0x33, 0x5c, 0x48, 0x30,
	0xb3, 0x1a, 0x63, 0xcf, 0x54, 0x51, 0xd4, 0x1b, 0x18, 0xb8, 0xc8, 0x28, 0xd2, 0x5a, 0x72, 0x35,
	0x1f, 0xf8, 0x26, 0x5a, 0xdb, 0xdc, 0x3a, 0x50, 0xde, 0x02, 0x77, 0xbe, 0x05, 0xd2, 0x55, 0x99,
	0x2f, 0x23, 0xad, 0xb6, 0x36, 0x39, 0x37, 0xb1, 0x8f, 0xb4, 0x94, 0x22, 0xc8, 0xa2, 0x34, 0xdf,
	0xc2, 0x1e, 0x84, 0x4a, 0xc7, 0x5a, 0x1b, 0xaf, 0x53, 0x50, 0x74, 0x05, 0x15, 0x90, 0x40, 0xdf,
	0x27, 0xfc, 0x19, 0x37, 0x31, 0xbd, 0xcd, 0xee, 0x90, 0x7a, 0x82, 0x8d, 0x88, 0xc1, 0x46, 0xa0,
	0xdc, 0x19, 0x3d, 0xad, 0x62, 0x88, 0x60, 0xb7, 0x8e, 0x1b, 0xe7, 0xb9, 0xb5, 0xc2, 0x9e, 0x32,
	0xa8, 0xb3, 0x9b, 0x84, 0x4a, 0xae, 0xc0, 0x77, 0x01, 0x52, 0xa1, 0x3a, 0x9e, 0x8b, 0x98, 0xae,
	0xa1, 0xbe, 0xa5, 0xe8, 0x24, 0x0e, 0x4f, 0xaa, 0x3a, 0x0b, 0x99, 0x19, 0xa0, 0xeb, 0xc1, 0xa8,
	0xf3, 0xb2, 0x8d, 0x5a, 0x84, 0xa2, 0x59, 0x8b, 0xc6, 0x0d, 0x76, 0x85, 0x5c, 0x46, 0x23, 0xd2,
	0xa6, 0x77, 0x82, 0x77, 0x9d, 0xfb, 0x9c, 0xdb, 0x04, 0xcc, 0xac, 0x6f, 0xf4, 0x2e, 0x4a, 0xa1,
	0xc7, 0x65, 0x3b, 0x53, 0xa1, 0x84, 0x73, 0x83, 0x75, 0x0f, 0xb5, 0x9b, 0x72, 0xd3, 0x45, 0x1a,
	0x7e, 0x0e, 0x43, 0x37, 0x51, 0x60, 0x29, 0xb2, 0xef, 0x83, 0x07, 0xe5, 0xcc, 0xbe, 0x2f, 0xbf,
	0x42, 0x62, 0xf4, 0x3e, 0xab, 0x93, 0x5b, 0xa9, 0xce, 0xc1, 0x60, 0xab, 0xd4, 0xc2, 0xe5, 0x07,
	0xec, 0x1e, 0x59, 0x37, 0x60, 0x83, 0xc5, 0xd9, 0xd3, 0x80, 0x55, 0x1d, 0xb6, 0x90, 0xa9, 0x05,
	0xd3, 0x17, 0x11, 0x3a, 0x04, 0x30, 0x78, 0xba, 0x8d, 0xc5, 0x9b, 0xd5, 0xcd, 0x3a, 0x9d, 0x9e,
	0x15, 0xee, 0x43, 0x34, 0x3b, 0xc3, 0x67, 0x59, 0x94, 0x75, 0xaf, 0x86, 0xf5, 0x3d, 0x64, 0xf4,
	0x9a, 0x19, 0xfa, 0xa0, 0x1c, 0xdd, 0x61, 0x3b, 0x64, 0xcb, 0x09, 0x03, 0x55, 0xe9, 0x7c, 0x4f,
	0x2b, 0x81, 0x94, 0x8a, 0x72, 0x97, 0x41, 0x1e, 0x31, 0x4a, 0xae, 0x64, 0x06, 0x78, 0x75, 0xf2,
	0x18, 0xc7, 0x38, 0xe7, 0xb8, 0x5b, 0x84, 0x9a, 0x15, 0xfa, 0x09, 0xaa, 0xa0, 0x10, 0x70, 0xb1,
	0x76, 0x0c, 0x74, 0x40, 0x81, 0xe1, 0xa1, 0x00, 0xef, 0x63, 0x01, 0xce, 0x1d, 0xa2, 0x1f, 0x33,
	0x2e, 0x85, 0xdb, 0xa7, 0xbb, 0x6c, 0x93, 0xac, 0x42, 0xdf, 0xb7, 0xb8, 0x73, 0x3e, 0x11, 0x9d,
	0x24, 0xe8, 0xb7, 0x8a, 0xd9, 0xa8, 0x5f, 0xb8, 0x5c, 0x63, 0x8f, 0xc8, 0x66, 0x09, 0x11, 0xca,
	0xe2, 0x08, 0x85, 0x2d, 0x06, 0x0b, 0x92, 0x6a, 0x06, 0xe8, 0x9c, 0xb7, 0x14, 0xc5, 0x1b, 0x01,
	0xee, 0xb4, 0x02, 0xf2, 0x41, 0x80, 0xec, 0x92, 0xed, 0x12, 0x52, 0x6c, 0x20, 0x24, 0xdc, 0xd9,
	0xc7, 0x5a, 0x1b, 0xde, 0x39, 0xad, 0xf3, 0x87, 0x01, 0x8e, 0x23, 0x1a, 0x2d, 0xb2, 0xfa, 0x88,
	0xdd, 0x27, 0xf5, 0x05, 0xd2, 0x12, 0xfa, 0x70, 0x1a, 0x67, 0x2f, 0x5c, 0x7c, 0x48, 0xee, 0x24,
	0x7d, 0x2f, 0x94, 0x03, 0x83, 0x6b, 0x32, 0x48, 0xda, 0x71, 0x97, 0xd9, 0x0a, 0xf7, 0xf1, 0x59,
	0xca, 0xa8, 0xd9, 0x45, 0x57, 0x9f, 0x04, 0xc8, 0x36, 0xd9, 0xa8, 0x28, 0x63, 0x42, 0x3d, 0xfe,
	0xd4, 0xf7, 0xf5, 0xdc, 0x44, 0x7f, 0x7a, 0x3e, 0x4c, 0xa8, 0x45, 0xd8, 0x67, 0x01, 0x36, 0x47,
	0xfe, 0x87, 0xac, 0x97, 0x9e, 0x89, 0xf8, 0xf9, 0x59, 0x52, 0x21, 0xf5, 0x58, 0xb4, 0xdb, 0x15,
	0xe4, 0x8b, 0x00, 0x79, 0x40, 0xd6, 0x4a, 0x48, 0x58, 0x3c, 0x45, 0x31, 0x4b, 0xd0, 0x97, 0x01,
	0xb4, 0x4c, 0x96, 0xe2, 0xa8, 0x3c, 0x2b, 0xf2, 0xa7, 0x5f, 0xb1, 0x27, 0xe4, 0x41, 0x79, 0xb7,
	0xca, 0xa5, 0x0f, 0x49, 0xd9, 0x89, 0xd2, 0xc7, 0xd7, 0xc1, 0xc7, 0x3c, 0xb8, 0xcc, 0xe8, 0x1c,
	0xf0, 0x37, 0x05, 0xf1, 0x8d, 0x9e, 0x46, 0x3d, 0x17, 0xc2, 0x97, 0xc5, 0x4a, 0xac, 0x60, 0xdf,
	0xe2, 0x44, 0x57, 0xe3, 0x31, 0xc3, 0x2e, 0x00, 0xbe, 0x43, 0xd2, 0x8b, 0xbf, 0x1b, 0x5c, 0xd6,
	0xdf, 0xb3, 0x2d, 0xb2, 0x3e, 0xdb, 0xcb, 0x5a, 0x89, 0x08, 0x73, 0x69, 0x09, 0xd4, 0x70, 0xf5,
	0x07, 0xe6, 0x65, 0xdb, 0xb1, 0x8d, 0x36, 0x05, 0x98, 0x5f, 0xae, 0x3d, 0xe1, 0x44, 0x67, 0x36,
	0x15, 0xad, 0x80, 0xdb, 0x21, 0xf7, 0xde, 0xb8, 0x82, 0x8b, 0x61, 0xa4, 0x51, 0x29, 0xd8, 0x7f,
	0xf7, 0x58, 0xc2, 0xe3, 0x00, 0xaf, 0x13, 0x66, 0x80, 0x1b, 0x1f, 0x19, 0x6d, 0x4f, 0xb5, 0x06,
	0x68, 0x6b, 0xfd, 0x51, 0x23, 0xdb, 0x87, 0x93, 0xa3, 0xc6, 0x7f, 0xbe, 0x4b, 0x5a, 0x2b, 0xaf,
	0x3f, 0x4a, 0x52, 0x7c, 0xd3, 0xa4, 0xb5, 0x67, 0xea, 0xe7, 0xe1, 0xf4, 0x97, 0x17, 0xcf, 0x1b,
	0x87, 0x93, 0xa3, 0x66, 0xe5, 0x68, 0xf7, 0x60, 0x88, 0x0f, 0xa3, 0xc1, 0xf1, 0xf8, 0x60, 0xb4,
	0x1b, 0x5e, 0x3f, 0x27, 0xcd, 0x93, 0xe3, 0xc3, 0xe6, 0xd1, 0xc1, 0x70, 0xdc, 0x0c, 0xdf, 0xcd,
	0x37, 0xbc, 0xa0, 0xfe, 0xae, 0xd5, 0xfe, 0xbc, 0x70, 0x91, 0x67, 0xee, 0xf9, 0xa5, 0x80, 0xdb,
	0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x40, 0xe7, 0x6d, 0x6c, 0x09, 0x00, 0x00,
}