package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes(ch CourseHandlers) http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/course/:id", ch.GetCourse)

	return r
}
