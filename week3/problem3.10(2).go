package main

import "fmt"

//3.10*. Удалить из предложения слова которые повторяются.

func main() {
	//approach avoiding the usage of strings.Spit and maps

	givenString := "I bla nikitapoc go nikitapoc bla to taras school every taras bla day "
	var stringSlice []string
	var arrayElement string
	for _, x := range givenString {

		if string(x) != " " {
			arrayElement += string(x)
		}
		if string(x) == " " {
			stringSlice = append(stringSlice, arrayElement)
			arrayElement = ""
		}

	}

	var stringSliceChecker []string

	for i, x := range stringSlice {
		stringSliceChecker = append(stringSliceChecker, "empty")
		for j, _ := range stringSliceChecker {
			if stringSlice[i] == stringSliceChecker[j] {
				stringSlice[i] = ""
				stringSlice[j] = ""
			}
		}

		stringSliceChecker[i] = x

	}
	var fixedString string
	for i, x := range stringSlice {
		if i == len(stringSlice)-1 {
			fixedString += x + "."
			break
		}
		if x != "" {
			fixedString += x + " "
		}

	}
	fmt.Println(fixedString)
}
