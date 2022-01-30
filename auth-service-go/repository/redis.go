package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dinel13/thesis-ac/auth/domain"
	"github.com/go-redis/redis/v8"
)

func NewAuthRepoRedisImpl(dbClient *redis.Client) domain.AuthRepository {
	return authRepositoryImpl{Rds: dbClient}
}

type authRepositoryImpl struct {
	Rds *redis.Client
}

//	CreateAuth creates a new auth
func (m authRepositoryImpl) Login(ctx context.Context, auth *domain.LoginSignupRequest) error {
	val, err := m.Rds.Get(ctx, auth.Username).Bytes()
	if err == redis.Nil {
		return errors.New("user not found")
	} else if err != nil {
		return err
	} else {
		user := domain.LoginSignupRequest{}
		err = json.Unmarshal(val, &user)
		if err != nil {
			return err
		}
		if user.Password != auth.Password {
			return errors.New("wrong password")
		}
		return nil
	}
}
