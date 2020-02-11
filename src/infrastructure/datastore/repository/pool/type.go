package pool

import (
	"github.com/jinzhu/gorm"
	"time"
)

type PoolEntity struct {
	db *gorm.DB

	Id int `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug" uri:"slug"`
	IsPublic bool `json:"is_public"`
	Password string `json:"password"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

