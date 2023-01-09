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
// source: autonomic/ext/event/state/deployment.proto

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

// Those well-known from_state and to_state fields in StateTransition events
type WellKnownDeploymentState int32

const (
	WellKnownDeploymentState_UNKNOWN_STATE WellKnownDeploymentState = 0
	// Deployment request was received by cloud through API
	// Emitted by: cloud
	WellKnownDeploymentState_REQUESTED WellKnownDeploymentState = 1
	// Request was queued in the cloud, waiting for device
	// to connect.
	// Emitted by: cloud
	// Triggered by: REQUEST_QUEUED_FOR_DELIVERY
	WellKnownDeploymentState_REQUEST_DELIVERY_QUEUED WellKnownDeploymentState = 2
	// Cloud attempted delivery of request to device
	// Emitted by: cloud
	// Triggered by: DEVICE_CONNECTED_TO_CLOUD
	WellKnownDeploymentState_REQUEST_DELIVERY_IN_PROGRESS WellKnownDeploymentState = 3
	// Delivery attempt was not acknowledged by device
	// Emitted by: cloud
	// Triggered by: DELIVERY_TIMEOUT_EXCEEDED
	WellKnownDeploymentState_REQUEST_DELIVERY_TIMED_OUT WellKnownDeploymentState = 4
	// Delivery of request attempt failed due to device disconnect
	// Emitted by: cloud
	// Triggered by: DELIVERY_FAILED
	WellKnownDeploymentState_REQUEST_DELIVERY_FAILURE WellKnownDeploymentState = 5
	// Cancellation removes the request from the list of
	// in-progress commands in the cloud, but does not
	// cancel the operation on the device. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUESTED
	WellKnownDeploymentState_CANCELLATION_IN_PROGRESS WellKnownDeploymentState = 6
	// Command was successfully removed from set of in-progress
	// commands from the cloud. Emitted by cloud.
	// Triggered by: CANCELLATION_REQUEST_FULFILLED
	WellKnownDeploymentState_CANCELLED WellKnownDeploymentState = 7
	// Command expired before reaching a terminal state.
	// This event can be emitted by the device or the cloud.
	// In either case, it will cause the command to be removed
	// from the set of in-progress commands on the cloud.
	// However the cloud does not send this event to the device.
	// This state is terminal.
	// Emitted by: cloud or device
	// Triggered by: EXPIRATION_EXCEEDED
	WellKnownDeploymentState_EXPIRED WellKnownDeploymentState = 30
	// Request is acknowledged by device, queued, and the device will begin
	// processing the request when the device is ready.
	// Emitted by: device
	// Triggered by: REQUEST_QUEUED_ON_DEVICE
	WellKnownDeploymentState_REQUEST_QUEUED WellKnownDeploymentState = 10
	// Request is being decrypted and validated by device
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_STARTED
	WellKnownDeploymentState_REQUEST_VALIDATION_IN_PROGRESS WellKnownDeploymentState = 9
	// Request failed validation or decryption
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_FAILED
	WellKnownDeploymentState_REQUEST_VALIDATION_FAILURE WellKnownDeploymentState = 8
	// Request is acknowledged by device, queued, and the device may begin
	// downloading artifacts when it is ready to begin the deployment
	// process.
	// Emitted by: device
	// Triggered by: REQUEST_VALIDATION_SUCCEEDED
	WellKnownDeploymentState_ARTIFACT_RETRIEVAL_QUEUED WellKnownDeploymentState = 60
	// Device has begun actively downloading artifacts from cloud
	// Emitted by: device
	// Triggered by: ARTIFACT_RETRIEVAL_STARTED
	WellKnownDeploymentState_ARTIFACT_RETRIEVAL_IN_PROGRESS WellKnownDeploymentState = 61
	// Device encountered an error downlaoding artifacts. It may either re-try at
	// its discretion or transition into a terminal failure.
	// Emitted by: device
	// Triggered by:  ARTIFACT_RETRIEVAL_FAILED
	WellKnownDeploymentState_ARTIFACT_RETRIEVAL_FAILURE WellKnownDeploymentState = 62
	// Artifact retrieval is successful, and the device has queued installation
	// to begin at its discretion.
	// Emitted by: device
	// Triggered by: ARTIFACT_RETRIVAL_SUCCEDED
	WellKnownDeploymentState_INSTALLATION_QUEUED WellKnownDeploymentState = 63
	// Device has begin installation of artifacts
	// Emitted by: device
	// Triggered by: INSTALLER_STARTED_ON_DEVICE
	WellKnownDeploymentState_INSTALLATION_IN_PROGRESS WellKnownDeploymentState = 64
	// Installation failed. The device may retry or tansition into a terminal failed state
	// Emitted by: device
	// Triggered by: INSTALLER_FAILED_ON_DEVICE
	WellKnownDeploymentState_INSTALLATION_FAILURE WellKnownDeploymentState = 65
	// Post-installation tasks are in-progress. This state covers any work that may be needed
	// after installation, before final success or failure
	// Emitted by: device
	// Triggered by: STARTED_ON_DEVICE
	WellKnownDeploymentState_DEPLOYING WellKnownDeploymentState = 66
	// A non-fatal deployment error has occurred on the device. This error may transition into
	// any other state on the device, depending on what caused the failure. This state allows
	// OEM Developers to easily identify non-fatal issues with OTA. OEM developers may attach
	// additional oem-specific context to this state in order to articulate the root cause
	// information of the failure.
	WellKnownDeploymentState_NON_FATAL_DEVICE_ERROR_STATE WellKnownDeploymentState = 22
	// The device has successfully completed the deployment in question.
	// Triggered by: DEPLOYMENT_SUCCEEDED_ON_DEVICE
	WellKnownDeploymentState_SUCCEEDED_ON_DEVICE WellKnownDeploymentState = 23
	// The cloud is in the process of validating the event emitted by the device indicating that the
	// command has succeeded.
	// Triggered by: SUCCEEDED_ON_DEVICE_EVENT_RECEIVED
	WellKnownDeploymentState_CLOUD_VALIDATION_IN_PROGRESS WellKnownDeploymentState = 24
	// Cloud failed to validate data provided by the device
	// Emitted by: cloud
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED
	WellKnownDeploymentState_CLOUD_VALIDATION_FAILURE WellKnownDeploymentState = 12
	// Deployment was successful on device. This state is terminal.
	// Emitted by: device
	// Triggered by: CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED
	WellKnownDeploymentState_SUCCESS WellKnownDeploymentState = 67
	// Deployment failed. This state is reachable from all intermediate failed states, if the
	// device determines that the failure is non-recoverable. This state is terminal.
	// Emitted by: device
	// Triggered by: DEPLOYMENT_FAILED_ON_DEVICE
	WellKnownDeploymentState_FAILURE WellKnownDeploymentState = 68
)

