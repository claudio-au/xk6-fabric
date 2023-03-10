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
// source: autonomic/ext/common/schedule.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	date "google.golang.org/genproto/googleapis/type/date"
	dayofweek "google.golang.org/genproto/googleapis/type/dayofweek"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
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

type TimeZone int32

const (
	TimeZone_UNKNOWN    TimeZone = 0
	TimeZone_UTC        TimeZone = 1
	TimeZone_LOCAL_TIME TimeZone = 2
)

var TimeZone_name = map[int32]string{
	0: "UNKNOWN",
	1: "UTC",
	2: "LOCAL_TIME",
}

var TimeZone_value = map[string]int32{
	"UNKNOWN":    0,
	"UTC":        1,
	"LOCAL_TIME": 2,
}

func (x TimeZone) String() string {
	return proto.EnumName(TimeZone_name, int32(x))
}

func (TimeZone) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c54d7c732fc3473e, []int{0}
}

type Schedule_ScheduleType int32

const (
	Schedule_UNKNOWN   Schedule_ScheduleType = 0
	Schedule_NONE      Schedule_ScheduleType = 1
	Schedule_ONE_TIME  Schedule_ScheduleType = 2
	Schedule_IMMEDIATE Schedule_ScheduleType = 3
	Schedule_RECURRING Schedule_ScheduleType = 4
)

var Schedule_ScheduleType_name = map[int32]string{
	0: "UNKNOWN",
	1: "NONE",
	2: "ONE_TIME",
	3: "IMMEDIATE",
	4: "RECURRING",
}

var Schedule_ScheduleType_value = map[string]int32{
	"UNKNOWN":   0,
	"NONE":      1,
	"ONE_TIME":  2,
	"IMMEDIATE": 3,
	"RECURRING": 4,
}

func (x Schedule_ScheduleType) String() string {
	return proto.EnumName(Schedule_ScheduleType_name, int32(x))
}

func (Schedule_ScheduleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c54d7c732fc3473e, []int{0, 0}
}

type Schedule_ScheduleExecutor int32

const (
	Schedule_UNKNOWN_EXECUTOR Schedule_ScheduleExecutor = 0
	Schedule_DEVICE           Schedule_ScheduleExecutor = 1
	Schedule_CLOUD            Schedule_ScheduleExecutor = 2
)

var Schedule_ScheduleExecutor_name = map[int32]string{
	0: "UNKNOWN_EXECUTOR",
	1: "DEVICE",
	2: "CLOUD",
}

var Schedule_ScheduleExecutor_value = map[string]int32{
	"UNKNOWN_EXECUTOR": 0,
	"DEVICE":           1,
	"CLOUD":            2,
}

func (x Schedule_ScheduleExecutor) String() string {
	return proto.EnumName(Schedule_ScheduleExecutor_name, int32(x))
}

func (Schedule_ScheduleExecutor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c54d7c732fc3473e, []int{0, 1}
}

