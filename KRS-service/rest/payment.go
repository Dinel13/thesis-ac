package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type payment struct {
	IsPay bool `json:"isPay"`
}

type resDataPayment struct {
	Payment payment `json:"payment"`
}

func VerifyPayment(userId int) (bool, error) {
	var client = &http.Client{}

	ip := os.Getenv("IP_PAYMENT")

	request, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8082/verify/%d", ip, userId), nil)
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
