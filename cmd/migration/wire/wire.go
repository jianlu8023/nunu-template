//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"nunu-template/cmd/migration/internal"
	"nunu-template/internal/repository"
	"nunu-template/pkg/log"
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

func NewApp(*viper.Viper, *log.Logger) (*internal.Migrate, func(), error) {
	panic(wire.Build(
		RepositorySet,
		internal.NewMigrate,
	))
}
