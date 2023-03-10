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
// source: autonomic/ext/event/state/vehicle_claim_validation.proto

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

type WellKnownVehicleClaimValidationState int32

const (
	WellKnownVehicleClaimValidationState_UNKNOWN_STATE WellKnownVehicleClaimValidationState = 0
	// Request was received by cloud through API
	// Emitted by: cloud
	WellKnownVehicleClaimValidationState_REQUESTED WellKnownVehicleClaimValidationState = 1
	// Request was queued in the cloud, waiting for device
	// to connect.
	// Emitted by: cloud
	// Triggered by: REQUEST_QUEUED_FOR_DELIVERY
	WellKnownVehicleClaimValidationState_REQUEST_DELIVERY_QUEUED WellKnownVehicleClaimValidationState = 2
	// Cloud attempted delivery of request to device
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_TRANSLATION_SUCCEEDED
	WellKnownVehicleClaimValidationState_REQUEST_DELIVERY_IN_PROGRESS WellKnownVehicleClaimValidationState = 3
	// Delivery of request attempt was not acknowledged by device in a timely fashion
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_TIMEOUT_EXCEEDED
	WellKnownVehicleClaimValidationState_REQUEST_DELIVERY_TIMED_OUT WellKnownVehicleClaimValidationState = 4
	// Delivery of request attempt failed due to device disconnect
	// Emitted by: cloud
	// Triggered by: REQUEST_DELIVERY_FAILED
	WellKnownVehicleClaimValidationState_REQUEST_DELIVERY_FAILURE WellKnownVehicleClaimValidationState = 5
	// Request is acknowledged by device, queued, and the device will begin
	// processing the request when the device is ready.
	// Emitted by: device
	// Triggered by: REQUEST_QUEUED_ON_DEVICE
	WellKnownVehicleClaimValidationState_REQUEST_QUEUED WellKnownVehicleClaimValidationState = 6
	// Request is being decrypted and validated by device
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_STARTED
	WellKnownVehicleClaimValidationState_REQUEST_VALIDATION_IN_PROGRESS WellKnownVehicleClaimValidationState = 7
	// Request failed validation or decryption
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_FAILED
	WellKnownVehicleClaimValidationState_REQUEST_VALIDATION_FAILURE WellKnownVehicleClaimValidationState = 8
	// Vehicle claim is in the process of being validated
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED
	WellKnownVehicleClaimValidationState_VEHICLE_CLAIM_VALIDATION_IN_PROGRESS WellKnownVehicleClaimValidationState = 9
	// The device has successfully completed the vehicle claim validation commmand in question.
	// Triggered by: VEHICLE_CLAIM_VALIDATION_SUCCEEDED
	WellKnownVehicleClaimValidationState_SUCCEEDED_ON_DEVICE WellKnownVehicleClaimValidationState = 23
	// The cloud is in the process of validating the event emitted by the device indicating that the
	// command has succeeded.
	// Triggered by: SUCCEEDED_ON_DEVICE_EVENT_RECEIVED
	WellKnownVehicleClaimValidationState_CLOUD_VALIDATION_IN_PROGRESS WellKnownVehicleClaimValidationState = 24
	// Oem cloud attempted delivery of request
	// Emitted by: cloud
	// Triggered by: DEVICE_CONNECTED_TO_CLOUD
	WellKnownVehicleClaimValidationState_OEM_CLOUD_DELIVERY_IN_PROGRESS WellKnownVehicleClaimValidationState = 10
	// Cloud attempted translation of request
	// Emitted by: cloud
	// Triggered by: OEM_CLOUD_DELIVERY_SUCCEEDED
	WellKnownVehicleClaimValidationState_OEM_CLOUD_TRANSLATION_IN_PROGRESS WellKnownVehicleClaimValidationState = 11
	// Cloud failed to validate data provided by the device
	// Emitted by: cloud
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED
	WellKnownVehicleClaimValidationState_CLOUD_VALIDATION_FAILURE WellKnownVehicleClaimValidationState = 12
	// VehicleClaimValidation upload succeeded, and the URI(s) of the archive(s) has been
	// added to the event containing this state.
	// This state is terminal
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED
	WellKnownVehicleClaimValidationState_SUCCESS WellKnownVehicleClaimValidationState = 17
	// VehicleClaimValidation failed. All intermediate failure states should transition to
	// this state.
	// Emitted by: device
	// Triggered by: VEHICLE_CLAIM_VALIDATION_FAILED
	WellKnownVehicleClaimValidationState_FAILURE WellKnownVehicleClaimValidationState = 18
	// VehicleClaimValidation command expired. This event can be emitted by
	// the device or the cloud. In either case, it will cause
	// the command to be removed from the set of in-progress commands
	// on the cloud. However the cloud does not send this event to
	// the device. This state is terminal.
	// Triggered by: EXPIRATION_EXCEEDED
	WellKnownVehicleClaimValidationState_EXPIRED WellKnownVehicleClaimValidationState = 19
	// Command is being removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_ACKNOWLEDGED
	WellKnownVehicleClaimValidationState_CANCELLATION_IN_PROGRESS WellKnownVehicleClaimValidationState = 20
	// Command was successfully removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_FULFILLED
	WellKnownVehicleClaimValidationState_CANCELLED WellKnownVehicleClaimValidationState = 21
)

