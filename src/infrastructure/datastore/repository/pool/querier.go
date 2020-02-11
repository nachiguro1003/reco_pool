package pool

import (
	"github.com/jinzhu/gorm"
)

func (u *PoolEntity)GetBySlug(slug string, db *gorm.DB) error {
	return db.Find(&u,"slug = ?",slug).Error
}
