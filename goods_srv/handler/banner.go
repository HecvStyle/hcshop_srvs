package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hcshop_srvs/goods_srv/global"
	"hcshop_srvs/goods_srv/model"
	"hcshop_srvs/goods_srv/proto"
)

func (s *GoodsServer) BannerList(ctx context.Context, req *emptypb.Empty) (*proto.BannerListResponse, error) {
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

func (s *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	banner := model.Banner{
		Image: req.Image,
		Url:   req.Url,
		Index: req.Index,
	}
	global.DB.Save(&banner)
	return &proto.BannerResponse{Id: banner.ID}, nil
}

func (s *GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {

	if result := global.DB.Delete(&model.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "轮播图不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	if result := global.DB.First(&model.Banner{}); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "轮播图不存在")
	}
	banner := model.Banner{}
	if req.Url != "" {
		banner.Url = req.Url
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	global.DB.Save(&banner)
	return &emptypb.Empty{}, nil
}
