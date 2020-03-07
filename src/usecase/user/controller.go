package user_service

import (
	"github.com/jinzhu/gorm"
)

func (us *UserService)CreatePool(db *gorm.DB) error {
	if _,err := us.User.Create(db); err != nil {
		return err
	}
	return nil
}

