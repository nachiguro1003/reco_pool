package user_service

import (
	"github.com/jinzhu/gorm"
)

func (us *UserService)CreatePool(db *gorm.DB) error {
	if err := us.UserRepository.Create(db); err != nil {
		return err
	}
	return nil
}

