package pool

import (
	"github.com/jinzhu/gorm"
)

func (u *Pool)Create(db *gorm.DB) error {
	return db.Create(u).Error
}


