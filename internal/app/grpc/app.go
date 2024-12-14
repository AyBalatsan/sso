package grpcAppAuth

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "sso/internal/grpc/auth"
	"sso/pkg/logging"

	"google.golang.org/grpc"
)

type App struct {
	logger     *logging.Logger
	gRPCServer *grpc.Server
	port       int
}

// Создание нового grpc server для приложения.
func New(logger *logging.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.RegisterServerAPI(gRPCServer)

	return &App{
		logger:     logger, // использовать переданный логгер
		gRPCServer: gRPCServer,
		port:       port,
	}
}
func (a *App) MastRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
func (a *App) Run() error {
	const op = "grpcAppAuth.run"

	logger := a.logger.WithField("op", op)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("Запуск gRPC сервера", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() error {
	const op = "grpcAppAuth.stop"

	a.logger.WithField("op", op).Info("Остановка gRPC сервера", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
	return nil
}
