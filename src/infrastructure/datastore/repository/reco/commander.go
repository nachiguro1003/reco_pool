package reco_repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco-pool/src/domain/model"
)

func (r *Reco)Create(db *gorm.DB) (*model.Reco,error) {
	if err := db.Create(r).Error; err!= nil {
		return nil,err
	}
	return &r.Reco,nil
}


