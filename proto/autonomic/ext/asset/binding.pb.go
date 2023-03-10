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
// source: autonomic/ext/asset/binding.proto

package asset

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type BindingChangeEvent_EventType int32

const (
	// Implies that a new binding has been created between a Vehicle and a Device.
	// An event of this type is NOT guaranteed to be produced for all newly created
	// bindings. Sometimes, an event with `type` set to UPDATED and `end_time` set to
	// `infinite` will also imply that a new binding has been created. So, anyone
	// interested in acting on newly created bindings should check this condition:
	// "if (type == CREATED or end_time == infinite)"
	BindingChangeEvent_CREATED BindingChangeEvent_EventType = 0
	// Implies that the `end_time` has been updated. There are scenarios (usually involving
	// pipeline monitoring systems) where an UPDATE event will be published but no attribute
	// of the binding has changed. The timestamp of the envelope which contains this message
	// will be updated in all cases.
	// Remember that an `unbind` operation is implemented by setting the `end_time` to the
	// current timestamp. So, whenever a Device is unbound from a Vehicle, an event with
	// type set to UPDATED will be generated.
	BindingChangeEvent_UPDATED BindingChangeEvent_EventType = 1
	// This is only for internal operations and should be ignored.
	BindingChangeEvent_DELETED BindingChangeEvent_EventType = 2
)

var BindingChangeEvent_EventType_name = map[int32]string{
	0: "CREATED",
	1: "UPDATED",
	2: "DELETED",
}

var BindingChangeEvent_EventType_value = map[string]int32{
	"CREATED": 0,
	"UPDATED": 1,
	"DELETED": 2,
}

func (x BindingChangeEvent_EventType) String() string {
	return proto.EnumName(BindingChangeEvent_EventType_name, int32(x))
}

func (BindingChangeEvent_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c606a41702e52224, []int{0, 0}
}

