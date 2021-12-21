package grpc

import (
	"context"
	"fmt"
	"log"

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
func (h grpcHandler) Read(ctx context.Context, req *proto.ReadKRSRequest) (*proto.KRSResponse, error) {
	token := req.GetToken()
	id_mahasiswa := req.GetIdMahasiswa()
	idMahasiswa := int(id_mahasiswa)

	fmt.Println("token: ", token)
	fmt.Println("idMahasiswa: ", idMahasiswa)

	// parse int32 to int64
	krs, err := h.service.Read(idMahasiswa)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.KRSResponse{
		Matakuliah: []*proto.MataKuliah{
			&proto.MataKuliah{
				Kode:     krs.MataKuliah[0].Kode,
				Nama:     krs.MataKuliah[0].Nama,
				Sks:      int32(krs.MataKuliah[0].Sks),
				Dosen:    krs.MataKuliah[0].Dosen,
				Semester: krs.MataKuliah[0].Semester,
			},
		}}

	return res, nil
}

// Create is a method that implements the Create method of the KrsGrpcHandler interface
func (h grpcHandler) Create(ctx context.Context, req *proto.CreateUpdateKRSRequest) (*proto.KRSResponse, error) {
	krs := &domain.Krs{
		Token:       req.GetToken(),
		IdMahasiswa: int(req.GetIdMahasiswa()),
		MataKuliah: []*domain.MataKuliah{
			&domain.MataKuliah{
				Kode:     req.GetMataKuliah().GetKode(),
				Nama:     req.GetMataKuliah().GetNama(),
				Sks:      int(req.GetMataKuliah().GetSks()),
				Dosen:    req.GetMataKuliah().GetDosen(),
				Semester: req.GetMataKuliah().GetSemester(),
			},
		},
	}

	// parse int32 to int64
	krs, err := h.service.Create(krs)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.KRSResponse{
		MataKuliahs: []*proto.MataKuliah{
			&proto.MataKuliah{
				Kode:     krs.MataKuliah[0].Kode,
				Nama:     krs.MataKuliah[0].Nama,
				Sks:      int32(krs.MataKuliah[0].Sks),
				Dosen:    krs.MataKuliah[0].Dosen,
				Semester: krs.MataKuliah[0].Semester,
			},
		},
	}

	return res, nil
}
