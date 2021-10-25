package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc("GET", "/", index)

	return r
}

// handler for index
func index(w http.ResponseWriter, r *http.Request) {

}
