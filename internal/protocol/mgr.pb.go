// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mgr.proto

package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CMD_MGR_ENUM int32

const (
	CMD_MGR_INVALID         CMD_MGR_ENUM = 0
	CMD_MGR_PING            CMD_MGR_ENUM = 1
	CMD_MGR_REGISTER_SERVER CMD_MGR_ENUM = 2
	CMD_MGR_LOSE_SERVER     CMD_MGR_ENUM = 3
)

var CMD_MGR_ENUM_name = map[int32]string{
	0: "INVALID",
	1: "PING",
	2: "REGISTER_SERVER",
	3: "LOSE_SERVER",
}
var CMD_MGR_ENUM_value = map[string]int32{
	"INVALID":         0,
	"PING":            1,
	"REGISTER_SERVER": 2,
	"LOSE_SERVER":     3,
}

func (x CMD_MGR_ENUM) String() string {
	return proto.EnumName(CMD_MGR_ENUM_name, int32(x))
}
func (CMD_MGR_ENUM) EnumDescriptor() ([]byte, []int) { return fileDescriptorMgr, []int{0, 0} }

type CMD_MGR struct {
}

func (m *CMD_MGR) Reset()                    { *m = CMD_MGR{} }
func (m *CMD_MGR) String() string            { return proto.CompactTextString(m) }
func (*CMD_MGR) ProtoMessage()               {}
func (*CMD_MGR) Descriptor() ([]byte, []int) { return fileDescriptorMgr, []int{0} }

// 服务器信息
type SERVER_INFO struct {
	Id       *SERVER_ID `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Type     uint32     `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Addrs    []string   `protobuf:"bytes,3,rep,name=Addrs" json:"Addrs,omitempty"`
	Ports    []int32    `protobuf:"varint,4,rep,packed,name=Ports" json:"Ports,omitempty"`
	Overload []uint32   `protobuf:"varint,5,rep,packed,name=Overload" json:"Overload,omitempty"`
	Version  string     `protobuf:"bytes,6,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (m *SERVER_INFO) Reset()                    { *m = SERVER_INFO{} }
func (m *SERVER_INFO) String() string            { return proto.CompactTextString(m) }
func (*SERVER_INFO) ProtoMessage()               {}
func (*SERVER_INFO) Descriptor() ([]byte, []int) { return fileDescriptorMgr, []int{1} }

func (m *SERVER_INFO) GetId() *SERVER_ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SERVER_INFO) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SERVER_INFO) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func (m *SERVER_INFO) GetPorts() []int32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *SERVER_INFO) GetOverload() []uint32 {
	if m != nil {
		return m.Overload
	}
	return nil
}

func (m *SERVER_INFO) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// PING ( MGR -> S )
type MSG_MGR_PING struct {
}

func (m *MSG_MGR_PING) Reset()                    { *m = MSG_MGR_PING{} }
func (m *MSG_MGR_PING) String() string            { return proto.CompactTextString(m) }
func (*MSG_MGR_PING) ProtoMessage()               {}
func (*MSG_MGR_PING) Descriptor() ([]byte, []int) { return fileDescriptorMgr, []int{2} }

// 注册服务器信息( S -> MGR -> S )
type MSG_MGR_REGISTER_SERVER struct {
	Data             *SERVER_INFO `protobuf:"bytes,1,opt,name=Data" json:"Data,omitempty"`
	Token            string       `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	TargetServerType uint32       `protobuf:"varint,3,opt,name=TargetServerType,proto3" json:"TargetServerType,omitempty"`
	TargetServerID   *SERVER_ID   `protobuf:"bytes,4,opt,name=TargetServerID" json:"TargetServerID,omitempty"`
}

func (m *MSG_MGR_REGISTER_SERVER) Reset()                    { *m = MSG_MGR_REGISTER_SERVER{} }
func (m *MSG_MGR_REGISTER_SERVER) String() string            { return proto.CompactTextString(m) }
func (*MSG_MGR_REGISTER_SERVER) ProtoMessage()               {}
func (*MSG_MGR_REGISTER_SERVER) Descriptor() ([]byte, []int) { return fileDescriptorMgr, []int{3} }

