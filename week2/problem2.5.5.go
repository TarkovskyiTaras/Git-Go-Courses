package main

import "fmt"

func theEnd(nonEmString string, boolFront bool) {
	var resultString string
	if boolFront == true {
		resultString = string(nonEmString[0])
	} else {
		resultString = string(nonEmString[len(nonEmString)-1])
	}
	fmt.Println(resultString)
}

func main() {
	nonEmString := "Taras"
	boolFront := false
	theEnd(nonEmString, boolFront)
}
