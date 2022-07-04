package main

import "fmt"

//Given an array of ints, return true if it contains no 1's or it contains no 4's.

func noOneOrFourInSlice(givenSlice []int) {
	unluckyNumb1 := 1
	unluckyNumb2 := 4
	luckMeasurer := true
	for _, c := range givenSlice {
		if c == unluckyNumb1 || c == unluckyNumb2 {
			luckMeasurer = false
			break
		}

	}
	fmt.Println(luckMeasurer)
}

func main() {
	givenSlice := []int{11, 2, 3, 14, 5, 6}
	noOneOrFourInSlice(givenSlice)

}
