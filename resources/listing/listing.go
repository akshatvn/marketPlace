package listing

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Listing defines the structure of a listing record
type Listing struct {
	id           string
	username     string
	title        string
	description  string
	priceInCents int
	categoryName string
	createdAt    time.Time
}

var listingsByID map[string]*Listing

// var listingsByUser map[string][]*Listing

func init() {
	listingsByID = make(map[string]*Listing)
	// listingsByUser = make(map[string][]*Listing)
}

// print prints the listing in the following format
// <title>|<description>|<price>|<created_at>|<category>|<username>
func (l *Listing) print() {
	fmt.Println(
		strings.Join([]string{
			l.title,
			l.description,
			strconv.FormatFloat(float64(l.priceInCents)/100.0, 'f', 2, 64),
			l.createdAt.Format(time.RFC3339),
			l.categoryName,
			l.username,
		}, "|"),
	)
}
