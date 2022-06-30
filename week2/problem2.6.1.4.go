package main

import (
	"fmt"
)

//Given a string and an int n, return a string made of n repetitions of the last n characters of the string.
//You may assume that n is between 0 and the length of the string, inclusive.

func repeatEnd(givenString string, numbOfRep int) {

	endString := givenString[len(givenString)-numbOfRep:]
	var endStringNTimes string
	for x := 1; x <= numbOfRep; x++ {
		endStringNTimes += endString
	}
	fmt.Println(endStringNTimes)
}

func main() {
	givenString := "123456789T"
	numbOfRep := 5
	repeatEnd(givenString, numbOfRep)
}
