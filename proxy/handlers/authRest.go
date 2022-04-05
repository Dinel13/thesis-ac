package handlers

import (
	"net/http"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/rest"
)

type AuthRestHandlers struct {
	ip string
}

func (a *AuthRestHandlers) Login(w http.ResponseWriter, r *http.Request) {
	loginSignupRequest := &domain.LoginSignupRequest{}
	err := ReadJson(r, loginSignupRequest)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	resBody, err := rest.Login(loginSignupRequest, a.ip)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, resBody.Token, "token")

}

func (a *AuthRestHandlers) VerifyToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	resBody, err := rest.VerifyToken(a.ip, token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, resBody.IsAuth, "isAuth")
}

func NewRestAuthHandlers(ip string) domain.AuthHandlers {
	return &AuthRestHandlers{
		ip: ip,
	}
}
