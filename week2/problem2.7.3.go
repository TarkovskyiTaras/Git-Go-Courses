package main

import "fmt"

//2.3 Посчитать колличество букв 'a' в строке (искомый символ может быть разным)

func main() {
	givenString := "123a123a123a123a123"
	searchedChar := "a"
	matchCount := 0
	for _, c := range givenString {
		if string(c) == searchedChar {
			matchCount += 1
		}
	}
	fmt.Printf("There are %d searched chars in the string", matchCount)
}