var WellKnownVehicleClaimValidationState_name = map[int32]string{
	0:  "UNKNOWN_STATE",
	1:  "REQUESTED",
	2:  "REQUEST_DELIVERY_QUEUED",
	3:  "REQUEST_DELIVERY_IN_PROGRESS",
	4:  "REQUEST_DELIVERY_TIMED_OUT",
	5:  "REQUEST_DELIVERY_FAILURE",
	6:  "REQUEST_QUEUED",
	7:  "REQUEST_VALIDATION_IN_PROGRESS",
	8:  "REQUEST_VALIDATION_FAILURE",
	9:  "VEHICLE_CLAIM_VALIDATION_IN_PROGRESS",
	23: "SUCCEEDED_ON_DEVICE",
	24: "CLOUD_VALIDATION_IN_PROGRESS",
	10: "OEM_CLOUD_DELIVERY_IN_PROGRESS",
	11: "OEM_CLOUD_TRANSLATION_IN_PROGRESS",
	12: "CLOUD_VALIDATION_FAILURE",
	17: "SUCCESS",
	18: "FAILURE",
	19: "EXPIRED",
	20: "CANCELLATION_IN_PROGRESS",
	21: "CANCELLED",
}

var WellKnownVehicleClaimValidationState_value = map[string]int32{
	"UNKNOWN_STATE":                        0,
	"REQUESTED":                            1,
	"REQUEST_DELIVERY_QUEUED":              2,
	"REQUEST_DELIVERY_IN_PROGRESS":         3,
	"REQUEST_DELIVERY_TIMED_OUT":           4,
	"REQUEST_DELIVERY_FAILURE":             5,
	"REQUEST_QUEUED":                       6,
	"REQUEST_VALIDATION_IN_PROGRESS":       7,
	"REQUEST_VALIDATION_FAILURE":           8,
	"VEHICLE_CLAIM_VALIDATION_IN_PROGRESS": 9,
	"SUCCEEDED_ON_DEVICE":                  23,
	"CLOUD_VALIDATION_IN_PROGRESS":         24,
	"OEM_CLOUD_DELIVERY_IN_PROGRESS":       10,
	"OEM_CLOUD_TRANSLATION_IN_PROGRESS":    11,
	"CLOUD_VALIDATION_FAILURE":             12,
	"SUCCESS":                              17,
	"FAILURE":                              18,
	"EXPIRED":                              19,
	"CANCELLATION_IN_PROGRESS":             20,
	"CANCELLED":                            21,
}

func (x WellKnownVehicleClaimValidationState) String() string {
	return proto.EnumName(WellKnownVehicleClaimValidationState_name, int32(x))
}

func (WellKnownVehicleClaimValidationState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eea5bdf9410f713, []int{0}
}

