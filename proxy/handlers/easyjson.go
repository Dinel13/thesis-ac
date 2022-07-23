package handlers

import (
	"net/http"

	"github.com/dinel13/thesis-ac/test/model"
	"github.com/mailru/easyjson"
)

func ReadKrsJson(request *http.Request, data *model.Krs) error {
	err := easyjson.UnmarshalFromReader(request.Body, data)
	if err != nil {
		return err
	}
	return nil
}

func ReadAuthJson(request *http.Request, data *model.LoginSignupRequest) error {
	err := easyjson.UnmarshalFromReader(request.Body, data)
	if err != nil {
		return err
	}
	return nil
}

func ReadPayJson(request *http.Request, data *model.PaymentRequest) error {
	err := easyjson.UnmarshalFromReader(request.Body, data)
	if err != nil {
		return err
	}
	return nil
}

func WriteKrsJson(w http.ResponseWriter, status int, data *model.KrsResponse) error {
	js, err := easyjson.Marshal(data)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func WriteDelKrsEasyJson(w http.ResponseWriter, status int, data *model.ResKrsDelete) error {
	js, err := easyjson.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func WriteVerifyJson(w http.ResponseWriter, status int, data *model.VerifyTokenResponse) error {
	js, err := easyjson.Marshal(data)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func WriteSignupJson(w http.ResponseWriter, status int, data *model.LoginSignupResponse) error {
	js, err := easyjson.Marshal(data)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func WritePayJson(w http.ResponseWriter, status int, data *model.PaymentWraperResponse) error {
	js, err := easyjson.Marshal(data)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return nil
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

	theError := model.Error{
		Message: err.Error(),
	}

	js, _ := easyjson.Marshal(theError)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}
