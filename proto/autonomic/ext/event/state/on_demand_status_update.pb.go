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
// source: autonomic/ext/event/state/on_demand_status_update.proto

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

type WellKnownOnDemandStatusUpdateState int32

const (
	WellKnownOnDemandStatusUpdateState_UNKNOWN_STATE WellKnownOnDemandStatusUpdateState = 0
	// Request was received by cloud through API
	// Emitted by: cloud
	WellKnownOnDemandStatusUpdateState_REQUESTED WellKnownOnDemandStatusUpdateState = 1
	// Request was queued in the cloud, waiting for device
	// to connect.
	// Emitted by: cloud
	// Triggered by: REQUEST_QUEUED_FOR_DELIVERY
	WellKnownOnDemandStatusUpdateState_REQUEST_DELIVERY_QUEUED WellKnownOnDemandStatusUpdateState = 2
	// Cloud attempted delivery of request to device
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_TRANSLATION_SUCCEEDED
	WellKnownOnDemandStatusUpdateState_REQUEST_DELIVERY_IN_PROGRESS WellKnownOnDemandStatusUpdateState = 3
	// Delivery of request attempt was not acknowledged by device in a timely fashion
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_TIMEOUT_EXCEEDED
	WellKnownOnDemandStatusUpdateState_REQUEST_DELIVERY_TIMED_OUT WellKnownOnDemandStatusUpdateState = 4
	// Delivery of request attempt failed due to device disconnect
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_FAILED
	WellKnownOnDemandStatusUpdateState_REQUEST_DELIVERY_FAILURE WellKnownOnDemandStatusUpdateState = 5
	// Request is acknowledged by device, queued, and the device will begin
	// processing the request when the device is ready.
	// Emitted by: device
	// Triggered by: REQUEST_QUEUED_ON_DEVICE
	WellKnownOnDemandStatusUpdateState_REQUEST_QUEUED WellKnownOnDemandStatusUpdateState = 6
	// Request is being decrypted and validated by device
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_STARTED
	WellKnownOnDemandStatusUpdateState_REQUEST_VALIDATION_IN_PROGRESS WellKnownOnDemandStatusUpdateState = 7
	// Request failed validation or decryption
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_FAILED
	WellKnownOnDemandStatusUpdateState_REQUEST_VALIDATION_FAILURE WellKnownOnDemandStatusUpdateState = 8
	// OnDemandStatusUpdate is in the process of being reset
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED
	WellKnownOnDemandStatusUpdateState_ON_DEMAND_STATUS_UPDATE_IN_PROGRESS WellKnownOnDemandStatusUpdateState = 9
	// Oem cloud attempted delivery of request
	// Emitted by: cloud
	// Triggered by: DEVICE_CONNECTED_TO_CLOUD
	WellKnownOnDemandStatusUpdateState_OEM_CLOUD_DELIVERY_IN_PROGRESS WellKnownOnDemandStatusUpdateState = 10
	// Cloud attempted translation of request
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_DELIVERY_SUCCEEDED
	WellKnownOnDemandStatusUpdateState_OEM_CLOUD_TRANSLATION_IN_PROGRESS WellKnownOnDemandStatusUpdateState = 11
	// The cloud is in the process of validating the event emitted by the device indicating that the
	// command has succeeded.
	// Triggered by: SUCCEEDED_ON_DEVICE_EVENT_RECEIVED
	WellKnownOnDemandStatusUpdateState_CLOUD_VALIDATION_IN_PROGRESS WellKnownOnDemandStatusUpdateState = 24
	// Cloud failed to validate data provided by the device
	// Emitted by: cloud
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED
	WellKnownOnDemandStatusUpdateState_CLOUD_VALIDATION_FAILURE WellKnownOnDemandStatusUpdateState = 12
	// The device has successfully completed the actuation in question.
	// Triggered by: ON_DEMAND_STATUS_UPDATE_SUCCEEDED
	WellKnownOnDemandStatusUpdateState_SUCCEEDED_ON_DEVICE WellKnownOnDemandStatusUpdateState = 23
	// OnDemandStatusUpdate upload succeeded, and the URI(s) of the archive(s) has been
	// added to the event containing this state.
	// This state is terminal
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED
	WellKnownOnDemandStatusUpdateState_SUCCESS WellKnownOnDemandStatusUpdateState = 17
	// OnDemandStatusUpdate failed. All intermediate failure states should transition to
	// this state.
	// Emitted by: device
	// Triggered by: ON_DEMAND_STATUS_UPDATE_FAILED
	WellKnownOnDemandStatusUpdateState_FAILURE WellKnownOnDemandStatusUpdateState = 18
	// OnDemandStatusUpdate command expired. This event can be emitted by
	// the device or the cloud. In either case, it will cause
	// the command to be removed from the set of in-progress commands
	// on the cloud. However the cloud does not send this event to
	// the device. This state is terminal.
	// Triggered by: EXPIRATION_EXCEEDED
	WellKnownOnDemandStatusUpdateState_EXPIRED WellKnownOnDemandStatusUpdateState = 19
	// Command is being removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_ACKNOWLEDGED
	WellKnownOnDemandStatusUpdateState_CANCELLATION_IN_PROGRESS WellKnownOnDemandStatusUpdateState = 21
	// Command was successfully removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_FULFILLED
	WellKnownOnDemandStatusUpdateState_CANCELLED WellKnownOnDemandStatusUpdateState = 22
)

