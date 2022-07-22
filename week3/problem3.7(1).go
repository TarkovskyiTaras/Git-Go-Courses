package main

import "fmt"

//3.7. Найти самое длинное слово в строке(разделенное одним пробелом)

func main() {

	givenString := "I go to  dfsad dsadsa dsadsaddasdsdasdadasdsa school every day."
	longestWordCharNumb := 0
	charCount := 0
	yPoint := 0

	for i, _ := range givenString {
		charCount += 1
		fmt.Println("-------------------")
		fmt.Println("i= ", i)
		fmt.Println("charCount++ :", charCount)

		if string(givenString[i]) == " " {
			fmt.Println("givenString[i], i: ", string(givenString[i]), i)

			if longestWordCharNumb < charCount {
				longestWordCharNumb = charCount - 1
				yPoint = i - 1
				charCount = 0
				fmt.Println("LongestWordCharNumb: ", longestWordCharNumb)
				fmt.Println("yPoint: ", yPoint)
				fmt.Println("charCount: ", charCount)
			}
			charCount = 0

		}

	}
	xPoint := yPoint - longestWordCharNumb + 1
	fmt.Println("")
	fmt.Println("----------------------------------------------------------")
	fmt.Printf("FINAL RESULT: Longest word in the string = %d chars, (xPoint: %d ; yPoint: %d)\n", longestWordCharNumb, xPoint, yPoint)
	fmt.Println("----------------------------------------------------------")
}
