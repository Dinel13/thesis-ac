package handlers

import (
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/test/domain"
	"github.com/dinel13/thesis-ac/test/model"
	"github.com/julienschmidt/httprouter"
)

func NewGrpcKrsHandlers(kgc domain.KrsGrpcClients) domain.KrsHandlers {
	return &grpcKrs{
		kgc,
	}
}

type grpcKrs struct {
	kgc domain.KrsGrpcClients
}

func (k *grpcKrs) Read(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")[7:]

	krs, err := k.kgc.ReadKrs(id, token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, krs, "krs")
}

func (k *grpcKrs) Create(w http.ResponseWriter, r *http.Request) {
	krs := &model.Krs{}
	err := ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	token := r.Header.Get("Authorization")[7:]

	createdkrs, err := k.kgc.CreateKrs(krs, token)
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

	krs := &model.Krs{}
	err = ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	token := r.Header.Get("Authorization")[7:]

	updatedkrs, err := k.kgc.UpdateKrs(krs, id, token)
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

	token := r.Header.Get("Authorization")[7:]

	err = k.kgc.DeleteKrs(id, token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}

	WriteJson(w, http.StatusOK, nil, "krs")
}
