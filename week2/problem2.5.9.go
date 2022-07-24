package main

import (
	"fmt"
	"strings"
)

//Given a string, if the string begins with "red" or "blue" return that color string, otherwise return the empty string.

func seeColor(givenString string) {
	if strings.Contains(givenString[0:3], "red") {
		fmt.Println("red")
	} else if strings.Contains(givenString[0:4], "blue") {
		fmt.Println("blue")
	} else {
		fmt.Println("")
	}
}

func main() {
	givenString := "dblueaefsfds f sdf"
	seeColor(givenString)
}
