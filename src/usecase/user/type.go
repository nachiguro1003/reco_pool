package user_service

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco-pool/src/domain/model"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/repository/user"
)

type (
	UserService struct {
		User UserServicer
	}

	UserServicer interface {
		Get(db *gorm.DB) (*model.User,error)
		Create(db *gorm.DB) (*model.User,error)
	}
)

func NewUserService(repo *user_repository.User) *UserService {
	return &UserService{
		User: repo,
	}
}
