package model

import (
	"github.com/reco_pool/src/infrastructure/datastore/repository/pool"
	"github.com/reco_pool/src/infrastructure/datastore/repository/reco"
	"github.com/reco_pool/src/infrastructure/datastore/repository/user"
)

type (
	PoolService struct {
		pool.Pool `json:"pool"`

		Joiners []*user.User `json:"joiners"`
		Recos []*reco.Reco `json:"recos"`
	}
)

func NewPoolService() *PoolService {
	return &PoolService{}
}
