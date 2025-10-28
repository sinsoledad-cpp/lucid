// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package svc

import (
	"lucid/app/user/api/internal/config"
	"lucid/app/user/api/internal/logic"
	"lucid/app/user/data/model"
	"lucid/app/user/domain/repository"
	datarepo "lucid/app/user/domain/repository"
	"lucid/common/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	AuthzMiddleware rest.Middleware
	UserRepo        repository.UserRepository
	Converter       *logic.Converter
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 1. 创建数据库连接
	conn := sqlx.NewMysql(c.Database.DataSource)

	// 2. 初始化 goctl model
	userModel := model.NewUsersModel(conn)

	// 3. 初始化 data.repository 实现
	// 注意：这里的 NewUserRepoImpl 是你手写的
	userRepo := datarepo.NewUserRepoImpl(userModel)
	e, err := casbin.NewEnforcer(c.Casbin.ModelPath, c.Casbin.PolicyPath)
	if err != nil {
		logx.Must(err)
	}

	// 6. 从文件加载策略
	if err := e.LoadPolicy(); err != nil {
		logx.Must(err)
	}
	converter := logic.NewConverter()
	return &ServiceContext{
		Config:          c,
		AuthzMiddleware: middleware.NewAuthzMiddleware(e).Handle,
		UserRepo:        userRepo,
		Converter:       converter,
	}
}
