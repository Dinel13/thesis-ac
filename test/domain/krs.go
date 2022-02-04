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

type ResKrsDelete struct {
	Krs interface{} `json:"krs"`
}

type KrsRestHandlers interface {
	RestKrsRead(http.ResponseWriter, *http.Request)
	RestKrsCreate(http.ResponseWriter, *http.Request)
	RestKrsUpdate(http.ResponseWriter, *http.Request)
	RestKrsDelete(http.ResponseWriter, *http.Request)
}
