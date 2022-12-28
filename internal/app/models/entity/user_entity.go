package entities

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey" gorm:"autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password string			`json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name string, email string) *User {
	return &User{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
