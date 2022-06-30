package main

import "fmt"

func endsLy(givenString string) {
	if givenString[len(givenString)-2:len(givenString)] == "ly" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func main() {
	givenString := "oddly"
	endsLy(givenString)
}
