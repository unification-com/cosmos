// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mainchain/stream/v1/stream.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// StreamPeriod enumerates the valid periods for calculating flow rates
type StreamPeriod int32

const (
	// STREAM_PERIOD_UNSPECIFIED defines unspecified
	StreamPeriodUnspecified StreamPeriod = 0
	// STREAM_PERIOD_SECOND defines second
	StreamPeriodSecond StreamPeriod = 1
	// STREAM_PERIOD_MINUTE defines minute
	StreamPeriodMinute StreamPeriod = 2
	// STREAM_PERIOD_HOUR defines hour
	StreamPeriodHour StreamPeriod = 3
	// STREAM_PERIOD_DAY defines day
	StreamPeriodDay StreamPeriod = 4
	// STREAM_PERIOD_WEEK defines week
	StreamPeriodWeek StreamPeriod = 5
	// STREAM_PERIOD_MONTH defines month
	StreamPeriodMonth StreamPeriod = 6
	// STREAM_PERIOD_YEAR defines year
	StreamPeriodYear StreamPeriod = 7
)

var StreamPeriod_name = map[int32]string{
	0: "STREAM_PERIOD_UNSPECIFIED",
	1: "STREAM_PERIOD_SECOND",
	2: "STREAM_PERIOD_MINUTE",
	3: "STREAM_PERIOD_HOUR",
	4: "STREAM_PERIOD_DAY",
	5: "STREAM_PERIOD_WEEK",
	6: "STREAM_PERIOD_MONTH",
	7: "STREAM_PERIOD_YEAR",
}

var StreamPeriod_value = map[string]int32{
	"STREAM_PERIOD_UNSPECIFIED": 0,
	"STREAM_PERIOD_SECOND":      1,
	"STREAM_PERIOD_MINUTE":      2,
	"STREAM_PERIOD_HOUR":        3,
	"STREAM_PERIOD_DAY":         4,
	"STREAM_PERIOD_WEEK":        5,
	"STREAM_PERIOD_MONTH":       6,
	"STREAM_PERIOD_YEAR":        7,
}

func (x StreamPeriod) String() string {
	return proto.EnumName(StreamPeriod_name, int32(x))
}

func (StreamPeriod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_835c4cdaca46b43c, []int{0}
}

// Stream holds data about a stream
type Stream struct {
	// deposit tracks the total amount the user has deposited to cover the stream, including any updates to the stream
	Deposit types.Coin `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit"`
	// flow_rate is the current rate of nund per second
	FlowRate int64 `protobuf:"varint,2,opt,name=flow_rate,json=flowRate,proto3" json:"flow_rate,omitempty"`
	// create_time is the timestamp a stream was first created
	CreateTime time.Time `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3,stdtime" json:"create_time" yaml:"create_time"`
	// last_updated_time is the timestamp the stream was last updated - flow_rate, deposit etc.
	LastUpdatedTime time.Time `protobuf:"bytes,4,opt,name=last_updated_time,json=lastUpdatedTime,proto3,stdtime" json:"last_updated_time" yaml:"last_updated_time"`
	// last_outflow_time is the timestamp of the last claim. Allows for a start point to calculate the next claim
	LastOutflowTime time.Time `protobuf:"bytes,5,opt,name=last_outflow_time,json=lastOutflowTime,proto3,stdtime" json:"last_outflow_time" yaml:"last_outflow_time"`
	// deposit_zero_time is the timestamp for when the current deposited amount will run out
	DepositZeroTime time.Time `protobuf:"bytes,6,opt,name=deposit_zero_time,json=depositZeroTime,proto3,stdtime" json:"deposit_zero_time" yaml:"deposit_zero_time"`
	// total_streamed tracks the total amount streamed. Calculated when a stream withdraw/update/cancel occurs
	TotalStreamed types.Coin `protobuf:"bytes,7,opt,name=total_streamed,json=totalStreamed,proto3" json:"total_streamed"`
	// cancellable is whether a stream can be cancelled. Default is true, but will be false for example id eFUND is used
	Cancellable bool `protobuf:"varint,8,opt,name=cancellable,proto3" json:"cancellable,omitempty"`
}

