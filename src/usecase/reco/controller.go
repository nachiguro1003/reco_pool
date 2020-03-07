package reco_service

import (
	"github.com/jinzhu/gorm"
)

func (re *RecoService)CreateReco(db *gorm.DB) error {
	if _,err := re.RecoRepository.Create(db); err != nil {
		return err
	}
	return nil
}

//func (ps *PoolService)RegisterUserIntoPool(slug string,db *gorm.DB) error {
//	if err := db.Transaction(func(tx *gorm.DB) error {
//		ps,err := ps.PoolRepository.Get(tx)
//		if err != nil {
//			return err
//		}
//
//		u := user_service.NewUserService()
//
//		if err := u.UserRepository.Create(tx); err != nil {
//			return err
//		}
//
//		pu := &pool_user_relation.PoolUserRelation{
//			PoolId:int(ps),
//			UserId:int(u.ID),
//		}
//		if err := pu.Create(tx); err != nil {
//			return err
//		}
//		return nil
//	});err != nil {
//		return err
//	}
//
//	return nil
//}
