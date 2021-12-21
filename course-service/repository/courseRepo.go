package repository

import (
	"context"
	"database/sql"

	"github.com/dinel13/thesis-ac/course/domain"
)

func NewCourseRepositoryImpl(dbClient *sql.DB) domain.CourseRepository {
	return courseRepositoryImpl{SQL: dbClient}
}

type courseRepositoryImpl struct {
	SQL *sql.DB
}

// GetCourse returns a course by id
func (m courseRepositoryImpl) Read(ctx context.Context, id int) (*domain.Course, error) {
	var course *domain.Course
	row := m.SQL.QueryRowContext(ctx, "SELECT * FROM courses WHERE id = $1", id)
	err := row.Scan(
		&course.Id,
		&course.Name,
		&course.Description,
		&course.Teacher,
		&course.Capacity,
		&course.Remain,
		&course.CreatedAt,
		&course.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return course, nil
}

//	CreateCourse creates a new course
func (m courseRepositoryImpl) Create(ctx context.Context, course *domain.Course) (*domain.Course, error) {
	var newCourse *domain.Course
	stmn := "INSERT INTO course (name, description) VALUES ($1, $2) RETURNING *"
	err := m.SQL.QueryRowContext(ctx, stmn, course.Name, course.Description).Scan(&course.Id,
		&course.Name,
		&course.Description,
		&course.Teacher,
		&course.Capacity,
		&course.Remain,
		&course.CreatedAt,
		&course.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return newCourse, nil
}

// UpdateCourse updates an existing course
func (m courseRepositoryImpl) Update(ctx context.Context, course *domain.Course) (*domain.Course, error) {
	var newCourse *domain.Course
	stmn := "UPDATE course SET name = $1, description = $2 WHERE id = $3 RETURNING *"
	err := m.SQL.QueryRowContext(ctx, stmn, course.Name, course.Description, course.Id).Scan(&course.Id,
		&course.Name,
		&course.Description,
		&course.Teacher,
		&course.Capacity,
		&course.Remain,
		&course.CreatedAt,
		&course.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return newCourse, nil
}

// DeleteCourse deletes an existing course
func (m courseRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := m.SQL.Exec("DELETE FROM course WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
