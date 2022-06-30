package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func return2() (float64, float64, float64) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter 3 sides of a triangle:")
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

		if math.Pow(number1, 2)+math.Pow(number2, 2) == math.Pow(number3, 2) ||
			math.Pow(number1, 2)+math.Pow(number3, 2) == math.Pow(number2, 2) ||
			math.Pow(number2, 2)+math.Pow(number3, 2) == math.Pow(number1, 2) {
			fmt.Println("You have a right triangle!")
		} else {
			fmt.Println("Not a right triangle")
		}
		fmt.Println(int(number1) % 2)
	}
}
