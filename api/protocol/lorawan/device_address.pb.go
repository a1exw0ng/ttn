// Code generated by protoc-gen-gogo.
// source: github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device_address.proto
// DO NOT EDIT!

/*
	Package lorawan is a generated protocol buffer package.

	It is generated from these files:
		github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device_address.proto

	It has these top-level messages:
		PrefixesRequest
		PrefixesResponse
		DevAddrRequest
		DevAddrResponse
*/
package lorawan

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_TheThingsNetwork_ttn_core_types "github.com/TheThingsNetwork/ttn/core/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PrefixesRequest struct {
}

func (m *PrefixesRequest) Reset()                    { *m = PrefixesRequest{} }
func (*PrefixesRequest) ProtoMessage()               {}
func (*PrefixesRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAddress, []int{0} }

type PrefixesResponse struct {
	// The prefixes that are in use or available
	Prefixes []*PrefixesResponse_PrefixMapping `protobuf:"bytes,1,rep,name=prefixes" json:"prefixes,omitempty"`
}

func (m *PrefixesResponse) Reset()                    { *m = PrefixesResponse{} }
func (*PrefixesResponse) ProtoMessage()               {}
func (*PrefixesResponse) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAddress, []int{1} }

func (m *PrefixesResponse) GetPrefixes() []*PrefixesResponse_PrefixMapping {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

type PrefixesResponse_PrefixMapping struct {
	// The prefix that can be used
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Usage constraints of this prefix (see activation_constraints in device.proto)
	Usage []string `protobuf:"bytes,2,rep,name=usage" json:"usage,omitempty"`
}

func (m *PrefixesResponse_PrefixMapping) Reset()      { *m = PrefixesResponse_PrefixMapping{} }
func (*PrefixesResponse_PrefixMapping) ProtoMessage() {}
func (*PrefixesResponse_PrefixMapping) Descriptor() ([]byte, []int) {
	return fileDescriptorDeviceAddress, []int{1, 0}
}

func (m *PrefixesResponse_PrefixMapping) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *PrefixesResponse_PrefixMapping) GetUsage() []string {
	if m != nil {
		return m.Usage
	}
	return nil
}

type DevAddrRequest struct {
	// The usage constraints (see activation_constraints in device.proto)
	Usage []string `protobuf:"bytes,1,rep,name=usage" json:"usage,omitempty"`
}

func (m *DevAddrRequest) Reset()                    { *m = DevAddrRequest{} }
func (*DevAddrRequest) ProtoMessage()               {}
func (*DevAddrRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAddress, []int{2} }

func (m *DevAddrRequest) GetUsage() []string {
	if m != nil {
		return m.Usage
	}
	return nil
}

type DevAddrResponse struct {
	DevAddr *github_com_TheThingsNetwork_ttn_core_types.DevAddr `protobuf:"bytes,1,opt,name=dev_addr,json=devAddr,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevAddr" json:"dev_addr,omitempty"`
}

func (m *DevAddrResponse) Reset()                    { *m = DevAddrResponse{} }
func (*DevAddrResponse) ProtoMessage()               {}
func (*DevAddrResponse) Descriptor() ([]byte, []int) { return fileDescriptorDeviceAddress, []int{3} }

