package listing

import (
	"errors"
	"fmt"
	"github.com/akshatvn/marketPlace/resources/category"
	"github.com/akshatvn/marketPlace/resources/user"
	"strings"
)

// Delete deletes a given listing
func Delete(args ...string) error {
	if len(args) < 2 {
		return errors.New("Listing Delete needs 2 arguments")
	}
	username, listingID := args[0], args[1]
	l, ok := listingsByID[listingID]
	if !ok {
		return errors.New("Error - listing does not exist")
	}
	if !user.DoesExist(username) || strings.ToLower(l.username) != username {
		return errors.New("Error - listing owner mismatch")
	}
	delete(listingsByID, l.id)
	category.UpdateHeap(l.categoryName, false)
	fmt.Println("Success")
	return nil
}