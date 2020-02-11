package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/reco_pool/src/domain/model"
	"github.com/reco_pool/src/infrastructure/datastore/repository/category"
	"github.com/reco_pool/src/infrastructure/datastore/repository/pool_reco_relations"
	"github.com/reco_pool/src/infrastructure/datastore/repository/pool_user_relation"
	"github.com/reco_pool/src/infrastructure/datastore/repository/reco"
	"github.com/reco_pool/src/infrastructure/datastore/repository/user"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=postgres dbname=reco-pool password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(user.User{})
	db.AutoMigrate(model.Pool{})
	db.AutoMigrate(category.Category{})
	db.AutoMigrate(reco.Reco{})
	db.AutoMigrate(pool_reco_relations.PoolRecoRelation{})
	db.AutoMigrate(pool_user_relation.PoolUserRelation{})
}