package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/dinel13/thesis-ac/payment/domain"
	"github.com/dinel13/thesis-ac/payment/repository"
	"github.com/go-redis/redis/v8"

	mygrpc "github.com/dinel13/thesis-ac/payment/grpc"

	"github.com/dinel13/thesis-ac/payment/proto"
	"github.com/dinel13/thesis-ac/payment/rest"
	"github.com/dinel13/thesis-ac/payment/service"
)

// StartRestServer starts the REST server
func StartRestServer() {
	port := ":8082"
	fmt.Printf("Staring REST server on port %s\n", port)

	dbClient, paymentRepo := startRepoRedis()
	defer dbClient.Close()

	cs := rest.NewPaymentRestHandlers(service.NewPaymentService(paymentRepo))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: rest.Routes(cs),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// StartGRPCServer starts the gRPC server
func StartGRPCServer() {
	port := ":9092"
	fmt.Printf("Staring gRPC server on port %s\n", port)

	dbClient, paymentRepo := startRepoRedis()
	defer dbClient.Close()
	cs := mygrpc.NewGrpcHandler(service.NewPaymentService(paymentRepo))

	// create gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterPaymentServiceServer(s, cs)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startRepoRedis() (*redis.Client, domain.PaymentRepository) {
	dbClient := connectRedis()
	crDb := repository.NewPaymentRepoRedisImpl(dbClient)
	return dbClient, crDb
}

func connectRedis() *redis.Client {
	// connect to redis
	log.Println("Connecting to redis...")
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Cannot connect to redis! Dying...")
	}
	log.Println("Connected to redis!")
	return client
}

func main() {
	go StartRestServer()
	go StartGRPCServer()
	time.Sleep(10 * time.Minute)
}
