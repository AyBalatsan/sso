package app

import (
	"time"

	grpcAppAuth "sso/internal/app/grpc"
	"sso/pkg/logging"
)

type App struct {
	GRPCSrv *grpcAppAuth.App
}

func New(
	logger *logging.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO инициализировать хранилище (storage)

	// TODO init auth service (auth)

	grpcApp := grpcAppAuth.New(logger, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
