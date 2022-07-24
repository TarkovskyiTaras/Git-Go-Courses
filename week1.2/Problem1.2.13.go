package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func return2() (float64, float64, float64) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter 3 numbers:")
	input1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input3, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputStripped1 := strings.TrimSpace(input1)
	inputStripped2 := strings.TrimSpace(input2)
	inputStripped3 := strings.TrimSpace(input3)
	inputFloat1, err := strconv.ParseFloat(inputStripped1, 64)
	if err != nil {
		log.Fatal(err)
	}
	inputFloat2, err := strconv.ParseFloat(inputStripped2, 64)
	if err != nil {
		log.Fatal(err)
	}
	inputFloat3, err := strconv.ParseFloat(inputStripped3, 64)
	if err != nil {
		log.Fatal(err)
	}
	return inputFloat1, inputFloat2, inputFloat3
}
func main() {
	for x := 1; true; x++ {
		number1, number2, number3 := return2()

		if number1+number2+number3 > 0 {
			number1 = number1 * 2
			number2 = number2 * 2
			number3 = number3 * 3
			fmt.Println(number1, number2, number3)
		} else {
			number1 = 0
			number2 = 0
			number3 = 0
			fmt.Println(number1, number2, number3)
		}
	}
}
