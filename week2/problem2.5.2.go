package main

import (
	"fmt"
)

func makeOutWord(outString string, wordString string) {
	resultString := outString[:2] + wordString + outString[2:]
	fmt.Println(resultString)

}

func main() {
	outString := "||<<"
	wordString := "NikitaPoc"
	makeOutWord(outString, wordString)

}