func init() {
	proto.RegisterType((*PrefixesRequest)(nil), "lorawan.PrefixesRequest")
	proto.RegisterType((*PrefixesResponse)(nil), "lorawan.PrefixesResponse")
	proto.RegisterType((*PrefixesResponse_PrefixMapping)(nil), "lorawan.PrefixesResponse.PrefixMapping")
	proto.RegisterType((*DevAddrRequest)(nil), "lorawan.DevAddrRequest")
	proto.RegisterType((*DevAddrResponse)(nil), "lorawan.DevAddrResponse")
}
func (this *PrefixesRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*PrefixesRequest)
	if !ok {
		that2, ok := that.(PrefixesRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *PrefixesRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *PrefixesRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *PrefixesRequest but is not nil && this == nil")
	}
	return nil
}
func (this *PrefixesRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*PrefixesRequest)
	if !ok {
		that2, ok := that.(PrefixesRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *PrefixesResponse) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*PrefixesResponse)
	if !ok {
		that2, ok := that.(PrefixesResponse)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *PrefixesResponse")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *PrefixesResponse but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *PrefixesResponse but is not nil && this == nil")
	}
	if len(this.Prefixes) != len(that1.Prefixes) {
		return fmt.Errorf("Prefixes this(%v) Not Equal that(%v)", len(this.Prefixes), len(that1.Prefixes))
	}
	for i := range this.Prefixes {
		if !this.Prefixes[i].Equal(that1.Prefixes[i]) {
			return fmt.Errorf("Prefixes this[%v](%v) Not Equal that[%v](%v)", i, this.Prefixes[i], i, that1.Prefixes[i])
		}
	}
	return nil
}
func (this *PrefixesResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*PrefixesResponse)
	if !ok {
		that2, ok := that.(PrefixesResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Prefixes) != len(that1.Prefixes) {
		return false
	}
	for i := range this.Prefixes {
		if !this.Prefixes[i].Equal(that1.Prefixes[i]) {
			return false
		}
	}
	return true
}
func (this *PrefixesResponse_PrefixMapping) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*PrefixesResponse_PrefixMapping)
	if !ok {
		that2, ok := that.(PrefixesResponse_PrefixMapping)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *PrefixesResponse_PrefixMapping")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *PrefixesResponse_PrefixMapping but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *PrefixesResponse_PrefixMapping but is not nil && this == nil")
	}
	if this.Prefix != that1.Prefix {
		return fmt.Errorf("Prefix this(%v) Not Equal that(%v)", this.Prefix, that1.Prefix)
	}
	if len(this.Usage) != len(that1.Usage) {
		return fmt.Errorf("Usage this(%v) Not Equal that(%v)", len(this.Usage), len(that1.Usage))
	}
	for i := range this.Usage {
		if this.Usage[i] != that1.Usage[i] {
			return fmt.Errorf("Usage this[%v](%v) Not Equal that[%v](%v)", i, this.Usage[i], i, that1.Usage[i])
		}
	}
	return nil
}
func (this *PrefixesResponse_PrefixMapping) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*PrefixesResponse_PrefixMapping)
	if !ok {
		that2, ok := that.(PrefixesResponse_PrefixMapping)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Prefix != that1.Prefix {
		return false
	}
	if len(this.Usage) != len(that1.Usage) {
		return false
	}
	for i := range this.Usage {
		if this.Usage[i] != that1.Usage[i] {
			return false
		}
	}
	return true
}
func (this *DevAddrRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*DevAddrRequest)
	if !ok {
		that2, ok := that.(DevAddrRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *DevAddrRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *DevAddrRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *DevAddrRequest but is not nil && this == nil")
	}
	if len(this.Usage) != len(that1.Usage) {
		return fmt.Errorf("Usage this(%v) Not Equal that(%v)", len(this.Usage), len(that1.Usage))
	}
	for i := range this.Usage {
		if this.Usage[i] != that1.Usage[i] {
			return fmt.Errorf("Usage this[%v](%v) Not Equal that[%v](%v)", i, this.Usage[i], i, that1.Usage[i])
		}
	}
	return nil
}
func (this *DevAddrRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DevAddrRequest)
	if !ok {
		that2, ok := that.(DevAddrRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Usage) != len(that1.Usage) {
		return false
	}
	for i := range this.Usage {
		if this.Usage[i] != that1.Usage[i] {
			return false
		}
	}
	return true
}
func (this *DevAddrResponse) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*DevAddrResponse)
	if !ok {
		that2, ok := that.(DevAddrResponse)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *DevAddrResponse")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *DevAddrResponse but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *DevAddrResponse but is not nil && this == nil")
	}
	if that1.DevAddr == nil {
		if this.DevAddr != nil {
			return fmt.Errorf("this.DevAddr != nil && that1.DevAddr == nil")
		}
	} else if !this.DevAddr.Equal(*that1.DevAddr) {
		return fmt.Errorf("DevAddr this(%v) Not Equal that(%v)", this.DevAddr, that1.DevAddr)
	}
	return nil
}
func (this *DevAddrResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DevAddrResponse)
	if !ok {
		that2, ok := that.(DevAddrResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if that1.DevAddr == nil {
		if this.DevAddr != nil {
			return false
		}
	} else if !this.DevAddr.Equal(*that1.DevAddr) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DevAddrManager service

type DevAddrManagerClient interface {
	// Get all prefixes that are in use or available
	GetPrefixes(ctx context.Context, in *PrefixesRequest, opts ...grpc.CallOption) (*PrefixesResponse, error)
	// Request a device address
	GetDevAddr(ctx context.Context, in *DevAddrRequest, opts ...grpc.CallOption) (*DevAddrResponse, error)
}

type devAddrManagerClient struct {
	cc *grpc.ClientConn
}

func NewDevAddrManagerClient(cc *grpc.ClientConn) DevAddrManagerClient {
	return &devAddrManagerClient{cc}
}

func (c *devAddrManagerClient) GetPrefixes(ctx context.Context, in *PrefixesRequest, opts ...grpc.CallOption) (*PrefixesResponse, error) {
	out := new(PrefixesResponse)
	err := grpc.Invoke(ctx, "/lorawan.DevAddrManager/GetPrefixes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devAddrManagerClient) GetDevAddr(ctx context.Context, in *DevAddrRequest, opts ...grpc.CallOption) (*DevAddrResponse, error) {
	out := new(DevAddrResponse)
	err := grpc.Invoke(ctx, "/lorawan.DevAddrManager/GetDevAddr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DevAddrManager service

type DevAddrManagerServer interface {
	// Get all prefixes that are in use or available
	GetPrefixes(context.Context, *PrefixesRequest) (*PrefixesResponse, error)
	// Request a device address
	GetDevAddr(context.Context, *DevAddrRequest) (*DevAddrResponse, error)
}

func RegisterDevAddrManagerServer(s *grpc.Server, srv DevAddrManagerServer) {
	s.RegisterService(&_DevAddrManager_serviceDesc, srv)
}

func _DevAddrManager_GetPrefixes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrefixesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevAddrManagerServer).GetPrefixes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorawan.DevAddrManager/GetPrefixes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevAddrManagerServer).GetPrefixes(ctx, req.(*PrefixesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevAddrManager_GetDevAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevAddrManagerServer).GetDevAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorawan.DevAddrManager/GetDevAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevAddrManagerServer).GetDevAddr(ctx, req.(*DevAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DevAddrManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lorawan.DevAddrManager",
	HandlerType: (*DevAddrManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrefixes",
			Handler:    _DevAddrManager_GetPrefixes_Handler,
		},
		{
			MethodName: "GetDevAddr",
			Handler:    _DevAddrManager_GetDevAddr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device_address.proto",
}

func (m *PrefixesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrefixesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *PrefixesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrefixesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Prefixes) > 0 {
		for _, msg := range m.Prefixes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeviceAddress(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PrefixesResponse_PrefixMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrefixesResponse_PrefixMapping) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Prefix) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeviceAddress(dAtA, i, uint64(len(m.Prefix)))
		i += copy(dAtA[i:], m.Prefix)
	}
	if len(m.Usage) > 0 {
		for _, s := range m.Usage {
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
	return i, nil
}

func (m *DevAddrRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevAddrRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Usage) > 0 {
		for _, s := range m.Usage {
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

func (m *DevAddrResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevAddrResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DevAddr != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeviceAddress(dAtA, i, uint64(m.DevAddr.Size()))
		n1, err := m.DevAddr.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeFixed64DeviceAddress(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32DeviceAddress(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDeviceAddress(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PrefixesRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *PrefixesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Prefixes) > 0 {
		for _, e := range m.Prefixes {
			l = e.Size()
			n += 1 + l + sovDeviceAddress(uint64(l))
		}
	}
	return n
}

func (m *PrefixesResponse_PrefixMapping) Size() (n int) {
	var l int
	_ = l
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovDeviceAddress(uint64(l))
	}
	if len(m.Usage) > 0 {
		for _, s := range m.Usage {
			l = len(s)
			n += 1 + l + sovDeviceAddress(uint64(l))
		}
	}
	return n
}

func (m *DevAddrRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Usage) > 0 {
		for _, s := range m.Usage {
			l = len(s)
			n += 1 + l + sovDeviceAddress(uint64(l))
		}
	}
	return n
}

func (m *DevAddrResponse) Size() (n int) {
	var l int
	_ = l
	if m.DevAddr != nil {
		l = m.DevAddr.Size()
		n += 1 + l + sovDeviceAddress(uint64(l))
	}
	return n
}

func sovDeviceAddress(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeviceAddress(x uint64) (n int) {
	return sovDeviceAddress(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PrefixesRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PrefixesRequest{`,
		`}`,
	}, "")
	return s
}
func (this *PrefixesResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PrefixesResponse{`,
		`Prefixes:` + strings.Replace(fmt.Sprintf("%v", this.Prefixes), "PrefixesResponse_PrefixMapping", "PrefixesResponse_PrefixMapping", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PrefixesResponse_PrefixMapping) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PrefixesResponse_PrefixMapping{`,
		`Prefix:` + fmt.Sprintf("%v", this.Prefix) + `,`,
		`Usage:` + fmt.Sprintf("%v", this.Usage) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DevAddrRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DevAddrRequest{`,
		`Usage:` + fmt.Sprintf("%v", this.Usage) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DevAddrResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DevAddrResponse{`,
		`DevAddr:` + fmt.Sprintf("%v", this.DevAddr) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDeviceAddress(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PrefixesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAddress
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
			return fmt.Errorf("proto: PrefixesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrefixesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAddress
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
func (m *PrefixesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAddress
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
			return fmt.Errorf("proto: PrefixesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrefixesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefixes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAddress
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
				return ErrInvalidLengthDeviceAddress
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefixes = append(m.Prefixes, &PrefixesResponse_PrefixMapping{})
			if err := m.Prefixes[len(m.Prefixes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAddress
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
func (m *PrefixesResponse_PrefixMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAddress
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
			return fmt.Errorf("proto: PrefixMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrefixMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAddress
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
				return ErrInvalidLengthDeviceAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Usage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAddress
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
				return ErrInvalidLengthDeviceAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Usage = append(m.Usage, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAddress
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
func (m *DevAddrRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAddress
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
			return fmt.Errorf("proto: DevAddrRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevAddrRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Usage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAddress
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
				return ErrInvalidLengthDeviceAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Usage = append(m.Usage, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAddress
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
func (m *DevAddrResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceAddress
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
			return fmt.Errorf("proto: DevAddrResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevAddrResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDeviceAddress
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevAddr
			m.DevAddr = &v
			if err := m.DevAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceAddress
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
func skipDeviceAddress(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceAddress
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
					return 0, ErrIntOverflowDeviceAddress
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
					return 0, ErrIntOverflowDeviceAddress
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
				return 0, ErrInvalidLengthDeviceAddress
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeviceAddress
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
				next, err := skipDeviceAddress(dAtA[start:])
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
	ErrInvalidLengthDeviceAddress = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceAddress   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device_address.proto", fileDescriptorDeviceAddress)
}

var fileDescriptorDeviceAddress = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0f, 0xc9, 0x48, 0x0d, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f,
	0xf6, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0x29, 0xc9, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0xcf, 0xc9, 0x2f, 0x4a, 0x2c, 0x4f, 0xcc,
	0xd3, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0x8d, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0xd6,
	0x03, 0xcb, 0x0b, 0xb1, 0x43, 0x65, 0xa5, 0x74, 0x91, 0xcc, 0x4c, 0xcf, 0x4f, 0xcf, 0x87, 0xe8,
	0x4f, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x4f, 0x49, 0x90, 0x8b, 0x3f, 0xa0,
	0x28, 0x35, 0x2d, 0xb3, 0x22, 0xb5, 0x38, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x69, 0x1a,
	0x23, 0x97, 0x00, 0x42, 0xac, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x99, 0x8b, 0xa3, 0x00,
	0x2a, 0x26, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xae, 0x07, 0xb5, 0x52, 0x0f, 0x5d, 0x31,
	0x54, 0xc0, 0x37, 0xb1, 0xa0, 0x20, 0x33, 0x2f, 0x3d, 0x08, 0xae, 0x51, 0xca, 0x96, 0x8b, 0x17,
	0x45, 0x4a, 0x48, 0x8c, 0x8b, 0x0d, 0x22, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5,
	0x09, 0x89, 0x70, 0xb1, 0x96, 0x16, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x06,
	0x41, 0x38, 0x4a, 0x6a, 0x5c, 0x7c, 0x2e, 0xa9, 0x65, 0x8e, 0x29, 0x29, 0x45, 0x50, 0xa7, 0x22,
	0xd4, 0x31, 0x22, 0xab, 0x4b, 0xe1, 0xe2, 0x87, 0xab, 0x83, 0x3a, 0x3f, 0x90, 0x8b, 0x23, 0x25,
	0xb5, 0x0c, 0x1c, 0x66, 0x60, 0xab, 0x78, 0x9c, 0xcc, 0x6e, 0xdd, 0x93, 0x37, 0x22, 0x14, 0xfe,
	0xc9, 0xf9, 0x45, 0xa9, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x7a, 0x30, 0x13, 0xd9, 0x53, 0x20,
	0x0c, 0xa3, 0xa9, 0x8c, 0x70, 0xe7, 0xf8, 0x26, 0xe6, 0x25, 0xa6, 0xa7, 0x16, 0x09, 0x39, 0x71,
	0x71, 0xbb, 0xa7, 0x96, 0xc0, 0x82, 0x43, 0x48, 0x02, 0x4b, 0x08, 0x81, 0xdd, 0x2d, 0x25, 0x89,
	0x33, 0xec, 0x84, 0xec, 0xb9, 0xb8, 0xdc, 0x53, 0x4b, 0xa0, 0x06, 0x0b, 0x89, 0xc3, 0x15, 0xa2,
	0xfa, 0x5c, 0x4a, 0x02, 0x53, 0x02, 0x62, 0x80, 0x53, 0xd8, 0x8d, 0x87, 0x72, 0x0c, 0x0f, 0x1e,
	0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x21, 0xca, 0x84, 0x9c, 0x64,
	0x97, 0xc4, 0x06, 0x16, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x05, 0xf6, 0x9f, 0xb5,
	0x02, 0x00, 0x00,
}