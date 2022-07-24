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

func lengthDay() (float64, float64) {
	var inputInt1 float64
	var inputInt2 float64
	var wrongValueErr error
	reader := bufio.NewReader(os.Stdin)
	for x := 1; true; x++ {
		fmt.Println("How long is the appointment going to be:")
		input1, err := reader.ReadString('\n')
		fmt.Println("Please enter the day of the week (1-7):")
		input2, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNoSpace1 := strings.TrimSpace(input1)
		inputNoSpace2 := strings.TrimSpace(input2)
		inputInt1, wrongValueErr = strconv.ParseFloat(inputNoSpace1, 64)
		inputInt2, wrongValueErr = strconv.ParseFloat(inputNoSpace2, 64)
		if wrongValueErr != nil || inputInt1 <= 0 || inputInt2 <= 0 || inputInt2 > 7 {
			wrongValueErr = errors.New("unacceptable value's been entered")
			fmt.Println(wrongValueErr.Error())
			continue
		}
		break
	}

	return inputInt1, inputInt2

}

func main() {
	var pricePerH float64 = 30.0
	var discount float64
	var length, day float64
	for x := 1; true; x++ {
		length, day = lengthDay()
		if day == 6 || day == 7 {
			discount = 0.2
			fmt.Printf("\nYour 20%% discount is %.2f$. Your total is going to be %.2f$.\n\n", pricePerH*length*discount, pricePerH*length-pricePerH*length*discount)
		} else {
			fmt.Printf("\nYour total is going to be %.2f$.\n\n", pricePerH*length)
		}
	}

}
