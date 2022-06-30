package main

import "fmt"

//Return true if the given string contains a "bob" string, but where the middle 'o' char can be any char.

func main() {
	givenString := "codeefeebbsebcbsefsdcoee"
	searchedString := "bob"
	matchStatus := false
	for i, _ := range givenString {
		if i > len(givenString)-3 {
			break
		}
		x := givenString[i : i+3]
		if x[0] == searchedString[0] && x[2] == searchedString[2] {
			matchStatus = true
			break
		}

	}
	fmt.Println(matchStatus)
}
