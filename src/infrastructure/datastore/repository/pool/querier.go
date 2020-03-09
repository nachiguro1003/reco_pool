package pool_repository

import (
	"github.com/jinzhu/gorm"
	 "github.com/nachiguro1003/reco_pool/src/domain/model"
)

func (p *Pool)Get(db *gorm.DB) (*model.Pool,error) {
	if err := db.Find(p,"slug = ?",p.Slug).Error; err != nil {
		return nil,err
	}
	return &p.Pool,nil
}
