package rest

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/dinel13/thesis-ac/test/model"
	"github.com/mailru/easyjson"
)

func VerifyToken(client *http.Client, ip, token string) (*model.VerifyTokenResponse, error) {
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/verify", ip), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", token)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer response.Body.Close()

	resBody := &model.VerifyTokenResponse{}
	err = easyjson.UnmarshalFromReader(response.Body, resBody)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func Login(client *http.Client, req *model.LoginSignupRequest, ip string) (*model.LoginSignupResponse, error) {
	body, err := easyjson.Marshal(req)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/login", ip), bytes.NewBuffer(body))
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

	resBody := &model.LoginSignupResponse{}
	err = easyjson.UnmarshalFromReader(response.Body, resBody)
	if err != nil {
		return nil, err
	}
	if resBody.Token == "" {
		return nil, fmt.Errorf("gagal login, perikasa username dan password")
	}

	return resBody, nil
}
