package rest

import (
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
	}

	// get the krs from service
	krs, err := h.service.Read(id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	WriteJson(w, http.StatusOK, krs, "krs")
}

// Create is handler for POST /krs to create COurse
func (h krsHandlers) Create(w http.ResponseWriter, r *http.Request) {
	// get the krs from request body
	krs := &domain.Krs{}
	err := ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

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
	// get the krs from request body
	krs := &domain.Krs{}
	err := ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

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

	// delete the krs
	err = h.service.Delete(id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	WriteJson(w, http.StatusOK, nil, "krs")
}
