package pool_repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco-pool/src/domain/model"
)

func (u *Pool)Create(db *gorm.DB) (*model.Pool,error) {
	if err := db.Create(u).Error; err!= nil {
		return nil,err
	}
	return &u.Pool,nil
}


