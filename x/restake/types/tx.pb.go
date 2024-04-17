// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: restake/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/bandprotocol/chain/v2/x/oracle/types"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgUpdateParams is the Msg/WithdrawRewards request type.
type MsgWithdrawRewards struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgWithdrawRewards) Reset()         { *m = MsgWithdrawRewards{} }
func (m *MsgWithdrawRewards) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawRewards) ProtoMessage()    {}
func (*MsgWithdrawRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6a768b7b224a47, []int{0}
}
func (m *MsgWithdrawRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawRewards.Merge(m, src)
}
func (m *MsgWithdrawRewards) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawRewards.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawRewards proto.InternalMessageInfo

func (m *MsgWithdrawRewards) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// MsgWithdrawRewardsResponse defines ...
type MsgWithdrawRewardsResponse struct {
}

func (m *MsgWithdrawRewardsResponse) Reset()         { *m = MsgWithdrawRewardsResponse{} }
func (m *MsgWithdrawRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawRewardsResponse) ProtoMessage()    {}
func (*MsgWithdrawRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6a768b7b224a47, []int{1}
}
func (m *MsgWithdrawRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawRewardsResponse.Merge(m, src)
}
func (m *MsgWithdrawRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawRewardsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgWithdrawRewards)(nil), "restake.v1beta1.MsgWithdrawRewards")
	proto.RegisterType((*MsgWithdrawRewardsResponse)(nil), "restake.v1beta1.MsgWithdrawRewardsResponse")
}

func init() { proto.RegisterFile("restake/v1beta1/tx.proto", fileDescriptor_4c6a768b7b224a47) }

var fileDescriptor_4c6a768b7b224a47 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4a, 0x03, 0x41,
	0x10, 0xc6, 0x73, 0x08, 0x8a, 0xd7, 0x04, 0x8f, 0xa0, 0xf1, 0x90, 0x45, 0x62, 0x23, 0x09, 0xde,
	0x9a, 0xd8, 0xd9, 0x99, 0x52, 0x48, 0x13, 0x0b, 0xc1, 0x46, 0xf6, 0xf6, 0x96, 0xcd, 0x69, 0x6e,
	0x27, 0xec, 0x6c, 0xfe, 0xd8, 0x5a, 0x5a, 0xf9, 0x28, 0x29, 0x7c, 0x08, 0xcb, 0x60, 0x65, 0x29,
	0x49, 0x91, 0xd7, 0x90, 0xdc, 0xee, 0x21, 0xe4, 0x1a, 0x9b, 0x63, 0x66, 0x7e, 0xdf, 0x7c, 0x73,
	0xfb, 0xf9, 0x75, 0x2d, 0xd0, 0xb0, 0x67, 0x41, 0x27, 0xed, 0x58, 0x18, 0xd6, 0xa6, 0x66, 0x16,
	0x8d, 0x34, 0x18, 0x08, 0xaa, 0x8e, 0x44, 0x8e, 0x84, 0x35, 0x09, 0x12, 0x72, 0x46, 0x37, 0x95,
	0x95, 0x85, 0x87, 0xa0, 0x19, 0x1f, 0x6e, 0xf6, 0xa9, 0xad, 0xdc, 0x9c, 0x70, 0xc0, 0x0c, 0x90,
	0xc6, 0x0c, 0xff, 0xcc, 0x39, 0xa4, 0xca, 0xf1, 0x63, 0xcb, 0x1f, 0xad, 0xa1, 0x6d, 0x1c, 0x3a,
	0x72, 0xab, 0x19, 0xca, 0x8d, 0x6d, 0x86, 0xd2, 0x81, 0x03, 0x96, 0xa5, 0x0a, 0x68, 0xfe, 0xb5,
	0xa3, 0xc6, 0xd8, 0x0f, 0x7a, 0x28, 0xef, 0x53, 0x33, 0x48, 0x34, 0x9b, 0xf6, 0xc5, 0x94, 0xe9,
	0x04, 0x83, 0x8e, 0xbf, 0xc7, 0x92, 0x44, 0x0b, 0xc4, 0xba, 0x77, 0xea, 0x9d, 0xef, 0x77, 0xeb,
	0x5f, 0x1f, 0x17, 0x35, 0x77, 0xe4, 0xc6, 0x92, 0x3b, 0xa3, 0x53, 0x25, 0xfb, 0x85, 0xf0, 0xba,
	0xf5, 0xba, 0x9e, 0x37, 0x8b, 0xee, 0x6d, 0x3d, 0x6f, 0x86, 0x45, 0x34, 0xe5, 0x03, 0x8d, 0x13,
	0x3f, 0x2c, 0x4f, 0xfb, 0x02, 0x47, 0xa0, 0x50, 0x74, 0x9e, 0xfc, 0x9d, 0x1e, 0xca, 0x80, 0xfb,
	0xd5, 0xed, 0x1f, 0x3b, 0x8b, 0xb6, 0x52, 0x8d, 0xca, 0x36, 0x61, 0xeb, 0x1f, 0xa2, 0xe2, 0x56,
	0xf7, 0xf6, 0x73, 0x49, 0xbc, 0xc5, 0x92, 0x78, 0x3f, 0x4b, 0xe2, 0xbd, 0xaf, 0x48, 0x65, 0xb1,
	0x22, 0x95, 0xef, 0x15, 0xa9, 0x3c, 0x5c, 0xca, 0xd4, 0x0c, 0xc6, 0x71, 0xc4, 0x21, 0xa3, 0x31,
	0x53, 0x49, 0x1e, 0x18, 0x87, 0x21, 0xe5, 0x03, 0x96, 0x2a, 0x3a, 0xe9, 0xd0, 0x19, 0x2d, 0x9e,
	0x68, 0x5e, 0x46, 0x02, 0xe3, 0xdd, 0x5c, 0x72, 0xf5, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x9b,
	0x5e, 0xba, 0x15, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	WithdrawRewards(ctx context.Context, in *MsgWithdrawRewards, opts ...grpc.CallOption) (*MsgWithdrawRewardsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) WithdrawRewards(ctx context.Context, in *MsgWithdrawRewards, opts ...grpc.CallOption) (*MsgWithdrawRewardsResponse, error) {
	out := new(MsgWithdrawRewardsResponse)
	err := c.cc.Invoke(ctx, "/restake.v1beta1.Msg/WithdrawRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	WithdrawRewards(context.Context, *MsgWithdrawRewards) (*MsgWithdrawRewardsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) WithdrawRewards(ctx context.Context, req *MsgWithdrawRewards) (*MsgWithdrawRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawRewards not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_WithdrawRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawRewards)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restake.v1beta1.Msg/WithdrawRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawRewards(ctx, req.(*MsgWithdrawRewards))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "restake.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WithdrawRewards",
			Handler:    _Msg_WithdrawRewards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "restake/v1beta1/tx.proto",
}

func (m *MsgWithdrawRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgWithdrawRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawRewardsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgWithdrawRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawRewardsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
