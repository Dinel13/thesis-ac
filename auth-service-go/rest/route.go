package rest

import (
	"net/http"

	"github.com/dinel13/thesis-ac/auth/domain"
	"github.com/julienschmidt/httprouter"
)

func Routes(ah domain.AuthRestHandlers) http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodPost, "/verify", ah.Verify)
	r.HandlerFunc(http.MethodPost, "/login", ah.Login)
	r.HandlerFunc(http.MethodPost, "/signup", ah.Signup)

	return r
}
