package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type resData struct {
	IsAuth bool `json:"isAuth"`
}

var url = os.Getenv("URL_AUTH")

func VerifyToken(token string) (bool, error) {
	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s:8081/verify", url), nil)
	if err != nil {
		return false, err
	}
	request.Header.Add("Authorization", token)

	// send req
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
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
