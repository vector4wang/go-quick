package logic

import (
	"context"

	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetPostLogic {
	return &GreetPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetPostLogic) GreetPost(req *types.PostRequest) (resp *types.PostResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.PostResponse{
		Address: "Hello go-zero",
		Number:  "17777777777",
	}, nil
}
