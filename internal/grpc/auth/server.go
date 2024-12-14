package auth

import (
	"context"

	ssov1 "github.com/AyBalatsan/protoc/gen/go/sso"
	"google.golang.org/grpc"
)

type ServerAPI struct {
	ssov1.UnimplementedOauthServer
}

// регистрирует обработчик
func RegisterServerAPI(gRPC *grpc.Server) {
	ssov1.RegisterOauthServer(gRPC, &ServerAPI{})
}

func (s *ServerAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("impement me")
}
func (s *ServerAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("impement me")
}
func (s *ServerAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("impement me")
}
func (s *ServerAPI) Logout(ctx context.Context, req *ssov1.LogoutRequest) (*ssov1.LogoutResponse, error) {
	panic("impement me")
}
