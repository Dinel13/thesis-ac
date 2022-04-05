package handlers

import (
	"net/http"

	"github.com/dinel13/thesis-ac/test/domain"
)

type grpcAuthHandlers struct {
	agc domain.AuthGrpcClients
}

func (g *grpcAuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	reqBody := &domain.LoginSignupRequest{}
	err := ReadJson(r, reqBody)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	resBody, err := g.agc.Login(reqBody)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, resBody.Token, "token")
}

func (g *grpcAuthHandlers) VerifyToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")[7:]
	resBody, err := g.agc.VerifyToken(token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, resBody.IsAuth, "isAuth")
}

func NewGrpcAuthHandlers(gac domain.AuthGrpcClients) domain.AuthHandlers {
	return &grpcAuthHandlers{
		gac,
	}
}
