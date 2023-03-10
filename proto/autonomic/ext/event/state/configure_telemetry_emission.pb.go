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
// source: autonomic/ext/event/state/configure_telemetry_emission.proto

package state

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

type WellKnownConfigureTelemetryEmissionState int32

const (
	WellKnownConfigureTelemetryEmissionState_UNKNOWN_STATE WellKnownConfigureTelemetryEmissionState = 0
	// Request was received by cloud through API.
	// Emitted by: cloud.
	WellKnownConfigureTelemetryEmissionState_REQUESTED WellKnownConfigureTelemetryEmissionState = 1
	// Request was queued in the cloud, waiting for device to connect.
	// Emitted by: cloud.
	// Triggered by: REQUEST_QUEUED_FOR_DELIVERY.
	WellKnownConfigureTelemetryEmissionState_REQUEST_DELIVERY_QUEUED WellKnownConfigureTelemetryEmissionState = 2
	// Cloud attempted delivery of request to device.
	// Emitted by: cloud.
	// Triggered by: OEM_CLOUD_TRANSLATION_SUCCEEDED.
	WellKnownConfigureTelemetryEmissionState_REQUEST_DELIVERY_IN_PROGRESS WellKnownConfigureTelemetryEmissionState = 3
	// Delivery of request attempt was not acknowledged by device in a timely
	// fashion.
	// Emitted by: cloud.
	// Triggered by: REQUEST_DELIVERY_TIMEOUT_EXCEEDED.
	WellKnownConfigureTelemetryEmissionState_REQUEST_DELIVERY_TIMED_OUT WellKnownConfigureTelemetryEmissionState = 4
	// Delivery of request attempt failed due to device disconnect.
	// Emitted by: cloud.
	// Triggered by: REQUEST_DELIVERY_FAILED.
	WellKnownConfigureTelemetryEmissionState_REQUEST_DELIVERY_FAILURE WellKnownConfigureTelemetryEmissionState = 5
	// Request is acknowledged by device, queued, and the device will begin
	// processing the request when the device is ready.
	// Emitted by: device.
	// Triggered by: REQUEST_QUEUED_ON_DEVICE.
	WellKnownConfigureTelemetryEmissionState_REQUEST_QUEUED WellKnownConfigureTelemetryEmissionState = 6
	// Request is being decrypted and validated by device.
	// Emitted by: device.
	// Triggered by: REQUEST_VALIDATION_STARTED.
	WellKnownConfigureTelemetryEmissionState_REQUEST_VALIDATION_IN_PROGRESS WellKnownConfigureTelemetryEmissionState = 7
	// Request failed validation or decryption.
	// Emitted by: device.
	// Triggered by: REQUEST_VALIDATION_FAILED.
	WellKnownConfigureTelemetryEmissionState_REQUEST_VALIDATION_FAILURE WellKnownConfigureTelemetryEmissionState = 8
	// The device has successfully completed the actuation in question.
	// Triggered by: CONFIGURE_TELEMETRY_EMISSION_SUCCEEDED
	WellKnownConfigureTelemetryEmissionState_SUCCEEDED_ON_DEVICE WellKnownConfigureTelemetryEmissionState = 23
	// Telemetry Emission settings in progress
	// Emitted by: device.
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED.
	WellKnownConfigureTelemetryEmissionState_CONFIGURE_TELEMETRY_EMISSION_IN_PROGRESS WellKnownConfigureTelemetryEmissionState = 9
	// Oem cloud attempted delivery of request.
	// Emitted by: cloud.
	// Triggered by: DEVICE_CONNECTED_TO_CLOUD.
	WellKnownConfigureTelemetryEmissionState_OEM_CLOUD_DELIVERY_IN_PROGRESS WellKnownConfigureTelemetryEmissionState = 10
	// Cloud attempted translation of request.
	// Emitted by: cloud.
	// Triggered by: OEM_CLOUD_DELIVERY_SUCCEEDED.
	WellKnownConfigureTelemetryEmissionState_OEM_CLOUD_TRANSLATION_IN_PROGRESS WellKnownConfigureTelemetryEmissionState = 11
	// The cloud is in the process of validating the event emitted by the device indicating that the
	// command has succeeded.
	// Triggered by: SUCCEEDED_ON_DEVICE_EVENT_RECEIVED
	WellKnownConfigureTelemetryEmissionState_CLOUD_VALIDATION_IN_PROGRESS WellKnownConfigureTelemetryEmissionState = 24
	// Cloud failed to validate data provided by the device
	// Emitted by: cloud
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED
	WellKnownConfigureTelemetryEmissionState_CLOUD_VALIDATION_FAILURE WellKnownConfigureTelemetryEmissionState = 12
	// Telemetry Emission upload succeeded
	// This state is terminal.
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED.
	WellKnownConfigureTelemetryEmissionState_SUCCESS WellKnownConfigureTelemetryEmissionState = 17
	// Telemetry Emission failed. All intermediate failure states should
	// transition to this state.
	// Emitted by: device.
	// Triggered by: CONFIGURE_TELEMETRY_EMISSION_FAILED.
	WellKnownConfigureTelemetryEmissionState_FAILURE WellKnownConfigureTelemetryEmissionState = 18
	// Telemetry Emission command expired. This event can be emitted by the
	// device or the cloud. In either case, it will cause the command to be
	// removed from the set of in-progress commands on the cloud. However the
	// cloud does not send this event to the device.
	// This state is terminal.
	// Triggered by: EXPIRATION_EXCEEDED.
	WellKnownConfigureTelemetryEmissionState_EXPIRED WellKnownConfigureTelemetryEmissionState = 19
	// Command is being removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUESTED
	WellKnownConfigureTelemetryEmissionState_CANCELLATION_IN_PROGRESS WellKnownConfigureTelemetryEmissionState = 20
	// Telemetry Emission Command removed from set of in-progress commands
	// Emitted by: cloud
	// Triggered by: CANCELLATION_REQUEST_FULFILLED
	WellKnownConfigureTelemetryEmissionState_CANCELLED WellKnownConfigureTelemetryEmissionState = 21
)

