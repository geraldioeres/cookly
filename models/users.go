package models

import "time"

type Users struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Created_At  time.Time `json:"created_at"`
	Modified_At time.Time `json:"updated_at"`
}