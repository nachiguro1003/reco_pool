package pool

import (
	"github.com/jinzhu/gorm"
)

func (p *Pool)GetBySlug(db *gorm.DB) error {
	return db.Find(p,"slug = ?",p.Slug).Error
}
