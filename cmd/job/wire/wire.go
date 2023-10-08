//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"nunu-template/internal/job"
	"nunu-template/pkg/log"
)

var JobSet = wire.NewSet(job.NewJob)

func NewApp(*viper.Viper, *log.Logger) (*job.Job, func(), error) {
	panic(wire.Build(
		JobSet,
	))
}
