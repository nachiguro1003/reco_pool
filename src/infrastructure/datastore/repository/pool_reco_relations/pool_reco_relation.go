package pool_reco_relations

import "time"

type PoolRecoRelation struct {
	PoolId int
	RecoId int

	CreatedAt time.Time
	UpdatedAt time.Time
}
