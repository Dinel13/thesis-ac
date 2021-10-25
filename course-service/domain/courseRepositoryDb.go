package domain

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type CourseRepository interface {
	GetCourseList() ([]Course, error)
	GetCourse(int) (Course, error)
	CreateCourse(*Course) (int, error)
	UpdateCourse(*Course) error
	DeleteCourse(int) error
}

type CourseRepositoryDb struct {
	SQL *sql.DB
}

// GetCourseList returns a list of courses
func (m CourseRepositoryDb) GetCourseList() ([]Course, error) {
	rows, err := m.SQL.Query("SELECT * FROM course")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courseList []Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description)
		if err != nil {
			return nil, err
		}
		courseList = append(courseList, course)
	}
	return courseList, nil
}

// GetCourse returns a course by id
func (m CourseRepositoryDb) GetCourse(id int) (Course, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var course Course
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
func (m CourseRepositoryDb) CreateCourse(course *Course) (int, error) {
	var id int
	err := m.SQL.QueryRow("INSERT INTO course (name, description) VALUES ($1, $2) RETURNING id", course.Name, course.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateCourse updates an existing course
func (m CourseRepositoryDb) UpdateCourse(course *Course) error {
	_, err := m.SQL.Exec("UPDATE course SET name = $1, description = $2 WHERE id = $3", course.Name, course.Description, course.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCourse deletes an existing course
func (m CourseRepositoryDb) DeleteCourse(id int) error {
	_, err := m.SQL.Exec("DELETE FROM course WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func NewCourseRepositoryDb(dbClient *sql.DB) CourseRepositoryDb {
	return CourseRepositoryDb{SQL: dbClient}
}
