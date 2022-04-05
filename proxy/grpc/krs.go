package rest

import (
	"context"
	"errors"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/proto"
)

type krsGrpcClient struct {
	s proto.KrsServiceClient
}

func NewKrsGrpcClient(s proto.KrsServiceClient) domain.KrsGrpcClients {
	return &krsGrpcClient{s: s}
}

func (k *krsGrpcClient) ReadKrs(id int, token string) (*domain.Krs, error) {
	r, err := k.s.Read(context.Background(), &proto.ReadKRSRequest{
		Token:       token,
		IdMahasiswa: int32(id),
	})

	if err != nil {
		return nil, err
	}

	// get krs
	var mataKuliahs []*domain.MataKuliah
	for _, mataKuliah := range r.GetMataKuliahs() {
		mataKuliahs = append(mataKuliahs, &domain.MataKuliah{
			Kode:     mataKuliah.GetKode(),
			Nama:     mataKuliah.GetNama(),
			Sks:      int(mataKuliah.GetSks()),
			Dosen:    mataKuliah.GetDosen(),
			Semester: mataKuliah.GetSemester(),
		})
	}

	return &domain.Krs{
		IdMahasiswa: id,
		MataKuliahs: mataKuliahs,
	}, nil

}

func (k *krsGrpcClient) CreateKrs(krs *domain.Krs, token string) (*domain.Krs, error) {
	// convert to proto
	var mataKuliahsProto []*proto.MataKuliah
	for _, mataKuliah := range krs.MataKuliahs {
		mataKuliahsProto = append(mataKuliahsProto, &proto.MataKuliah{
			Kode:     mataKuliah.Kode,
			Nama:     mataKuliah.Nama,
			Sks:      int32(mataKuliah.Sks),
			Dosen:    mataKuliah.Dosen,
			Semester: mataKuliah.Semester,
		})
	}

	r, err := k.s.Create(context.Background(), &proto.CreateUpdateKRSRequest{
		Token:       token,
		IdMahasiswa: int32(krs.IdMahasiswa),
		MataKuliahs: mataKuliahsProto,
	})

	if err != nil {
		return nil, err
	}

	// get krs
	var mataKuliahs []*domain.MataKuliah
	for _, mataKuliah := range r.GetMataKuliahs() {
		mataKuliahs = append(mataKuliahs, &domain.MataKuliah{
			Kode:     mataKuliah.GetKode(),
			Nama:     mataKuliah.GetNama(),
			Sks:      int(mataKuliah.GetSks()),
			Dosen:    mataKuliah.GetDosen(),
			Semester: mataKuliah.GetSemester(),
		})
	}

	return &domain.Krs{
		IdMahasiswa: krs.IdMahasiswa,
		MataKuliahs: mataKuliahs,
	}, nil
}

func (k *krsGrpcClient) UpdateKrs(krs *domain.Krs, id int, token string) (*domain.Krs, error) {
	// convert to proto
	var mataKuliahsProto []*proto.MataKuliah
	for _, mataKuliah := range krs.MataKuliahs {
		mataKuliahsProto = append(mataKuliahsProto, &proto.MataKuliah{
			Kode:     mataKuliah.Kode,
			Nama:     mataKuliah.Nama,
			Sks:      int32(mataKuliah.Sks),
			Dosen:    mataKuliah.Dosen,
			Semester: mataKuliah.Semester,
		})
	}

	r, err := k.s.Create(context.Background(), &proto.CreateUpdateKRSRequest{
		Token:       token,
		IdMahasiswa: int32(id),
		MataKuliahs: mataKuliahsProto,
	})

	if err != nil {
		return nil, err
	}

	// get krs
	var mataKuliahs []*domain.MataKuliah
	for _, mataKuliah := range r.GetMataKuliahs() {
		mataKuliahs = append(mataKuliahs, &domain.MataKuliah{
			Kode:     mataKuliah.GetKode(),
			Nama:     mataKuliah.GetNama(),
			Sks:      int(mataKuliah.GetSks()),
			Dosen:    mataKuliah.GetDosen(),
			Semester: mataKuliah.GetSemester(),
		})
	}

	return &domain.Krs{
		IdMahasiswa: id,
		MataKuliahs: mataKuliahs,
	}, nil
}

func (k *krsGrpcClient) DeleteKrs(id int, token string) error {
	r, err := k.s.Delete(context.Background(), &proto.DeleteKRSRequest{
		Token:       token,
		IdMahasiswa: int32(id),
	})

	if err != nil {
		return err
	}

	status := r.GetStatus()
	if status != "success" {
		return errors.New("failed to delete krs")
	}

	return nil

}
