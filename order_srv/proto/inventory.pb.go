// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GoodsInvInfo struct {
	GoodsId              int32    `protobuf:"varint,1,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsInvInfo) Reset()         { *m = GoodsInvInfo{} }
func (m *GoodsInvInfo) String() string { return proto.CompactTextString(m) }
func (*GoodsInvInfo) ProtoMessage()    {}
func (*GoodsInvInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7173caedb7c6ae96, []int{0}
}

func (m *GoodsInvInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsInvInfo.Unmarshal(m, b)
}
func (m *GoodsInvInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsInvInfo.Marshal(b, m, deterministic)
}
func (m *GoodsInvInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsInvInfo.Merge(m, src)
}
func (m *GoodsInvInfo) XXX_Size() int {
	return xxx_messageInfo_GoodsInvInfo.Size(m)
}
func (m *GoodsInvInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsInvInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsInvInfo proto.InternalMessageInfo

func (m *GoodsInvInfo) GetGoodsId() int32 {
	if m != nil {
		return m.GoodsId
	}
	return 0
}

func (m *GoodsInvInfo) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type SellInfo struct {
	GoodsInfo            []*GoodsInvInfo `protobuf:"bytes,1,rep,name=goodsInfo,proto3" json:"goodsInfo,omitempty"`
	OrderSn              string          `protobuf:"bytes,2,opt,name=orderSn,proto3" json:"orderSn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SellInfo) Reset()         { *m = SellInfo{} }
func (m *SellInfo) String() string { return proto.CompactTextString(m) }
func (*SellInfo) ProtoMessage()    {}
func (*SellInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7173caedb7c6ae96, []int{1}
}

