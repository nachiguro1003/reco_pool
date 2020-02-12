package pool_service

import (
	"github.com/jinzhu/gorm"
	"github.com/reco_pool/src/domain/model"
)

func Create(pool *model.Pool,db *gorm.DB) error {
	if err := pool.Create(db); err != nil {
		return err
	}
	return nil
}