var WellKnownConfigureTelemetryEmissionState_name = map[int32]string{
	0:  "UNKNOWN_STATE",
	1:  "REQUESTED",
	2:  "REQUEST_DELIVERY_QUEUED",
	3:  "REQUEST_DELIVERY_IN_PROGRESS",
	4:  "REQUEST_DELIVERY_TIMED_OUT",
	5:  "REQUEST_DELIVERY_FAILURE",
	6:  "REQUEST_QUEUED",
	7:  "REQUEST_VALIDATION_IN_PROGRESS",
	8:  "REQUEST_VALIDATION_FAILURE",
	23: "SUCCEEDED_ON_DEVICE",
	9:  "CONFIGURE_TELEMETRY_EMISSION_IN_PROGRESS",
	10: "OEM_CLOUD_DELIVERY_IN_PROGRESS",
	11: "OEM_CLOUD_TRANSLATION_IN_PROGRESS",
	24: "CLOUD_VALIDATION_IN_PROGRESS",
	12: "CLOUD_VALIDATION_FAILURE",
	17: "SUCCESS",
	18: "FAILURE",
	19: "EXPIRED",
	20: "CANCELLATION_IN_PROGRESS",
	21: "CANCELLED",
}

var WellKnownConfigureTelemetryEmissionState_value = map[string]int32{
	"UNKNOWN_STATE":                            0,
	"REQUESTED":                                1,
	"REQUEST_DELIVERY_QUEUED":                  2,
	"REQUEST_DELIVERY_IN_PROGRESS":             3,
	"REQUEST_DELIVERY_TIMED_OUT":               4,
	"REQUEST_DELIVERY_FAILURE":                 5,
	"REQUEST_QUEUED":                           6,
	"REQUEST_VALIDATION_IN_PROGRESS":           7,
	"REQUEST_VALIDATION_FAILURE":               8,
	"SUCCEEDED_ON_DEVICE":                      23,
	"CONFIGURE_TELEMETRY_EMISSION_IN_PROGRESS": 9,
	"OEM_CLOUD_DELIVERY_IN_PROGRESS":           10,
	"OEM_CLOUD_TRANSLATION_IN_PROGRESS":        11,
	"CLOUD_VALIDATION_IN_PROGRESS":             24,
	"CLOUD_VALIDATION_FAILURE":                 12,
	"SUCCESS":                                  17,
	"FAILURE":                                  18,
	"EXPIRED":                                  19,
	"CANCELLATION_IN_PROGRESS":                 20,
	"CANCELLED":                                21,
}

func (x WellKnownConfigureTelemetryEmissionState) String() string {
	return proto.EnumName(WellKnownConfigureTelemetryEmissionState_name, int32(x))
}

func (WellKnownConfigureTelemetryEmissionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_50fc17a07c0567ab, []int{0}
}

// Configure Telemetry Emission state transition triggers. See
// WellKnownConfigureTelemetryEmissionState enum for documentation on when these
// triggers occur.
type WellKnownConfigureTelemetryEmissionTrigger int32

const (
	WellKnownConfigureTelemetryEmissionTrigger_UNKNOWN_TRANSITION_TRIGGER                WellKnownConfigureTelemetryEmissionTrigger = 0
	WellKnownConfigureTelemetryEmissionTrigger_REQUEST_QUEUED_FOR_DELIVERY               WellKnownConfigureTelemetryEmissionTrigger = 1
	WellKnownConfigureTelemetryEmissionTrigger_DEVICE_CONNECTED_TO_CLOUD                 WellKnownConfigureTelemetryEmissionTrigger = 2
	WellKnownConfigureTelemetryEmissionTrigger_REQUEST_DELIVERY_FAILED                   WellKnownConfigureTelemetryEmissionTrigger = 3
	WellKnownConfigureTelemetryEmissionTrigger_REQUEST_DELIVERY_TIMEOUT_EXCEEDED         WellKnownConfigureTelemetryEmissionTrigger = 4
	WellKnownConfigureTelemetryEmissionTrigger_REQUEST_QUEUED_ON_DEVICE                  WellKnownConfigureTelemetryEmissionTrigger = 5
	WellKnownConfigureTelemetryEmissionTrigger_REQUEST_VALIDATION_STARTED                WellKnownConfigureTelemetryEmissionTrigger = 6
	WellKnownConfigureTelemetryEmissionTrigger_REQUEST_VALIDATION_FAILED                 WellKnownConfigureTelemetryEmissionTrigger = 7
	WellKnownConfigureTelemetryEmissionTrigger_REQUEST_VALIDATION_SUCCEEDED              WellKnownConfigureTelemetryEmissionTrigger = 23
	WellKnownConfigureTelemetryEmissionTrigger_CONFIGURE_TELEMETRY_EMISSION_STARTED      WellKnownConfigureTelemetryEmissionTrigger = 8
	WellKnownConfigureTelemetryEmissionTrigger_CONFIGURE_TELEMETRY_EMISSION_SUCCEEDED    WellKnownConfigureTelemetryEmissionTrigger = 9
	WellKnownConfigureTelemetryEmissionTrigger_CONFIGURE_TELEMETRY_EMISSION_FAILED       WellKnownConfigureTelemetryEmissionTrigger = 10
	WellKnownConfigureTelemetryEmissionTrigger_OEM_CLOUD_DELIVERY_SUCCEEDED              WellKnownConfigureTelemetryEmissionTrigger = 11
	WellKnownConfigureTelemetryEmissionTrigger_OEM_CLOUD_TRANSLATION_SUCCEEDED           WellKnownConfigureTelemetryEmissionTrigger = 12
	WellKnownConfigureTelemetryEmissionTrigger_SUCCEEDED_ON_DEVICE_EVENT_RECEIVED        WellKnownConfigureTelemetryEmissionTrigger = 24
	WellKnownConfigureTelemetryEmissionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED WellKnownConfigureTelemetryEmissionTrigger = 13
	WellKnownConfigureTelemetryEmissionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED    WellKnownConfigureTelemetryEmissionTrigger = 14
	WellKnownConfigureTelemetryEmissionTrigger_FAILED                                    WellKnownConfigureTelemetryEmissionTrigger = 15
	WellKnownConfigureTelemetryEmissionTrigger_EXPIRATION_EXCEEDED                       WellKnownConfigureTelemetryEmissionTrigger = 16
	WellKnownConfigureTelemetryEmissionTrigger_CANCELLATION_REQUESTED                    WellKnownConfigureTelemetryEmissionTrigger = 17
	WellKnownConfigureTelemetryEmissionTrigger_CANCELLATION_REQUEST_FULFILLED            WellKnownConfigureTelemetryEmissionTrigger = 18
)

