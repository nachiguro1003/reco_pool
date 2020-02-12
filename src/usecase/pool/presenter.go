package pool_service

import (
	"github.com/jinzhu/gorm"
	"github.com/reco_pool/src/domain/model"
)

func GetPool(ps *model.Pool,db *gorm.DB) (*model.Pool,error) {
	err := ps.GetBySlug(db)
	if err != nil {
		return nil,err
	}

	return ps,err
}
