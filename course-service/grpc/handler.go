package grpc

import (
	"context"

	"github.com/dinel13/thesis-ac/course/proto"
	"github.com/dinel13/thesis-ac/course/service"
)

type GrpcHandler interface {
	GetCourse(ctx context.Context, req *proto.CourseRequest) (*proto.CourseResponse, error)
}

type DefaultGrpcHandler struct {
	service service.CourseService
}

func (h *DefaultGrpcHandler) GetCourse(ctx context.Context, req *proto.CourseRequest) (*proto.CourseResponse, error) {
	courseId := req.GetId()
	id := int(courseId)

	// parse int32 to int64
	course, err := h.service.GetCourse(id)
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

func NewGrpcHandler(s service.CourseService) DefaultGrpcHandler {
	return DefaultGrpcHandler{s}
}
