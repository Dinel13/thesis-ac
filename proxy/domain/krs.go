package domain

import (
	"net/http"
)

type MataKuliah struct {
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	Sks      int    `json:"sks"`
	Dosen    string `json:"dosen"`
	Semester string `json:"semester"`
}

type Krs struct {
	IdMahasiswa int           `json:"id_mahasiswa"`
	MataKuliahs []*MataKuliah `json:"mata_kuliahs"`
}

type KrsResponse struct {
	Krs Krs `json:"krs"`
}

type ResKrsDelete struct {
	Krs interface{} `json:"krs"`
}

type KrsHandlers interface {
	Read(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type KrsGrpcClients interface {
	ReadKrs(id int, token string) (*Krs, error)
	CreateKrs(krs *Krs, token string) (*Krs, error)
	UpdateKrs(krs *Krs, id int, token string) (*Krs, error)
	DeleteKrs(id int, token string) error
}