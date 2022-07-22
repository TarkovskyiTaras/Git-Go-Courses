package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//4.3. Пользователь вводит число. Вывести на экран его удвоенное значение, если число делится на 7 нацело.

func main() {
	reader := bufio.NewReader(os.Stdin)
	number, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	numberStripped := strings.TrimSpace(number)
	numberParsed, err := strconv.Atoi(numberStripped)
	if err != nil {
		log.Fatal(err)
	}
	if numberParsed%7 == 0 {
		numberDoubled := numberParsed * 2
		fmt.Println(numberDoubled)
	} else {
		fmt.Println("The number is not divisible by 7.")
	}
}
