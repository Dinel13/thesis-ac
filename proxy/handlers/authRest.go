package handlers

import (
	"net/http"
	"time"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/model"
	"github.com/dinel13/thesis-ac/test/rest-outbound"
)

type AuthRestHandlers struct {
	ip         string
	authClient *http.Client
}

func (a *AuthRestHandlers) Login(w http.ResponseWriter, r *http.Request) {
	loginSignupRequest := &model.LoginSignupRequest{}
	err := ReadJson(r, loginSignupRequest)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	resBody, err := rest.Login(a.authClient, loginSignupRequest, a.ip)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, resBody.Token, "token")

}

func (a *AuthRestHandlers) VerifyToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	resBody, err := rest.VerifyToken(a.authClient, a.ip, token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, resBody.IsAuth, "isAuth")
}

func NewRestAuthHandlers(ip string) domain.AuthHandlers {
	authClient := &http.Client{Timeout: 15 * time.Second}
	return &AuthRestHandlers{
		ip:         ip,
		authClient: authClient,
	}
}