var WellKnownDeploymentState_name = map[int32]string{
	0:  "UNKNOWN_STATE",
	1:  "REQUESTED",
	2:  "REQUEST_DELIVERY_QUEUED",
	3:  "REQUEST_DELIVERY_IN_PROGRESS",
	4:  "REQUEST_DELIVERY_TIMED_OUT",
	5:  "REQUEST_DELIVERY_FAILURE",
	6:  "CANCELLATION_IN_PROGRESS",
	7:  "CANCELLED",
	30: "EXPIRED",
	10: "REQUEST_QUEUED",
	9:  "REQUEST_VALIDATION_IN_PROGRESS",
	8:  "REQUEST_VALIDATION_FAILURE",
	60: "ARTIFACT_RETRIEVAL_QUEUED",
	61: "ARTIFACT_RETRIEVAL_IN_PROGRESS",
	62: "ARTIFACT_RETRIEVAL_FAILURE",
	63: "INSTALLATION_QUEUED",
	64: "INSTALLATION_IN_PROGRESS",
	65: "INSTALLATION_FAILURE",
	66: "DEPLOYING",
	22: "NON_FATAL_DEVICE_ERROR_STATE",
	23: "SUCCEEDED_ON_DEVICE",
	24: "CLOUD_VALIDATION_IN_PROGRESS",
	12: "CLOUD_VALIDATION_FAILURE",
	67: "SUCCESS",
	68: "FAILURE",
}

