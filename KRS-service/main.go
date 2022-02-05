package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/dinel13/thesis-ac/krs/proto"
	"github.com/dinel13/thesis-ac/krs/repository"
	"github.com/dinel13/thesis-ac/krs/rest"
	"github.com/dinel13/thesis-ac/krs/service"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"

	mygrpc "github.com/dinel13/thesis-ac/krs/grpc"
)

var urlAuth = os.Getenv("URL_AUTH")
var urlPay = os.Getenv("URL_PAYMENT")

// startRestServer starts the REST server
func startRestServer() {
	port := ":8080"
	fmt.Printf("Staring REST server on port %s\n", port)

	// connect to SQL database
	// dbClient, krsRepo := startRepoSql()
	// defer dbClient.SQL.Close()

	// connect to Reddis database
	dbClient, krsRepo := startRepoRedis()

	defer dbClient.Close()

	// cs := rest.NewCoursRestHandlers(service.NewKrsService(krsRepo))
	ks := service.NewKrsService(krsRepo)
	cs := rest.NewCoursRestHandlers(ks)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: rest.Routes(cs),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// startGRPCServer starts the gRPC server
func startGRPCServer() {
	port := ":9090"
	fmt.Printf("Staring gRPC server on port %s\n", port)

	// connect to SQL database
	// dbClient, krsRepo := startRepoSql()
	// defer dbClient.SQL.Close()

	// connect to Reddis database
	dbClient, krsRepo := startRepoRedis()

	connAuth, err := grpc.Dial(fmt.Sprintf("%s:9091", urlAuth), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	connPayment, err := grpc.Dial(fmt.Sprintf("%s:9092", urlPay), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		dbClient.Close()
		connAuth.Close()
		connPayment.Close()
	}()

	clientAuth := proto.NewAuthServiceClient(connAuth)
	clientPayment := proto.NewPaymentServiceClient(connPayment)

	ks := service.NewKrsService(krsRepo)

	cs := mygrpc.NewGrpcHandler(clientAuth, clientPayment, ks)

	// create gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterKrsServiceServer(s, cs)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func startRepoSql() (*driver.DB, domain.KrsRepository) {
// 	dbClient := connectDB()
// 	crDb := repository.NewKrsRepositoryImpl(dbClient.SQL)
// 	return dbClient, crDb
// }

func startRepoRedis() (*redis.Client, domain.KrsRepository) {
	dbClient := connectRedis()
	crDb := repository.NewKrsRepoRedisImpl(dbClient)
	return dbClient, crDb
}

// func connectDB() *driver.DB {
// 	// connect to database
// 	log.Println("Connecting to database...")
// 	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=thesis user=din password=postgres")
// 	if err != nil {
// 		log.Fatal("Cannot connect to database! Dying...")
// 	}
// 	log.Println("Connected to database!")
// 	return db
// }

func connectRedis() *redis.Client {
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
	go startRestServer()
	go startGRPCServer()
	time.Sleep(113880 * time.Hour)
}
