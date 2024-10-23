package models

import "time"

type User struct {
	Id        int32
	Email     string
	UserName  string
	Password  string
	CreatedAt time.Time
}