// protos for binding event
type BindingChangeEvent struct {
	Type BindingChangeEvent_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=autonomic.ext.asset.BindingChangeEvent_EventType" json:"type,omitempty"`
	// The time when the binding becomes active. This is always before the `end_time`.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time which the binding becomes inactive.
	EndTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// The Au assigned id of the subject in the relationship.
	// This is usually the id of the device.
	SubjectId string `protobuf:"bytes,4,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	// The Au assigned id of the object in the relationship.
	// This is usually the id of the vehicle.
	ObjectId string `protobuf:"bytes,5,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	// Types that are valid to be assigned to ObjectDetails:
	//	*BindingChangeEvent_VehicleDetails
	ObjectDetails isBindingChangeEvent_ObjectDetails `protobuf_oneof:"object_details"`
	// Types that are valid to be assigned to SubjectDetails:
	//	*BindingChangeEvent_DeviceDetails
	SubjectDetails       isBindingChangeEvent_SubjectDetails `protobuf_oneof:"subject_details"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *BindingChangeEvent) Reset()         { *m = BindingChangeEvent{} }
func (m *BindingChangeEvent) String() string { return proto.CompactTextString(m) }
func (*BindingChangeEvent) ProtoMessage()    {}
func (*BindingChangeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_c606a41702e52224, []int{0}
}

func (m *BindingChangeEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindingChangeEvent.Unmarshal(m, b)
}
func (m *BindingChangeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindingChangeEvent.Marshal(b, m, deterministic)
}
func (m *BindingChangeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingChangeEvent.Merge(m, src)
}
func (m *BindingChangeEvent) XXX_Size() int {
	return xxx_messageInfo_BindingChangeEvent.Size(m)
}
func (m *BindingChangeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingChangeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BindingChangeEvent proto.InternalMessageInfo

func (m *BindingChangeEvent) GetType() BindingChangeEvent_EventType {
	if m != nil {
		return m.Type
	}
	return BindingChangeEvent_CREATED
}

func (m *BindingChangeEvent) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *BindingChangeEvent) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *BindingChangeEvent) GetSubjectId() string {
	if m != nil {
		return m.SubjectId
	}
	return ""
}

func (m *BindingChangeEvent) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

type isBindingChangeEvent_ObjectDetails interface {
	isBindingChangeEvent_ObjectDetails()
}

type BindingChangeEvent_VehicleDetails struct {
	VehicleDetails *VehicleDetails `protobuf:"bytes,6,opt,name=vehicle_details,json=vehicleDetails,proto3,oneof"`
}

func (*BindingChangeEvent_VehicleDetails) isBindingChangeEvent_ObjectDetails() {}

func (m *BindingChangeEvent) GetObjectDetails() isBindingChangeEvent_ObjectDetails {
	if m != nil {
		return m.ObjectDetails
	}
	return nil
}

func (m *BindingChangeEvent) GetVehicleDetails() *VehicleDetails {
	if x, ok := m.GetObjectDetails().(*BindingChangeEvent_VehicleDetails); ok {
		return x.VehicleDetails
	}
	return nil
}

type isBindingChangeEvent_SubjectDetails interface {
	isBindingChangeEvent_SubjectDetails()
}

type BindingChangeEvent_DeviceDetails struct {
	DeviceDetails *DeviceDetails `protobuf:"bytes,7,opt,name=device_details,json=deviceDetails,proto3,oneof"`
}

func (*BindingChangeEvent_DeviceDetails) isBindingChangeEvent_SubjectDetails() {}

func (m *BindingChangeEvent) GetSubjectDetails() isBindingChangeEvent_SubjectDetails {
	if m != nil {
		return m.SubjectDetails
	}
	return nil
}

func (m *BindingChangeEvent) GetDeviceDetails() *DeviceDetails {
	if x, ok := m.GetSubjectDetails().(*BindingChangeEvent_DeviceDetails); ok {
		return x.DeviceDetails
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BindingChangeEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BindingChangeEvent_VehicleDetails)(nil),
		(*BindingChangeEvent_DeviceDetails)(nil),
	}
}

type VehicleDetails struct {
	// The VIN
	Vin                  string   `protobuf:"bytes,1,opt,name=vin,proto3" json:"vin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VehicleDetails) Reset()         { *m = VehicleDetails{} }
func (m *VehicleDetails) String() string { return proto.CompactTextString(m) }
func (*VehicleDetails) ProtoMessage()    {}
func (*VehicleDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_c606a41702e52224, []int{1}
}

func (m *VehicleDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehicleDetails.Unmarshal(m, b)
}
func (m *VehicleDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehicleDetails.Marshal(b, m, deterministic)
}
func (m *VehicleDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehicleDetails.Merge(m, src)
}
func (m *VehicleDetails) XXX_Size() int {
	return xxx_messageInfo_VehicleDetails.Size(m)
}
func (m *VehicleDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_VehicleDetails.DiscardUnknown(m)
}

var xxx_messageInfo_VehicleDetails proto.InternalMessageInfo

func (m *VehicleDetails) GetVin() string {
	if m != nil {
		return m.Vin
	}
	return ""
}

type DeviceDetails struct {
	// One of
	// - epid2
	// - tcu
	// - ecg
	// - sync4
	// More device types may be added without prior warning.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The serial number of the device. Guaranteed to be unique per the type of device.
	Esn string `protobuf:"bytes,2,opt,name=esn,proto3" json:"esn,omitempty"`
	// Optional. IMEI field of the device
	Imei string `protobuf:"bytes,4,opt,name=imei,proto3" json:"imei,omitempty"`
	// Optional. ICCID field of the device
	Iccid string `protobuf:"bytes,5,opt,name=iccid,proto3" json:"iccid,omitempty"`
	// Optional. IMSI field of the device
	Imsi string `protobuf:"bytes,6,opt,name=imsi,proto3" json:"imsi,omitempty"`
	// Optional. HardwarePartNumber of the device
	HardwarePartNumber string `protobuf:"bytes,7,opt,name=hardwarePartNumber,proto3" json:"hardwarePartNumber,omitempty"`
	// Other surfaced properties that may be exposed later
	Properties           map[string]string `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeviceDetails) Reset()         { *m = DeviceDetails{} }
func (m *DeviceDetails) String() string { return proto.CompactTextString(m) }
func (*DeviceDetails) ProtoMessage()    {}
func (*DeviceDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_c606a41702e52224, []int{2}
}

func (m *DeviceDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceDetails.Unmarshal(m, b)
}
func (m *DeviceDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceDetails.Marshal(b, m, deterministic)
}
func (m *DeviceDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceDetails.Merge(m, src)
}
func (m *DeviceDetails) XXX_Size() int {
	return xxx_messageInfo_DeviceDetails.Size(m)
}
func (m *DeviceDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceDetails.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceDetails proto.InternalMessageInfo

func (m *DeviceDetails) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DeviceDetails) GetEsn() string {
	if m != nil {
		return m.Esn
	}
	return ""
}

func (m *DeviceDetails) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *DeviceDetails) GetIccid() string {
	if m != nil {
		return m.Iccid
	}
	return ""
}

func (m *DeviceDetails) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *DeviceDetails) GetHardwarePartNumber() string {
	if m != nil {
		return m.HardwarePartNumber
	}
	return ""
}

func (m *DeviceDetails) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterEnum("autonomic.ext.asset.BindingChangeEvent_EventType", BindingChangeEvent_EventType_name, BindingChangeEvent_EventType_value)
	proto.RegisterType((*BindingChangeEvent)(nil), "autonomic.ext.asset.BindingChangeEvent")
	proto.RegisterType((*VehicleDetails)(nil), "autonomic.ext.asset.VehicleDetails")
	proto.RegisterType((*DeviceDetails)(nil), "autonomic.ext.asset.DeviceDetails")
	proto.RegisterMapType((map[string]string)(nil), "autonomic.ext.asset.DeviceDetails.PropertiesEntry")
}

func init() { proto.RegisterFile("autonomic/ext/asset/binding.proto", fileDescriptor_c606a41702e52224) }

var fileDescriptor_c606a41702e52224 = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xae, 0x93, 0xb4, 0x8d, 0xa7, 0x6f, 0xd3, 0xbc, 0x0b, 0x12, 0x51, 0x11, 0x22, 0x84, 0x4b,
	0x2e, 0x5d, 0x8b, 0x20, 0x24, 0x40, 0xe2, 0xd0, 0x34, 0x96, 0xa0, 0xa0, 0x2a, 0xb2, 0x52, 0x0e,
	0x5c, 0xaa, 0xb5, 0x3d, 0x24, 0x4b, 0xe3, 0xb5, 0xe5, 0x5d, 0x9b, 0xe6, 0xc2, 0x8d, 0x3f, 0xc2,
	0x8f, 0xe0, 0xb7, 0x71, 0x44, 0xbb, 0xfe, 0xa0, 0x01, 0x4b, 0xbd, 0x58, 0x3b, 0x33, 0xcf, 0xf3,
	0xec, 0x7c, 0xad, 0xe1, 0x09, 0xcb, 0x54, 0x2c, 0xe2, 0x88, 0x07, 0x0e, 0xde, 0x28, 0x87, 0x49,
	0x89, 0xca, 0xf1, 0xb9, 0x08, 0xb9, 0x58, 0xd2, 0x24, 0x8d, 0x55, 0x4c, 0xee, 0xd5, 0x10, 0x8a,
	0x37, 0x8a, 0x1a, 0xc8, 0xf1, 0xe3, 0x65, 0x1c, 0x2f, 0xd7, 0xe8, 0x18, 0x88, 0x9f, 0x7d, 0x76,
	0x14, 0x8f, 0x50, 0x2a, 0x16, 0x25, 0x05, 0x6b, 0xf4, 0xbd, 0x03, 0x64, 0x5a, 0xe8, 0x9c, 0xad,
	0x98, 0x58, 0xa2, 0x9b, 0xa3, 0x50, 0xc4, 0x85, 0x8e, 0xda, 0x24, 0x38, 0xb0, 0x86, 0xd6, 0xb8,
	0x37, 0x79, 0x46, 0x1b, 0xb4, 0xe9, 0xbf, 0x34, 0x6a, 0xbe, 0x8b, 0x4d, 0x82, 0x9e, 0xa1, 0x93,
	0x57, 0x00, 0x52, 0xb1, 0x54, 0x5d, 0xe9, 0x6b, 0x07, 0xad, 0xa1, 0x35, 0x3e, 0x98, 0x1c, 0xd3,
	0x22, 0x27, 0x5a, 0xe5, 0x44, 0x17, 0x55, 0x4e, 0x9e, 0x6d, 0xd0, 0xda, 0x26, 0x2f, 0xa0, 0x8b,
	0x22, 0x2c, 0x88, 0xed, 0x3b, 0x89, 0xfb, 0x28, 0x42, 0x43, 0x7b, 0x04, 0x20, 0x33, 0xff, 0x0b,
	0x06, 0xea, 0x8a, 0x87, 0x83, 0xce, 0xd0, 0x1a, 0xdb, 0x9e, 0x5d, 0x7a, 0xde, 0x85, 0xe4, 0x21,
	0xd8, 0x71, 0x1d, 0xdd, 0x35, 0xd1, 0x6e, 0x5c, 0x05, 0x2f, 0xe0, 0x28, 0xc7, 0x15, 0x0f, 0xd6,
	0x78, 0x15, 0xa2, 0x62, 0x7c, 0x2d, 0x07, 0x7b, 0xe6, 0xe6, 0xa7, 0x8d, 0xf5, 0x7f, 0x2c, 0xb0,
	0xb3, 0x02, 0xfa, 0x76, 0xc7, 0xeb, 0xe5, 0x5b, 0x1e, 0xf2, 0x1e, 0x7a, 0x21, 0xe6, 0x3c, 0xf8,
	0x23, 0xb7, 0x6f, 0xe4, 0x46, 0x8d, 0x72, 0x33, 0x03, 0xad, 0xd4, 0x2c, 0xef, 0x30, 0xbc, 0xed,
	0x18, 0x4d, 0xc0, 0xae, 0xbb, 0x4b, 0x0e, 0x60, 0xff, 0xcc, 0x73, 0x4f, 0x17, 0xee, 0xac, 0xbf,
	0xa3, 0x8d, 0xcb, 0xf9, 0xcc, 0x18, 0x96, 0x36, 0x66, 0xee, 0x07, 0x57, 0x1b, 0xad, 0x69, 0x1f,
	0x7a, 0x65, 0xb5, 0x65, 0x02, 0xd3, 0xff, 0xe1, 0xa8, 0x6a, 0x4f, 0xe9, 0x3a, 0xef, 0x74, 0xbb,
	0x7d, 0x7b, 0x34, 0x82, 0xde, 0x76, 0x3d, 0xa4, 0x0f, 0xed, 0x9c, 0x0b, 0xb3, 0x01, 0xb6, 0xa7,
	0x8f, 0xa3, 0x9f, 0x2d, 0x38, 0xdc, 0xca, 0x92, 0x90, 0x5b, 0x6b, 0x62, 0x97, 0x33, 0xef, 0x43,
	0x1b, 0xa5, 0x30, 0xc3, 0xb6, 0x3d, 0x7d, 0xd4, 0x28, 0x1e, 0x21, 0x2f, 0xa7, 0x61, 0xce, 0xe4,
	0x3e, 0xec, 0xf2, 0x20, 0xa8, 0x87, 0x50, 0x18, 0x05, 0x52, 0x72, 0xd3, 0x76, 0x83, 0x94, 0x9c,
	0x50, 0x20, 0x2b, 0x96, 0x86, 0x5f, 0x59, 0x8a, 0x73, 0x96, 0xaa, 0x8b, 0x2c, 0xf2, 0x31, 0x35,
	0x9d, 0xb4, 0xbd, 0x86, 0x08, 0xf1, 0x00, 0x92, 0x34, 0x4e, 0x30, 0x55, 0x1c, 0xe5, 0xa0, 0x3d,
	0x6c, 0x8f, 0x0f, 0x26, 0x93, 0xbb, 0x3b, 0x4e, 0xe7, 0x35, 0xc9, 0x15, 0x2a, 0xdd, 0x78, 0xb7,
	0x54, 0x8e, 0xdf, 0xc0, 0xd1, 0x5f, 0x61, 0x5d, 0xe6, 0x35, 0x6e, 0xaa, 0xf6, 0x5c, 0xe3, 0x46,
	0x97, 0x94, 0xb3, 0x75, 0x86, 0x65, 0xe9, 0x85, 0xf1, 0xba, 0xf5, 0xd2, 0x9a, 0x7e, 0x83, 0x07,
	0x41, 0x1c, 0x35, 0xe5, 0x30, 0xfd, 0xaf, 0x7c, 0x45, 0x73, 0xbd, 0xd3, 0x73, 0xeb, 0xd3, 0xf9,
	0x92, 0xab, 0x55, 0xe6, 0xd3, 0x20, 0x8e, 0x9c, 0x1a, 0x7f, 0xc2, 0xb8, 0x7e, 0xf6, 0x98, 0x0a,
	0xb6, 0x3e, 0x31, 0xdb, 0x2f, 0x1d, 0x99, 0x06, 0x4e, 0xc4, 0xb8, 0x28, 0x9e, 0xb6, 0xd3, 0xf0,
	0x7f, 0xf8, 0x65, 0x59, 0x3f, 0x5a, 0xed, 0xd3, 0xcb, 0x85, 0xbf, 0x67, 0x30, 0xcf, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0xd2, 0xbf, 0x3a, 0x46, 0x04, 0x00, 0x00,
}
