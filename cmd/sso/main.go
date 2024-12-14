package main

import (
	"fmt"

	"sso/internal/app"
	"sso/internal/config"
	"sso/pkg/logging"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(*cfg)

	logger := logging.GetLogger()

	logger.Info("Start application")

	application := app.New(logger, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	application.GRPCSrv.MastRun()
	//	app

	// запуск grpc
}
