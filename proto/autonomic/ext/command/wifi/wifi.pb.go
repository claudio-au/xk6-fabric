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
// source: autonomic/ext/command/wifi/wifi.proto

package wifi

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

// Enum representing the enable / disable state of Wireless Access Point
type AccessPointState int32

const (
	AccessPointState_UNKNOWN_ACCESS_POINT_STATE AccessPointState = 0
	AccessPointState_AP_ENABLED                 AccessPointState = 1
	AccessPointState_AP_DISABLED                AccessPointState = 2
)

var AccessPointState_name = map[int32]string{
	0: "UNKNOWN_ACCESS_POINT_STATE",
	1: "AP_ENABLED",
	2: "AP_DISABLED",
}

var AccessPointState_value = map[string]int32{
	"UNKNOWN_ACCESS_POINT_STATE": 0,
	"AP_ENABLED":                 1,
	"AP_DISABLED":                2,
}

func (x AccessPointState) String() string {
	return proto.EnumName(AccessPointState_name, int32(x))
}

func (AccessPointState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c73d799b819899d7, []int{0}
}

// wifi security protocol
type AccessPointSettings_WifiSecurityProtocol int32

const (
	AccessPointSettings_UKNOWN_SECURITY_PROTOCOL AccessPointSettings_WifiSecurityProtocol = 0
	AccessPointSettings_WPA2                     AccessPointSettings_WifiSecurityProtocol = 1
	AccessPointSettings_NONE                     AccessPointSettings_WifiSecurityProtocol = 2
)

var AccessPointSettings_WifiSecurityProtocol_name = map[int32]string{
	0: "UKNOWN_SECURITY_PROTOCOL",
	1: "WPA2",
	2: "NONE",
}

var AccessPointSettings_WifiSecurityProtocol_value = map[string]int32{
	"UKNOWN_SECURITY_PROTOCOL": 0,
	"WPA2":                     1,
	"NONE":                     2,
}

func (x AccessPointSettings_WifiSecurityProtocol) String() string {
	return proto.EnumName(AccessPointSettings_WifiSecurityProtocol_name, int32(x))
}

func (AccessPointSettings_WifiSecurityProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c73d799b819899d7, []int{0, 0}
}

