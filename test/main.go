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
	gkh := handlers.NewGrpcKrsHandlers(ipKrs)

	server := http.Server{
		Addr:    port,
		Handler: routes(rkh, gkh),
	}

	fmt.Println("Server is running on port", port)
	fmt.Println("Endpoint for rest : /rest")
	fmt.Println("Endpoint for grpc : /grpc")
	// http.ListenAndServe(port, router)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func routes(rkh, gkh domain.KrsHandlers) http.Handler {
	r := httprouter.New()

	// krs rest
	r.HandlerFunc(http.MethodGet, "/rest/krs/:id", rkh.Read)
	r.HandlerFunc(http.MethodPost, "/rest/krs", rkh.Create)
	r.HandlerFunc(http.MethodPut, "/rest/krs/:id", rkh.Update)
	r.HandlerFunc(http.MethodDelete, "/rest/krs/:id", rkh.Delete)

	// krs grpc
	r.HandlerFunc(http.MethodGet, "/grpc/krs/:id", rkh.Read)
	r.HandlerFunc(http.MethodPost, "/grpc/krs", rkh.Create)
	r.HandlerFunc(http.MethodPut, "/grpc/krs/:id", rkh.Update)
	r.HandlerFunc(http.MethodDelete, "/grpc/krs/:id", rkh.Delete)

	return r
}
