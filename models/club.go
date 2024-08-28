package models

// import ("")


type Club struct {
    ID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
    ClubName string  `gorm:"type:varchar(100);not null" json:"club_name"`
    Address  string  `gorm:"type:varchar(255)" json:"address"`
    Rating   string   `gorm:"type:varchar(255)" json:"rating"`
    Description string `gorm:"type:varchar(255)" json:"description"`
    Events   []Event `gorm:"foreignKey:ClubID" json:"events"`
}

