package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/dinel13/thesis-ac/test/domain"
	mygrpc "github.com/dinel13/thesis-ac/test/grpc"
	"github.com/dinel13/thesis-ac/test/handlers"
	"github.com/dinel13/thesis-ac/test/proto"
	"github.com/dinel13/thesis-ac/test/sqs"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port string = ":8085"

var ipKrs string = os.Getenv("IP_KRS")
var ipAuth string = os.Getenv("IP_AUTH")
var ipPay string = os.Getenv("IP_PAYMENT")
var qsProxy = os.Getenv("Q_PROXY")
var qsPay = os.Getenv("Q_PAY")

func startSQSServer() domain.PaymentSQSHandler {
	fmt.Printf("Staring SQS server on url %s\n", qsPay)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	waitTime := int64(20)

	ss := sqs.NewSQSHandler(sess, &qsProxy, &qsPay, &waitTime)

	return ss
}

func main() {
	if ipKrs == "" {
		log.Fatal("IP_KRS is not set")
	}
	if ipAuth == "" {
		log.Fatal("IP_AUTH is not set")
	}
	if ipPay == "" {
		log.Fatal("IP_PAYMENT is not set")
	}
	if qsPay == "" {
		log.Fatal("sqs url pay is not set")
	}
	if qsProxy == "" {
		log.Fatal("sqs url proxy is not set")
	}

	log.Printf("use ipKrs: %s, ipAuth: %s, ipPay: %s, qsPay : %s", ipKrs, ipAuth, ipPay, qsPay)

	// krs
	connKrs, err := grpc.Dial(fmt.Sprintf("%s:9090", ipKrs), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("can't connect grpc: %v", err)
	}

	gksc := proto.NewKrsServiceClient(connKrs) // create grpc client from proto
	gkc := mygrpc.NewKrsGrpcClient(gksc)       // create grpc client from grpc to other services

	rkh := handlers.NewRestKrsHandlers(ipKrs)
	gkh := handlers.NewGrpcKrsHandlers(gkc)

	// auth
	connAuth, err := grpc.Dial(fmt.Sprintf("%s:9091", ipAuth), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("can't connect grpc: %v", err)
	}

	gasc := proto.NewAuthServiceClient(connAuth)
	gac := mygrpc.NewAuthGrpcClient(gasc)
	guh := handlers.NewGrpcAuthHandlers(gac)
	ruh := handlers.NewRestAuthHandlers(ipAuth)

	// payment
	connPay, err := grpc.Dial(fmt.Sprintf("%s:9092", ipPay), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("can't connect grpc: %v", err)
	}

	gpsc := proto.NewPaymentServiceClient(connPay)
	gpc := mygrpc.NewPaymentGrpcClient(gpsc)
	gph := handlers.NewGrpcPaymentHandlers(gpc)
	rph := handlers.NewRestPaymentHandlers(ipPay)

	// sqs
	sqsPayClient := startSQSServer()
	sph := handlers.NewSqsPaymentHandlers(sqsPayClient)

	defer func() {
		connKrs.Close()
		connAuth.Close()
	}()

	server := http.Server{
		Addr:    port,
		Handler: routes(rkh, gkh, ruh, guh, rph, gph, sph),
	}

	fmt.Println("Server is running on port", port)
	fmt.Println("Endpoint for rest : /rest")
	fmt.Println("Endpoint for grpc : /grpc")
	fmt.Println("Endpoint for sqs : /sqs")
	// http.ListenAndServe(port, router)
	err = server.ListenAndServe()
	// err = server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func routes(rkh, gkh domain.KrsHandlers, ruh, guh domain.AuthHandlers, rph, gph, sph domain.PaymentHandlers) http.Handler {
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
	r.HandlerFunc(http.MethodPost, "/rest/login", ruh.Login)
	r.HandlerFunc(http.MethodGet, "/rest/verify-token", ruh.VerifyToken)

	// auth grpc
	r.HandlerFunc(http.MethodPost, "/grpc/login", guh.Login)
	r.HandlerFunc(http.MethodGet, "/grpc/verify-token", guh.VerifyToken)

	// pay rest
	r.HandlerFunc(http.MethodPost, "/rest/pay", rph.Pay)
	r.HandlerFunc(http.MethodGet, "/rest/verify-pay/:id", rph.VerifyPayment)

	// pay grpc
	r.HandlerFunc(http.MethodPost, "/grpc/pay", gph.Pay)
	r.HandlerFunc(http.MethodGet, "/grpc/verify-pay/:id", gph.VerifyPayment)

	// pay sqs
	r.HandlerFunc(http.MethodPost, "/sqs/pay", sph.Pay)
	r.HandlerFunc(http.MethodGet, "/sqs/verify-pay/:id", sph.VerifyPayment)

	return r
}
