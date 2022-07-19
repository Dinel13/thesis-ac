package rest

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/dinel13/thesis-ac/test/model"
	"github.com/mailru/easyjson"
)

func VerifyPayment(client *http.Client, ip string, id int) (*model.PaymentResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8082/verify/%d", ip, id), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer response.Body.Close()

	resBody := &model.PaymentWraperResponse{}
	err = easyjson.UnmarshalFromReader(response.Body, resBody)
	if err != nil {
		return nil, err
	}

	return &resBody.PaymentResponse, nil
}

func Pay(client *http.Client, req *model.PaymentRequest, ip string) (*model.PaymentResponse, error) {
	body, err := easyjson.Marshal(req)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s:8082/pay", ip), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer response.Body.Close()

	resBody := &model.PaymentWraperResponse{}
	err = easyjson.UnmarshalFromReader(response.Body, resBody)
	if err != nil {
		return nil, err
	}

	return &resBody.PaymentResponse, nil
}
