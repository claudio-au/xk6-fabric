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
// source: autonomic/ext/event/state/token_refresh.proto

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

type WellKnownTokenRefreshState int32

const (
	WellKnownTokenRefreshState_UNKNOWN_STATE WellKnownTokenRefreshState = 0
	// Request was received by cloud through API
	// Emitted by: cloud
	WellKnownTokenRefreshState_REQUESTED WellKnownTokenRefreshState = 1
	// Request was queued in the cloud, waiting for device
	// to connect.
	// Emitted by: cloud
	// Triggered by: REQUEST_QUEUED_FOR_DELIVERY
	WellKnownTokenRefreshState_REQUEST_DELIVERY_QUEUED WellKnownTokenRefreshState = 2
	// Cloud attempted delivery of request to device
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_TRANSLATION_SUCCEEDED
	WellKnownTokenRefreshState_REQUEST_DELIVERY_IN_PROGRESS WellKnownTokenRefreshState = 3
	// Delivery of request attempt was not acknowledged by device in a timely fashion
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_TIMEOUT_EXCEEDED
	WellKnownTokenRefreshState_REQUEST_DELIVERY_TIMED_OUT WellKnownTokenRefreshState = 4
	// Delivery of request attempt failed due to device disconnect
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_FAILED
	WellKnownTokenRefreshState_REQUEST_DELIVERY_FAILURE WellKnownTokenRefreshState = 5
	// Request is acknowledged by device, queued, and the device will begin
	// processing the request when the device is ready.
	// Emitted by: device
	// Triggered by: REQUEST_QUEUED_ON_DEVICE
	WellKnownTokenRefreshState_REQUEST_QUEUED WellKnownTokenRefreshState = 6
	// Request is being decrypted and validated by device
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_STARTED
	WellKnownTokenRefreshState_REQUEST_VALIDATION_IN_PROGRESS WellKnownTokenRefreshState = 7
	// Request failed validation or decryption
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_FAILED
	WellKnownTokenRefreshState_REQUEST_VALIDATION_FAILURE WellKnownTokenRefreshState = 8
	// TokenRefresh is in the process of being reset
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED
	WellKnownTokenRefreshState_TOKEN_REFRESH_IN_PROGRESS WellKnownTokenRefreshState = 9
	// The device has successfully completed the actuation in question.
	// Triggered by: TOKEN_REFRESH_SUCCEEDED
	WellKnownTokenRefreshState_SUCCEEDED_ON_DEVICE WellKnownTokenRefreshState = 23
	// Oem cloud attempted delivery of request
	// Emitted by: cloud
	// Triggered by: DEVICE_CONNECTED_TO_CLOUD
	WellKnownTokenRefreshState_OEM_CLOUD_DELIVERY_IN_PROGRESS WellKnownTokenRefreshState = 10
	// Cloud attempted translation of request
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_DELIVERY_SUCCEEDED
	WellKnownTokenRefreshState_OEM_CLOUD_TRANSLATION_IN_PROGRESS WellKnownTokenRefreshState = 11
	// The cloud is in the process of validating the event emitted by the device indicating that the
	// command has succeeded.
	// Triggered by: SUCCEEDED_ON_DEVICE_EVENT_RECEIVED
	WellKnownTokenRefreshState_CLOUD_VALIDATION_IN_PROGRESS WellKnownTokenRefreshState = 24
	// Cloud failed to validate data provided by the device
	// Emitted by: cloud
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED
	WellKnownTokenRefreshState_CLOUD_VALIDATION_FAILURE WellKnownTokenRefreshState = 12
	// TokenRefresh upload succeeded, and the URI(s) of the archive(s) has been
	// added to the event containing this state.
	// This state is terminal
	// Triggered by: TOKEN_REFRESH_SUCCEEDED
	WellKnownTokenRefreshState_SUCCESS WellKnownTokenRefreshState = 17
	// TokenRefresh failed. All intermediate failure states should transition to
	// this state.
	// Emitted by: device
	// Triggered by: TOKEN_REFRESH_FAILED
	WellKnownTokenRefreshState_FAILURE WellKnownTokenRefreshState = 18
	// TokenRefresh command expired. This event can be emitted by
	// the device or the cloud. In either case, it will cause
	// the command to be removed from the set of in-progress commands
	// on the cloud. However the cloud does not send this event to
	// the device. This state is terminal.
	// Triggered by: EXPIRATION_EXCEEDED
	WellKnownTokenRefreshState_EXPIRED WellKnownTokenRefreshState = 19
	// Command is being removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_ACKNOWLEDGED
	WellKnownTokenRefreshState_CANCELLATION_IN_PROGRESS WellKnownTokenRefreshState = 21
	// Command was successfully removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_FULFILLED
	WellKnownTokenRefreshState_CANCELLED WellKnownTokenRefreshState = 22
)

