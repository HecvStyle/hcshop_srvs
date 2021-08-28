package handler

import (
	"context"
	"hcshop_srvs/goods_srv/global"
	"hcshop_srvs/goods_srv/model"
	"hcshop_srvs/goods_srv/proto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func ModelToResponse(goods model.Goods) proto.GoodsInfoResponse {
	return proto.GoodsInfoResponse{
		Id:              goods.ID,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsBrief:      goods.GoodsBrief,
		ShipFree:        goods.ShipFree,
		Images:          goods.Images,
		DescImages:      goods.DescImages,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		Category: &proto.CategoryBriefInfoResponse{
			Id:   goods.CategoryID,
			Name: goods.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   goods.BrandsID,
			Name: goods.Brands.Name,
		},
	}
}

//GoodsList(context.Context, *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error){
//
//}

////现在用户提交订单有多个商品，你得批量查询商品的信息吧
func (s *GoodsServer) BatchGetGoods(ctx context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	goodsListResp := proto.GoodsListResponse{}
	var goodsList []model.Goods
	result := global.DB.Where(req.Id).Find(&goodsList)

	for _, goods := range goodsList {
		goodsInfo := ModelToResponse(goods)
		goodsListResp.Data = append(goodsListResp.Data, &goodsInfo)
	}
	goodsListResp.Total = int32(result.RowsAffected)
	return &goodsListResp, nil
}

//func (s *GoodsServer) CreateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
//
//}

//DeleteGoods(context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error)
//UpdateGoods(context.Context, *CreateGoodsInfo) (*emptypb.Empty, error)
//GetGoodsDetail(context.Context, *GoodInfoRequest) (*GoodsInfoResponse, error)
