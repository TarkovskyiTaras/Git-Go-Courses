package main

import "fmt"

//Return the number of times that the string "code" appears anywhere in the given string, except we'll accept any letter for the 'd', so "cope" and "cooe" count.

func main() {
	givenString := "codeefeebcosebbsefsdcoee"
	searchedString := "code"
	var searchedStringDiffSecChar string
	matchCount := 0
	for i, _ := range givenString {
		if i > len(givenString)-4 {
			break
		}
		x := givenString[i : i+4]
		fmt.Println("given string: ", x)
		fmt.Println("searched string before if: ", searchedString)
		if string(x[2]) != "d" {
			searchedStringDiffSecChar = fmt.Sprintf("co%se", string(x[2]))
		}
		fmt.Println("searched string diff second char: ", searchedStringDiffSecChar)
		if x == searchedString || x == searchedStringDiffSecChar {
			matchCount++
		}

	}
	fmt.Printf("Searched string has apeeared %d times in the given string.", matchCount)
}
