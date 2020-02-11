package model

import (
	"github.com/jinzhu/gorm"
	"github.com/reco_pool/src/infrastructure/datastore/repository/pool"
	"github.com/reco_pool/src/infrastructure/datastore/repository/reco"
	"github.com/reco_pool/src/infrastructure/datastore/repository/user"
	"github.com/reco_pool/src/usecase/pool"
)

type (
	Pool struct {
		Pool *pool.PoolEntity `json:"pool"`

		Joiners []*user.UserEntity `json:"joiners"`
		Recos []*reco.RecoEntity `json:"recos"`
	}
)

func NewPoolService() (ps *Pool) {
	return ps
}

func (ps *Pool)CreateNewPool(db *gorm.DB) error {
	err := pool_service.Create(ps,db)
	if err != nil {
		return err
	}
	return nil
}
