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
	//fmt.Println("cls")
	fmt.Print("Сколько вам лет: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	old, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Вводите только цыфры")
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	result := ""
	if old >= 18 {
		result = "Достаточный"
	} else {
		result = "Не достаточный"
	}

	fmt.Print("Вш возраст: " + result)
}
