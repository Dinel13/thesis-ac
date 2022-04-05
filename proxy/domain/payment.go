package domain

import (
	"net/http"
)

type PaymentRequest struct {
	IdMahasiswa int     `json:"id_mahasiswa"`
	Jumlah      float64 `json:"jumlah"`
	Metode      string  `json:"metode"`
}

type VerifyPaymentRequest struct {
	IdMahasiswa int `json:"id_mahasiswa"`
}

type PaymentResponse struct {
	IsPay bool `json:"isPay"`
}

type PaymentWraperResponse struct {
	PaymentResponse PaymentResponse `json:"payment"`
}

type PaymentHandlers interface {
	Pay(http.ResponseWriter, *http.Request)
	VerifyPayment(http.ResponseWriter, *http.Request)
}

type PaymentGrpcClients interface {
	Pay(req *PaymentRequest) (*PaymentResponse, error)
	VerifyPayment(id int) (*PaymentResponse, error)
}
