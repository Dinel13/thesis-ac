package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dinel13/thesis-ac/course/db/driver"
	"github.com/dinel13/thesis-ac/course/domain"
	"github.com/dinel13/thesis-ac/course/service"
	"github.com/julienschmidt/httprouter"
)

func StartRestServer() {
	fmt.Println("Starting REST server")
	port := ":8080"

	router := httprouter.New()
	dbClient := connectDB()

	fmt.Printf("Staring application on port %s\n", port)

	CourseRepositoryDb := domain.NewCourseRepositoryDb(dbClient.SQL)

	ch := CourseHandlers{service.NewCourseService(CourseRepositoryDb)}

	router.HandlerFunc(http.MethodGet, "/course/:id", ch.GetCourse)

	// srv := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: routes(),
	// }

	// err := srv.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	addres := "localhost"
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%s", addres, port), router))

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
