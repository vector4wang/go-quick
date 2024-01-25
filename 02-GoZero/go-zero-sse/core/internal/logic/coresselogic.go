package logic

import (
	"context"
	"fmt"
	"gopkg.in/antage/eventsource.v1"
	"log"
	"time"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreSseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreSseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreSseLogic {
	return &CoreSseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreSseLogic) CoreSse() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	es := eventsource.New(nil, nil)
	defer es.Close()
	for {
		// 只设置发送数据，不添加事件名
		es.SendEventMessage(fmt.Sprintf("send data: %s", time.Now().Format("2006-01-02 15:04:05")), "", "")
		log.Printf("客户端连接数: %d", es.ConsumersCount())
		time.Sleep(2 * time.Second)
	}

	return &types.Response{
		Name: "123",
	}, nil
}
