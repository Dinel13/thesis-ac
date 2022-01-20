package service

import (
	"context"
	"time"

	"github.com/dinel13/thesis-ac/krs/domain"
)

func NewKrsService(Repo domain.KrsRepository) domain.KrsService {
	return DefaultKrsService{Repo}
}

type DefaultKrsService struct {
	Repo domain.KrsRepository
}

// GetKrs returns a krs by id
func (s DefaultKrsService) Read(token string, id int) (*domain.Krs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c, err := s.Repo.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// 	AddKrs adds a new krs
func (s DefaultKrsService) Create(krs *domain.Krs) (*domain.Krs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c, err := s.Repo.Create(ctx, krs)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// 	UpdateKrs
func (s DefaultKrsService) Update(krs *domain.Krs) (*domain.Krs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c, err := s.Repo.Update(ctx, krs)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// 	DeleteKrs
func (s DefaultKrsService) Delete(token string, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
