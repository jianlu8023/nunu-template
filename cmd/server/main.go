package main

import (
	"fmt"

	"go.uber.org/zap"
	"nunu-template/cmd/server/wire"
	"nunu-template/pkg/config"
	"nunu-template/pkg/http"
	"nunu-template/pkg/log"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewApp(conf, logger)
	if err != nil {
		panic(err)
	}
	logger.Info("server start", zap.String("host", "http://localhost:"+conf.GetString("http.port")))

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
	defer cleanup()

}
