package grpc

import (
	"context"

	"github.com/dinel13/thesis-ac/course/domain"
	"github.com/dinel13/thesis-ac/course/proto"
)

func NewGrpcHandler(s domain.CourseService) domain.CourseGrpcHandler {
	return grpcHandler{s}
}

type grpcHandler struct {
	service domain.CourseService
}

func (h grpcHandler) Read(ctx context.Context, req *proto.CourseRequest) (*proto.CourseResponse, error) {
	courseId := req.GetId()
	id := int(courseId)

	// parse int32 to int64
	course, err := h.service.Read(id)
	if err != nil {
		return nil, err
	}

	res := &proto.CourseResponse{
		Courses: &proto.Course{
			Id:          int32(course.Id),
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
