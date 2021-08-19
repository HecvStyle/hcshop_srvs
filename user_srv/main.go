package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"hcshop_srvs/user_srv/handler"
	"hcshop_srvs/user_srv/initialize"
	"hcshop_srvs/user_srv/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// 配置文件初始化
	initialize.InitConfig()

	// 日志文件初始化
	initialize.InitLogger()

	// 初始化数据里链接
	initialize.InitDB()

	IP := flag.String("ip", "0.0.0.0", "ip 地址")
	Port := flag.Int("port", 50051, "端口号")
	flag.Parse()

	zap.S().Info("准备启动服务...")
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
	zap.S().Info("启动服务完毕")
}
