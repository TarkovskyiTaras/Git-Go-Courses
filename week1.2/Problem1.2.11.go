package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputIntReturn() int {
	var inputInt int
	reader := bufio.NewReader(os.Stdin)
	for x := 1; true; x++ {
		fmt.Println("Enter a number from 1 to 99:")
		input, wrongValueErr := reader.ReadString('\n')
		if wrongValueErr != nil {
			wrongValueErr = errors.New("unacceptable value's been entered")
			fmt.Println(wrongValueErr.Error())
		}
		inputNoSpace := strings.TrimSpace(input)
		inputInt, wrongValueErr = strconv.Atoi(inputNoSpace)
		if wrongValueErr != nil || inputInt < 0 || inputInt > 99 {
			wrongValueErr = errors.New("unacceptable value's been entered")
			fmt.Println(wrongValueErr.Error())
			continue
		}
		break
	}
	return inputInt
}

func main() {
	for x := 1; true; x++ {

		numbOfCoins := inputIntReturn()

		if numbOfCoins < 11 || numbOfCoins > 14 {
			if numbOfCoins%10 == 1 && numbOfCoins != 11 {
				fmt.Printf("%d копейка", numbOfCoins)
			}

			if numbOfCoins%10 >= 2 && numbOfCoins%10 <= 4 {
				fmt.Printf("%d копейки", numbOfCoins)
			}
			if numbOfCoins%10 >= 5 && numbOfCoins%10 <= 9 || numbOfCoins%10 == 0 {
				fmt.Printf("%d копеек", numbOfCoins)
			}

		} else {
			fmt.Printf("%d копеек", numbOfCoins)
		}
		fmt.Println()
	}
}
