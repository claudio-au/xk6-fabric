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
// source: autonomic/ext/event/state/wifi_status.proto

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

type WellKnownWifiStatusState int32

const (
	WellKnownWifiStatusState_UNKNOWN_STATE WellKnownWifiStatusState = 0
	// Request was received by cloud through API
	// Emitted by: cloud
	WellKnownWifiStatusState_REQUESTED WellKnownWifiStatusState = 1
	// Request was queued in the cloud, waiting for device
	// to connect.
	// Emitted by: cloud
	// Triggered by: REQUEST_QUEUED_FOR_DELIVERY
	WellKnownWifiStatusState_REQUEST_DELIVERY_QUEUED WellKnownWifiStatusState = 2
	// Cloud attempted delivery of request to device
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_TRANSLATION_SUCCEEDED
	WellKnownWifiStatusState_REQUEST_DELIVERY_IN_PROGRESS WellKnownWifiStatusState = 3
	// Delivery of request attempt was not acknowledged by device in a timely fashion
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_TIMEOUT_EXCEEDED
	WellKnownWifiStatusState_REQUEST_DELIVERY_TIMED_OUT WellKnownWifiStatusState = 4
	// Delivery of request attempt failed due to device disconnect
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_FAILED
	WellKnownWifiStatusState_REQUEST_DELIVERY_FAILURE WellKnownWifiStatusState = 5
	// Request is acknowledged by device, queued, and the device will begin
	// processing the request when the device is ready.
	// Emitted by: device
	// Triggered by: REQUEST_QUEUED_ON_DEVICE
	WellKnownWifiStatusState_REQUEST_QUEUED WellKnownWifiStatusState = 6
	// Request is being decrypted and validated by device
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_STARTED
	WellKnownWifiStatusState_REQUEST_VALIDATION_IN_PROGRESS WellKnownWifiStatusState = 7
	// Request failed validation or decryption
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_FAILED
	WellKnownWifiStatusState_REQUEST_VALIDATION_FAILURE WellKnownWifiStatusState = 8
	// Wifi configuration in progress
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED
	WellKnownWifiStatusState_WIFI_STATUS_IN_PROGRESS WellKnownWifiStatusState = 9
	// The device has successfully completed the wifi status command in question.
	// Triggered by: WIFI_STATUS_SUCCEEDED
	WellKnownWifiStatusState_SUCCEEDED_ON_DEVICE WellKnownWifiStatusState = 23
	// Oem cloud attempted delivery of request
	// Emitted by: cloud
	// Triggered by: DEVICE_CONNECTED_TO_CLOUD
	WellKnownWifiStatusState_OEM_CLOUD_DELIVERY_IN_PROGRESS WellKnownWifiStatusState = 10
	// Cloud attempted translation of request
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_DELIVERY_SUCCEEDED
	WellKnownWifiStatusState_OEM_CLOUD_TRANSLATION_IN_PROGRESS WellKnownWifiStatusState = 11
	// The cloud is in the process of validating the event emitted by the device indicating that the
	// command has succeeded.
	// Triggered by: SUCCEEDED_ON_DEVICE_EVENT_RECEIVED
	WellKnownWifiStatusState_CLOUD_VALIDATION_IN_PROGRESS WellKnownWifiStatusState = 24
	// Cloud failed to validate data provided by the device
	// Emitted by: cloud
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED
	WellKnownWifiStatusState_CLOUD_VALIDATION_FAILURE WellKnownWifiStatusState = 12
	// Wifi configuration succeeded
	// added to the event containing this state.
	// This state is terminal
	// Triggered by: WIFI_STATUS_SUCCEEDED
	WellKnownWifiStatusState_SUCCESS WellKnownWifiStatusState = 17
	// Wifi configuration failed
	// this state.
	// Emitted by: device
	// Triggered by: WIFI_STATUS_FAILED
	WellKnownWifiStatusState_FAILURE WellKnownWifiStatusState = 18
	// Wifi configuration command expired. This event can be emitted by
	// the device or the cloud. In either case, it will cause
	// the command to be removed from the set of in-progress commands
	// on the cloud. However the cloud does not send this event to
	// the device. This state is terminal.
	// Triggered by: EXPIRATION_EXCEEDED
	WellKnownWifiStatusState_EXPIRED WellKnownWifiStatusState = 19
	// Command is being removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_ACKNOWLEDGED
	WellKnownWifiStatusState_CANCELLATION_IN_PROGRESS WellKnownWifiStatusState = 21
	// Command was successfully removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_FULFILLED
	WellKnownWifiStatusState_CANCELLED WellKnownWifiStatusState = 22
)

