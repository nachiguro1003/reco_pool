package user_repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco_pool/src/domain/model"
)

type User struct {
	gorm.Model

	model.User
}

func NewUserRepository() *User {
	return &User{}
}
