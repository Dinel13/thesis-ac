package httpjson

import (
	"encoding/json"
	"log"
	"net/http"
)

// Start entrypoint
func Start() {
	http.HandleFunc("/", CreateKrs)
	log.Println(http.ListenAndServe(":60001", nil))
}

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

func CreateKrs(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var krs Krs
	decoder.Decode(&krs)
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{
		Code:        200,
		Message:     "OK",
		MataKuliahs: krs.MataKuliahs,
	})
}
