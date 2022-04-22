package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dinel13/thesis-ac/test/domain"
)

func ReadKrs(ip, token string, id int) (*domain.KrsResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8080/krs/%d", ip, id), nil)
	if err != nil {
		return nil, err
	}
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

	resBody := &domain.KrsResponse{}
	err = json.NewDecoder(response.Body).Decode(resBody)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func CreateKrs(krs *domain.Krs, ip, token string) (*domain.KrsResponse, error) {
	body, err := json.Marshal(krs)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s:8080/krs", ip), bytes.NewBuffer(body))
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

	resBody := &domain.KrsResponse{}
	err = json.NewDecoder(response.Body).Decode(resBody)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func UpdateKrs(krs *domain.Krs, ip, token string, id int) (*domain.KrsResponse, error) {
	body, err := json.Marshal(krs)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("PUT", fmt.Sprintf("http://%s:8080/krs/%d", ip, id), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
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

	resBody := &domain.KrsResponse{}
	err = json.NewDecoder(response.Body).Decode(resBody)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func DeleteKrs(ip, token string, id int) error {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:8080/krs/%d", ip, id), nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", token)

	// send req
	client := &http.Client{
		Timeout: time.Duration(15) * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error make request ", err)
		return err
	}

	defer response.Body.Close()

	resBody := &domain.ResKrsDelete{}
	err = json.NewDecoder(response.Body).Decode(resBody)
	if err != nil {
		log.Println("error decoded", err)
		return err
	}

	return nil
}