var WellKnownOnDemandStatusUpdateState_name = map[int32]string{
	0:  "UNKNOWN_STATE",
	1:  "REQUESTED",
	2:  "REQUEST_DELIVERY_QUEUED",
	3:  "REQUEST_DELIVERY_IN_PROGRESS",
	4:  "REQUEST_DELIVERY_TIMED_OUT",
	5:  "REQUEST_DELIVERY_FAILURE",
	6:  "REQUEST_QUEUED",
	7:  "REQUEST_VALIDATION_IN_PROGRESS",
	8:  "REQUEST_VALIDATION_FAILURE",
	9:  "ON_DEMAND_STATUS_UPDATE_IN_PROGRESS",
	10: "OEM_CLOUD_DELIVERY_IN_PROGRESS",
	11: "OEM_CLOUD_TRANSLATION_IN_PROGRESS",
	24: "CLOUD_VALIDATION_IN_PROGRESS",
	12: "CLOUD_VALIDATION_FAILURE",
	23: "SUCCEEDED_ON_DEVICE",
	17: "SUCCESS",
	18: "FAILURE",
	19: "EXPIRED",
	21: "CANCELLATION_IN_PROGRESS",
	22: "CANCELLED",
}

var WellKnownOnDemandStatusUpdateState_value = map[string]int32{
	"UNKNOWN_STATE":                       0,
	"REQUESTED":                           1,
	"REQUEST_DELIVERY_QUEUED":             2,
	"REQUEST_DELIVERY_IN_PROGRESS":        3,
	"REQUEST_DELIVERY_TIMED_OUT":          4,
	"REQUEST_DELIVERY_FAILURE":            5,
	"REQUEST_QUEUED":                      6,
	"REQUEST_VALIDATION_IN_PROGRESS":      7,
	"REQUEST_VALIDATION_FAILURE":          8,
	"ON_DEMAND_STATUS_UPDATE_IN_PROGRESS": 9,
	"OEM_CLOUD_DELIVERY_IN_PROGRESS":      10,
	"OEM_CLOUD_TRANSLATION_IN_PROGRESS":   11,
	"CLOUD_VALIDATION_IN_PROGRESS":        24,
	"CLOUD_VALIDATION_FAILURE":            12,
	"SUCCEEDED_ON_DEVICE":                 23,
	"SUCCESS":                             17,
	"FAILURE":                             18,
	"EXPIRED":                             19,
	"CANCELLATION_IN_PROGRESS":            21,
	"CANCELLED":                           22,
}

func (x WellKnownOnDemandStatusUpdateState) String() string {
	return proto.EnumName(WellKnownOnDemandStatusUpdateState_name, int32(x))
}

func (WellKnownOnDemandStatusUpdateState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c96966bb606ab1bd, []int{0}
}

