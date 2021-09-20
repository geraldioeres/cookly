package users

import "time"

type Users struct {
	ID          int
	Name        string
	Email       string
	Password    string
	// Location 
	// date of birth
	Created_At  time.Time
	Modified_At time.Time
}
