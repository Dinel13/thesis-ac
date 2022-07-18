package httpjson

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

type Response struct {
	Code        int           `json:"code"`
	Message     string        `json:"message"`
	MataKuliahs []*MataKuliah `json:"mata_kuliahs"`
}
