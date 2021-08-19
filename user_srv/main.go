package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"hcshop_srvs/user_srv/global"
	"hcshop_srvs/user_srv/handler"
	"hcshop_srvs/user_srv/initialize"
	"hcshop_srvs/user_srv/proto"
	"net"
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

	// 注册健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("192.168.1.75:50051"),
		Interval:                       "5s",
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = global.ServerConfig.Name
	registration.Port = *Port
	registration.Tags = []string{"user-srv"}
	// 这里别瞎鸡毛加 "http://" scheme，这可是rpc 协议，踩坑 +1
	registration.Address = fmt.Sprintf("192.168.1.75")
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
	zap.S().Info("启动服务完毕")
}
