// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shop.proto

package shop

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetShopInfoReq struct {
	ShopID               int64    `protobuf:"varint,1,opt,name=shopID,proto3" json:"shopID,omitempty"`
	UserID               int64    `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShopInfoReq) Reset()         { *m = GetShopInfoReq{} }
func (m *GetShopInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetShopInfoReq) ProtoMessage()    {}
func (*GetShopInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{0}
}
func (m *GetShopInfoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetShopInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetShopInfoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetShopInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShopInfoReq.Merge(m, src)
}
func (m *GetShopInfoReq) XXX_Size() int {
	return m.Size()
}
func (m *GetShopInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShopInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetShopInfoReq proto.InternalMessageInfo

func (m *GetShopInfoReq) GetShopID() int64 {
	if m != nil {
		return m.ShopID
	}
	return 0
}

func (m *GetShopInfoReq) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type GetShopInfoRes struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID               int64    `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=Address,proto3" json:"Address,omitempty"`
	Status               int32    `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShopInfoRes) Reset()         { *m = GetShopInfoRes{} }
func (m *GetShopInfoRes) String() string { return proto.CompactTextString(m) }
func (*GetShopInfoRes) ProtoMessage()    {}
func (*GetShopInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{1}
}
func (m *GetShopInfoRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetShopInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetShopInfoRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetShopInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShopInfoRes.Merge(m, src)
}
func (m *GetShopInfoRes) XXX_Size() int {
	return m.Size()
}
func (m *GetShopInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShopInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetShopInfoRes proto.InternalMessageInfo

func (m *GetShopInfoRes) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetShopInfoRes) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GetShopInfoRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetShopInfoRes) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GetShopInfoRes) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type RegisterShopReq struct {
	UserID               int64    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterShopReq) Reset()         { *m = RegisterShopReq{} }
func (m *RegisterShopReq) String() string { return proto.CompactTextString(m) }
func (*RegisterShopReq) ProtoMessage()    {}
func (*RegisterShopReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{2}
}
func (m *RegisterShopReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterShopReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterShopReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterShopReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterShopReq.Merge(m, src)
}
func (m *RegisterShopReq) XXX_Size() int {
	return m.Size()
}
func (m *RegisterShopReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterShopReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterShopReq proto.InternalMessageInfo

func (m *RegisterShopReq) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *RegisterShopReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterShopReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type RegisterShopResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterShopResp) Reset()         { *m = RegisterShopResp{} }
func (m *RegisterShopResp) String() string { return proto.CompactTextString(m) }
func (*RegisterShopResp) ProtoMessage()    {}
func (*RegisterShopResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f3030369b20fd61, []int{3}
}
func (m *RegisterShopResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterShopResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterShopResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterShopResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterShopResp.Merge(m, src)
}
func (m *RegisterShopResp) XXX_Size() int {
	return m.Size()
}
func (m *RegisterShopResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterShopResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterShopResp proto.InternalMessageInfo

func (m *RegisterShopResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*GetShopInfoReq)(nil), "shop.GetShopInfoReq")
	proto.RegisterType((*GetShopInfoRes)(nil), "shop.GetShopInfoRes")
	proto.RegisterType((*RegisterShopReq)(nil), "shop.RegisterShopReq")
	proto.RegisterType((*RegisterShopResp)(nil), "shop.RegisterShopResp")
}

func init() { proto.RegisterFile("shop.proto", fileDescriptor_0f3030369b20fd61) }

var fileDescriptor_0f3030369b20fd61 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xed, 0x96, 0x82, 0x32, 0x1a, 0x24, 0x1b, 0x25, 0x4d, 0x0f, 0xb5, 0xe9, 0x89, 0x83, 0x40,
	0xa2, 0x89, 0x17, 0x0f, 0x2a, 0x21, 0x31, 0x5c, 0x3c, 0x94, 0x8b, 0xd7, 0x42, 0x57, 0xd8, 0x28,
	0x6e, 0xed, 0x6c, 0xe5, 0xe6, 0xc9, 0x8f, 0xf0, 0x93, 0x3c, 0xfa, 0x07, 0x9a, 0xfa, 0x23, 0x66,
	0xb7, 0x05, 0x29, 0xc1, 0xdb, 0xbc, 0xb7, 0xf3, 0xe6, 0xbd, 0x99, 0x16, 0x00, 0x67, 0x22, 0xee,
	0xc6, 0x89, 0x90, 0x82, 0x5a, 0xaa, 0x76, 0xce, 0xa7, 0x5c, 0xce, 0xd2, 0x71, 0x77, 0x22, 0xe6,
	0xbd, 0xf9, 0x82, 0xcb, 0x07, 0xb1, 0xe8, 0x4d, 0x45, 0x47, 0xb7, 0x74, 0x5e, 0xc2, 0x47, 0x1e,
	0x85, 0x52, 0x24, 0xd8, 0x5b, 0x95, 0xb9, 0xda, 0xbf, 0x82, 0xc6, 0x0d, 0x93, 0xa3, 0x99, 0x88,
	0x87, 0x4f, 0xf7, 0x22, 0x60, 0xcf, 0xb4, 0x05, 0x35, 0x35, 0x71, 0x38, 0xb0, 0x89, 0x47, 0xda,
	0x95, 0xa0, 0x40, 0x8a, 0x4f, 0x91, 0x25, 0xc3, 0x81, 0x6d, 0xe6, 0x7c, 0x8e, 0xfc, 0xd7, 0x8d,
	0x09, 0x48, 0x1b, 0x60, 0xf2, 0xa8, 0x50, 0x9b, 0x3c, 0xfa, 0x4f, 0x49, 0x29, 0x58, 0xb7, 0xe1,
	0x9c, 0xd9, 0x15, 0x8f, 0xb4, 0xeb, 0x81, 0xae, 0xa9, 0x0d, 0x3b, 0xd7, 0x51, 0x94, 0x30, 0x44,
	0xdb, 0xd2, 0xf4, 0x12, 0xaa, 0x29, 0x23, 0x19, 0xca, 0x14, 0xed, 0xaa, 0x47, 0xda, 0xd5, 0xa0,
	0x40, 0xbe, 0x80, 0x83, 0x80, 0x4d, 0x39, 0x4a, 0x96, 0xa8, 0x10, 0x6a, 0x05, 0x77, 0x65, 0xa8,
	0x43, 0xf4, 0x6b, 0xd9, 0xd7, 0xb1, 0xd9, 0x34, 0x56, 0xc6, 0x4e, 0x61, 0xac, 0xe2, 0xd4, 0xf3,
	0xd7, 0x3b, 0x52, 0x04, 0xf0, 0xfe, 0x02, 0x54, 0x4a, 0xcf, 0x4b, 0xda, 0x3f, 0x81, 0x66, 0xd9,
	0x10, 0x63, 0x15, 0x1b, 0xd3, 0xc9, 0x44, 0xa9, 0x94, 0xe5, 0x6e, 0xb0, 0x84, 0xa7, 0x6f, 0x04,
	0x2c, 0xd5, 0x46, 0x2f, 0x60, 0x6f, 0xed, 0x4e, 0xf4, 0xb0, 0xab, 0xbf, 0x61, 0xf9, 0xf8, 0xce,
	0x36, 0x16, 0x7d, 0x83, 0x5e, 0xc2, 0xfe, 0xba, 0x27, 0x3d, 0xca, 0xfb, 0x36, 0x16, 0x77, 0x5a,
	0xdb, 0x68, 0x8c, 0x7d, 0xa3, 0xdf, 0xfc, 0xc8, 0x5c, 0xf2, 0x99, 0xb9, 0xe4, 0x3b, 0x73, 0xc9,
	0xfb, 0x8f, 0x6b, 0x8c, 0x6b, 0xfa, 0x07, 0x38, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x16, 0xfb,
	0xd8, 0x2d, 0x4c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShopClient is the client API for Shop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShopClient interface {
	GetShopInfo(ctx context.Context, in *GetShopInfoReq, opts ...grpc.CallOption) (*GetShopInfoRes, error)
	RegisterShop(ctx context.Context, in *RegisterShopReq, opts ...grpc.CallOption) (*RegisterShopResp, error)
}

type shopClient struct {
	cc *grpc.ClientConn
}

func NewShopClient(cc *grpc.ClientConn) ShopClient {
	return &shopClient{cc}
}

func (c *shopClient) GetShopInfo(ctx context.Context, in *GetShopInfoReq, opts ...grpc.CallOption) (*GetShopInfoRes, error) {
	out := new(GetShopInfoRes)
	err := c.cc.Invoke(ctx, "/shop.Shop/GetShopInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) RegisterShop(ctx context.Context, in *RegisterShopReq, opts ...grpc.CallOption) (*RegisterShopResp, error) {
	out := new(RegisterShopResp)
	err := c.cc.Invoke(ctx, "/shop.Shop/RegisterShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServer is the server API for Shop service.
type ShopServer interface {
	GetShopInfo(context.Context, *GetShopInfoReq) (*GetShopInfoRes, error)
	RegisterShop(context.Context, *RegisterShopReq) (*RegisterShopResp, error)
}

// UnimplementedShopServer can be embedded to have forward compatible implementations.
type UnimplementedShopServer struct {
}

func (*UnimplementedShopServer) GetShopInfo(ctx context.Context, req *GetShopInfoReq) (*GetShopInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopInfo not implemented")
}
func (*UnimplementedShopServer) RegisterShop(ctx context.Context, req *RegisterShopReq) (*RegisterShopResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterShop not implemented")
}

func RegisterShopServer(s *grpc.Server, srv ShopServer) {
	s.RegisterService(&_Shop_serviceDesc, srv)
}

func _Shop_GetShopInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).GetShopInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.Shop/GetShopInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).GetShopInfo(ctx, req.(*GetShopInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_RegisterShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterShopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).RegisterShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.Shop/RegisterShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).RegisterShop(ctx, req.(*RegisterShopReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shop_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shop.Shop",
	HandlerType: (*ShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShopInfo",
			Handler:    _Shop_GetShopInfo_Handler,
		},
		{
			MethodName: "RegisterShop",
			Handler:    _Shop_RegisterShop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

func (m *GetShopInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetShopInfoReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetShopInfoReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UserID != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.UserID))
		i--
		dAtA[i] = 0x10
	}
	if m.ShopID != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.ShopID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetShopInfoRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetShopInfoRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetShopInfoRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.UserID != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.UserID))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RegisterShopReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterShopReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterShopReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintShop(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.UserID != 0 {
		i = encodeVarintShop(dAtA, i, uint64(m.UserID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RegisterShopResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterShopResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterShopResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintShop(dAtA []byte, offset int, v uint64) int {
	offset -= sovShop(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetShopInfoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ShopID != 0 {
		n += 1 + sovShop(uint64(m.ShopID))
	}
	if m.UserID != 0 {
		n += 1 + sovShop(uint64(m.UserID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetShopInfoRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovShop(uint64(m.Id))
	}
	if m.UserID != 0 {
		n += 1 + sovShop(uint64(m.UserID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovShop(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegisterShopReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserID != 0 {
		n += 1 + sovShop(uint64(m.UserID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovShop(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegisterShopResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovShop(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShop(x uint64) (n int) {
	return sovShop(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetShopInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShop
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
			return fmt.Errorf("proto: GetShopInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetShopInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShopID", wireType)
			}
			m.ShopID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShopID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShop
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShop
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
func (m *GetShopInfoRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShop
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
			return fmt.Errorf("proto: GetShopInfoRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetShopInfoRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
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
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
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
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShop
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShop
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
func (m *RegisterShopReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShop
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
			return fmt.Errorf("proto: RegisterShopReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterShopReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
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
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
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
				return ErrInvalidLengthShop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShop
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShop
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
func (m *RegisterShopResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShop
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
			return fmt.Errorf("proto: RegisterShopResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterShopResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipShop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShop
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShop
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
func skipShop(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShop
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
					return 0, ErrIntOverflowShop
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
					return 0, ErrIntOverflowShop
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
				return 0, ErrInvalidLengthShop
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthShop
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowShop
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
				next, err := skipShop(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthShop
				}
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
	ErrInvalidLengthShop = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShop   = fmt.Errorf("proto: integer overflow")
)