package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/db/postgresql"
)

type RecoPoolDBManager struct {
	DBs []DBManager
}

type DBManager interface {
	Open() error
	Close() error
}

func RecoPoolDbList() []DBManager{
	postgres := postgresql.Init()

	pdb,_ := interface{}(postgres).(DBManager)

	return []DBManager{
		pdb,
	}
}

// Init is initialize db from main function
func Init() (*RecoPoolDBManager,error) {
	dbs := RecoPoolDbList()
	for _,v := range dbs {
		if err := v.Open();err != nil {
			return nil,err
		}
	}
	return &RecoPoolDBManager{
		DBs:dbs,
	},nil
}

// Close is closing db
func (rdbm *RecoPoolDBManager) Close() {
	for _,v := range rdbm.DBs {
		if err:= v.Close(); err != nil {
			panic(err)
		}
	}
}
