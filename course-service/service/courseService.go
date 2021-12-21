package service

import (
	"context"
	"time"

	"github.com/dinel13/thesis-ac/course/domain"
)

func NewCourseService(Repo domain.CourseRepository) domain.CourseService {
	return DefaultCourseService{Repo}
}

type DefaultCourseService struct {
	Repo domain.CourseRepository
}

// GetCourse returns a course by id
func (s DefaultCourseService) Read(id int) (*domain.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c, err := s.Repo.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// 	AddCourse adds a new course
func (s DefaultCourseService) Create(course *domain.Course) (*domain.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c, err := s.Repo.Create(ctx, course)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// 	UpdateCourse
func (s DefaultCourseService) Update(course *domain.Course) (*domain.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c, err := s.Repo.Update(ctx, course)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// 	DeleteCourse
func (s DefaultCourseService) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
