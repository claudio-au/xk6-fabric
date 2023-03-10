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
// source: autonomic/ext/event/state/command_cancellation.proto

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

type WellKnownCommandCancellationState int32

const (
	WellKnownCommandCancellationState_UNKNOWN_STATE WellKnownCommandCancellationState = 0
	// Request was received by cloud through API.
	// Emitted by: cloud.
	WellKnownCommandCancellationState_REQUESTED WellKnownCommandCancellationState = 1
	// Request was queued in the cloud, waiting for device to connect.
	// Emitted by: cloud.
	// Triggered by: REQUEST_QUEUED_FOR_DELIVERY.
	WellKnownCommandCancellationState_REQUEST_DELIVERY_QUEUED WellKnownCommandCancellationState = 2
	// Cloud attempted delivery of request to device.
	// Emitted by: cloud.
	// Triggered by: OEM_CLOUD_TRANSLATION_SUCCEEDED.
	WellKnownCommandCancellationState_REQUEST_DELIVERY_IN_PROGRESS WellKnownCommandCancellationState = 3
	// Delivery of request attempt was not acknowledged by device in a timely
	// fashion.
	// Emitted by: cloud.
	// Triggered by: REQUEST_DELIVERY_TIMEOUT_EXCEEDED.
	WellKnownCommandCancellationState_REQUEST_DELIVERY_TIMED_OUT WellKnownCommandCancellationState = 4
	// Delivery of request attempt failed due to device disconnect.
	// Emitted by: cloud.
	// Triggered by: REQUEST_DELIVERY_FAILED.
	WellKnownCommandCancellationState_REQUEST_DELIVERY_FAILURE WellKnownCommandCancellationState = 5
	// Request is acknowledged by device, queued, and the device will begin
	// processing the request when the device is ready.
	// Emitted by: device.
	// Triggered by: REQUEST_QUEUED_ON_DEVICE.
	WellKnownCommandCancellationState_REQUEST_QUEUED WellKnownCommandCancellationState = 6
	// Request is being decrypted and validated by device.
	// Emitted by: device.
	// Triggered by: REQUEST_VALIDATION_STARTED.
	WellKnownCommandCancellationState_REQUEST_VALIDATION_IN_PROGRESS WellKnownCommandCancellationState = 7
	// Request failed validation or decryption.
	// Emitted by: device.
	// Triggered by: REQUEST_VALIDATION_FAILED.
	WellKnownCommandCancellationState_REQUEST_VALIDATION_FAILURE WellKnownCommandCancellationState = 8
	// The device has successfully completed the actuation in question.
	// Triggered by: COMMAND_CANCELLATION_SUCCEEDED
	WellKnownCommandCancellationState_SUCCEEDED_ON_DEVICE WellKnownCommandCancellationState = 23
	// Command is in the process of being cancelled.
	// Emitted by: device.
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED.
	WellKnownCommandCancellationState_COMMAND_CANCELLATION_IN_PROGRESS WellKnownCommandCancellationState = 9
	// Oem cloud attempted delivery of request.
	// Emitted by: cloud.
	// Triggered by: DEVICE_CONNECTED_TO_CLOUD.
	WellKnownCommandCancellationState_OEM_CLOUD_DELIVERY_IN_PROGRESS WellKnownCommandCancellationState = 10
	// Cloud attempted translation of request.
	// Emitted by: cloud.
	// Triggered by: OEM_CLOUD_DELIVERY_SUCCEEDED.
	WellKnownCommandCancellationState_OEM_CLOUD_TRANSLATION_IN_PROGRESS WellKnownCommandCancellationState = 11
	// The cloud is in the process of validating the event emitted by the device indicating that the
	// command has succeeded.
	// Triggered by: SUCCEEDED_ON_DEVICE_EVENT_RECEIVED
	WellKnownCommandCancellationState_CLOUD_VALIDATION_IN_PROGRESS WellKnownCommandCancellationState = 24
	// Cloud failed to validate data provided by the device
	// Emitted by: cloud
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED
	WellKnownCommandCancellationState_CLOUD_VALIDATION_FAILURE WellKnownCommandCancellationState = 12
	// Command cancellation upload succeeded, and the URI(s) of the archive(s)
	// has been added to the event containing this state.
	// This state is terminal.
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED.
	WellKnownCommandCancellationState_SUCCESS WellKnownCommandCancellationState = 17
	// Command cancellation failed. All intermediate failure states should
	// transition to this state.
	// Emitted by: device.
	// Triggered by: COMMAND_CANCELLATION_FAILED.
	WellKnownCommandCancellationState_FAILURE WellKnownCommandCancellationState = 18
	// Command cancellation command expired. This event can be emitted by the
	// device or the cloud. In either case, it will cause the command to be
	// removed from the set of in-progress commands on the cloud. However the
	// cloud does not send this event to the device.
	// This state is terminal.
	// Triggered by: EXPIRATION_EXCEEDED.
	WellKnownCommandCancellationState_EXPIRED WellKnownCommandCancellationState = 19
)

