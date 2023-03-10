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
// source: autonomic/ext/event/state/hmi_display.proto

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

type WellKnownHMIDisplayState int32

const (
	WellKnownHMIDisplayState_UNKNOWN_STATE WellKnownHMIDisplayState = 0
	// Request was received by cloud through API
	// Emitted by: cloud
	WellKnownHMIDisplayState_REQUESTED WellKnownHMIDisplayState = 1
	// Request was queued in the cloud, waiting for device
	// to connect.
	// Emitted by: cloud
	// Triggered by: REQUEST_QUEUED_FOR_DELIVERY
	WellKnownHMIDisplayState_REQUEST_DELIVERY_QUEUED WellKnownHMIDisplayState = 2
	// Cloud attempted delivery of request to device
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_TRANSLATION_SUCCEEDED
	WellKnownHMIDisplayState_REQUEST_DELIVERY_IN_PROGRESS WellKnownHMIDisplayState = 3
	// Delivery of request attempt was not acknowledged by device in a timely fashion
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_TIMEOUT_EXCEEDED
	WellKnownHMIDisplayState_REQUEST_DELIVERY_TIMED_OUT WellKnownHMIDisplayState = 4
	// Delivery of request attempt failed due to device disconnect
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_FAILED
	WellKnownHMIDisplayState_REQUEST_DELIVERY_FAILURE WellKnownHMIDisplayState = 5
	// Request is acknowledged by device, queued, and the device will begin
	// processing the request when the device is ready.
	// Emitted by: device
	// Triggered by: REQUEST_QUEUED_ON_DEVICE
	WellKnownHMIDisplayState_REQUEST_QUEUED WellKnownHMIDisplayState = 6
	// Request is being decrypted and validated by device
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_STARTED
	WellKnownHMIDisplayState_REQUEST_VALIDATION_IN_PROGRESS WellKnownHMIDisplayState = 7
	// Request failed validation or decryption
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_FAILED
	WellKnownHMIDisplayState_REQUEST_VALIDATION_FAILURE WellKnownHMIDisplayState = 8
	// HMIDisplay is in the process of being reset
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED
	WellKnownHMIDisplayState_HMI_DISPLAY_IN_PROGRESS WellKnownHMIDisplayState = 9
	// The device has successfully completed the actuation in question.
	// Triggered by: HMI_DISPLAY_SUCCEEDED
	WellKnownHMIDisplayState_SUCCEEDED_ON_DEVICE WellKnownHMIDisplayState = 23
	// Oem cloud attempted delivery of request
	// Emitted by: cloud
	// Triggered by: DEVICE_CONNECTED_TO_CLOUD
	WellKnownHMIDisplayState_OEM_CLOUD_DELIVERY_IN_PROGRESS WellKnownHMIDisplayState = 10
	// Cloud attempted translation of request
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_DELIVERY_SUCCEEDED
	WellKnownHMIDisplayState_OEM_CLOUD_TRANSLATION_IN_PROGRESS WellKnownHMIDisplayState = 11
	// The cloud is in the process of validating the event emitted by the device indicating that the
	// command has succeeded.
	// Triggered by: SUCCEEDED_ON_DEVICE_EVENT_RECEIVED
	WellKnownHMIDisplayState_CLOUD_VALIDATION_IN_PROGRESS WellKnownHMIDisplayState = 24
	// Cloud failed to validate data provided by the device
	// Emitted by: cloud
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED
	WellKnownHMIDisplayState_CLOUD_VALIDATION_FAILURE WellKnownHMIDisplayState = 12
	// HMIDisplay upload succeeded, and the URI(s) of the archive(s) has been
	// added to the event containing this state.
	// This state is terminal
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED
	WellKnownHMIDisplayState_SUCCESS WellKnownHMIDisplayState = 17
	// HMIDisplay failed. All intermediate failure states should transition to
	// this state.
	// Emitted by: device
	// Triggered by: HMI_DISPLAY_FAILED
	WellKnownHMIDisplayState_FAILURE WellKnownHMIDisplayState = 18
	// HMIDisplay command expired. This event can be emitted by
	// the device or the cloud. In either case, it will cause
	// the command to be removed from the set of in-progress commands
	// on the cloud. However the cloud does not send this event to
	// the device. This state is terminal.
	// Triggered by: EXPIRATION_EXCEEDED
	WellKnownHMIDisplayState_EXPIRED WellKnownHMIDisplayState = 19
	// Command is being removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_ACKNOWLEDGED
	WellKnownHMIDisplayState_CANCELLATION_IN_PROGRESS WellKnownHMIDisplayState = 21
	// Command was successfully removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_FULFILLED
	WellKnownHMIDisplayState_CANCELLED WellKnownHMIDisplayState = 22
)

