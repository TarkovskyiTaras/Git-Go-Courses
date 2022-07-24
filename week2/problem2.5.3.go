package main

import "fmt"

func firstHalf(stringOfEvenLength string) {
	firstHalfString := stringOfEvenLength[:len(stringOfEvenLength)/2]
	fmt.Println(firstHalfString)
}

func main() {
	stringOfEvenLength := "HelloThere"
	firstHalf(stringOfEvenLength)
}
