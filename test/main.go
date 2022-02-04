package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dinel13/thesis-ac/test/domain"
	mygrpc "github.com/dinel13/thesis-ac/test/grpc"
	"github.com/dinel13/thesis-ac/test/handlers"
	"github.com/dinel13/thesis-ac/test/proto"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port string = ":8085"

var ipKrs string = os.Getenv("IP_KRS")
var ipAuth string = os.Getenv("IP_AUTH")

func main() {

	// krs
	conn, err := grpc.Dial(fmt.Sprintf("%s:9090", ipKrs), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("can't connect grpc: %v", err)
	}

	gksc := proto.NewKrsServiceClient(conn) // create grpc client from proto
	gkc := mygrpc.NewKrsGrpcClient(gksc)    // create grpc client from grpc to other services

	rkh := handlers.NewRestKrsHandlers(ipKrs)
	gkh := handlers.NewGrpcKrsHandlers(gkc)

	// auth
	ruh := handlers.NewRestAuthHandlers(ipAuth)

	defer func() {
		conn.Close()
	}()

	server := http.Server{
		Addr:    port,
		Handler: routes(rkh, gkh, ruh),
	}

	fmt.Println("Server is running on port", port)
	fmt.Println("Endpoint for rest : /rest")
	fmt.Println("Endpoint for grpc : /grpc")
	// http.ListenAndServe(port, router)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func routes(rkh, gkh domain.KrsHandlers, ruh domain.AuthHandlers) http.Handler {
	r := httprouter.New()

	// krs rest
	r.HandlerFunc(http.MethodGet, "/rest/krs/:id", rkh.Read)
	r.HandlerFunc(http.MethodPost, "/rest/krs", rkh.Create)
	r.HandlerFunc(http.MethodPut, "/rest/krs/:id", rkh.Update)
	r.HandlerFunc(http.MethodDelete, "/rest/krs/:id", rkh.Delete)

	// krs grpc
	r.HandlerFunc(http.MethodGet, "/grpc/krs/:id", gkh.Read)
	r.HandlerFunc(http.MethodPost, "/grpc/krs", gkh.Create)
	r.HandlerFunc(http.MethodPut, "/grpc/krs/:id", gkh.Update)
	r.HandlerFunc(http.MethodDelete, "/grpc/krs/:id", gkh.Delete)

	// auth rest
	r.HandlerFunc(http.MethodPost, "/rest/auth/login", ruh.Login)
	r.HandlerFunc(http.MethodPost, "/rest/auth/verify", ruh.Verify)

	return r
}