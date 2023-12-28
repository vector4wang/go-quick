package logic

import (
	"context"
	"database/sql"
	"fmt"
	"go-zero-demo/greet/internal/model/user"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetAddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetAddUserLogic {
	return &GreetAddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetAddUserLogic) GreetAddUser(req *types.AddUserRequest) (resp *types.BaseResponse, err error) {
	//user, err := svc.AddUser(req)
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &user.User{
		Name:     sql.NullString{String: req.Name, Valid: true},
		Password: req.Password,
		Mobile:   req.Mobile,
		Gender:   req.Gender,
		Nickname: req.Nickname,
		Type:     req.Type,
	})
	if err != nil {
		return &types.BaseResponse{Code: -1, Message: "fail"}, nil

	}
	fmt.Println(res)
	return &types.BaseResponse{Code: 0, Message: "success"}, nil
}
