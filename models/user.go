package models

// import("clubApi/models")
type User struct {
    ID      uint    `gorm:"primaryKey;autoIncrement" json:"id"`
    Name    string  `gorm:"type:varchar(100);not null" json:"name"`
    Email   string  `gorm:"type:varchar(100);unique;not null" json:"email"`
    Phone   string  `gorm:"type:varchar(100);unique;not null" json:"phone"`
    Events  []Event `gorm:"many2many:event_users" json:"events"`
}