type AccessPointSettings struct {
	// wifi ssid
	Ssid []byte `protobuf:"bytes,1,opt,name=ssid,proto3" json:"ssid,omitempty"`
	// wifi password
	Password []byte `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Whether or not to broadcast the SSID. Some devices may not support toggling this value.
	// Those devices should:
	// 1) ignore this field when it is sent as a WifiConfigureResponse
	// 2) set this value to the non-configurable behavior of the device, when sending
	//    AccessPointSettings as part of telemetry. see: test/java/com/autonomic/ext/event/state/WifiCommandStateExamples.java
	Broadcast            bool                                     `protobuf:"varint,3,opt,name=broadcast,proto3" json:"broadcast,omitempty"`
	WifiSecurityProtocol AccessPointSettings_WifiSecurityProtocol `protobuf:"varint,4,opt,name=wifi_security_protocol,json=wifiSecurityProtocol,proto3,enum=autonomic.ext.command.wifi.AccessPointSettings_WifiSecurityProtocol" json:"wifi_security_protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *AccessPointSettings) Reset()         { *m = AccessPointSettings{} }
func (m *AccessPointSettings) String() string { return proto.CompactTextString(m) }
func (*AccessPointSettings) ProtoMessage()    {}
func (*AccessPointSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_c73d799b819899d7, []int{0}
}

func (m *AccessPointSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPointSettings.Unmarshal(m, b)
}
func (m *AccessPointSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPointSettings.Marshal(b, m, deterministic)
}
func (m *AccessPointSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPointSettings.Merge(m, src)
}
func (m *AccessPointSettings) XXX_Size() int {
	return xxx_messageInfo_AccessPointSettings.Size(m)
}
func (m *AccessPointSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPointSettings.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPointSettings proto.InternalMessageInfo

func (m *AccessPointSettings) GetSsid() []byte {
	if m != nil {
		return m.Ssid
	}
	return nil
}

func (m *AccessPointSettings) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *AccessPointSettings) GetBroadcast() bool {
	if m != nil {
		return m.Broadcast
	}
	return false
}

func (m *AccessPointSettings) GetWifiSecurityProtocol() AccessPointSettings_WifiSecurityProtocol {
	if m != nil {
		return m.WifiSecurityProtocol
	}
	return AccessPointSettings_UKNOWN_SECURITY_PROTOCOL
}

// A request to update wifi AP settings, which may also change the AP state (enable disable)
type WifiConfigRequest struct {
	ApSettings *AccessPointSettings `protobuf:"bytes,1,opt,name=ap_settings,json=apSettings,proto3" json:"ap_settings,omitempty"`
	// Configure the AP state (enable / disable)
	ApState              AccessPointState `protobuf:"varint,2,opt,name=ap_state,json=apState,proto3,enum=autonomic.ext.command.wifi.AccessPointState" json:"ap_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *WifiConfigRequest) Reset()         { *m = WifiConfigRequest{} }
func (m *WifiConfigRequest) String() string { return proto.CompactTextString(m) }
func (*WifiConfigRequest) ProtoMessage()    {}
func (*WifiConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c73d799b819899d7, []int{1}
}

func (m *WifiConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiConfigRequest.Unmarshal(m, b)
}
func (m *WifiConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiConfigRequest.Marshal(b, m, deterministic)
}
func (m *WifiConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiConfigRequest.Merge(m, src)
}
func (m *WifiConfigRequest) XXX_Size() int {
	return xxx_messageInfo_WifiConfigRequest.Size(m)
}
func (m *WifiConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WifiConfigRequest proto.InternalMessageInfo

func (m *WifiConfigRequest) GetApSettings() *AccessPointSettings {
	if m != nil {
		return m.ApSettings
	}
	return nil
}

func (m *WifiConfigRequest) GetApState() AccessPointState {
	if m != nil {
		return m.ApState
	}
	return AccessPointState_UNKNOWN_ACCESS_POINT_STATE
}

// Request to receive the various wifi access point settings. AcccessPointSettings should be sent
// as part of normal telemetry. see: test/java/com/autonomic/ext/event/state/WifiCommandStateExamples.java
type WifiStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WifiStatusRequest) Reset()         { *m = WifiStatusRequest{} }
func (m *WifiStatusRequest) String() string { return proto.CompactTextString(m) }
func (*WifiStatusRequest) ProtoMessage()    {}
func (*WifiStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c73d799b819899d7, []int{2}
}

func (m *WifiStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiStatusRequest.Unmarshal(m, b)
}
func (m *WifiStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiStatusRequest.Marshal(b, m, deterministic)
}
func (m *WifiStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiStatusRequest.Merge(m, src)
}
func (m *WifiStatusRequest) XXX_Size() int {
	return xxx_messageInfo_WifiStatusRequest.Size(m)
}
func (m *WifiStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WifiStatusRequest proto.InternalMessageInfo

// The response of WifiStatus command, needs to be populated by device at StateTransitionData when transitioned to success
type WifiStatusSettings struct {
	ApSettings           *AccessPointSettings `protobuf:"bytes,1,opt,name=ap_settings,json=apSettings,proto3" json:"ap_settings,omitempty"`
	ApState              AccessPointState     `protobuf:"varint,2,opt,name=ap_state,json=apState,proto3,enum=autonomic.ext.command.wifi.AccessPointState" json:"ap_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WifiStatusSettings) Reset()         { *m = WifiStatusSettings{} }
func (m *WifiStatusSettings) String() string { return proto.CompactTextString(m) }
func (*WifiStatusSettings) ProtoMessage()    {}
func (*WifiStatusSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_c73d799b819899d7, []int{3}
}

func (m *WifiStatusSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiStatusSettings.Unmarshal(m, b)
}
func (m *WifiStatusSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiStatusSettings.Marshal(b, m, deterministic)
}
func (m *WifiStatusSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiStatusSettings.Merge(m, src)
}
func (m *WifiStatusSettings) XXX_Size() int {
	return xxx_messageInfo_WifiStatusSettings.Size(m)
}
func (m *WifiStatusSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiStatusSettings.DiscardUnknown(m)
}

var xxx_messageInfo_WifiStatusSettings proto.InternalMessageInfo

func (m *WifiStatusSettings) GetApSettings() *AccessPointSettings {
	if m != nil {
		return m.ApSettings
	}
	return nil
}

func (m *WifiStatusSettings) GetApState() AccessPointState {
	if m != nil {
		return m.ApState
	}
	return AccessPointState_UNKNOWN_ACCESS_POINT_STATE
}

func init() {
	proto.RegisterEnum("autonomic.ext.command.wifi.AccessPointState", AccessPointState_name, AccessPointState_value)
	proto.RegisterEnum("autonomic.ext.command.wifi.AccessPointSettings_WifiSecurityProtocol", AccessPointSettings_WifiSecurityProtocol_name, AccessPointSettings_WifiSecurityProtocol_value)
	proto.RegisterType((*AccessPointSettings)(nil), "autonomic.ext.command.wifi.AccessPointSettings")
	proto.RegisterType((*WifiConfigRequest)(nil), "autonomic.ext.command.wifi.WifiConfigRequest")
	proto.RegisterType((*WifiStatusRequest)(nil), "autonomic.ext.command.wifi.WifiStatusRequest")
	proto.RegisterType((*WifiStatusSettings)(nil), "autonomic.ext.command.wifi.WifiStatusSettings")
}

func init() {
	proto.RegisterFile("autonomic/ext/command/wifi/wifi.proto", fileDescriptor_c73d799b819899d7)
}

var fileDescriptor_c73d799b819899d7 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0xdd, 0xc9, 0x2e, 0xda, 0xbd, 0x95, 0x1a, 0x67, 0x17, 0x09, 0x65, 0x59, 0x4a, 0x40, 0x28,
	0xe2, 0x26, 0x50, 0xbf, 0x20, 0x4d, 0x83, 0x16, 0x97, 0x24, 0x26, 0x29, 0x45, 0x5f, 0xc2, 0x34,
	0x4d, 0xeb, 0xc0, 0x26, 0x13, 0x33, 0x13, 0xba, 0xfa, 0x07, 0xfe, 0x86, 0x6f, 0xbe, 0xe8, 0xef,
	0xf9, 0x28, 0x33, 0x0d, 0x59, 0x90, 0xae, 0xe8, 0x9b, 0x2f, 0xe1, 0xde, 0x93, 0x33, 0x67, 0xce,
	0x9c, 0x99, 0x0b, 0xcf, 0x48, 0x23, 0x58, 0xc9, 0x0a, 0x9a, 0xd9, 0xf9, 0xad, 0xb0, 0x33, 0x56,
	0x14, 0xa4, 0x5c, 0xdb, 0x3b, 0xba, 0xa1, 0xea, 0x63, 0x55, 0x35, 0x13, 0x0c, 0x0f, 0x3b, 0x9a,
	0x95, 0xdf, 0x0a, 0xab, 0xa5, 0x59, 0x92, 0x61, 0x7e, 0xd3, 0xe0, 0xcc, 0xc9, 0xb2, 0x9c, 0xf3,
	0x90, 0xd1, 0x52, 0xc4, 0xb9, 0x10, 0xb4, 0xdc, 0x72, 0x8c, 0xe1, 0x84, 0x73, 0xba, 0x36, 0xd0,
	0x08, 0x8d, 0x1f, 0x45, 0xaa, 0xc6, 0x43, 0xe8, 0x55, 0x84, 0xf3, 0x1d, 0xab, 0xd7, 0x86, 0xa6,
	0xf0, 0xae, 0xc7, 0x17, 0x70, 0xba, 0xaa, 0x19, 0x59, 0x67, 0x84, 0x0b, 0xe3, 0x78, 0x84, 0xc6,
	0xbd, 0xe8, 0x0e, 0xc0, 0x9f, 0xe1, 0xa9, 0xdc, 0x2d, 0xe5, 0x79, 0xd6, 0xd4, 0x54, 0x7c, 0x4a,
	0x95, 0xb1, 0x8c, 0xdd, 0x18, 0x27, 0x23, 0x34, 0x1e, 0x4c, 0x66, 0xd6, 0xfd, 0x16, 0xad, 0x03,
	0xf6, 0xac, 0x25, 0xdd, 0xd0, 0xb8, 0x15, 0x0b, 0x5b, 0xad, 0xe8, 0x7c, 0x77, 0x00, 0x35, 0x5f,
	0xc3, 0xf9, 0x21, 0x36, 0xbe, 0x00, 0x63, 0xf1, 0xc6, 0x0f, 0x96, 0x7e, 0x1a, 0x7b, 0xee, 0x22,
	0x9a, 0x27, 0xef, 0xd2, 0x30, 0x0a, 0x92, 0xc0, 0x0d, 0xae, 0xf5, 0x23, 0xdc, 0x83, 0x93, 0x65,
	0xe8, 0x4c, 0x74, 0x24, 0x2b, 0x3f, 0xf0, 0x3d, 0x5d, 0x33, 0xbf, 0x23, 0x78, 0x22, 0xa5, 0x5c,
	0x56, 0x6e, 0xe8, 0x36, 0xca, 0x3f, 0x36, 0x39, 0x17, 0x38, 0x84, 0x3e, 0xa9, 0x52, 0xde, 0x3a,
	0x53, 0x81, 0xf5, 0x27, 0xf6, 0x3f, 0x1e, 0x28, 0x02, 0x52, 0x75, 0xd9, 0xbf, 0x82, 0x9e, 0x54,
	0x14, 0x44, 0xe4, 0x2a, 0xe7, 0xc1, 0xe4, 0xc5, 0xdf, 0xca, 0xc9, 0x35, 0xd1, 0x43, 0x52, 0xa9,
	0xc2, 0x3c, 0xdb, 0xfb, 0x95, 0x4d, 0xc3, 0x5b, 0xbf, 0xe6, 0x0f, 0x04, 0xf8, 0x0e, 0xed, 0x36,
	0xfd, 0x7f, 0x8f, 0xf1, 0x3c, 0x06, 0xfd, 0xf7, 0x9f, 0xf8, 0x12, 0x86, 0x0b, 0x7f, 0x7f, 0x7d,
	0x8e, 0xeb, 0x7a, 0x71, 0x9c, 0x86, 0xc1, 0xdc, 0x4f, 0xd2, 0x38, 0x71, 0x12, 0x4f, 0x3f, 0xc2,
	0x03, 0x00, 0x27, 0x4c, 0x3d, 0xdf, 0x99, 0x5e, 0x7b, 0x33, 0x1d, 0xe1, 0xc7, 0xd0, 0x77, 0xc2,
	0x74, 0x36, 0x8f, 0xf7, 0x80, 0x36, 0xfd, 0x82, 0xe0, 0x32, 0x63, 0xc5, 0x1f, 0x1c, 0x4d, 0x4f,
	0x65, 0x4c, 0xea, 0xbd, 0x84, 0xe8, 0xfd, 0xdb, 0x2d, 0x15, 0x1f, 0x9a, 0x95, 0x64, 0xd8, 0xdd,
	0x9a, 0x2b, 0x42, 0xe5, 0xe4, 0xe5, 0x75, 0x49, 0x6e, 0xae, 0xd4, 0x9b, 0xe6, 0x36, 0xaf, 0x33,
	0xbb, 0x20, 0xb4, 0xb4, 0x55, 0x6f, 0xdf, 0x3f, 0xa2, 0x3f, 0x11, 0xfa, 0xaa, 0x1d, 0x3b, 0x8b,
	0x64, 0xf5, 0x40, 0x51, 0x5f, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xec, 0x06, 0x46, 0xd0,
	0x03, 0x00, 0x00,
}