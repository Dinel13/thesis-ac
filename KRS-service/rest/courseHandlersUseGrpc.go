package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dinel13/thesis-ac/krs/domain"
	"github.com/dinel13/thesis-ac/krs/grpc"
	"github.com/dinel13/thesis-ac/krs/proto"
	"github.com/julienschmidt/httprouter"
)

func NewCoursRestGrpcHandlers(cu proto.AuthServiceClient, cp proto.PaymentServiceClient, s domain.KrsService) domain.KrsRestGrpcHandlers {
	return krsGrpcHandlers{clientAuth: cu, clientPay: cp, service: s}
}

type krsGrpcHandlers struct {
	clientAuth proto.AuthServiceClient
	clientPay  proto.PaymentServiceClient
	service    domain.KrsService
}

// GetKrs is handler for GET /krs/{id}
func (h krsGrpcHandlers) Read(w http.ResponseWriter, r *http.Request) {
	// get the krs id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")
	ctx := r.Context()

	token = token[7:]
	isAuth, err := grpc.VerifyToken(ctx, h.clientAuth, token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}
	if !isAuth {
		WriteJsonError(w, errors.New("token tidak valid"), http.StatusBadRequest)
		return
	}

	isPay, err := grpc.VerifyPayment(ctx, h.clientPay, id)
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
func (h krsGrpcHandlers) Create(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	token = token[7:]

	ctx := r.Context()
	isAuth, err := grpc.VerifyToken(ctx, h.clientAuth, token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}
	if !isAuth {
		WriteJsonError(w, errors.New("token tidak valid"), http.StatusBadRequest)
		return
	}

	// get the krs from request body
	krs := &domain.Krs{}
	err = ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	isPay, err := grpc.VerifyPayment(ctx, h.clientPay, krs.IdMahasiswa)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if !isPay {
		WriteJsonError(w, errors.New("belum melakukan pembayaran"), http.StatusBadRequest)
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
func (h krsGrpcHandlers) Update(w http.ResponseWriter, r *http.Request) {
	// get the krs id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")
	ctx := r.Context()
	token = token[7:]

	isAuth, err := grpc.VerifyToken(ctx, h.clientAuth, token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}
	if !isAuth {
		WriteJsonError(w, errors.New("token tidak valid"), http.StatusBadRequest)
		return
	}

	isPay, err := grpc.VerifyPayment(ctx, h.clientPay, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if !isPay {
		WriteJsonError(w, errors.New("belum melakukan pembayaran"), http.StatusBadRequest)
		return
	}

	// get the krs from request body
	krs := &domain.Krs{}
	err = ReadJson(r, krs)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	krs.IdMahasiswa = id

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
func (h krsGrpcHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	// get the krs id from request param
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")
	ctx := r.Context()
	token = token[7:]

	isAuth, err := grpc.VerifyToken(ctx, h.clientAuth, token)
	if err != nil {
		WriteJsonError(w, err, http.StatusBadRequest)
		return
	}
	if !isAuth {
		WriteJsonError(w, errors.New("token tidak valid"), http.StatusBadRequest)
		return
	}

	isPay, err := grpc.VerifyPayment(ctx, h.clientPay, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}
	if !isPay {
		WriteJsonError(w, errors.New("belum melakukan pembayaran"), http.StatusBadRequest)
		return
	}

	// delete the krs
	err = h.service.Delete(token, id)
	if err != nil {
		WriteJsonError(w, err, http.StatusInternalServerError)
		return
	}

	// write the krs to response
	WriteJson(w, http.StatusOK, nil, "krs")
}
