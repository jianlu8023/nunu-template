//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"nunu-template/internal/handler"
	"nunu-template/internal/repository"
	"nunu-template/internal/server"
	"nunu-template/internal/service"
	"nunu-template/pkg/helper/sid"
	"nunu-template/pkg/jwt"
	"nunu-template/pkg/log"
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

func NewApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		RepositorySet,
		ServiceSet,
		HandlerSet,
		server.NewServerHTTP,
		sid.NewSid,
		jwt.NewJwt,
	))
}
