package grpc

import (
	"context"
	"log"

	"github.com/dinel13/thesis-ac/course/proto"
	"google.golang.org/grpc"
)

func Main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewCourseServiceClient(conn)

	r, err := c.GetCourse(context.Background(), &proto.CourseRequest{Id: int32(1)})
	if err != nil {
		log.Fatalf("could not get course: %v", err)
	}

	log.Printf("Course: %s", r.Courses.Name)
	log.Printf("Course: %s", r.Courses.Description)
	log.Printf("Course: %s", r.Courses.UpdatedAt)

}