var WellKnownTokenRefreshState_name = map[int32]string{
	0:  "UNKNOWN_STATE",
	1:  "REQUESTED",
	2:  "REQUEST_DELIVERY_QUEUED",
	3:  "REQUEST_DELIVERY_IN_PROGRESS",
	4:  "REQUEST_DELIVERY_TIMED_OUT",
	5:  "REQUEST_DELIVERY_FAILURE",
	6:  "REQUEST_QUEUED",
	7:  "REQUEST_VALIDATION_IN_PROGRESS",
	8:  "REQUEST_VALIDATION_FAILURE",
	9:  "TOKEN_REFRESH_IN_PROGRESS",
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

var WellKnownTokenRefreshState_value = map[string]int32{
	"UNKNOWN_STATE":                     0,
	"REQUESTED":                         1,
	"REQUEST_DELIVERY_QUEUED":           2,
	"REQUEST_DELIVERY_IN_PROGRESS":      3,
	"REQUEST_DELIVERY_TIMED_OUT":        4,
	"REQUEST_DELIVERY_FAILURE":          5,
	"REQUEST_QUEUED":                    6,
	"REQUEST_VALIDATION_IN_PROGRESS":    7,
	"REQUEST_VALIDATION_FAILURE":        8,
	"TOKEN_REFRESH_IN_PROGRESS":         9,
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

func (x WellKnownTokenRefreshState) String() string {
	return proto.EnumName(WellKnownTokenRefreshState_name, int32(x))
}

func (WellKnownTokenRefreshState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12bef689c0468891, []int{0}
}

// TokenRefresh state transition triggers. See WellKnownTokenRefreshState enum
// for documentation on when these triggers occur.
type WellKnownTokenRefreshTransitionTrigger int32

const (
	WellKnownTokenRefreshTransitionTrigger_UNKNOWN_TRANSITION_TRIGGER                WellKnownTokenRefreshTransitionTrigger = 0
	WellKnownTokenRefreshTransitionTrigger_REQUEST_QUEUED_FOR_DELIVERY               WellKnownTokenRefreshTransitionTrigger = 1
	WellKnownTokenRefreshTransitionTrigger_DEVICE_CONNECTED_TO_CLOUD                 WellKnownTokenRefreshTransitionTrigger = 2
	WellKnownTokenRefreshTransitionTrigger_REQUEST_DELIVERY_FAILED                   WellKnownTokenRefreshTransitionTrigger = 3
	WellKnownTokenRefreshTransitionTrigger_REQUEST_DELIVERY_TIMEOUT_EXCEEDED         WellKnownTokenRefreshTransitionTrigger = 4
	WellKnownTokenRefreshTransitionTrigger_REQUEST_QUEUED_ON_DEVICE                  WellKnownTokenRefreshTransitionTrigger = 5
	WellKnownTokenRefreshTransitionTrigger_REQUEST_VALIDATION_STARTED                WellKnownTokenRefreshTransitionTrigger = 6
	WellKnownTokenRefreshTransitionTrigger_REQUEST_VALIDATION_FAILED                 WellKnownTokenRefreshTransitionTrigger = 7
	WellKnownTokenRefreshTransitionTrigger_REQUEST_VALIDATION_SUCCEEDED              WellKnownTokenRefreshTransitionTrigger = 23
	WellKnownTokenRefreshTransitionTrigger_TOKEN_REFRESH_STARTED                     WellKnownTokenRefreshTransitionTrigger = 8
	WellKnownTokenRefreshTransitionTrigger_TOKEN_REFRESH_SUCCEEDED                   WellKnownTokenRefreshTransitionTrigger = 9
	WellKnownTokenRefreshTransitionTrigger_TOKEN_REFRESH_FAILED                      WellKnownTokenRefreshTransitionTrigger = 10
	WellKnownTokenRefreshTransitionTrigger_OEM_CLOUD_DELIVERY_SUCCEEDED              WellKnownTokenRefreshTransitionTrigger = 11
	WellKnownTokenRefreshTransitionTrigger_OEM_CLOUD_TRANSLATION_SUCCEEDED           WellKnownTokenRefreshTransitionTrigger = 12
	WellKnownTokenRefreshTransitionTrigger_SUCCEEDED_ON_DEVICE_EVENT_RECEIVED        WellKnownTokenRefreshTransitionTrigger = 24
	WellKnownTokenRefreshTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED WellKnownTokenRefreshTransitionTrigger = 13
	WellKnownTokenRefreshTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED    WellKnownTokenRefreshTransitionTrigger = 14
	WellKnownTokenRefreshTransitionTrigger_FAILED                                    WellKnownTokenRefreshTransitionTrigger = 19
	WellKnownTokenRefreshTransitionTrigger_EXPIRATION_EXCEEDED                       WellKnownTokenRefreshTransitionTrigger = 20
	WellKnownTokenRefreshTransitionTrigger_CANCELLATION_REQUESTED                    WellKnownTokenRefreshTransitionTrigger = 21
	WellKnownTokenRefreshTransitionTrigger_CANCELLATION_REQUEST_FULFILLED            WellKnownTokenRefreshTransitionTrigger = 22
)

var WellKnownTokenRefreshTransitionTrigger_name = map[int32]string{
	0:  "UNKNOWN_TRANSITION_TRIGGER",
	1:  "REQUEST_QUEUED_FOR_DELIVERY",
	2:  "DEVICE_CONNECTED_TO_CLOUD",
	3:  "REQUEST_DELIVERY_FAILED",
	4:  "REQUEST_DELIVERY_TIMEOUT_EXCEEDED",
	5:  "REQUEST_QUEUED_ON_DEVICE",
	6:  "REQUEST_VALIDATION_STARTED",
	7:  "REQUEST_VALIDATION_FAILED",
	23: "REQUEST_VALIDATION_SUCCEEDED",
	8:  "TOKEN_REFRESH_STARTED",
	9:  "TOKEN_REFRESH_SUCCEEDED",
	10: "TOKEN_REFRESH_FAILED",
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

var WellKnownTokenRefreshTransitionTrigger_value = map[string]int32{
	"UNKNOWN_TRANSITION_TRIGGER":                0,
	"REQUEST_QUEUED_FOR_DELIVERY":               1,
	"DEVICE_CONNECTED_TO_CLOUD":                 2,
	"REQUEST_DELIVERY_FAILED":                   3,
	"REQUEST_DELIVERY_TIMEOUT_EXCEEDED":         4,
	"REQUEST_QUEUED_ON_DEVICE":                  5,
	"REQUEST_VALIDATION_STARTED":                6,
	"REQUEST_VALIDATION_FAILED":                 7,
	"REQUEST_VALIDATION_SUCCEEDED":              23,
	"TOKEN_REFRESH_STARTED":                     8,
	"TOKEN_REFRESH_SUCCEEDED":                   9,
	"TOKEN_REFRESH_FAILED":                      10,
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

func (x WellKnownTokenRefreshTransitionTrigger) String() string {
	return proto.EnumName(WellKnownTokenRefreshTransitionTrigger_name, int32(x))
}

func (WellKnownTokenRefreshTransitionTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12bef689c0468891, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.event.state.tokenrefresh.WellKnownTokenRefreshState", WellKnownTokenRefreshState_name, WellKnownTokenRefreshState_value)
	proto.RegisterEnum("autonomic.ext.event.state.tokenrefresh.WellKnownTokenRefreshTransitionTrigger", WellKnownTokenRefreshTransitionTrigger_name, WellKnownTokenRefreshTransitionTrigger_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/event/state/token_refresh.proto", fileDescriptor_12bef689c0468891)
}

var fileDescriptor_12bef689c0468891 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdb, 0x6e, 0xda, 0x4c,
	0x10, 0xc7, 0x43, 0xce, 0x99, 0x1c, 0xb4, 0xd9, 0x1c, 0x20, 0xe7, 0xef, 0x4b, 0x55, 0xa4, 0x22,
	0x81, 0x2f, 0xfa, 0x04, 0x8e, 0x77, 0x48, 0xad, 0x10, 0x9b, 0xac, 0xd7, 0x24, 0xed, 0xcd, 0x8a,
	0x20, 0x37, 0x41, 0x0d, 0x76, 0x65, 0x9c, 0x36, 0xcf, 0xd3, 0xcb, 0x3c, 0x4b, 0x1f, 0xa8, 0x97,
	0xd5, 0xfa, 0x00, 0xb8, 0x18, 0xb5, 0x77, 0x2c, 0xff, 0xff, 0xce, 0xec, 0xcc, 0xfc, 0x3c, 0x50,
	0xef, 0x3e, 0x47, 0x81, 0x1f, 0x0c, 0xfa, 0x3d, 0xcd, 0x7b, 0x89, 0x34, 0xef, 0x9b, 0xe7, 0x47,
	0xda, 0x30, 0xea, 0x46, 0x9e, 0x16, 0x05, 0x5f, 0x3c, 0x5f, 0x86, 0xde, 0xe7, 0xd0, 0x1b, 0x3e,
	0x36, 0xbe, 0x86, 0x41, 0x14, 0xd0, 0xea, 0xc8, 0xde, 0xf0, 0x5e, 0xa2, 0x46, 0x6c, 0x6f, 0xc4,
	0xf6, 0x46, 0x6c, 0x4f, 0xdd, 0xb5, 0xd7, 0x45, 0x38, 0xbc, 0xf5, 0x9e, 0x9e, 0xae, 0xfc, 0xe0,
	0xbb, 0x2f, 0x94, 0xc2, 0x13, 0xc5, 0x51, 0x5e, 0xba, 0x0d, 0x9b, 0xae, 0x75, 0x65, 0xd9, 0xb7,
	0x96, 0x74, 0x84, 0x2e, 0x90, 0xcc, 0xd1, 0x4d, 0x58, 0xe3, 0x78, 0xe3, 0xa2, 0x23, 0x90, 0x91,
	0x12, 0x3d, 0x82, 0x72, 0x7a, 0x94, 0x0c, 0x5b, 0x66, 0x07, 0xf9, 0x47, 0x79, 0xe3, 0xa2, 0x8b,
	0x8c, 0xcc, 0xd3, 0xff, 0xe0, 0x78, 0x4a, 0x34, 0x2d, 0xd9, 0xe6, 0xf6, 0x25, 0x47, 0xc7, 0x21,
	0x0b, 0xf4, 0x14, 0x0e, 0xa7, 0x1c, 0xc2, 0xbc, 0x46, 0x26, 0x6d, 0x57, 0x90, 0x45, 0x7a, 0x0c,
	0x95, 0x29, 0xbd, 0xa9, 0x9b, 0x2d, 0x97, 0x23, 0x59, 0xa2, 0x14, 0xb6, 0x32, 0x35, 0xcd, 0xb9,
	0x4c, 0xcf, 0xe1, 0x34, 0xfb, 0xaf, 0xa3, 0xb7, 0x4c, 0xa6, 0x0b, 0xd3, 0xb6, 0x72, 0x59, 0x57,
	0x26, 0xb3, 0x4e, 0x78, 0xb2, 0xb8, 0xab, 0xf4, 0x04, 0x0e, 0x84, 0x7d, 0x85, 0x96, 0xe4, 0xd8,
	0xe4, 0xe8, 0x7c, 0xc8, 0x5d, 0x5f, 0xa3, 0x65, 0xd8, 0x71, 0x5c, 0xc3, 0x40, 0x64, 0xea, 0x9d,
	0x96, 0x64, 0xd8, 0x31, 0x0d, 0x24, 0x65, 0x95, 0xdb, 0xc6, 0x6b, 0x69, 0xb4, 0x6c, 0x97, 0x15,
	0x57, 0x0c, 0xf4, 0x2d, 0xfc, 0x3f, 0xf6, 0x08, 0xae, 0x5b, 0x4e, 0x6b, 0xfa, 0x89, 0xeb, 0xaa,
	0x75, 0x89, 0x65, 0x46, 0x11, 0x15, 0xd5, 0x9a, 0x29, 0x47, 0x56, 0xc2, 0x06, 0x5d, 0x87, 0x95,
	0xf8, 0x8d, 0x8e, 0x43, 0xb6, 0xd5, 0x21, 0x53, 0xa8, 0x3a, 0xe0, 0x5d, 0xdb, 0xe4, 0xc8, 0xc8,
	0x4e, 0x1c, 0x44, 0xb7, 0x0c, 0x6c, 0x15, 0x3c, 0x62, 0x4f, 0xcd, 0x3a, 0x55, 0x91, 0x91, 0xfd,
	0xda, 0xcf, 0x25, 0xa8, 0x16, 0xc2, 0x22, 0xc2, 0xae, 0x3f, 0xec, 0x47, 0xfd, 0xc0, 0x17, 0x61,
	0xff, 0xe1, 0xc1, 0x0b, 0x55, 0x87, 0x33, 0x70, 0xe2, 0x1a, 0xcd, 0x38, 0xba, 0xe0, 0xe6, 0xe5,
	0x25, 0x72, 0x32, 0x47, 0xcf, 0xe0, 0x28, 0x3f, 0x39, 0xd9, 0xb4, 0xf9, 0xa8, 0x65, 0xa4, 0xa4,
	0x46, 0x90, 0xb4, 0x55, 0x1a, 0xb6, 0x65, 0xa1, 0x21, 0x90, 0x49, 0x61, 0x27, 0x6d, 0x23, 0xf3,
	0x85, 0xd8, 0xa9, 0x12, 0x91, 0x91, 0x05, 0xd5, 0xe2, 0x42, 0xa8, 0x6c, 0x57, 0x48, 0xbc, 0x4b,
	0xe6, 0x96, 0x67, 0x2b, 0x7d, 0xc3, 0x78, 0x96, 0x4b, 0x33, 0x18, 0x71, 0x84, 0xce, 0x45, 0xcc,
	0xd9, 0x09, 0x1c, 0xcc, 0x60, 0x08, 0x19, 0x59, 0x99, 0x44, 0x7f, 0xf2, 0x7a, 0x86, 0x0d, 0x29,
	0xd3, 0x03, 0xd8, 0xcb, 0x43, 0x96, 0xc5, 0x5e, 0x55, 0xd5, 0xfd, 0x21, 0x8d, 0xee, 0xad, 0xd1,
	0x0a, 0xec, 0xe6, 0xc5, 0x34, 0x27, 0xa8, 0x9c, 0x05, 0xf8, 0x8d, 0xef, 0xae, 0xd3, 0x37, 0x70,
	0x56, 0x0c, 0xdf, 0xd8, 0xb4, 0x41, 0xab, 0x70, 0x5e, 0x80, 0xb7, 0xc4, 0x0e, 0x5a, 0x42, 0x72,
	0x34, 0xd0, 0xec, 0x20, 0x23, 0x15, 0x5a, 0x87, 0x77, 0x53, 0x00, 0xda, 0xcd, 0xcc, 0xce, 0x74,
	0xa1, 0x4f, 0x84, 0xdd, 0xa4, 0x35, 0xa8, 0xfe, 0xcd, 0x9e, 0x56, 0xb2, 0x45, 0x01, 0x96, 0xd3,
	0xdf, 0x3b, 0xea, 0x6b, 0x8b, 0x79, 0x4d, 0x6e, 0x8c, 0xe6, 0xb7, 0x4b, 0x0f, 0x61, 0x3f, 0xc7,
	0xee, 0x78, 0x2d, 0xed, 0xa9, 0x2f, 0xb1, 0x48, 0x93, 0x4d, 0xb7, 0xd5, 0x34, 0x13, 0x9c, 0x2f,
	0x5e, 0x4b, 0x50, 0xeb, 0x05, 0x83, 0xc6, 0xbf, 0xad, 0xca, 0x8b, 0xb3, 0xd9, 0x7b, 0xb2, 0xad,
	0x76, 0x6e, 0xbb, 0xf4, 0xa9, 0xfd, 0xd0, 0x8f, 0x1e, 0x9f, 0xef, 0x1b, 0xbd, 0x60, 0xa0, 0x8d,
	0xa2, 0xd6, 0xbb, 0x7d, 0xb5, 0xb2, 0xbd, 0xd0, 0xef, 0x3e, 0xd5, 0xe3, 0xed, 0x3c, 0xd4, 0x86,
	0x61, 0x4f, 0x1b, 0x74, 0xfb, 0xbe, 0x16, 0x9f, 0xb5, 0x99, 0xbb, 0xfd, 0x57, 0xa9, 0xf4, 0x63,
	0x7e, 0x41, 0x77, 0xc5, 0xfd, 0x72, 0xec, 0x7c, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x7e,
	0x34, 0x30, 0x08, 0x06, 0x00, 0x00,
}