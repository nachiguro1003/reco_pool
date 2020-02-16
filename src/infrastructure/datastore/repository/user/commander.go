package user_repository

import (
	"github.com/jinzhu/gorm"
	"github.com/reco_pool/src/domain/model"
)

func (u *UserRepository)Create(db *gorm.DB) (*model.User,error) {
	if err := db.Create(&u).Error; err != nil {
		return nil,err
	}
	return &u.User,nil
}
