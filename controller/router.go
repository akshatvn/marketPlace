package controller

import (
	"errors"
	"fmt"
	"github.com/akshatvn/marketPlace/resources/category"
	"github.com/akshatvn/marketPlace/resources/listing"
	"github.com/akshatvn/marketPlace/resources/user"
	"github.com/akshatvn/marketPlace/utils"
	"os"
	"regexp"
	"strings"
)

// ProcessInput takes a line of input and processes it
func ProcessInput(text string) {
	inputStrs := splitIntoArgs(text)
	action := strings.ToUpper(inputStrs[0])
	args := inputStrs[1:]
	processor, err := getFunc(action)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err = processor(args...); err != nil {
		fmt.Println(err.Error())
	}
}

// splitIntoArgs takes a line of input and returns a slice of
// strings. The first element of the slice is the command, and
// the following elements are its arguments
func splitIntoArgs(text string) []string {
	text = strings.TrimSpace(text)
	// FOR SIMPLICITY OF REGEX, QUOTES INSIDE AN ARGUMENT ARE NOT HANDLED.
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)
	ans :=  r.FindAllString(text, -1 )
	r = regexp.MustCompile(`^['"]{1}|['"]{1}$`)
	for i, arg := range ans {
		ans[i] = r.ReplaceAllString(arg, "")
	}
	return ans
}


// getFunc takes the input command and returns the corresponding
// handler function
func getFunc(action string) (func(args ...string) error, error) {
	funcMap := map[string]func(args ...string) error{
		"HELP":             utils.Print,
		"REGISTER":         user.Create,
		"CREATE_LISTING":   listing.Create,
		"DELETE_LISTING":   listing.Delete,
		"GET_LISTING":      listing.GetByUserAndId,
		"GET_CATEGORY":     listing.GetAllByCategory,
		"GET_TOP_CATEGORY": category.GetTopCategory,
		"EXIT":             exit,
	}
	if f, ok := funcMap[action]; ok {
		return f, nil
	}
	return nil, errors.New(fmt.Sprintf("Unknown command %s\n", action))
}

// exit exits the program with code 0
func exit(args ...string) error {
	os.Exit(0)
	return nil
}

