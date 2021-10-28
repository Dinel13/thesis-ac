package app

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/dinel13/thesis-ac/course/db/driver"
	"github.com/dinel13/thesis-ac/course/domain"
	"google.golang.org/grpc"

	mygrpc "github.com/dinel13/thesis-ac/course/grpc"
	"github.com/dinel13/thesis-ac/course/proto"
	"github.com/dinel13/thesis-ac/course/rest"
	"github.com/dinel13/thesis-ac/course/service"
)

// StartRestServer starts the REST server
func StartRestServer() {
	port := ":8080"
	fmt.Printf("Staring REST server on port %s\n", port)

	dbClient := connectDB()

	// crete course repository db
	crDb := domain.NewCourseRepositoryDb(dbClient.SQL)

	// create course service
	cs := rest.NewCoursRestHandlers(service.NewCourseService(crDb))

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

	dbClient := connectDB()

	// crete course repository db
	crDb := domain.NewCourseRepositoryDb(dbClient.SQL)

	// create course service
	cs := mygrpc.NewGrpcHandler(service.NewCourseService(crDb))

	// create gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCourseServiceServer(s, &cs)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func connectDB() *driver.DB {
	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=thesis user=din password=postgres")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")
	return db
}
