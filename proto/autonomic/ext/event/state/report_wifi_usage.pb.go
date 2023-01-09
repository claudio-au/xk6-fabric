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
// source: autonomic/ext/event/state/report_wifi_usage.proto

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

type WellKnownReportWifiUsageState int32

const (
	WellKnownReportWifiUsageState_UNKNOWN_STATE WellKnownReportWifiUsageState = 0
	// Request was received by cloud through API
	// Emitted by: cloud
	WellKnownReportWifiUsageState_REQUESTED WellKnownReportWifiUsageState = 1
	// Request was queued in the cloud, waiting for device
	// to connect.
	// Emitted by: cloud
	// Triggered by: REQUEST_QUEUED_FOR_DELIVERY
	WellKnownReportWifiUsageState_REQUEST_DELIVERY_QUEUED WellKnownReportWifiUsageState = 2
	// Cloud attempted delivery of request to device
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_TRANSLATION_SUCCEEDED
	WellKnownReportWifiUsageState_REQUEST_DELIVERY_IN_PROGRESS WellKnownReportWifiUsageState = 3
	// Delivery of request attempt was not acknowledged by device in a timely fashion
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_TIMEOUT_EXCEEDED
	WellKnownReportWifiUsageState_REQUEST_DELIVERY_TIMED_OUT WellKnownReportWifiUsageState = 4
	// Delivery of request attempt failed due to device disconnect
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_FAILED
	WellKnownReportWifiUsageState_REQUEST_DELIVERY_FAILURE WellKnownReportWifiUsageState = 5
	// Request is acknowledged by device, queued, and the device will begin
	// processing the request when the device is ready.
	// Emitted by: device
	// Triggered by: REQUEST_QUEUED_ON_DEVICE
	WellKnownReportWifiUsageState_REQUEST_QUEUED WellKnownReportWifiUsageState = 6
	// Request is being decrypted and validated by device
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_STARTED
	WellKnownReportWifiUsageState_REQUEST_VALIDATION_IN_PROGRESS WellKnownReportWifiUsageState = 7
	// Request failed validation or decryption
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_FAILED
	WellKnownReportWifiUsageState_REQUEST_VALIDATION_FAILURE WellKnownReportWifiUsageState = 8
	// Wifi usage reporting in progress
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED
	WellKnownReportWifiUsageState_REPORT_WIFI_USAGE_IN_PROGRESS WellKnownReportWifiUsageState = 9
	// The device has successfully completed the report WiFi usage command in question.
	// Triggered by: REPORT_WIFI_USAGE_SUCCEEDED
	WellKnownReportWifiUsageState_SUCCEEDED_ON_DEVICE WellKnownReportWifiUsageState = 23
	// Oem cloud attempted delivery of request
	// Emitted by: cloud
	// Triggered by: DEVICE_CONNECTED_TO_CLOUD
	WellKnownReportWifiUsageState_OEM_CLOUD_DELIVERY_IN_PROGRESS WellKnownReportWifiUsageState = 10
	// Cloud attempted translation of request
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_DELIVERY_SUCCEEDED
	WellKnownReportWifiUsageState_OEM_CLOUD_TRANSLATION_IN_PROGRESS WellKnownReportWifiUsageState = 11
	// The cloud is in the process of validating the event emitted by the device indicating that the
	// command has succeeded.
	// Triggered by: SUCCEEDED_ON_DEVICE_EVENT_RECEIVED
	WellKnownReportWifiUsageState_CLOUD_VALIDATION_IN_PROGRESS WellKnownReportWifiUsageState = 24
	// Cloud failed to validate data provided by the device
	// Emitted by: cloud
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED
	WellKnownReportWifiUsageState_CLOUD_VALIDATION_FAILURE WellKnownReportWifiUsageState = 12
	// Wifi usage reporting succeeded
	// added to the event containing this state.
	// This state is terminal
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED
	WellKnownReportWifiUsageState_SUCCESS WellKnownReportWifiUsageState = 17
	// Wifi usage reporting failed
	// this state.
	// Emitted by: cloud
	// Triggered by: REPORT_WIFI_USAGE_FAILED
	WellKnownReportWifiUsageState_FAILURE WellKnownReportWifiUsageState = 18
	// Wifi usage reporting command expired. This event can be emitted by
	// the device or the cloud. In either case, it will cause
	// the command to be removed from the set of in-progress commands
	// on the cloud. However the cloud does not send this event to
	// the device. This state is terminal.
	// Triggered by: EXPIRATION_EXCEEDED
	WellKnownReportWifiUsageState_EXPIRED WellKnownReportWifiUsageState = 19
	// Command is being removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUESTED
	WellKnownReportWifiUsageState_CANCELLATION_IN_PROGRESS WellKnownReportWifiUsageState = 21
	// Command was successfully removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_FULFILLED
	WellKnownReportWifiUsageState_CANCELLED WellKnownReportWifiUsageState = 22
)

