package rest

import (
	"context"
	"errors"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/proto"
)

type authGrpcClient struct {
	s proto.AuthServiceClient
}

func (a *authGrpcClient) Login(req *domain.LoginSignupRequest) (*domain.LoginSignupResponse, error) {
	r, err := a.s.Login(context.Background(), &proto.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	if r.GetToken() == "" {
		return nil, errors.New("gagal login, pastikan username dan password benar")
	}

	return &domain.LoginSignupResponse{
		Token: r.GetToken(),
	}, nil
}

func (a *authGrpcClient) VerifyToken(token string) (*domain.VerifyTokenResponse, error) {
	r, err := a.s.Verify(context.Background(), &proto.VerifyRequest{
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	return &domain.VerifyTokenResponse{
		IsAuth: r.GetIsAuth(),
	}, nil
}

func NewAuthGrpcClient(s proto.AuthServiceClient) domain.AuthGrpcClients {
	return &authGrpcClient{s: s}
}
