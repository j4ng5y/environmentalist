package envrionmentalistgrpc

import (
	"context"

	"github.com/j4ng5y/environmentalist/srv/environmentalistpb"
)

type Server struct{}

func NewServer() *Server {
	var S Server
	return &S
}

func (*Server) NewSecret(ctx context.Context, req *environmentalistpb.NewSecretRequest) (*environmentalistpb.NewSecretResponse, error) {
	resp := &environmentalistpb.NewSecretResponse{}
	return resp, nil
}

func (*Server) UpdateSecret(ctx context.Context, req *environmentalistpb.UpdateSecretRequest) (*environmentalistpb.UpdateSecretResponse, error) {
	resp := &environmentalistpb.UpdateSecretResponse{}
	return resp, nil
}

func (*Server) DeleteSecret(ctx context.Context, req *environmentalistpb.DeleteSecretRequest) (*environmentalistpb.DeleteSecretResponse, error) {
	resp := &environmentalistpb.DeleteSecretResponse{}
	return resp, nil
}

func (*Server) ViewSecret(ctx context.Context, req *environmentalistpb.ViewSecretRequest) (*environmentalistpb.ViewSecretResponse, error) {
	resp := &environmentalistpb.ViewSecretResponse{}
	return resp, nil
}
