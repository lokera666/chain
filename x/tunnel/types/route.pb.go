// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: band/tunnel/v1beta1/route.proto

package types

import (
	fmt "fmt"
	github_com_bandprotocol_chain_v3_x_bandtss_types "github.com/bandprotocol/chain/v3/x/bandtss/types"
	types "github.com/bandprotocol/chain/v3/x/feeds/types"
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

// TSSRoute represents a route for TSS packets and implements the RouteI interface.
type TSSRoute struct {
	// destination_chain_id is the destination chain ID
	DestinationChainID string `protobuf:"bytes,1,opt,name=destination_chain_id,json=destinationChainId,proto3" json:"destination_chain_id,omitempty"`
	// destination_contract_address is the destination contract address
	DestinationContractAddress string `protobuf:"bytes,2,opt,name=destination_contract_address,json=destinationContractAddress,proto3" json:"destination_contract_address,omitempty"`
}

func (m *TSSRoute) Reset()         { *m = TSSRoute{} }
func (m *TSSRoute) String() string { return proto.CompactTextString(m) }
func (*TSSRoute) ProtoMessage()    {}
func (*TSSRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_543238289d94b7a6, []int{0}
}
func (m *TSSRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TSSRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TSSRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TSSRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSSRoute.Merge(m, src)
}
func (m *TSSRoute) XXX_Size() int {
	return m.Size()
}
func (m *TSSRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_TSSRoute.DiscardUnknown(m)
}

var xxx_messageInfo_TSSRoute proto.InternalMessageInfo

func (m *TSSRoute) GetDestinationChainID() string {
	if m != nil {
		return m.DestinationChainID
	}
	return ""
}

func (m *TSSRoute) GetDestinationContractAddress() string {
	if m != nil {
		return m.DestinationContractAddress
	}
	return ""
}

// TSSPacketReceipt represents a receipt for a TSS packet and implements the PacketReceiptI interface.
type TSSPacketReceipt struct {
	// signing_id is the signing ID
	SigningID github_com_bandprotocol_chain_v3_x_bandtss_types.SigningID `protobuf:"varint,1,opt,name=signing_id,json=signingId,proto3,casttype=github.com/bandprotocol/chain/v3/x/bandtss/types.SigningID" json:"signing_id,omitempty"`
}

func (m *TSSPacketReceipt) Reset()         { *m = TSSPacketReceipt{} }
func (m *TSSPacketReceipt) String() string { return proto.CompactTextString(m) }
func (*TSSPacketReceipt) ProtoMessage()    {}
func (*TSSPacketReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_543238289d94b7a6, []int{1}
}
func (m *TSSPacketReceipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TSSPacketReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TSSPacketReceipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TSSPacketReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSSPacketReceipt.Merge(m, src)
}
func (m *TSSPacketReceipt) XXX_Size() int {
	return m.Size()
}
func (m *TSSPacketReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_TSSPacketReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_TSSPacketReceipt proto.InternalMessageInfo

func (m *TSSPacketReceipt) GetSigningID() github_com_bandprotocol_chain_v3_x_bandtss_types.SigningID {
	if m != nil {
		return m.SigningID
	}
	return 0
}

// IBCRoute is the type for an IBC route
type IBCRoute struct {
	// channel_id is the IBC channel ID
	ChannelID string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (m *IBCRoute) Reset()         { *m = IBCRoute{} }
func (m *IBCRoute) String() string { return proto.CompactTextString(m) }
func (*IBCRoute) ProtoMessage()    {}
func (*IBCRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_543238289d94b7a6, []int{2}
}
func (m *IBCRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCRoute.Merge(m, src)
}
func (m *IBCRoute) XXX_Size() int {
	return m.Size()
}
func (m *IBCRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCRoute.DiscardUnknown(m)
}

var xxx_messageInfo_IBCRoute proto.InternalMessageInfo

func (m *IBCRoute) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

// IBCPacketReceipt represents a receipt for a IBC packet and implements the PacketReceiptI interface.
type IBCPacketReceipt struct {
	// sequence is representing the sequence of the IBC packet.
	Sequence uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *IBCPacketReceipt) Reset()         { *m = IBCPacketReceipt{} }
func (m *IBCPacketReceipt) String() string { return proto.CompactTextString(m) }
func (*IBCPacketReceipt) ProtoMessage()    {}
func (*IBCPacketReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_543238289d94b7a6, []int{3}
}
func (m *IBCPacketReceipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCPacketReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCPacketReceipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCPacketReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCPacketReceipt.Merge(m, src)
}
func (m *IBCPacketReceipt) XXX_Size() int {
	return m.Size()
}
func (m *IBCPacketReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCPacketReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_IBCPacketReceipt proto.InternalMessageInfo

func (m *IBCPacketReceipt) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// IBCPacket is the type for an IBC packet
type IBCPacket struct {
	// tunnel_id is the tunnel ID
	TunnelID uint64 `protobuf:"varint,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	// sequence is representing the sequence of the tunnel packet.
	Sequence uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// prices is the list of prices information from feeds module.
	Prices []types.Price `protobuf:"bytes,3,rep,name=prices,proto3" json:"prices"`
	// created_at is the timestamp when the packet is created
	CreatedAt int64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *IBCPacket) Reset()         { *m = IBCPacket{} }
func (m *IBCPacket) String() string { return proto.CompactTextString(m) }
func (*IBCPacket) ProtoMessage()    {}
func (*IBCPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_543238289d94b7a6, []int{4}
}
func (m *IBCPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCPacket.Merge(m, src)
}
func (m *IBCPacket) XXX_Size() int {
	return m.Size()
}
func (m *IBCPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCPacket.DiscardUnknown(m)
}

var xxx_messageInfo_IBCPacket proto.InternalMessageInfo

func (m *IBCPacket) GetTunnelID() uint64 {
	if m != nil {
		return m.TunnelID
	}
	return 0
}

func (m *IBCPacket) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *IBCPacket) GetPrices() []types.Price {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *IBCPacket) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*TSSRoute)(nil), "band.tunnel.v1beta1.TSSRoute")
	proto.RegisterType((*TSSPacketReceipt)(nil), "band.tunnel.v1beta1.TSSPacketReceipt")
	proto.RegisterType((*IBCRoute)(nil), "band.tunnel.v1beta1.IBCRoute")
	proto.RegisterType((*IBCPacketReceipt)(nil), "band.tunnel.v1beta1.IBCPacketReceipt")
	proto.RegisterType((*IBCPacket)(nil), "band.tunnel.v1beta1.IBCPacket")
}

func init() { proto.RegisterFile("band/tunnel/v1beta1/route.proto", fileDescriptor_543238289d94b7a6) }

var fileDescriptor_543238289d94b7a6 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xbd, 0x8e, 0x13, 0x31,
	0x10, 0xc7, 0xb3, 0x24, 0x44, 0x59, 0xf3, 0xa1, 0xd3, 0x72, 0x42, 0xb9, 0x08, 0x76, 0xa3, 0x54,
	0x41, 0x22, 0x6b, 0x1d, 0x57, 0x20, 0xa5, 0xe2, 0x36, 0x29, 0xb0, 0x10, 0xd2, 0x69, 0x93, 0x8a,
	0x26, 0x72, 0x6c, 0x93, 0x2c, 0xdc, 0xd9, 0x61, 0xed, 0x9c, 0xe0, 0x2d, 0x10, 0x8f, 0x40, 0x85,
	0xa8, 0xf3, 0x10, 0xa7, 0x54, 0x57, 0x52, 0xad, 0xd0, 0xe6, 0x2d, 0xa8, 0x90, 0x3f, 0x48, 0x2e,
	0x14, 0x88, 0x6e, 0xfd, 0x9f, 0x9f, 0x67, 0xc6, 0xff, 0x99, 0x05, 0xd1, 0x14, 0x73, 0x0a, 0xd5,
	0x92, 0x73, 0x76, 0x0e, 0x2f, 0x8f, 0xa7, 0x4c, 0xe1, 0x63, 0x98, 0x8b, 0xa5, 0x62, 0xf1, 0x22,
	0x17, 0x4a, 0x04, 0x0f, 0x34, 0x10, 0x5b, 0x20, 0x76, 0x40, 0xeb, 0x88, 0x08, 0x79, 0x21, 0xe4,
	0xc4, 0x20, 0xd0, 0x1e, 0x2c, 0xdf, 0x3a, 0x9c, 0x89, 0x99, 0xb0, 0xba, 0xfe, 0x72, 0x6a, 0x68,
	0xca, 0xbc, 0x65, 0x8c, 0xca, 0x6d, 0x15, 0x73, 0xb2, 0xf1, 0xce, 0x57, 0x0f, 0x34, 0xc6, 0xa3,
	0x51, 0xaa, 0x0b, 0x07, 0x2f, 0xc1, 0x21, 0x65, 0x52, 0x65, 0x1c, 0xab, 0x4c, 0xf0, 0x09, 0x99,
	0xe3, 0x8c, 0x4f, 0x32, 0xda, 0xf4, 0xda, 0x5e, 0xd7, 0x4f, 0x1e, 0x96, 0x45, 0x14, 0x0c, 0x77,
	0xf1, 0x81, 0x0e, 0xa3, 0x61, 0x1a, 0xd0, 0xbf, 0x35, 0x1a, 0xbc, 0x00, 0x8f, 0xf6, 0x32, 0x09,
	0xae, 0x72, 0x4c, 0xd4, 0x04, 0x53, 0x9a, 0x33, 0x29, 0x9b, 0xb7, 0x74, 0xc6, 0xb4, 0x75, 0xf3,
	0xa6, 0x43, 0x4e, 0x2d, 0xd1, 0x07, 0xeb, 0x55, 0xaf, 0x6e, 0xda, 0x42, 0x9d, 0x2f, 0x1e, 0x38,
	0x18, 0x8f, 0x46, 0x67, 0x98, 0xbc, 0x67, 0x2a, 0x65, 0x84, 0x65, 0x0b, 0x15, 0xbc, 0x03, 0x40,
	0x66, 0x33, 0x9e, 0xf1, 0xd9, 0x9f, 0x16, 0x6b, 0xc9, 0xab, 0xb2, 0x88, 0xfc, 0x91, 0x55, 0xd1,
	0xf0, 0x57, 0x11, 0xf5, 0x67, 0x99, 0x9a, 0x2f, 0xa7, 0x31, 0x11, 0x17, 0x50, 0x3b, 0x61, 0x1e,
	0x4d, 0xc4, 0x39, 0x34, 0x6f, 0x83, 0x97, 0x27, 0xf0, 0xa3, 0xd1, 0x95, 0x94, 0x50, 0x7d, 0x5a,
	0x30, 0x19, 0x6f, 0x6f, 0xa7, 0xbe, 0x4b, 0x8f, 0x68, 0x3f, 0x58, 0xaf, 0x7a, 0xf7, 0xf7, 0xca,
	0xa3, 0xce, 0x00, 0x34, 0x50, 0x32, 0xb0, 0xc6, 0x3d, 0x05, 0x80, 0xcc, 0xb1, 0x9e, 0xd4, 0xce,
	0xae, 0x7b, 0xba, 0x97, 0x81, 0x55, 0x75, 0x36, 0x07, 0x20, 0xda, 0xf7, 0xd7, 0xab, 0xde, 0x6d,
	0x73, 0xb1, 0x93, 0x80, 0x03, 0x94, 0x0c, 0xf6, 0x1f, 0xd6, 0x02, 0x0d, 0xc9, 0x3e, 0x2c, 0x19,
	0x27, 0xcc, 0xf8, 0x54, 0x4b, 0xb7, 0xe7, 0x9b, 0x8d, 0x68, 0xc3, 0x18, 0x57, 0xa8, 0xf3, 0xdd,
	0x03, 0xfe, 0x36, 0x49, 0xf0, 0x04, 0xf8, 0x76, 0x67, 0x76, 0xae, 0xdc, 0x2d, 0x8b, 0xa8, 0x31,
	0x5e, 0xba, 0x46, 0x1a, 0x36, 0x8c, 0xe8, 0xbf, 0x0a, 0x05, 0xcf, 0x41, 0x7d, 0x91, 0x67, 0x84,
	0xc9, 0x66, 0xb5, 0x5d, 0xed, 0xde, 0x79, 0x76, 0x14, 0x9b, 0x75, 0xb4, 0xab, 0xe3, 0x16, 0x29,
	0x3e, 0xd3, 0x44, 0x52, 0xbb, 0x2a, 0xa2, 0x4a, 0xea, 0xf0, 0xe0, 0x31, 0x00, 0x24, 0x67, 0x58,
	0x31, 0x3a, 0xc1, 0xaa, 0x59, 0x6b, 0x7b, 0xdd, 0x6a, 0xea, 0x3b, 0xe5, 0x54, 0x25, 0xaf, 0xbf,
	0x95, 0xa1, 0x77, 0x55, 0x86, 0xde, 0x75, 0x19, 0x7a, 0x3f, 0xcb, 0xd0, 0xfb, 0xbc, 0x09, 0x2b,
	0xd7, 0x9b, 0xb0, 0xf2, 0x63, 0x13, 0x56, 0xde, 0xc0, 0xff, 0x18, 0x97, 0xfb, 0x65, 0xcc, 0xb4,
	0xa6, 0x75, 0x43, 0x9c, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x00, 0xf7, 0x98, 0x0d, 0x4e, 0x03,
	0x00, 0x00,
}

func (this *TSSRoute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TSSRoute)
	if !ok {
		that2, ok := that.(TSSRoute)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.DestinationChainID != that1.DestinationChainID {
		return false
	}
	if this.DestinationContractAddress != that1.DestinationContractAddress {
		return false
	}
	return true
}
func (this *TSSPacketReceipt) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TSSPacketReceipt)
	if !ok {
		that2, ok := that.(TSSPacketReceipt)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SigningID != that1.SigningID {
		return false
	}
	return true
}
func (this *IBCRoute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IBCRoute)
	if !ok {
		that2, ok := that.(IBCRoute)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ChannelID != that1.ChannelID {
		return false
	}
	return true
}
func (this *IBCPacketReceipt) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IBCPacketReceipt)
	if !ok {
		that2, ok := that.(IBCPacketReceipt)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Sequence != that1.Sequence {
		return false
	}
	return true
}
func (this *IBCPacket) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IBCPacket)
	if !ok {
		that2, ok := that.(IBCPacket)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TunnelID != that1.TunnelID {
		return false
	}
	if this.Sequence != that1.Sequence {
		return false
	}
	if len(this.Prices) != len(that1.Prices) {
		return false
	}
	for i := range this.Prices {
		if !this.Prices[i].Equal(&that1.Prices[i]) {
			return false
		}
	}
	if this.CreatedAt != that1.CreatedAt {
		return false
	}
	return true
}
func (m *TSSRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TSSRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TSSRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DestinationContractAddress) > 0 {
		i -= len(m.DestinationContractAddress)
		copy(dAtA[i:], m.DestinationContractAddress)
		i = encodeVarintRoute(dAtA, i, uint64(len(m.DestinationContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DestinationChainID) > 0 {
		i -= len(m.DestinationChainID)
		copy(dAtA[i:], m.DestinationChainID)
		i = encodeVarintRoute(dAtA, i, uint64(len(m.DestinationChainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TSSPacketReceipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TSSPacketReceipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TSSPacketReceipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SigningID != 0 {
		i = encodeVarintRoute(dAtA, i, uint64(m.SigningID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *IBCRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintRoute(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IBCPacketReceipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCPacketReceipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCPacketReceipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintRoute(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func (m *IBCPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintRoute(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Prices) > 0 {
		for iNdEx := len(m.Prices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRoute(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Sequence != 0 {
		i = encodeVarintRoute(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x10
	}
	if m.TunnelID != 0 {
		i = encodeVarintRoute(dAtA, i, uint64(m.TunnelID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRoute(dAtA []byte, offset int, v uint64) int {
	offset -= sovRoute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TSSRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DestinationChainID)
	if l > 0 {
		n += 1 + l + sovRoute(uint64(l))
	}
	l = len(m.DestinationContractAddress)
	if l > 0 {
		n += 1 + l + sovRoute(uint64(l))
	}
	return n
}

func (m *TSSPacketReceipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SigningID != 0 {
		n += 1 + sovRoute(uint64(m.SigningID))
	}
	return n
}

func (m *IBCRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovRoute(uint64(l))
	}
	return n
}

func (m *IBCPacketReceipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovRoute(uint64(m.Sequence))
	}
	return n
}

func (m *IBCPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TunnelID != 0 {
		n += 1 + sovRoute(uint64(m.TunnelID))
	}
	if m.Sequence != 0 {
		n += 1 + sovRoute(uint64(m.Sequence))
	}
	if len(m.Prices) > 0 {
		for _, e := range m.Prices {
			l = e.Size()
			n += 1 + l + sovRoute(uint64(l))
		}
	}
	if m.CreatedAt != 0 {
		n += 1 + sovRoute(uint64(m.CreatedAt))
	}
	return n
}

func sovRoute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRoute(x uint64) (n int) {
	return sovRoute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TSSRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
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
			return fmt.Errorf("proto: TSSRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TSSRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoute
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
func (m *TSSPacketReceipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
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
			return fmt.Errorf("proto: TSSPacketReceipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TSSPacketReceipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningID", wireType)
			}
			m.SigningID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigningID |= github_com_bandprotocol_chain_v3_x_bandtss_types.SigningID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoute
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
func (m *IBCRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
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
			return fmt.Errorf("proto: IBCRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoute
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
func (m *IBCPacketReceipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
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
			return fmt.Errorf("proto: IBCPacketReceipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCPacketReceipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoute
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
func (m *IBCPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoute
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
			return fmt.Errorf("proto: IBCPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TunnelID", wireType)
			}
			m.TunnelID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TunnelID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
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
				return ErrInvalidLengthRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prices = append(m.Prices, types.Price{})
			if err := m.Prices[len(m.Prices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoute
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
func skipRoute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoute
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
					return 0, ErrIntOverflowRoute
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
					return 0, ErrIntOverflowRoute
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
				return 0, ErrInvalidLengthRoute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRoute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRoute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRoute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRoute = fmt.Errorf("proto: unexpected end of group")
)
