package main

import "fmt"

//Given a string, return a string where for every char in the original, there are two chars.

func main() {

	givenString := "tarasnikita"
	var doubleCharString string
	for i, _ := range givenString {
		doubleCharString += string(givenString[i]) + string(givenString[i])

	}
	fmt.Println(doubleCharString)
}
