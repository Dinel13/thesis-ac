package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type resData struct {
	IsAuth bool `json:"isAuth"`
}

func VerifyToken(token string) (bool, error) {
	ip := os.Getenv("IP_AUTH")

	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s:8081/verify", ip), nil)
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