// VehicleClaimValidation state transition triggers. See WellKnownVehicleClaimValidationState enum
// for documentation on when these triggers occur.
type WellKnownVehicleClaimValidationTransitionTrigger int32

const (
	WellKnownVehicleClaimValidationTransitionTrigger_UNKNOWN_TRANSITION_TRIGGER                WellKnownVehicleClaimValidationTransitionTrigger = 0
	WellKnownVehicleClaimValidationTransitionTrigger_REQUEST_QUEUED_FOR_DELIVERY               WellKnownVehicleClaimValidationTransitionTrigger = 1
	WellKnownVehicleClaimValidationTransitionTrigger_DEVICE_CONNECTED_TO_CLOUD                 WellKnownVehicleClaimValidationTransitionTrigger = 2
	WellKnownVehicleClaimValidationTransitionTrigger_REQUEST_DELIVERY_FAILED                   WellKnownVehicleClaimValidationTransitionTrigger = 3
	WellKnownVehicleClaimValidationTransitionTrigger_REQUEST_DELIVERY_TIMEOUT_EXCEEDED         WellKnownVehicleClaimValidationTransitionTrigger = 4
	WellKnownVehicleClaimValidationTransitionTrigger_REQUEST_QUEUED_ON_DEVICE                  WellKnownVehicleClaimValidationTransitionTrigger = 5
	WellKnownVehicleClaimValidationTransitionTrigger_REQUEST_VALIDATION_STARTED                WellKnownVehicleClaimValidationTransitionTrigger = 6
	WellKnownVehicleClaimValidationTransitionTrigger_REQUEST_VALIDATION_FAILED                 WellKnownVehicleClaimValidationTransitionTrigger = 7
	WellKnownVehicleClaimValidationTransitionTrigger_REQUEST_VALIDATION_SUCCEEDED              WellKnownVehicleClaimValidationTransitionTrigger = 23
	WellKnownVehicleClaimValidationTransitionTrigger_VEHICLE_CLAIM_VALIDATION_STARTED          WellKnownVehicleClaimValidationTransitionTrigger = 8
	WellKnownVehicleClaimValidationTransitionTrigger_VEHICLE_CLAIM_VALIDATION_SUCCEEDED        WellKnownVehicleClaimValidationTransitionTrigger = 9
	WellKnownVehicleClaimValidationTransitionTrigger_VEHICLE_CLAIM_VALIDATION_FAILED           WellKnownVehicleClaimValidationTransitionTrigger = 10
	WellKnownVehicleClaimValidationTransitionTrigger_OEM_CLOUD_DELIVERY_SUCCEEDED              WellKnownVehicleClaimValidationTransitionTrigger = 11
	WellKnownVehicleClaimValidationTransitionTrigger_OEM_CLOUD_TRANSLATION_SUCCEEDED           WellKnownVehicleClaimValidationTransitionTrigger = 12
	WellKnownVehicleClaimValidationTransitionTrigger_SUCCEEDED_ON_DEVICE_EVENT_RECEIVED        WellKnownVehicleClaimValidationTransitionTrigger = 24
	WellKnownVehicleClaimValidationTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED WellKnownVehicleClaimValidationTransitionTrigger = 13
	WellKnownVehicleClaimValidationTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED    WellKnownVehicleClaimValidationTransitionTrigger = 14
	WellKnownVehicleClaimValidationTransitionTrigger_FAILED                                    WellKnownVehicleClaimValidationTransitionTrigger = 19
	WellKnownVehicleClaimValidationTransitionTrigger_EXPIRATION_EXCEEDED                       WellKnownVehicleClaimValidationTransitionTrigger = 20
	WellKnownVehicleClaimValidationTransitionTrigger_CANCELLATION_REQUESTED                    WellKnownVehicleClaimValidationTransitionTrigger = 21
	WellKnownVehicleClaimValidationTransitionTrigger_CANCELLATION_REQUEST_FULFILLED            WellKnownVehicleClaimValidationTransitionTrigger = 22
)

