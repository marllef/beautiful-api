package dto

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey" gorm:"autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
