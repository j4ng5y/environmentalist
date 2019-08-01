package envrionmentalistgrpc

import (
	"context"

	"github.com/j4ng5y/environmentalist/srv/environmentalistpb"
)

// Server is a simple, empty struct to pin methods to
type Server struct{}

// NewServer returns an address pointer to a new, empty instance of Server
//
// Arguments:
//     None
//
// Returns:
//     (*Server)
func NewServer() *Server {
	var S Server
	return &S
}

// NewSecret is a gRPC method that adds a secret to the backing data stores
//
// Arguments:
//     ctx (context.Context):
//     req (*environmentalistpb.NewSecretRequest):
//
// Returns:
//     (*environmentalistpb.NewSecretResponse):
//     (error):
func (*Server) NewSecret(ctx context.Context, req *environmentalistpb.NewSecretRequest) (*environmentalistpb.NewSecretResponse, error) {
	resp := &environmentalistpb.NewSecretResponse{}
	return resp, nil
}

// UpdateSecret is a gRPC method that adds a secret to the backing data stores
//
// Arguments:
//     ctx (context.Context):
//     req (*environmentalistpb.NewSecretRequest):
//
// Returns:
//     (*environmentalistpb.NewSecretResponse):
//     (error):
func (*Server) UpdateSecret(ctx context.Context, req *environmentalistpb.UpdateSecretRequest) (*environmentalistpb.UpdateSecretResponse, error) {
	resp := &environmentalistpb.UpdateSecretResponse{}
	return resp, nil
}

// DeleteSecret is a gRPC method that adds a secret to the backing data stores
//
// Arguments:
//     ctx (context.Context):
//     req (*environmentalistpb.NewSecretRequest):
//
// Returns:
//     (*environmentalistpb.NewSecretResponse):
//     (error):
func (*Server) DeleteSecret(ctx context.Context, req *environmentalistpb.DeleteSecretRequest) (*environmentalistpb.DeleteSecretResponse, error) {
	resp := &environmentalistpb.DeleteSecretResponse{}
	return resp, nil
}

// ViewSecret is a gRPC method that adds a secret to the backing data stores
//
// Arguments:
//     ctx (context.Context):
//     req (*environmentalistpb.NewSecretRequest):
//
// Returns:
//     (*environmentalistpb.NewSecretResponse):
//     (error):
func (*Server) ViewSecret(ctx context.Context, req *environmentalistpb.ViewSecretRequest) (*environmentalistpb.ViewSecretResponse, error) {
	resp := &environmentalistpb.ViewSecretResponse{}
	return resp, nil
}
