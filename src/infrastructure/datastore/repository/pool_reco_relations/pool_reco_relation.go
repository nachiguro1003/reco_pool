package pool_reco_relations

import (
	"github.com/jinzhu/gorm"
)

type PoolRecoRelation struct {
	gorm.Model

	PoolId int `gorm:"primary_key"`
	RecoId int `gorm:"primary_key"`

}
