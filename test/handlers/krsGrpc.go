package handlers

import (
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/rest"
	"github.com/julienschmidt/httprouter"
)

func NewGrpcKrsHandlers(ip string) domain.KrsHandlers {
	return &grpcKrs{
		ip: ip,
	}
}

type grpcKrs struct {
	ip string
}

func (k *grpcKrs) Read(w http.ResponseWriter, r *http.Request) {
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

	WriteJson(w, http.StatusOK, krs, "krs")
}

func (k *grpcKrs) Create(w http.ResponseWriter, r *http.Request) {
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

	WriteJson(w, http.StatusOK, createdkrs, "krs")
}

func (k *grpcKrs) Update(w http.ResponseWriter, r *http.Request) {
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

	WriteJson(w, http.StatusOK, updatedkrs, "krs")
}

func (k *grpcKrs) Delete(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

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
