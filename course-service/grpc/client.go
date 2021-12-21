package grpc

import (
	"context"
	"log"

	"github.com/dinel13/thesis-ac/course/proto"
	"google.golang.org/grpc"
)

func Create() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewCourseServiceClient(conn)

	r, err := c.Create(context.Background(), &proto.Course{
		Id:          int32(1),
		Name:        "Go",
		Description: "Go is a programming language",
	})
	if err != nil {
		log.Fatalf("could not get course: %v", err)
	}

	log.Printf("Course: %s", r.Courses.Name)
	log.Printf("Course: %s", r.Courses.Description)
}

func Read() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewCourseServiceClient(conn)

	r, err := c.Read(context.Background(), &proto.CourseRequest{Id: int32(1)})
	if err != nil {
		log.Fatalf("could not get course: %v", err)
	}

	log.Printf("Course: %s", r.Courses.Name)
	log.Printf("Course: %s", r.Courses.Description)
}
