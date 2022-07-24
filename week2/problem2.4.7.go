package main

import "fmt"

func main() {

	array1 := [10]int{124, 222222222225, 32, 4, 5, 6, 7, 8, 999999999, 10}
	var sum1 int
	var sum2 int

	for i, _ := range array1 {

		if i < len(array1)/2 {
			sum1 += array1[i]
		}
		if i >= len(array1)/2 {
			sum2 += array1[i]
		}

	}

	if sum1 > sum2 {
		for i := 0; i < len(array1)/2; i++ {
			fmt.Print(array1[i], " ")
		}
	}
	if sum2 > sum1 {
		for i := len(array1) / 2; i < len(array1); i++ {
			fmt.Print(array1[i], " ")
		}

	}
}
