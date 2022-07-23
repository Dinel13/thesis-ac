package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/model"
	"github.com/dinel13/thesis-ac/test/rest-outbound"
	"github.com/julienschmidt/httprouter"
)

type paymentRestHandlers struct {
	ip        string
	payClient *http.Client
}

func (p *paymentRestHandlers) Pay(w http.ResponseWriter, r *http.Request) {
	paymentRequest := &model.PaymentRequest{}
	err := ReadPayJson(r, paymentRequest)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	resBody, err := rest.Pay(p.payClient, paymentRequest, p.ip)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusBadRequest)
		return
	}

	WritePayJson(w, http.StatusOK, &model.PaymentWraperResponse{PaymentResponse: *resBody})

}

func (p *paymentRestHandlers) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, _ := strconv.Atoi(params.ByName("id"))

	resBody, err := rest.VerifyPayment(p.payClient, p.ip, id)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusBadRequest)
		return
	}

	WritePayJson(w, http.StatusOK, &model.PaymentWraperResponse{PaymentResponse: *resBody})
}

func NewRestPaymentHandlers(ip string) domain.PaymentHandlers {
	payClient := &http.Client{Timeout: 15 * time.Second}
	return &paymentRestHandlers{
		ip:        ip,
		payClient: payClient,
	}
}
