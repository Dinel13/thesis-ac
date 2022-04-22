package handlers

import (
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/rest"
	"github.com/julienschmidt/httprouter"
)

func NewRestKrsHandlers(ip string) domain.KrsHandlers {
	return &RestKrs{
		ip: ip,
	}
}

type RestKrs struct {
	ip string
}

func (k *RestKrs) Read(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")

	krs, err := rest.ReadKrs(k.ip, token, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, krs.Krs, "krs")
}

func (k *RestKrs) Create(w http.ResponseWriter, r *http.Request) {
	krs := &domain.Krs{}
	err := ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	token := r.Header.Get("Authorization")

	createdkrs, err := rest.CreateKrs(krs, k.ip, token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, createdkrs.Krs, "krs")
}

func (k *RestKrs) Update(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	krs := &domain.Krs{}
	err = ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	token := r.Header.Get("Authorization")

	updatedkrs, err := rest.UpdateKrs(krs, k.ip, token, id)
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

	err = rest.DeleteKrs(k.ip, token, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, nil, "krs")
}
