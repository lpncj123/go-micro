package handler

import (
	"context"
	"go-micro.dev/v4/logger"
	pb "goodsinfo/proto"
)

type Goodsinfo struct{}

func (e *Goodsinfo) AddGoods(ctx context.Context, req *pb.AddRequest, rsp *pb.AddResponse) error {
	logger.Infof("request: %v", req)
	//书写返回的逻辑结果
	rsp.Message = "增加成功"
	rsp.Success = true
	return nil
}
