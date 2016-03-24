// Code generated by protoc-gen-gogo.
// source: etcdserver.proto
// DO NOT EDIT!

/*
	Package etcdserverpb is a generated protocol buffer package.

	It is generated from these files:
		etcdserver.proto
		raft_internal.proto
		rpc.proto

	It has these top-level messages:
		Request
		Metadata
		InternalRaftRequest
		EmptyResponse
		ResponseHeader
		RangeRequest
		RangeResponse
		PutRequest
		PutResponse
		DeleteRangeRequest
		DeleteRangeResponse
		RequestUnion
		ResponseUnion
		Compare
		TxnRequest
		TxnResponse
		CompactionRequest
		CompactionResponse
		HashRequest
		HashResponse
		WatchRequest
		WatchCreateRequest
		WatchCancelRequest
		WatchResponse
		LeaseCreateRequest
		LeaseCreateResponse
		LeaseRevokeRequest
		LeaseRevokeResponse
		LeaseKeepAliveRequest
		LeaseKeepAliveResponse
		Member
		MemberAddRequest
		MemberAddResponse
		MemberRemoveRequest
		MemberRemoveResponse
		MemberUpdateRequest
		MemberUpdateResponse
		MemberListRequest
		MemberListResponse
		DefragmentRequest
		DefragmentResponse
		AuthEnableRequest
		AuthDisableRequest
		AuthenticateRequest
		UserAddRequest
		UserGetRequest
		UserDeleteRequest
		UserChangePasswordRequest
		UserGrantRequest
		UserRevokeRequest
		RoleAddRequest
		RoleGetRequest
		RoleDeleteRequest
		RoleGrantRequest
		RoleRevokeRequest
		AuthEnableResponse
		AuthDisableResponse
		AuthenticateResponse
		UserAddResponse
		UserGetResponse
		UserDeleteResponse
		UserChangePasswordResponse
		UserGrantResponse
		UserRevokeResponse
		RoleAddResponse
		RoleGetResponse
		RoleDeleteResponse
		RoleGrantResponse
		RoleRevokeResponse
*/
package etcdserverpb

import (
	"fmt"

	proto "github.com/gogo/protobuf/proto"
)

import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request struct {
	ID               uint64 `protobuf:"varint,1,opt,name=ID" json:"ID"`
	Method           string `protobuf:"bytes,2,opt,name=Method" json:"Method"`
	Path             string `protobuf:"bytes,3,opt,name=Path" json:"Path"`
	Val              string `protobuf:"bytes,4,opt,name=Val" json:"Val"`
	Dir              bool   `protobuf:"varint,5,opt,name=Dir" json:"Dir"`
	PrevValue        string `protobuf:"bytes,6,opt,name=PrevValue" json:"PrevValue"`
	PrevIndex        uint64 `protobuf:"varint,7,opt,name=PrevIndex" json:"PrevIndex"`
	PrevExist        *bool  `protobuf:"varint,8,opt,name=PrevExist" json:"PrevExist,omitempty"`
	Expiration       int64  `protobuf:"varint,9,opt,name=Expiration" json:"Expiration"`
	Wait             bool   `protobuf:"varint,10,opt,name=Wait" json:"Wait"`
	Since            uint64 `protobuf:"varint,11,opt,name=Since" json:"Since"`
	Recursive        bool   `protobuf:"varint,12,opt,name=Recursive" json:"Recursive"`
	Sorted           bool   `protobuf:"varint,13,opt,name=Sorted" json:"Sorted"`
	Quorum           bool   `protobuf:"varint,14,opt,name=Quorum" json:"Quorum"`
	Time             int64  `protobuf:"varint,15,opt,name=Time" json:"Time"`
	Stream           bool   `protobuf:"varint,16,opt,name=Stream" json:"Stream"`
	Refresh          *bool  `protobuf:"varint,17,opt,name=Refresh" json:"Refresh,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

type Metadata struct {
	NodeID           uint64 `protobuf:"varint,1,opt,name=NodeID" json:"NodeID"`
	ClusterID        uint64 `protobuf:"varint,2,opt,name=ClusterID" json:"ClusterID"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}

