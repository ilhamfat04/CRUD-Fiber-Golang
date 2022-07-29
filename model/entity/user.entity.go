package entity

import "time"

type User struct {
	ID        uint      `gorm:"type: int" json:"id"`
	Name      string    `gorm:"type: varchar(255)" json:"name"`
	Email     string    `gorm:"type: varchar(255)" json:"email"`
	Phone     string    `gorm:"type: varchar(255)" json:"phone "`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
