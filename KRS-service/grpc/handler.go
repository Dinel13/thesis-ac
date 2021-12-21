package grpc

import (
	"context"
	"log"
	"time"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/dinel13/thesis-ac/krs/proto"
)

func NewGrpcHandler(s domain.KrsService) domain.KrsGrpcHandler {
	return grpcHandler{s}
}

type grpcHandler struct {
	service domain.KrsService
}

// read is a method that implements the Read method of the KrsGrpcHandler interface
func (h grpcHandler) Read(ctx context.Context, req *proto.KrsRequest) (*proto.KrsResponse, error) {
	krsId := req.GetId()
	id := int(krsId)

	// parse int32 to int64
	krs, err := h.service.Read(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.KrsResponse{
		Krss: &proto.Krs{
			Id:          int32(krs.Id),
			Name:        krs.Name,
			Description: krs.Description,
			// Teacher:     int32(krs.Teacher),
			// Capacity:    int32(krs.Capacity),
			// Remain:      int32(krs.Remain),
			// CreatedAt:   krs.CreatedAt,
			// UpdatedAt:   krs.UpdatedAt,
		},
	}

	return res, nil
}

// Create is a method that implements the Create method of the KrsGrpcHandler interface
func (h grpcHandler) Create(ctx context.Context, req *proto.Krs) (*proto.KrsResponse, error) {
	krs := &domain.Krs{
		Id:          int(req.GetId()),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Teacher:     int(1),
		Capacity:    int(2),
		Remain:      int(3),
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	// parse int32 to int64
	krs, err := h.service.Create(krs)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.KrsResponse{
		Krss: &proto.Krs{
			Id:          int32(krs.Id),
			Name:        krs.Name,
			Description: krs.Description,
			// Teacher:     int32(krs.Teacher),
			// Capacity:    int32(krs.Capacity),
			// Remain:      int32(krs.Remain),
			// CreatedAt:   krs.CreatedAt,
			// UpdatedAt:   krs.UpdatedAt,
		},
	}

	return res, nil
}
