package rest

import (
	"net/http"

	"github.com/dinel13/thesis-ac/payment/domain"
	"github.com/julienschmidt/httprouter"
)

func Routes(ph domain.PaymentRestHandlers) http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/verify/:id", ph.Verify)
	r.HandlerFunc(http.MethodPost, "/pay", ph.Create)

	return r
}
