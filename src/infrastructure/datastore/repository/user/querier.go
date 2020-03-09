package user_repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco_pool/src/domain/model"
)

func (u *User)Get(db *gorm.DB) (*model.User,error) {
	if err := db.Find(u,"id = ?",u.ID).Error; err != nil {
		return nil,err
	}
	return &u.User,nil
}