// OnDemandStatusUpdate state transition triggers. See WellKnownOnDemandStatusUpdateState enum
// for documentation on when these triggers occur.
type WellKnownOnDemandStatusUpdateTransitionTrigger int32

const (
	WellKnownOnDemandStatusUpdateTransitionTrigger_UNKNOWN_TRANSITION_TRIGGER                WellKnownOnDemandStatusUpdateTransitionTrigger = 0
	WellKnownOnDemandStatusUpdateTransitionTrigger_REQUEST_QUEUED_FOR_DELIVERY               WellKnownOnDemandStatusUpdateTransitionTrigger = 1
	WellKnownOnDemandStatusUpdateTransitionTrigger_DEVICE_CONNECTED_TO_CLOUD                 WellKnownOnDemandStatusUpdateTransitionTrigger = 2
	WellKnownOnDemandStatusUpdateTransitionTrigger_REQUEST_DELIVERY_FAILED                   WellKnownOnDemandStatusUpdateTransitionTrigger = 3
	WellKnownOnDemandStatusUpdateTransitionTrigger_REQUEST_DELIVERY_TIMEOUT_EXCEEDED         WellKnownOnDemandStatusUpdateTransitionTrigger = 4
	WellKnownOnDemandStatusUpdateTransitionTrigger_REQUEST_QUEUED_ON_DEVICE                  WellKnownOnDemandStatusUpdateTransitionTrigger = 5
	WellKnownOnDemandStatusUpdateTransitionTrigger_REQUEST_VALIDATION_STARTED                WellKnownOnDemandStatusUpdateTransitionTrigger = 6
	WellKnownOnDemandStatusUpdateTransitionTrigger_REQUEST_VALIDATION_FAILED                 WellKnownOnDemandStatusUpdateTransitionTrigger = 7
	WellKnownOnDemandStatusUpdateTransitionTrigger_REQUEST_VALIDATION_SUCCEEDED              WellKnownOnDemandStatusUpdateTransitionTrigger = 23
	WellKnownOnDemandStatusUpdateTransitionTrigger_ON_DEMAND_STATUS_UPDATE_STARTED           WellKnownOnDemandStatusUpdateTransitionTrigger = 8
	WellKnownOnDemandStatusUpdateTransitionTrigger_ON_DEMAND_STATUS_UPDATE_SUCCEEDED         WellKnownOnDemandStatusUpdateTransitionTrigger = 9
	WellKnownOnDemandStatusUpdateTransitionTrigger_ON_DEMAND_STATUS_UPDATE_FAILED            WellKnownOnDemandStatusUpdateTransitionTrigger = 10
	WellKnownOnDemandStatusUpdateTransitionTrigger_OEM_CLOUD_DELIVERY_SUCCEEDED              WellKnownOnDemandStatusUpdateTransitionTrigger = 11
	WellKnownOnDemandStatusUpdateTransitionTrigger_OEM_CLOUD_TRANSLATION_SUCCEEDED           WellKnownOnDemandStatusUpdateTransitionTrigger = 12
	WellKnownOnDemandStatusUpdateTransitionTrigger_SUCCEEDED_ON_DEVICE_EVENT_RECEIVED        WellKnownOnDemandStatusUpdateTransitionTrigger = 24
	WellKnownOnDemandStatusUpdateTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED WellKnownOnDemandStatusUpdateTransitionTrigger = 13
	WellKnownOnDemandStatusUpdateTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED    WellKnownOnDemandStatusUpdateTransitionTrigger = 14
	WellKnownOnDemandStatusUpdateTransitionTrigger_FAILED                                    WellKnownOnDemandStatusUpdateTransitionTrigger = 19
	WellKnownOnDemandStatusUpdateTransitionTrigger_EXPIRATION_EXCEEDED                       WellKnownOnDemandStatusUpdateTransitionTrigger = 20
	WellKnownOnDemandStatusUpdateTransitionTrigger_CANCELLATION_REQUESTED                    WellKnownOnDemandStatusUpdateTransitionTrigger = 21
	WellKnownOnDemandStatusUpdateTransitionTrigger_CANCELLATION_REQUEST_FULFILLED            WellKnownOnDemandStatusUpdateTransitionTrigger = 22
)

