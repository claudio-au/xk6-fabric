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
// source: autonomic/ext/event/state/wifi_config.proto

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

type WellKnownWifiConfigState int32

const (
	WellKnownWifiConfigState_UNKNOWN_STATE WellKnownWifiConfigState = 0
	// Request was received by cloud through API
	// Emitted by: cloud
	WellKnownWifiConfigState_REQUESTED WellKnownWifiConfigState = 1
	// Request was queued in the cloud, waiting for device
	// to connect.
	// Emitted by: cloud
	// Triggered by: REQUEST_QUEUED_FOR_DELIVERY
	WellKnownWifiConfigState_REQUEST_DELIVERY_QUEUED WellKnownWifiConfigState = 2
	// Cloud attempted delivery of request to device
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_TRANSLATION_SUCCEEDED
	WellKnownWifiConfigState_REQUEST_DELIVERY_IN_PROGRESS WellKnownWifiConfigState = 3
	// Delivery of request attempt was not acknowledged by device in a timely fashion
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_TIMEOUT_EXCEEDED
	WellKnownWifiConfigState_REQUEST_DELIVERY_TIMED_OUT WellKnownWifiConfigState = 4
	// Delivery of request attempt failed due to device disconnect
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_FAILED
	WellKnownWifiConfigState_REQUEST_DELIVERY_FAILURE WellKnownWifiConfigState = 5
	// Request is acknowledged by device, queued, and the device will begin
	// processing the request when the device is ready.
	// Emitted by: device
	// Triggered by: REQUEST_QUEUED_ON_DEVICE
	WellKnownWifiConfigState_REQUEST_QUEUED WellKnownWifiConfigState = 6
	// Request is being decrypted and validated by device
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_STARTED
	WellKnownWifiConfigState_REQUEST_VALIDATION_IN_PROGRESS WellKnownWifiConfigState = 7
	// Request failed validation or decryption
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_FAILED
	WellKnownWifiConfigState_REQUEST_VALIDATION_FAILURE WellKnownWifiConfigState = 8
	// Wifi configuration in progress
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED
	WellKnownWifiConfigState_WIFI_CONFIG_IN_PROGRESS WellKnownWifiConfigState = 9
	// The device has successfully completed the wifi config command in question.
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED
	WellKnownWifiConfigState_SUCCEEDED_ON_DEVICE WellKnownWifiConfigState = 23
	// Oem cloud attempted delivery of request
	// Emitted by: cloud
	// Triggered by: DEVICE_CONNECTED_TO_CLOUD
	WellKnownWifiConfigState_OEM_CLOUD_DELIVERY_IN_PROGRESS WellKnownWifiConfigState = 10
	// Cloud attempted translation of request
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_DELIVERY_SUCCEEDED
	WellKnownWifiConfigState_OEM_CLOUD_TRANSLATION_IN_PROGRESS WellKnownWifiConfigState = 11
	// The cloud is in the process of validating the event emitted by the device indicating that the
	// command has succeeded.
	// Triggered by: SUCCEEDED_ON_DEVICE_EVENT_RECEIVED
	WellKnownWifiConfigState_CLOUD_VALIDATION_IN_PROGRESS WellKnownWifiConfigState = 24
	// Cloud failed to validate data provided by the device
	// Emitted by: cloud
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED
	WellKnownWifiConfigState_CLOUD_VALIDATION_FAILURE WellKnownWifiConfigState = 12
	// Wifi configuration succeeded
	// added to the event containing this state.
	// This state is terminal
	// Triggered by: WIFI_CONFIG_SUCCEEDED
	WellKnownWifiConfigState_SUCCESS WellKnownWifiConfigState = 17
	// Wifi configuration failed
	// this state.
	// Emitted by: device
	// Triggered by: WIFI_CONFIG_FAILED
	WellKnownWifiConfigState_FAILURE WellKnownWifiConfigState = 18
	// Wifi configuration command expired. This event can be emitted by
	// the device or the cloud. In either case, it will cause
	// the command to be removed from the set of in-progress commands
	// on the cloud. However the cloud does not send this event to
	// the device. This state is terminal.
	// Triggered by: EXPIRATION_EXCEEDED
	WellKnownWifiConfigState_EXPIRED WellKnownWifiConfigState = 19
	// Command is being removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_ACKNOWLEDGED
	WellKnownWifiConfigState_CANCELLATION_IN_PROGRESS WellKnownWifiConfigState = 21
	// Command was successfully removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_FULFILLED
	WellKnownWifiConfigState_CANCELLED WellKnownWifiConfigState = 22
)

