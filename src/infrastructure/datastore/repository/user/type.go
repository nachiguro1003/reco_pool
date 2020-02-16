package user_repository

import (
	"github.com/jinzhu/gorm"
	"github.com/reco_pool/src/domain/model"
)

type UserRepository struct {
	gorm.Model

	model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
