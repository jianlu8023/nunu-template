//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	handler2 "nunu-template/internal/cn/cas/xjipc/blockchain/handler"
	repository2 "nunu-template/internal/cn/cas/xjipc/blockchain/repository"
	"nunu-template/internal/cn/cas/xjipc/blockchain/server"
	service2 "nunu-template/internal/cn/cas/xjipc/blockchain/service"
	"nunu-template/pkg/helper/sid"
	"nunu-template/pkg/jwt"
	"nunu-template/pkg/log"
)

var HandlerSet = wire.NewSet(
	handler2.NewHandler,
	handler2.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service2.NewService,
	service2.NewUserService,
)

var RepositorySet = wire.NewSet(
	repository2.NewDB,
	repository2.NewRedis,
	repository2.NewRepository,
	repository2.NewUserRepository,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		RepositorySet,
		ServiceSet,
		HandlerSet,
		server.NewServerHTTP,
		sid.NewSid,
		jwt.NewJwt,
	))
}
