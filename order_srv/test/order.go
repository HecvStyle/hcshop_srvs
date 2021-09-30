package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"hcshop_srvs/order_srv/proto"
)

var OrderClient proto.OrderClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err := grpc.Dial("127.0.0.1:50054", grpc.WithInsecure())
	if err != nil {
		return
	}
	OrderClient = proto.NewOrderClient(conn)
}

func TestCreateCartItem(nums int32, userId int32, goodsId int32) {
	resp, err := OrderClient.CreateCartItem(context.Background(), &proto.CartItemRequest{
		Nums:    nums,
		UserId:  userId,
		GoodsId: goodsId,
	})
	if err != nil {
		return
	}
	fmt.Println(resp.Id)
}

func TestCartItemList(userId int32) {
	resp, err := OrderClient.CartItemList(context.Background(), &proto.UserInfo{Id: userId})
	if err != nil {
		return
	}
	fmt.Println(resp.Data)
}

func TestCreateOrder(userId int32) {
	resp, err := OrderClient.CreateOrder(context.Background(), &proto.OrderRequest{

		UserId:  userId,
		Address: "湖南",
		Name:    "callme",
		Mobile:  "17707411111",
		Post:    "快点发货",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}


func TestOrderDetail(orderId int32) {
	resp, err := OrderClient.OrderDetail(context.Background(), &proto.OrderRequest{
		Id: orderId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}

func main() {
	Init()

	//TestCreateCartItem(1,21,3)

	//TestCartItemList(21)

	//TestCreateOrder(21)

	//TestOrderDetail(1)
}


