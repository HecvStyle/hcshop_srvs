package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hcshop_srvs/user_srv/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func main() {
	Init()
	//TestGetUserList()
	//TestCreateUser()
	conn.Close()
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{Pn: 1, PSize: 10})
	if err != nil {
		panic(err)
	}

	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.Nickname, user.Password)
		checkResp, err := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:          "admin123",
			EncryptedPassword: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkResp.GetSuccess())
	}
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			Nickname: fmt.Sprintf("hecv%d", i),
			Mobile:   fmt.Sprintf("1770712345%d", i),
			Password: "admin123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.GetId())
	}
}
