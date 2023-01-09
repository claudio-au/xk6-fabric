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
// source: autonomic/ext/telemetry/report.proto

package telemetry

import (
	fmt "fmt"
	modules "xk6-fabric/proto/autonomic/ext/telemetry/modules"
	//modules "xk6-fabric/proto/autonomic/ext/telemetry/modules"
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

type Report struct {
	// Au service generating the information - aui:mqs, aui:gfs, aui:feed,
	// aui:deploy, aui:bytestream, aui:edge:vehicle:<vin>, etc
	Source    string               `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Metrics   []*Metric            `protobuf:"bytes,3,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// discovered modules on the vehicle, including the device reporting this data
	Modules []*modules.Module `protobuf:"bytes,4,rep,name=modules,proto3" json:"modules,omitempty"`
	// telemetry report metric intent
	Tags []*ReportTag `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// Unique identifier set by the customer (opaque to Au)
	OemCorrelationId     string   `protobuf:"bytes,6,opt,name=oem_correlation_id,json=oemCorrelationId,proto3" json:"oem_correlation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_c541f640f5708aef, []int{0}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Report) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Report) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Report) GetModules() []*modules.Module {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *Report) GetTags() []*ReportTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Report) GetOemCorrelationId() string {
	if m != nil {
		return m.OemCorrelationId
	}
	return ""
}

func init() {
	proto.RegisterType((*Report)(nil), "autonomic.ext.telemetry.Report")
}

func init() {
	proto.RegisterFile("autonomic/ext/telemetry/report.proto", fileDescriptor_c541f640f5708aef)
}

var fileDescriptor_c541f640f5708aef = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0xb6, 0x04, 0xd5, 0x5d, 0x50, 0x06, 0x88, 0xca, 0xd0, 0xaa, 0x62, 0xc8, 0xd0,
	0xda, 0x52, 0x91, 0x10, 0x8c, 0x94, 0x89, 0x01, 0x54, 0x45, 0x65, 0x61, 0xa9, 0x5c, 0xf7, 0x08,
	0x96, 0xe2, 0x5c, 0x65, 0x3b, 0x52, 0x79, 0x00, 0x5e, 0x84, 0x87, 0xe0, 0xd9, 0x18, 0x51, 0xed,
	0x24, 0x9d, 0x02, 0x4b, 0xa2, 0x53, 0xbe, 0xef, 0x57, 0xee, 0x3f, 0x72, 0xc5, 0x4b, 0x8b, 0x05,
	0x2a, 0x29, 0x18, 0xec, 0x2d, 0xb3, 0x90, 0x83, 0x02, 0xab, 0x3f, 0x98, 0x86, 0x1d, 0x6a, 0x4b,
	0x77, 0x1a, 0x2d, 0x46, 0x17, 0x0d, 0x45, 0x61, 0x6f, 0x69, 0x43, 0x0d, 0x5b, 0xf5, 0xc3, 0x53,
	0x0a, 0xaf, 0x0f, 0xa7, 0xad, 0x14, 0x6e, 0xcb, 0x1c, 0x4c, 0xf5, 0xae, 0xe8, 0xe4, 0xef, 0x5f,
	0x5a, 0x5b, 0x9e, 0x55, 0xe4, 0x28, 0x43, 0xcc, 0x72, 0x60, 0x6e, 0xda, 0x94, 0x6f, 0xcc, 0x4a,
	0x05, 0xc6, 0x72, 0xb5, 0xf3, 0xc0, 0xe4, 0xbb, 0x43, 0xc2, 0xd4, 0x59, 0xd1, 0x39, 0x09, 0x0d,
	0x96, 0x5a, 0x40, 0x1c, 0x8c, 0x83, 0xa4, 0x9f, 0x56, 0x53, 0x74, 0x4b, 0xfa, 0x8d, 0x15, 0x77,
	0xc6, 0x41, 0x32, 0x98, 0x0f, 0xa9, 0xcf, 0xa5, 0x75, 0x2e, 0x5d, 0xd5, 0x44, 0x7a, 0x84, 0xa3,
	0x3b, 0x72, 0xea, 0xb7, 0x34, 0x71, 0x77, 0xdc, 0x4d, 0x06, 0xf3, 0x11, 0x6d, 0xa9, 0x89, 0x3e,
	0x39, 0x2e, 0xad, 0x79, 0xa7, 0xfa, 0xd5, 0xe3, 0xde, 0x7f, 0xaa, 0xe3, 0xd2, 0x9a, 0x8f, 0x6e,
	0x48, 0xcf, 0xf2, 0xcc, 0xc4, 0x27, 0xce, 0x9b, 0xb4, 0x7a, 0x7e, 0xed, 0x15, 0xcf, 0x52, 0xc7,
	0x47, 0x53, 0x12, 0x21, 0xa8, 0xb5, 0x40, 0xad, 0x21, 0xe7, 0x56, 0x62, 0xb1, 0x96, 0xdb, 0x38,
	0x74, 0x5d, 0x9c, 0x21, 0xa8, 0x87, 0xe3, 0x87, 0xc7, 0xed, 0xe2, 0x33, 0x20, 0x97, 0x02, 0x55,
	0x5b, 0xfa, 0x62, 0xe0, 0xe3, 0x97, 0x87, 0x82, 0x96, 0xc1, 0xeb, 0x73, 0x26, 0xed, 0x7b, 0xb9,
	0xa1, 0x02, 0x15, 0x6b, 0x94, 0x19, 0x97, 0x87, 0x03, 0x82, 0x2e, 0x78, 0x3e, 0x73, 0x55, 0x1a,
	0x66, 0xb4, 0x60, 0x8a, 0xcb, 0xc2, 0x9f, 0x8c, 0xb5, 0x5c, 0xfa, 0x27, 0x08, 0xbe, 0x3a, 0xdd,
	0xfb, 0x97, 0xd5, 0x26, 0x74, 0xdc, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xed, 0x16,
	0x6a, 0xa7, 0x02, 0x00, 0x00,
}