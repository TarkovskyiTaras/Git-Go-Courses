package main

import "fmt"

func main() {

	givenString := "WASITARATISAW"
	fmt.Println(givenString)

	var backwardString string

	for i, _ := range givenString {
		backwardString += string(givenString[len(givenString)-1-i])
	}
	fmt.Println(backwardString)
	if givenString == backwardString {
		fmt.Println("given string is a palindrome")
	} else {
		fmt.Println("not a palindrome")
	}
}