var WellKnownVehicleClaimValidationTransitionTrigger_name = map[int32]string{
	0:  "UNKNOWN_TRANSITION_TRIGGER",
	1:  "REQUEST_QUEUED_FOR_DELIVERY",
	2:  "DEVICE_CONNECTED_TO_CLOUD",
	3:  "REQUEST_DELIVERY_FAILED",
	4:  "REQUEST_DELIVERY_TIMEOUT_EXCEEDED",
	5:  "REQUEST_QUEUED_ON_DEVICE",
	6:  "REQUEST_VALIDATION_STARTED",
	7:  "REQUEST_VALIDATION_FAILED",
	23: "REQUEST_VALIDATION_SUCCEEDED",
	8:  "VEHICLE_CLAIM_VALIDATION_STARTED",
	9:  "VEHICLE_CLAIM_VALIDATION_SUCCEEDED",
	10: "VEHICLE_CLAIM_VALIDATION_FAILED",
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

var WellKnownVehicleClaimValidationTransitionTrigger_value = map[string]int32{
	"UNKNOWN_TRANSITION_TRIGGER":                0,
	"REQUEST_QUEUED_FOR_DELIVERY":               1,
	"DEVICE_CONNECTED_TO_CLOUD":                 2,
	"REQUEST_DELIVERY_FAILED":                   3,
	"REQUEST_DELIVERY_TIMEOUT_EXCEEDED":         4,
	"REQUEST_QUEUED_ON_DEVICE":                  5,
	"REQUEST_VALIDATION_STARTED":                6,
	"REQUEST_VALIDATION_FAILED":                 7,
	"REQUEST_VALIDATION_SUCCEEDED":              23,
	"VEHICLE_CLAIM_VALIDATION_STARTED":          8,
	"VEHICLE_CLAIM_VALIDATION_SUCCEEDED":        9,
	"VEHICLE_CLAIM_VALIDATION_FAILED":           10,
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

func (x WellKnownVehicleClaimValidationTransitionTrigger) String() string {
	return proto.EnumName(WellKnownVehicleClaimValidationTransitionTrigger_name, int32(x))
}

func (WellKnownVehicleClaimValidationTransitionTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eea5bdf9410f713, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.event.state.vehicleclaimvalidation.WellKnownVehicleClaimValidationState", WellKnownVehicleClaimValidationState_name, WellKnownVehicleClaimValidationState_value)
	proto.RegisterEnum("autonomic.ext.event.state.vehicleclaimvalidation.WellKnownVehicleClaimValidationTransitionTrigger", WellKnownVehicleClaimValidationTransitionTrigger_name, WellKnownVehicleClaimValidationTransitionTrigger_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/event/state/vehicle_claim_validation.proto", fileDescriptor_3eea5bdf9410f713)
}

var fileDescriptor_3eea5bdf9410f713 = []byte{
	// 696 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xd9, 0x4e, 0xe3, 0x48,
	0x14, 0x25, 0xec, 0x5c, 0x16, 0x15, 0xc5, 0x12, 0x76, 0x18, 0x86, 0x41, 0x03, 0x12, 0x31, 0xd2,
	0xcc, 0xc3, 0xbc, 0x1a, 0xd7, 0x0d, 0x63, 0x61, 0xec, 0x60, 0x97, 0x0d, 0x33, 0x2f, 0x25, 0x93,
	0xb6, 0xc0, 0x52, 0x62, 0xb7, 0x12, 0x43, 0xf3, 0x3d, 0xfd, 0x21, 0xdd, 0x9f, 0xd5, 0xfd, 0xd8,
	0x2a, 0xc7, 0xce, 0xa2, 0xd8, 0xa2, 0xfb, 0xcd, 0xe5, 0x7b, 0xce, 0x5d, 0x4f, 0xdd, 0x82, 0x7f,
	0xfc, 0x97, 0x24, 0x8e, 0xe2, 0x76, 0xd8, 0x54, 0x82, 0xb7, 0x44, 0x09, 0x5e, 0x83, 0x28, 0x51,
	0xba, 0x89, 0x9f, 0x04, 0xca, 0x6b, 0xf0, 0x1c, 0x36, 0x5b, 0x81, 0x68, 0xb6, 0xfc, 0xb0, 0x2d,
	0x5e, 0xfd, 0x56, 0xf8, 0xc1, 0x4f, 0xc2, 0x38, 0xaa, 0x7d, 0xec, 0xc4, 0x49, 0x4c, 0x2f, 0xfb,
	0xcc, 0x5a, 0xf0, 0x96, 0xd4, 0x52, 0x66, 0x2d, 0x65, 0xd6, 0x32, 0x66, 0x4a, 0x1c, 0xf0, 0xce,
	0xbf, 0x4e, 0xc3, 0xc9, 0x7d, 0xd0, 0x6a, 0xdd, 0x44, 0xf1, 0xa7, 0xc8, 0xeb, 0x61, 0x34, 0x89,
	0xf1, 0xfa, 0x18, 0x47, 0xf2, 0xe9, 0x2a, 0x2c, 0xbb, 0xe6, 0x8d, 0x69, 0xdd, 0x9b, 0xc2, 0xe1,
	0x2a, 0x47, 0x32, 0x41, 0x97, 0x61, 0xc1, 0xc6, 0x3b, 0x17, 0x1d, 0x8e, 0x8c, 0x54, 0xe8, 0x2e,
	0x54, 0xb3, 0xa3, 0x60, 0x68, 0xe8, 0x1e, 0xda, 0xff, 0x89, 0x3b, 0x17, 0x5d, 0x64, 0x64, 0x92,
	0x1e, 0xc1, 0xde, 0x98, 0x51, 0x37, 0x45, 0xc3, 0xb6, 0xae, 0x6d, 0x74, 0x1c, 0x32, 0x45, 0x0f,
	0x60, 0x67, 0x0c, 0xc1, 0xf5, 0x5b, 0x64, 0xc2, 0x72, 0x39, 0x99, 0xa6, 0x7b, 0xb0, 0x35, 0x66,
	0xaf, 0xab, 0xba, 0xe1, 0xda, 0x48, 0x66, 0x28, 0x85, 0x95, 0xdc, 0x9a, 0xc5, 0x9c, 0xa5, 0xc7,
	0x70, 0x90, 0xff, 0xf3, 0x54, 0x43, 0x67, 0x2a, 0xd7, 0x2d, 0x73, 0x24, 0xea, 0xdc, 0x70, 0xd4,
	0x21, 0x4c, 0xee, 0x77, 0x9e, 0xfe, 0x09, 0x27, 0x1e, 0xfe, 0xab, 0x6b, 0x06, 0x0a, 0xcd, 0x50,
	0xf5, 0xdb, 0x32, 0x4f, 0x0b, 0xb4, 0x0a, 0x6b, 0x8e, 0xab, 0x69, 0x88, 0x4c, 0xa6, 0x6c, 0x0a,
	0x86, 0x9e, 0xae, 0x21, 0xa9, 0xca, 0xd2, 0x35, 0xc3, 0x72, 0x59, 0x19, 0x75, 0x4b, 0x26, 0x6a,
	0xe1, 0xad, 0xe8, 0xa1, 0x0a, 0xdb, 0x03, 0xf4, 0x0f, 0xf8, 0x6d, 0x80, 0xe1, 0xb6, 0x6a, 0x3a,
	0xc6, 0xb8, 0xab, 0x45, 0xd9, 0xa5, 0xb1, 0x60, 0x79, 0x35, 0x4b, 0x74, 0x11, 0xe6, 0xd2, 0x1c,
	0x1d, 0x87, 0xac, 0xca, 0x43, 0x6e, 0xa1, 0xf2, 0x80, 0x0f, 0x0d, 0xdd, 0x46, 0x46, 0xd6, 0x52,
	0x27, 0xaa, 0xa9, 0xa1, 0x51, 0x10, 0x62, 0x5d, 0x8e, 0x3d, 0xb3, 0x22, 0x23, 0x1b, 0xe7, 0xdf,
	0x66, 0xe0, 0xf2, 0x1d, 0x05, 0xf1, 0x8e, 0x1f, 0x75, 0xc3, 0xde, 0x57, 0xf8, 0xf4, 0x14, 0x74,
	0x64, 0xdb, 0x73, 0x35, 0xa5, 0xb5, 0xe8, 0x69, 0x1c, 0x6e, 0xeb, 0xd7, 0xd7, 0x68, 0x93, 0x09,
	0x7a, 0x08, 0xbb, 0xa3, 0xe3, 0x14, 0x75, 0xcb, 0xee, 0xb7, 0x86, 0x54, 0xe8, 0x3e, 0x6c, 0xf7,
	0x1a, 0x2c, 0x34, 0xcb, 0x34, 0x51, 0xe3, 0xc8, 0x04, 0xb7, 0x7a, 0xed, 0x21, 0x93, 0x85, 0x5a,
	0x94, 0xc5, 0x22, 0x23, 0x53, 0xb2, 0x95, 0x85, 0x4a, 0xb3, 0x5c, 0x2e, 0xf0, 0xa1, 0x37, 0xc1,
	0x51, 0xc1, 0x65, 0x39, 0x0c, 0xa6, 0x3a, 0x53, 0x22, 0x1c, 0x87, 0xab, 0x36, 0x4f, 0xc5, 0xb7,
	0x0f, 0xdb, 0x25, 0xc2, 0x42, 0x46, 0xe6, 0x86, 0xef, 0xc3, 0x30, 0x3d, 0x17, 0x10, 0xa9, 0xd2,
	0x13, 0x38, 0x2a, 0x55, 0x5e, 0x1e, 0x66, 0x9e, 0x9e, 0xc2, 0x71, 0x39, 0xaa, 0xef, 0x6d, 0x81,
	0xfe, 0x0e, 0x87, 0xa5, 0xb8, 0x2c, 0x29, 0x90, 0x49, 0x15, 0xe8, 0x70, 0xe0, 0x66, 0x51, 0xba,
	0x29, 0x56, 0xe1, 0x00, 0xb4, 0x24, 0x73, 0x2a, 0xb8, 0x09, 0x02, 0x3d, 0x34, 0xb9, 0xb0, 0x51,
	0x43, 0xdd, 0x43, 0x46, 0xb6, 0xe8, 0x05, 0x9c, 0x8d, 0x69, 0xd5, 0xaa, 0xe7, 0x70, 0xa6, 0x72,
	0x75, 0xc8, 0xed, 0x32, 0x3d, 0x87, 0xd3, 0xf7, 0xe0, 0x59, 0x25, 0x2b, 0x14, 0x60, 0x36, 0xfb,
	0x5e, 0x93, 0x17, 0x33, 0x95, 0x76, 0x8f, 0xd1, 0x1f, 0xf0, 0x3a, 0xdd, 0x81, 0xcd, 0x11, 0x99,
	0x0f, 0x96, 0xd9, 0x86, 0xbc, 0x92, 0x45, 0x36, 0x51, 0x77, 0x8d, 0xba, 0x9e, 0x2a, 0x7f, 0xf3,
	0xea, 0x4b, 0x05, 0xfe, 0x6e, 0xc6, 0xed, 0xda, 0xaf, 0x2e, 0xdd, 0xab, 0xb3, 0x9f, 0xd9, 0xb8,
	0x0d, 0xb9, 0xd1, 0x1b, 0x95, 0xff, 0x1b, 0x4f, 0x61, 0xf2, 0xfc, 0xf2, 0x58, 0x6b, 0xc6, 0x6d,
	0xa5, 0x1f, 0xe9, 0xc2, 0x0f, 0xe5, 0xdb, 0x10, 0x74, 0x22, 0xbf, 0x75, 0x91, 0xee, 0xfe, 0xae,
	0xd2, 0xed, 0x34, 0x95, 0xb6, 0x1f, 0x46, 0x4a, 0x7a, 0x56, 0x4a, 0x1f, 0x91, 0xef, 0x95, 0xca,
	0xe7, 0xc9, 0x29, 0xd5, 0xe5, 0x8f, 0xb3, 0x29, 0xf2, 0xaf, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0x22, 0x23, 0x47, 0x71, 0x06, 0x00, 0x00,
}
