package main

import (
	"context"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"time"

	pb "goodsinfo-client/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "goodsinfo"
	version = "latest"
)

func main() {
	consulReg := consul.NewRegistry(
		//指定微服务的ip:  选择注册服务器地址,默认为本机,也可以选择consul集群中的client
		registry.Addrs("127.0.0.1:8500"),
	)
	// Create service

	srv := micro.NewService(
		micro.Registry(consulReg))

	srv.Init()

	// Create client
	c := pb.NewGoodsinfoService(service, srv.Client())

	for {
		// Call service
		rsp, err := c.AddGoods(context.Background(), &pb.AddRequest{
			Title:   "我是一个商品",
			Price:   "20.22",
			Content: "内容展示",
		})
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info(rsp)

		time.Sleep(1 * time.Second)
	}
}
