package main

import "fmt"

func main() {

	array1 := [10]int{1, 0, 3, 4, 45, 6, 7, 20, 9, 10}
	array2 := [10]int{1, 0, 3, 4, 5, 6, 7, 20, 9, 110}
	arraySum := [10]int{}
	for z, _ := range arraySum {
		for i, _ := range array1 {
			for j, _ := range array2 {
				if i == j && j == z {
					arraySum[z] = array1[i] + array2[j]
				}
			}
		}
	}
	fmt.Println(arraySum)
}
