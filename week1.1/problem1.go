package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//4.1. Пользователь вводит параметр с консоли (текущий час от 0 до 24)
//				Если время от 9 до 18, то выводим "Я на работе", в другом случае "Я отдыхаю")

func main() {
	for x := 5; x >= 0; x-- {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a time value 0-24: ")
		user_input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input_stripped_of_space := strings.TrimSpace(user_input)
		input_integer, err := strconv.Atoi(input_stripped_of_space)

		if input_integer >= 9 && input_integer <= 18 {
			fmt.Println("I'm out at work")
		} else if input_integer >= 0 && input_integer < 9 {
			fmt.Println("I'm chilling at home")
		} else if input_integer > 18 && input_integer <= 24 {
			fmt.Println("I'm chilling at home")
		} else {
			fmt.Println("Oops, you've entered a wrong time value.")
		}
	}
}
