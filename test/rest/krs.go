package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dinel13/thesis-ac/test/domain"
)

func ReadKrs(ip, token string, id int) (*domain.Krs, error) {
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

	krs := &domain.Krs{}
	err = json.NewDecoder(response.Body).Decode(krs)
	if err != nil {
		return nil, err
	}

	return krs, nil
}

func CreateKrs(krs *domain.Krs, ip, token string) (*domain.Krs, error) {
	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s:8080/krs", ip), nil)
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

	createdKrs := &domain.Krs{}
	err = json.NewDecoder(response.Body).Decode(createdKrs)
	if err != nil {
		return nil, err
	}

	return createdKrs, nil
}

func UpdateKrs(krs *domain.Krs, ip, token string, id int) (*domain.Krs, error) {
	request, err := http.NewRequest("PUT", fmt.Sprintf("http://%s:8080/krs/%d", ip, id), nil)
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

	updatedKrs := &domain.Krs{}
	err = json.NewDecoder(response.Body).Decode(updatedKrs)
	if err != nil {
		return nil, err
	}

	return updatedKrs, nil
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
		fmt.Println(err)
		return err
	}

	defer response.Body.Close()

	delKrs := &domain.ResKrsDelete{}
	err = json.NewDecoder(response.Body).Decode(delKrs)
	if err != nil {
		return err
	}

	return nil
}
