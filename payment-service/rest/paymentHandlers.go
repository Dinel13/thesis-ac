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
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
	}

	// get the payment from service
	payment, err := h.service.Verify(id)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the payment to response
	WriteEasyJson(w, http.StatusOK, &domain.PaymentWrapper{Payment: *payment})
}

// Create is handler for POST /payment to create COurse
func (h paymentHandlers) Create(w http.ResponseWriter, r *http.Request) {
	// get the payment from request body
	payment := &domain.PaymentRequest{}
	err := ReadEasyJson(r, payment)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// create the payment
	paymentCreated, err := h.service.Create(payment)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the payment to response
	WriteEasyJson(w, http.StatusCreated, &domain.PaymentWrapper{Payment: *paymentCreated})
}
