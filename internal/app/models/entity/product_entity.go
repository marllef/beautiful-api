package entities

import "time"

type Product struct {
	ID        int64      `gorm:"primaryKey" gorm:"autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
