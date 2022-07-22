package main

import "fmt"

//Given an array of positive and negative numbers. Find if there is a subarray (of size at-least one) with 0 sum.
//Example 1:
//
//Input:
//5
//4 2 -3 1 6
//
//Output:
//Yes
//
//Explanation:
//2, -3, 1 is the subarray
//with sum 0.

func main() {

	givenSlice := []int{1, 3, -1, -103, 150, 24, -5, -20, +1, 234, -1245}

	prefixSumOfEl := 0
	prefixSumMap := make(map[int]int)
	zeroSubstringStatus := false

	for _, x := range givenSlice {
		prefixSumOfEl += x
		prefixSumMap[prefixSumOfEl]++
		if prefixSumMap[prefixSumOfEl] == 2 {
			zeroSubstringStatus = true
			break
		}
	}

	if prefixSumOfEl == 0 {
		zeroSubstringStatus = true
	}

	fmt.Println(zeroSubstringStatus)
}