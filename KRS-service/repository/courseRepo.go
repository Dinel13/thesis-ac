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
	row := m.SQL.QueryRowContext(ctx, "SELECT mata_kuliah FROM krss WHERE id = $1", id)
	err := row.Scan(
		&krs.MataKuliahs,
	)
	if err != nil {
		return nil, err
	}
	return krs, nil
}

//	CreateKrs creates a new krs
func (m krsRepositoryImpl) Create(ctx context.Context, krs *domain.Krs) (*domain.Krs, error) {
	var newKrs *domain.Krs
	stmn := "INSERT INTO krs (matakuliah) VALUES ($1, $2) RETURNING id"
	err := m.SQL.QueryRowContext(ctx, stmn, krs.MataKuliahs).Scan(&krs.IdMahasiswa)
	if err != nil {
		return nil, err
	}
	return newKrs, nil
}

// UpdateKrs updates an existing krs
func (m krsRepositoryImpl) Update(ctx context.Context, krs *domain.Krs) (*domain.Krs, error) {
	var newKrs *domain.Krs
	stmn := "UPDATE krs SET matakuliah = $1 WHERE id_mahasiswa = $2 RETURNING id"
	err := m.SQL.QueryRowContext(ctx, stmn, krs.MataKuliahs, krs.IdMahasiswa).Scan(&krs.IdMahasiswa)
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
