package domain

import (
	"context"
	"net/http"

	"github.com/dinel13/thesis-ac/krs/proto"
)

type Krs struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Teacher     int    `json:"teacher"`
	Capacity    int    `json:"capacity"`
	Remain      int    `json:"remain"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type KrsRepository interface {
	Create(context.Context, *Krs) (*Krs, error)
	Read(context.Context, int) (*Krs, error)
	Update(context.Context, *Krs) (*Krs, error)
	Delete(context.Context, int) error
}

type KrsService interface {
	Read(int) (*Krs, error)
	Create(*Krs) (*Krs, error)
	Update(*Krs) (*Krs, error)
	Delete(int) error
}

type KrsRestHandlers interface {
	Read(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type KrsGrpcHandler interface {
	Read(context.Context, *proto.ReadKRSRequest) (*proto.KRSResponse, error)
	Create(context.Context, *proto.CreateUpdateKRSRequest) (*proto.KRSResponse, error)
	Update(context.Context, *proto.CreateUpdateKRSRequest) (*proto.KRSResponse, error)
	Delete(context.Context, *proto.DeleteKRSRequest) (*proto.KRSResponse, error)
}
