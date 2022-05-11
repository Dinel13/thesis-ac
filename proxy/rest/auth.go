package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dinel13/thesis-ac/test/domain"
)

func VerifyToken(ip, token string) (*domain.VerifyTokenResponse, error) {
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/verify", ip), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", token)

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

	resBody := &domain.VerifyTokenResponse{}
	err = json.NewDecoder(response.Body).Decode(resBody)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func Login(req *domain.LoginSignupRequest, ip string) (*domain.LoginSignupResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/login", ip), bytes.NewBuffer(body))
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

	resBody := &domain.LoginSignupResponse{}
	err = json.NewDecoder(response.Body).Decode(resBody)
	if err != nil {
		return nil, err
	}
	if resBody.Token == "" {
		return nil, fmt.Errorf("gagal login, perikasa username dan password")
	}

	return resBody, nil
}
