package main

import (
	"fmt"
)

func main() {

	notes := [7]string{
		"do",
		"re",
		"mi",
	}

	for _, value := range notes {
		fmt.Println(value)
	}
}