func (m *MSG_MGR_REGISTER_SERVER) GetData() *SERVER_INFO {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MSG_MGR_REGISTER_SERVER) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *MSG_MGR_REGISTER_SERVER) GetTargetServerType() uint32 {
	if m != nil {
		return m.TargetServerType
	}
	return 0
}

func (m *MSG_MGR_REGISTER_SERVER) GetTargetServerID() *SERVER_ID {
	if m != nil {
		return m.TargetServerID
	}
	return nil
}

// 丢失服务器信息( MGR -> S )
type MSG_MGR_LOSE_SERVER struct {
	Id   *SERVER_ID `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Type uint32     `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *MSG_MGR_LOSE_SERVER) Reset()                    { *m = MSG_MGR_LOSE_SERVER{} }
func (m *MSG_MGR_LOSE_SERVER) String() string            { return proto.CompactTextString(m) }
func (*MSG_MGR_LOSE_SERVER) ProtoMessage()               {}
func (*MSG_MGR_LOSE_SERVER) Descriptor() ([]byte, []int) { return fileDescriptorMgr, []int{4} }

func (m *MSG_MGR_LOSE_SERVER) GetId() *SERVER_ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MSG_MGR_LOSE_SERVER) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*CMD_MGR)(nil), "protocol.CMD_MGR")
	proto.RegisterType((*SERVER_INFO)(nil), "protocol.SERVER_INFO")
	proto.RegisterType((*MSG_MGR_PING)(nil), "protocol.MSG_MGR_PING")
	proto.RegisterType((*MSG_MGR_REGISTER_SERVER)(nil), "protocol.MSG_MGR_REGISTER_SERVER")
	proto.RegisterType((*MSG_MGR_LOSE_SERVER)(nil), "protocol.MSG_MGR_LOSE_SERVER")
	proto.RegisterEnum("protocol.CMD_MGR_ENUM", CMD_MGR_ENUM_name, CMD_MGR_ENUM_value)
}
func (m *CMD_MGR) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CMD_MGR) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SERVER_INFO) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SERVER_INFO) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMgr(dAtA, i, uint64(m.Id.Size()))
		n1, err := m.Id.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMgr(dAtA, i, uint64(m.Type))
	}
	if len(m.Addrs) > 0 {
		for _, s := range m.Addrs {
			dAtA[i] = 0x1a
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
	if len(m.Ports) > 0 {
		dAtA3 := make([]byte, len(m.Ports)*10)
		var j2 int
		for _, num1 := range m.Ports {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintMgr(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	if len(m.Overload) > 0 {
		dAtA5 := make([]byte, len(m.Overload)*10)
		var j4 int
		for _, num := range m.Overload {
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMgr(dAtA, i, uint64(j4))
		i += copy(dAtA[i:], dAtA5[:j4])
	}
	if len(m.Version) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMgr(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	return i, nil
}

func (m *MSG_MGR_PING) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MSG_MGR_PING) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *MSG_MGR_REGISTER_SERVER) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MSG_MGR_REGISTER_SERVER) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMgr(dAtA, i, uint64(m.Data.Size()))
		n6, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMgr(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if m.TargetServerType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMgr(dAtA, i, uint64(m.TargetServerType))
	}
	if m.TargetServerID != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMgr(dAtA, i, uint64(m.TargetServerID.Size()))
		n7, err := m.TargetServerID.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func (m *MSG_MGR_LOSE_SERVER) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MSG_MGR_LOSE_SERVER) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMgr(dAtA, i, uint64(m.Id.Size()))
		n8, err := m.Id.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMgr(dAtA, i, uint64(m.Type))
	}
	return i, nil
}

