// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: craft/exp/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// Params holds parameters for the incentives module
type Params struct {
	MaxCoinMint uint64 `protobuf:"varint,1,opt,name=max_coin_mint,json=maxCoinMint,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_coin_mint,omitempty" yaml:"max_coin_mint"`
	DaoAccount  string `protobuf:"bytes,2,opt,name=DaoAccount,proto3" json:"DaoAccount,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eb7710cccfbf9b, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxCoinMint() uint64 {
	if m != nil {
		return m.MaxCoinMint
	}
	return 0
}

func (m *Params) GetDaoAccount() string {
	if m != nil {
		return m.DaoAccount
	}
	return ""
}

// GenesisState defines the exp module's genesis state.
type GenesisState struct {
	Params    Params           `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	WhiteList []*AccountRecord `protobuf:"bytes,2,rep,name=white_list,json=whiteList,proto3" json:"white_list,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9eb7710cccfbf9b, []int{1}
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

func (m *GenesisState) GetWhiteList() []*AccountRecord {
	if m != nil {
		return m.WhiteList
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "craft.exp.v1beta1.Params")
	proto.RegisterType((*GenesisState)(nil), "craft.exp.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("craft/exp/v1beta1/genesis.proto", fileDescriptor_d9eb7710cccfbf9b) }

var fileDescriptor_d9eb7710cccfbf9b = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xbf, 0x4e, 0xfb, 0x30,
	0x18, 0x8c, 0x7f, 0xbf, 0xaa, 0x52, 0x1d, 0x18, 0x88, 0x3a, 0xb4, 0x1d, 0x9c, 0x2a, 0x12, 0x52,
	0x16, 0x12, 0xb5, 0x0c, 0x48, 0x2c, 0x88, 0x82, 0x04, 0x03, 0x48, 0x28, 0x6c, 0x2c, 0x91, 0x93,
	0xba, 0xa9, 0x45, 0x63, 0x47, 0xb1, 0x0b, 0xe9, 0x1b, 0x30, 0x32, 0xf1, 0x10, 0x3c, 0x49, 0xc7,
	0x8e, 0x4c, 0x05, 0xb5, 0x6f, 0xc0, 0x13, 0xa0, 0x38, 0x56, 0xc5, 0x9f, 0x4e, 0xf9, 0x2e, 0x77,
	0x9f, 0xef, 0xbe, 0x83, 0x76, 0x9c, 0xe3, 0x91, 0xf4, 0x49, 0x91, 0xf9, 0x0f, 0xbd, 0x88, 0x48,
	0xdc, 0xf3, 0x13, 0xc2, 0x88, 0xa0, 0xc2, 0xcb, 0x72, 0x2e, 0xb9, 0xb5, 0xa7, 0x04, 0x1e, 0x29,
	0x32, 0x4f, 0x0b, 0x3a, 0xcd, 0x84, 0x27, 0x5c, 0xb1, 0x7e, 0x39, 0x55, 0xc2, 0x4e, 0x3b, 0xe6,
	0x22, 0xe5, 0x22, 0xac, 0x88, 0x0a, 0x68, 0x0a, 0x55, 0xc8, 0x8f, 0xb0, 0x20, 0x1b, 0x9b, 0x98,
	0x53, 0xf6, 0x8b, 0xc7, 0x53, 0x39, 0xde, 0xf0, 0x25, 0xd0, 0xfc, 0x96, 0x90, 0x43, 0xcc, 0x29,
	0x1b, 0x69, 0x6f, 0xe7, 0x05, 0xc0, 0xfa, 0x0d, 0xce, 0x71, 0x2a, 0xac, 0x09, 0xdc, 0x4d, 0x71,
	0x11, 0x96, 0xaf, 0x87, 0x29, 0x65, 0xb2, 0x05, 0xba, 0xc0, 0xad, 0x0d, 0x2e, 0xe7, 0x4b, 0x1b,
	0x7c, 0x2e, 0xed, 0xe6, 0x0c, 0xa7, 0x93, 0x63, 0xe7, 0x87, 0xc4, 0x79, 0x7d, 0xb7, 0xdd, 0x84,
	0xca, 0xf1, 0x34, 0xf2, 0x62, 0x9e, 0xea, 0xec, 0xfa, 0x73, 0x20, 0x86, 0xf7, 0xbe, 0x9c, 0x65,
	0x44, 0x78, 0x67, 0x9c, 0x32, 0x11, 0x98, 0x29, 0x2e, 0xca, 0xe9, 0x9a, 0x32, 0x69, 0x21, 0x08,
	0xcf, 0x31, 0x3f, 0x8d, 0x63, 0x3e, 0x65, 0xb2, 0xf5, 0xaf, 0x0b, 0xdc, 0x46, 0xf0, 0xed, 0x8f,
	0xf3, 0x04, 0xe0, 0xce, 0x45, 0xd5, 0xe7, 0xad, 0xc4, 0x92, 0x58, 0x47, 0xb0, 0x9e, 0xa9, 0xa0,
	0x2a, 0x97, 0xd9, 0x6f, 0x7b, 0x7f, 0xfa, 0xf5, 0xaa, 0x4b, 0x06, 0xb5, 0xf9, 0xd2, 0x36, 0x02,
	0x2d, 0xb7, 0x4e, 0x20, 0x7c, 0x1c, 0x53, 0x49, 0xc2, 0x09, 0x15, 0xa5, 0xd3, 0x7f, 0xd7, 0xec,
	0x77, 0xb7, 0x2c, 0x6b, 0xe7, 0x80, 0xc4, 0x3c, 0x1f, 0x06, 0x0d, 0xb5, 0x73, 0x45, 0x85, 0x1c,
	0xec, 0xcf, 0x57, 0x08, 0x2c, 0x56, 0x08, 0x7c, 0xac, 0x10, 0x78, 0x5e, 0x23, 0x63, 0xb1, 0x46,
	0xc6, 0xdb, 0x1a, 0x19, 0x77, 0x66, 0xa1, 0xaa, 0x55, 0x47, 0x46, 0x75, 0xd5, 0xe8, 0xe1, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x2f, 0x75, 0xd2, 0x19, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DaoAccount) > 0 {
		i -= len(m.DaoAccount)
		copy(dAtA[i:], m.DaoAccount)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DaoAccount)))
		i--
		dAtA[i] = 0x12
	}
	if m.MaxCoinMint != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxCoinMint))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
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
	if len(m.WhiteList) > 0 {
		for iNdEx := len(m.WhiteList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WhiteList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxCoinMint != 0 {
		n += 1 + sovGenesis(uint64(m.MaxCoinMint))
	}
	l = len(m.DaoAccount)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.WhiteList) > 0 {
		for _, e := range m.WhiteList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCoinMint", wireType)
			}
			m.MaxCoinMint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCoinMint |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaoAccount", wireType)
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
			m.DaoAccount = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field WhiteList", wireType)
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
			m.WhiteList = append(m.WhiteList, &AccountRecord{})
			if err := m.WhiteList[len(m.WhiteList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
