package main

import "fmt"

func middleThree(givenString string) {
	if len(givenString) > 3 {

		resultString := givenString[len(givenString)/2-1 : len(givenString)/2+2]
		fmt.Println(resultString)
	}
}

func main() {
	givenString := "Taras"
	middleThree(givenString)
}
