package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//4.2. Пользователь вводит три числа с консоли, нужно вывести на консоль наибольшее, наименьшее

func main() {
	for x := 5; x >= 0; x-- {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Yor're to enter 3 numbers")
		fmt.Print("1) ")
		user_input1, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("2) ")
		user_input2, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("3) ")
		user_input3, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input_stripped_of_space1 := strings.TrimSpace(user_input1)
		input_stripped_of_space2 := strings.TrimSpace(user_input2)
		input_stripped_of_space3 := strings.TrimSpace(user_input3)

		input_integer1, err := strconv.Atoi(input_stripped_of_space1)
		if err != nil {
			log.Fatal(err)
		}
		input_integer2, err := strconv.Atoi(input_stripped_of_space2)
		if err != nil {
			log.Fatal(err)
		}
		input_integer3, err := strconv.Atoi(input_stripped_of_space3)
		if err != nil {
			log.Fatal(err)
		}
		if input_integer1 > input_integer2 && input_integer1 > input_integer3 {
			fmt.Println(input_integer1)
		} else if input_integer2 > input_integer1 && input_integer2 > input_integer3 {
			fmt.Println(input_integer2)
		} else if input_integer3 > input_integer1 && input_integer3 > input_integer2 {
			fmt.Println(input_integer3)
		}
		if input_integer1 < input_integer2 && input_integer1 < input_integer3 {
			fmt.Println(input_integer1)
		} else if input_integer2 < input_integer1 && input_integer2 < input_integer3 {
			fmt.Println(input_integer2)
		} else if input_integer3 < input_integer1 && input_integer3 < input_integer2 {
			fmt.Println(input_integer3)
		}
	}
}
