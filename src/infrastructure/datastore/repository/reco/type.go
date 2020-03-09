package reco_repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco_pool/src/domain/model"
)

type Reco struct {
	gorm.Model

	model.Reco
}

func NewRecoRepository() *Reco {
	return &Reco{}
}
