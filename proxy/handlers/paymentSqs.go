package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/julienschmidt/httprouter"
)

type paymentSqsHandlers struct {
	psc domain.PaymentSQSHandler
}

func (p *paymentSqsHandlers) Pay(w http.ResponseWriter, r *http.Request) {
	paymentRequest := &domain.PaymentRequest{}
	err := ReadJson(r, paymentRequest)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, false, "isPay")

}

func (p *paymentSqsHandlers) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")

	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	_, err := p.psc.SendMsg(&timestamp, &id)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	responChan := make(chan domain.ResponMsgWaitChan)
	go p.psc.WaitMsgSqs(responChan, &timestamp)

	respon := <-responChan
	if respon.Err != nil {
		WriteJsonError(w, respon.Err, http.StatusBadRequest)
		return
	}
	WriteJson(w, http.StatusOK, respon.IsPay, "isPay")
	return
}

func NewSqsPaymentHandlers(psc domain.PaymentSQSHandler) domain.PaymentHandlers {
	return &paymentSqsHandlers{
		psc: psc,
	}
}
