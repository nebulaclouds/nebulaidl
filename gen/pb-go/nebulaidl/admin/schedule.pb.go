// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/admin/schedule.proto

package admin

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

// Represents a frequency at which to run a schedule.
type FixedRateUnit int32

const (
	FixedRateUnit_MINUTE FixedRateUnit = 0
	FixedRateUnit_HOUR   FixedRateUnit = 1
	FixedRateUnit_DAY    FixedRateUnit = 2
)

var FixedRateUnit_name = map[int32]string{
	0: "MINUTE",
	1: "HOUR",
	2: "DAY",
}

var FixedRateUnit_value = map[string]int32{
	"MINUTE": 0,
	"HOUR":   1,
	"DAY":    2,
}

func (x FixedRateUnit) String() string {
	return proto.EnumName(FixedRateUnit_name, int32(x))
}

func (FixedRateUnit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_556686655346ecb6, []int{0}
}

// Option for schedules run at a certain frequency e.g. every 2 minutes.
type FixedRate struct {
	Value                uint32        `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit                 FixedRateUnit `protobuf:"varint,2,opt,name=unit,proto3,enum=nebulaidl.admin.FixedRateUnit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FixedRate) Reset()         { *m = FixedRate{} }
func (m *FixedRate) String() string { return proto.CompactTextString(m) }
func (*FixedRate) ProtoMessage()    {}
func (*FixedRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_556686655346ecb6, []int{0}
}

func (m *FixedRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixedRate.Unmarshal(m, b)
}
func (m *FixedRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixedRate.Marshal(b, m, deterministic)
}
func (m *FixedRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedRate.Merge(m, src)
}
func (m *FixedRate) XXX_Size() int {
	return xxx_messageInfo_FixedRate.Size(m)
}
func (m *FixedRate) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedRate.DiscardUnknown(m)
}

var xxx_messageInfo_FixedRate proto.InternalMessageInfo

func (m *FixedRate) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *FixedRate) GetUnit() FixedRateUnit {
	if m != nil {
		return m.Unit
	}
	return FixedRateUnit_MINUTE
}

// Options for schedules to run according to a cron expression.
type CronSchedule struct {
	// Standard/default cron implementation as described by https://en.wikipedia.org/wiki/Cron#CRON_expression;
	// Also supports nonstandard predefined scheduling definitions
	// as described by https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions
	// except @reboot
	Schedule string `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
	// ISO 8601 duration as described by https://en.wikipedia.org/wiki/ISO_8601#Durations
	Offset               string   `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CronSchedule) Reset()         { *m = CronSchedule{} }
func (m *CronSchedule) String() string { return proto.CompactTextString(m) }
func (*CronSchedule) ProtoMessage()    {}
func (*CronSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_556686655346ecb6, []int{1}
}

func (m *CronSchedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CronSchedule.Unmarshal(m, b)
}
func (m *CronSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CronSchedule.Marshal(b, m, deterministic)
}
func (m *CronSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronSchedule.Merge(m, src)
}
func (m *CronSchedule) XXX_Size() int {
	return xxx_messageInfo_CronSchedule.Size(m)
}
func (m *CronSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_CronSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_CronSchedule proto.InternalMessageInfo

func (m *CronSchedule) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *CronSchedule) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

