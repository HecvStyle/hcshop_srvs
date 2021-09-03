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

// CategoryBrandList 所有品牌
func (s *GoodsServer) CategoryBrandList(ctx context.Context, req *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	var categoryBrands []model.GoodsCategoryBrand
	categoryBrandListResp := proto.CategoryBrandListResponse{}

	var total int64
	global.DB.Model(&model.GoodsCategoryBrand{}).Count(&total)
	categoryBrandListResp.Total = int32(total)
	global.DB.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&categoryBrands)

	var categoryBrandsResponses []*proto.CategoryBrandResponse
	for _, categoryBrand := range categoryBrands {
		categoryBrandsResponses = append(categoryBrandsResponses, &proto.CategoryBrandResponse{
			Brand: &proto.BrandInfoResponse{
				Id:   categoryBrand.Brands.ID,
				Name: categoryBrand.Brands.Name,
				Logo: categoryBrand.Brands.Logo,
			},
			Category: &proto.CategoryInfoResponse{
				Id:             categoryBrand.Category.ID,
				Name:           categoryBrand.Category.Name,
				ParentCategory: categoryBrand.Category.ParentCategoryID,
				Level:          categoryBrand.Category.Level,
				IsTab:          categoryBrand.Category.IsTab,
			},
		})
	}
	categoryBrandListResp.Data = categoryBrandsResponses
	return &categoryBrandListResp, nil

}

// GetCategoryBrandList 通过category获取brands
func (s *GoodsServer) GetCategoryBrandList(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	brandListResp := proto.BrandListResponse{}
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	var categoryBrands []model.GoodsCategoryBrand
	if result := global.DB.Preload("Brands").Where(&model.Goods{CategoryID: req.Id}).Find(&categoryBrands); result.RowsAffected > 0 {
		brandListResp.Total = int32(result.RowsAffected)
	}

	var brandInfoResponses []*proto.BrandInfoResponse

	for _, categoryBrand := range categoryBrands {
		brandInfoResponses = append(brandInfoResponses, &proto.BrandInfoResponse{
			Id:   categoryBrand.Brands.ID,
			Name: categoryBrand.Brands.Name,
			Logo: categoryBrand.Brands.Logo,
		})
	}
	brandListResp.Data = brandInfoResponses
	return &brandListResp, nil

}

func (s *GoodsServer) CreateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "分类不存在")
	}

	var brands model.Brands
	if result := global.DB.First(&brands, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}

	categoryBrands := model.GoodsCategoryBrand{BrandsID: brands.ID, CategoryID: category.ID}
	global.DB.Save(&categoryBrands)
	return &proto.CategoryBrandResponse{Id: categoryBrands.BrandsID}, nil

}

func (s *GoodsServer) DeleteCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	var categoryBrands model.GoodsCategoryBrand
	if result := global.DB.Delete(&categoryBrands, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*emptypb.Empty, error) {
	var categoryBrands model.GoodsCategoryBrand
	if result := global.DB.First(&categoryBrands, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
	}

	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "分类不存在")
	}

	var brands model.Brands
	if result := global.DB.First(&brands, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}

	categoryBrands.CategoryID = category.ID
	categoryBrands.BrandsID = brands.ID

	global.DB.Save(&categoryBrands)
	return &emptypb.Empty{}, nil
}