var WellKnownCommandCancellationState_name = map[int32]string{
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
	9:  "COMMAND_CANCELLATION_IN_PROGRESS",
	10: "OEM_CLOUD_DELIVERY_IN_PROGRESS",
	11: "OEM_CLOUD_TRANSLATION_IN_PROGRESS",
	24: "CLOUD_VALIDATION_IN_PROGRESS",
	12: "CLOUD_VALIDATION_FAILURE",
	17: "SUCCESS",
	18: "FAILURE",
	19: "EXPIRED",
}

var WellKnownCommandCancellationState_value = map[string]int32{
	"UNKNOWN_STATE":                     0,
	"REQUESTED":                         1,
	"REQUEST_DELIVERY_QUEUED":           2,
	"REQUEST_DELIVERY_IN_PROGRESS":      3,
	"REQUEST_DELIVERY_TIMED_OUT":        4,
	"REQUEST_DELIVERY_FAILURE":          5,
	"REQUEST_QUEUED":                    6,
	"REQUEST_VALIDATION_IN_PROGRESS":    7,
	"REQUEST_VALIDATION_FAILURE":        8,
	"SUCCEEDED_ON_DEVICE":               23,
	"COMMAND_CANCELLATION_IN_PROGRESS":  9,
	"OEM_CLOUD_DELIVERY_IN_PROGRESS":    10,
	"OEM_CLOUD_TRANSLATION_IN_PROGRESS": 11,
	"CLOUD_VALIDATION_IN_PROGRESS":      24,
	"CLOUD_VALIDATION_FAILURE":          12,
	"SUCCESS":                           17,
	"FAILURE":                           18,
	"EXPIRED":                           19,
}

func (x WellKnownCommandCancellationState) String() string {
	return proto.EnumName(WellKnownCommandCancellationState_name, int32(x))
}

func (WellKnownCommandCancellationState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a94fea03155cccd3, []int{0}
}

// Command cancellation state transition triggers. See
// WellKnownCommandCancellationState enum for documentation on when these
// triggers occur.
type WellKnownCommandCancellationTransitionTrigger int32

const (
	WellKnownCommandCancellationTransitionTrigger_UNKNOWN_TRANSITION_TRIGGER                WellKnownCommandCancellationTransitionTrigger = 0
	WellKnownCommandCancellationTransitionTrigger_REQUEST_QUEUED_FOR_DELIVERY               WellKnownCommandCancellationTransitionTrigger = 1
	WellKnownCommandCancellationTransitionTrigger_DEVICE_CONNECTED_TO_CLOUD                 WellKnownCommandCancellationTransitionTrigger = 2
	WellKnownCommandCancellationTransitionTrigger_REQUEST_DELIVERY_FAILED                   WellKnownCommandCancellationTransitionTrigger = 3
	WellKnownCommandCancellationTransitionTrigger_REQUEST_DELIVERY_TIMEOUT_EXCEEDED         WellKnownCommandCancellationTransitionTrigger = 4
	WellKnownCommandCancellationTransitionTrigger_REQUEST_QUEUED_ON_DEVICE                  WellKnownCommandCancellationTransitionTrigger = 5
	WellKnownCommandCancellationTransitionTrigger_REQUEST_VALIDATION_STARTED                WellKnownCommandCancellationTransitionTrigger = 6
	WellKnownCommandCancellationTransitionTrigger_REQUEST_VALIDATION_FAILED                 WellKnownCommandCancellationTransitionTrigger = 7
	WellKnownCommandCancellationTransitionTrigger_REQUEST_VALIDATION_SUCCEEDED              WellKnownCommandCancellationTransitionTrigger = 23
	WellKnownCommandCancellationTransitionTrigger_COMMAND_CANCELLATION_STARTED              WellKnownCommandCancellationTransitionTrigger = 8
	WellKnownCommandCancellationTransitionTrigger_COMMAND_CANCELLATION_SUCCEEDED            WellKnownCommandCancellationTransitionTrigger = 9
	WellKnownCommandCancellationTransitionTrigger_COMMAND_CANCELLATION_FAILED               WellKnownCommandCancellationTransitionTrigger = 10
	WellKnownCommandCancellationTransitionTrigger_OEM_CLOUD_DELIVERY_SUCCEEDED              WellKnownCommandCancellationTransitionTrigger = 11
	WellKnownCommandCancellationTransitionTrigger_OEM_CLOUD_TRANSLATION_SUCCEEDED           WellKnownCommandCancellationTransitionTrigger = 12
	WellKnownCommandCancellationTransitionTrigger_SUCCEEDED_ON_DEVICE_EVENT_RECEIVED        WellKnownCommandCancellationTransitionTrigger = 24
	WellKnownCommandCancellationTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED WellKnownCommandCancellationTransitionTrigger = 13
	WellKnownCommandCancellationTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED    WellKnownCommandCancellationTransitionTrigger = 14
	WellKnownCommandCancellationTransitionTrigger_FAILED                                    WellKnownCommandCancellationTransitionTrigger = 19
	WellKnownCommandCancellationTransitionTrigger_EXPIRATION_EXCEEDED                       WellKnownCommandCancellationTransitionTrigger = 20
)

