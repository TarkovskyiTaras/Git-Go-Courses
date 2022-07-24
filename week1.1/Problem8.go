package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func return2() (float64, float64) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter 2 numbers:")
	input1, err := reader.ReadString('\n')
	input2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputStripped1 := strings.TrimSpace(input1)
	inputStripped2 := strings.TrimSpace(input2)
	inputFloat1, err := strconv.ParseFloat(inputStripped1, 64)
	if err != nil {
		log.Fatal(err)
	}
	inputFloat2, err := strconv.ParseFloat(inputStripped2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return inputFloat1, inputFloat2
}

func main() {
	for x := 1; true; x++ {
		number1, number2 := return2()
		if int(number1)%10 == int(number2)%10 {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
