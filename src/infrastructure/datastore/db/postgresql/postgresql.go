package postgresql

import (
	"github.com/jinzhu/gorm"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/repository/category"
	pool_repository "github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/repository/pool"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/repository/pool_joiner"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/repository/pool_reco_relations"
	reco_repository "github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/repository/reco"
	user_repository "github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/repository/user"
)

var (
	postgresDB *gorm.DB
	err error
)
type PostgresManager struct {
	DB *gorm.DB
}

func(pm *PostgresManager) Open() error {
	// todo configから持ってきたい

	postgresDB,err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=reco-pool password=postgres sslmode=disable")
	postgresDB.AutoMigrate(user_repository.User{})
	postgresDB.AutoMigrate(pool_repository.Pool{})
	postgresDB.AutoMigrate(category.Category{})
	postgresDB.AutoMigrate(reco_repository.Reco{})
	postgresDB.AutoMigrate(pool_reco_relations.PoolRecoRelation{})
	postgresDB.AutoMigrate(pool_joiner.PoolUserRelation{})

	return err
}

func (pm *PostgresManager) Close() error {
	return pm.DB.Close()
}

func Init() *PostgresManager {
	return &PostgresManager{}
}

func GetPostgresInstance() *PostgresManager {
	return &PostgresManager{DB:postgresDB}
}
