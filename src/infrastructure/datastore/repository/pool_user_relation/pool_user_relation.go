package pool_user_relation

import "time"

type PoolUserRelation struct {
	PoolId int
	UserId int

	CreatedAt time.Time
	UpdatedAt time.Time
}
