package rest

import (
	"net/http"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/mailru/easyjson"
)

func ReadEasyJson(request *http.Request, data *domain.Krs) error {
	err := easyjson.UnmarshalFromReader(request.Body, data)
	if err != nil {
		return err
	}
	return nil
}

func WriteEasyJson(w http.ResponseWriter, status int, data *domain.Krs) error {
	js, err := easyjson.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func WriteEasyJsonError(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	logErorr(err)

	theError := domain.Error{
		Message: err.Error(),
	}

	js, _ := easyjson.Marshal(theError)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}
