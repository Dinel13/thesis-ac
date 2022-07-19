package model

type PaymentRequest struct {
	IdMahasiswa int     `json:"id_mahasiswa"`
	Jumlah      float64 `json:"jumlah"`
	Metode      string  `json:"metode"`
}

type VerifyPaymentRequest struct {
	IdMahasiswa int `json:"id_mahasiswa"`
}

type PaymentResponse struct {
	IsPay bool `json:"isPay"`
}

type PaymentWraperResponse struct {
	PaymentResponse PaymentResponse `json:"payment"`
}
