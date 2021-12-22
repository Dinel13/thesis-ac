package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/dinel13/thesis-ac/payment/domain"
	"github.com/go-redis/redis/v8"
)

func NewPaymentRepoRedisImpl(dbClient *redis.Client) domain.PaymentRepository {
	return paymentRepositoryImpl{Rds: dbClient}
}

type paymentRepositoryImpl struct {
	Rds *redis.Client
}

//	CreatePayment creates a new payment
func (m paymentRepositoryImpl) Create(ctx context.Context, payment *domain.PaymentRequest) (*domain.PaymentResponse, error) {
	paymentJson, err := json.Marshal(payment)
	if err != nil {
		return nil, err
	}

	err = m.Rds.Set(ctx, strconv.Itoa(payment.IdMahasiswa), paymentJson, 24*time.Hour).Err()
	if err != nil {
		return nil, err
	}
	res := &domain.PaymentResponse{IsPay: true}
	return res, nil
}

// GetPayment returns a payment by id
func (m paymentRepositoryImpl) Verify(ctx context.Context, id int) (*domain.PaymentResponse, error) {
	err := m.Rds.Get(ctx, strconv.Itoa(id)).Err()
	if err != nil {
		return nil, err
	}
	payment := domain.PaymentResponse{
		IsPay: true,
	}
	if err != nil {
		return nil, err
	}
	return &payment, nil
}
