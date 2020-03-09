package pool_service

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco_pool/src/domain/model"
)

func (ps *PoolService)GetPool(db *gorm.DB) (pool *model.Pool, err error) {
	pool,err = ps.PoolRepository.Get(db)
	if err != nil {
		return
	}

	return
}
