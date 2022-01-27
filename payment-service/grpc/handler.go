package grpc

import (
	"context"
	"log"

	"github.com/dinel13/thesis-ac/payment/domain"
	"github.com/dinel13/thesis-ac/payment/proto"
)

func NewGrpcHandler(s domain.PaymentService) domain.PaymentGrpcHandler {
	return grpcHandler{s}
}

type grpcHandler struct {
	service domain.PaymentService
}

// read is a method that implements the Read method of the PaymentGrpcHandler interface
func (h grpcHandler) Verify(ctx context.Context, req *proto.VerifyPaymentRequest) (*proto.PaymentResponse, error) {
	id_mahasiswa := req.GetIdMahasiswa()
	idMahasiswa := int(id_mahasiswa)

	// parse int32 to int64
	payment, err := h.service.Verify(idMahasiswa)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.PaymentResponse{
		IsPay: payment.IsPay,
	}

	return res, nil
}

// Create is a method that implements the Create method of the PaymentGrpcHandler interface
func (h grpcHandler) Create(ctx context.Context, req *proto.CreatePaymentRequest) (*proto.PaymentResponse, error) {

	IdMahasiswa := int(req.GetIdMahasiswa())
	Jumlah := float64(req.GetJumlah())
	Metode := req.GetMetode()

	payment := &domain.PaymentRequest{
		IdMahasiswa: IdMahasiswa,
		Jumlah:      Jumlah,
		Metode:      Metode,
	}

	paymentRespon, err := h.service.Create(payment)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.PaymentResponse{
		IsPay: paymentRespon.IsPay,
	}

	return res, nil
}