var WellKnownConfigureTelemetryEmissionTrigger_name = map[int32]string{
	0:  "UNKNOWN_TRANSITION_TRIGGER",
	1:  "REQUEST_QUEUED_FOR_DELIVERY",
	2:  "DEVICE_CONNECTED_TO_CLOUD",
	3:  "REQUEST_DELIVERY_FAILED",
	4:  "REQUEST_DELIVERY_TIMEOUT_EXCEEDED",
	5:  "REQUEST_QUEUED_ON_DEVICE",
	6:  "REQUEST_VALIDATION_STARTED",
	7:  "REQUEST_VALIDATION_FAILED",
	23: "REQUEST_VALIDATION_SUCCEEDED",
	8:  "CONFIGURE_TELEMETRY_EMISSION_STARTED",
	9:  "CONFIGURE_TELEMETRY_EMISSION_SUCCEEDED",
	10: "CONFIGURE_TELEMETRY_EMISSION_FAILED",
	11: "OEM_CLOUD_DELIVERY_SUCCEEDED",
	12: "OEM_CLOUD_TRANSLATION_SUCCEEDED",
	24: "SUCCEEDED_ON_DEVICE_EVENT_RECEIVED",
	13: "CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED",
	14: "CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED",
	15: "FAILED",
	16: "EXPIRATION_EXCEEDED",
	17: "CANCELLATION_REQUESTED",
	18: "CANCELLATION_REQUEST_FULFILLED",
}

var WellKnownConfigureTelemetryEmissionTrigger_value = map[string]int32{
	"UNKNOWN_TRANSITION_TRIGGER":                0,
	"REQUEST_QUEUED_FOR_DELIVERY":               1,
	"DEVICE_CONNECTED_TO_CLOUD":                 2,
	"REQUEST_DELIVERY_FAILED":                   3,
	"REQUEST_DELIVERY_TIMEOUT_EXCEEDED":         4,
	"REQUEST_QUEUED_ON_DEVICE":                  5,
	"REQUEST_VALIDATION_STARTED":                6,
	"REQUEST_VALIDATION_FAILED":                 7,
	"REQUEST_VALIDATION_SUCCEEDED":              23,
	"CONFIGURE_TELEMETRY_EMISSION_STARTED":      8,
	"CONFIGURE_TELEMETRY_EMISSION_SUCCEEDED":    9,
	"CONFIGURE_TELEMETRY_EMISSION_FAILED":       10,
	"OEM_CLOUD_DELIVERY_SUCCEEDED":              11,
	"OEM_CLOUD_TRANSLATION_SUCCEEDED":           12,
	"SUCCEEDED_ON_DEVICE_EVENT_RECEIVED":        24,
	"CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED": 13,
	"CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED":    14,
	"FAILED":                                    15,
	"EXPIRATION_EXCEEDED":                       16,
	"CANCELLATION_REQUESTED":                    17,
	"CANCELLATION_REQUEST_FULFILLED":            18,
}

func (x WellKnownConfigureTelemetryEmissionTrigger) String() string {
	return proto.EnumName(WellKnownConfigureTelemetryEmissionTrigger_name, int32(x))
}

func (WellKnownConfigureTelemetryEmissionTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_50fc17a07c0567ab, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.event.state.configuretelemetryemission.WellKnownConfigureTelemetryEmissionState", WellKnownConfigureTelemetryEmissionState_name, WellKnownConfigureTelemetryEmissionState_value)
	proto.RegisterEnum("autonomic.ext.event.state.configuretelemetryemission.WellKnownConfigureTelemetryEmissionTrigger", WellKnownConfigureTelemetryEmissionTrigger_name, WellKnownConfigureTelemetryEmissionTrigger_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/event/state/configure_telemetry_emission.proto", fileDescriptor_50fc17a07c0567ab)
}

var fileDescriptor_50fc17a07c0567ab = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4b, 0x6f, 0xd3, 0x4a,
	0x14, 0x6e, 0xfa, 0xee, 0xe9, 0xe3, 0x4e, 0xa7, 0xf7, 0xde, 0xf4, 0xdd, 0x7b, 0x5b, 0x28, 0x25,
	0x22, 0xf1, 0x02, 0x16, 0x2c, 0xd8, 0xb8, 0x9e, 0x93, 0xc8, 0xaa, 0x63, 0xa7, 0xf6, 0x38, 0x6d,
	0xd9, 0x8c, 0xd2, 0xc8, 0x04, 0x4b, 0x89, 0x8d, 0x12, 0x17, 0xca, 0xdf, 0xe1, 0xc7, 0xc0, 0x5f,
	0x62, 0x83, 0x84, 0xc6, 0xb1, 0xf3, 0x20, 0x4e, 0xa9, 0xd8, 0x79, 0x7c, 0xbe, 0xef, 0x3c, 0xbf,
	0x39, 0x03, 0x6f, 0x1a, 0x77, 0x51, 0x18, 0x84, 0x1d, 0xbf, 0xa9, 0x78, 0xf7, 0x91, 0xe2, 0x7d,
	0xf4, 0x82, 0x48, 0xe9, 0x45, 0x8d, 0xc8, 0x53, 0x9a, 0x61, 0xf0, 0xce, 0x6f, 0xdd, 0x75, 0x3d,
	0x11, 0x79, 0x6d, 0xaf, 0xe3, 0x45, 0xdd, 0xcf, 0xc2, 0xeb, 0xf8, 0xbd, 0x9e, 0x1f, 0x06, 0xa5,
	0x0f, 0xdd, 0x30, 0x0a, 0xe9, 0xab, 0x01, 0xbb, 0xe4, 0xdd, 0x47, 0xa5, 0x98, 0x5d, 0x8a, 0xd9,
	0xa5, 0x01, 0x7b, 0x40, 0x4e, 0xb9, 0x85, 0x6f, 0xf3, 0x70, 0x76, 0xe5, 0xb5, 0xdb, 0x17, 0x41,
	0xf8, 0x29, 0xd0, 0x52, 0x1c, 0x4f, 0x71, 0x98, 0xe0, 0x1c, 0xe9, 0x87, 0x6e, 0xc2, 0xba, 0x6b,
	0x5e, 0x98, 0xd6, 0x95, 0x29, 0x1c, 0xae, 0x72, 0x24, 0x33, 0x74, 0x1d, 0x56, 0x6c, 0xbc, 0x74,
	0xd1, 0xe1, 0xc8, 0x48, 0x8e, 0xee, 0x41, 0x3e, 0x39, 0x0a, 0x86, 0x86, 0x5e, 0x47, 0xfb, 0x46,
	0x5c, 0xba, 0xe8, 0x22, 0x23, 0xb3, 0xf4, 0x3f, 0xd8, 0x9f, 0x30, 0xea, 0xa6, 0xa8, 0xd9, 0x56,
	0xc5, 0x46, 0xc7, 0x21, 0x73, 0xf4, 0x10, 0x76, 0x27, 0x10, 0x5c, 0xaf, 0x22, 0x13, 0x96, 0xcb,
	0xc9, 0x3c, 0xdd, 0x87, 0xed, 0x09, 0x7b, 0x59, 0xd5, 0x0d, 0xd7, 0x46, 0xb2, 0x40, 0x29, 0x6c,
	0xa4, 0xd6, 0x24, 0xe6, 0x22, 0x3d, 0x86, 0xc3, 0xf4, 0x5f, 0x5d, 0x35, 0x74, 0xa6, 0x72, 0xdd,
	0x32, 0xc7, 0xa2, 0x2e, 0x8d, 0x46, 0x1d, 0xc1, 0xa4, 0x7e, 0x97, 0x69, 0x1e, 0xb6, 0x1c, 0x57,
	0xd3, 0x10, 0x99, 0x4c, 0xc4, 0x14, 0x0c, 0xeb, 0xba, 0x86, 0x24, 0x4f, 0x5f, 0xc0, 0x99, 0x66,
	0x99, 0x65, 0xbd, 0xe2, 0xda, 0x28, 0x38, 0x1a, 0x58, 0x45, 0x6e, 0xdf, 0x08, 0xac, 0xea, 0x8e,
	0xf3, 0x6b, 0x98, 0x15, 0x99, 0x8a, 0x85, 0x55, 0xa1, 0x19, 0x96, 0xcb, 0xb2, 0x1b, 0x00, 0xf4,
	0x29, 0xfc, 0x3f, 0xc4, 0x70, 0x5b, 0x35, 0x1d, 0x63, 0x32, 0xe3, 0x55, 0xd9, 0xc9, 0x3e, 0x64,
	0x4a, 0x4d, 0xdb, 0xb2, 0x53, 0x13, 0x88, 0xb4, 0xa2, 0x35, 0xba, 0x0a, 0x4b, 0x71, 0x45, 0x8e,
	0x43, 0x36, 0xe5, 0x21, 0xb5, 0x50, 0x79, 0xc0, 0xeb, 0x9a, 0x6e, 0x23, 0x23, 0x5b, 0xb1, 0x13,
	0xd5, 0xd4, 0xd0, 0xc8, 0x48, 0xe2, 0x6f, 0x39, 0xfa, 0xc4, 0x8a, 0x8c, 0xfc, 0x53, 0xf8, 0xb1,
	0x00, 0x85, 0x47, 0x28, 0x89, 0x77, 0xfd, 0x56, 0xcb, 0xeb, 0xca, 0xa6, 0xa7, 0x5a, 0x8a, 0xeb,
	0xd4, 0xe3, 0x08, 0xdc, 0xd6, 0x2b, 0x15, 0xb4, 0xc9, 0x0c, 0x3d, 0x82, 0xbd, 0xf1, 0x61, 0x8a,
	0xb2, 0x65, 0x0f, 0xda, 0x46, 0x72, 0xf4, 0x00, 0x76, 0xfa, 0x83, 0x10, 0x9a, 0x65, 0x9a, 0xa8,
	0x71, 0x64, 0x82, 0x5b, 0xfd, 0xd6, 0x91, 0xd9, 0x4c, 0x25, 0xca, 0x32, 0x91, 0x91, 0x39, 0xd9,
	0xe6, 0x4c, 0x9d, 0x59, 0x2e, 0x17, 0x78, 0xdd, 0x9f, 0xf4, 0xb8, 0xdc, 0x92, 0x1c, 0x86, 0xd3,
	0x5f, 0x98, 0x22, 0x1b, 0x87, 0xab, 0x36, 0x8f, 0xa5, 0x77, 0x00, 0x3b, 0x53, 0x64, 0x85, 0x8c,
	0x2c, 0x8d, 0xde, 0x86, 0x51, 0x7a, 0x2a, 0x34, 0x92, 0xa7, 0x67, 0xf0, 0xe4, 0x41, 0x79, 0xa5,
	0xa1, 0x96, 0x69, 0x01, 0x4e, 0x1f, 0x46, 0x0e, 0xbc, 0xae, 0xd0, 0x67, 0x70, 0xf2, 0x20, 0x36,
	0x49, 0x10, 0x64, 0x82, 0x19, 0x7a, 0x1d, 0xba, 0x5a, 0xa5, 0x27, 0x70, 0x94, 0xad, 0xd6, 0x21,
	0x68, 0x8d, 0x9e, 0xc2, 0x71, 0xc6, 0xed, 0x11, 0x58, 0x47, 0x93, 0x0b, 0x1b, 0x35, 0xd4, 0xeb,
	0xc8, 0xc8, 0x36, 0x2d, 0xc2, 0xf3, 0x09, 0xc5, 0x5a, 0xe5, 0x14, 0xce, 0x54, 0xae, 0x8e, 0xb8,
	0x5d, 0x8f, 0x4b, 0xfe, 0x0d, 0x3c, 0xa9, 0x64, 0x83, 0x02, 0x2c, 0x26, 0xdf, 0x7f, 0xc9, 0xcb,
	0x1c, 0x0b, 0xbc, 0xcf, 0x18, 0x0c, 0x9b, 0xd0, 0x5d, 0xf8, 0x77, 0x4c, 0xec, 0xc3, 0xb5, 0xb6,
	0x29, 0xaf, 0x6e, 0x96, 0x4d, 0x94, 0x5d, 0xa3, 0xac, 0xc7, 0xfa, 0xa7, 0xe7, 0x5f, 0x73, 0xf0,
	0xba, 0x19, 0x76, 0x4a, 0x7f, 0xb2, 0x86, 0xcf, 0x8b, 0x8f, 0xdd, 0xc1, 0x35, 0xb9, 0xeb, 0x6b,
	0xb9, 0xb7, 0xb5, 0x96, 0x1f, 0xbd, 0xbf, 0xbb, 0x2d, 0x35, 0xc3, 0x8e, 0x32, 0x88, 0x58, 0x6c,
	0xf8, 0xf2, 0xe5, 0xf0, 0xba, 0x41, 0xa3, 0x5d, 0x8c, 0x5f, 0x85, 0x9e, 0xd2, 0xeb, 0x36, 0x95,
	0x4e, 0xc3, 0x0f, 0x94, 0xf8, 0xac, 0x4c, 0x7d, 0x62, 0xbe, 0xe7, 0x72, 0x5f, 0x66, 0xe7, 0x54,
	0x97, 0xdf, 0x2e, 0xc6, 0xc8, 0x97, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x88, 0xa4, 0xab, 0xb5,
	0x8f, 0x06, 0x00, 0x00,
}
