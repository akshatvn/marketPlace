package listing

import (
	"errors"
	"fmt"
	"github.com/akshatvn/marketPlace/resources/category"
	"github.com/akshatvn/marketPlace/resources/user"
	"sort"
	"strings"
)

// GetAllByCategory gets all listings of a particular category.
// there is an provide sort parameter and direction.
func GetAllByCategory(args ...string) error {
	if len(args) < 2 {
		return errors.New("Category Get needs at least 2 arguments")
	}
	username, categoryName, sortBy, sortDir := strings.ToLower(args[0]), args[1], "", ""

	_, ok := category.GetCategoryIndex()[categoryName]
	if !ok {
		return errors.New("Error - category not found ")
	}
	if !user.DoesExist(username) {
		return errors.New("Error - unknown user")
	}

	listings := make([]*Listing, 0)
	for _, l := range listingsByID {
		if l.categoryName == categoryName {
			listings = append(listings, l)
		}
	}
	if len(args) >= 4 {
		sortBy, sortDir = strings.ToLower(args[2]), strings.ToLower(args[3])
		if sortDir != "asc" && sortDir != "desc" {
			fmt.Println("unknown sort direction, ignoring sort")
		} else if sortBy == "sort_price" {
			if sortDir == "asc" {
				sort.Slice(listings[:], func(i, j int) bool {
					return listings[i].priceInCents < listings[j].priceInCents
				})
			} else {
				sort.Slice(listings[:], func(i, j int) bool {
					return listings[i].priceInCents > listings[j].priceInCents
				})
			}

		} else if sortBy == "sort_time" {
			if sortDir == "asc" {
				sort.Slice(listings[:], func(i, j int) bool {
					return listings[i].createdAt.UnixNano() < listings[j].createdAt.UnixNano()
				})
			} else {
				sort.Slice(listings[:], func(i, j int) bool {
					return listings[i].createdAt.UnixNano() > listings[j].createdAt.UnixNano()
				})
			}

		}
	}
	for _, l := range listings {
		l.print()
	}
	return nil
}
