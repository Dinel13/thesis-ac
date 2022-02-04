package handlers

import (
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/rest"
	"github.com/julienschmidt/httprouter"
)

type paymentRestHandlers struct {
	ip string
}

func (p *paymentRestHandlers) Pay(w http.ResponseWriter, r *http.Request) {
	paymentRequest := &domain.PaymentRequest{}
	err := ReadJson(r, paymentRequest)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	resBody, err := rest.Pay(paymentRequest, p.ip)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, resBody.IsPay, "isPay")

}

func (p *paymentRestHandlers) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, _ := strconv.Atoi(params.ByName("id"))

	resBody, err := rest.VerifyPayment(p.ip, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, resBody.IsPay, "isPay")
}

func NewRestPaymentHandlers(ip string) domain.PaymentHandlers {
	return &paymentRestHandlers{
		ip: ip,
	}
}
