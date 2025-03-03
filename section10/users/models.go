package users

import "time"

type User struct {
	Id             int64
	Username       string
	Password       string
	DateRegistered time.Time
}