var WellKnownDeploymentState_value = map[string]int32{
	"UNKNOWN_STATE":                  0,
	"REQUESTED":                      1,
	"REQUEST_DELIVERY_QUEUED":        2,
	"REQUEST_DELIVERY_IN_PROGRESS":   3,
	"REQUEST_DELIVERY_TIMED_OUT":     4,
	"REQUEST_DELIVERY_FAILURE":       5,
	"CANCELLATION_IN_PROGRESS":       6,
	"CANCELLED":                      7,
	"EXPIRED":                        30,
	"REQUEST_QUEUED":                 10,
	"REQUEST_VALIDATION_IN_PROGRESS": 9,
	"REQUEST_VALIDATION_FAILURE":     8,
	"ARTIFACT_RETRIEVAL_QUEUED":      60,
	"ARTIFACT_RETRIEVAL_IN_PROGRESS": 61,
	"ARTIFACT_RETRIEVAL_FAILURE":     62,
	"INSTALLATION_QUEUED":            63,
	"INSTALLATION_IN_PROGRESS":       64,
	"INSTALLATION_FAILURE":           65,
	"DEPLOYING":                      66,
	"NON_FATAL_DEVICE_ERROR_STATE":   22,
	"SUCCEEDED_ON_DEVICE":            23,
	"CLOUD_VALIDATION_IN_PROGRESS":   24,
	"CLOUD_VALIDATION_FAILURE":       12,
	"SUCCESS":                        67,
	"FAILURE":                        68,
}

func (x WellKnownDeploymentState) String() string {
	return proto.EnumName(WellKnownDeploymentState_name, int32(x))
}

func (WellKnownDeploymentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb85fcfacee82825, []int{0}
}

// Those well-Known trigger field values in StateTransition events
// See WellKnownDeploymentState enum for information as to when these triggers are used
type WellKnownDeploymentTransitionTrigger int32

const (
	WellKnownDeploymentTransitionTrigger_UNKNOWN_TRIGGER                           WellKnownDeploymentTransitionTrigger = 0
	WellKnownDeploymentTransitionTrigger_REQUEST_QUEUED_FOR_DELIVERY               WellKnownDeploymentTransitionTrigger = 1
	WellKnownDeploymentTransitionTrigger_DEVICE_CONNECTED_TO_CLOUD                 WellKnownDeploymentTransitionTrigger = 2
	WellKnownDeploymentTransitionTrigger_REQUEST_DELIVERY_FAILED                   WellKnownDeploymentTransitionTrigger = 3
	WellKnownDeploymentTransitionTrigger_REQUEST_DELIVERY_TIMEOUT_EXCEEDED         WellKnownDeploymentTransitionTrigger = 4
	WellKnownDeploymentTransitionTrigger_CANCELLATION_REQUEST_RECEIVED             WellKnownDeploymentTransitionTrigger = 5
	WellKnownDeploymentTransitionTrigger_CANCELLATION_REQUEST_FULFILLED            WellKnownDeploymentTransitionTrigger = 6
	WellKnownDeploymentTransitionTrigger_EXPIRATION_EXCEEDED                       WellKnownDeploymentTransitionTrigger = 30
	WellKnownDeploymentTransitionTrigger_REQUEST_QUEUED_ON_DEVICE                  WellKnownDeploymentTransitionTrigger = 60
	WellKnownDeploymentTransitionTrigger_REQUEST_VALIDATION_FAILED                 WellKnownDeploymentTransitionTrigger = 7
	WellKnownDeploymentTransitionTrigger_REQUEST_VALIDATION_STARTED                WellKnownDeploymentTransitionTrigger = 8
	WellKnownDeploymentTransitionTrigger_REQUEST_VALIDATION_SUCCEEDED              WellKnownDeploymentTransitionTrigger = 23
	WellKnownDeploymentTransitionTrigger_NON_FATAL_DEVICE_ERROR_OCCURRED           WellKnownDeploymentTransitionTrigger = 22
	WellKnownDeploymentTransitionTrigger_STARTED_ON_DEVICE                         WellKnownDeploymentTransitionTrigger = 61
	WellKnownDeploymentTransitionTrigger_DEPLOYMENT_SUCCEEDED_ON_DEVICE            WellKnownDeploymentTransitionTrigger = 62
	WellKnownDeploymentTransitionTrigger_DEPLOYMENT_FAILED_ON_DEVICE               WellKnownDeploymentTransitionTrigger = 63
	WellKnownDeploymentTransitionTrigger_INSTALLER_STARTED_ON_DEVICE               WellKnownDeploymentTransitionTrigger = 64
	WellKnownDeploymentTransitionTrigger_INSTALLER_FAILED_ON_DEVICE                WellKnownDeploymentTransitionTrigger = 65
	WellKnownDeploymentTransitionTrigger_ARTIFACT_RETRIEVAL_STARTED                WellKnownDeploymentTransitionTrigger = 80
	WellKnownDeploymentTransitionTrigger_ARTIFACT_RETRIEVAL_SUCCEEDED              WellKnownDeploymentTransitionTrigger = 81
	WellKnownDeploymentTransitionTrigger_ARTIFACT_RETRIEVAL_FAILED                 WellKnownDeploymentTransitionTrigger = 82
	WellKnownDeploymentTransitionTrigger_SUCCEEDED_ON_DEVICE_EVENT_RECEIVED        WellKnownDeploymentTransitionTrigger = 24
	WellKnownDeploymentTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED WellKnownDeploymentTransitionTrigger = 13
	WellKnownDeploymentTransitionTrigger_CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED    WellKnownDeploymentTransitionTrigger = 14
)