func (m *Stream) Reset()         { *m = Stream{} }
func (m *Stream) String() string { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()    {}
func (*Stream) Descriptor() ([]byte, []int) {
	return fileDescriptor_835c4cdaca46b43c, []int{0}
}
func (m *Stream) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stream.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stream.Merge(m, src)
}
func (m *Stream) XXX_Size() int {
	return m.Size()
}
func (m *Stream) XXX_DiscardUnknown() {
	xxx_messageInfo_Stream.DiscardUnknown(m)
}

var xxx_messageInfo_Stream proto.InternalMessageInfo

func (m *Stream) GetDeposit() types.Coin {
	if m != nil {
		return m.Deposit
	}
	return types.Coin{}
}

func (m *Stream) GetFlowRate() int64 {
	if m != nil {
		return m.FlowRate
	}
	return 0
}

func (m *Stream) GetCreateTime() time.Time {
	if m != nil {
		return m.CreateTime
	}
	return time.Time{}
}

func (m *Stream) GetLastUpdatedTime() time.Time {
	if m != nil {
		return m.LastUpdatedTime
	}
	return time.Time{}
}

func (m *Stream) GetLastOutflowTime() time.Time {
	if m != nil {
		return m.LastOutflowTime
	}
	return time.Time{}
}

func (m *Stream) GetDepositZeroTime() time.Time {
	if m != nil {
		return m.DepositZeroTime
	}
	return time.Time{}
}

func (m *Stream) GetTotalStreamed() types.Coin {
	if m != nil {
		return m.TotalStreamed
	}
	return types.Coin{}
}

func (m *Stream) GetCancellable() bool {
	if m != nil {
		return m.Cancellable
	}
	return false
}

func init() {
	proto.RegisterEnum("mainchain.stream.v1.StreamPeriod", StreamPeriod_name, StreamPeriod_value)
	proto.RegisterType((*Stream)(nil), "mainchain.stream.v1.Stream")
}

func init() { proto.RegisterFile("mainchain/stream/v1/stream.proto", fileDescriptor_835c4cdaca46b43c) }

