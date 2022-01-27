package rest

import (
	"encoding/json"
	"net/http"
)

type resData struct {
	IsAuth bool `json:"isAuth"`
}

func VerifyToken(token string) (bool, error) {
	request, err := http.NewRequest("POST", "http://localhost:8081/verify", nil)
	if err != nil {
		return false, err
	}

	request.Header.Add("Authorization", token)

	// send req
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	var data resData
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return false, err
	}

	return data.IsAuth, nil
}
