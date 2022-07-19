package handlers

import (
	"log"
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
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")

	krs, err := rest.ReadKrs(k.krsClient, k.ip, token, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, krs.Krs, "krs")
}

func (k *RestKrs) Create(w http.ResponseWriter, r *http.Request) {
	krs := &model.Krs{}
	err := ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	token := r.Header.Get("Authorization")

	log.Println(krs)
	createdkrs, err := rest.CreateKrs(k.krsClient, krs, k.ip, token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}
	log.Println(createdkrs.Krs.IdMahasiswa)
	WriteJson(w, http.StatusOK, createdkrs.Krs, "krs")
}

func (k *RestKrs) Update(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	krs := &model.Krs{}
	err = ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	token := r.Header.Get("Authorization")

	updatedkrs, err := rest.UpdateKrs(k.krsClient, krs, k.ip, token, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, updatedkrs.Krs, "krs")
}

func (k *RestKrs) Delete(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")

	err = rest.DeleteKrs(k.krsClient, k.ip, token, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, nil, "krs")
}
