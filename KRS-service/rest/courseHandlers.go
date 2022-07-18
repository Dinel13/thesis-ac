package rest

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/julienschmidt/httprouter"
)

func NewCoursRestHandlers(s domain.KrsService) domain.KrsRestHandlers {
	authClient := &http.Client{Timeout: 15 * time.Second}
	payClient := &http.Client{Timeout: 15 * time.Second}
	return krsHandlers{s, authClient, payClient}
}

type krsHandlers struct {
	service    domain.KrsService
	authClient *http.Client
	payClient  *http.Client
}

// GetKrs is handler for GET /krs/{id}
func (h krsHandlers) Read(w http.ResponseWriter, r *http.Request) {
	// get the krs id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// GET TOKEN FROM HEADER USE BEARER TOKEN
	token := r.Header.Get("Authorization")

	isAuth, err := verifyToken(h.authClient, token)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusBadRequest)
		return
	}
	if !isAuth {
		WriteEasyJsonError(w, errors.New("token tidak valid"), http.StatusBadRequest)
		return
	}

	isPay, err := verifyPayment(h.payClient, id)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if !isPay {
		WriteEasyJsonError(w, errors.New("belum melakukan pembayaran"), http.StatusBadRequest)
		return
	}

	// get the krs from service
	krs, err := h.service.Read(token, id)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	WriteEasyJson(w, http.StatusOK, krs)
}

// Create is handler for POST /krs to create COurse
func (h krsHandlers) Create(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	isAuth, err := verifyToken(h.authClient, token)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusBadRequest)
		return
	}
	if !isAuth {
		WriteEasyJsonError(w, errors.New("token tidak valid"), http.StatusBadRequest)
		return
	}

	// get the krs from request body
	krs := &domain.Krs{}
	err = ReadEasyJson(r, krs)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	isPay, err := verifyPayment(h.payClient, krs.IdMahasiswa)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if !isPay {
		WriteEasyJsonError(w, errors.New("belum melakukan pembayaran"), http.StatusBadRequest)
		return
	}

	// create the krs
	krsCreated, err := h.service.Create(krs)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	WriteEasyJson(w, http.StatusCreated, krsCreated)
}

// Update is handler for PUT /krs to update Krs
func (h krsHandlers) Update(w http.ResponseWriter, r *http.Request) {
	// get the krs id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")

	isAuth, err := verifyToken(h.authClient, token)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusBadRequest)
		return
	}
	if !isAuth {
		WriteEasyJsonError(w, errors.New("token tidak valid"), http.StatusBadRequest)
		return
	}

	isPay, err := verifyPayment(h.payClient, id)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if !isPay {
		WriteEasyJsonError(w, errors.New("belum melakukan pembayaran"), http.StatusBadRequest)
		return
	}

	// get the krs from request body
	krs := &domain.Krs{}
	err = ReadEasyJson(r, krs)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}
	krs.IdMahasiswa = id

	// update the krs
	c, err := h.service.Update(krs)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	WriteEasyJson(w, http.StatusOK, c)
}

// Delete is handler for DELETE /krs/{id} to delete Krs
func (h krsHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	// get the krs id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")
	isAuth, err := verifyToken(h.authClient, token)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusBadRequest)
		return
	}
	if !isAuth {
		WriteEasyJsonError(w, errors.New("token tidak valid"), http.StatusBadRequest)
		return
	}

	isPay, err := verifyPayment(h.payClient, id)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if !isPay {
		WriteEasyJsonError(w, errors.New("belum melakukan pembayaran"), http.StatusBadRequest)
		return
	}

	// delete the krs
	err = h.service.Delete(token, id)
	if err != nil {
		WriteEasyJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
