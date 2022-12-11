package dto

import "time"

type Product struct {
	ID int64
	Name string
	Code string
	CreatedAt time.Time
	UpdatedAt time.Time
}
