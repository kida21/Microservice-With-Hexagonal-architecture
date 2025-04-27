package domain

import "time"

type UserModel struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Creadted_at time.Time `json:"created_at"`
}