// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

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

// Injectors from wire.go:

func newApp(viperViper *viper.Viper, logger *log.Logger) (*gin.Engine, func(), error) {
	jwtJWT := jwt.NewJwt(viperViper)
	handlerHandler := handler.NewHandler(logger)
	sidSid := sid.NewSid()
	serviceService := service.NewService(logger, sidSid, jwtJWT)
	db := repository.NewDB(viperViper)
	client := repository.NewRedis(viperViper)
	repositoryRepository := repository.NewRepository(db, client, logger)
	userRepository := repository.NewUserRepository(repositoryRepository)
	userService := service.NewUserService(serviceService, userRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	engine := server.NewServerHTTP(logger, jwtJWT, userHandler)
	return engine, func() {
	}, nil
}

// wire.go:

var HandlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler)

var ServiceSet = wire.NewSet(service.NewService, service.NewUserService)

var RepositorySet = wire.NewSet(repository.NewDB, repository.NewRedis, repository.NewRepository, repository.NewUserRepository)