var WellKnownReportWifiUsageState_name = map[int32]string{
	0:  "UNKNOWN_STATE",
	1:  "REQUESTED",
	2:  "REQUEST_DELIVERY_QUEUED",
	3:  "REQUEST_DELIVERY_IN_PROGRESS",
	4:  "REQUEST_DELIVERY_TIMED_OUT",
	5:  "REQUEST_DELIVERY_FAILURE",
	6:  "REQUEST_QUEUED",
	7:  "REQUEST_VALIDATION_IN_PROGRESS",
	8:  "REQUEST_VALIDATION_FAILURE",
	9:  "REPORT_WIFI_USAGE_IN_PROGRESS",
	23: "SUCCEEDED_ON_DEVICE",
	10: "OEM_CLOUD_DELIVERY_IN_PROGRESS",
	11: "OEM_CLOUD_TRANSLATION_IN_PROGRESS",
	24: "CLOUD_VALIDATION_IN_PROGRESS",
	12: "CLOUD_VALIDATION_FAILURE",
	17: "SUCCESS",
	18: "FAILURE",
	19: "EXPIRED",
	21: "CANCELLATION_IN_PROGRESS",
	22: "CANCELLED",
}

var WellKnownReportWifiUsageState_value = map[string]int32{
	"UNKNOWN_STATE":                     0,
	"REQUESTED":                         1,
	"REQUEST_DELIVERY_QUEUED":           2,
	"REQUEST_DELIVERY_IN_PROGRESS":      3,
	"REQUEST_DELIVERY_TIMED_OUT":        4,
	"REQUEST_DELIVERY_FAILURE":          5,
	"REQUEST_QUEUED":                    6,
	"REQUEST_VALIDATION_IN_PROGRESS":    7,
	"REQUEST_VALIDATION_FAILURE":        8,
	"REPORT_WIFI_USAGE_IN_PROGRESS":     9,
	"SUCCEEDED_ON_DEVICE":               23,
	"OEM_CLOUD_DELIVERY_IN_PROGRESS":    10,
	"OEM_CLOUD_TRANSLATION_IN_PROGRESS": 11,
	"CLOUD_VALIDATION_IN_PROGRESS":      24,
	"CLOUD_VALIDATION_FAILURE":          12,
	"SUCCESS":                           17,
	"FAILURE":                           18,
	"EXPIRED":                           19,
	"CANCELLATION_IN_PROGRESS":          21,
	"CANCELLED":                         22,
}

func (x WellKnownReportWifiUsageState) String() string {
	return proto.EnumName(WellKnownReportWifiUsageState_name, int32(x))
}

func (WellKnownReportWifiUsageState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7f54877d3bc13075, []int{0}
}

// Wifi usage display state transition triggers. See WellKnownReportWifiUsageState enum
// for documentation on when these triggers occur.
type WellKnownReportWifiUsageTransitionTrigger int32

