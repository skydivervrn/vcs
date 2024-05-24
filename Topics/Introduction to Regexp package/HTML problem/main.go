package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// DO NOT modify the code block below:
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	node := scanner.Text()

	// Create a regexp to match the attribute below:
	reBackground := regexp.MustCompile(`.*style="background-color: red;".*`)

	if reBackground.MatchString(node) {
		fmt.Println("reject")
	} else {
		fmt.Println("accept")
	}
}
