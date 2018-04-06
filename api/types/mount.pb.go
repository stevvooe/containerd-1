// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/api/types/mount.proto

package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Mount describes mounts for a container.
//
// This type is the lingua franca of ContainerD. All services provide mounts
// to be used with the container at creation time.
//
// The Mount type follows the structure of the mount syscall, including a type,
// source, target and options.
type Mount struct {
	// Type defines the nature of the mount.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Source specifies the name of the mount. Depending on mount type, this
	// may be a volume name or a host path, or even ignored.
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// Target path in container
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	// Options specifies zero or more fstab style mount options.
	Options              []string `protobuf:"bytes,4,rep,name=options" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `protobuf_unrecognized:"proto3" json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mount) Reset()                    { *m = Mount{} }
func (*Mount) ProtoMessage()               {}
func (*Mount) Descriptor() ([]byte, []int) { return fileDescriptorMount, []int{0} }
func (dst *Mount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mount.Merge(dst, src)
}
func (m *Mount) XXX_Size() int {
	return xxx_messageInfo_Mount.Size(m)
}
func (m *Mount) XXX_DiscardUnknown() {
	xxx_messageInfo_Mount.DiscardUnknown(m)
}

var xxx_messageInfo_Mount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Mount)(nil), "containerd.types.Mount")
}
func (m *Mount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mount) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMount(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Source) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMount(dAtA, i, uint64(len(m.Source)))
		i += copy(dAtA[i:], m.Source)
	}
	if len(m.Target) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMount(dAtA, i, uint64(len(m.Target)))
		i += copy(dAtA[i:], m.Target)
	}
	if len(m.Options) > 0 {
		for _, s := range m.Options {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		if proto.Proto3UnknownFields {
			i += copy(dAtA[i:], m.XXX_unrecognized)
		}
	}
	return i, nil
}

func encodeVarintMount(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Mount) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMount(uint64(l))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovMount(uint64(l))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovMount(uint64(l))
	}
	if len(m.Options) > 0 {
		for _, s := range m.Options {
			l = len(s)
			n += 1 + l + sovMount(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		if proto.Proto3UnknownFields {
			n += len(m.XXX_unrecognized)
		}
	}
	return n
}

func sovMount(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMount(x uint64) (n int) {
	return sovMount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Mount) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Mount{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Source:` + fmt.Sprintf("%v", this.Source) + `,`,
		`Target:` + fmt.Sprintf("%v", this.Target) + `,`,
		`Options:` + fmt.Sprintf("%v", this.Options) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMount(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Mount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Mount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = append(m.Options, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMount
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
					return 0, ErrIntOverflowMount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
				if shift >= 64 {
					return 0, ErrIntOverflowMount
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMount
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMount
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipMount(dAtA[start:])
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
	ErrInvalidLengthMount = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMount   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/types/mount.proto", fileDescriptorMount)
}

var fileDescriptorMount = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d,
	0x4a, 0x41, 0x66, 0x26, 0x16, 0x64, 0xea, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0xe7, 0xe6, 0x97,
	0xe6, 0x95, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x20, 0x54, 0xe8, 0x81, 0x65, 0xa5,
	0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x92, 0xfa, 0x20, 0x16, 0x44, 0x9d, 0x52, 0x2a, 0x17, 0xab,
	0x2f, 0x48, 0x9b, 0x90, 0x10, 0x17, 0x0b, 0x48, 0x9d, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x98, 0x2d, 0x24, 0xc6, 0xc5, 0x56, 0x9c, 0x5f, 0x5a, 0x94, 0x9c, 0x2a, 0xc1, 0x04, 0x16, 0x85,
	0xf2, 0x40, 0xe2, 0x25, 0x89, 0x45, 0xe9, 0xa9, 0x25, 0x12, 0xcc, 0x10, 0x71, 0x08, 0x4f, 0x48,
	0x82, 0x8b, 0x3d, 0xbf, 0xa0, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0x82, 0x45, 0x81, 0x59, 0x83, 0x33,
	0x08, 0xc6, 0x75, 0xf2, 0x3a, 0xf1, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x86, 0x47, 0x72,
	0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x94, 0x01,
	0xf1, 0x1e, 0xb4, 0x06, 0x93, 0x11, 0x0c, 0x49, 0x6c, 0x60, 0xb7, 0x1b, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x82, 0x1c, 0x02, 0x18, 0x1d, 0x01, 0x00, 0x00,
}
