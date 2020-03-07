package model

type PoolJoiner struct {
	UserId int `gorm:"primary_key"`
	PoolId int `gorm:"primary_key"`

	DisplayName string `json:"display_name"`
	IconId int `json:"icon_id"`
}
