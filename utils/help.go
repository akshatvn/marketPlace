package utils

import "fmt"

// Print prints the static HELP output
func Print(args ...string) error {
	fmt.Printf(
		`HELP
REGISTER <username>
CREATE_LISTING <username> <title> <description> <price> <category>
DELETE_LISTING <username> <listing_id>
GET_LISTING <username> <listing_id>
GET_CATEGORY <username> <category> {sort_price|sort_time} {asc|dsc}
GET_TOP_CATEGORY <username>
EXIT
`)
	return nil
}
