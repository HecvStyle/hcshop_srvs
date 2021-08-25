package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"hcshop_srvs/goods_srv/proto"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func TestGetBrandList() {
	rsp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
	}
}

func Init() {
	var err error
	conn, err = grpc.Dial("192.168.1.175:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	brandClient = proto.NewGoodsClient(conn)
}

func main() {
	Init()
	//TestGetBrandList()

	AddBrands()

	conn.Close()
}

func AddBrands() {
	for i := 0; i < 10; i++ {
		brandClient.CreateBrand(context.Background(), &proto.BrandRequest{
			Name: fmt.Sprintf("小米%d", i),
			Logo: "https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Xiaomi_logo_%282021-%29.svg/512px-Xiaomi_logo_%282021-%29.svg.png",
		})
	}
}
