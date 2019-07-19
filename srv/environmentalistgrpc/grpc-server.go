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

func (*Server) NewHashicorpVaultSecret(ctx context.Context, req *environmentalistpb.HashicorpVault_NewSecretRequest) (*environmentalistpb.HashicorpVault_NewSecretResponse, error) {
	resp := &environmentalistpb.HashicorpVault_NewSecretResponse{}
	return resp, nil
}

func (*Server) UpdateHashicorpVaultSecret(ctx context.Context, req *environmentalistpb.HashicorpVault_UpdateSecretRequest) (*environmentalistpb.HashicorpVault_UpdateSecretResponse, error) {
	resp := &environmentalistpb.HashicorpVault_UpdateSecretResponse{}
	return resp, nil
}

func (*Server) DeleteHashicorpVaultSecret(ctx context.Context, req *environmentalistpb.HashicorpVault_DeleteSecretRequest) (*environmentalistpb.HashicorpVault_DeleteSecretResponse, error) {
	resp := &environmentalistpb.HashicorpVault_DeleteSecretResponse{}
	return resp, nil
}

func (*Server) ViewHashicorpVaultSecret(ctx context.Context, req *environmentalistpb.HashicorpVault_ViewSecretRequest) (*environmentalistpb.HashicorpVault_ViewSecretResponse, error) {
	resp := &environmentalistpb.HashicorpVault_ViewSecretResponse{}
	return resp, nil
}

func (*Server) NewAWSSSMSecret(ctx context.Context, req *environmentalistpb.AWSSSM_NewSecretRequest) (*environmentalistpb.AWSSSM_NewSecretResponse, error) {
	resp := &environmentalistpb.AWSSSM_NewSecretResponse{}
	return resp, nil
}

func (*Server) UpdateAWSSSMSecret(ctx context.Context, req *environmentalistpb.AWSSSM_UpdateSecretRequest) (*environmentalistpb.AWSSSM_UpdateSecretResponse, error) {
	resp := &environmentalistpb.AWSSSM_UpdateSecretResponse{}
	return resp, nil
}

func (*Server) DeleteAWSSSMSecret(ctx context.Context, req *environmentalistpb.AWSSSM_DeleteSecretRequest) (*environmentalistpb.AWSSSM_DeleteSecretResponse, error) {
	resp := &environmentalistpb.AWSSSM_DeleteSecretResponse{}
	return resp, nil
}

func (*Server) ViewAWSSSMSecret(ctx context.Context, req *environmentalistpb.AWSSSM_ViewSecretRequest) (*environmentalistpb.AWSSSM_ViewSecretResponse, error) {
	resp := &environmentalistpb.AWSSSM_ViewSecretResponse{}
	return resp, nil
}
