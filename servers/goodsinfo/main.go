package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"goodsinfo/handler"
	pb "goodsinfo/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "goodsinfo"
	version = "latest"
)

func main() {
	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService(
		micro.Address("127.0.0.1:9002"), //指定微服务的ip:  选择注册服务器地址,也可以不配置,默认为本机,也可以选择consul集群中的client
		micro.Name(service),
		micro.Version(version),
		//注册consul
		micro.Registry(consulReg),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterGoodsinfoHandler(srv.Server(), new(handler.Goodsinfo)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
