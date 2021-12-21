package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/dinel13/thesis-ac/course/domain"
	"github.com/go-redis/redis/v8"
)

func NewCourseRepoRedisImpl(dbClient *redis.Client) domain.CourseRepository {
	return courseRepositoryRedisImpl{Rds: dbClient}
}

type courseRepositoryRedisImpl struct {
	Rds *redis.Client
}

//	CreateCourse creates a new course
func (m courseRepositoryRedisImpl) Create(ctx context.Context, course *domain.Course) (*domain.Course, error) {
	err := m.Rds.Set(ctx, strconv.Itoa(course.Id), course, 24*time.Hour).Err()
	if err != nil {
		return nil, err
	}
	return course, nil
}

// GetCourse returns a course by id
func (m courseRepositoryRedisImpl) Read(ctx context.Context, id int) (*domain.Course, error) {
	val, err := m.Rds.Get(ctx, strconv.Itoa(id)).Bytes()
	if err != nil {
		return nil, err
	}
	course := domain.Course{}
	err = json.Unmarshal(val, &course)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

// UpdateCourse updates an existing course
func (m courseRepositoryRedisImpl) Update(ctx context.Context, course *domain.Course) (*domain.Course, error) {
	return course, nil
}

// DeleteCourse deletes an existing course
func (m courseRepositoryRedisImpl) Delete(ctx context.Context, id int) error {
	return nil
}
