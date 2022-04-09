package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type payment struct {
	IsPay bool `json:"isPay"`
}

type resDataPayment struct {
	Payment payment `json:"payment"`
}

var urlPay = os.Getenv("IP_PAYMENT")

func verifyPayment(userId int) (bool, error) {
	var client = &http.Client{
		Timeout: time.Duration(15) * time.Second,
	}
	request, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8082/verify/%d", urlPay, userId), nil)
	if err != nil {
		return false, err
	}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	var dataPayment resDataPayment
	err = json.NewDecoder(response.Body).Decode(&dataPayment)
	if err != nil {
		return false, err
	}

	return dataPayment.Payment.IsPay, nil
}
