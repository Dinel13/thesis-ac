package handlers

import (
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/model"
	"github.com/julienschmidt/httprouter"
)

type grpcPaymentHandlers struct {
	pgc domain.PaymentGrpcClients
}

func (g *grpcPaymentHandlers) Pay(w http.ResponseWriter, r *http.Request) {
	reqBody := &model.PaymentRequest{}
	err := ReadJson(r, reqBody)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	resBody, err := g.pgc.Pay(reqBody)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, resBody.IsPay, "isPay")
}

func (g *grpcPaymentHandlers) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, _ := strconv.Atoi(params.ByName("id"))

	resBody, err := g.pgc.VerifyPayment(id)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, resBody.IsPay, "isPay")
}

func NewGrpcPaymentHandlers(gac domain.PaymentGrpcClients) domain.PaymentHandlers {
	return &grpcPaymentHandlers{
		gac,
	}
}