var WellKnownDeploymentTransitionTrigger_name = map[int32]string{
	0:  "UNKNOWN_TRIGGER",
	1:  "REQUEST_QUEUED_FOR_DELIVERY",
	2:  "DEVICE_CONNECTED_TO_CLOUD",
	3:  "REQUEST_DELIVERY_FAILED",
	4:  "REQUEST_DELIVERY_TIMEOUT_EXCEEDED",
	5:  "CANCELLATION_REQUEST_RECEIVED",
	6:  "CANCELLATION_REQUEST_FULFILLED",
	30: "EXPIRATION_EXCEEDED",
	60: "REQUEST_QUEUED_ON_DEVICE",
	7:  "REQUEST_VALIDATION_FAILED",
	8:  "REQUEST_VALIDATION_STARTED",
	23: "REQUEST_VALIDATION_SUCCEEDED",
	22: "NON_FATAL_DEVICE_ERROR_OCCURRED",
	61: "STARTED_ON_DEVICE",
	62: "DEPLOYMENT_SUCCEEDED_ON_DEVICE",
	63: "DEPLOYMENT_FAILED_ON_DEVICE",
	64: "INSTALLER_STARTED_ON_DEVICE",
	65: "INSTALLER_FAILED_ON_DEVICE",
	80: "ARTIFACT_RETRIEVAL_STARTED",
	81: "ARTIFACT_RETRIEVAL_SUCCEEDED",
	82: "ARTIFACT_RETRIEVAL_FAILED",
	24: "SUCCEEDED_ON_DEVICE_EVENT_RECEIVED",
	13: "CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED",
	14: "CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED",
}

