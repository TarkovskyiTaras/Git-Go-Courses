package main

import "fmt"

//2.1 Инвертировать массив (без использования допольнительного массива)

func main() {

	givenSlice := []int{12, 242, 31, 4, 595, 67, 702, 18, 29, 410}
	var lastIndex int

	for _, x := range givenSlice {
		givenSlice = append(givenSlice, x)
		lastIndex = len(givenSlice) - 1
	}
	for i, _ := range givenSlice {
		if i == lastIndex-1/2 {
			break
		}
		givenSlice[i] = givenSlice[lastIndex-i]
	}
	givenSlice = givenSlice[0 : len(givenSlice)/2]

	fmt.Println(givenSlice)
}
