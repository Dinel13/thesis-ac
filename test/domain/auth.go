package domain

import (
	"context"
	"net/http"

	"github.com/dinel13/thesis-ac/test/proto"
)

type LoginSignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginSignupResponse struct {
	Token string `json:"token"`
}

type VerifyTokenResponse struct {
	IsAuth bool `json:"isAuth"`
}

type AuthHandlers interface {
	Login(http.ResponseWriter, *http.Request)
	Verify(http.ResponseWriter, *http.Request)
}

type AuthGrpcHandler interface {
	Login(context.Context, *proto.LoginRequest) (*proto.LoginResponse, error)
	Verify(context.Context, *proto.VerifyRequest) (*proto.VerifyResponse, error)
}
