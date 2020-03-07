package user_service

import (
	"github.com/jinzhu/gorm"
)

func (us *UserService)GetPool(db *gorm.DB) error {
	_,err := us.User.Get(db)
	if err != nil {
		return err
	}

	return nil
}
