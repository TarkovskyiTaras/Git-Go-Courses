package main

import (
	"fmt"
	"strings"
)

//Given a string, consider the prefix string made of the first N chars of the string.
//Does that prefix string appear somewhere else in the string?
//Assume that the string is not empty and that N is in the range 1..str.length().

func prefixAgain(givenString string, numbOfChars int) {
	prefixString := givenString[:numbOfChars]
	result := strings.Contains(givenString[numbOfChars:], prefixString)
	fmt.Println(result)
}

func main() {

	givenString := "abXYabc"
	numbOfChars := 3
	prefixAgain(givenString, numbOfChars)

}
