// Code generated by protoc-gen-gogo.
// source: node.proto
// DO NOT EDIT!

package vcrypt

import proto "github.com/gogo/protobuf/proto"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"
import cryptex9 "github.com/vcrypt/vcrypt/cryptex"
import secret3 "github.com/vcrypt/vcrypt/secret"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Node struct {
	Nonce   []byte             `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Inputs  [][]byte           `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty"`
	cryptex *cryptex9.Envelope `protobuf:"bytes,3,opt,name=cryptex" json:"cryptex,omitempty"`
	secret  *secret3.Envelope  `protobuf:"bytes,4,opt,name=secret" json:"secret,omitempty"`
	Marker  *Marker            `protobuf:"bytes,5,opt,name=marker" json:"marker,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}

func (m *Node) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Node) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Nonce != nil {
		if len(m.Nonce) > 0 {
			data[i] = 0xa
			i++
			i = encodeVarintNode(data, i, uint64(len(m.Nonce)))
			i += copy(data[i:], m.Nonce)
		}
	}
	if len(m.Inputs) > 0 {
		for _, b := range m.Inputs {
			data[i] = 0x12
			i++
			i = encodeVarintNode(data, i, uint64(len(b)))
			i += copy(data[i:], b)
		}
	}
	if m.cryptex != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintNode(data, i, uint64(m.cryptex.Size()))
		n1, err := m.cryptex.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.secret != nil {
		data[i] = 0x22
		i++
		i = encodeVarintNode(data, i, uint64(m.secret.Size()))
		n2, err := m.secret.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Marker != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintNode(data, i, uint64(m.Marker.Size()))
		n3, err := m.Marker.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeFixed64Node(data []byte, offset int, v uint64) int {
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
func encodeFixed32Node(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintNode(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Node) Size() (n int) {
	var l int
	_ = l
	if m.Nonce != nil {
		l = len(m.Nonce)
		if l > 0 {
			n += 1 + l + sovNode(uint64(l))
		}
	}
	if len(m.Inputs) > 0 {
		for _, b := range m.Inputs {
			l = len(b)
			n += 1 + l + sovNode(uint64(l))
		}
	}
	if m.cryptex != nil {
		l = m.cryptex.Size()
		n += 1 + l + sovNode(uint64(l))
	}
	if m.secret != nil {
		l = m.secret.Size()
		n += 1 + l + sovNode(uint64(l))
	}
	if m.Marker != nil {
		l = m.Marker.Size()
		n += 1 + l + sovNode(uint64(l))
	}
	return n
}

func sovNode(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Node) GetValue() interface{} {
	if this.Nonce != nil {
		return this.Nonce
	}
	if this.Inputs != nil {
		return this.Inputs
	}
	if this.cryptex != nil {
		return this.cryptex
	}
	if this.secret != nil {
		return this.secret
	}
	if this.Marker != nil {
		return this.Marker
	}
	return nil
}

func (this *Node) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case []byte:
		this.Nonce = vt
	case [][]byte:
		this.Inputs = vt
	case *cryptex9.Envelope:
		this.cryptex = vt
	case *secret3.Envelope:
		this.secret = vt
	case *Marker:
		this.Marker = vt
	default:
		this.cryptex = new(cryptex9.Envelope)
		if set := this.cryptex.SetValue(value); set {
			return true
		}
		this.cryptex = nil
		this.secret = new(secret3.Envelope)
		if set := this.secret.SetValue(value); set {
			return true
		}
		this.secret = nil
		return false
	}
	return true
}
func (m *Node) Unmarshal(data []byte) error {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inputs = append(m.Inputs, make([]byte, postIndex-iNdEx))
			copy(m.Inputs[len(m.Inputs)-1], data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field cryptex", wireType)
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
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.cryptex == nil {
				m.cryptex = &cryptex9.Envelope{}
			}
			if err := m.cryptex.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field secret", wireType)
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
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.secret == nil {
				m.secret = &secret3.Envelope{}
			}
			if err := m.secret.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Marker", wireType)
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
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Marker == nil {
				m.Marker = &Marker{}
			}
			if err := m.Marker.Unmarshal(data[iNdEx:postIndex]); err != nil {
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
			skippy, err := skipNode(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipNode(data []byte) (n int, err error) {
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
				return 0, ErrInvalidLengthNode
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
				next, err := skipNode(data[start:])
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
	ErrInvalidLengthNode = fmt.Errorf("proto: negative length found during unmarshaling")
)
