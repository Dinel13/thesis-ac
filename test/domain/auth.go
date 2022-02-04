package domain

import (
	"net/http"
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
	VerifyToken(http.ResponseWriter, *http.Request)
}

type AuthGrpcClients interface {
	Login(req *LoginSignupRequest) (*LoginSignupResponse, error)
	VerifyToken(token string) (*VerifyTokenResponse, error)
}
