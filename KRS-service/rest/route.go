package rest

import (
	"net/http"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/julienschmidt/httprouter"
)

func Routes(ch domain.KrsRestHandlers, chg domain.KrsRestGrpcHandlers) http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/krs/:id", ch.Read)
	r.HandlerFunc(http.MethodPost, "/krs", ch.Create)
	r.HandlerFunc(http.MethodPut, "/krs/:id", ch.Update)
	r.HandlerFunc(http.MethodDelete, "/krs/:id", ch.Delete)

	// grpc
	r.HandlerFunc(http.MethodGet, "/v2/krs/:id", chg.Read)
	r.HandlerFunc(http.MethodPost, "/v2/krs", chg.Create)
	r.HandlerFunc(http.MethodPut, "/v2/krs/:id", chg.Update)
	r.HandlerFunc(http.MethodDelete, "/v2/krs/:id", chg.Delete)

	return r
}
