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
		fmt.Println("Enter 3 numbers:")
		numb1, numb2, numb3 := return2()

		array1 := [3]float64{numb1, numb2, numb3}

		for _, x := range array1 {
			if x >= 0 {
				fmt.Println(math.Pow(x, 3))
			} else {
				x = 0
				fmt.Println(x)
			}
		}
	}
}
