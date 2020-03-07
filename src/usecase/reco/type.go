package reco_service

import (
	"github.com/jinzhu/gorm"
	"github.com/reco_pool/src/domain/model"
)

type (
	RecoService struct {
		RecoRepository RecoServicer
	}

	RecoServicer interface {
		Get(db *gorm.DB) (*model.Reco,error)
		Create(db *gorm.DB) (*model.Reco,error)
	}
)

func NewRecoService(repo RecoServicer) *RecoService {
	return &RecoService{
		RecoRepository: repo,
	}
}