var WellKnownDeploymentTransitionTrigger_value = map[string]int32{
	"UNKNOWN_TRIGGER":                           0,
	"REQUEST_QUEUED_FOR_DELIVERY":               1,
	"DEVICE_CONNECTED_TO_CLOUD":                 2,
	"REQUEST_DELIVERY_FAILED":                   3,
	"REQUEST_DELIVERY_TIMEOUT_EXCEEDED":         4,
	"CANCELLATION_REQUEST_RECEIVED":             5,
	"CANCELLATION_REQUEST_FULFILLED":            6,
	"EXPIRATION_EXCEEDED":                       30,
	"REQUEST_QUEUED_ON_DEVICE":                  60,
	"REQUEST_VALIDATION_FAILED":                 7,
	"REQUEST_VALIDATION_STARTED":                8,
	"REQUEST_VALIDATION_SUCCEEDED":              23,
	"NON_FATAL_DEVICE_ERROR_OCCURRED":           22,
	"STARTED_ON_DEVICE":                         61,
	"DEPLOYMENT_SUCCEEDED_ON_DEVICE":            62,
	"DEPLOYMENT_FAILED_ON_DEVICE":               63,
	"INSTALLER_STARTED_ON_DEVICE":               64,
	"INSTALLER_FAILED_ON_DEVICE":                65,
	"ARTIFACT_RETRIEVAL_STARTED":                80,
	"ARTIFACT_RETRIEVAL_SUCCEEDED":              81,
	"ARTIFACT_RETRIEVAL_FAILED":                 82,
	"SUCCEEDED_ON_DEVICE_EVENT_RECEIVED":        24,
	"CLOUD_VALIDATION_OF_DEVICE_DATA_SUCCEEDED": 13,
	"CLOUD_VALIDATION_OF_DEVICE_DATA_FAILED":    14,
}

func (x WellKnownDeploymentTransitionTrigger) String() string {
	return proto.EnumName(WellKnownDeploymentTransitionTrigger_name, int32(x))
}

func (WellKnownDeploymentTransitionTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb85fcfacee82825, []int{1}
}

func init() {
	proto.RegisterEnum("autonomic.ext.event.state.deployment.WellKnownDeploymentState", WellKnownDeploymentState_name, WellKnownDeploymentState_value)
	proto.RegisterEnum("autonomic.ext.event.state.deployment.WellKnownDeploymentTransitionTrigger", WellKnownDeploymentTransitionTrigger_name, WellKnownDeploymentTransitionTrigger_value)
}

func init() {
	proto.RegisterFile("autonomic/ext/event/state/deployment.proto", fileDescriptor_bb85fcfacee82825)
}

