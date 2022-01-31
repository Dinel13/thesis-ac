package grpc

import (
	"context"
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

	// parse int32 to int64
	krs, err := h.service.Read(token, idMahasiswa)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// loop krs.MataKuliah to convert to proto.MataKuliah
	var mataKuliahs []*proto.MataKuliah
	for _, mataKuliah := range krs.MataKuliahs {
		mataKuliahs = append(mataKuliahs, &proto.MataKuliah{
			Kode:     mataKuliah.Kode,
			Nama:     mataKuliah.Nama,
			Sks:      int32(mataKuliah.Sks),
			Dosen:    mataKuliah.Dosen,
			Semester: mataKuliah.Semester,
		})
	}

	res := &proto.KRSResponse{
		MataKuliahs: mataKuliahs,
	}

	return res, nil
}

// Create is a method that implements the Create method of the KrsGrpcHandler interface
func (h grpcHandler) Create(ctx context.Context, req *proto.CreateUpdateKRSRequest) (*proto.KRSResponse, error) {

	_ = req.GetToken()
	idMahasiswa := int(req.GetIdMahasiswa())

	krs := &domain.Krs{
		IdMahasiswa: idMahasiswa,
		MataKuliahs: nil,
	}

	// loop proto.MataKuliah to convert to domain.MataKuliah untuk ambil data yg dikirm
	for _, mataKuliah := range req.GetMataKuliahs() {
		krs.MataKuliahs = append(krs.MataKuliahs, &domain.MataKuliah{
			Kode:     mataKuliah.Kode,
			Nama:     mataKuliah.Nama,
			Sks:      int(mataKuliah.Sks),
			Dosen:    mataKuliah.Dosen,
			Semester: mataKuliah.Semester,
		})
	}

	krs, err := h.service.Create(krs)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// loop krs.MataKuliah to convert to proto.MataKuliah
	var mataKuliahs []*proto.MataKuliah
	for _, mataKuliah := range krs.MataKuliahs {
		mataKuliahs = append(mataKuliahs, &proto.MataKuliah{
			Kode:     mataKuliah.Kode,
			Nama:     mataKuliah.Nama,
			Sks:      int32(mataKuliah.Sks),
			Dosen:    mataKuliah.Dosen,
			Semester: mataKuliah.Semester,
		})
	}

	res := &proto.KRSResponse{
		MataKuliahs: mataKuliahs,
	}

	return res, nil
}

// Update is a method that implements the Update method of the KrsGrpcHandler interface
func (h grpcHandler) Update(ctx context.Context, req *proto.CreateUpdateKRSRequest) (*proto.KRSResponse, error) {
	_ = req.GetToken()
	idMahasiswa := int(req.GetIdMahasiswa())

	// parse proto.KRS to domain.Krs
	krs := &domain.Krs{
		IdMahasiswa: idMahasiswa,
		MataKuliahs: nil,
	}

	// loop proto.MataKuliah to convert to domain.MataKuliah
	for _, mataKuliah := range req.GetMataKuliahs() {
		krs.MataKuliahs = append(krs.MataKuliahs, &domain.MataKuliah{
			Kode:     mataKuliah.Kode,
			Nama:     mataKuliah.Nama,
			Sks:      int(mataKuliah.Sks),
			Dosen:    mataKuliah.Dosen,
			Semester: mataKuliah.Semester,
		})
	}

	// parse int32 to int64
	krs, err := h.service.Update(krs)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// loop krs.MataKuliah to convert to proto.MataKuliah
	var mataKuliahs []*proto.MataKuliah
	for _, mataKuliah := range krs.MataKuliahs {
		mataKuliahs = append(mataKuliahs, &proto.MataKuliah{
			Kode:     mataKuliah.Kode,
			Nama:     mataKuliah.Nama,
			Sks:      int32(mataKuliah.Sks),
			Dosen:    mataKuliah.Dosen,
			Semester: mataKuliah.Semester,
		})
	}

	res := &proto.KRSResponse{
		MataKuliahs: mataKuliahs,
	}

	return res, nil
}

// Delete is a method that implements the Delete method of the KrsGrpcHandler interface
func (h grpcHandler) Delete(ctx context.Context, req *proto.DeleteKRSRequest) (*proto.DeleteKRSResponse, error) {
	token := req.GetToken()
	idMahasiswa := int(req.GetIdMahasiswa())

	err := h.service.Delete(token, idMahasiswa)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &proto.DeleteKRSResponse{
		Status: "success",
	}

	return res, nil
}
