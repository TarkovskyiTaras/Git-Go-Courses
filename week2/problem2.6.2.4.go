package main

import "fmt"

//Given a number n, create and return a new int array of length n, containing the numbers 0, 1, 2, ... n-1.

func fizzSlice(numbOfSliceEl int) {
	resultSlice := make([]int, numbOfSliceEl)
	elementValueSpacing := 0
	for i, _ := range resultSlice {
		resultSlice[i] = elementValueSpacing
		elementValueSpacing++

	}
	fmt.Println(resultSlice)
}

func main() {
	fizzSlice(1000)
}
