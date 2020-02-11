package reco

import "time"

type RecoEntity struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ImageId int `json:"image_id"`
	Url string `json:"url"`
	CategoryId int `json:"category_id"`
	Status int `json:"status"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

