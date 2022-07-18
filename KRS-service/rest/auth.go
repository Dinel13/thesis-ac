package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/mailru/easyjson"
)

var url = os.Getenv("IP_AUTH")

func verifyToken(client *http.Client, token string) (bool, error) {
	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s:8081/verify", url), nil)
	if err != nil {
		return false, err
	}
	request.Header.Add("Authorization", token)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	defer response.Body.Close()

	var authData domain.AuthData
	err = easyjson.UnmarshalFromReader(response.Body, &authData)
	if err != nil {
		return false, err
	}

	return authData.IsAuth, nil
}
