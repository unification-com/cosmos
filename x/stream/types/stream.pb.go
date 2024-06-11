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
	// last_outflow_time is the timestamp of the last claim. Allows for a start point to calculate the next claim
	LastOutflowTime time.Time `protobuf:"bytes,3,opt,name=last_outflow_time,json=lastOutflowTime,proto3,stdtime" json:"last_outflow_time" yaml:"last_outflow_time"`
	// deposit_zero_time is the timestamp for when the current deposited amount will run out
	DepositZeroTime time.Time `protobuf:"bytes,4,opt,name=deposit_zero_time,json=depositZeroTime,proto3,stdtime" json:"deposit_zero_time" yaml:"deposit_zero_time"`
	// cancellable is whether a stream can be cancelled. Default is true, but will be false for example id eFUND is used
	Cancellable bool `protobuf:"varint,5,opt,name=cancellable,proto3" json:"cancellable,omitempty"`
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
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xbd, 0x6e, 0xd3, 0x40,
	0x1c, 0xb7, 0x9b, 0xf4, 0x83, 0x2b, 0x52, 0x1d, 0xb7, 0x40, 0xea, 0x4a, 0x8e, 0x55, 0x31, 0x54,
	0x15, 0xd8, 0x84, 0x4e, 0x74, 0x6b, 0x13, 0xa3, 0x44, 0x28, 0x1f, 0x72, 0x12, 0x55, 0xe9, 0x62,
	0x9d, 0x9d, 0x4b, 0x72, 0xc2, 0xbe, 0x8b, 0xec, 0x73, 0x20, 0x3c, 0x01, 0xca, 0xd4, 0x19, 0x29,
	0x13, 0x2f, 0xd3, 0xb1, 0x23, 0x53, 0x40, 0xc9, 0x1b, 0xf0, 0x04, 0xc8, 0x1f, 0x81, 0x04, 0xb3,
	0xb0, 0x58, 0xf7, 0xff, 0xdd, 0xef, 0x4b, 0xd6, 0xfd, 0x81, 0xe2, 0x42, 0x4c, 0xec, 0x21, 0xc4,
	0x44, 0xf3, 0x99, 0x87, 0xa0, 0xab, 0x8d, 0x8b, 0xc9, 0x49, 0x1d, 0x79, 0x94, 0x51, 0xf1, 0xf0,
	0x37, 0x43, 0x4d, 0xf0, 0x71, 0x51, 0x92, 0x6d, 0xea, 0xbb, 0xd4, 0xd7, 0x2c, 0xe8, 0x23, 0x6d,
	0x5c, 0xb4, 0x10, 0x83, 0x45, 0xcd, 0xa6, 0x98, 0xc4, 0x22, 0xe9, 0x38, 0xbe, 0x37, 0xa3, 0x49,
	0x8b, 0x87, 0xe4, 0xea, 0x68, 0x40, 0x07, 0x34, 0xc6, 0xc3, 0x53, 0x82, 0x16, 0x06, 0x94, 0x0e,
	0x1c, 0xa4, 0x45, 0x93, 0x15, 0xf4, 0x35, 0x86, 0x5d, 0xe4, 0x33, 0xe8, 0x8e, 0x12, 0x42, 0x0e,
	0xba, 0x98, 0x50, 0x2d, 0xfa, 0xc6, 0xd0, 0xe9, 0x7c, 0x0b, 0xec, 0xb4, 0xa2, 0x4a, 0xe2, 0x1b,
	0xb0, 0xdb, 0x43, 0x23, 0xea, 0x63, 0x96, 0xe7, 0x15, 0xfe, 0x6c, 0xff, 0xf5, 0xb1, 0x9a, 0x84,
	0x86, 0x0d, 0xd5, 0xa4, 0xa1, 0x5a, 0xa2, 0x98, 0x5c, 0x67, 0xef, 0xe7, 0x05, 0xce, 0x58, 0xf1,
	0xc5, 0x13, 0xf0, 0xa8, 0xef, 0xd0, 0x0f, 0xa6, 0x07, 0x19, 0xca, 0x6f, 0x29, 0xfc, 0x59, 0xc6,
	0xd8, 0x0b, 0x01, 0x03, 0x32, 0x24, 0x3a, 0x20, 0xe7, 0x40, 0x9f, 0x99, 0x34, 0x60, 0x11, 0x29,
	0x6c, 0x95, 0xcf, 0x44, 0x09, 0x92, 0x1a, 0x57, 0x56, 0x57, 0x95, 0xd5, 0xf6, 0xaa, 0xf2, 0xf5,
	0xf3, 0x30, 0xe2, 0xe7, 0xbc, 0x90, 0x9f, 0x40, 0xd7, 0xb9, 0x3c, 0x4d, 0x59, 0x9c, 0xde, 0x7d,
	0x2f, 0xf0, 0xc6, 0x41, 0x88, 0x37, 0x62, 0x38, 0xd4, 0x86, 0x69, 0x49, 0x2b, 0xf3, 0x13, 0xf2,
	0x68, 0x9c, 0x96, 0xfd, 0xdf, 0xb4, 0x94, 0x45, 0x92, 0x96, 0xe0, 0xb7, 0xc8, 0xa3, 0x51, 0x9a,
	0x02, 0xf6, 0x6d, 0x48, 0x6c, 0xe4, 0x38, 0xd0, 0x72, 0x50, 0x7e, 0x5b, 0xe1, 0xcf, 0xf6, 0x8c,
	0x75, 0xe8, 0xfc, 0x4b, 0x06, 0x3c, 0x8e, 0x7f, 0x70, 0x13, 0x79, 0x98, 0xf6, 0xc4, 0x4b, 0x70,
	0xdc, 0x6a, 0x1b, 0xfa, 0x55, 0xcd, 0x6c, 0xea, 0x46, 0xb5, 0x51, 0x36, 0x3b, 0xf5, 0x56, 0x53,
	0x2f, 0x55, 0xdf, 0x56, 0xf5, 0xb2, 0xc0, 0x49, 0x27, 0xd3, 0x99, 0xf2, 0x6c, 0x5d, 0xd0, 0x21,
	0xfe, 0x08, 0xd9, 0xb8, 0x8f, 0x51, 0x4f, 0x7c, 0x05, 0x8e, 0x36, 0xb5, 0x2d, 0xbd, 0xd4, 0xa8,
	0x97, 0x05, 0x5e, 0x7a, 0x3a, 0x9d, 0x29, 0xe2, 0xba, 0xac, 0x85, 0x6c, 0x4a, 0xfe, 0xa1, 0xa8,
	0x55, 0xeb, 0x9d, 0xb6, 0x2e, 0x6c, 0xa5, 0x15, 0x35, 0x4c, 0x02, 0x86, 0xc4, 0x17, 0x40, 0xdc,
	0x54, 0x54, 0x1a, 0x1d, 0x43, 0xc8, 0x48, 0x47, 0xd3, 0x99, 0x22, 0xac, 0xf3, 0x2b, 0x34, 0xf0,
	0xc4, 0x73, 0x90, 0xdb, 0x64, 0x97, 0xaf, 0xba, 0x42, 0x56, 0x3a, 0x9c, 0xce, 0x94, 0x83, 0x75,
	0x72, 0x19, 0x4e, 0xd2, 0xce, 0x37, 0xba, 0xfe, 0x4e, 0xd8, 0x4e, 0x3b, 0xdf, 0x20, 0xf4, 0x5e,
	0x54, 0xc1, 0xe1, 0x5f, 0xcd, 0x1b, 0xf5, 0x76, 0x45, 0xd8, 0x91, 0x9e, 0x4c, 0x67, 0x4a, 0x6e,
	0xa3, 0x38, 0x25, 0x6c, 0x98, 0x76, 0xef, 0xea, 0x57, 0x86, 0xb0, 0x9b, 0x76, 0xef, 0x22, 0xe8,
	0x49, 0xd9, 0xcf, 0x5f, 0x65, 0xee, 0xba, 0x76, 0xbf, 0x90, 0xf9, 0x87, 0x85, 0xcc, 0xff, 0x58,
	0xc8, 0xfc, 0xdd, 0x52, 0xe6, 0x1e, 0x96, 0x32, 0xf7, 0x6d, 0x29, 0x73, 0xb7, 0x17, 0x03, 0xcc,
	0x86, 0x81, 0xa5, 0xda, 0xd4, 0xd5, 0x02, 0x82, 0xfb, 0xd8, 0x86, 0x0c, 0x53, 0xf2, 0x32, 0x9c,
	0xff, 0xac, 0xfb, 0xc7, 0xd5, 0xc2, 0xb3, 0xc9, 0x08, 0xf9, 0xd6, 0x4e, 0xf4, 0xb0, 0x2e, 0x7e,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xb4, 0x3c, 0x37, 0x11, 0x04, 0x00, 0x00,
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
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.DepositZeroTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.DepositZeroTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintStream(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.LastOutflowTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastOutflowTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintStream(dAtA, i, uint64(n2))
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
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastOutflowTime)
	n += 1 + l + sovStream(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.DepositZeroTime)
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
		case 4:
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
		case 5:
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
