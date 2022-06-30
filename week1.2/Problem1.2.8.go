package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func returnNumb() float64 {
	var inputInt float64
	var wrongValueErr error
	reader := bufio.NewReader(os.Stdin)
	for x := 1; true; x++ {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNoSpace := strings.TrimSpace(input)
		inputInt, wrongValueErr = strconv.ParseFloat(inputNoSpace, 64)
		if wrongValueErr != nil || inputInt <= 0 {
			wrongValueErr = errors.New("unacceptable value's been entered")
			fmt.Println(wrongValueErr)
			continue
		}
		break
	}
	fmt.Println(inputInt, "before return")
	return inputInt
}

func main() {
	var price float64
	var discount float64
	for x := 1; true; x++ {
		price = returnNumb()
		if price > 1000 {
			discount = 0.10
			fmt.Printf("\nYour 10%% discount is %.2f$. Your total is going to be %.2f$.\n\n", price*discount, price-price*discount)
		}
	}

}
