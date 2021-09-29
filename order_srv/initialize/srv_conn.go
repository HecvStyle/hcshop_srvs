package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"hcshop_srvs/order_srv/global"
	"hcshop_srvs/order_srv/proto"
)

func InitGoodsSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	conn, err := grpc.Dial(
		// 记得添加 consul:// 协议，别搞错了 踩坑+1
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("商品服务注册连接失败")
	}
	// 这里涉及到了多个链接，都只用了一个gorutine,考虑使用连接池才可以
	global.GoodsSrvClient = proto.NewGoodsClient(conn)
}

func InitInventorySrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	conn, err := grpc.Dial(
		// 记得添加 consul:// 协议，别搞错了 踩坑+1
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.InventorySrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("库存服务注册连接失败")
	}
	// 这里涉及到了多个链接，都只用了一个gorutine,考虑使用连接池才可以
	global.GoodsSrvClient = proto.NewGoodsClient(conn)
}
