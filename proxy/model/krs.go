package model

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
	Status interface{} `json:"status"`
}
