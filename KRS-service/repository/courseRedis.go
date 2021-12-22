package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/go-redis/redis/v8"
)

func NewKrsRepoRedisImpl(dbClient *redis.Client) domain.KrsRepository {
	return krsRepositoryRedisImpl{Rds: dbClient}
}

type krsRepositoryRedisImpl struct {
	Rds *redis.Client
}

//	CreateKrs creates a new krs
func (m krsRepositoryRedisImpl) Create(ctx context.Context, krs *domain.Krs) (*domain.Krs, error) {
	krsJson, err := json.Marshal(krs)
	if err != nil {
		return nil, err
	}

	err = m.Rds.Set(ctx, strconv.Itoa(krs.IdMahasiswa), krsJson, 24*time.Hour).Err()
	if err != nil {
		return nil, err
	}
	return krs, nil
}

// GetKrs returns a krs by id
func (m krsRepositoryRedisImpl) Read(ctx context.Context, id int) (*domain.Krs, error) {
	val, err := m.Rds.Get(ctx, strconv.Itoa(id)).Bytes()
	if err != nil {
		return nil, err
	}
	krs := domain.Krs{}
	err = json.Unmarshal(val, &krs)
	if err != nil {
		return nil, err
	}
	return &krs, nil
}

// UpdateKrs updates an existing krs
func (m krsRepositoryRedisImpl) Update(ctx context.Context, krs *domain.Krs) (*domain.Krs, error) {
	krsJson, err := json.Marshal(krs)
	if err != nil {
		return nil, err
	}

	err = m.Rds.Set(ctx, strconv.Itoa(krs.IdMahasiswa), krsJson, 24*time.Hour).Err()
	if err != nil {
		return nil, err
	}
	return krs, nil
}

// DeleteKrs deletes an existing krs
func (m krsRepositoryRedisImpl) Delete(ctx context.Context, id int) error {
	err := m.Rds.Del(ctx, strconv.Itoa(id)).Err()
	if err != nil {
		return err
	}
	return nil
}
