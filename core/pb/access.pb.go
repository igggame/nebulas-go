// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: access.proto

package corepb

import (
	fmt "fmt"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type Access struct {
	// WhiteList of transactions, some of which are accepted only when conditions permit
	Whitelist *Whitelist `protobuf:"bytes,1,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
	// Blacklists of transactions, some of which are not accepted on the blacklists
	Blacklist *Blacklist `protobuf:"bytes,2,opt,name=blacklist,proto3" json:"blacklist,omitempty"`
}

func (m *Access) Reset() { *m = Access{} }
func (m *Access) String() string {
	if b, err := m.Marshal(); err == nil {
		return string(b)
	}
	return string("")
}
func (*Access) ProtoMessage() {}

func (m *Access) GetWhitelist() *Whitelist {
	if m != nil {
		return m.Whitelist
	}
	return nil
}

func (m *Access) GetBlacklist() *Blacklist {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

type Whitelist struct {
	// TODO(larry): later may change
	NbrePublisher []string `protobuf:"bytes,1,rep,name=nbre_publisher,json=nbrePublisher,proto3" json:"nbre_publisher,omitempty"`
}

func (m *Whitelist) Reset() { *m = Whitelist{} }
func (m *Whitelist) String() string {
	if b, err := m.Marshal(); err == nil {
		return string(b)
	}
	return string("")
}
func (*Whitelist) ProtoMessage() {}

func (m *Whitelist) GetNbrePublisher() []string {
	if m != nil {
		return m.NbrePublisher
	}
	return nil
}

type Blacklist struct {
	// Hex string of the transaction's from address.
	From []string `protobuf:"bytes,1,rep,name=from,proto3" json:"from,omitempty"`
	// Hex string of the transaction's to address.
	To []string `protobuf:"bytes,2,rep,name=to,proto3" json:"to,omitempty"`
	// contract access data
	Contracts []*Contract `protobuf:"bytes,3,rep,name=contracts,proto3" json:"contracts,omitempty"`
}

func (m *Blacklist) Reset() { *m = Blacklist{} }
func (m *Blacklist) String() string {
	if b, err := m.Marshal(); err == nil {
		return string(b)
	}
	return string("")
}
func (*Blacklist) ProtoMessage() {}

func (m *Blacklist) GetFrom() []string {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Blacklist) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Blacklist) GetContracts() []*Contract {
	if m != nil {
		return m.Contracts
	}
	return nil
}

type Contract struct {
	// Hex string of the call type transaction's to address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// slice of call type transaction's function string.
	Functions []string `protobuf:"bytes,2,rep,name=functions,proto3" json:"functions,omitempty"`
	// slice of deploy type transaction's data keyword.
	Keywords []string `protobuf:"bytes,3,rep,name=keywords,proto3" json:"keywords,omitempty"`
}

func (m *Contract) Reset() { *m = Contract{} }
func (m *Contract) String() string {
	if b, err := m.Marshal(); err == nil {
		return string(b)
	}
	return string("")
}
func (*Contract) ProtoMessage() {}

func (m *Contract) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Contract) GetFunctions() []string {
	if m != nil {
		return m.Functions
	}
	return nil
}

func (m *Contract) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

func (m *Access) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Access) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Whitelist != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAccess(dAtA, i, uint64(m.Whitelist.Size()))
		n1, err := m.Whitelist.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Blacklist != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAccess(dAtA, i, uint64(m.Blacklist.Size()))
		n2, err := m.Blacklist.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Whitelist) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Whitelist) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NbrePublisher) > 0 {
		for _, s := range m.NbrePublisher {
			dAtA[i] = 0xa
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
	return i, nil
}

func (m *Blacklist) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Blacklist) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		for _, s := range m.From {
			dAtA[i] = 0xa
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
	if len(m.To) > 0 {
		for _, s := range m.To {
			dAtA[i] = 0x12
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
	if len(m.Contracts) > 0 {
		for _, msg := range m.Contracts {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintAccess(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAccess(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Functions) > 0 {
		for _, s := range m.Functions {
			dAtA[i] = 0x12
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
	if len(m.Keywords) > 0 {
		for _, s := range m.Keywords {
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
	return i, nil
}

func encodeVarintAccess(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Access) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Whitelist != nil {
		l = m.Whitelist.Size()
		n += 1 + l + sovAccess(uint64(l))
	}
	if m.Blacklist != nil {
		l = m.Blacklist.Size()
		n += 1 + l + sovAccess(uint64(l))
	}
	return n
}

func (m *Whitelist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.NbrePublisher) > 0 {
		for _, s := range m.NbrePublisher {
			l = len(s)
			n += 1 + l + sovAccess(uint64(l))
		}
	}
	return n
}

func (m *Blacklist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.From) > 0 {
		for _, s := range m.From {
			l = len(s)
			n += 1 + l + sovAccess(uint64(l))
		}
	}
	if len(m.To) > 0 {
		for _, s := range m.To {
			l = len(s)
			n += 1 + l + sovAccess(uint64(l))
		}
	}
	if len(m.Contracts) > 0 {
		for _, e := range m.Contracts {
			l = e.Size()
			n += 1 + l + sovAccess(uint64(l))
		}
	}
	return n
}

func (m *Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccess(uint64(l))
	}
	if len(m.Functions) > 0 {
		for _, s := range m.Functions {
			l = len(s)
			n += 1 + l + sovAccess(uint64(l))
		}
	}
	if len(m.Keywords) > 0 {
		for _, s := range m.Keywords {
			l = len(s)
			n += 1 + l + sovAccess(uint64(l))
		}
	}
	return n
}

func sovAccess(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAccess(x uint64) (n int) {
	return sovAccess(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Access) Unmarshal(dAtA []byte) error {
	l := uint32(len(dAtA))
	iNdEx := uint32(0)
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccess
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
		fieldNum := uint32(wire >> 3)
		wireType := uint32(wire & 0x7)
		_ = wireType
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Whitelist", wireType)
			}
			var msglen uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccess
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Whitelist == nil {
				m.Whitelist = &Whitelist{}
			}
			if err := m.Whitelist.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blacklist", wireType)
			}
			var msglen uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccess
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Blacklist == nil {
				m.Blacklist = &Blacklist{}
			}
			if err := m.Blacklist.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			return fmt.Errorf("proto: invalid data ")
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Whitelist) Unmarshal(dAtA []byte) error {
	l := uint32(len(dAtA))
	iNdEx := uint32(0)
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccess
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
		fieldNum := uint32(wire >> 3)
		wireType := uint32(wire & 0x7)
		_ = wireType
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NbrePublisher", wireType)
			}
			var stringLen uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccess
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + stringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NbrePublisher = append(m.NbrePublisher, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			return fmt.Errorf("proto: invalid data ")
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Blacklist) Unmarshal(dAtA []byte) error {
	l := uint32(len(dAtA))
	iNdEx := uint32(0)
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccess
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
		fieldNum := uint32(wire >> 3)
		wireType := uint32(wire & 0x7)
		_ = wireType
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccess
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + stringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = append(m.From, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccess
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + stringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = append(m.To, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
			}
			var msglen uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccess
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contracts = append(m.Contracts, &Contract{})
			if err := m.Contracts[len(m.Contracts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			return fmt.Errorf("proto: invalid data ")
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Contract) Unmarshal(dAtA []byte) error {
	l := uint32(len(dAtA))
	iNdEx := uint32(0)
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccess
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
		fieldNum := uint32(wire >> 3)
		wireType := uint32(wire & 0x7)
		_ = wireType
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccess
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + stringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Functions", wireType)
			}
			var stringLen uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccess
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + stringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Functions = append(m.Functions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keywords", wireType)
			}
			var stringLen uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccess
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + stringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keywords = append(m.Keywords, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			return fmt.Errorf("proto: invalid data ")
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

var (
	ErrInvalidLengthAccess = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccess   = fmt.Errorf("proto: integer overflow")
)