type Schedule struct {
	ScheduleType Schedule_ScheduleType `protobuf:"varint,1,opt,name=schedule_type,json=scheduleType,proto3,enum=autonomic.ext.common.schedule.Schedule_ScheduleType" json:"schedule_type,omitempty"`
	// Types that are valid to be assigned to ScheduledTime:
	//	*Schedule_OneTimeSchedule
	//	*Schedule_DailySchedule
	//	*Schedule_WeeklySchedule
	ScheduledTime        isSchedule_ScheduledTime  `protobuf_oneof:"scheduled_time"`
	TimeZone             TimeZone                  `protobuf:"varint,5,opt,name=time_zone,json=timeZone,proto3,enum=autonomic.ext.common.schedule.TimeZone" json:"time_zone,omitempty"`
	ScheduleExecutor     Schedule_ScheduleExecutor `protobuf:"varint,6,opt,name=schedule_executor,json=scheduleExecutor,proto3,enum=autonomic.ext.common.schedule.Schedule_ScheduleExecutor" json:"schedule_executor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c54d7c732fc3473e, []int{0}
}

func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schedule.Unmarshal(m, b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
}
func (m *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(m, src)
}
func (m *Schedule) XXX_Size() int {
	return xxx_messageInfo_Schedule.Size(m)
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

func (m *Schedule) GetScheduleType() Schedule_ScheduleType {
	if m != nil {
		return m.ScheduleType
	}
	return Schedule_UNKNOWN
}

type isSchedule_ScheduledTime interface {
	isSchedule_ScheduledTime()
}

type Schedule_OneTimeSchedule struct {
	OneTimeSchedule *OneTimeSchedule `protobuf:"bytes,2,opt,name=one_time_schedule,json=oneTimeSchedule,proto3,oneof"`
}

type Schedule_DailySchedule struct {
	DailySchedule *timeofday.TimeOfDay `protobuf:"bytes,3,opt,name=daily_schedule,json=dailySchedule,proto3,oneof"`
}

type Schedule_WeeklySchedule struct {
	WeeklySchedule *DayOfWeekAndTime `protobuf:"bytes,4,opt,name=weekly_schedule,json=weeklySchedule,proto3,oneof"`
}

func (*Schedule_OneTimeSchedule) isSchedule_ScheduledTime() {}

func (*Schedule_DailySchedule) isSchedule_ScheduledTime() {}

func (*Schedule_WeeklySchedule) isSchedule_ScheduledTime() {}

func (m *Schedule) GetScheduledTime() isSchedule_ScheduledTime {
	if m != nil {
		return m.ScheduledTime
	}
	return nil
}

func (m *Schedule) GetOneTimeSchedule() *OneTimeSchedule {
	if x, ok := m.GetScheduledTime().(*Schedule_OneTimeSchedule); ok {
		return x.OneTimeSchedule
	}
	return nil
}

func (m *Schedule) GetDailySchedule() *timeofday.TimeOfDay {
	if x, ok := m.GetScheduledTime().(*Schedule_DailySchedule); ok {
		return x.DailySchedule
	}
	return nil
}

func (m *Schedule) GetWeeklySchedule() *DayOfWeekAndTime {
	if x, ok := m.GetScheduledTime().(*Schedule_WeeklySchedule); ok {
		return x.WeeklySchedule
	}
	return nil
}

func (m *Schedule) GetTimeZone() TimeZone {
	if m != nil {
		return m.TimeZone
	}
	return TimeZone_UNKNOWN
}

func (m *Schedule) GetScheduleExecutor() Schedule_ScheduleExecutor {
	if m != nil {
		return m.ScheduleExecutor
	}
	return Schedule_UNKNOWN_EXECUTOR
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Schedule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Schedule_OneTimeSchedule)(nil),
		(*Schedule_DailySchedule)(nil),
		(*Schedule_WeeklySchedule)(nil),
	}
}

type OneTimeSchedule struct {
	// Types that are valid to be assigned to Value:
	//	*OneTimeSchedule_DateAndTime
	//	*OneTimeSchedule_DayOfWeekAndTime
	Value                isOneTimeSchedule_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *OneTimeSchedule) Reset()         { *m = OneTimeSchedule{} }
func (m *OneTimeSchedule) String() string { return proto.CompactTextString(m) }
func (*OneTimeSchedule) ProtoMessage()    {}
func (*OneTimeSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c54d7c732fc3473e, []int{1}
}

func (m *OneTimeSchedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneTimeSchedule.Unmarshal(m, b)
}
func (m *OneTimeSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneTimeSchedule.Marshal(b, m, deterministic)
}
func (m *OneTimeSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneTimeSchedule.Merge(m, src)
}
func (m *OneTimeSchedule) XXX_Size() int {
	return xxx_messageInfo_OneTimeSchedule.Size(m)
}
func (m *OneTimeSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_OneTimeSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_OneTimeSchedule proto.InternalMessageInfo

type isOneTimeSchedule_Value interface {
	isOneTimeSchedule_Value()
}

type OneTimeSchedule_DateAndTime struct {
	DateAndTime *DateAndTime `protobuf:"bytes,1,opt,name=date_and_time,json=dateAndTime,proto3,oneof"`
}

type OneTimeSchedule_DayOfWeekAndTime struct {
	DayOfWeekAndTime *DayOfWeekAndTime `protobuf:"bytes,2,opt,name=day_of_week_and_time,json=dayOfWeekAndTime,proto3,oneof"`
}

func (*OneTimeSchedule_DateAndTime) isOneTimeSchedule_Value() {}

func (*OneTimeSchedule_DayOfWeekAndTime) isOneTimeSchedule_Value() {}

func (m *OneTimeSchedule) GetValue() isOneTimeSchedule_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *OneTimeSchedule) GetDateAndTime() *DateAndTime {
	if x, ok := m.GetValue().(*OneTimeSchedule_DateAndTime); ok {
		return x.DateAndTime
	}
	return nil
}

func (m *OneTimeSchedule) GetDayOfWeekAndTime() *DayOfWeekAndTime {
	if x, ok := m.GetValue().(*OneTimeSchedule_DayOfWeekAndTime); ok {
		return x.DayOfWeekAndTime
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OneTimeSchedule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OneTimeSchedule_DateAndTime)(nil),
		(*OneTimeSchedule_DayOfWeekAndTime)(nil),
	}
}

type DateAndTime struct {
	YearMonthDate        *date.Date           `protobuf:"bytes,1,opt,name=year_month_date,json=yearMonthDate,proto3" json:"year_month_date,omitempty"`
	TimeOfDay            *timeofday.TimeOfDay `protobuf:"bytes,2,opt,name=time_of_day,json=timeOfDay,proto3" json:"time_of_day,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DateAndTime) Reset()         { *m = DateAndTime{} }
func (m *DateAndTime) String() string { return proto.CompactTextString(m) }
func (*DateAndTime) ProtoMessage()    {}
func (*DateAndTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_c54d7c732fc3473e, []int{2}
}

func (m *DateAndTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateAndTime.Unmarshal(m, b)
}
func (m *DateAndTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateAndTime.Marshal(b, m, deterministic)
}
func (m *DateAndTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateAndTime.Merge(m, src)
}
func (m *DateAndTime) XXX_Size() int {
	return xxx_messageInfo_DateAndTime.Size(m)
}
func (m *DateAndTime) XXX_DiscardUnknown() {
	xxx_messageInfo_DateAndTime.DiscardUnknown(m)
}

var xxx_messageInfo_DateAndTime proto.InternalMessageInfo

func (m *DateAndTime) GetYearMonthDate() *date.Date {
	if m != nil {
		return m.YearMonthDate
	}
	return nil
}

func (m *DateAndTime) GetTimeOfDay() *timeofday.TimeOfDay {
	if m != nil {
		return m.TimeOfDay
	}
	return nil
}

type DayOfWeekAndTime struct {
	DayOfWeek            dayofweek.DayOfWeek  `protobuf:"varint,1,opt,name=day_of_week,json=dayOfWeek,proto3,enum=google.type.DayOfWeek" json:"day_of_week,omitempty"`
	TimeOfDay            *timeofday.TimeOfDay `protobuf:"bytes,2,opt,name=time_of_day,json=timeOfDay,proto3" json:"time_of_day,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DayOfWeekAndTime) Reset()         { *m = DayOfWeekAndTime{} }
func (m *DayOfWeekAndTime) String() string { return proto.CompactTextString(m) }
func (*DayOfWeekAndTime) ProtoMessage()    {}
func (*DayOfWeekAndTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_c54d7c732fc3473e, []int{3}
}

func (m *DayOfWeekAndTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DayOfWeekAndTime.Unmarshal(m, b)
}
func (m *DayOfWeekAndTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DayOfWeekAndTime.Marshal(b, m, deterministic)
}
func (m *DayOfWeekAndTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DayOfWeekAndTime.Merge(m, src)
}
func (m *DayOfWeekAndTime) XXX_Size() int {
	return xxx_messageInfo_DayOfWeekAndTime.Size(m)
}
func (m *DayOfWeekAndTime) XXX_DiscardUnknown() {
	xxx_messageInfo_DayOfWeekAndTime.DiscardUnknown(m)
}

var xxx_messageInfo_DayOfWeekAndTime proto.InternalMessageInfo

func (m *DayOfWeekAndTime) GetDayOfWeek() dayofweek.DayOfWeek {
	if m != nil {
		return m.DayOfWeek
	}
	return dayofweek.DayOfWeek_DAY_OF_WEEK_UNSPECIFIED
}

func (m *DayOfWeekAndTime) GetTimeOfDay() *timeofday.TimeOfDay {
	if m != nil {
		return m.TimeOfDay
	}
	return nil
}

func init() {
	proto.RegisterEnum("autonomic.ext.common.schedule.TimeZone", TimeZone_name, TimeZone_value)
	proto.RegisterEnum("autonomic.ext.common.schedule.Schedule_ScheduleType", Schedule_ScheduleType_name, Schedule_ScheduleType_value)
	proto.RegisterEnum("autonomic.ext.common.schedule.Schedule_ScheduleExecutor", Schedule_ScheduleExecutor_name, Schedule_ScheduleExecutor_value)
	proto.RegisterType((*Schedule)(nil), "autonomic.ext.common.schedule.Schedule")
	proto.RegisterType((*OneTimeSchedule)(nil), "autonomic.ext.common.schedule.OneTimeSchedule")
	proto.RegisterType((*DateAndTime)(nil), "autonomic.ext.common.schedule.DateAndTime")
	proto.RegisterType((*DayOfWeekAndTime)(nil), "autonomic.ext.common.schedule.DayOfWeekAndTime")
}

func init() {
	proto.RegisterFile("autonomic/ext/common/schedule.proto", fileDescriptor_c54d7c732fc3473e)
}

var fileDescriptor_c54d7c732fc3473e = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x6e, 0xda, 0x75, 0x6b, 0x4f, 0xd7, 0xd6, 0xb3, 0xa6, 0xa9, 0x1a, 0x42, 0x1a, 0xe5, 0x82,
	0x69, 0xd2, 0x12, 0x34, 0x10, 0x82, 0xab, 0xa9, 0x3f, 0x11, 0xab, 0xb6, 0x26, 0xc3, 0x4b, 0x19,
	0x4c, 0x48, 0x91, 0xd7, 0xb8, 0x5b, 0xb4, 0x26, 0x9e, 0xda, 0x14, 0x1a, 0xae, 0x10, 0x97, 0x3c,
	0x06, 0xaf, 0xc0, 0x83, 0xf0, 0x3a, 0x5c, 0x22, 0xbb, 0x49, 0x7f, 0xa2, 0xb1, 0x8a, 0xdd, 0xf9,
	0xd8, 0xdf, 0xf9, 0xfc, 0x1d, 0x9f, 0xcf, 0x07, 0x9e, 0xd2, 0x51, 0xc0, 0x7d, 0xee, 0xb9, 0x5d,
	0x8d, 0x8d, 0x03, 0xad, 0xcb, 0x3d, 0x8f, 0xfb, 0xda, 0xb0, 0x7b, 0xcd, 0x9c, 0x51, 0x9f, 0xa9,
	0xb7, 0x03, 0x1e, 0x70, 0xfc, 0x78, 0x0a, 0x52, 0xd9, 0x38, 0x50, 0x27, 0x20, 0x35, 0x06, 0x6d,
	0x6f, 0x5d, 0x71, 0x7e, 0xd5, 0x67, 0x5a, 0x10, 0xde, 0x32, 0xcd, 0xa1, 0x41, 0x94, 0xb6, 0xfd,
	0x68, 0x71, 0x3f, 0xe4, 0xbd, 0x2f, 0x8c, 0xdd, 0xdc, 0x75, 0x18, 0xb8, 0x1e, 0xe3, 0x3d, 0x87,
	0x86, 0x93, 0xc3, 0xea, 0xaf, 0x2c, 0xe4, 0xce, 0x22, 0x7a, 0xfc, 0x11, 0x8a, 0xf1, 0x55, 0xb6,
	0x40, 0x57, 0x94, 0x1d, 0x65, 0xb7, 0x74, 0xf0, 0x52, 0xbd, 0x57, 0x95, 0x7a, 0x96, 0x5c, 0x58,
	0xe1, 0x2d, 0x23, 0xeb, 0xc3, 0xb9, 0x08, 0x7f, 0x82, 0x0d, 0xee, 0x33, 0x5b, 0x5c, 0x6f, 0xc7,
	0x07, 0x95, 0xf4, 0x8e, 0xb2, 0x5b, 0x38, 0x50, 0x97, 0xd0, 0x9b, 0x3e, 0xb3, 0x5c, 0x8f, 0xc5,
	0xe4, 0x47, 0x29, 0x52, 0xe6, 0x8b, 0x5b, 0xf8, 0x10, 0x4a, 0x0e, 0x75, 0xfb, 0xe1, 0x8c, 0x3a,
	0x23, 0xa9, 0xb7, 0xd4, 0x49, 0xed, 0xaa, 0xa8, 0x46, 0x15, 0x29, 0x66, 0xaf, 0x49, 0xc3, 0xa3,
	0x14, 0x29, 0x4a, 0xfc, 0x94, 0xe0, 0x02, 0xca, 0xe2, 0xc5, 0xe6, 0x19, 0x56, 0x24, 0x83, 0xb6,
	0x44, 0x5c, 0x93, 0x86, 0x66, 0xef, 0x9c, 0xb1, 0x9b, 0x9a, 0xef, 0x08, 0xfe, 0xa3, 0x14, 0x29,
	0x4d, 0x98, 0xa6, 0xdc, 0x4d, 0xc8, 0xcb, 0xb2, 0xbf, 0x72, 0x9f, 0x55, 0xb2, 0xf2, 0x45, 0x9f,
	0x2d, 0x61, 0x15, 0x4c, 0x17, 0xdc, 0x67, 0x24, 0x17, 0x44, 0x2b, 0xcc, 0x60, 0x63, 0xda, 0x1b,
	0x36, 0x66, 0xdd, 0x51, 0xc0, 0x07, 0x95, 0x55, 0xc9, 0xf6, 0xfa, 0x7f, 0xfb, 0xa3, 0x47, 0xf9,
	0x04, 0x0d, 0x13, 0x3b, 0xd5, 0x77, 0xb0, 0x3e, 0xdf, 0x45, 0x5c, 0x80, 0xb5, 0x8e, 0x71, 0x6c,
	0x98, 0xe7, 0x06, 0x4a, 0xe1, 0x1c, 0xac, 0x18, 0xa6, 0xa1, 0x23, 0x05, 0xaf, 0x43, 0xce, 0x34,
	0x74, 0xdb, 0x6a, 0xb5, 0x75, 0x94, 0xc6, 0x45, 0xc8, 0xb7, 0xda, 0x6d, 0xbd, 0xd9, 0xaa, 0x59,
	0x3a, 0xca, 0x88, 0x90, 0xe8, 0x8d, 0x0e, 0x21, 0x2d, 0xe3, 0x2d, 0x5a, 0xa9, 0x1e, 0x02, 0x4a,
	0x5e, 0x8c, 0x37, 0x01, 0x45, 0xb4, 0xb6, 0xfe, 0x41, 0x6f, 0x74, 0x2c, 0x93, 0xa0, 0x14, 0x06,
	0x58, 0x6d, 0xea, 0xef, 0x5b, 0x0d, 0x71, 0x43, 0x1e, 0xb2, 0x8d, 0x13, 0xb3, 0xd3, 0x44, 0xe9,
	0x3a, 0x82, 0x52, 0xac, 0xd3, 0x91, 0x0e, 0xaa, 0xfe, 0x56, 0xa0, 0x9c, 0xb0, 0x05, 0x3e, 0x85,
	0xa2, 0xf8, 0x11, 0x36, 0xf5, 0x27, 0x20, 0x69, 0xde, 0xc2, 0xc1, 0xde, 0xd2, 0x06, 0x06, 0x6c,
	0xd6, 0xbb, 0x82, 0x33, 0x0b, 0x31, 0x85, 0x4d, 0x87, 0x86, 0x36, 0xef, 0xd9, 0xa2, 0xa3, 0x33,
	0xe2, 0xf4, 0x43, 0x9d, 0x81, 0x9c, 0xc4, 0x5e, 0x7d, 0x0d, 0xb2, 0x9f, 0x69, 0x7f, 0xc4, 0xaa,
	0xdf, 0x14, 0x28, 0xcc, 0x49, 0xc1, 0x6f, 0xa0, 0x1c, 0x32, 0x3a, 0xb0, 0x3d, 0xee, 0x07, 0xd7,
	0xb6, 0x50, 0x15, 0xd5, 0xb3, 0xb1, 0x60, 0x69, 0x91, 0x42, 0x8a, 0x02, 0xd9, 0x16, 0x40, 0x11,
	0xe2, 0x57, 0x50, 0x90, 0x7e, 0xe3, 0x3d, 0xdb, 0xa1, 0x61, 0xa4, 0xf6, 0x1f, 0x3f, 0x81, 0x48,
	0x6b, 0xca, 0x65, 0xf5, 0xbb, 0x02, 0x28, 0x29, 0x5a, 0x90, 0xcd, 0xbd, 0x41, 0x34, 0x10, 0xb6,
	0x12, 0x1a, 0xa2, 0x1c, 0x92, 0x9f, 0xd6, 0xf7, 0x50, 0x11, 0x7b, 0xcf, 0x21, 0x17, 0x9b, 0x7f,
	0xd1, 0x7b, 0x6b, 0x90, 0xe9, 0x58, 0x0d, 0xa4, 0xe0, 0x12, 0xc0, 0x89, 0xd9, 0xa8, 0x9d, 0x44,
	0xe6, 0xab, 0xff, 0x50, 0xe0, 0x49, 0x97, 0x7b, 0xf7, 0x77, 0xa3, 0x5e, 0x8c, 0x7d, 0x72, 0x2a,
	0xc6, 0xde, 0xa9, 0x72, 0x71, 0x7c, 0xe5, 0x06, 0xd7, 0xa3, 0x4b, 0x01, 0xd4, 0xa6, 0xa9, 0xfb,
	0xd4, 0x15, 0xc3, 0x99, 0x0d, 0x7c, 0xda, 0xdf, 0x97, 0x03, 0x72, 0xa8, 0x0d, 0x07, 0x5d, 0xcd,
	0xa3, 0xae, 0xaf, 0xc9, 0x58, 0xbb, 0x6b, 0x8a, 0xff, 0x51, 0x94, 0x9f, 0xe9, 0x4c, 0xad, 0x63,
	0x5d, 0xae, 0x4a, 0xd0, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xae, 0xfd, 0x46, 0xed,
	0x05, 0x00, 0x00,
}
