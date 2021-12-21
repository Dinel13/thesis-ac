package repository

import (
	"context"
	"database/sql"

	"github.com/dinel13/thesis-ac/krs/domain"
)

func NewKrsRepositoryImpl(dbClient *sql.DB) domain.KrsRepository {
	return krsRepositoryImpl{SQL: dbClient}
}

type krsRepositoryImpl struct {
	SQL *sql.DB
}

// GetKrs returns a krs by id
func (m krsRepositoryImpl) Read(ctx context.Context, id int) (*domain.Krs, error) {
	var krs *domain.Krs
	row := m.SQL.QueryRowContext(ctx, "SELECT * FROM krss WHERE id = $1", id)
	err := row.Scan(
		&krs.Id,
		&krs.Name,
		&krs.Description,
		&krs.Teacher,
		&krs.Capacity,
		&krs.Remain,
		&krs.CreatedAt,
		&krs.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return krs, nil
}

//	CreateKrs creates a new krs
func (m krsRepositoryImpl) Create(ctx context.Context, krs *domain.Krs) (*domain.Krs, error) {
	var newKrs *domain.Krs
	stmn := "INSERT INTO krs (name, description) VALUES ($1, $2) RETURNING *"
	err := m.SQL.QueryRowContext(ctx, stmn, krs.Name, krs.Description).Scan(&krs.Id,
		&krs.Name,
		&krs.Description,
		&krs.Teacher,
		&krs.Capacity,
		&krs.Remain,
		&krs.CreatedAt,
		&krs.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return newKrs, nil
}

// UpdateKrs updates an existing krs
func (m krsRepositoryImpl) Update(ctx context.Context, krs *domain.Krs) (*domain.Krs, error) {
	var newKrs *domain.Krs
	stmn := "UPDATE krs SET name = $1, description = $2 WHERE id = $3 RETURNING *"
	err := m.SQL.QueryRowContext(ctx, stmn, krs.Name, krs.Description, krs.Id).Scan(&krs.Id,
		&krs.Name,
		&krs.Description,
		&krs.Teacher,
		&krs.Capacity,
		&krs.Remain,
		&krs.CreatedAt,
		&krs.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return newKrs, nil
}

// DeleteKrs deletes an existing krs
func (m krsRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := m.SQL.Exec("DELETE FROM krs WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
