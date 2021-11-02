package repository

import (
	"context"
	"database/sql"

	"github.com/dinel13/thesis-ac/course/domain"
)

type CourseRepository interface {
	GetCourseList(context.Context) ([]domain.Course, error)
	GetCourse(context.Context, int) (domain.Course, error)
	CreateCourse(context.Context, *domain.Course) (int, error)
	UpdateCourse(context.Context, *domain.Course) error
	DeleteCourse(context.Context, int) error
}

type CourseRepositoryImpl struct {
	SQL *sql.DB
}

// GetCourseList returns a list of courses
func (m CourseRepositoryImpl) GetCourseList(ctx context.Context) ([]domain.Course, error) {
	tx, err := m.SQL.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmn := "SELECT * FROM courses"
	rows, err := tx.QueryContext(ctx, stmn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courseList []domain.Course
	for rows.Next() {
		var course domain.Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description)
		if err != nil {
			return nil, err
		}
		courseList = append(courseList, course)
	}
	return courseList, nil
}

// GetCourse returns a course by id
func (m CourseRepositoryImpl) GetCourse(ctx context.Context, id int) (domain.Course, error) {
	var course domain.Course
	row := m.SQL.QueryRowContext(ctx, "SELECT * FROM courses WHERE id = $1", id)
	err := row.Scan(
		&course.ID,
		&course.Name,
		&course.Description,
		&course.Teacher,
		&course.Capacity,
		&course.Remain,
		&course.CreatedAt,
		&course.UpdatedAt,
	)
	if err != nil {
		return course, err
	}
	return course, nil
}

//	CreateCourse creates a new course
func (m CourseRepositoryImpl) CreateCourse(ctx context.Context, course *domain.Course) (int, error) {
	var id int
	stmn := "INSERT INTO course (name, description) VALUES ($1, $2) RETURNING id"
	err := m.SQL.QueryRowContext(ctx, stmn, course.Name, course.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateCourse updates an existing course
func (m CourseRepositoryImpl) UpdateCourse(ctx context.Context, course *domain.Course) error {
	_, err := m.SQL.Exec("UPDATE course SET name = $1, description = $2 WHERE id = $3", course.Name, course.Description, course.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCourse deletes an existing course
func (m CourseRepositoryImpl) DeleteCourse(ctx context.Context, id int) error {
	_, err := m.SQL.Exec("DELETE FROM course WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func NewCourseRepositoryImpl(dbClient *sql.DB) CourseRepositoryImpl {
	return CourseRepositoryImpl{SQL: dbClient}
}
