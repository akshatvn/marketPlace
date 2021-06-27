package listing

import (
	"errors"
	"github.com/akshatvn/marketPlace/resources/user"
)

// GetByUserAndId gets a listing given a username and an id
func GetByUserAndId(args ...string) error {
	if len(args) < 2 {
		return errors.New("Listing GetByUserAndId needs 2 arguments")
	}
	username, listingID := args[0], args[1]
	l, ok := listingsByID[listingID]
	if !ok {
		return errors.New("Error - not found")
	}
	if !user.DoesExist(username) {
		return errors.New("Error - unknown user")
	}
	l.print()
	return nil
}