var fileDescriptor_835c4cdaca46b43c = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x6e, 0xd3, 0x40,
	0x18, 0xc5, 0xe3, 0x26, 0x4d, 0xcb, 0x14, 0xa8, 0x33, 0x2d, 0x90, 0xba, 0x92, 0x63, 0x55, 0x2c,
	0xaa, 0x0a, 0x6c, 0x42, 0x57, 0x74, 0xd7, 0x36, 0xae, 0x1a, 0xa1, 0x24, 0x95, 0x93, 0xa8, 0x6a,
	0x59, 0x58, 0x13, 0x67, 0x92, 0x8e, 0xb0, 0x3d, 0x96, 0x3d, 0x2e, 0x94, 0x13, 0xa0, 0xac, 0xba,
	0x46, 0xca, 0x8a, 0xcb, 0x74, 0xd9, 0x25, 0xab, 0x82, 0x9a, 0x13, 0xc0, 0x09, 0x90, 0x3d, 0x0e,
	0x38, 0x18, 0x89, 0x3f, 0x9b, 0xc8, 0xdf, 0xf3, 0xf7, 0xde, 0xef, 0xc5, 0x1a, 0x0d, 0x50, 0x1c,
	0x44, 0x5c, 0xeb, 0x0c, 0x11, 0x57, 0x0b, 0x98, 0x8f, 0x91, 0xa3, 0x9d, 0x57, 0x93, 0x27, 0xd5,
	0xf3, 0x29, 0xa3, 0x70, 0xe5, 0xc7, 0x86, 0x9a, 0xe8, 0xe7, 0x55, 0x49, 0xb6, 0x68, 0xe0, 0xd0,
	0x40, 0xeb, 0xa1, 0x00, 0x6b, 0xe7, 0xd5, 0x1e, 0x66, 0xa8, 0xaa, 0x59, 0x94, 0xb8, 0xdc, 0x24,
	0xad, 0xf1, 0xf7, 0x66, 0x3c, 0x69, 0x7c, 0x48, 0x5e, 0xad, 0x0e, 0xe9, 0x90, 0x72, 0x3d, 0x7a,
	0x4a, 0xd4, 0xca, 0x90, 0xd2, 0xa1, 0x8d, 0xb5, 0x78, 0xea, 0x85, 0x03, 0x8d, 0x11, 0x07, 0x07,
	0x0c, 0x39, 0x5e, 0xb2, 0x50, 0x42, 0x0e, 0x71, 0xa9, 0x16, 0xff, 0x72, 0x69, 0xe3, 0x6b, 0x01,
	0x14, 0xdb, 0x71, 0x25, 0xf8, 0x02, 0x2c, 0xf4, 0xb1, 0x47, 0x03, 0xc2, 0xca, 0x82, 0x22, 0x6c,
	0x2e, 0x3d, 0x5f, 0x53, 0x13, 0x68, 0xd4, 0x50, 0x4d, 0x1a, 0xaa, 0xfb, 0x94, 0xb8, 0x7b, 0x85,
	0xab, 0x9b, 0x4a, 0xce, 0x98, 0xee, 0xc3, 0x75, 0x70, 0x67, 0x60, 0xd3, 0x37, 0xa6, 0x8f, 0x18,
	0x2e, 0xcf, 0x29, 0xc2, 0x66, 0xde, 0x58, 0x8c, 0x04, 0x03, 0x31, 0x0c, 0x5f, 0x81, 0x25, 0xcb,
	0xc7, 0x88, 0x61, 0x33, 0xea, 0x53, 0xce, 0xc7, 0xd9, 0x92, 0xca, 0xcb, 0xaa, 0xd3, 0xb2, 0x6a,
	0x67, 0x5a, 0x76, 0x4f, 0x8e, 0xc2, 0xbf, 0xdd, 0x54, 0xe0, 0x05, 0x72, 0xec, 0x9d, 0x8d, 0x94,
	0x79, 0xe3, 0xf2, 0x73, 0x45, 0x30, 0x00, 0x57, 0x22, 0x03, 0xb4, 0x41, 0xc9, 0x46, 0x01, 0x33,
	0x43, 0xaf, 0x8f, 0x18, 0xee, 0x73, 0x44, 0xe1, 0x8f, 0x88, 0xc7, 0x09, 0xa2, 0xcc, 0x11, 0x99,
	0x08, 0x0e, 0x5a, 0x8e, 0xf4, 0x2e, 0x97, 0x67, 0x68, 0x34, 0x64, 0xf1, 0xff, 0x8d, 0x69, 0xf3,
	0xff, 0x45, 0x4b, 0x47, 0xa4, 0x68, 0x2d, 0x2e, 0x4f, 0x69, 0xc9, 0x07, 0x36, 0xdf, 0x61, 0x9f,
	0x72, 0x5a, 0xf1, 0x5f, 0x69, 0x99, 0x88, 0x84, 0x96, 0xe8, 0xa7, 0xd8, 0xa7, 0x31, 0xed, 0x00,
	0xdc, 0x67, 0x94, 0x21, 0xdb, 0xe4, 0x27, 0x14, 0xf7, 0xcb, 0x0b, 0x7f, 0x77, 0x0a, 0xee, 0xc5,
	0xb6, 0x76, 0xe2, 0x82, 0x0a, 0x58, 0xb2, 0x90, 0x6b, 0x61, 0xdb, 0x46, 0x3d, 0x1b, 0x97, 0x17,
	0x15, 0x61, 0x73, 0xd1, 0x48, 0x4b, 0x5b, 0x1f, 0xf2, 0xe0, 0x2e, 0x5f, 0x3f, 0xc2, 0x3e, 0xa1,
	0x7d, 0xb8, 0x03, 0xd6, 0xda, 0x1d, 0x43, 0xdf, 0x6d, 0x98, 0x47, 0xba, 0x51, 0x6f, 0xd5, 0xcc,
	0x6e, 0xb3, 0x7d, 0xa4, 0xef, 0xd7, 0x0f, 0xea, 0x7a, 0x4d, 0xcc, 0x49, 0xeb, 0xa3, 0xb1, 0xf2,
	0x28, 0x6d, 0xe8, 0xba, 0x81, 0x87, 0x2d, 0x32, 0x20, 0xb8, 0x0f, 0x9f, 0x81, 0xd5, 0x59, 0x6f,
	0x5b, 0xdf, 0x6f, 0x35, 0x6b, 0xa2, 0x20, 0x3d, 0x1c, 0x8d, 0x15, 0x98, 0xb6, 0xb5, 0xb1, 0x45,
	0xdd, 0xdf, 0x38, 0x1a, 0xf5, 0x66, 0xb7, 0xa3, 0x8b, 0x73, 0x59, 0x47, 0x83, 0xb8, 0x21, 0xc3,
	0xf0, 0x09, 0x80, 0xb3, 0x8e, 0xc3, 0x56, 0xd7, 0x10, 0xf3, 0xd2, 0xea, 0x68, 0xac, 0x88, 0xe9,
	0xfd, 0x43, 0x1a, 0xfa, 0x70, 0x0b, 0x94, 0x66, 0xb7, 0x6b, 0xbb, 0x27, 0x62, 0x41, 0x5a, 0x19,
	0x8d, 0x95, 0xe5, 0xf4, 0x72, 0x0d, 0x5d, 0x64, 0x93, 0x8f, 0x75, 0xfd, 0xa5, 0x38, 0x9f, 0x4d,
	0x3e, 0xc6, 0xf8, 0x35, 0x54, 0xc1, 0xca, 0x2f, 0xcd, 0x5b, 0xcd, 0xce, 0xa1, 0x58, 0x94, 0x1e,
	0x8c, 0xc6, 0x4a, 0x69, 0xa6, 0x38, 0x75, 0xd9, 0x59, 0x36, 0xfd, 0x44, 0xdf, 0x35, 0xc4, 0x85,
	0x6c, 0xfa, 0x09, 0x46, 0xbe, 0x54, 0x78, 0xff, 0x51, 0xce, 0xed, 0x35, 0xae, 0x6e, 0x65, 0xe1,
	0xfa, 0x56, 0x16, 0xbe, 0xdc, 0xca, 0xc2, 0xe5, 0x44, 0xce, 0x5d, 0x4f, 0xe4, 0xdc, 0xa7, 0x89,
	0x9c, 0x3b, 0xdd, 0x1e, 0x12, 0x76, 0x16, 0xf6, 0x54, 0x8b, 0x3a, 0x5a, 0xe8, 0x92, 0x01, 0xb1,
	0x10, 0x23, 0xd4, 0x7d, 0x1a, 0xcd, 0x3f, 0x6f, 0xc0, 0xb7, 0xd3, 0x3b, 0x90, 0x5d, 0x78, 0x38,
	0xe8, 0x15, 0xe3, 0x03, 0xba, 0xfd, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xd7, 0x6d, 0x8f, 0x24,
	0x05, 0x00, 0x00,
}

