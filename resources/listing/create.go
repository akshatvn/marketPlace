package listing

import (
	"errors"
	"fmt"
	"github.com/akshatvn/marketPlace/resources/category"
	"github.com/akshatvn/marketPlace/resources/user"
	"github.com/akshatvn/marketPlace/utils"
	"strconv"
	"time"
)

// Create creates a new listing
func Create(args ...string) error {
	if len(args) < 5 {
		return errors.New("Listing Create needs 5 arguments")
	}
	username, title, desciption, categoryName := args[0], args[1], args[2], args[4]
	price, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return errors.New("Error - price is invalid")
	}
	if !user.DoesExist(username) {
		return errors.New("Error - unknown user")
	}

	l := Listing{
		id:           utils.GenerateID(),
		username:     username,
		title:        title,
		description:  desciption,
		priceInCents: int(price * 100),
		categoryName: categoryName,
		createdAt:    time.Now(),
	}
	listingsByID[l.id] = &l
	// listingsByUser[username] =append(listingsByUser[username], &l)

	category.EnsureExists(categoryName)
	category.UpdateHeap(l.categoryName, true)

	fmt.Println(l.id)
	return nil
}

