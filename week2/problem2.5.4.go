package main

import "fmt"

func nonStart(firstString string, secondString string) {
	resultString := firstString[1:] + secondString[1:]
	fmt.Println(resultString)
}

func main() {

	firstString := "shotl"
	secondString := "java"
	nonStart(firstString, secondString)

}
