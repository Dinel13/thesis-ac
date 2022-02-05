package rest

import (
	"net/http"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/julienschmidt/httprouter"
)

func Routes(ch domain.KrsRestHandlers) http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/krs/:id", ch.Read)
	r.HandlerFunc(http.MethodPost, "/krs", ch.Create)
	r.HandlerFunc(http.MethodPut, "/krs/:id", ch.Update)
	r.HandlerFunc(http.MethodDelete, "/krs/:id", ch.Delete)

	return r
}
