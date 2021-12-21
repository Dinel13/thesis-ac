package domain

import (
	"context"
	"net/http"

	"github.com/dinel13/thesis-ac/course/proto"
)

type Course struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Teacher     int    `json:"teacher"`
	Capacity    int    `json:"capacity"`
	Remain      int    `json:"remain"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CourseRepository interface {
	Create(context.Context, *Course) (*Course, error)
	Read(context.Context, int) (*Course, error)
	Update(context.Context, *Course) (*Course, error)
	Delete(context.Context, int) error
}

type CourseService interface {
	Read(int) (*Course, error)
	Create(*Course) (*Course, error)
	Update(*Course) (*Course, error)
	Delete(int) error
}

type CourseRestHandlers interface {
	Read(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type CourseGrpcHandler interface {
	Read(context.Context, *proto.CourseRequest) (*proto.CourseResponse, error)
	Create(context.Context, *proto.Course) (*proto.CourseResponse, error)
}
