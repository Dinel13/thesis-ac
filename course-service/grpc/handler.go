package grpc

import (
	"context"
	"log"
	"time"

	"github.com/dinel13/thesis-ac/course/domain"
	"github.com/dinel13/thesis-ac/course/proto"
)

func NewGrpcHandler(s domain.CourseService) domain.CourseGrpcHandler {
	return grpcHandler{s}
}

type grpcHandler struct {
	service domain.CourseService
}

// read is a method that implements the Read method of the CourseGrpcHandler interface
func (h grpcHandler) Read(ctx context.Context, req *proto.CourseRequest) (*proto.CourseResponse, error) {
	courseId := req.GetId()
	id := int(courseId)

	// parse int32 to int64
	course, err := h.service.Read(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.CourseResponse{
		Courses: &proto.Course{
			Id:          int32(course.Id),
			Name:        course.Name,
			Description: course.Description,
			// Teacher:     int32(course.Teacher),
			// Capacity:    int32(course.Capacity),
			// Remain:      int32(course.Remain),
			// CreatedAt:   course.CreatedAt,
			// UpdatedAt:   course.UpdatedAt,
		},
	}

	return res, nil
}

// Create is a method that implements the Create method of the CourseGrpcHandler interface
func (h grpcHandler) Create(ctx context.Context, req *proto.Course) (*proto.CourseResponse, error) {
	course := &domain.Course{
		Id:          int(req.GetId()),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Teacher:     int(1),
		Capacity:    int(2),
		Remain:      int(3),
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	// parse int32 to int64
	course, err := h.service.Create(course)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.CourseResponse{
		Courses: &proto.Course{
			Id:          int32(course.Id),
			Name:        course.Name,
			Description: course.Description,
			// Teacher:     int32(course.Teacher),
			// Capacity:    int32(course.Capacity),
			// Remain:      int32(course.Remain),
			// CreatedAt:   course.CreatedAt,
			// UpdatedAt:   course.UpdatedAt,
		},
	}

	return res, nil
}
