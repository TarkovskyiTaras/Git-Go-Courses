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

func returnWeight() (float64, float64) {
	var inputFloat1 float64
	var inputFloat2 float64
	var wrongValueErr error
	reader := bufio.NewReader(os.Stdin)
	for x := 1; true; x++ {
		fmt.Println("Please enter your height:")
		input1, err := reader.ReadString('\n')
		fmt.Println("Please enter your weight:")
		input2, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNoSpace1 := strings.TrimSpace(input1)
		inputNoSpace2 := strings.TrimSpace(input2)
		inputFloat1, wrongValueErr = strconv.ParseFloat(inputNoSpace1, 64)
		inputFloat2, wrongValueErr = strconv.ParseFloat(inputNoSpace2, 64)
		if wrongValueErr != nil || inputFloat1 <= 0 || inputFloat2 <= 0 {
			wrongValueErr = errors.New("unacceptable value's been entered")
			fmt.Println(wrongValueErr.Error())
			continue
		}

		break
	}

	return inputFloat1, inputFloat2

}

func main() {
	for x := 1; true; x++ {
		usersHeight, usersWeight := returnWeight()
		properWeight := usersHeight - 100

		if usersWeight > properWeight {
			fmt.Println("You might wanna consider losing some weight.")
		}
		if usersWeight < properWeight {
			fmt.Println("You might wanna consider gaining some weight.")
		}
		if usersWeight == properWeight {
			fmt.Println("Keep it up! You are good!")
		}
	}
}
