package user

import (
	"time"
)

// User defines the structure of a user record
type User struct {
	username  string
	createdAt time.Time
}

var users map[string]*User

func init() {
	users = make(map[string]*User)
}
