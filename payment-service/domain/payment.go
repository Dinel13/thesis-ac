package domain

import (
	"context"
	"net/http"
	"sync"

	"github.com/aws/aws-sdk-go/service/sqs"
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
	IsPay bool `json:"isPay"`
}

type PaymentWrapper struct {
	Payment PaymentResponse `json:"payment"`
}

type PaymentRepository interface {
	Create(context.Context, *PaymentRequest) (*PaymentResponse, error)
	Verify(context.Context, int) (*PaymentResponse, error)
}

type PaymentService interface {
	Create(*PaymentRequest) (*PaymentResponse, error)
	Verify(int) (*PaymentResponse, error)
}

type PaymentRestHandlers interface {
	Create(http.ResponseWriter, *http.Request)
	Verify(http.ResponseWriter, *http.Request)
}

type PaymentGrpcHandler interface {
	Create(context.Context, *proto.CreatePaymentRequest) (*proto.PaymentResponse, error)
	Verify(context.Context, *proto.VerifyPaymentRequest) (*proto.PaymentResponse, error)
}

type PaymentSQSHandler interface {
	WaitMsgSqs(sync.WaitGroup)
	GetLPMessages() (*sqs.ReceiveMessageOutput, error)
	SendMsg(*string, bool) (*string, error)
	DeleteMessage(*string) error
}