var WellKnownCommandCancellationTransitionTrigger_name = map[int32]string{
	0:  "UNKNOWN_TRANSITION_TRIGGER",
	1:  "REQUEST_QUEUED_FOR_DELIVERY",
	2:  "DEVICE_CONNECTED_TO_CLOUD",
	3:  "REQUEST_DELIVERY_FAILED",
	4:  "REQUEST_DELIVERY_TIMEOUT_EXCEEDED",
	5:  "REQUEST_QUEUED_ON_DEVICE",
	6:  "REQUEST_VALIDATION_STARTED",
	7:  "REQUEST_VALIDATION_FAILED",
	23: "REQUEST_VALIDATION_SUCCEEDED",
	8:  "COMMAND_CANCELLATION_STARTED",
	9:  "COMMAND_CANCELLATION_SUCCEEDED",
	10: "COMMAND_CANCELLATION_FAILED",
	11: "OEM_CLOUD_DELIVERY_SUCCEEDED",
	12: "OEM_CLOUD_TRANSLATION_SUCCEEDED",
	24: "SUCCEEDED_ON_DEVICE_EVENT_RECEIVED",
	13: "CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED",
	14: "CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED",
	19: "FAILED",
	20: "EXPIRATION_EXCEEDED",
}

var WellKnownCommandCancellationTransitionTrigger_value = map[string]int32{
	"UNKNOWN_TRANSITION_TRIGGER":                0,
	"REQUEST_QUEUED_FOR_DELIVERY":               1,
	"DEVICE_CONNECTED_TO_CLOUD":                 2,
	"REQUEST_DELIVERY_FAILED":                   3,
	"REQUEST_DELIVERY_TIMEOUT_EXCEEDED":         4,
	"REQUEST_QUEUED_ON_DEVICE":                  5,
	"REQUEST_VALIDATION_STARTED":                6,
	"REQUEST_VALIDATION_FAILED":                 7,
	"REQUEST_VALIDATION_SUCCEEDED":              23,
	"COMMAND_CANCELLATION_STARTED":              8,
	"COMMAND_CANCELLATION_SUCCEEDED":            9,
	"COMMAND_CANCELLATION_FAILED":               10,
	"OEM_CLOUD_DELIVERY_SUCCEEDED":              11,
	"OEM_CLOUD_TRANSLATION_SUCCEEDED":           12,
	"SUCCEEDED_ON_DEVICE_EVENT_RECEIVED":        24,
	"CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED": 13,
	"CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED":    14,
	"FAILED":                                    19,
	"EXPIRATION_EXCEEDED":                       20,
}

func (x WellKnownCommandCancellationTransitionTrigger) String() string {
	return proto.EnumName(WellKnownCommandCancellationTransitionTrigger_name, int32(x))
}

func (WellKnownCommandCancellationTransitionTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a94fea03155cccd3, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.event.state.cancellation.WellKnownCommandCancellationState", WellKnownCommandCancellationState_name, WellKnownCommandCancellationState_value)
	proto.RegisterEnum("autonomic.ext.event.state.cancellation.WellKnownCommandCancellationTransitionTrigger", WellKnownCommandCancellationTransitionTrigger_name, WellKnownCommandCancellationTransitionTrigger_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/event/state/command_cancellation.proto", fileDescriptor_a94fea03155cccd3)
}

var fileDescriptor_a94fea03155cccd3 = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xc7, 0x09, 0x04, 0x08, 0xc3, 0x87, 0x96, 0xa5, 0x10, 0x28, 0xdf, 0x69, 0x1b, 0xa9, 0x48,
	0x89, 0x55, 0xb5, 0x2f, 0x60, 0xbc, 0x03, 0xb2, 0x48, 0xec, 0x60, 0xaf, 0x03, 0xed, 0x65, 0x65,
	0x52, 0x8b, 0x5a, 0x4a, 0xec, 0x2a, 0x31, 0x2d, 0x6f, 0xd2, 0x7b, 0x9f, 0xa1, 0x87, 0x3e, 0x5a,
	0x8f, 0xd5, 0xda, 0x71, 0x62, 0x64, 0x07, 0x6e, 0xde, 0xcc, 0x7f, 0xbe, 0x7f, 0x19, 0xf8, 0xe4,
	0x3e, 0x44, 0x61, 0x10, 0x0e, 0xfc, 0x9e, 0xe2, 0x3d, 0x46, 0x8a, 0xf7, 0xc3, 0x0b, 0x22, 0x65,
	0x14, 0xb9, 0x91, 0xa7, 0xf4, 0xc2, 0xc1, 0xc0, 0x0d, 0xbe, 0x8a, 0x9e, 0x1b, 0xf4, 0xbc, 0x7e,
	0xdf, 0x8d, 0xfc, 0x30, 0x68, 0x7e, 0x1f, 0x86, 0x51, 0x48, 0xeb, 0x13, 0xaf, 0xa6, 0xf7, 0x18,
	0x35, 0x63, 0xaf, 0x66, 0xec, 0xd5, 0xcc, 0xaa, 0xcf, 0x7e, 0x95, 0xe1, 0xf4, 0xc6, 0xeb, 0xf7,
	0xaf, 0x82, 0xf0, 0x67, 0xa0, 0x25, 0xf1, 0xb4, 0x8c, 0xc0, 0x96, 0x2e, 0x74, 0x13, 0xd6, 0x1d,
	0xe3, 0xca, 0x30, 0x6f, 0x0c, 0x61, 0x73, 0x95, 0x23, 0x99, 0xa3, 0xeb, 0xb0, 0x62, 0xe1, 0xb5,
	0x83, 0x36, 0x47, 0x46, 0x4a, 0x74, 0x1f, 0xaa, 0xe3, 0xa7, 0x60, 0xd8, 0xd2, 0xbb, 0x68, 0x7d,
	0x16, 0xd7, 0x0e, 0x3a, 0xc8, 0xc8, 0x3c, 0x3d, 0x81, 0x83, 0x9c, 0x51, 0x37, 0x44, 0xc7, 0x32,
	0x2f, 0x2d, 0xb4, 0x6d, 0xb2, 0x40, 0x8f, 0xe0, 0x75, 0x4e, 0xc1, 0xf5, 0x36, 0x32, 0x61, 0x3a,
	0x9c, 0x94, 0xe9, 0x01, 0xec, 0xe6, 0xec, 0x17, 0xaa, 0xde, 0x72, 0x2c, 0x24, 0x8b, 0x94, 0xc2,
	0x46, 0x6a, 0x1d, 0xe7, 0x5c, 0xa2, 0x35, 0x38, 0x4a, 0x7f, 0xeb, 0xaa, 0x2d, 0x9d, 0xa9, 0x5c,
	0x37, 0x8d, 0x27, 0x59, 0x97, 0xb3, 0x59, 0x33, 0x9a, 0x34, 0x6e, 0x85, 0x56, 0x61, 0xcb, 0x76,
	0x34, 0x0d, 0x91, 0xc9, 0x42, 0x0c, 0xc1, 0xb0, 0xab, 0x6b, 0x48, 0xaa, 0xf4, 0x2d, 0x9c, 0x68,
	0x66, 0xbb, 0xad, 0x1a, 0x4c, 0x68, 0xaa, 0xa1, 0x61, 0xab, 0x95, 0x0f, 0xbf, 0x22, 0x4b, 0x30,
	0xb1, 0x2d, 0xb4, 0x96, 0xe9, 0xb0, 0xe2, 0xc6, 0x81, 0xbe, 0x83, 0xd3, 0xa9, 0x86, 0x5b, 0xaa,
	0x61, 0x17, 0x84, 0x5a, 0x95, 0x13, 0x4c, 0x24, 0x33, 0x7a, 0xd9, 0x95, 0x13, 0xca, 0x29, 0xd2,
	0x4e, 0xd6, 0xe8, 0x2a, 0x2c, 0xc7, 0x9d, 0xd8, 0x36, 0xd9, 0x94, 0x8f, 0xd4, 0x42, 0xe5, 0x03,
	0x6f, 0x3b, 0xba, 0x85, 0x8c, 0x6c, 0xd5, 0xca, 0x95, 0x6d, 0xb2, 0x5d, 0x2b, 0x57, 0x76, 0xc8,
	0xce, 0xd9, 0x9f, 0x45, 0x68, 0x3c, 0x47, 0x06, 0x1f, 0xba, 0xc1, 0xc8, 0x4f, 0xbe, 0xfc, 0xfb,
	0x7b, 0x6f, 0x28, 0xc7, 0x99, 0x52, 0x12, 0x77, 0xa2, 0xc7, 0x45, 0x70, 0x4b, 0xbf, 0xbc, 0x44,
	0x8b, 0xcc, 0xd1, 0x63, 0xd8, 0x7f, 0xba, 0x26, 0x71, 0x61, 0x5a, 0x93, 0xc1, 0x90, 0x12, 0x3d,
	0x84, 0xbd, 0x64, 0xc4, 0x42, 0x33, 0x0d, 0x03, 0x35, 0x8e, 0x4c, 0x70, 0x33, 0x19, 0x0e, 0x99,
	0x2f, 0x64, 0x4c, 0x36, 0x82, 0x8c, 0x2c, 0xc8, 0x41, 0x16, 0x12, 0x64, 0x3a, 0x5c, 0xe0, 0x6d,
	0xb2, 0xc3, 0xa7, 0x20, 0x8d, 0x6b, 0x98, 0xee, 0x75, 0x71, 0x06, 0x10, 0x36, 0x57, 0x2d, 0x1e,
	0x43, 0x75, 0x08, 0x7b, 0x33, 0x80, 0x41, 0x46, 0x96, 0xb3, 0x9c, 0x67, 0xdd, 0x53, 0x84, 0x48,
	0x35, 0xde, 0x63, 0x11, 0x38, 0x69, 0x8a, 0x8a, 0x84, 0xa6, 0x58, 0x31, 0x89, 0xb2, 0x22, 0x07,
	0x59, 0xa8, 0x19, 0x17, 0x02, 0x32, 0x4d, 0x01, 0x79, 0xd3, 0x10, 0xab, 0xf4, 0x0d, 0x1c, 0x17,
	0x73, 0x37, 0x15, 0xad, 0xd1, 0x3a, 0xd4, 0x0a, 0xf8, 0x17, 0xd8, 0x45, 0x83, 0x0b, 0x0b, 0x35,
	0xd4, 0xbb, 0xc8, 0xc8, 0x2e, 0x6d, 0xc0, 0xfb, 0x1c, 0x7b, 0xe6, 0x45, 0x2a, 0x67, 0x2a, 0x57,
	0x33, 0x61, 0xd7, 0xe9, 0x19, 0xd4, 0x5f, 0x92, 0x8f, 0x3b, 0xd9, 0xa0, 0x00, 0x4b, 0xe3, 0xef,
	0x2d, 0xf9, 0x77, 0x8c, 0x51, 0x4d, 0x3c, 0x26, 0x4b, 0x7d, 0x95, 0xc5, 0xf6, 0xfc, 0x6f, 0x09,
	0x3e, 0xf4, 0xc2, 0x41, 0xf3, 0x99, 0xfb, 0x97, 0xb0, 0x9c, 0x3d, 0x83, 0xe7, 0xf5, 0x17, 0x6f,
	0x60, 0x47, 0x9e, 0xd5, 0x4e, 0xe9, 0x4b, 0xe7, 0xde, 0x8f, 0xbe, 0x3d, 0xdc, 0xc9, 0x28, 0xca,
	0x24, 0x47, 0xc3, 0xf5, 0xe5, 0x71, 0xf6, 0x86, 0x81, 0xdb, 0x6f, 0xc4, 0x07, 0x78, 0xa4, 0x8c,
	0x86, 0x3d, 0x65, 0xe0, 0xfa, 0x81, 0x12, 0xbf, 0x95, 0x99, 0x57, 0xfc, 0x5f, 0xa9, 0xf4, 0x7b,
	0x7e, 0x41, 0x75, 0xf8, 0xdd, 0x52, 0xac, 0xfc, 0xf8, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x38, 0x17,
	0x10, 0xb0, 0xf2, 0x05, 0x00, 0x00,
}
