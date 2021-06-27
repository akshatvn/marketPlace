package user

// DoesExist tells whether a user with a given username exists or not
func DoesExist(username string) bool {
	_, ok := users[username]
	return ok
}
