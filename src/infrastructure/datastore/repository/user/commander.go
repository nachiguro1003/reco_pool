package user_repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco_pool/src/domain/model"
	"log"
)

func (u *User)Create(db *gorm.DB) (*model.User,error) {
	log.Print(u)
	if err := db.Create(&u).Error; err != nil {
		return nil,err
	}
	return &u.User,nil
}