var WellKnownHMIDisplayState_name = map[int32]string{
	0:  "UNKNOWN_STATE",
	1:  "REQUESTED",
	2:  "REQUEST_DELIVERY_QUEUED",
	3:  "REQUEST_DELIVERY_IN_PROGRESS",
	4:  "REQUEST_DELIVERY_TIMED_OUT",
	5:  "REQUEST_DELIVERY_FAILURE",
	6:  "REQUEST_QUEUED",
	7:  "REQUEST_VALIDATION_IN_PROGRESS",
	8:  "REQUEST_VALIDATION_FAILURE",
	9:  "HMI_DISPLAY_IN_PROGRESS",
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

var WellKnownHMIDisplayState_value = map[string]int32{
	"UNKNOWN_STATE":                     0,
	"REQUESTED":                         1,
	"REQUEST_DELIVERY_QUEUED":           2,
	"REQUEST_DELIVERY_IN_PROGRESS":      3,
	"REQUEST_DELIVERY_TIMED_OUT":        4,
	"REQUEST_DELIVERY_FAILURE":          5,
	"REQUEST_QUEUED":                    6,
	"REQUEST_VALIDATION_IN_PROGRESS":    7,
	"REQUEST_VALIDATION_FAILURE":        8,
	"HMI_DISPLAY_IN_PROGRESS":           9,
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

func (x WellKnownHMIDisplayState) String() string {
	return proto.EnumName(WellKnownHMIDisplayState_name, int32(x))
}

func (WellKnownHMIDisplayState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa47896cb7a40671, []int{0}
}

// HMIDisplay state transition triggers. See WellKnownHMIDisplayState enum
// for documentation on when these triggers occur.
type WellKnownHMIDisplayTransitionTrigger int32

const (
	WellKnownHMIDisplayTransitionTrigger_UNKNOWN_TRANSITION_TRIGGER                WellKnownHMIDisplayTransitionTrigger = 0
	WellKnownHMIDisplayTransitionTrigger_REQUEST_QUEUED_FOR_DELIVERY               WellKnownHMIDisplayTransitionTrigger = 1
	WellKnownHMIDisplayTransitionTrigger_DEVICE_CONNECTED_TO_CLOUD                 WellKnownHMIDisplayTransitionTrigger = 2
	WellKnownHMIDisplayTransitionTrigger_REQUEST_DELIVERY_FAILED                   WellKnownHMIDisplayTransitionTrigger = 3
	WellKnownHMIDisplayTransitionTrigger_REQUEST_DELIVERY_TIMEOUT_EXCEEDED         WellKnownHMIDisplayTransitionTrigger = 4
	WellKnownHMIDisplayTransitionTrigger_REQUEST_QUEUED_ON_DEVICE                  WellKnownHMIDisplayTransitionTrigger = 5
	WellKnownHMIDisplayTransitionTrigger_REQUEST_VALIDATION_STARTED                WellKnownHMIDisplayTransitionTrigger = 6
	WellKnownHMIDisplayTransitionTrigger_REQUEST_VALIDATION_FAILED                 WellKnownHMIDisplayTransitionTrigger = 7
	WellKnownHMIDisplayTransitionTrigger_REQUEST_VALIDATION_SUCCEEDED              WellKnownHMIDisplayTransitionTrigger = 23
	WellKnownHMIDisplayTransitionTrigger_HMI_DISPLAY_STARTED                       WellKnownHMIDisplayTransitionTrigger = 8
	WellKnownHMIDisplayTransitionTrigger_HMI_DISPLAY_SUCCEEDED                     WellKnownHMIDisplayTransitionTrigger = 9
	WellKnownHMIDisplayTransitionTrigger_HMI_DISPLAY_FAILED                        WellKnownHMIDisplayTransitionTrigger = 10
	WellKnownHMIDisplayTransitionTrigger_OEM_CLOUD_DELIVERY_SUCCEEDED              WellKnownHMIDisplayTransitionTrigger = 11
	WellKnownHMIDisplayTransitionTrigger_OEM_CLOUD_TRANSLATION_SUCCEEDED           WellKnownHMIDisplayTransitionTrigger = 12
	WellKnownHMIDisplayTransitionTrigger_SUCCEEDED_ON_DEVICE_EVENT_RECEIVED        WellKnownHMIDisplayTransitionTrigger = 24
	WellKnownHMIDisplayTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED WellKnownHMIDisplayTransitionTrigger = 13
	WellKnownHMIDisplayTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED    WellKnownHMIDisplayTransitionTrigger = 14
	WellKnownHMIDisplayTransitionTrigger_FAILED                                    WellKnownHMIDisplayTransitionTrigger = 19
	WellKnownHMIDisplayTransitionTrigger_EXPIRATION_EXCEEDED                       WellKnownHMIDisplayTransitionTrigger = 20
	WellKnownHMIDisplayTransitionTrigger_CANCELLATION_REQUESTED                    WellKnownHMIDisplayTransitionTrigger = 21
	WellKnownHMIDisplayTransitionTrigger_CANCELLATION_REQUEST_FULFILLED            WellKnownHMIDisplayTransitionTrigger = 22
)

var WellKnownHMIDisplayTransitionTrigger_name = map[int32]string{
	0:  "UNKNOWN_TRANSITION_TRIGGER",
	1:  "REQUEST_QUEUED_FOR_DELIVERY",
	2:  "DEVICE_CONNECTED_TO_CLOUD",
	3:  "REQUEST_DELIVERY_FAILED",
	4:  "REQUEST_DELIVERY_TIMEOUT_EXCEEDED",
	5:  "REQUEST_QUEUED_ON_DEVICE",
	6:  "REQUEST_VALIDATION_STARTED",
	7:  "REQUEST_VALIDATION_FAILED",
	23: "REQUEST_VALIDATION_SUCCEEDED",
	8:  "HMI_DISPLAY_STARTED",
	9:  "HMI_DISPLAY_SUCCEEDED",
	10: "HMI_DISPLAY_FAILED",
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

var WellKnownHMIDisplayTransitionTrigger_value = map[string]int32{
	"UNKNOWN_TRANSITION_TRIGGER":                0,
	"REQUEST_QUEUED_FOR_DELIVERY":               1,
	"DEVICE_CONNECTED_TO_CLOUD":                 2,
	"REQUEST_DELIVERY_FAILED":                   3,
	"REQUEST_DELIVERY_TIMEOUT_EXCEEDED":         4,
	"REQUEST_QUEUED_ON_DEVICE":                  5,
	"REQUEST_VALIDATION_STARTED":                6,
	"REQUEST_VALIDATION_FAILED":                 7,
	"REQUEST_VALIDATION_SUCCEEDED":              23,
	"HMI_DISPLAY_STARTED":                       8,
	"HMI_DISPLAY_SUCCEEDED":                     9,
	"HMI_DISPLAY_FAILED":                        10,
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

func (x WellKnownHMIDisplayTransitionTrigger) String() string {
	return proto.EnumName(WellKnownHMIDisplayTransitionTrigger_name, int32(x))
}

func (WellKnownHMIDisplayTransitionTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa47896cb7a40671, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.event.state.hmidisplay.WellKnownHMIDisplayState", WellKnownHMIDisplayState_name, WellKnownHMIDisplayState_value)
	proto.RegisterEnum("autonomic.ext.event.state.hmidisplay.WellKnownHMIDisplayTransitionTrigger", WellKnownHMIDisplayTransitionTrigger_name, WellKnownHMIDisplayTransitionTrigger_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/event/state/hmi_display.proto", fileDescriptor_fa47896cb7a40671)
}

var fileDescriptor_fa47896cb7a40671 = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0xc7, 0x09, 0xdf, 0x1c, 0x3e, 0x34, 0x0c, 0x0b, 0x09, 0xdf, 0xbb, 0x2c, 0x8b, 0xb6, 0x54,
	0xc4, 0x17, 0x7d, 0x02, 0xe3, 0x39, 0x81, 0x11, 0x8e, 0x1d, 0xec, 0x71, 0xa0, 0xbd, 0x19, 0x85,
	0xd4, 0x02, 0x4b, 0x89, 0x8d, 0x12, 0xd3, 0xd2, 0xd7, 0xe9, 0x05, 0x4f, 0xd2, 0x87, 0xea, 0x65,
	0x35, 0x8e, 0x9d, 0x38, 0x8a, 0xa3, 0xf6, 0x2e, 0xe3, 0xf3, 0x3f, 0xdf, 0xbf, 0x1c, 0x78, 0xdf,
	0x7a, 0x89, 0xa3, 0x30, 0xea, 0x06, 0x6d, 0xcd, 0x7f, 0x8d, 0x35, 0xff, 0x8b, 0x1f, 0xc6, 0x5a,
	0x3f, 0x6e, 0xc5, 0xbe, 0xf6, 0xd4, 0x0d, 0xe4, 0xe7, 0xa0, 0xff, 0xdc, 0x69, 0x7d, 0xab, 0x3e,
	0xf7, 0xa2, 0x38, 0xa2, 0xa7, 0x43, 0x71, 0xd5, 0x7f, 0x8d, 0xab, 0x89, 0xb8, 0x9a, 0x88, 0xab,
	0x4f, 0xdd, 0x20, 0xd5, 0x9e, 0xbf, 0xcd, 0x43, 0xe5, 0xce, 0xef, 0x74, 0x6e, 0xc2, 0xe8, 0x6b,
	0x78, 0x5d, 0xe7, 0x6c, 0xf0, 0xdd, 0x55, 0x3a, 0xba, 0x09, 0xeb, 0x9e, 0x75, 0x63, 0xd9, 0x77,
	0x96, 0x74, 0x85, 0x2e, 0x90, 0xcc, 0xd0, 0x75, 0x58, 0x71, 0xf0, 0xd6, 0x43, 0x57, 0x20, 0x23,
	0x25, 0xba, 0x0f, 0xe5, 0xf4, 0x29, 0x19, 0x9a, 0xbc, 0x89, 0xce, 0x47, 0x79, 0xeb, 0xa1, 0x87,
	0x8c, 0xcc, 0xd2, 0xbf, 0xe1, 0x60, 0xc2, 0xc8, 0x2d, 0xd9, 0x70, 0xec, 0x2b, 0x07, 0x5d, 0x97,
	0xcc, 0xd1, 0x23, 0xd8, 0x9b, 0x50, 0x08, 0x5e, 0x47, 0x26, 0x6d, 0x4f, 0x90, 0x79, 0x7a, 0x00,
	0x95, 0x09, 0x7b, 0x4d, 0xe7, 0xa6, 0xe7, 0x20, 0x59, 0xa0, 0x14, 0x36, 0x32, 0x6b, 0x9a, 0x73,
	0x91, 0x9e, 0xc0, 0x51, 0xf6, 0xad, 0xa9, 0x9b, 0x9c, 0xe9, 0x82, 0xdb, 0xd6, 0x58, 0xd6, 0xa5,
	0x7c, 0xd6, 0x9c, 0x26, 0x8b, 0xbb, 0xac, 0x9a, 0xba, 0xae, 0x73, 0xc9, 0xb8, 0xdb, 0x30, 0xf5,
	0xf1, 0x92, 0x57, 0x68, 0x19, 0xb6, 0x5c, 0xcf, 0x30, 0x10, 0x99, 0xaa, 0xd2, 0x92, 0x0c, 0x9b,
	0xdc, 0x40, 0x52, 0x56, 0x99, 0x6d, 0xac, 0x4b, 0xc3, 0xb4, 0x3d, 0x56, 0xdc, 0x2f, 0xd0, 0xff,
	0xe0, 0x9f, 0x91, 0x46, 0x38, 0xba, 0xe5, 0x9a, 0x93, 0x05, 0xae, 0xaa, 0xc1, 0x0d, 0x24, 0x53,
	0x5a, 0xa8, 0xa8, 0xc1, 0x4c, 0x28, 0xb2, 0x06, 0xd6, 0xe8, 0x2a, 0x2c, 0x25, 0x35, 0xba, 0x2e,
	0xd9, 0x54, 0x8f, 0xcc, 0x42, 0xd5, 0x03, 0xef, 0x1b, 0xdc, 0x41, 0x46, 0xb6, 0x92, 0x20, 0xba,
	0x65, 0xa0, 0x59, 0x50, 0xc4, 0xb6, 0xda, 0x74, 0x6a, 0x45, 0x46, 0x76, 0xce, 0x7f, 0x2c, 0xc0,
	0x69, 0x01, 0x28, 0xa2, 0xd7, 0x0a, 0xfb, 0x41, 0x1c, 0x44, 0xa1, 0xe8, 0x05, 0x8f, 0x8f, 0x7e,
	0x4f, 0x4d, 0x37, 0x83, 0x26, 0xe9, 0x90, 0x27, 0xb1, 0x85, 0xc3, 0xaf, 0xae, 0xd0, 0x21, 0x33,
	0xf4, 0x18, 0xf6, 0xc7, 0xb7, 0x26, 0x6b, 0xb6, 0x33, 0x1c, 0x18, 0x29, 0xd1, 0x43, 0xd8, 0x1d,
	0x0c, 0x55, 0x1a, 0xb6, 0x65, 0xa1, 0x21, 0x90, 0x49, 0x61, 0x0f, 0x86, 0x46, 0x66, 0x0b, 0x91,
	0x53, 0x0d, 0x22, 0x23, 0x73, 0x6a, 0xc0, 0x85, 0x40, 0xd9, 0x9e, 0x90, 0x78, 0x3f, 0xd8, 0xda,
	0x38, 0x57, 0x69, 0x0d, 0xa3, 0x4d, 0x2e, 0x4c, 0xe1, 0xc3, 0x15, 0xba, 0x23, 0x12, 0xc6, 0x0e,
	0x61, 0x77, 0x0a, 0x3f, 0xc8, 0xc8, 0x52, 0x1e, 0xfb, 0xbc, 0x7b, 0x06, 0x0d, 0x29, 0x2b, 0x86,
	0xf2, 0x80, 0x65, 0x91, 0x97, 0xe9, 0x2e, 0x6c, 0x8f, 0x19, 0x86, 0x3e, 0x2b, 0x74, 0x07, 0x68,
	0xde, 0x94, 0x66, 0x03, 0x95, 0xad, 0x00, 0xbb, 0x91, 0xe7, 0x2a, 0xfd, 0x17, 0x8e, 0x8b, 0xa1,
	0x1b, 0x89, 0xd6, 0xe8, 0x19, 0x9c, 0x14, 0x60, 0x2d, 0xb1, 0x89, 0x96, 0x90, 0x0e, 0x1a, 0xc8,
	0x9b, 0xc8, 0x48, 0x85, 0x5e, 0xc0, 0xbb, 0x09, 0xf0, 0xec, 0x5a, 0x26, 0x67, 0xba, 0xd0, 0x73,
	0x61, 0xd7, 0xe9, 0x39, 0x9c, 0xfd, 0x4e, 0x9e, 0x76, 0xb2, 0x41, 0x01, 0x16, 0xd3, 0xdf, 0x5b,
	0x6a, 0x42, 0x09, 0xa7, 0x03, 0x8f, 0xe1, 0xe6, 0xfe, 0xa2, 0x7b, 0xb0, 0x33, 0xc6, 0xec, 0xe8,
	0x18, 0x6d, 0xab, 0x7f, 0x60, 0x91, 0x4d, 0xd6, 0x3c, 0xb3, 0xc6, 0x07, 0x18, 0x5f, 0xbe, 0x95,
	0xe0, 0xff, 0x76, 0xd4, 0xad, 0xfe, 0xc9, 0x71, 0xbc, 0x3c, 0x9c, 0x76, 0x19, 0x1b, 0xea, 0xc2,
	0x36, 0x4a, 0x9f, 0x1a, 0x8f, 0x41, 0xfc, 0xf4, 0xf2, 0x50, 0x6d, 0x47, 0x5d, 0x6d, 0x18, 0xf1,
	0xa2, 0x15, 0xa8, 0xf3, 0xec, 0xf7, 0xc2, 0x56, 0xe7, 0x22, 0xb9, 0xc5, 0x7d, 0xad, 0xdf, 0x6b,
	0x6b, 0xdd, 0x56, 0x10, 0x6a, 0xc9, 0x5b, 0x9b, 0x7a, 0xc7, 0x7f, 0x96, 0x4a, 0xdf, 0x67, 0xe7,
	0x74, 0x4f, 0x3c, 0x2c, 0x26, 0xca, 0x0f, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x43, 0x30, 0x32,
	0x5f, 0xf4, 0x05, 0x00, 0x00,
}
