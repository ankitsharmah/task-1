package models


type Club struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	ClubName        string `json:"clubname"`
	Address string `json:"address"`
	
}