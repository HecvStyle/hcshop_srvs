package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"hcshop_srvs/user_srv/global"
	"hcshop_srvs/user_srv/handler"
	"hcshop_srvs/user_srv/initialize"
	"hcshop_srvs/user_srv/proto"
	"hcshop_srvs/user_srv/utils"
	"net"
	"os"
	"os/signal"
	"syscall"
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

	// 如果没有通过命令行参数传递端口进来，则动态生成一个端口来使用
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}

	zap.S().Info("准备启动服务...")
	zap.S().Info("使用的端口", *Port)
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
		GRPC:                           fmt.Sprintf("192.168.1.75:%d", *Port),
		Interval:                       "5s",
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID, _ = uuid.GenerateUUID()
	registration.Port = *Port
	registration.Tags = []string{"user-srv"}
	// 这里别瞎鸡毛加 "http://" scheme，这可是rpc 协议，踩坑 +1
	registration.Address = fmt.Sprintf("192.168.1.75")
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()
	zap.S().Info("服务启动完成")
	quit := make(chan os.Signal)

	// 这里是没有监听到 goland 停止debug/build 的信号，所以停止时候，并不立刻注销服务
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(registration.ID); err != nil {
		zap.S().Info("服务注销失败")
		panic(err)
	}
	zap.S().Info("注销成功")
}
