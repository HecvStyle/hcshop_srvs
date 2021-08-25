package handler

import (
	"hcshop_srvs/goods_srv/proto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

//GoodsList(context.Context, *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error){
//
//}
////现在用户提交订单有多个商品，你得批量查询商品的信息吧
//BatchGetGoods(context.Context, *BatchGoodsIdInfo) (*GoodsListResponse, error)
//CreateGoods(context.Context, *CreateGoodsInfo) (*GoodsInfoResponse, error)
//DeleteGoods(context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error)
//UpdateGoods(context.Context, *CreateGoodsInfo) (*emptypb.Empty, error)
//GetGoodsDetail(context.Context, *GoodInfoRequest) (*GoodsInfoResponse, error)