const (
	WellKnownReportWifiUsageTransitionTrigger_UNKNOWN_TRANSITION_TRIGGER                WellKnownReportWifiUsageTransitionTrigger = 0
	WellKnownReportWifiUsageTransitionTrigger_REQUEST_QUEUED_FOR_DELIVERY               WellKnownReportWifiUsageTransitionTrigger = 1
	WellKnownReportWifiUsageTransitionTrigger_DEVICE_CONNECTED_TO_CLOUD                 WellKnownReportWifiUsageTransitionTrigger = 2
	WellKnownReportWifiUsageTransitionTrigger_REQUEST_DELIVERY_FAILED                   WellKnownReportWifiUsageTransitionTrigger = 3
	WellKnownReportWifiUsageTransitionTrigger_REQUEST_DELIVERY_TIMEOUT_EXCEEDED         WellKnownReportWifiUsageTransitionTrigger = 4
	WellKnownReportWifiUsageTransitionTrigger_REQUEST_QUEUED_ON_DEVICE                  WellKnownReportWifiUsageTransitionTrigger = 5
	WellKnownReportWifiUsageTransitionTrigger_REQUEST_VALIDATION_STARTED                WellKnownReportWifiUsageTransitionTrigger = 6
	WellKnownReportWifiUsageTransitionTrigger_REQUEST_VALIDATION_FAILED                 WellKnownReportWifiUsageTransitionTrigger = 7
	WellKnownReportWifiUsageTransitionTrigger_REQUEST_VALIDATION_SUCCEEDED              WellKnownReportWifiUsageTransitionTrigger = 23
	WellKnownReportWifiUsageTransitionTrigger_REPORT_WIFI_USAGE_STARTED                 WellKnownReportWifiUsageTransitionTrigger = 8
	WellKnownReportWifiUsageTransitionTrigger_REPORT_WIFI_USAGE_SUCCEEDED               WellKnownReportWifiUsageTransitionTrigger = 9
	WellKnownReportWifiUsageTransitionTrigger_REPORT_WIFI_USAGE_FAILED                  WellKnownReportWifiUsageTransitionTrigger = 10
	WellKnownReportWifiUsageTransitionTrigger_OEM_CLOUD_DELIVERY_SUCCEEDED              WellKnownReportWifiUsageTransitionTrigger = 11
	WellKnownReportWifiUsageTransitionTrigger_OEM_CLOUD_TRANSLATION_SUCCEEDED           WellKnownReportWifiUsageTransitionTrigger = 12
	WellKnownReportWifiUsageTransitionTrigger_SUCCEEDED_ON_DEVICE_EVENT_RECEIVED        WellKnownReportWifiUsageTransitionTrigger = 24
	WellKnownReportWifiUsageTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED WellKnownReportWifiUsageTransitionTrigger = 13
	WellKnownReportWifiUsageTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED    WellKnownReportWifiUsageTransitionTrigger = 14
	WellKnownReportWifiUsageTransitionTrigger_FAILED                                    WellKnownReportWifiUsageTransitionTrigger = 19
	WellKnownReportWifiUsageTransitionTrigger_EXPIRATION_EXCEEDED                       WellKnownReportWifiUsageTransitionTrigger = 20
	WellKnownReportWifiUsageTransitionTrigger_CANCELLATION_REQUESTED                    WellKnownReportWifiUsageTransitionTrigger = 21
	WellKnownReportWifiUsageTransitionTrigger_CANCELLATION_REQUEST_FULFILLED            WellKnownReportWifiUsageTransitionTrigger = 22
)

var WellKnownReportWifiUsageTransitionTrigger_name = map[int32]string{
	0:  "UNKNOWN_TRANSITION_TRIGGER",
	1:  "REQUEST_QUEUED_FOR_DELIVERY",
	2:  "DEVICE_CONNECTED_TO_CLOUD",
	3:  "REQUEST_DELIVERY_FAILED",
	4:  "REQUEST_DELIVERY_TIMEOUT_EXCEEDED",
	5:  "REQUEST_QUEUED_ON_DEVICE",
	6:  "REQUEST_VALIDATION_STARTED",
	7:  "REQUEST_VALIDATION_FAILED",
	23: "REQUEST_VALIDATION_SUCCEEDED",
	8:  "REPORT_WIFI_USAGE_STARTED",
	9:  "REPORT_WIFI_USAGE_SUCCEEDED",
	10: "REPORT_WIFI_USAGE_FAILED",
	11: "OEM_CLOUD_DELIVERY_SUCCEEDED",
	12: "OEM_CLOUD_TRANSLATION_SUCCEEDED",
	24: "SUCCEEDED_ON_DEVICE_EVENT_RECEIVED",
	13: "CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED",
	14: "CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED",
	19: "FAILED",
	20: "EXPIRATION_EXCEEDED",
	21: "CANCELLATION_REQUESTED",
	22: "CANCELLATION_REQUEST_FULFILLED",
}

