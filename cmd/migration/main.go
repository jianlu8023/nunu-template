package main

import (
	"nunu-template/cmd/migration/wire"
	"nunu-template/pkg/config"
	"nunu-template/pkg/log"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewApp(conf, logger)
	if err != nil {
		panic(err)
	}
	app.Run()
	defer cleanup()
}
