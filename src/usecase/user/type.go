package user_service

import (
	"github.com/jinzhu/gorm"
	"github.com/reco_pool/src/domain/model"
	"github.com/reco_pool/src/infrastructure/datastore/repository/user"
)

type (
	UserService struct {
		UserRepository UserServicer
	}

	UserServicer interface {
		Get(db *gorm.DB) (*model.User,error)
		Create(db *gorm.DB) (*model.User,error)
	}
)

func NewUserService() *UserService {
	return &UserService{
		UserRepository: user_repository.NewUserRepository(),
	}
}
