package service

import "github.com/dinel13/thesis-ac/course/domain"

type CourseService interface {
	GetCourses() ([]domain.Course, error)
	GetCourse(int) (*domain.Course, error)
	AddCourse(*domain.Course) error
	UpdateCourse(*domain.Course) error
	DeleteCourse(int) error
}

type DefaultCourseService struct {
	repo domain.CourseRepository
}

// GetCourse returns a course by id
func (s DefaultCourseService) GetCourse(id int) (course *domain.Course, err error) {
	c, err := s.repo.GetCourse(id)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// 	GetCourses returns all courses
func (s DefaultCourseService) GetCourses() (courses []domain.Course, err error) {
	courses, err = s.repo.GetCourseList()
	if err != nil {
		return nil, err
	}
	return courses, nil
}

// 	AddCourse adds a new course
func (s DefaultCourseService) AddCourse(course *domain.Course) (err error) {
	_, err = s.repo.CreateCourse(course)
	if err != nil {
		return err
	}
	return nil
}

// 	UpdateCourse
func (s DefaultCourseService) UpdateCourse(course *domain.Course) (err error) {
	err = s.repo.UpdateCourse(course)
	if err != nil {
		return err
	}
	return nil
}

// 	DeleteCourse
func (s DefaultCourseService) DeleteCourse(id int) (err error) {
	err = s.repo.DeleteCourse(id)
	if err != nil {
		return err
	}
	return nil
}

func NewCourseService(repo domain.CourseRepository) DefaultCourseService {
	return DefaultCourseService{repo}
}
