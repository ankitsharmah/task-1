package models

// import "time"

type Event struct {
    ID      uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Name    string    `gorm:"type:varchar(100);not null" json:"name"`
    ClubID  uint      `gorm:"not null" json:"club_id"`
    // Date    time.Time `gorm:"type:datetime" json:"date"`
    Users   []User    `gorm:"many2many:event_users" json:"users"`
}
