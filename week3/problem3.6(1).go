package main

import "fmt"

//3.6 Требуется найти самую длинную непрерывную цепочку нулей  и единиц в последовательности цифр.
func main() {
	givenString := "1234004500"
	maxZeroNumb := 0
	maxZeroSliceX := 0
	maxZeroSliceY := 0

	for j := 0; j < len(givenString); j++ {

		for i, _ := range givenString {
			if i+j < len(givenString) {

				zeroCharCount := 0
				slicedString := givenString[j : i+j+1]
				for _, z := range slicedString {

					if string(z) != "0" {
						break
					}

					if string(z) == "0" {
						zeroCharCount += 1
					}

				}

				if zeroCharCount > maxZeroNumb {
					maxZeroNumb = zeroCharCount
					maxZeroSliceX = j + 1
					maxZeroSliceY = i + j + 1
				}

			}

		}
	}
	fmt.Println("-----------------")
	fmt.Println("maxzeronumb: ", maxZeroNumb)
	fmt.Println("x: ", maxZeroSliceX)
	fmt.Println("y: ", maxZeroSliceY)
}
