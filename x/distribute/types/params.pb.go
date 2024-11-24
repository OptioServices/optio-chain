// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: optio/distribute/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// Params defines the parameters for the module.
type Params struct {
	MaxSupply          uint64   `protobuf:"varint,1,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty" yaml:"max_supply"`
	Denom              string   `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	AuthorizedAccounts []string `protobuf:"bytes,3,rep,name=authorized_accounts,json=authorizedAccounts,proto3" json:"authorized_accounts,omitempty" yaml:"authorized_accounts"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_15d225345473907a, []int{0}
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

func (m *Params) GetMaxSupply() uint64 {
	if m != nil {
		return m.MaxSupply
	}
	return 0
}

func (m *Params) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Params) GetAuthorizedAccounts() []string {
	if m != nil {
		return m.AuthorizedAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "optio.distribute.Params")
}

func init() { proto.RegisterFile("optio/distribute/params.proto", fileDescriptor_15d225345473907a) }

var fileDescriptor_15d225345473907a = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x2f, 0x28, 0xc9,
	0xcc, 0xd7, 0x4f, 0xc9, 0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0xd5, 0x2f, 0x48, 0x2c,
	0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x00, 0x4b, 0xeb, 0x21, 0xa4,
	0xa5, 0x04, 0x13, 0x73, 0x33, 0xf3, 0xf2, 0xf5, 0xc1, 0x24, 0x44, 0x91, 0x94, 0x48, 0x7a, 0x7e,
	0x7a, 0x3e, 0x98, 0xa9, 0x0f, 0x62, 0x41, 0x44, 0x95, 0x6e, 0x33, 0x72, 0xb1, 0x05, 0x80, 0xcd,
	0x12, 0x32, 0xe1, 0xe2, 0xca, 0x4d, 0xac, 0x88, 0x2f, 0x2e, 0x2d, 0x28, 0xc8, 0xa9, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x71, 0x12, 0xfd, 0x74, 0x4f, 0x5e, 0xb0, 0x32, 0x31, 0x37, 0xc7, 0x4a,
	0x09, 0x21, 0xa7, 0x14, 0xc4, 0x99, 0x9b, 0x58, 0x11, 0x0c, 0x66, 0x0b, 0xa9, 0x71, 0xb1, 0xa6,
	0xa4, 0xe6, 0xe5, 0xe7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x3a, 0x09, 0x7c, 0xba, 0x27, 0xcf,
	0x03, 0xd1, 0x00, 0x16, 0x56, 0x0a, 0x82, 0x48, 0x0b, 0xf9, 0x73, 0x09, 0x27, 0x96, 0x96, 0x64,
	0xe4, 0x17, 0x65, 0x56, 0xa5, 0xa6, 0xc4, 0x27, 0x26, 0x27, 0xe7, 0x97, 0xe6, 0x95, 0x14, 0x4b,
	0x30, 0x2b, 0x30, 0x6b, 0x70, 0x3a, 0xc9, 0x7d, 0xba, 0x27, 0x2f, 0x05, 0xd1, 0x85, 0x45, 0x91,
	0x52, 0x90, 0x10, 0x42, 0xd4, 0x11, 0x2a, 0x68, 0xa5, 0xf4, 0x62, 0x81, 0x3c, 0x63, 0xd7, 0xf3,
	0x0d, 0x5a, 0x92, 0x90, 0xc0, 0xa9, 0x40, 0x0e, 0x1e, 0x88, 0x97, 0x9c, 0xbc, 0x4f, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c,
	0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x30, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0xdf, 0x1f, 0xa4, 0x3f, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0xb5, 0x58, 0x1f,
	0x8b, 0x69, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0x10, 0x33, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x82, 0xef, 0x90, 0x7b, 0x8d, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.MaxSupply != that1.MaxSupply {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if len(this.AuthorizedAccounts) != len(that1.AuthorizedAccounts) {
		return false
	}
	for i := range this.AuthorizedAccounts {
		if this.AuthorizedAccounts[i] != that1.AuthorizedAccounts[i] {
			return false
		}
	}
	return true
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
	if len(m.AuthorizedAccounts) > 0 {
		for iNdEx := len(m.AuthorizedAccounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthorizedAccounts[iNdEx])
			copy(dAtA[i:], m.AuthorizedAccounts[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.AuthorizedAccounts[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.MaxSupply != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	if m.MaxSupply != 0 {
		n += 1 + sovParams(uint64(m.MaxSupply))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.AuthorizedAccounts) > 0 {
		for _, s := range m.AuthorizedAccounts {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedAccounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedAccounts = append(m.AuthorizedAccounts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
