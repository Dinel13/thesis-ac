package service

import (
	"context"
	"time"

	"github.com/dinel13/thesis-ac/payment/domain"
)

func NewPaymentService(repo domain.PaymentRepository) domain.PaymentService {
	return paymentService{repo}
}

type paymentService struct {
	repo domain.PaymentRepository
}

// GetPayment returns a payment by id
func (p paymentService) Verify(id int) (*domain.PaymentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c, err := p.repo.Verify(ctx, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// 	AddPayment adds a new payment
func (p paymentService) Create(payment *domain.PaymentRequest) (*domain.PaymentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c, err := p.repo.Create(ctx, payment)
	if err != nil {
		return nil, err
	}
	return c, nil
}
