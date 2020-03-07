package pool_joiner

import "github.com/jinzhu/gorm"

func(pu *PoolUserRelation)Create(db *gorm.DB) error {
	return db.Create(&pu).Error
}
