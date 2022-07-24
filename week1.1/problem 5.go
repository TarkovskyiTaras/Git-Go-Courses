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
	for x := 1; x <= 5; x++ {
		fmt.Println("Enter two numbers: ")
		reader := bufio.NewReader(os.Stdin)
		input1, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input2, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNoSpace1 := strings.TrimSpace(input1)
		inputNoSpace2 := strings.TrimSpace(input2)
		inputFloat1, err := strconv.ParseFloat(inputNoSpace1, 64)
		if err != nil {
			log.Fatal(err)
		}
		inputFloat2, err := strconv.ParseFloat(inputNoSpace2, 64)
		if err != nil {
			log.Fatal(err)
		}
		if inputFloat1 > inputFloat2 {
			fmt.Println(inputFloat1 - inputFloat2)
		} else if inputFloat1 == inputFloat2 {
			fmt.Println("you've entered equal numbers")
		} else {
			fmt.Println(inputFloat1 + inputFloat2)
		}
	}
}
