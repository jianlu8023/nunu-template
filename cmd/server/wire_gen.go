// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func newApp(viperViper *viper.Viper, logger *log.Logger) (*gin.Engine, func(), error) {
	jwtJWT := jwt.NewJwt(viperViper)
	handlerHandler := handler2.NewHandler(logger)
	sidSid := sid.NewSid()
	serviceService := service2.NewService(logger, sidSid, jwtJWT)
	db := repository2.NewDB(viperViper)
	client := repository2.NewRedis(viperViper)
	repositoryRepository := repository2.NewRepository(db, client, logger)
	userRepository := repository2.NewUserRepository(repositoryRepository)
	userService := service2.NewUserService(serviceService, userRepository)
	userHandler := handler2.NewUserHandler(handlerHandler, userService)
	engine := server.NewServerHTTP(logger, jwtJWT, userHandler)
	return engine, func() {
	}, nil
}

// wire.go:

var HandlerSet = wire.NewSet(handler2.NewHandler, handler2.NewUserHandler)

var ServiceSet = wire.NewSet(service2.NewService, service2.NewUserService)

var RepositorySet = wire.NewSet(repository2.NewDB, repository2.NewRedis, repository2.NewRepository, repository2.NewUserRepository)
