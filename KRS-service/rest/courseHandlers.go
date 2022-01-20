package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/julienschmidt/httprouter"
)

func NewCoursRestHandlers(s domain.KrsService) domain.KrsRestHandlers {
	return krsHandlers{s}
}

type krsHandlers struct {
	service domain.KrsService
}

// GetKrs is handler for GET /krs/{id}
func (h krsHandlers) Read(w http.ResponseWriter, r *http.Request) {
	// get the krs id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// GET TOKEN FROM HEADER USE BEARER TOKEN
	token := r.Header.Get("Authorization")

	isAuth, err := VerifyToken(token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}
	if !isAuth {
		WriteJsonError(w, errors.New("token tidak valid"), http.StatusBadRequest)
		return
	}

	isPay, err := VerifyPayment(id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if !isPay {
		WriteJsonError(w, errors.New("belum melakukan pembayaran"), http.StatusBadRequest)
		return
	}

	// get the krs from service
	krs, err := h.service.Read(token, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	WriteJson(w, http.StatusOK, krs, "krs")
}

// Create is handler for POST /krs to create COurse
func (h krsHandlers) Create(w http.ResponseWriter, r *http.Request) {
	// GET TOKEN FROM HEADER USE BEARER TOKEN
	token := r.Header.Get("Authorization")

	// get the krs from request body
	krs := &domain.Krs{}
	err := ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	krs.Token = token

	// create the krs
	c, err := h.service.Create(krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	WriteJson(w, http.StatusCreated, c, "krs")
}

// Update is handler for PUT /krs to update Krs
func (h krsHandlers) Update(w http.ResponseWriter, r *http.Request) {
	// GET TOKEN FROM HEADER USE BEARER TOKEN
	token := r.Header.Get("Authorization")

	// get the krs from request body
	krs := &domain.Krs{}
	err := ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	krs.Token = token

	// update the krs
	c, err := h.service.Update(krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	WriteJson(w, http.StatusOK, c, "krs")
}

// Delete is handler for DELETE /krs/{id} to delete Krs
func (h krsHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	// get the krs id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// GET TOKEN FROM HEADER USE BEARER TOKEN
	token := r.Header.Get("Authorization")

	// delete the krs
	err = h.service.Delete(token, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	WriteJson(w, http.StatusOK, nil, "krs")
}
