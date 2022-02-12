package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dinel13/thesis-ac/auth/domain"
)

func NewAuthRestHandlers(s domain.AuthService) domain.AuthRestHandlers {
	return authHandlers{s}
}

type authHandlers struct {
	service domain.AuthService
}

// GetAuth is handler for GET /auth/{id}
func (h authHandlers) Verify(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		WriteJsonError(w, errors.New("invalid token"), http.StatusBadRequest)
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	err := h.service.Verify(tokenString)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	WriteJson(w, http.StatusOK, true, "isAuth")
}

// Create is handler for POST /auth to create COurse
func (h authHandlers) Login(w http.ResponseWriter, r *http.Request) {
	// get the auth from request body
	auth := &domain.LoginSignupRequest{}
	err := ReadJson(r, auth)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// create the auth
	c, err := h.service.Login(auth)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the auth to response
	WriteJson(w, http.StatusCreated, c.Token, "token")
}

func (h authHandlers) Signup(w http.ResponseWriter, r *http.Request) {
	// get the auth from request body
	auth := &domain.LoginSignupRequest{}
	err := ReadJson(r, auth)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// create the auth
	c, err := h.service.Signup(auth)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the auth to response
	WriteJson(w, http.StatusCreated, c.Token, "token")
}