var WellKnownOnDemandStatusUpdateTransitionTrigger_name = map[int32]string{
	0:  "UNKNOWN_TRANSITION_TRIGGER",
	1:  "REQUEST_QUEUED_FOR_DELIVERY",
	2:  "DEVICE_CONNECTED_TO_CLOUD",
	3:  "REQUEST_DELIVERY_FAILED",
	4:  "REQUEST_DELIVERY_TIMEOUT_EXCEEDED",
	5:  "REQUEST_QUEUED_ON_DEVICE",
	6:  "REQUEST_VALIDATION_STARTED",
	7:  "REQUEST_VALIDATION_FAILED",
	23: "REQUEST_VALIDATION_SUCCEEDED",
	8:  "ON_DEMAND_STATUS_UPDATE_STARTED",
	9:  "ON_DEMAND_STATUS_UPDATE_SUCCEEDED",
	10: "ON_DEMAND_STATUS_UPDATE_FAILED",
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

var WellKnownOnDemandStatusUpdateTransitionTrigger_value = map[string]int32{
	"UNKNOWN_TRANSITION_TRIGGER":                0,
	"REQUEST_QUEUED_FOR_DELIVERY":               1,
	"DEVICE_CONNECTED_TO_CLOUD":                 2,
	"REQUEST_DELIVERY_FAILED":                   3,
	"REQUEST_DELIVERY_TIMEOUT_EXCEEDED":         4,
	"REQUEST_QUEUED_ON_DEVICE":                  5,
	"REQUEST_VALIDATION_STARTED":                6,
	"REQUEST_VALIDATION_FAILED":                 7,
	"REQUEST_VALIDATION_SUCCEEDED":              23,
	"ON_DEMAND_STATUS_UPDATE_STARTED":           8,
	"ON_DEMAND_STATUS_UPDATE_SUCCEEDED":         9,
	"ON_DEMAND_STATUS_UPDATE_FAILED":            10,
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

func (x WellKnownOnDemandStatusUpdateTransitionTrigger) String() string {
	return proto.EnumName(WellKnownOnDemandStatusUpdateTransitionTrigger_name, int32(x))
}

func (WellKnownOnDemandStatusUpdateTransitionTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c96966bb606ab1bd, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.event.state.ondemandstatusupdate.WellKnownOnDemandStatusUpdateState", WellKnownOnDemandStatusUpdateState_name, WellKnownOnDemandStatusUpdateState_value)
	proto.RegisterEnum("autonomic.ext.event.state.ondemandstatusupdate.WellKnownOnDemandStatusUpdateTransitionTrigger", WellKnownOnDemandStatusUpdateTransitionTrigger_name, WellKnownOnDemandStatusUpdateTransitionTrigger_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/event/state/on_demand_status_update.proto", fileDescriptor_c96966bb606ab1bd)
}

var fileDescriptor_c96966bb606ab1bd = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5b, 0x4f, 0xdb, 0x4a,
	0x10, 0x26, 0xdc, 0x19, 0x2e, 0x5a, 0x96, 0x03, 0xe1, 0xce, 0x39, 0xa0, 0x03, 0x3a, 0x48, 0xd8,
	0xd2, 0xe9, 0x43, 0x9f, 0x8d, 0x77, 0x82, 0x2c, 0x82, 0x1d, 0xec, 0xdd, 0x40, 0xfb, 0xb2, 0x32,
	0xc1, 0xa2, 0x91, 0x88, 0x8d, 0x12, 0xa7, 0xe5, 0xf7, 0xf4, 0x67, 0x54, 0xfd, 0x55, 0x7d, 0xea,
	0x63, 0xb5, 0xeb, 0x38, 0x17, 0xc5, 0x29, 0xea, 0x9b, 0xc7, 0xf3, 0xcd, 0xed, 0x9b, 0xcf, 0x63,
	0x78, 0x1f, 0x76, 0xd3, 0x24, 0x4e, 0x5a, 0xcd, 0x86, 0x19, 0xbd, 0xa6, 0x66, 0xf4, 0x39, 0x8a,
	0x53, 0xb3, 0x93, 0x86, 0x69, 0x64, 0x26, 0xb1, 0x7c, 0x8c, 0x5a, 0x61, 0xfc, 0x28, 0x95, 0xdd,
	0xed, 0xc8, 0xee, 0xcb, 0x63, 0x98, 0x46, 0xc6, 0x4b, 0x3b, 0x49, 0x13, 0x6a, 0xf4, 0x03, 0x8d,
	0xe8, 0x35, 0x35, 0x74, 0xa0, 0xa1, 0x03, 0x8d, 0x24, 0xce, 0xe2, 0xb2, 0xb0, 0x2c, 0xea, 0xfc,
	0xfb, 0x2c, 0x1c, 0xdf, 0x45, 0xcf, 0xcf, 0xd7, 0x71, 0xf2, 0x25, 0xf6, 0x62, 0xa6, 0x11, 0x81,
	0x46, 0x08, 0x8d, 0x50, 0xcf, 0x11, 0x5d, 0x87, 0x55, 0xe1, 0x5e, 0xbb, 0xde, 0x9d, 0x2b, 0x03,
	0x6e, 0x71, 0x24, 0x53, 0x74, 0x15, 0x96, 0x7c, 0xbc, 0x15, 0x18, 0x70, 0x64, 0xa4, 0x44, 0xf7,
	0xa0, 0xdc, 0x33, 0x25, 0xc3, 0xaa, 0x53, 0x47, 0xff, 0x83, 0xbc, 0x15, 0x28, 0x90, 0x91, 0x69,
	0xfa, 0x37, 0xec, 0x8f, 0x39, 0x1d, 0x57, 0xd6, 0x7c, 0xef, 0xca, 0xc7, 0x20, 0x20, 0x33, 0xf4,
	0x10, 0x76, 0xc7, 0x10, 0xdc, 0xb9, 0x41, 0x26, 0x3d, 0xc1, 0xc9, 0x2c, 0xdd, 0x87, 0xed, 0x31,
	0x7f, 0xc5, 0x72, 0xaa, 0xc2, 0x47, 0x32, 0x47, 0x29, 0xac, 0xe5, 0xde, 0x5e, 0xcd, 0x79, 0x7a,
	0x0c, 0x87, 0xf9, 0xbb, 0xba, 0x55, 0x75, 0x98, 0xc5, 0x1d, 0xcf, 0x1d, 0xa9, 0xba, 0x30, 0x5c,
	0x75, 0x08, 0x93, 0xe7, 0x5d, 0xa4, 0x67, 0x70, 0xe2, 0xb9, 0x92, 0xe1, 0x8d, 0xe5, 0x32, 0x3d,
	0xb8, 0x08, 0xa4, 0xa8, 0x31, 0x8b, 0xe3, 0x48, 0xa2, 0x25, 0x55, 0xcc, 0xc3, 0x1b, 0x69, 0x57,
	0x3d, 0xc1, 0x8a, 0x47, 0x04, 0xfa, 0x2f, 0xfc, 0x33, 0xc0, 0x70, 0xdf, 0x72, 0x83, 0xea, 0x78,
	0x4f, 0xcb, 0x8a, 0xab, 0x0c, 0x32, 0xa1, 0xeb, 0x6d, 0xc5, 0xc5, 0x18, 0x22, 0xef, 0x79, 0x85,
	0x96, 0x61, 0x23, 0x10, 0xb6, 0x8d, 0xc8, 0x14, 0x79, 0xaa, 0xfb, 0xba, 0x63, 0x23, 0x29, 0xd3,
	0x65, 0x58, 0xd0, 0x8e, 0x20, 0x20, 0xeb, 0xca, 0xc8, 0x43, 0xa8, 0x32, 0xf0, 0xbe, 0xe6, 0xf8,
	0xc8, 0xc8, 0x86, 0xce, 0x6e, 0xb9, 0x36, 0x56, 0x0b, 0xba, 0xdb, 0x54, 0x5b, 0xef, 0x79, 0x91,
	0x91, 0xad, 0xf3, 0x1f, 0x73, 0x60, 0xfc, 0x56, 0x3e, 0xbc, 0x1d, 0xc6, 0x9d, 0x66, 0xda, 0x4c,
	0x62, 0xde, 0x6e, 0x3e, 0x3d, 0x45, 0x6d, 0xc5, 0x79, 0x2e, 0x25, 0x4d, 0x82, 0xa3, 0xab, 0x70,
	0xdf, 0xb9, 0xba, 0x42, 0x9f, 0x4c, 0xd1, 0x23, 0xd8, 0x1b, 0xdd, 0xa5, 0xac, 0x78, 0x7e, 0x9f,
	0x53, 0x52, 0xa2, 0x07, 0xb0, 0x93, 0xcd, 0x24, 0x6d, 0xcf, 0x75, 0xd1, 0xe6, 0xc8, 0x24, 0xf7,
	0x32, 0x5e, 0xc9, 0x74, 0xa1, 0x10, 0xd5, 0xa8, 0xc8, 0xc8, 0x8c, 0xda, 0x41, 0xa1, 0xcc, 0x3c,
	0xc1, 0x25, 0xde, 0x67, 0xa4, 0x8d, 0xaa, 0xad, 0xd7, 0xc3, 0x80, 0xc8, 0xb9, 0x09, 0xaa, 0x09,
	0xb8, 0xe5, 0x73, 0xad, 0xbc, 0x03, 0xd8, 0x99, 0xa0, 0x2a, 0x64, 0x64, 0x61, 0xf8, 0x63, 0x18,
	0x0e, 0xcf, 0x77, 0x46, 0xca, 0xf4, 0x04, 0x8e, 0x26, 0xc9, 0x2e, 0xaf, 0xb2, 0xa8, 0xe5, 0x34,
	0x09, 0xd4, 0xcf, 0x95, 0x29, 0x73, 0x02, 0xac, 0xd7, 0x11, 0xa8, 0x8e, 0x0a, 0xd4, 0x3b, 0xc8,
	0xb2, 0xac, 0x3b, 0x2a, 0xd4, 0xee, 0x00, 0xb4, 0x42, 0x4f, 0xe1, 0xb8, 0x40, 0x79, 0x12, 0xeb,
	0xe8, 0x72, 0xe9, 0xa3, 0x8d, 0x4e, 0x1d, 0x19, 0xd9, 0xa6, 0x17, 0xf0, 0xdf, 0x98, 0x7e, 0xbd,
	0x4a, 0x0e, 0x67, 0x16, 0xb7, 0x86, 0xd2, 0xae, 0xd2, 0x73, 0x38, 0x7d, 0x0b, 0xde, 0x9b, 0x64,
	0x8d, 0x02, 0xcc, 0xf7, 0x9e, 0x37, 0xd4, 0x87, 0xa0, 0x55, 0x9d, 0x45, 0xf4, 0xb7, 0xfb, 0x17,
	0xdd, 0x85, 0xad, 0x11, 0x85, 0x0f, 0xce, 0xd8, 0xa6, 0xa2, 0xab, 0xc8, 0x27, 0x2b, 0xa2, 0x5a,
	0x71, 0x32, 0xd1, 0x5f, 0x7e, 0x2b, 0xc1, 0xff, 0x8d, 0xa4, 0xf5, 0x87, 0xa7, 0xf6, 0xf2, 0xec,
	0xed, 0x3b, 0x5b, 0x53, 0x37, 0xbc, 0x56, 0xfa, 0x58, 0x7b, 0x6a, 0xa6, 0x9f, 0xba, 0x0f, 0x46,
	0x23, 0x69, 0x99, 0xfd, 0x2a, 0x17, 0x61, 0x53, 0xfd, 0x0c, 0xa2, 0x76, 0x1c, 0x3e, 0x5f, 0xe8,
	0x6b, 0xdf, 0x31, 0x3b, 0xed, 0x86, 0xd9, 0x0a, 0x9b, 0xb1, 0xa9, 0x6d, 0x73, 0xe2, 0x5f, 0xe3,
	0x67, 0xa9, 0xf4, 0x75, 0x7a, 0xc6, 0x12, 0xfc, 0x61, 0x5e, 0x23, 0xdf, 0xfd, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x51, 0x1c, 0x1b, 0x68, 0x62, 0x06, 0x00, 0x00,
}