package rest

import (
	"context"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/model"
	"github.com/dinel13/thesis-ac/test/proto"
)

type paymentGrpcClient struct {
	s proto.PaymentServiceClient
}

func (p *paymentGrpcClient) Pay(req *model.PaymentRequest) (*model.PaymentResponse, error) {
	r, err := p.s.Create(context.Background(), &proto.CreatePaymentRequest{
		IdMahasiswa: int32(req.IdMahasiswa),
		Jumlah:      float32(req.Jumlah),
		Metode:      req.Metode,
	})

	if err != nil {
		return nil, err
	}

	return &model.PaymentResponse{
		IsPay: r.GetIsPay(),
	}, nil
}

func (a *paymentGrpcClient) VerifyPayment(id int) (*model.PaymentResponse, error) {
	r, err := a.s.Verify(context.Background(), &proto.VerifyPaymentRequest{
		IdMahasiswa: int32(id),
	})

	if err != nil {
		return nil, err
	}

	return &model.PaymentResponse{
		IsPay: r.GetIsPay(),
	}, nil
}

func NewPaymentGrpcClient(s proto.PaymentServiceClient) domain.PaymentGrpcClients {
	return &paymentGrpcClient{s: s}
}
