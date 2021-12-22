package domain

import (
	"context"
	"net/http"

	"github.com/dinel13/thesis-ac/payment/proto"
)

type PaymentRequest struct {
	IdMahasiswa int     `json:"id_mahasiswa"`
	Jumlah      float64 `json:"jumlah"`
	Metode      string  `json:"metode"`
}

type VerifyRequest struct {
	IdMahasiswa int `json:"id_mahasiswa"`
}

type PaymentResponse struct {
	IsPay bool `json:"is_pay"`
}

type PaymentRepository interface {
	Create(context.Context, *PaymentRequest) (*PaymentResponse, error)
	Verify(context.Context, *VerifyRequest) (*PaymentResponse, error)
}

type PaymentService interface {
	Create(*PaymentRequest) (*PaymentResponse, error)
	Verify(string, int) (*PaymentResponse, error)
}

type PaymentRestHandlers interface {
	Create(http.ResponseWriter, *http.Request)
	Verify(http.ResponseWriter, *http.Request)
}

type PaymentGrpcHandler interface {
	Create(context.Context, *proto.CreatePaymentRequest) (*proto.PaymentResponse, error)
	Verify(context.Context, *proto.VerifyPaymentRequest) (*proto.PaymentResponse, error)
}
