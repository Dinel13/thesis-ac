package rest

import (
	"net/http"

	"github.com/dinel13/thesis-ac/course/domain"
	"github.com/julienschmidt/httprouter"
)

func Routes(ch domain.CourseRestHandlers) http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/course/:id", ch.Read)
	r.HandlerFunc(http.MethodPost, "/course", ch.Create)
	r.HandlerFunc(http.MethodPut, "/course", ch.Update)
	r.HandlerFunc(http.MethodDelete, "/course/:id", ch.Delete)

	return r
}