func (m *Stream) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stream) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stream) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Cancellable {
		i--
		if m.Cancellable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	{
		size, err := m.TotalStreamed.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStream(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.DepositZeroTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.DepositZeroTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintStream(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.LastOutflowTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastOutflowTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintStream(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x2a
	n4, err4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.LastUpdatedTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastUpdatedTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintStream(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x22
	n5, err5 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CreateTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreateTime):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintStream(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x1a
	if m.FlowRate != 0 {
		i = encodeVarintStream(dAtA, i, uint64(m.FlowRate))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStream(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintStream(dAtA []byte, offset int, v uint64) int {
	offset -= sovStream(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Stream) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Deposit.Size()
	n += 1 + l + sovStream(uint64(l))
	if m.FlowRate != 0 {
		n += 1 + sovStream(uint64(m.FlowRate))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreateTime)
	n += 1 + l + sovStream(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastUpdatedTime)
	n += 1 + l + sovStream(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastOutflowTime)
	n += 1 + l + sovStream(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.DepositZeroTime)
	n += 1 + l + sovStream(uint64(l))
	l = m.TotalStreamed.Size()
	n += 1 + l + sovStream(uint64(l))
	if m.Cancellable {
		n += 2
	}
	return n
}

func sovStream(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStream(x uint64) (n int) {
	return sovStream(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Stream) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStream
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Stream: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stream: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlowRate", wireType)
			}
			m.FlowRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FlowRate |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CreateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdatedTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.LastUpdatedTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastOutflowTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.LastOutflowTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositZeroTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.DepositZeroTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStreamed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStream
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStream
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalStreamed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cancellable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStream
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Cancellable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipStream(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStream
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStream(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStream
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStream
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStream
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStream
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStream
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStream
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStream        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStream          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStream = fmt.Errorf("proto: unexpected end of group")
)
