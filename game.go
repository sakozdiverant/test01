package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	second := time.Now().Unix()
	rand.Seed(second)
	targer := rand.Intn(100) + 1
	fmt.Println("Я Выбираю число от 1 до 100")

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("У вас осталось ", 10-guesses, " попыток")
		fmt.Print("Напишите ваше число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)

		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess > targer {
			fmt.Println("Ваше значение больше чем загаданное: ")
		} else if guess < targer {
			fmt.Println("Ваше значение меньше чем загаданное: ")
		} else {
			success = true
			fmt.Print("Поздравляю вы угадали!!!")
			break
		}

	}
	if !success {
		fmt.Println("Извините ваше попытки закончились! Это число было: ", targer)
	}
}
