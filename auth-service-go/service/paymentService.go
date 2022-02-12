package service

import (
	"context"
	"time"

	"github.com/dinel13/thesis-ac/auth/domain"
	"github.com/dinel13/thesis-ac/auth/helper"
)

func NewAuthService(repo domain.AuthRepository) domain.AuthService {
	return authService{repo}
}

type authService struct {
	repo domain.AuthRepository
}

// GetAuth returns a auth by id
func (p authService) Verify(tokenString string) error {
	err := helper.ParseToken(tokenString, "secretKey@123")
	if err != nil {
		return err
	}
	return nil
}

// 	AddAuth adds a new auth
func (p authService) Login(auth *domain.LoginSignupRequest) (*domain.LoginSignupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.repo.Login(ctx, auth)
	if err != nil {
		return nil, err
	}

	// create token
	token, err := helper.CreateToken(auth.Username, "secretKey@123")
	if err != nil {
		return nil, err
	}

	res := &domain.LoginSignupResponse{
		Token: token,
	}

	return res, nil
}

func (p authService) Signup(auth *domain.LoginSignupRequest) (*domain.LoginSignupResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.repo.Signup(ctx, auth)
	if err != nil {
		return nil, err
	}

	// create token
	token, err := helper.CreateToken(auth.Username, "secretKey@123")
	if err != nil {
		return nil, err
	}

	res := &domain.LoginSignupResponse{
		Token: token,
	}

	return res, nil
}
