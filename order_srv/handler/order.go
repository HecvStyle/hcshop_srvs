package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"hcshop_srvs/order_srv/proto"
)

type OrderServer struct {
	proto.UnimplementedOrderServer
}

func (o *OrderServer) CartItemList(ctx context.Context, req *proto.UserInfo ) (*proto.CartItemListResponse, error) {
	panic("implement me")
}

func (o *OrderServer) CreateCartItem(ctx context.Context, in *proto.CartItemRequest) (*proto.ShopCartInfoResponse, error) {
	panic("implement me")
}

func (o *OrderServer) UpdateCartItem(ctx context.Context, in *proto.CartItemRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (o *OrderServer) DeleteCartItem(ctx context.Context, in *proto.CartItemRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (o *OrderServer) CreateOrder(ctx context.Context, in *proto.OrderRequest) (*proto.OrderInfoResponse, error) {
	panic("implement me")
}

func (o *OrderServer) OrderList(ctx context.Context, in *proto.OrderFilterRequest) (*proto.OrderListResponse, error) {
	panic("implement me")
}

func (o *OrderServer) OrderDetail(ctx context.Context, in *proto.OrderRequest) (*proto.OrderInfoDetailResponse, error) {
	panic("implement me")
}

func (o *OrderServer) UpdateOrderStatus(ctx context.Context, in *proto.OrderStatus) (*emptypb.Empty, error) {
	panic("implement me")
}
