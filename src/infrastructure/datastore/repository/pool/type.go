package pool

import (
	"github.com/jinzhu/gorm"
)

type Pool struct {
	gorm.Model

	Name string `json:"name" gorm:"type:VARCHAR(100)"`
	Slug string `json:"slug" uri:"slug" gorm:"unique;not null;unique_index"`
	IsPublic bool `json:"is_public"gorm:"default:true"`
	Password string `json:"password" gorm:"type:VARCHAR(100)"`

}

