package grpc

import (
	"context"
	"log"

	"github.com/dinel13/thesis-ac/auth/domain"
	"github.com/dinel13/thesis-ac/auth/proto"
)

func NewGrpcHandler(s domain.AuthService) domain.AuthGrpcHandler {
	return grpcHandler{s}
}

type grpcHandler struct {
	service domain.AuthService
}

// read is a method that implements the Read method of the AuthGrpcHandler interface
func (h grpcHandler) Verify(ctx context.Context, req *proto.VerifyRequest) (*proto.VerifyResponse, error) {
	token := req.GetToken()

	// parse int32 to int64
	err := h.service.Verify(token)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.VerifyResponse{
		IsAuth: true,
	}

	return res, nil
}

// Create is a method that implements the Create method of the AuthGrpcHandler interface
func (h grpcHandler) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	username := req.GetUsername()
	password := req.GetPassword()

	auth := &domain.LoginSignupRequest{
		Username: username,
		Password: password,
	}

	authRespon, err := h.service.Login(auth)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.LoginResponse{
		Token: authRespon.Token,
	}

	return res, nil
}