var WellKnownReportWifiUsageTransitionTrigger_value = map[string]int32{
	"UNKNOWN_TRANSITION_TRIGGER":                0,
	"REQUEST_QUEUED_FOR_DELIVERY":               1,
	"DEVICE_CONNECTED_TO_CLOUD":                 2,
	"REQUEST_DELIVERY_FAILED":                   3,
	"REQUEST_DELIVERY_TIMEOUT_EXCEEDED":         4,
	"REQUEST_QUEUED_ON_DEVICE":                  5,
	"REQUEST_VALIDATION_STARTED":                6,
	"REQUEST_VALIDATION_FAILED":                 7,
	"REQUEST_VALIDATION_SUCCEEDED":              23,
	"REPORT_WIFI_USAGE_STARTED":                 8,
	"REPORT_WIFI_USAGE_SUCCEEDED":               9,
	"REPORT_WIFI_USAGE_FAILED":                  10,
	"OEM_CLOUD_DELIVERY_SUCCEEDED":              11,
	"OEM_CLOUD_TRANSLATION_SUCCEEDED":           12,
	"SUCCEEDED_ON_DEVICE_EVENT_RECEIVED":        24,
	"CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED": 13,
	"CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED":    14,
	"FAILED":                                    19,
	"EXPIRATION_EXCEEDED":                       20,
	"CANCELLATION_REQUESTED":                    21,
	"CANCELLATION_REQUEST_FULFILLED":            22,
}

func (x WellKnownReportWifiUsageTransitionTrigger) String() string {
	return proto.EnumName(WellKnownReportWifiUsageTransitionTrigger_name, int32(x))
}

func (WellKnownReportWifiUsageTransitionTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7f54877d3bc13075, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.event.state.reportwifiusage.WellKnownReportWifiUsageState", WellKnownReportWifiUsageState_name, WellKnownReportWifiUsageState_value)
	proto.RegisterEnum("autonomic.ext.event.state.reportwifiusage.WellKnownReportWifiUsageTransitionTrigger", WellKnownReportWifiUsageTransitionTrigger_name, WellKnownReportWifiUsageTransitionTrigger_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/event/state/report_wifi_usage.proto", fileDescriptor_7f54877d3bc13075)
}

var fileDescriptor_7f54877d3bc13075 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x4e, 0xdb, 0x4a,
	0x14, 0x25, 0xdc, 0xd9, 0x5c, 0x34, 0x0c, 0x07, 0xc2, 0x2d, 0x70, 0xe0, 0xe8, 0x20, 0x81, 0x94,
	0x58, 0x47, 0xe7, 0x0b, 0x8c, 0x67, 0x27, 0xb2, 0x08, 0x76, 0x18, 0x8f, 0x13, 0xda, 0x97, 0x51,
	0x88, 0x4c, 0x6a, 0x89, 0xd8, 0xc8, 0x31, 0x85, 0xef, 0xe9, 0x17, 0xb4, 0xdf, 0xd3, 0x1f, 0xe9,
	0x63, 0x35, 0x76, 0x9c, 0x0b, 0x71, 0xda, 0xbe, 0x65, 0xb2, 0xd6, 0xda, 0xd7, 0xe5, 0x0d, 0xff,
	0xb5, 0x5f, 0xe2, 0x30, 0x08, 0x7b, 0x7e, 0x47, 0xf3, 0xde, 0x62, 0xcd, 0xfb, 0xec, 0x05, 0xb1,
	0xd6, 0x8f, 0xdb, 0xb1, 0xa7, 0x45, 0xde, 0x73, 0x18, 0xc5, 0xf2, 0xd5, 0x7f, 0xf4, 0xe5, 0x4b,
	0xbf, 0xdd, 0xf5, 0x2a, 0xcf, 0x51, 0x18, 0x87, 0xf4, 0x72, 0x28, 0xa9, 0x78, 0x6f, 0x71, 0x25,
	0x91, 0x54, 0x12, 0x49, 0x25, 0x95, 0x28, 0x45, 0x22, 0xb8, 0xfa, 0xb6, 0x08, 0xa5, 0x96, 0xf7,
	0xf4, 0x74, 0x13, 0x84, 0xaf, 0x01, 0x4f, 0xc0, 0x96, 0xff, 0xe8, 0xbb, 0x0a, 0x74, 0x94, 0x82,
	0x6e, 0xc3, 0xa6, 0x6b, 0xdd, 0x58, 0x76, 0xcb, 0x92, 0x8e, 0xd0, 0x05, 0x92, 0x39, 0xba, 0x09,
	0x6b, 0x1c, 0xef, 0x5c, 0x74, 0x04, 0x32, 0x52, 0xa0, 0x47, 0x50, 0x1c, 0x3c, 0x25, 0xc3, 0xba,
	0xd9, 0x44, 0xfe, 0x41, 0xde, 0xb9, 0xe8, 0x22, 0x23, 0xf3, 0xf4, 0x6f, 0x38, 0x9e, 0x02, 0x4d,
	0x4b, 0x36, 0xb8, 0x5d, 0xe3, 0xe8, 0x38, 0x64, 0x81, 0x9e, 0xc0, 0xe1, 0x14, 0x43, 0x98, 0xb7,
	0xc8, 0xa4, 0xed, 0x0a, 0xb2, 0x48, 0x8f, 0x61, 0x7f, 0x0a, 0xaf, 0xea, 0x66, 0xdd, 0xe5, 0x48,
	0x96, 0x28, 0x85, 0xad, 0x0c, 0x1d, 0xe4, 0x5c, 0xa6, 0xe7, 0x70, 0x92, 0xfd, 0xd7, 0xd4, 0xeb,
	0x26, 0xd3, 0x85, 0x69, 0x5b, 0x13, 0x59, 0x57, 0xc6, 0xb3, 0x8e, 0x71, 0xb2, 0xb8, 0xab, 0xf4,
	0x0c, 0x4a, 0x1c, 0x1b, 0x36, 0x17, 0xb2, 0x65, 0x56, 0x4d, 0xe9, 0x3a, 0x7a, 0x0d, 0x27, 0x42,
	0xac, 0xd1, 0x22, 0xec, 0x38, 0xae, 0x61, 0x20, 0x32, 0x55, 0xab, 0x25, 0x19, 0x36, 0x4d, 0x03,
	0x49, 0x51, 0xe5, 0xb7, 0xf1, 0x56, 0x1a, 0x75, 0xdb, 0x65, 0xf9, 0x5d, 0x03, 0xfd, 0x17, 0xce,
	0x46, 0x1c, 0xc1, 0x75, 0xcb, 0xa9, 0x4f, 0x97, 0xb9, 0xae, 0xc6, 0x97, 0x52, 0x66, 0x34, 0xb2,
	0xaf, 0xc6, 0x33, 0xc5, 0xc8, 0xda, 0xd8, 0xa0, 0xeb, 0xb0, 0x92, 0xd4, 0xe8, 0x38, 0x64, 0x5b,
	0x3d, 0x32, 0x84, 0xaa, 0x07, 0xde, 0x37, 0x4c, 0x8e, 0x8c, 0xec, 0x24, 0x41, 0x74, 0xcb, 0xc0,
	0x7a, 0x4e, 0x11, 0xbb, 0x6a, 0xdf, 0x03, 0x14, 0x19, 0xd9, 0xbb, 0xfa, 0xbe, 0x04, 0x97, 0xb3,
	0x3c, 0x23, 0xa2, 0x76, 0xd0, 0xf7, 0x63, 0x3f, 0x0c, 0x44, 0xe4, 0x77, 0xbb, 0x5e, 0xa4, 0x06,
	0x9d, 0xf9, 0x27, 0x69, 0xd3, 0x4c, 0x12, 0x08, 0x6e, 0xd6, 0x6a, 0xc8, 0xc9, 0x1c, 0x3d, 0x85,
	0xa3, 0xc9, 0x05, 0xca, 0xaa, 0xcd, 0x87, 0x53, 0x23, 0x05, 0x5a, 0x82, 0x83, 0x74, 0xb2, 0xd2,
	0xb0, 0x2d, 0x0b, 0x0d, 0x81, 0x4c, 0x0a, 0x3b, 0x9d, 0x1c, 0x99, 0xcf, 0x75, 0x9f, 0xea, 0x12,
	0x19, 0x59, 0x50, 0x53, 0xce, 0xf5, 0x96, 0xed, 0x0a, 0x89, 0xf7, 0xe9, 0xea, 0x26, 0x2d, 0x36,
	0xa8, 0x61, 0xb4, 0xce, 0xa5, 0x19, 0x56, 0x71, 0x84, 0xce, 0x45, 0x62, 0xb7, 0x12, 0x1c, 0xcc,
	0xb0, 0x12, 0x32, 0xb2, 0x32, 0xfe, 0x05, 0x8c, 0xcb, 0x33, 0xe7, 0x90, 0x62, 0x1a, 0xe0, 0xbd,
	0xd7, 0xb2, 0xf8, 0xab, 0xe9, 0x84, 0xa6, 0xe0, 0xa1, 0x7e, 0x2d, 0x2d, 0xff, 0x3d, 0x61, 0x90,
	0x1f, 0x54, 0xfe, 0x1c, 0x37, 0x8e, 0xf4, 0xeb, 0xf4, 0x1f, 0x38, 0xcd, 0xf7, 0xe2, 0x88, 0xb4,
	0x41, 0x2f, 0xe0, 0x3c, 0xc7, 0xed, 0x12, 0x9b, 0x68, 0x09, 0xc9, 0xd1, 0x40, 0xb3, 0x89, 0x8c,
	0xec, 0xd3, 0x32, 0x5c, 0x4e, 0xf9, 0xd1, 0xae, 0x66, 0x74, 0xa6, 0x0b, 0x7d, 0x2c, 0xec, 0x26,
	0xbd, 0x82, 0x8b, 0xdf, 0xd1, 0x07, 0x9d, 0x6c, 0x51, 0x80, 0xe5, 0xc1, 0xef, 0x1d, 0xf5, 0xf1,
	0x25, 0xf6, 0x4d, 0x15, 0xc3, 0x5d, 0xfe, 0x45, 0x0f, 0x61, 0x6f, 0xc2, 0xca, 0xa3, 0x4b, 0xb5,
	0xab, 0x3e, 0xcc, 0x3c, 0x4c, 0x56, 0xdd, 0x7a, 0xd5, 0x4c, 0xdd, 0x7d, 0xfd, 0xb5, 0x00, 0xe5,
	0x4e, 0xd8, 0xab, 0xfc, 0xf1, 0x0d, 0xbd, 0x3e, 0xff, 0xe5, 0x01, 0x6d, 0xa8, 0x93, 0xdc, 0x28,
	0x7c, 0x6c, 0x74, 0xfd, 0xf8, 0xd3, 0xcb, 0x43, 0xa5, 0x13, 0xf6, 0xb4, 0x61, 0xec, 0x72, 0xdb,
	0x57, 0x57, 0xdd, 0x8b, 0x82, 0xf6, 0x53, 0x39, 0x39, 0xde, 0x7d, 0xad, 0x1f, 0x75, 0xb4, 0x5e,
	0xdb, 0x0f, 0xb4, 0xe4, 0xad, 0xcd, 0x3c, 0xff, 0x3f, 0x0a, 0x85, 0x2f, 0xf3, 0x0b, 0xba, 0x2b,
	0x1e, 0x96, 0x13, 0xe6, 0xff, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x66, 0x85, 0x15, 0xfb, 0x2b,
	0x06, 0x00, 0x00,
}
