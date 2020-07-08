package pool_service

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco-pool/src/domain/model"
)

type (
	PoolService struct {
		PoolRepository PoolServicer
	}

	PoolServicer interface {
		Get(db *gorm.DB) (*model.Pool,error)
		Create(db *gorm.DB) (*model.Pool,error)
	}
)

func NewPoolService(repo PoolServicer) *PoolService {
	return &PoolService{
		PoolRepository: repo,
	}
}
