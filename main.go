package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(\w+)\s(\d+)`)

	// (\w+) -> word followed by space '\s' followed by number (\d+)
	// Find 'John 1233"' matches
	matches := re.MatchString("John 1233")

	if matches {
		fmt.Println("Input matched")
	} else {
		fmt.Println("Input not matched")
	}
}
