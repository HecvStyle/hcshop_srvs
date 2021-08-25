package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"hcshop_srvs/goods_srv/global"
	"hcshop_srvs/goods_srv/model"
	"hcshop_srvs/goods_srv/proto"
)

//轮播图
func (g *GoodsServer) BannerList(ctx context.Context, req *emptypb.Empty) (*proto.BannerListResponse, error) {
	resp := proto.BannerListResponse{}
	var banners []model.Banner
	result := global.DB.Find(&banners)
	resp.Total = int32(result.RowsAffected)
	for _, banner := range banners {
		resp.Data = append(resp.Data, &proto.BannerResponse{
			Id:    banner.ID,
			Index: banner.Index,
			Image: banner.Image,
			Url:   banner.Url,
		})
	}
	return &resp, nil
}

//func (g *GoodsServer) CreateBanner(context.Context, *proto.BannerRequest) (*proto.BannerResponse, error) {
//
//}
//func (g *GoodsServer) DeleteBanner(context.Context, *proto.BannerRequest) (*emptypb.Empty, error) {
//
//}
//func (g *GoodsServer) UpdateBanner(context.Context, *proto.BannerRequest) (*emptypb.Empty, error) {
//
//}
