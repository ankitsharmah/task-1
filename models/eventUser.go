package models

type EventUser struct {
    EventID uint `gorm:"primaryKey" json:"event_id"`
    UserID  uint `gorm:"primaryKey" json:"user_id"`
}