func encodeFixed64Mgr(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Mgr(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMgr(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CMD_MGR) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *SERVER_INFO) Size() (n int) {
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovMgr(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovMgr(uint64(m.Type))
	}
	if len(m.Addrs) > 0 {
		for _, s := range m.Addrs {
			l = len(s)
			n += 1 + l + sovMgr(uint64(l))
		}
	}
	if len(m.Ports) > 0 {
		l = 0
		for _, e := range m.Ports {
			l += sovMgr(uint64(e))
		}
		n += 1 + sovMgr(uint64(l)) + l
	}
	if len(m.Overload) > 0 {
		l = 0
		for _, e := range m.Overload {
			l += sovMgr(uint64(e))
		}
		n += 1 + sovMgr(uint64(l)) + l
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovMgr(uint64(l))
	}
	return n
}

func (m *MSG_MGR_PING) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *MSG_MGR_REGISTER_SERVER) Size() (n int) {
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovMgr(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovMgr(uint64(l))
	}
	if m.TargetServerType != 0 {
		n += 1 + sovMgr(uint64(m.TargetServerType))
	}
	if m.TargetServerID != nil {
		l = m.TargetServerID.Size()
		n += 1 + l + sovMgr(uint64(l))
	}
	return n
}

func (m *MSG_MGR_LOSE_SERVER) Size() (n int) {
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovMgr(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovMgr(uint64(m.Type))
	}
	return n
}

