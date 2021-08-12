package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"hcshop_srvs/user_srv/global"
	"hcshop_srvs/user_srv/model"
	"hcshop_srvs/user_srv/proto"
)

type UserServer struct{}

func ModelToResponse(user model.User) proto.UserResponse {
	// grpc 的message中，字段不能随便赋值nil，容易出错
	userResp := proto.UserResponse{
		Id:       user.ID,
		Password: user.Password,
		Nickname: user.NickName,
		Mobile:   user.Mobile,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userResp.Birthday = uint64(user.Birthday.Unix())
	}
	return userResp
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (u *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)
	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
	for _, user := range users {
		userResp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userResp)
	}
	return rsp, nil
}

//CreateUser(context.Context, *CreateUserInfo) (*UserResponse, error)
//UpdateUser(context.Context, *UpdateUserInfo) (*emptypb.Empty, error)
//CheckPassword(context.Context, *PasswordCheckInfo) (*CheckResponse, error)

func (u *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有找到用户")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userResp := ModelToResponse(user)
	return &userResp, nil
}

func (u *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserResponse, error) {
	var user model.User
	// 主键查询用户
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有找到用户")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userResp := ModelToResponse(user)
	return &userResp, nil
}

func (u *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserResponse, error) {
	// 新建用户
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	return nil, nil
}