var WellKnownWifiStatusState_name = map[int32]string{
	0:  "UNKNOWN_STATE",
	1:  "REQUESTED",
	2:  "REQUEST_DELIVERY_QUEUED",
	3:  "REQUEST_DELIVERY_IN_PROGRESS",
	4:  "REQUEST_DELIVERY_TIMED_OUT",
	5:  "REQUEST_DELIVERY_FAILURE",
	6:  "REQUEST_QUEUED",
	7:  "REQUEST_VALIDATION_IN_PROGRESS",
	8:  "REQUEST_VALIDATION_FAILURE",
	9:  "WIFI_STATUS_IN_PROGRESS",
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

var WellKnownWifiStatusState_value = map[string]int32{
	"UNKNOWN_STATE":                     0,
	"REQUESTED":                         1,
	"REQUEST_DELIVERY_QUEUED":           2,
	"REQUEST_DELIVERY_IN_PROGRESS":      3,
	"REQUEST_DELIVERY_TIMED_OUT":        4,
	"REQUEST_DELIVERY_FAILURE":          5,
	"REQUEST_QUEUED":                    6,
	"REQUEST_VALIDATION_IN_PROGRESS":    7,
	"REQUEST_VALIDATION_FAILURE":        8,
	"WIFI_STATUS_IN_PROGRESS":           9,
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

func (x WellKnownWifiStatusState) String() string {
	return proto.EnumName(WellKnownWifiStatusState_name, int32(x))
}

func (WellKnownWifiStatusState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_330a863691c5dd04, []int{0}
}

// Wifi configuration state transition triggers. See WellKnownWifiStatusState enum
// for documentation on when these triggers occur.
type WellKnownWifiStatusTransitionTrigger int32

const (
	WellKnownWifiStatusTransitionTrigger_UNKNOWN_TRANSITION_TRIGGER                WellKnownWifiStatusTransitionTrigger = 0
	WellKnownWifiStatusTransitionTrigger_REQUEST_QUEUED_FOR_DELIVERY               WellKnownWifiStatusTransitionTrigger = 1
	WellKnownWifiStatusTransitionTrigger_DEVICE_CONNECTED_TO_CLOUD                 WellKnownWifiStatusTransitionTrigger = 2
	WellKnownWifiStatusTransitionTrigger_REQUEST_DELIVERY_FAILED                   WellKnownWifiStatusTransitionTrigger = 3
	WellKnownWifiStatusTransitionTrigger_REQUEST_DELIVERY_TIMEOUT_EXCEEDED         WellKnownWifiStatusTransitionTrigger = 4
	WellKnownWifiStatusTransitionTrigger_REQUEST_QUEUED_ON_DEVICE                  WellKnownWifiStatusTransitionTrigger = 5
	WellKnownWifiStatusTransitionTrigger_REQUEST_VALIDATION_STARTED                WellKnownWifiStatusTransitionTrigger = 6
	WellKnownWifiStatusTransitionTrigger_REQUEST_VALIDATION_FAILED                 WellKnownWifiStatusTransitionTrigger = 7
	WellKnownWifiStatusTransitionTrigger_REQUEST_VALIDATION_SUCCEEDED              WellKnownWifiStatusTransitionTrigger = 23
	WellKnownWifiStatusTransitionTrigger_WIFI_STATUS_STARTED                       WellKnownWifiStatusTransitionTrigger = 8
	WellKnownWifiStatusTransitionTrigger_WIFI_STATUS_SUCCEEDED                     WellKnownWifiStatusTransitionTrigger = 9
	WellKnownWifiStatusTransitionTrigger_WIFI_STATUS_FAILED                        WellKnownWifiStatusTransitionTrigger = 10
	WellKnownWifiStatusTransitionTrigger_OEM_CLOUD_DELIVERY_SUCCEEDED              WellKnownWifiStatusTransitionTrigger = 11
	WellKnownWifiStatusTransitionTrigger_OEM_CLOUD_TRANSLATION_SUCCEEDED           WellKnownWifiStatusTransitionTrigger = 12
	WellKnownWifiStatusTransitionTrigger_SUCCEEDED_ON_DEVICE_EVENT_RECEIVED        WellKnownWifiStatusTransitionTrigger = 24
	WellKnownWifiStatusTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED WellKnownWifiStatusTransitionTrigger = 13
	WellKnownWifiStatusTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED    WellKnownWifiStatusTransitionTrigger = 14
	WellKnownWifiStatusTransitionTrigger_FAILED                                    WellKnownWifiStatusTransitionTrigger = 19
	WellKnownWifiStatusTransitionTrigger_EXPIRATION_EXCEEDED                       WellKnownWifiStatusTransitionTrigger = 20
	WellKnownWifiStatusTransitionTrigger_CANCELLATION_REQUESTED                    WellKnownWifiStatusTransitionTrigger = 21
	WellKnownWifiStatusTransitionTrigger_CANCELLATION_REQUEST_FULFILLED            WellKnownWifiStatusTransitionTrigger = 22
)

var WellKnownWifiStatusTransitionTrigger_name = map[int32]string{
	0:  "UNKNOWN_TRANSITION_TRIGGER",
	1:  "REQUEST_QUEUED_FOR_DELIVERY",
	2:  "DEVICE_CONNECTED_TO_CLOUD",
	3:  "REQUEST_DELIVERY_FAILED",
	4:  "REQUEST_DELIVERY_TIMEOUT_EXCEEDED",
	5:  "REQUEST_QUEUED_ON_DEVICE",
	6:  "REQUEST_VALIDATION_STARTED",
	7:  "REQUEST_VALIDATION_FAILED",
	23: "REQUEST_VALIDATION_SUCCEEDED",
	8:  "WIFI_STATUS_STARTED",
	9:  "WIFI_STATUS_SUCCEEDED",
	10: "WIFI_STATUS_FAILED",
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

var WellKnownWifiStatusTransitionTrigger_value = map[string]int32{
	"UNKNOWN_TRANSITION_TRIGGER":                0,
	"REQUEST_QUEUED_FOR_DELIVERY":               1,
	"DEVICE_CONNECTED_TO_CLOUD":                 2,
	"REQUEST_DELIVERY_FAILED":                   3,
	"REQUEST_DELIVERY_TIMEOUT_EXCEEDED":         4,
	"REQUEST_QUEUED_ON_DEVICE":                  5,
	"REQUEST_VALIDATION_STARTED":                6,
	"REQUEST_VALIDATION_FAILED":                 7,
	"REQUEST_VALIDATION_SUCCEEDED":              23,
	"WIFI_STATUS_STARTED":                       8,
	"WIFI_STATUS_SUCCEEDED":                     9,
	"WIFI_STATUS_FAILED":                        10,
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

func (x WellKnownWifiStatusTransitionTrigger) String() string {
	return proto.EnumName(WellKnownWifiStatusTransitionTrigger_name, int32(x))
}

func (WellKnownWifiStatusTransitionTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_330a863691c5dd04, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.event.state.wifistatus.WellKnownWifiStatusState", WellKnownWifiStatusState_name, WellKnownWifiStatusState_value)
	proto.RegisterEnum("autonomic.ext.event.state.wifistatus.WellKnownWifiStatusTransitionTrigger", WellKnownWifiStatusTransitionTrigger_name, WellKnownWifiStatusTransitionTrigger_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/event/state/wifi_status.proto", fileDescriptor_330a863691c5dd04)
}

var fileDescriptor_330a863691c5dd04 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdb, 0x4e, 0xdb, 0x4a,
	0x14, 0x86, 0x09, 0x67, 0x16, 0x07, 0x0d, 0xc3, 0x86, 0x84, 0xf3, 0xde, 0x6c, 0x8a, 0x5a, 0x2a,
	0xe2, 0x8b, 0x3e, 0x81, 0xf1, 0xac, 0xa0, 0x11, 0xc6, 0x0e, 0xf6, 0x38, 0xa1, 0xbd, 0x19, 0x85,
	0xc8, 0x50, 0x4b, 0xc4, 0xae, 0x12, 0x53, 0x78, 0x9e, 0x5e, 0xf0, 0x24, 0x7d, 0xa8, 0x5e, 0x56,
	0xe3, 0x43, 0xe2, 0x28, 0x8e, 0xda, 0x9b, 0xc8, 0x33, 0xeb, 0x9f, 0x75, 0xfc, 0xb2, 0xe0, 0x63,
	0xe7, 0x39, 0x8e, 0xc2, 0xa8, 0x17, 0x74, 0x35, 0xff, 0x35, 0xd6, 0xfc, 0xef, 0x7e, 0x18, 0x6b,
	0x83, 0xb8, 0x13, 0xfb, 0xda, 0x4b, 0xf0, 0x10, 0x48, 0xf5, 0xf9, 0x3c, 0xa8, 0x7f, 0xeb, 0x47,
	0x71, 0x44, 0x4f, 0x87, 0xe2, 0xba, 0xff, 0x1a, 0xd7, 0x13, 0x71, 0x3d, 0x11, 0xd7, 0x95, 0x38,
	0xd5, 0x9e, 0xbf, 0xcd, 0x43, 0xad, 0xed, 0x3f, 0x3d, 0x5d, 0x87, 0xd1, 0x4b, 0xd8, 0x0e, 0x1e,
	0x02, 0x37, 0xb9, 0x57, 0xbf, 0x3e, 0xdd, 0x84, 0x75, 0xcf, 0xba, 0xb6, 0xec, 0xb6, 0x25, 0x5d,
	0xa1, 0x0b, 0x24, 0x33, 0x74, 0x1d, 0x56, 0x1c, 0xbc, 0xf5, 0xd0, 0x15, 0xc8, 0x48, 0x85, 0xee,
	0x43, 0x35, 0x3b, 0x4a, 0x86, 0x26, 0x6f, 0xa1, 0xf3, 0x59, 0xde, 0x7a, 0xe8, 0x21, 0x23, 0xb3,
	0xf4, 0x5f, 0x38, 0x98, 0x30, 0x72, 0x4b, 0x36, 0x1d, 0xfb, 0xca, 0x41, 0xd7, 0x25, 0x73, 0xf4,
	0x08, 0xf6, 0x26, 0x14, 0x82, 0xdf, 0x20, 0x93, 0xb6, 0x27, 0xc8, 0x3c, 0x3d, 0x80, 0xda, 0x84,
	0xbd, 0xa1, 0x73, 0xd3, 0x73, 0x90, 0x2c, 0x50, 0x0a, 0x1b, 0xb9, 0x35, 0x8b, 0xb9, 0x48, 0x4f,
	0xe0, 0x28, 0xbf, 0x6b, 0xe9, 0x26, 0x67, 0xba, 0xe0, 0xb6, 0x35, 0x16, 0x75, 0xa9, 0x18, 0xb5,
	0xa0, 0xc9, 0xfd, 0x2e, 0xab, 0xa2, 0xda, 0xbc, 0xc1, 0x93, 0x9a, 0x3d, 0x77, 0xec, 0xf1, 0x0a,
	0xad, 0xc2, 0x96, 0xeb, 0x19, 0x06, 0x22, 0x53, 0x59, 0x5a, 0x92, 0x61, 0x8b, 0x1b, 0x48, 0xaa,
	0x2a, 0xb2, 0x8d, 0x37, 0xd2, 0x30, 0x6d, 0x8f, 0x95, 0xd7, 0x0b, 0xf4, 0x1d, 0xfc, 0x37, 0xd2,
	0x08, 0x47, 0xb7, 0x5c, 0x73, 0x32, 0xc1, 0x55, 0xd5, 0xb8, 0x54, 0x32, 0xa5, 0x84, 0x9a, 0x6a,
	0xcc, 0x84, 0x22, 0x2f, 0x60, 0x8d, 0xae, 0xc2, 0x52, 0x92, 0xa3, 0xeb, 0x92, 0x4d, 0x75, 0xc8,
	0x2d, 0x54, 0x1d, 0xf0, 0xae, 0xc9, 0x1d, 0x64, 0x64, 0x2b, 0x71, 0xa2, 0x5b, 0x06, 0x9a, 0x25,
	0x49, 0x6c, 0xab, 0x49, 0x67, 0x56, 0x64, 0x64, 0xe7, 0xfc, 0xe7, 0x02, 0x9c, 0x96, 0x80, 0x22,
	0xfa, 0x9d, 0x70, 0x10, 0xc4, 0x41, 0x14, 0x8a, 0x7e, 0xf0, 0xf8, 0xe8, 0xf7, 0x55, 0x77, 0x73,
	0x68, 0x92, 0x0a, 0x79, 0xe2, 0x5b, 0x38, 0xfc, 0xea, 0x0a, 0x1d, 0x32, 0x43, 0x8f, 0x61, 0x7f,
	0x7c, 0x6a, 0xb2, 0x61, 0x3b, 0xc3, 0x86, 0x91, 0x0a, 0x3d, 0x84, 0xdd, 0xb4, 0xa9, 0xd2, 0xb0,
	0x2d, 0x0b, 0x0d, 0x81, 0x4c, 0x0a, 0x3b, 0x6d, 0x1a, 0x99, 0x2d, 0x45, 0x4e, 0x15, 0x88, 0x8c,
	0xcc, 0xa9, 0x06, 0x97, 0x02, 0x65, 0x7b, 0x42, 0xe2, 0x5d, 0x3a, 0xb5, 0x71, 0xae, 0xb2, 0x1c,
	0x46, 0x93, 0x5c, 0x98, 0xc2, 0x87, 0x2b, 0x74, 0x47, 0x24, 0x8c, 0x1d, 0xc2, 0xee, 0x14, 0x7e,
	0x90, 0x91, 0xa5, 0x22, 0xf6, 0xc5, 0xe7, 0x39, 0x34, 0xa4, 0xaa, 0x18, 0x2a, 0x02, 0x96, 0x7b,
	0x5e, 0xa6, 0xbb, 0xb0, 0x3d, 0x66, 0x18, 0xbe, 0x59, 0xa1, 0x3b, 0x40, 0x8b, 0xa6, 0x2c, 0x1a,
	0xa8, 0x68, 0x25, 0xd8, 0x8d, 0x5e, 0xae, 0xd2, 0xff, 0xe1, 0xb8, 0x1c, 0xba, 0x91, 0x68, 0x8d,
	0x9e, 0xc1, 0x49, 0x09, 0xd6, 0x12, 0x5b, 0x68, 0x09, 0xe9, 0xa0, 0x81, 0xbc, 0x85, 0x8c, 0xd4,
	0xe8, 0x05, 0x7c, 0x98, 0x00, 0xcf, 0x6e, 0xe4, 0x72, 0xa6, 0x0b, 0xbd, 0xe0, 0x76, 0x9d, 0x9e,
	0xc3, 0xd9, 0x9f, 0xe4, 0x59, 0x25, 0x1b, 0x14, 0x60, 0x31, 0xfb, 0xde, 0x52, 0x1d, 0x4a, 0x38,
	0x4d, 0x5f, 0x0c, 0x27, 0xf7, 0x0f, 0xdd, 0x83, 0x9d, 0x31, 0x66, 0x47, 0xcb, 0x68, 0x5b, 0xfd,
	0x03, 0xcb, 0x6c, 0xb2, 0xe1, 0x99, 0x0d, 0x9e, 0x62, 0x7c, 0xf9, 0x56, 0x81, 0xf7, 0xdd, 0xa8,
	0x57, 0xff, 0x9b, 0xe5, 0x78, 0x79, 0x38, 0x6d, 0x33, 0x36, 0xd5, 0x86, 0x6d, 0x56, 0xbe, 0x34,
	0x1f, 0x83, 0xf8, 0xeb, 0xf3, 0x7d, 0xbd, 0x1b, 0xf5, 0xb4, 0xa1, 0xc7, 0x8b, 0x4e, 0xa0, 0xd6,
	0xb3, 0xdf, 0x0f, 0x3b, 0x4f, 0x17, 0xc9, 0x2e, 0x1e, 0x68, 0x83, 0x7e, 0x57, 0xeb, 0x75, 0x82,
	0x50, 0x4b, 0xce, 0xda, 0xd4, 0x3d, 0xfe, 0xab, 0x52, 0xf9, 0x31, 0x3b, 0xa7, 0x7b, 0xe2, 0x7e,
	0x31, 0x51, 0x7e, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x08, 0x01, 0x22, 0xce, 0xf4, 0x05, 0x00,
	0x00,
}