// Defines complete set of information required to trigger an execution on a schedule.
type Schedule struct {
	// Types that are valid to be assigned to ScheduleExpression:
	//	*Schedule_CronExpression
	//	*Schedule_Rate
	//	*Schedule_CronSchedule
	ScheduleExpression isSchedule_ScheduleExpression `protobuf_oneof:"ScheduleExpression"`
	// Name of the input variable that the kickoff time will be supplied to when the workflow is kicked off.
	KickoffTimeInputArg  string   `protobuf:"bytes,3,opt,name=kickoff_time_input_arg,json=kickoffTimeInputArg,proto3" json:"kickoff_time_input_arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_556686655346ecb6, []int{2}
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

type isSchedule_ScheduleExpression interface {
	isSchedule_ScheduleExpression()
}

type Schedule_CronExpression struct {
	CronExpression string `protobuf:"bytes,1,opt,name=cron_expression,json=cronExpression,proto3,oneof"`
}

type Schedule_Rate struct {
	Rate *FixedRate `protobuf:"bytes,2,opt,name=rate,proto3,oneof"`
}

type Schedule_CronSchedule struct {
	CronSchedule *CronSchedule `protobuf:"bytes,4,opt,name=cron_schedule,json=cronSchedule,proto3,oneof"`
}

func (*Schedule_CronExpression) isSchedule_ScheduleExpression() {}

func (*Schedule_Rate) isSchedule_ScheduleExpression() {}

func (*Schedule_CronSchedule) isSchedule_ScheduleExpression() {}

func (m *Schedule) GetScheduleExpression() isSchedule_ScheduleExpression {
	if m != nil {
		return m.ScheduleExpression
	}
	return nil
}

// Deprecated: Do not use.
func (m *Schedule) GetCronExpression() string {
	if x, ok := m.GetScheduleExpression().(*Schedule_CronExpression); ok {
		return x.CronExpression
	}
	return ""
}

func (m *Schedule) GetRate() *FixedRate {
	if x, ok := m.GetScheduleExpression().(*Schedule_Rate); ok {
		return x.Rate
	}
	return nil
}

func (m *Schedule) GetCronSchedule() *CronSchedule {
	if x, ok := m.GetScheduleExpression().(*Schedule_CronSchedule); ok {
		return x.CronSchedule
	}
	return nil
}

func (m *Schedule) GetKickoffTimeInputArg() string {
	if m != nil {
		return m.KickoffTimeInputArg
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Schedule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Schedule_CronExpression)(nil),
		(*Schedule_Rate)(nil),
		(*Schedule_CronSchedule)(nil),
	}
}

func init() {
	proto.RegisterEnum("nebulaidl.admin.FixedRateUnit", FixedRateUnit_name, FixedRateUnit_value)
	proto.RegisterType((*FixedRate)(nil), "nebulaidl.admin.FixedRate")
	proto.RegisterType((*CronSchedule)(nil), "nebulaidl.admin.CronSchedule")
	proto.RegisterType((*Schedule)(nil), "nebulaidl.admin.Schedule")
}

func init() { proto.RegisterFile("nebulaidl/admin/schedule.proto", fileDescriptor_556686655346ecb6) }

var fileDescriptor_556686655346ecb6 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4b, 0xeb, 0xda, 0x40,
	0x14, 0xc5, 0x13, 0x4d, 0xad, 0xde, 0xfa, 0x62, 0x2a, 0x12, 0x84, 0x8a, 0xb8, 0x92, 0x82, 0x49,
	0xd1, 0x65, 0xe9, 0xc2, 0x54, 0x4b, 0x5c, 0xb4, 0x85, 0xa9, 0x59, 0xb4, 0x9b, 0x90, 0xc7, 0x24,
	0x0e, 0x26, 0x33, 0x21, 0x99, 0x14, 0x3f, 0x7c, 0x17, 0xc5, 0x89, 0x46, 0x6b, 0xf9, 0x2f, 0x2f,
	0x67, 0xee, 0xef, 0x9e, 0x33, 0x1c, 0x98, 0x32, 0xe2, 0x97, 0x89, 0x47, 0xc3, 0xc4, 0xf4, 0xc2,
	0x94, 0x32, 0xb3, 0x08, 0x8e, 0x24, 0x2c, 0x13, 0x62, 0x64, 0x39, 0x17, 0x1c, 0x0d, 0x6a, 0xdd,
	0x90, 0xfa, 0xdc, 0x81, 0xce, 0x17, 0x7a, 0x26, 0x21, 0xf6, 0x04, 0x41, 0x23, 0x78, 0xf5, 0xdb,
	0x4b, 0x4a, 0xa2, 0xab, 0x33, 0x75, 0xd1, 0xc3, 0xd5, 0x80, 0x56, 0xa0, 0x95, 0x8c, 0x0a, 0xbd,
	0x31, 0x53, 0x17, 0xfd, 0xd5, 0xd4, 0x78, 0x42, 0x18, 0xf5, 0xbe, 0xc3, 0xa8, 0xc0, 0xf2, 0xed,
	0xdc, 0x82, 0xee, 0xe7, 0x9c, 0xb3, 0x1f, 0xd7, 0xeb, 0x68, 0x02, 0xed, 0x9b, 0x13, 0x09, 0xef,
	0xe0, 0x7a, 0x46, 0x63, 0x68, 0xf1, 0x28, 0x2a, 0x48, 0x75, 0xa1, 0x83, 0xaf, 0xd3, 0xfc, 0x8f,
	0x0a, 0xed, 0x1a, 0xb0, 0x84, 0x41, 0x90, 0x73, 0xe6, 0x92, 0x73, 0x96, 0x93, 0xa2, 0xa0, 0x9c,
	0x55, 0x1c, 0xab, 0xa1, 0xab, 0xb6, 0x82, 0xfb, 0x17, 0x71, 0x57, 0x6b, 0xe8, 0x03, 0x68, 0xb9,
	0x27, 0x88, 0x24, 0xbe, 0x59, 0x4d, 0x5e, 0xf6, 0x6c, 0x2b, 0x58, 0xbe, 0x44, 0x5b, 0xe8, 0xc9,
	0x03, 0xb5, 0x4d, 0x4d, 0xae, 0xbe, 0xfb, 0x6f, 0xf5, 0x31, 0x97, 0xad, 0xe0, 0x6e, 0xf0, 0x98,
	0x73, 0x0d, 0xe3, 0x13, 0x0d, 0x4e, 0x3c, 0x8a, 0x5c, 0x41, 0x53, 0xe2, 0x52, 0x96, 0x95, 0xc2,
	0xf5, 0xf2, 0x58, 0x6f, 0xca, 0x6c, 0x6f, 0xaf, 0xea, 0x81, 0xa6, 0x64, 0x7f, 0xd1, 0x36, 0x79,
	0x6c, 0x8d, 0x00, 0xdd, 0x00, 0xf7, 0x08, 0xef, 0x0d, 0xe8, 0xfd, 0xf3, 0xb3, 0x08, 0xa0, 0xf5,
	0x75, 0xff, 0xcd, 0x39, 0xec, 0x86, 0x0a, 0x6a, 0x83, 0x66, 0x7f, 0x77, 0xf0, 0x50, 0x45, 0xaf,
	0xa1, 0xb9, 0xdd, 0xfc, 0x1c, 0x36, 0xac, 0x4f, 0xbf, 0x3e, 0xc6, 0x54, 0x1c, 0x4b, 0xdf, 0x08,
	0x78, 0x6a, 0x56, 0xae, 0x83, 0x84, 0x97, 0x61, 0x61, 0xde, 0x4b, 0x11, 0x13, 0x66, 0x66, 0xfe,
	0x32, 0xe6, 0xe6, 0x53, 0x51, 0xfc, 0x96, 0x2c, 0xc8, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe8, 0x48, 0x33, 0x6f, 0x42, 0x02, 0x00, 0x00,
}
