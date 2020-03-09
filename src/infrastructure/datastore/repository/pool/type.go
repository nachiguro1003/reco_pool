package pool_repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco_pool/src/domain/model"
)

type Pool struct {
	gorm.Model

	model.Pool
}

func NewPoolRepository() *Pool{
	return &Pool{}
}
