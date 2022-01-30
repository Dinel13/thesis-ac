package domain

import (
	"context"
	"net/http"

	"github.com/dinel13/thesis-ac/auth/proto"
)

type LoginSignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginSignupResponse struct {
	Token string `json:"token"`
}

type VerifyResponse struct {
	IsAuth bool `json:"isAuth"`
}

type AuthRepository interface {
	Login(context.Context, *LoginSignupRequest) error
}

type AuthService interface {
	Login(*LoginSignupRequest) (*LoginSignupResponse, error)
	Verify(string) error
}

type AuthRestHandlers interface {
	Login(http.ResponseWriter, *http.Request)
	Verify(http.ResponseWriter, *http.Request)
}

type AuthGrpcHandler interface {
	Login(context.Context, *proto.LoginRequest) (*proto.LoginResponse, error)
	Verify(context.Context, *proto.VerifyRequest) (*proto.VerifyResponse, error)
}
