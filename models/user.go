package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID       int       `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"-" db:"password"` // "-" means this field won't be included in JSON
	DoB      time.Time `json:"dob" db:"dob"`
}

// For login/register requests
type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UnmarshalJSON implements custom JSON unmarshaling for User
func (u *User) UnmarshalJSON(data []byte) error {
	// Create an auxiliary type to avoid recursive UnmarshalJSON calls
	type Aux struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password,omitempty"`
		DoB      string `json:"dob"`
	}
	var aux Aux
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	u.ID = aux.ID
	u.Name = aux.Name
	u.Email = aux.Email
	u.Password = aux.Password

	// Parse the date string
	if aux.DoB != "" {
		date, err := time.Parse("2006-01-02", aux.DoB)
		if err != nil {
			return err
		}
		u.DoB = date
	}

	return nil
}

type Post struct {
	ID        int       `json:"id" db:"id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
