package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/model"
	"github.com/dinel13/thesis-ac/test/rest-outbound"
	"github.com/julienschmidt/httprouter"
)

func NewRestKrsHandlers(ip string) domain.KrsHandlers {
	krsClient := &http.Client{Timeout: 15 * time.Second}
	return &RestKrs{
		ip:        ip,
		krsClient: krsClient,
	}
}

type RestKrs struct {
	ip        string
	krsClient *http.Client
}

func (k *RestKrs) Read(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")

	krs, err := rest.ReadKrs(k.krsClient, k.ip, token, id)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteKrsJson(w, http.StatusOK, &model.KrsResponse{Krs: krs.Krs})
}

func (k *RestKrs) Create(w http.ResponseWriter, r *http.Request) {
	krs := &model.Krs{}
	err := ReadKrsJson(r, krs)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}
	token := r.Header.Get("Authorization")

	createdkrs, err := rest.CreateKrs(k.krsClient, krs, k.ip, token)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteKrsJson(w, http.StatusOK, &model.KrsResponse{Krs: createdkrs.Krs})
}

func (k *RestKrs) Update(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	krs := &model.Krs{}
	err = ReadKrsJson(r, krs)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}
	token := r.Header.Get("Authorization")

	updatedkrs, err := rest.UpdateKrs(k.krsClient, krs, k.ip, token, id)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteKrsJson(w, http.StatusOK, &model.KrsResponse{Krs: updatedkrs.Krs})
}

func (k *RestKrs) Delete(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")

	err = rest.DeleteKrs(k.krsClient, k.ip, token, id)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteDelKrsEasyJson(w, http.StatusOK, &model.ResKrsDelete{Status: "succes"})
}
