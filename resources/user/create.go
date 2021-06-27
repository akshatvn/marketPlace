package user

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Create creates a user with a given username unless it already exists.
func Create(args ...string) error {
	userName := strings.ToLower(args[0])
	if _, ok := users[userName]; ok {
		return errors.New("Error - user already existing")
	}
	users[userName] = &User{
		username:  userName,
		createdAt: time.Now(),
	}
	fmt.Println("Success")
	return nil
}