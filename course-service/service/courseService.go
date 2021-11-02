package service

import (
	"context"
	"time"

	"github.com/dinel13/thesis-ac/course/domain"
	"github.com/dinel13/thesis-ac/course/repository"
)

type CourseService interface {
	GetCourses() ([]domain.Course, error)
	GetCourse(int) (*domain.Course, error)
	AddCourse(*domain.Course) error
	UpdateCourse(*domain.Course) error
	DeleteCourse(int) error
}

type DefaultCourseService struct {
	Repo repository.CourseRepository
}

// GetCourse returns a course by id
func (s DefaultCourseService) GetCourse(id int) (course *domain.Course, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c, err := s.Repo.GetCourse(ctx, id)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// 	GetCourses returns all courses
func (s DefaultCourseService) GetCourses() (courses []domain.Course, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	courses, err = s.Repo.GetCourseList(ctx)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

// 	AddCourse adds a new course
func (s DefaultCourseService) AddCourse(course *domain.Course) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = s.Repo.CreateCourse(ctx, course)
	if err != nil {
		return err
	}
	return nil
}

// 	UpdateCourse
func (s DefaultCourseService) UpdateCourse(course *domain.Course) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = s.Repo.UpdateCourse(ctx, course)
	if err != nil {
		return err
	}
	return nil
}

// 	DeleteCourse
func (s DefaultCourseService) DeleteCourse(id int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = s.Repo.DeleteCourse(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewCourseService(Repo repository.CourseRepository) DefaultCourseService {
	return DefaultCourseService{Repo}
}
