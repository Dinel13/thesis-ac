package domain

import (
	"context"
	"net/http"

	"github.com/dinel13/thesis-ac/krs/proto"
)

type MataKuliah struct {
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	Sks      int    `json:"sks"`
	Dosen    string `json:"dosen"`
	Semester string `json:"semester"`
}

type Krs struct {
	Token       string        `json:"token"`
	IdMahasiswa int           `json:"id_mahasiswa"`
	MataKuliahs []*MataKuliah `json:"matakuliahs"`
}

type KrsRepository interface {
	Create(context.Context, *Krs) (*Krs, error)
	Read(context.Context, int) (*Krs, error)
	Update(context.Context, *Krs) (*Krs, error)
	Delete(context.Context, int) error
}

type KrsService interface {
	Read(int) (*Krs, error)
	Create(*Krs) (*Krs, error)
	Update(*Krs) (*Krs, error)
	Delete(int) error
}

type KrsRestHandlers interface {
	Read(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type KrsGrpcHandler interface {
	Read(context.Context, *proto.ReadKRSRequest) (*proto.KRSResponse, error)
	Create(context.Context, *proto.CreateUpdateKRSRequest) (*proto.KRSResponse, error)
	Update(context.Context, *proto.CreateUpdateKRSRequest) (*proto.KRSResponse, error)
	Delete(context.Context, *proto.DeleteKRSRequest) (*proto.KRSResponse, error)
}
