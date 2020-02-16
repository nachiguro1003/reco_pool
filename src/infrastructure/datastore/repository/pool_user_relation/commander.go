package pool_user_relation

import "github.com/jinzhu/gorm"

func(pu *PoolUserRelation)Create(db *gorm.DB) error {
	return db.Create(&pu).Error
}