func (m *SellInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfo.Unmarshal(m, b)
}
func (m *SellInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfo.Marshal(b, m, deterministic)
}
func (m *SellInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfo.Merge(m, src)
}
func (m *SellInfo) XXX_Size() int {
	return xxx_messageInfo_SellInfo.Size(m)
}
func (m *SellInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfo proto.InternalMessageInfo

func (m *SellInfo) GetGoodsInfo() []*GoodsInvInfo {
	if m != nil {
		return m.GoodsInfo
	}
	return nil
}

func (m *SellInfo) GetOrderSn() string {
	if m != nil {
		return m.OrderSn
	}
	return ""
}

func init() {
	proto.RegisterType((*GoodsInvInfo)(nil), "GoodsInvInfo")
	proto.RegisterType((*SellInfo)(nil), "SellInfo")
}

func init() { proto.RegisterFile("inventory.proto", fileDescriptor_7173caedb7c6ae96) }

var fileDescriptor_7173caedb7c6ae96 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xcc, 0x2b, 0x4b,
	0xcd, 0x2b, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4e, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0xa0, 0x92,
	0x4a, 0x56, 0x5c, 0x3c, 0xee, 0xf9, 0xf9, 0x29, 0xc5, 0x9e, 0x79, 0x65, 0x9e, 0x79, 0x69, 0xf9,
	0x42, 0x12, 0x5c, 0xec, 0xe9, 0x60, 0x7e, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x8c,
	0x2b, 0x24, 0xc0, 0xc5, 0x9c, 0x57, 0x9a, 0x2b, 0xc1, 0x04, 0x16, 0x05, 0x31, 0x95, 0x02, 0xb9,
	0x38, 0x82, 0x53, 0x73, 0x72, 0xc0, 0xfa, 0xb4, 0xb9, 0x38, 0x21, 0x0a, 0xf3, 0xd2, 0xf2, 0x25,
	0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x78, 0xf5, 0x90, 0x4d, 0x0e, 0x42, 0xc8, 0x83, 0x2c, 0xc9,
	0x2f, 0x4a, 0x49, 0x2d, 0x0a, 0xce, 0x03, 0x1b, 0xc7, 0x19, 0x04, 0xe3, 0x1a, 0xed, 0x67, 0xe4,
	0xe2, 0xf4, 0x84, 0xb9, 0x5f, 0x48, 0x9f, 0x8b, 0x2d, 0x38, 0xb5, 0xc4, 0x33, 0xaf, 0x4c, 0x08,
	0xd5, 0x2c, 0x29, 0x31, 0x3d, 0x88, 0x9f, 0xf4, 0x60, 0x7e, 0xd2, 0x73, 0x05, 0xf9, 0x49, 0x48,
	0x13, 0xac, 0xdb, 0x25, 0xb5, 0x24, 0x31, 0x33, 0x07, 0x5d, 0x0f, 0x2a, 0x57, 0x48, 0x93, 0x8b,
	0x05, 0xe4, 0x78, 0x21, 0x4e, 0x3d, 0x98, 0x1f, 0x70, 0x9a, 0xaa, 0xcd, 0xc5, 0x16, 0x94, 0x9a,
	0x94, 0x98, 0x9c, 0x4d, 0x84, 0x62, 0x27, 0xce, 0x28, 0x76, 0x3d, 0x6b, 0x88, 0x18, 0x1b, 0x98,
	0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x35, 0x84, 0x2f, 0xdb, 0x92, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InventoryClient interface {
	SetInv(ctx context.Context, in *GoodsInvInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	InvDetail(ctx context.Context, in *GoodsInvInfo, opts ...grpc.CallOption) (*GoodsInvInfo, error)
	Sell(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Reback(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type inventoryClient struct {
	cc *grpc.ClientConn
}

func NewInventoryClient(cc *grpc.ClientConn) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) SetInv(ctx context.Context, in *GoodsInvInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Inventory/SetInv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) InvDetail(ctx context.Context, in *GoodsInvInfo, opts ...grpc.CallOption) (*GoodsInvInfo, error) {
	out := new(GoodsInvInfo)
	err := c.cc.Invoke(ctx, "/Inventory/InvDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) Sell(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Inventory/Sell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) Reback(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Inventory/Reback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
type InventoryServer interface {
	SetInv(context.Context, *GoodsInvInfo) (*emptypb.Empty, error)
	InvDetail(context.Context, *GoodsInvInfo) (*GoodsInvInfo, error)
	Sell(context.Context, *SellInfo) (*emptypb.Empty, error)
	Reback(context.Context, *SellInfo) (*emptypb.Empty, error)
}

// UnimplementedInventoryServer can be embedded to have forward compatible implementations.
type UnimplementedInventoryServer struct {
}

func (*UnimplementedInventoryServer) SetInv(ctx context.Context, req *GoodsInvInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInv not implemented")
}
func (*UnimplementedInventoryServer) InvDetail(ctx context.Context, req *GoodsInvInfo) (*GoodsInvInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvDetail not implemented")
}
func (*UnimplementedInventoryServer) Sell(ctx context.Context, req *SellInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sell not implemented")
}
func (*UnimplementedInventoryServer) Reback(ctx context.Context, req *SellInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reback not implemented")
}

func RegisterInventoryServer(s *grpc.Server, srv InventoryServer) {
	s.RegisterService(&_Inventory_serviceDesc, srv)
}

func _Inventory_SetInv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsInvInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).SetInv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/SetInv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).SetInv(ctx, req.(*GoodsInvInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_InvDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsInvInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).InvDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/InvDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).InvDetail(ctx, req.(*GoodsInvInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_Sell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).Sell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/Sell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).Sell(ctx, req.(*SellInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_Reback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).Reback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/Reback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).Reback(ctx, req.(*SellInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Inventory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetInv",
			Handler:    _Inventory_SetInv_Handler,
		},
		{
			MethodName: "InvDetail",
			Handler:    _Inventory_InvDetail_Handler,
		},
		{
			MethodName: "Sell",
			Handler:    _Inventory_Sell_Handler,
		},
		{
			MethodName: "Reback",
			Handler:    _Inventory_Reback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}
