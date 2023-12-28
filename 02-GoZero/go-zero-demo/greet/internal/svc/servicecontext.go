package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/greet/internal/config"
	"go-zero-demo/greet/internal/model/user"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println(c.Mysql.DataSource)
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUserModel(mysqlConn)}
}
