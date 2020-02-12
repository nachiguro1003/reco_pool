package reco

import (
	"github.com/jinzhu/gorm"
)

type Reco struct {
	gorm.Model

	Title string `json:"title"`
	ImageId int `json:"image_id"`
	Url string `json:"url"`
	CategoryId int `json:"category_id"`
	Status int `json:"status"`

}

