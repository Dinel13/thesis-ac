package domain

import (
	"net/http"

	"github.com/dinel13/thesis-ac/test/model"
)

type AuthHandlers interface {
	Login(http.ResponseWriter, *http.Request)
	VerifyToken(http.ResponseWriter, *http.Request)
}

type AuthGrpcClients interface {
	Login(req *model.LoginSignupRequest) (*model.LoginSignupResponse, error)
	VerifyToken(token string) (*model.VerifyTokenResponse, error)
}
