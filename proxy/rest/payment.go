package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dinel13/thesis-ac/test/domain"
)

func VerifyPayment(ip string, id int) (*domain.PaymentResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8082/verify/%d", ip, id), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	// send req
	client := &http.Client{
		Timeout: time.Duration(15) * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer response.Body.Close()

	resBody := &domain.PaymentWraperResponse{}
	err = json.NewDecoder(response.Body).Decode(resBody)
	if err != nil {
		return nil, err
	}

	return &resBody.PaymentResponse, nil
}

func Pay(req *domain.PaymentRequest, ip string) (*domain.PaymentResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s:8082/pay", ip), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	// send req
	client := &http.Client{
		Timeout: time.Duration(15) * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer response.Body.Close()

	resBody := &domain.PaymentWraperResponse{}
	err = json.NewDecoder(response.Body).Decode(resBody)
	if err != nil {
		return nil, err
	}

	return &resBody.PaymentResponse, nil
}
