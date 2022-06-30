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

func returnNumb() (int, error) {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputNoSpace := strings.TrimSpace(input)
	inputInt, err := strconv.Atoi(inputNoSpace)
	if err != nil {
		err = errors.New("unacceptable value's been entered")
		fmt.Println(err.Error())
	}
	var wrongValueErr error
	if inputInt < 1 || inputInt > 12 {
		wrongValueErr = errors.New("unacceptable value's been entered")
	}

	return inputInt, wrongValueErr
}

func main() {
	var number int
	var valueError error
	for x := 1; true; x++ {
		number, valueError = returnNumb()

		if valueError != nil {
			fmt.Println(valueError.Error())
			continue
		}
		break
	}

	if number >= 3 && number <= 5 {
		fmt.Println("It's spring now.")
	} else if number >= 6 && number <= 8 {
		fmt.Println("It's summer now.")
	} else if number >= 9 && number <= 11 {
		fmt.Println("It's fall now.")
	} else {
		fmt.Println("It's winter now.")
	}

}
