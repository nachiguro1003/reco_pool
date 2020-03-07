package model

type Reco struct {
	Title      string `json:"title"`
	ImageId    int    `json:"image_id"`
	Url        string `json:"url"`
	CategoryId int    `json:"category_id"`
	Status     int    `json:"status"`
	UserId     int    `json:"user_id"`
	PoolId     int    `json:"pool_id"`
}
