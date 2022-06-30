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

		valueOfWind := inputIntReturn()

		if valueOfWind >= 1 && valueOfWind <= 4 {
			fmt.Println("Light breeze")
		}

		if valueOfWind >= 5 && valueOfWind <= 10 {
			fmt.Println("Mild wind")
		}
		if valueOfWind >= 11 && valueOfWind <= 18 {
			fmt.Println("Heavy Wind")
		}
		if valueOfWind >= 19 {
			fmt.Println("Tornado")
		}

	}
}