func sovMgr(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMgr(x uint64) (n int) {
	return sovMgr(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CMD_MGR) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMgr
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
			return fmt.Errorf("proto: CMD_MGR: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CMD_MGR: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMgr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMgr
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
func (m *SERVER_INFO) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMgr
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
			return fmt.Errorf("proto: SERVER_INFO: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SERVER_INFO: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMgr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMgr
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &SERVER_ID{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMgr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addrs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMgr
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
				return ErrInvalidLengthMgr
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addrs = append(m.Addrs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMgr
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ports = append(m.Ports, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMgr
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMgr
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMgr
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ports = append(m.Ports, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ports", wireType)
			}
		case 5:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMgr
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Overload = append(m.Overload, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMgr
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMgr
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMgr
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Overload = append(m.Overload, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Overload", wireType)
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMgr
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
				return ErrInvalidLengthMgr
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMgr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMgr
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
func (m *MSG_MGR_PING) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMgr
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
			return fmt.Errorf("proto: MSG_MGR_PING: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MSG_MGR_PING: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMgr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMgr
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
func (m *MSG_MGR_REGISTER_SERVER) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMgr
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
			return fmt.Errorf("proto: MSG_MGR_REGISTER_SERVER: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MSG_MGR_REGISTER_SERVER: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMgr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMgr
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &SERVER_INFO{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMgr
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
				return ErrInvalidLengthMgr
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetServerType", wireType)
			}
			m.TargetServerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMgr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetServerType |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetServerID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMgr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMgr
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TargetServerID == nil {
				m.TargetServerID = &SERVER_ID{}
			}
			if err := m.TargetServerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMgr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMgr
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
func (m *MSG_MGR_LOSE_SERVER) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMgr
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
			return fmt.Errorf("proto: MSG_MGR_LOSE_SERVER: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MSG_MGR_LOSE_SERVER: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMgr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMgr
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &SERVER_ID{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMgr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMgr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMgr
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
func skipMgr(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMgr
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
					return 0, ErrIntOverflowMgr
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
					return 0, ErrIntOverflowMgr
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
				return 0, ErrInvalidLengthMgr
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMgr
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
				next, err := skipMgr(dAtA[start:])
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
	ErrInvalidLengthMgr = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMgr   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mgr.proto", fileDescriptorMgr) }

var fileDescriptorMgr = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcd, 0xaa, 0x9b, 0x40,
	0x18, 0x86, 0x33, 0x6a, 0x7e, 0xfc, 0xcc, 0x8f, 0x4c, 0x5a, 0x3a, 0x64, 0x21, 0x62, 0x37, 0xb6,
	0x8b, 0x2c, 0xda, 0x65, 0x57, 0x69, 0x34, 0x32, 0x10, 0x35, 0x8c, 0x36, 0x5b, 0xb1, 0x51, 0x42,
	0x69, 0x92, 0x09, 0x13, 0x09, 0xf4, 0x2e, 0x7a, 0x19, 0xbd, 0x8f, 0x6e, 0xba, 0xec, 0x25, 0x94,
	0xf4, 0x46, 0x0e, 0xea, 0xf1, 0x10, 0x12, 0xce, 0xe6, 0xac, 0x66, 0xde, 0x87, 0x77, 0xe0, 0x7b,
	0xe6, 0x03, 0x75, 0xbf, 0x15, 0xd3, 0xa3, 0xe0, 0x05, 0xc7, 0xbd, 0xea, 0xd8, 0xf0, 0xdd, 0xa4,
	0xbf, 0xe1, 0xfb, 0x3d, 0x3f, 0xd4, 0xdc, 0x0a, 0xa0, 0x3b, 0xf7, 0x9d, 0xc4, 0xf7, 0x98, 0x35,
	0x07, 0xc5, 0x0d, 0xbe, 0xf8, 0x58, 0x83, 0x2e, 0x0d, 0xd6, 0xb3, 0x25, 0x75, 0xf4, 0x16, 0xee,
	0x81, 0xb2, 0xa2, 0x81, 0xa7, 0x23, 0x3c, 0x86, 0x11, 0x73, 0x3d, 0x1a, 0xc5, 0x2e, 0x4b, 0x22,
	0x97, 0xad, 0x5d, 0xa6, 0x4b, 0x78, 0x04, 0xda, 0x32, 0x8c, 0xdc, 0x06, 0xc8, 0xd6, 0x2f, 0x04,
	0x5a, 0x1d, 0x12, 0x1a, 0x2c, 0x42, 0xfc, 0x16, 0x24, 0x9a, 0x11, 0x64, 0x22, 0x5b, 0xfb, 0x30,
	0x9e, 0x36, 0x43, 0x4c, 0x9b, 0x8a, 0xc3, 0x24, 0x9a, 0x61, 0x0c, 0x4a, 0xfc, 0xe3, 0x98, 0x13,
	0xc9, 0x44, 0xf6, 0x80, 0x55, 0x77, 0xfc, 0x0a, 0xda, 0xb3, 0x2c, 0x13, 0x27, 0x22, 0x9b, 0xb2,
	0xad, 0xb2, 0x3a, 0x94, 0x74, 0xc5, 0x45, 0x71, 0x22, 0x8a, 0x29, 0xdb, 0x6d, 0x56, 0x07, 0x3c,
	0x81, 0x5e, 0x78, 0xce, 0xc5, 0x8e, 0xa7, 0x19, 0x69, 0x9b, 0xb2, 0x3d, 0x60, 0x4f, 0x19, 0x13,
	0xe8, 0xae, 0x73, 0x71, 0xfa, 0xc6, 0x0f, 0xa4, 0x63, 0x22, 0x5b, 0x65, 0x4d, 0xb4, 0x86, 0xd0,
	0xf7, 0x23, 0xaf, 0x54, 0x4f, 0x4a, 0x45, 0xeb, 0x37, 0x82, 0x37, 0x0d, 0xb8, 0x31, 0xc5, 0xef,
	0x40, 0x71, 0xd2, 0x22, 0x7d, 0x14, 0x79, 0x7d, 0x2f, 0x12, 0x2c, 0x42, 0x56, 0x55, 0xca, 0x11,
	0x63, 0xfe, 0x3d, 0x3f, 0x54, 0x36, 0x2a, 0xab, 0x03, 0x7e, 0x0f, 0x7a, 0x9c, 0x8a, 0x6d, 0x5e,
	0x44, 0xb9, 0x38, 0xe7, 0xa2, 0xd2, 0x95, 0x2b, 0xdd, 0x3b, 0x8e, 0x3f, 0xc1, 0xf0, 0x9a, 0x51,
	0x87, 0x28, 0xcf, 0xff, 0xdf, 0x4d, 0xd5, 0x0a, 0x60, 0xdc, 0x48, 0x5c, 0x6d, 0xe6, 0xc5, 0x7b,
	0xf8, 0xac, 0xff, 0xb9, 0x18, 0xe8, 0xef, 0xc5, 0x40, 0xff, 0x2e, 0x06, 0xfa, 0xf9, 0xdf, 0x68,
	0x7d, 0xed, 0x54, 0xaf, 0x3f, 0x3e, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x91, 0x1e, 0x5c, 0x5e,
	0x02, 0x00, 0x00,
}
