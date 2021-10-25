package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dinel13/thesis-ac/course/db/driver"
	"github.com/dinel13/thesis-ac/course/domain"
	"github.com/dinel13/thesis-ac/course/rest"
	"github.com/dinel13/thesis-ac/course/service"
)

func StartRestServer() {
	port := ":8080"
	fmt.Printf("Staring REST server on port %s\n", port)

	dbClient := connectDB()

	// crete course repository db
	crDb := domain.NewCourseRepositoryDb(dbClient.SQL)

	// create course service
	cs := rest.CourseHandlers{service.NewCourseService(crDb)}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: rest.Routes(cs),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
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
