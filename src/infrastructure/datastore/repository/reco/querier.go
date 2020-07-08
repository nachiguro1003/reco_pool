package reco_repository

import (
	"github.com/jinzhu/gorm"
	 "github.com/nachiguro1003/reco-pool/src/domain/model"
)

func (r *Reco)Get(db *gorm.DB) (*model.Reco,error) {
	if err := db.Find(r,"id = ?",r.ID).Error; err != nil {
		return nil,err
	}
	return &r.Reco,nil
}
