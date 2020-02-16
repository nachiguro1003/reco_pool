package reco_repository

import (
	"github.com/jinzhu/gorm"
	"github.com/reco_pool/src/domain/model"
)

type RecoRepository struct {
	gorm.Model

	model.Reco
}

