package user_service

import (
	"github.com/jinzhu/gorm"
)

func (us *UserService)GetPool(db *gorm.DB) error {
	err := us.UserRepository.Get(db)
	if err != nil {
		return err
	}

	return nil
}
