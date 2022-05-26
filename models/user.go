package models

import "time"

//User Data
type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at,omitempty"`
}
