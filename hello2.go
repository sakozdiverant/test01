package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Сколько вам лет: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	old, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	if old >= 18 {
		fmt.Println("Ого какой взрослый ! ")
	} else {
		fmt.Println("Вам здесь не место")
	}

	//fmt.Println(input)
}
