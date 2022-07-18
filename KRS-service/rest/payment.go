package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/mailru/easyjson"
)

var urlPay = os.Getenv("IP_PAYMENT")

func verifyPayment(client *http.Client, userId int) (bool, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8082/verify/%d", urlPay, userId), nil)
	if err != nil {
		return false, err
	}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	var dataPayment domain.PaymentData
	err = easyjson.UnmarshalFromReader(response.Body, &dataPayment)
	if err != nil {
		return false, err
	}

	return dataPayment.Payment.IsPay, nil
}
