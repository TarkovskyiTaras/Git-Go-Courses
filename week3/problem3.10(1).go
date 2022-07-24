package main

import (
	"fmt"
	"strings"
)

//3.10*. Удалить из предложения слова которые повторяются.

func main() {
	//approach through the usage of maps

	givenString := "bla bla my name is Taras Nikita Nikita go go very very"
	splitString := strings.Split(givenString, " ")
	splitStringMap := make(map[string]int)

	for _, x := range splitString {
		splitStringMap[x]++
	}
	var newString string
	for key1, count1 := range splitStringMap {
		if count1 == 1 {
			newString += key1 + " "
		}
	}

	fmt.Println(newString)
}
