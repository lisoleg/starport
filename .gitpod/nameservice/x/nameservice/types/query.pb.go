// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nameservice/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryGetWhoisRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetWhoisRequest) Reset()         { *m = QueryGetWhoisRequest{} }
func (m *QueryGetWhoisRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetWhoisRequest) ProtoMessage()    {}
func (*QueryGetWhoisRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37776ef2c2bc2f1b, []int{0}
}
func (m *QueryGetWhoisRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetWhoisRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetWhoisRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetWhoisRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetWhoisRequest.Merge(m, src)
}
func (m *QueryGetWhoisRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetWhoisRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetWhoisRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetWhoisRequest proto.InternalMessageInfo

func (m *QueryGetWhoisRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetWhoisResponse struct {
	Whois Whois `protobuf:"bytes,1,opt,name=whois,proto3" json:"whois"`
}

func (m *QueryGetWhoisResponse) Reset()         { *m = QueryGetWhoisResponse{} }
func (m *QueryGetWhoisResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetWhoisResponse) ProtoMessage()    {}
func (*QueryGetWhoisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_37776ef2c2bc2f1b, []int{1}
}
func (m *QueryGetWhoisResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetWhoisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetWhoisResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetWhoisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetWhoisResponse.Merge(m, src)
}
func (m *QueryGetWhoisResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetWhoisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetWhoisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetWhoisResponse proto.InternalMessageInfo

func (m *QueryGetWhoisResponse) GetWhois() Whois {
	if m != nil {
		return m.Whois
	}
	return Whois{}
}

type QueryAllWhoisRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllWhoisRequest) Reset()         { *m = QueryAllWhoisRequest{} }
func (m *QueryAllWhoisRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllWhoisRequest) ProtoMessage()    {}
func (*QueryAllWhoisRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37776ef2c2bc2f1b, []int{2}
}
func (m *QueryAllWhoisRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllWhoisRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllWhoisRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllWhoisRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllWhoisRequest.Merge(m, src)
}
func (m *QueryAllWhoisRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllWhoisRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllWhoisRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllWhoisRequest proto.InternalMessageInfo

func (m *QueryAllWhoisRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllWhoisResponse struct {
	Whois      []Whois             `protobuf:"bytes,1,rep,name=whois,proto3" json:"whois"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllWhoisResponse) Reset()         { *m = QueryAllWhoisResponse{} }
func (m *QueryAllWhoisResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllWhoisResponse) ProtoMessage()    {}
func (*QueryAllWhoisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_37776ef2c2bc2f1b, []int{3}
}
func (m *QueryAllWhoisResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllWhoisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllWhoisResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllWhoisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllWhoisResponse.Merge(m, src)
}
func (m *QueryAllWhoisResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllWhoisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllWhoisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllWhoisResponse proto.InternalMessageInfo

func (m *QueryAllWhoisResponse) GetWhois() []Whois {
	if m != nil {
		return m.Whois
	}
	return nil
}

func (m *QueryAllWhoisResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetWhoisRequest)(nil), "cosmonaut.nameservice.nameservice.QueryGetWhoisRequest")
	proto.RegisterType((*QueryGetWhoisResponse)(nil), "cosmonaut.nameservice.nameservice.QueryGetWhoisResponse")
	proto.RegisterType((*QueryAllWhoisRequest)(nil), "cosmonaut.nameservice.nameservice.QueryAllWhoisRequest")
	proto.RegisterType((*QueryAllWhoisResponse)(nil), "cosmonaut.nameservice.nameservice.QueryAllWhoisResponse")
}

func init() { proto.RegisterFile("nameservice/query.proto", fileDescriptor_37776ef2c2bc2f1b) }

var fileDescriptor_37776ef2c2bc2f1b = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x4b, 0xcc, 0x4d,
	0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0xcf, 0x4b, 0x2c, 0x2d, 0xd1, 0x43,
	0x52, 0x82, 0xcc, 0x96, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4,
	0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x18, 0x20, 0xa5, 0x05,
	0x36, 0xa0, 0x58, 0x3f, 0x29, 0xb1, 0x18, 0x6a, 0xb2, 0x7e, 0x99, 0x61, 0x52, 0x6a, 0x49, 0xa2,
	0xa1, 0x7e, 0x41, 0x62, 0x7a, 0x66, 0x1e, 0x58, 0x31, 0x54, 0x2d, 0x8a, 0x2b, 0xca, 0x33, 0xf2,
	0x33, 0x61, 0x86, 0x88, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20, 0x16, 0x44, 0x54, 0x49,
	0x87, 0x4b, 0x24, 0x10, 0x64, 0xa0, 0x7b, 0x6a, 0x49, 0x38, 0x48, 0x71, 0x50, 0x6a, 0x61, 0x69,
	0x6a, 0x71, 0x89, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x5e, 0x4a, 0x6a, 0x85, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x84, 0xa3, 0x14, 0xcb, 0x25, 0x8a, 0xa6, 0xba, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x55, 0xc8, 0x85, 0x8b, 0x15, 0x6c, 0x17, 0x58, 0x39, 0xb7, 0x91, 0x86, 0x1e, 0x41, 0x2f, 0xeb,
	0x81, 0x0d, 0x70, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0xa2, 0x59, 0x29, 0x0e, 0xea, 0x18,
	0xc7, 0x9c, 0x1c, 0x14, 0xc7, 0xb8, 0x71, 0x71, 0x21, 0xfc, 0x09, 0xb5, 0x42, 0x0d, 0x62, 0x45,
	0xb1, 0x1e, 0x28, 0x50, 0xf4, 0x20, 0xc1, 0x0d, 0x0d, 0x14, 0xbd, 0x80, 0xc4, 0xf4, 0x54, 0xa8,
	0xde, 0x20, 0x24, 0x9d, 0x4a, 0xcb, 0x18, 0xa1, 0xee, 0x47, 0x58, 0x80, 0xe9, 0x7e, 0x66, 0xb2,
	0xdd, 0x2f, 0xe4, 0x8e, 0xe2, 0x4e, 0x26, 0xb0, 0x3b, 0xd5, 0x09, 0xba, 0x13, 0xe2, 0x04, 0x64,
	0x87, 0x1a, 0x7d, 0x60, 0xe2, 0x62, 0x05, 0x3b, 0x54, 0x68, 0x0b, 0x23, 0x17, 0x2b, 0xd8, 0x26,
	0x21, 0x73, 0x22, 0xdc, 0x84, 0x2d, 0x2a, 0xa5, 0x2c, 0x48, 0xd7, 0x08, 0x71, 0x92, 0x92, 0x45,
	0xd3, 0xe5, 0x27, 0x93, 0x99, 0x8c, 0x84, 0x0c, 0xf4, 0xe1, 0x26, 0xe8, 0x23, 0x27, 0x2f, 0x8c,
	0xa4, 0xa6, 0x5f, 0x0d, 0x4e, 0x27, 0xb5, 0x42, 0xeb, 0x19, 0xb9, 0x38, 0xc0, 0x66, 0x39, 0xe6,
	0xe4, 0x10, 0xef, 0x72, 0xb4, 0x78, 0x27, 0xde, 0xe5, 0xe8, 0xf1, 0xa9, 0x64, 0x00, 0x76, 0xb9,
	0x96, 0x90, 0x06, 0xb1, 0x2e, 0x77, 0xf2, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0x93, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x1c, 0xa6, 0x55,
	0xa0, 0xf0, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xf9, 0xcb, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xcc, 0xe4, 0x4d, 0xc1, 0x16, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a whois by index.
	Whois(ctx context.Context, in *QueryGetWhoisRequest, opts ...grpc.CallOption) (*QueryGetWhoisResponse, error)
	// Queries a list of whois items.
	WhoisAll(ctx context.Context, in *QueryAllWhoisRequest, opts ...grpc.CallOption) (*QueryAllWhoisResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Whois(ctx context.Context, in *QueryGetWhoisRequest, opts ...grpc.CallOption) (*QueryGetWhoisResponse, error) {
	out := new(QueryGetWhoisResponse)
	err := c.cc.Invoke(ctx, "/cosmonaut.nameservice.nameservice.Query/Whois", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) WhoisAll(ctx context.Context, in *QueryAllWhoisRequest, opts ...grpc.CallOption) (*QueryAllWhoisResponse, error) {
	out := new(QueryAllWhoisResponse)
	err := c.cc.Invoke(ctx, "/cosmonaut.nameservice.nameservice.Query/WhoisAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a whois by index.
	Whois(context.Context, *QueryGetWhoisRequest) (*QueryGetWhoisResponse, error)
	// Queries a list of whois items.
	WhoisAll(context.Context, *QueryAllWhoisRequest) (*QueryAllWhoisResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Whois(ctx context.Context, req *QueryGetWhoisRequest) (*QueryGetWhoisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Whois not implemented")
}
func (*UnimplementedQueryServer) WhoisAll(ctx context.Context, req *QueryAllWhoisRequest) (*QueryAllWhoisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoisAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Whois_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetWhoisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Whois(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmonaut.nameservice.nameservice.Query/Whois",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Whois(ctx, req.(*QueryGetWhoisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_WhoisAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllWhoisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WhoisAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmonaut.nameservice.nameservice.Query/WhoisAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WhoisAll(ctx, req.(*QueryAllWhoisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmonaut.nameservice.nameservice.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Whois",
			Handler:    _Query_Whois_Handler,
		},
		{
			MethodName: "WhoisAll",
			Handler:    _Query_WhoisAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nameservice/query.proto",
}

func (m *QueryGetWhoisRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetWhoisRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetWhoisRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetWhoisResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetWhoisResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetWhoisResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Whois.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllWhoisRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllWhoisRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllWhoisRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllWhoisResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllWhoisResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllWhoisResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Whois) > 0 {
		for iNdEx := len(m.Whois) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Whois[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetWhoisRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetWhoisResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Whois.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllWhoisRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllWhoisResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Whois) > 0 {
		for _, e := range m.Whois {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetWhoisRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetWhoisRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetWhoisRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetWhoisResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetWhoisResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetWhoisResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Whois", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Whois.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllWhoisRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllWhoisRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllWhoisRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllWhoisResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllWhoisResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllWhoisResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Whois", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Whois = append(m.Whois, Whois{})
			if err := m.Whois[len(m.Whois)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
