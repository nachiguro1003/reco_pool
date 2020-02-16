package pool_user_relation

import (
	"github.com/jinzhu/gorm"
)

type PoolUserRelation struct {
	gorm.Model

	PoolId int `gorm:"primary_key"`
	UserId int `gorm:"primary_key"`

}
