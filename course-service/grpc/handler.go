package grpc

import (
	"context"
	"log"

	"github.com/dinel13/thesis-ac/course/proto"
	"github.com/dinel13/thesis-ac/course/service"
)

type Handler struct {
	Cs service.CourseService
}

func (h *Handler) GetCourse(ctx context.Context, req *proto.CourseRequest) (*proto.CourseResponse, error) {
	log.Printf("Greet function was invoked with %v", req)
	courseId := req.GetId()
	id := int(courseId)

	// parse int32 to int64
	course, err := h.Cs.GetCourse(id)
	if err != nil {
		return nil, err
	}

	res := &proto.CourseResponse{
		Courses: &proto.Course{
			Id:          int32(course.ID),
			Name:        course.Name,
			Description: course.Description,
			Teacher:     int32(course.Teacher),
			Capacity:    int32(course.Capacity),
			Remain:      int32(course.Remain),
			CreatedAt:   course.CreatedAt,
			UpdatedAt:   course.UpdatedAt,
		},
	}

	return res, nil
}
