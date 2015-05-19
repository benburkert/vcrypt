// Code generated by protoc-gen-gogo.
// source: cryptex/cryptex.proto
// DO NOT EDIT!

/*
	Package cryptex is a generated protocol buffer package.

	It is generated from these files:
		cryptex/cryptex.proto
		cryptex/sss.proto
		cryptex/xor.proto
		cryptex/secretbox.proto

	It has these top-level messages:
		Envelope
*/
package cryptex

import proto "github.com/gogo/protobuf/proto"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Envelope struct {
	SSS       *SSS       `protobuf:"bytes,1,opt,name=sss" json:"sss,omitempty"`
	XOR       *XOR       `protobuf:"bytes,2,opt,name=xor" json:"xor,omitempty"`
	SecretBox *SecretBox `protobuf:"bytes,3,opt,name=secretbox" json:"secretbox,omitempty"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}

func (m *Envelope) GetSSS() *SSS {
	if m != nil {
		return m.SSS
	}
	return nil
}

func (m *Envelope) GetXOR() *XOR {
	if m != nil {
		return m.XOR
	}
	return nil
}

func (m *Envelope) GetSecretBox() *SecretBox {
	if m != nil {
		return m.SecretBox
	}
	return nil
}

func (m *Envelope) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Envelope) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SSS != nil {
		data[i] = 0xa
		i++
		i = encodeVarintCryptex(data, i, uint64(m.SSS.Size()))
		n1, err := m.SSS.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XOR != nil {
		data[i] = 0x12
		i++
		i = encodeVarintCryptex(data, i, uint64(m.XOR.Size()))
		n2, err := m.XOR.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.SecretBox != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintCryptex(data, i, uint64(m.SecretBox.Size()))
		n3, err := m.SecretBox.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeFixed64Cryptex(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Cryptex(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCryptex(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Envelope) Size() (n int) {
	var l int
	_ = l
	if m.SSS != nil {
		l = m.SSS.Size()
		n += 1 + l + sovCryptex(uint64(l))
	}
	if m.XOR != nil {
		l = m.XOR.Size()
		n += 1 + l + sovCryptex(uint64(l))
	}
	if m.SecretBox != nil {
		l = m.SecretBox.Size()
		n += 1 + l + sovCryptex(uint64(l))
	}
	return n
}

func sovCryptex(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCryptex(x uint64) (n int) {
	return sovCryptex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Envelope) GetValue() interface{} {
	if this.SSS != nil {
		return this.SSS
	}
	if this.XOR != nil {
		return this.XOR
	}
	if this.SecretBox != nil {
		return this.SecretBox
	}
	return nil
}

func (this *Envelope) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *SSS:
		this.SSS = vt
	case *XOR:
		this.XOR = vt
	case *SecretBox:
		this.SecretBox = vt
	default:
		return false
	}
	return true
}
func (m *Envelope) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SSS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthCryptex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SSS == nil {
				m.SSS = &SSS{}
			}
			if err := m.SSS.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XOR", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthCryptex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.XOR == nil {
				m.XOR = &XOR{}
			}
			if err := m.XOR.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecretBox", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthCryptex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SecretBox == nil {
				m.SecretBox = &SecretBox{}
			}
			if err := m.SecretBox.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipCryptex(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCryptex
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipCryptex(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCryptex
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCryptex(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCryptex = fmt.Errorf("proto: negative length found during unmarshaling")
)
