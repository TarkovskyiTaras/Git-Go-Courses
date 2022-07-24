package main

import "fmt"

// Given an array of ints, return true if the array contains no 1's and no 3's.

func noOneOrThreeInSlice(givenSlice []int) {

	unluckyElement1 := 1
	unluckyElement2 := 3
	var luckMeasurer bool

	for _, c := range givenSlice {

		if c != unluckyElement1 && c != unluckyElement2 {
			luckMeasurer = true
		}
		if c == unluckyElement1 || c == unluckyElement2 {
			luckMeasurer = false
			break
		}
	}
	fmt.Println(luckMeasurer)
}

func main() {
	givenSlice := []int{11, 2, 4, 13, 12}
	noOneOrThreeInSlice(givenSlice)
}