var fileDescriptor_bb85fcfacee82825 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0x53, 0xd3, 0x40,
	0x14, 0xa5, 0x7c, 0x14, 0x58, 0x05, 0x97, 0x45, 0x69, 0x15, 0x28, 0x82, 0xc8, 0x68, 0x67, 0x68,
	0x1e, 0x7c, 0xe5, 0x6b, 0xc9, 0xde, 0x76, 0x32, 0x84, 0xa4, 0x6c, 0x36, 0x05, 0x7c, 0xd9, 0x29,
	0x35, 0x83, 0x9d, 0x69, 0x13, 0xa6, 0x0d, 0x8a, 0x7f, 0x47, 0x67, 0xfc, 0x25, 0xfe, 0x28, 0x1f,
	0x9d, 0x4d, 0x93, 0x92, 0xda, 0x74, 0xf4, 0x71, 0x73, 0x4f, 0xce, 0x3d, 0xf7, 0xde, 0x73, 0xe7,
	0xa2, 0x72, 0xf3, 0x3e, 0x0c, 0xfc, 0xa0, 0xdb, 0x6e, 0x69, 0xde, 0x43, 0xa8, 0x79, 0x5f, 0x3c,
	0x3f, 0xd4, 0xfa, 0x61, 0x33, 0xf4, 0xb4, 0x4f, 0xde, 0x5d, 0x27, 0xf8, 0xd6, 0xf5, 0xfc, 0xb0,
	0x72, 0xd7, 0x0b, 0xc2, 0x80, 0xec, 0x0e, 0xb1, 0x15, 0xef, 0x21, 0xac, 0x44, 0xd8, 0x4a, 0x84,
	0xad, 0x3c, 0x62, 0xcb, 0x3f, 0xe6, 0x50, 0xf1, 0xd2, 0xeb, 0x74, 0xce, 0xfc, 0xe0, 0xab, 0xcf,
	0x86, 0xdf, 0x1d, 0x85, 0x23, 0x2b, 0x68, 0xc9, 0xb5, 0xce, 0x2c, 0xfb, 0xd2, 0x92, 0x8e, 0xa0,
	0x02, 0xf0, 0x14, 0x59, 0x42, 0x8b, 0x1c, 0x2e, 0x5c, 0x70, 0x04, 0x30, 0x9c, 0x23, 0xeb, 0xa8,
	0x10, 0x3f, 0x25, 0x03, 0xd3, 0x68, 0x00, 0xbf, 0x96, 0x17, 0x2e, 0xb8, 0xc0, 0xf0, 0x34, 0x79,
	0x8d, 0x36, 0xc6, 0x82, 0x86, 0x25, 0xeb, 0xdc, 0xae, 0x71, 0x70, 0x1c, 0x3c, 0x43, 0x4a, 0xe8,
	0xd5, 0x18, 0x42, 0x18, 0xe7, 0xc0, 0xa4, 0xed, 0x0a, 0x3c, 0x4b, 0x36, 0x50, 0x71, 0x2c, 0x5e,
	0xa5, 0x86, 0xe9, 0x72, 0xc0, 0x73, 0x2a, 0xaa, 0x53, 0x4b, 0x07, 0xd3, 0xa4, 0xc2, 0xb0, 0xad,
	0x11, 0xee, 0xbc, 0x52, 0x1a, 0x47, 0x81, 0xe1, 0x79, 0xf2, 0x04, 0xcd, 0xc3, 0x55, 0xdd, 0xe0,
	0xc0, 0x70, 0x89, 0x10, 0xb4, 0x9c, 0xf0, 0xc6, 0x6a, 0x11, 0xd9, 0x41, 0xa5, 0xe4, 0x5b, 0x83,
	0x9a, 0x06, 0x1b, 0xe7, 0x5c, 0x4c, 0xeb, 0x4d, 0x61, 0x12, 0x45, 0x0b, 0x64, 0x13, 0xbd, 0xa4,
	0x5c, 0x18, 0x55, 0xaa, 0x0b, 0xc9, 0x41, 0x70, 0x03, 0x1a, 0xd4, 0x4c, 0x52, 0x1c, 0xa8, 0x14,
	0x19, 0xe1, 0x74, 0x8a, 0x43, 0x95, 0x22, 0x03, 0x93, 0xa4, 0x38, 0x22, 0x05, 0xb4, 0x6a, 0x58,
	0x8e, 0xa0, 0x49, 0xd1, 0x31, 0xf9, 0xb1, 0xea, 0xc6, 0x48, 0x20, 0x4d, 0x7b, 0x42, 0x8a, 0xe8,
	0xf9, 0x48, 0x34, 0x21, 0xa4, 0xaa, 0x4f, 0x0c, 0xea, 0xa6, 0x7d, 0x6d, 0x58, 0x35, 0x7c, 0xaa,
	0x86, 0x66, 0x45, 0x71, 0x41, 0x4d, 0xc9, 0xa0, 0x61, 0xe8, 0x20, 0x81, 0x73, 0x9b, 0xc7, 0x16,
	0x58, 0x53, 0x0a, 0x1c, 0x57, 0xd7, 0x01, 0x98, 0x9a, 0x93, 0x15, 0x83, 0x70, 0x41, 0xfd, 0xaa,
	0x9b, 0xb6, 0xcb, 0x26, 0xf5, 0xaf, 0x18, 0x4d, 0xec, 0x6f, 0x44, 0xa2, 0xe4, 0xa9, 0x1a, 0x51,
	0x44, 0xec, 0x38, 0x58, 0x57, 0x8f, 0x24, 0xc2, 0xca, 0xbf, 0xf2, 0x68, 0x37, 0xc3, 0xa5, 0xa2,
	0xd7, 0xf4, 0xfb, 0xed, 0xb0, 0x1d, 0xf8, 0xa2, 0xd7, 0xbe, 0xbd, 0xf5, 0x7a, 0x64, 0x15, 0x3d,
	0x4b, 0x1c, 0x2b, 0xb8, 0x51, 0xab, 0x01, 0xc7, 0x53, 0x64, 0x0b, 0xad, 0x8f, 0x4e, 0x5b, 0x56,
	0x6d, 0x3e, 0x34, 0x14, 0xce, 0xa9, 0xb1, 0xc5, 0x95, 0xea, 0xb6, 0x65, 0x81, 0x2e, 0x80, 0x49,
	0x61, 0xcb, 0x48, 0x2a, 0x9e, 0xce, 0x34, 0xb9, 0xd2, 0x06, 0x0c, 0xcf, 0x90, 0xb7, 0x68, 0x3b,
	0xd3, 0xc2, 0xb6, 0x2b, 0x24, 0x5c, 0x0d, 0xba, 0x84, 0x67, 0xc9, 0x36, 0xda, 0x1c, 0xf1, 0x6a,
	0xf2, 0x0f, 0x07, 0x1d, 0x8c, 0x06, 0x30, 0x3c, 0xa7, 0xdc, 0x91, 0x09, 0xa9, 0xba, 0x66, 0xd5,
	0x88, 0x5c, 0x9c, 0x57, 0xbd, 0x8f, 0x5c, 0x3c, 0x40, 0x0c, 0xf9, 0x4b, 0xe9, 0x4d, 0x89, 0x6b,
	0x7c, 0x9c, 0xcc, 0x81, 0x2a, 0x70, 0x82, 0x6f, 0xa3, 0xdd, 0xc8, 0xb6, 0xb5, 0x23, 0x28, 0x57,
	0x5b, 0xbe, 0x90, 0x5e, 0xe4, 0x74, 0x3c, 0x31, 0x01, 0x2e, 0x90, 0x37, 0x68, 0x6b, 0x82, 0x6b,
	0x6c, 0x5d, 0x77, 0xb9, 0xda, 0xba, 0x35, 0xf2, 0x02, 0xad, 0xc4, 0x9c, 0x29, 0x71, 0x87, 0xaa,
	0xee, 0x81, 0x01, 0xcf, 0xc1, 0x12, 0x32, 0xcb, 0x5a, 0x47, 0x6a, 0x84, 0x29, 0xcc, 0x40, 0x78,
	0x0a, 0x70, 0xac, 0x00, 0xb1, 0xbf, 0x81, 0xcb, 0xf1, 0x2c, 0x27, 0xaa, 0xc6, 0x47, 0xc0, 0x18,
	0x01, 0x9d, 0xb0, 0x77, 0x49, 0x0f, 0xea, 0xaa, 0x07, 0x59, 0xf1, 0x61, 0x0f, 0x2e, 0x26, 0x2c,
	0x7f, 0xdc, 0x64, 0x4e, 0xf6, 0xd0, 0x4e, 0x46, 0x6d, 0x12, 0x1a, 0xaa, 0xa2, 0xa1, 0x0d, 0x8a,
	0x64, 0x1f, 0xbd, 0x1f, 0xdb, 0x11, 0xbb, 0x9a, 0xc0, 0x19, 0x15, 0x34, 0x95, 0x75, 0x89, 0x94,
	0xd1, 0xde, 0xbf, 0xe0, 0xb1, 0x84, 0xe5, 0xd3, 0x9f, 0x39, 0xf4, 0xae, 0x15, 0x74, 0x2b, 0xff,
	0x73, 0x19, 0x4e, 0x37, 0x27, 0x9d, 0x85, 0xba, 0x3a, 0x2f, 0xf5, 0xdc, 0xc7, 0xfa, 0x6d, 0x3b,
	0xfc, 0x7c, 0x7f, 0x53, 0x69, 0x05, 0x5d, 0x6d, 0xc8, 0xb8, 0xdf, 0x6c, 0xab, 0xd3, 0xe4, 0xf5,
	0xfc, 0x66, 0x67, 0x3f, 0x3a, 0x44, 0x7d, 0xad, 0xdf, 0x6b, 0x69, 0xdd, 0x66, 0xdb, 0xd7, 0xa2,
	0xb7, 0x36, 0xf1, 0x86, 0xfd, 0xce, 0xe5, 0xbe, 0x4f, 0xcf, 0x50, 0x57, 0xdc, 0xe4, 0x23, 0xe4,
	0x87, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x96, 0x0d, 0x2f, 0x13, 0xf0, 0x06, 0x00, 0x00,
}