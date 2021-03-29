package main

import (
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.test2("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0

	for __, number := range numbers {
		sum += number
	}

	count := float64(len(numbers))
	fmt.Println(sum / count)
}
