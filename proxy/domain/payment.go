package domain

import (
	"net/http"

	"github.com/dinel13/thesis-ac/test/model"
)

type ResponMsgWaitChan struct {
	IsPay bool  `json:"isPay"`
	Err   error `json:"err"`
}

type PaymentHandlers interface {
	Pay(http.ResponseWriter, *http.Request)
	VerifyPayment(http.ResponseWriter, *http.Request)
}

type PaymentGrpcClients interface {
	Pay(req *model.PaymentRequest) (*model.PaymentResponse, error)
	VerifyPayment(id int) (*model.PaymentResponse, error)
}
