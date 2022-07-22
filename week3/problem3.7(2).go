package main

import (
	"fmt"
	"strings"
)

//3.7. Найти самое длинное слово в строке(разделенное одним пробелом)

func main() {

	givenString := "I go to school every day"
	fmt.Println(givenString)
	y := strings.Split(givenString, " ")
	fmt.Println("y: ", y)

	for _, x := range y {
		fmt.Println(x)
	}
	maxLengthString := 0
	xPoint := 0

	fmt.Println("----------------------------")

	for i, x := range y {
		if len(x) > maxLengthString {
			maxLengthString = len(x)
			xPoint = i
		}

	}
	fmt.Println(y[xPoint])
}