var WellKnownWifiConfigState_name = map[int32]string{
	0:  "UNKNOWN_STATE",
	1:  "REQUESTED",
	2:  "REQUEST_DELIVERY_QUEUED",
	3:  "REQUEST_DELIVERY_IN_PROGRESS",
	4:  "REQUEST_DELIVERY_TIMED_OUT",
	5:  "REQUEST_DELIVERY_FAILURE",
	6:  "REQUEST_QUEUED",
	7:  "REQUEST_VALIDATION_IN_PROGRESS",
	8:  "REQUEST_VALIDATION_FAILURE",
	9:  "WIFI_CONFIG_IN_PROGRESS",
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

var WellKnownWifiConfigState_value = map[string]int32{
	"UNKNOWN_STATE":                     0,
	"REQUESTED":                         1,
	"REQUEST_DELIVERY_QUEUED":           2,
	"REQUEST_DELIVERY_IN_PROGRESS":      3,
	"REQUEST_DELIVERY_TIMED_OUT":        4,
	"REQUEST_DELIVERY_FAILURE":          5,
	"REQUEST_QUEUED":                    6,
	"REQUEST_VALIDATION_IN_PROGRESS":    7,
	"REQUEST_VALIDATION_FAILURE":        8,
	"WIFI_CONFIG_IN_PROGRESS":           9,
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

func (x WellKnownWifiConfigState) String() string {
	return proto.EnumName(WellKnownWifiConfigState_name, int32(x))
}

func (WellKnownWifiConfigState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9bff5e8ae388ed8c, []int{0}
}

// Wifi configuration state transition triggers. See WellKnownRemoteStarState enum
// for documentation on when these triggers occur.
type WellKnownWifiConfigTransitionTrigger int32

const (
	WellKnownWifiConfigTransitionTrigger_UNKNOWN_TRANSITION_TRIGGER                WellKnownWifiConfigTransitionTrigger = 0
	WellKnownWifiConfigTransitionTrigger_REQUEST_QUEUED_FOR_DELIVERY               WellKnownWifiConfigTransitionTrigger = 1
	WellKnownWifiConfigTransitionTrigger_DEVICE_CONNECTED_TO_CLOUD                 WellKnownWifiConfigTransitionTrigger = 2
	WellKnownWifiConfigTransitionTrigger_REQUEST_DELIVERY_FAILED                   WellKnownWifiConfigTransitionTrigger = 3
	WellKnownWifiConfigTransitionTrigger_REQUEST_DELIVERY_TIMEOUT_EXCEEDED         WellKnownWifiConfigTransitionTrigger = 4
	WellKnownWifiConfigTransitionTrigger_REQUEST_QUEUED_ON_DEVICE                  WellKnownWifiConfigTransitionTrigger = 5
	WellKnownWifiConfigTransitionTrigger_REQUEST_VALIDATION_STARTED                WellKnownWifiConfigTransitionTrigger = 6
	WellKnownWifiConfigTransitionTrigger_REQUEST_VALIDATION_FAILED                 WellKnownWifiConfigTransitionTrigger = 7
	WellKnownWifiConfigTransitionTrigger_REQUEST_VALIDATION_SUCCEEDED              WellKnownWifiConfigTransitionTrigger = 23
	WellKnownWifiConfigTransitionTrigger_WIFI_CONFIG_STARTED                       WellKnownWifiConfigTransitionTrigger = 8
	WellKnownWifiConfigTransitionTrigger_WIFI_CONFIG_SUCCEEDED                     WellKnownWifiConfigTransitionTrigger = 9
	WellKnownWifiConfigTransitionTrigger_WIFI_CONFIG_FAILED                        WellKnownWifiConfigTransitionTrigger = 10
	WellKnownWifiConfigTransitionTrigger_OEM_CLOUD_DELIVERY_SUCCEEDED              WellKnownWifiConfigTransitionTrigger = 11
	WellKnownWifiConfigTransitionTrigger_OEM_CLOUD_TRANSLATION_SUCCEEDED           WellKnownWifiConfigTransitionTrigger = 12
	WellKnownWifiConfigTransitionTrigger_SUCCEEDED_ON_DEVICE_EVENT_RECEIVED        WellKnownWifiConfigTransitionTrigger = 24
	WellKnownWifiConfigTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED WellKnownWifiConfigTransitionTrigger = 13
	WellKnownWifiConfigTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED    WellKnownWifiConfigTransitionTrigger = 14
	WellKnownWifiConfigTransitionTrigger_FAILED                                    WellKnownWifiConfigTransitionTrigger = 19
	WellKnownWifiConfigTransitionTrigger_EXPIRATION_EXCEEDED                       WellKnownWifiConfigTransitionTrigger = 20
	WellKnownWifiConfigTransitionTrigger_CANCELLATION_REQUESTED                    WellKnownWifiConfigTransitionTrigger = 21
	WellKnownWifiConfigTransitionTrigger_CANCELLATION_REQUEST_FULFILLED            WellKnownWifiConfigTransitionTrigger = 22
)

var WellKnownWifiConfigTransitionTrigger_name = map[int32]string{
	0:  "UNKNOWN_TRANSITION_TRIGGER",
	1:  "REQUEST_QUEUED_FOR_DELIVERY",
	2:  "DEVICE_CONNECTED_TO_CLOUD",
	3:  "REQUEST_DELIVERY_FAILED",
	4:  "REQUEST_DELIVERY_TIMEOUT_EXCEEDED",
	5:  "REQUEST_QUEUED_ON_DEVICE",
	6:  "REQUEST_VALIDATION_STARTED",
	7:  "REQUEST_VALIDATION_FAILED",
	23: "REQUEST_VALIDATION_SUCCEEDED",
	8:  "WIFI_CONFIG_STARTED",
	9:  "WIFI_CONFIG_SUCCEEDED",
	10: "WIFI_CONFIG_FAILED",
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

var WellKnownWifiConfigTransitionTrigger_value = map[string]int32{
	"UNKNOWN_TRANSITION_TRIGGER":                0,
	"REQUEST_QUEUED_FOR_DELIVERY":               1,
	"DEVICE_CONNECTED_TO_CLOUD":                 2,
	"REQUEST_DELIVERY_FAILED":                   3,
	"REQUEST_DELIVERY_TIMEOUT_EXCEEDED":         4,
	"REQUEST_QUEUED_ON_DEVICE":                  5,
	"REQUEST_VALIDATION_STARTED":                6,
	"REQUEST_VALIDATION_FAILED":                 7,
	"REQUEST_VALIDATION_SUCCEEDED":              23,
	"WIFI_CONFIG_STARTED":                       8,
	"WIFI_CONFIG_SUCCEEDED":                     9,
	"WIFI_CONFIG_FAILED":                        10,
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

func (x WellKnownWifiConfigTransitionTrigger) String() string {
	return proto.EnumName(WellKnownWifiConfigTransitionTrigger_name, int32(x))
}

func (WellKnownWifiConfigTransitionTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9bff5e8ae388ed8c, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.event.state.wificonfig.WellKnownWifiConfigState", WellKnownWifiConfigState_name, WellKnownWifiConfigState_value)
	proto.RegisterEnum("autonomic.ext.event.state.wificonfig.WellKnownWifiConfigTransitionTrigger", WellKnownWifiConfigTransitionTrigger_name, WellKnownWifiConfigTransitionTrigger_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/event/state/wifi_config.proto", fileDescriptor_9bff5e8ae388ed8c)
}

var fileDescriptor_9bff5e8ae388ed8c = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdb, 0x4e, 0xdb, 0x4a,
	0x14, 0x86, 0x09, 0x67, 0x16, 0x07, 0x0d, 0xc3, 0x86, 0x84, 0xf3, 0xde, 0x6c, 0x8a, 0x5a, 0x2a,
	0xe2, 0x8b, 0x3e, 0x81, 0xf1, 0xac, 0x44, 0x23, 0x82, 0x1d, 0xec, 0x71, 0x42, 0x7b, 0x33, 0x0a,
	0x91, 0x49, 0x2d, 0x11, 0xbb, 0x4a, 0x4c, 0xe1, 0x79, 0x7a, 0xc1, 0x93, 0xf4, 0xa1, 0x7a, 0x59,
	0x8d, 0x0f, 0x89, 0xa3, 0x38, 0x6a, 0xef, 0x32, 0x5e, 0xff, 0x3a, 0x7f, 0x59, 0xf0, 0xb1, 0xf3,
	0x1c, 0x85, 0x41, 0xd8, 0xf7, 0xbb, 0x9a, 0xf7, 0x1a, 0x69, 0xde, 0x77, 0x2f, 0x88, 0xb4, 0x61,
	0xd4, 0x89, 0x3c, 0xed, 0xc5, 0x7f, 0xf4, 0x65, 0x37, 0x0c, 0x1e, 0xfd, 0x5e, 0xf5, 0xdb, 0x20,
	0x8c, 0x42, 0x7a, 0x3e, 0x12, 0x57, 0xbd, 0xd7, 0xa8, 0x1a, 0x8b, 0xab, 0xb1, 0xb8, 0xaa, 0xc4,
	0x89, 0xf6, 0xf2, 0x6d, 0x11, 0x2a, 0x6d, 0xef, 0xe9, 0xe9, 0x26, 0x08, 0x5f, 0x82, 0xb6, 0xff,
	0xe8, 0x1b, 0xf1, 0x77, 0x47, 0xe9, 0xe8, 0x36, 0x6c, 0xba, 0xe6, 0x8d, 0x69, 0xb5, 0x4d, 0xe9,
	0x08, 0x5d, 0x20, 0x99, 0xa3, 0x9b, 0xb0, 0x66, 0xe3, 0x9d, 0x8b, 0x8e, 0x40, 0x46, 0x4a, 0xf4,
	0x10, 0xca, 0xe9, 0x53, 0x32, 0x6c, 0xf0, 0x16, 0xda, 0x9f, 0xe5, 0x9d, 0x8b, 0x2e, 0x32, 0x32,
	0x4f, 0xff, 0x85, 0xa3, 0x29, 0x23, 0x37, 0x65, 0xd3, 0xb6, 0xea, 0x36, 0x3a, 0x0e, 0x59, 0xa0,
	0x27, 0x70, 0x30, 0xa5, 0x10, 0xfc, 0x16, 0x99, 0xb4, 0x5c, 0x41, 0x16, 0xe9, 0x11, 0x54, 0xa6,
	0xec, 0x35, 0x9d, 0x37, 0x5c, 0x1b, 0xc9, 0x12, 0xa5, 0xb0, 0x95, 0x59, 0xd3, 0x9c, 0xcb, 0xf4,
	0x0c, 0x4e, 0xb2, 0x6f, 0x2d, 0xbd, 0xc1, 0x99, 0x2e, 0xb8, 0x65, 0x4e, 0x64, 0x5d, 0xc9, 0x67,
	0xcd, 0x69, 0xb2, 0xb8, 0xab, 0xaa, 0xa9, 0x36, 0xaf, 0x71, 0x69, 0x58, 0x66, 0x8d, 0xd7, 0x27,
	0x9c, 0xd7, 0x68, 0x19, 0x76, 0x1c, 0xd7, 0x30, 0x10, 0x99, 0xaa, 0xd2, 0x94, 0x0c, 0x5b, 0xdc,
	0x40, 0x52, 0x56, 0x99, 0x2d, 0xbc, 0x95, 0x46, 0xc3, 0x72, 0x59, 0x71, 0xbf, 0x40, 0xdf, 0xc1,
	0x7f, 0x63, 0x8d, 0xb0, 0x75, 0xd3, 0x69, 0x4c, 0x17, 0xb8, 0xae, 0x06, 0x97, 0x48, 0x66, 0xb4,
	0x50, 0x51, 0x83, 0x99, 0x52, 0x64, 0x0d, 0x6c, 0xd0, 0x75, 0x58, 0x89, 0x6b, 0x74, 0x1c, 0xb2,
	0xad, 0x1e, 0x99, 0x85, 0xaa, 0x07, 0xde, 0x37, 0xb9, 0x8d, 0x8c, 0xec, 0xc4, 0x41, 0x74, 0xd3,
	0xc0, 0x46, 0x41, 0x11, 0xbb, 0x6a, 0xd3, 0xa9, 0x15, 0x19, 0xd9, 0xbb, 0xfc, 0xb9, 0x04, 0xe7,
	0x05, 0xa0, 0x88, 0x41, 0x27, 0x18, 0xfa, 0x91, 0x1f, 0x06, 0x62, 0xe0, 0xf7, 0x7a, 0xde, 0x40,
	0x4d, 0x37, 0x83, 0x26, 0xee, 0x90, 0xc7, 0xb1, 0x85, 0xcd, 0xeb, 0x75, 0xb4, 0xc9, 0x1c, 0x3d,
	0x85, 0xc3, 0xc9, 0xad, 0xc9, 0x9a, 0x65, 0x8f, 0x06, 0x46, 0x4a, 0xf4, 0x18, 0xf6, 0x93, 0xa1,
	0xaa, 0x05, 0x98, 0x68, 0x08, 0x64, 0x52, 0x58, 0xc9, 0xd0, 0xc8, 0x7c, 0x21, 0x72, 0xaa, 0x41,
	0x64, 0x64, 0x41, 0x0d, 0xb8, 0x10, 0x28, 0xcb, 0x15, 0x12, 0xef, 0x93, 0xad, 0x4d, 0x72, 0x95,
	0xd6, 0x30, 0xde, 0xe4, 0xd2, 0x0c, 0x3e, 0x1c, 0xa1, 0xdb, 0x22, 0x66, 0xec, 0x18, 0xf6, 0x67,
	0xf0, 0x83, 0x8c, 0xac, 0xe4, 0xb1, 0xcf, 0xbb, 0x67, 0xd0, 0x90, 0xb2, 0x62, 0x28, 0x0f, 0x58,
	0x16, 0x79, 0x95, 0xee, 0xc3, 0xee, 0x84, 0x61, 0xe4, 0xb3, 0x46, 0xf7, 0x80, 0xe6, 0x4d, 0x69,
	0x36, 0x50, 0xd9, 0x0a, 0xb0, 0x1b, 0x7b, 0xae, 0xd3, 0xff, 0xe1, 0xb4, 0x18, 0xba, 0xb1, 0x68,
	0x83, 0x5e, 0xc0, 0x59, 0x01, 0xd6, 0x12, 0x5b, 0x68, 0x0a, 0x69, 0xa3, 0x81, 0xbc, 0x85, 0x8c,
	0x54, 0xe8, 0x15, 0x7c, 0x98, 0x02, 0xcf, 0xaa, 0x65, 0x72, 0xa6, 0x0b, 0x3d, 0x17, 0x76, 0x93,
	0x5e, 0xc2, 0xc5, 0x9f, 0xe4, 0x69, 0x27, 0x5b, 0x14, 0x60, 0x39, 0xfd, 0xbd, 0xa3, 0x26, 0x14,
	0x73, 0x9a, 0x78, 0x8c, 0x36, 0xf7, 0x0f, 0x3d, 0x80, 0xbd, 0x09, 0x66, 0xc7, 0xc7, 0x68, 0x57,
	0xfd, 0x03, 0x8b, 0x6c, 0xb2, 0xe6, 0x36, 0x6a, 0x3c, 0xc1, 0xf8, 0xfa, 0xad, 0x04, 0xef, 0xbb,
	0x61, 0xbf, 0xfa, 0x37, 0xc7, 0xf1, 0xfa, 0x78, 0xd6, 0x65, 0x6c, 0xaa, 0x0b, 0xdb, 0x2c, 0x7d,
	0x69, 0xf6, 0xfc, 0xe8, 0xeb, 0xf3, 0x43, 0xb5, 0x1b, 0xf6, 0xb5, 0x51, 0xc4, 0xab, 0x8e, 0xaf,
	0xce, 0xb3, 0x37, 0x08, 0x3a, 0x4f, 0x57, 0xf1, 0x2d, 0x1e, 0x6a, 0xc3, 0x41, 0x57, 0xeb, 0x77,
	0xfc, 0x40, 0x8b, 0xdf, 0xda, 0xcc, 0x3b, 0xfe, 0xab, 0x54, 0xfa, 0x31, 0xbf, 0xa0, 0xbb, 0xe2,
	0x61, 0x39, 0x56, 0x7e, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xba, 0x34, 0xb3, 0xd6, 0xf4, 0x05,
	0x00, 0x00,
}
