package main

import "fmt"

//Given an array containing N integers and an integer K., Your task is to find the length of the longest Sub-Array with the sum of the elements equal to the given value K.
//Example 1:
//
//
//Input :
//A[] = {10, 5, 2, 7, 1, 9}
//K = 15
//Output : 4
//Explanation:
//The sub-array is {5, 2, 7, 1}.

func main() {

	n := 8
	k := 133

	givenSlice := make([]int, n)
	givenSlice = []int{132, 1, 3, 4, 5, 6, 7, 8}
	prefixSumMap := make(map[int]int)
	prefixSum := 0

	for _, x := range givenSlice {
		prefixSum += x
		prefixSumMap[prefixSum] = 1
		if prefixSum == k {
			fmt.Println("true")
		}
		if prefixSumMap[prefixSum-k] == prefixSumMap[prefixSum] {
			fmt.Println("true")
		}

	}

	fmt.Println(prefixSumMap)

}
