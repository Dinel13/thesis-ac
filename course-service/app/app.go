package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/dinel13/thesis-ac/course/domain"
	"github.com/dinel13/thesis-ac/course/repository"
	"google.golang.org/grpc"

	"github.com/go-redis/redis/v8"

	mygrpc "github.com/dinel13/thesis-ac/course/grpc"

	"github.com/dinel13/thesis-ac/course/proto"
	"github.com/dinel13/thesis-ac/course/rest"
	"github.com/dinel13/thesis-ac/course/service"
)

// StartRestServer starts the REST server
func StartRestServer() {
	port := ":8080"
	fmt.Printf("Staring REST server on port %s\n", port)

	// connect to SQL database
	// dbClient, courseRepo := startRepoSql()
	// defer dbClient.SQL.Close()

	// connect to Reddis database
	dbClient, courseRepo := startRepoRedis()
	defer dbClient.Close()

	cs := rest.NewCoursRestHandlers(service.NewCourseService(courseRepo))

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
	port := ":8081"
	fmt.Printf("Staring gRPC server on port %s\n", port)

	// connect to SQL database
	// dbClient, courseRepo := startRepoSql()
	// defer dbClient.SQL.Close()

	// connect to Reddis database
	dbClient, courseRepo := startRepoRedis()
	defer dbClient.Close()
	cs := mygrpc.NewGrpcHandler(service.NewCourseService(courseRepo))

	// create gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCourseServiceServer(s, cs)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func startRepoSql() (*driver.DB, domain.CourseRepository) {
// 	dbClient := connectDB()
// 	crDb := repository.NewCourseRepositoryImpl(dbClient.SQL)
// 	return dbClient, crDb
// }

func startRepoRedis() (*redis.Client, domain.CourseRepository) {
	dbClient := connectRedis()
	crDb := repository.NewCourseRepoRedisImpl(dbClient)
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
