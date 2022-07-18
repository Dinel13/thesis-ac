package domain

type payment struct {
	IsPay bool `json:"isPay"`
}

type PaymentData struct {
	Payment payment `json:"payment"`
}
