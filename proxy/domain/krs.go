package domain

import (
	"net/http"

	"github.com/dinel13/thesis-ac/test/model"
)

type KrsHandlers interface {
	Read(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type KrsGrpcClients interface {
	ReadKrs(id int, token string) (*model.Krs, error)
	CreateKrs(krs *model.Krs, token string) (*model.Krs, error)
	UpdateKrs(krs *model.Krs, id int, token string) (*model.Krs, error)
	DeleteKrs(id int, token string) error
}
