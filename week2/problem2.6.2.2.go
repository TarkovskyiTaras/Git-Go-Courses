package main

import "fmt"

// Return the sum of the numbers in the array, returning 0 for an empty array.
//Except the number 13 is very unlucky, so it does not count and numbers that come immediately after a 13 also do not count.

func sumOfSliceElements(givenSlice []int) {
	sumOfEl := 0
	unluckyNumb := 13
	for i, c := range givenSlice {
		lastIndex := len(givenSlice) - 1
		sumOfEl += c

		if c == unluckyNumb && i == lastIndex {
			sumOfEl -= givenSlice[lastIndex]

		}

		if c == unluckyNumb && i != lastIndex {
			sumOfEl -= c

			sumOfEl -= givenSlice[i+1]

		}

	}
	fmt.Println(sumOfEl)
}

func main() {
	givenSlice := []int{100, 100, 13, 100, 100, 100, 100, 100, 100, 13}
	sumOfSliceElements(givenSlice)
}
