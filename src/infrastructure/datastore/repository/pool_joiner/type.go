package pool_joiner

import (
	"github.com/jinzhu/gorm"
)

type PoolUserRelation struct {
	gorm.Model

	UserId int `gorm:"primary_key"`
	PoolId int `gorm:"primary_key"`

	DisplayName string

}
