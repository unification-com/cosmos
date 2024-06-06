// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mainchain/stream/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the stream module's genesis state.
type GenesisState struct {
	Params  Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Streams []StreamExport `protobuf:"bytes,2,rep,name=streams,proto3" json:"streams"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d898a81da94e0d10, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetStreams() []StreamExport {
	if m != nil {
		return m.Streams
	}
	return nil
}

// StreamExport holds genesis export data for a beacon, including submitted timestamps
type StreamExport struct {
	// receiver is the wallet that will receive stream payments
	Receiver string `protobuf:"bytes,1,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// sender is the wallet making the update
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	// stream is the stream data
	Stream Stream `protobuf:"bytes,3,opt,name=stream,proto3" json:"stream"`
}

func (m *StreamExport) Reset()         { *m = StreamExport{} }
func (m *StreamExport) String() string { return proto.CompactTextString(m) }
func (*StreamExport) ProtoMessage()    {}
func (*StreamExport) Descriptor() ([]byte, []int) {
	return fileDescriptor_d898a81da94e0d10, []int{1}
}
func (m *StreamExport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamExport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamExport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamExport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamExport.Merge(m, src)
}
func (m *StreamExport) XXX_Size() int {
	return m.Size()
}
func (m *StreamExport) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamExport.DiscardUnknown(m)
}

var xxx_messageInfo_StreamExport proto.InternalMessageInfo

func (m *StreamExport) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *StreamExport) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *StreamExport) GetStream() Stream {
	if m != nil {
		return m.Stream
	}
	return Stream{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "mainchain.stream.v1.GenesisState")
	proto.RegisterType((*StreamExport)(nil), "mainchain.stream.v1.StreamExport")
}

func init() { proto.RegisterFile("mainchain/stream/v1/genesis.proto", fileDescriptor_d898a81da94e0d10) }

var fileDescriptor_d898a81da94e0d10 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x4f, 0x3a, 0x31,
	0x18, 0xc6, 0xaf, 0xf0, 0x0f, 0x7f, 0x2d, 0x4c, 0x27, 0xc3, 0x89, 0x49, 0x05, 0x26, 0x16, 0x7a,
	0x02, 0x2e, 0x8e, 0x90, 0x18, 0x27, 0x13, 0x03, 0x9b, 0x8b, 0x29, 0x47, 0x3d, 0x3a, 0x5c, 0x7b,
	0x69, 0x0b, 0xc1, 0xef, 0xe0, 0xe0, 0x57, 0x31, 0xf1, 0x43, 0x30, 0x12, 0x27, 0x27, 0x63, 0xe0,
	0x8b, 0x98, 0x6b, 0x7b, 0xe8, 0x70, 0xc1, 0xad, 0xef, 0x3d, 0xbf, 0xe7, 0x7d, 0x9f, 0xf7, 0x5e,
	0xd8, 0x4a, 0x08, 0xe3, 0xd1, 0x9c, 0x30, 0x1e, 0x2a, 0x2d, 0x29, 0x49, 0xc2, 0x65, 0x2f, 0x8c,
	0x29, 0xa7, 0x8a, 0x29, 0x9c, 0x4a, 0xa1, 0x85, 0x7f, 0xb2, 0x47, 0xb0, 0x45, 0xf0, 0xb2, 0xd7,
	0xa8, 0xc7, 0x22, 0x16, 0x46, 0x0f, 0xb3, 0x97, 0x45, 0x1b, 0xa7, 0x91, 0x50, 0x89, 0x50, 0x0f,
	0x56, 0xb0, 0x85, 0x93, 0x9a, 0x45, 0x83, 0x52, 0x22, 0x49, 0x72, 0x90, 0x70, 0x13, 0x0d, 0xd1,
	0x7e, 0x06, 0xb0, 0x76, 0x63, 0xb3, 0x4d, 0x34, 0xd1, 0xd4, 0xbf, 0x82, 0x15, 0xdb, 0x22, 0x00,
	0x4d, 0xd0, 0xa9, 0xf6, 0xcf, 0x70, 0x41, 0x56, 0x7c, 0x67, 0x90, 0xd1, 0xbf, 0xf5, 0xe7, 0xb9,
	0x37, 0x76, 0x06, 0x7f, 0x08, 0xff, 0x5b, 0x42, 0x05, 0xa5, 0x66, 0xb9, 0x53, 0xed, 0xb7, 0x0a,
	0xbd, 0x13, 0xf3, 0xba, 0x5e, 0xa5, 0x42, 0x6a, 0xd7, 0x21, 0xf7, 0xb5, 0x5f, 0x01, 0xac, 0xfd,
	0xd6, 0xfd, 0x4b, 0x78, 0x24, 0x69, 0x44, 0xd9, 0x92, 0x4a, 0x13, 0xe8, 0x78, 0x14, 0xbc, 0xbf,
	0x75, 0xeb, 0xee, 0x3f, 0x0c, 0x67, 0x33, 0x49, 0x95, 0x9a, 0x68, 0xc9, 0x78, 0x3c, 0xde, 0x93,
	0xfe, 0x05, 0xac, 0x28, 0xca, 0x67, 0x54, 0x06, 0xa5, 0x3f, 0x3c, 0x8e, 0xcb, 0xd6, 0xb6, 0x19,
	0x82, 0xf2, 0x81, 0xb5, 0x6d, 0xb4, 0x7c, 0x6d, 0xfb, 0x7d, 0x74, 0xbb, 0xde, 0x22, 0xb0, 0xd9,
	0x22, 0xf0, 0xb5, 0x45, 0xe0, 0x65, 0x87, 0xbc, 0xcd, 0x0e, 0x79, 0x1f, 0x3b, 0xe4, 0xdd, 0x0f,
	0x62, 0xa6, 0xe7, 0x8b, 0x29, 0x8e, 0x44, 0x12, 0x2e, 0x38, 0x7b, 0x64, 0x11, 0xd1, 0x4c, 0xf0,
	0x6e, 0x56, 0xff, 0x5c, 0x66, 0x95, 0xdf, 0x46, 0x3f, 0xa5, 0x54, 0x4d, 0x2b, 0xe6, 0x30, 0x83,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x6e, 0x3d, 0x2e, 0x47, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Streams) > 0 {
		for iNdEx := len(m.Streams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Streams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *StreamExport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamExport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamExport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Stream.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Streams) > 0 {
		for _, e := range m.Streams {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *StreamExport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Stream.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Streams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Streams = append(m.Streams, StreamExport{})
			if err := m.Streams[len(m.Streams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *StreamExport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: StreamExport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamExport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stream", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stream.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
