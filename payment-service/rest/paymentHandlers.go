package rest

import (
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/payment/domain"
	"github.com/julienschmidt/httprouter"
)

func NewPaymentRestHandlers(s domain.PaymentService) domain.PaymentRestHandlers {
	return paymentHandlers{s}
}

type paymentHandlers struct {
	service domain.PaymentService
}

// GetPayment is handler for GET /payment/{id}
func (h paymentHandlers) Verify(w http.ResponseWriter, r *http.Request) {
	// get the payment id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
	}

	// get the payment from service
	payment, err := h.service.Verify(id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the payment to response
	WriteJson(w, http.StatusOK, payment, "payment")
}

// Create is handler for POST /payment to create COurse
func (h paymentHandlers) Create(w http.ResponseWriter, r *http.Request) {
	// get the payment from request body
	payment := &domain.PaymentRequest{}
	err := ReadJson(r, payment)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// create the payment
	c, err := h.service.Create(payment)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the payment to response
	WriteJson(w, http.StatusCreated, c, "payment")
}
