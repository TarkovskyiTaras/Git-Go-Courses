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
	fmt.Print("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputNoSpace := strings.TrimSpace(input)
	inputFloat, err := strconv.ParseFloat(inputNoSpace, 64)
	if err != nil {
		log.Fatal(err)
	}
	if inputFloat >= 0 && inputFloat <= 1 {
		fmt.Println("Within 0 to 1")
	} else {
		fmt.Println("Falls outside of 0 to 1")
	}
}
