package main

import (
	"bufio"
	"fmt"
	"github.com/akshatvn/marketPlace/controller"
	"os"
)

func main() {
	fmt.Printf("Marketplace is starting...\nType HELP for whole list of commands. Type EXIT to exit.\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("# ")
		text, _ := reader.ReadString('\n')
		controller.ProcessInput(text)
	}
}

