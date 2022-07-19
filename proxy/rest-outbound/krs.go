package rest

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/dinel13/thesis-ac/test/model"
	"github.com/mailru/easyjson"
)

func ReadKrs(client *http.Client, ip, token string, id int) (*model.KrsResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8080/krs/%d", ip, id), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", token)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer response.Body.Close()

	resBody := &model.KrsResponse{}
	err = easyjson.UnmarshalFromReader(response.Body, resBody)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func CreateKrs(client *http.Client, krs *model.Krs, ip, token string) (*model.KrsResponse, error) {
	body, err := easyjson.Marshal(krs)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s:8080/krs", ip), bytes.NewBuffer(body))
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

	resBody := &model.KrsResponse{}
	err = easyjson.UnmarshalFromReader(response.Body, resBody)
	if err != nil {
		return nil, err
	}

	log.Println(resBody.Krs.IdMahasiswa)
	return resBody, nil
}

func UpdateKrs(client *http.Client, krs *model.Krs, ip, token string, id int) (*model.KrsResponse, error) {
	body, err := easyjson.Marshal(krs)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("PUT", fmt.Sprintf("http://%s:8080/krs/%d", ip, id), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", token)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer response.Body.Close()

	resBody := &model.KrsResponse{}
	err = easyjson.UnmarshalFromReader(response.Body, resBody)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func DeleteKrs(client *http.Client, ip, token string, id int) error {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:8080/krs/%d", ip, id), nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", token)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error make request ", err)
		return err
	}

	defer response.Body.Close()

	resBody := &model.ResKrsDelete{}
	err = easyjson.UnmarshalFromReader(response.Body, resBody)
	if err != nil {
		log.Println("error decoded", err)
		return err
	}

	return nil
}
