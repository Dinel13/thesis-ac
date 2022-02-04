package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/handlers"
	"github.com/julienschmidt/httprouter"
)

const port string = ":8080"

var ipKrs string = os.Getenv("IP_KRS")

func main() {

	rkh := handlers.NewRestKrsHandlers(ipKrs)

	server := http.Server{
		Addr:    port,
		Handler: routes(rkh),
	}

	fmt.Println("Server is running on port", port)
	// http.ListenAndServe(port, router)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func routes(rkh domain.KrsRestHandlers) http.Handler {
	r := httprouter.New()

	// krs rest
	r.HandlerFunc(http.MethodGet, "/krs/:id", rkh.KrsRestRead)
	r.HandlerFunc(http.MethodPost, "/krs", rkh.KrsRestCreate)
	r.HandlerFunc(http.MethodPut, "/krs/:id", rkh.KrsRestUpdate)
	r.HandlerFunc(http.MethodDelete, "/krs/:id", rkh.KrsRestDelete)

	// krs grpc
	r.HandlerFunc(http.MethodGet, "/v2/krs/:id", chg.Read)
	r.HandlerFunc(http.MethodPost, "/v2/krs", chg.Create)
	r.HandlerFunc(http.MethodPut, "/v2/krs/:id", chg.Update)
	r.HandlerFunc(http.MethodDelete, "/v2/krs/:id", chg.Delete)

	return r
}