func init() {
	proto.RegisterType((*Request)(nil), "etcdserverpb.Request")
	proto.RegisterType((*Metadata)(nil), "etcdserverpb.Metadata")
}
func (m *Request) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Request) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintEtcdserver(data, i, uint64(m.ID))
	data[i] = 0x12
	i++
	i = encodeVarintEtcdserver(data, i, uint64(len(m.Method)))
	i += copy(data[i:], m.Method)
	data[i] = 0x1a
	i++
	i = encodeVarintEtcdserver(data, i, uint64(len(m.Path)))
	i += copy(data[i:], m.Path)
	data[i] = 0x22
	i++
	i = encodeVarintEtcdserver(data, i, uint64(len(m.Val)))
	i += copy(data[i:], m.Val)
	data[i] = 0x28
	i++
	if m.Dir {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x32
	i++
	i = encodeVarintEtcdserver(data, i, uint64(len(m.PrevValue)))
	i += copy(data[i:], m.PrevValue)
	data[i] = 0x38
	i++
	i = encodeVarintEtcdserver(data, i, uint64(m.PrevIndex))
	if m.PrevExist != nil {
		data[i] = 0x40
		i++
		if *m.PrevExist {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	data[i] = 0x48
	i++
	i = encodeVarintEtcdserver(data, i, uint64(m.Expiration))
	data[i] = 0x50
	i++
	if m.Wait {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x58
	i++
	i = encodeVarintEtcdserver(data, i, uint64(m.Since))
	data[i] = 0x60
	i++
	if m.Recursive {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x68
	i++
	if m.Sorted {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x70
	i++
	if m.Quorum {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x78
	i++
	i = encodeVarintEtcdserver(data, i, uint64(m.Time))
	data[i] = 0x80
	i++
	data[i] = 0x1
	i++
	if m.Stream {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	if m.Refresh != nil {
		data[i] = 0x88
		i++
		data[i] = 0x1
		i++
		if *m.Refresh {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Metadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Metadata) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintEtcdserver(data, i, uint64(m.NodeID))
	data[i] = 0x10
	i++
	i = encodeVarintEtcdserver(data, i, uint64(m.ClusterID))
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Etcdserver(data []byte, offset int, v uint64) int {
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
func encodeFixed32Etcdserver(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintEtcdserver(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovEtcdserver(uint64(m.ID))
	l = len(m.Method)
	n += 1 + l + sovEtcdserver(uint64(l))
	l = len(m.Path)
	n += 1 + l + sovEtcdserver(uint64(l))
	l = len(m.Val)
	n += 1 + l + sovEtcdserver(uint64(l))
	n += 2
	l = len(m.PrevValue)
	n += 1 + l + sovEtcdserver(uint64(l))
	n += 1 + sovEtcdserver(uint64(m.PrevIndex))
	if m.PrevExist != nil {
		n += 2
	}
	n += 1 + sovEtcdserver(uint64(m.Expiration))
	n += 2
	n += 1 + sovEtcdserver(uint64(m.Since))
	n += 2
	n += 2
	n += 2
	n += 1 + sovEtcdserver(uint64(m.Time))
	n += 3
	if m.Refresh != nil {
		n += 3
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Metadata) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovEtcdserver(uint64(m.NodeID))
	n += 1 + sovEtcdserver(uint64(m.ClusterID))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEtcdserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEtcdserver(x uint64) (n int) {
	return sovEtcdserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEtcdserver
			}
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
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEtcdserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEtcdserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEtcdserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Val = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dir", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Dir = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEtcdserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevValue = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevIndex", wireType)
			}
			m.PrevIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.PrevIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevExist", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.PrevExist = &b
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			m.Expiration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Expiration |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wait", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Wait = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Since", wireType)
			}
			m.Since = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Since |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recursive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Recursive = bool(v != 0)
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sorted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sorted = bool(v != 0)
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quorum", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Quorum = bool(v != 0)
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stream", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Stream = bool(v != 0)
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refresh", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Refresh = &b
		default:
			iNdEx = preIndex
			skippy, err := skipEtcdserver(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEtcdserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Metadata) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEtcdserver
			}
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
		if wireType == 4 {
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NodeID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			m.ClusterID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ClusterID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEtcdserver(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEtcdserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEtcdserver(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEtcdserver
			}
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
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEtcdserver
				}
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
				if shift >= 64 {
					return 0, ErrIntOverflowEtcdserver
				}
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
				return 0, ErrInvalidLengthEtcdserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEtcdserver
					}
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
				next, err := skipEtcdserver(data[start:])
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
	ErrInvalidLengthEtcdserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEtcdserver   = fmt.Errorf("proto: integer overflow")
)
