package main

import (
	"fmt"
)

func main() {
	sliceHeight := 2
	sliceLength := 2
	sliceOfSlices1 := make([][]int, sliceHeight)
	sliceOfSlices2 := make([][]int, sliceHeight)
	for i, _ := range sliceOfSlices1 {
		sliceOfSlices1[i] = make([]int, sliceLength)
	}
	for i, _ := range sliceOfSlices2 {
		sliceOfSlices2[i] = make([]int, sliceLength)
	}

	for i, x := range sliceOfSlices1 {
		for j, _ := range x {
			x[j] = 1 + 2*i + j
		}
		fmt.Println(x)
	}
	fmt.Println("----------------------")
	for i, x := range sliceOfSlices2 {
		for j, _ := range x {
			x[j] = 4 + 2*i + j + 1
		}
		fmt.Println(x)
	}
	fmt.Println("----------------------")
	resultSlice := make([][]int, sliceHeight)
	for i, _ := range resultSlice {
		resultSlice[i] = make([]int, sliceLength)
	}
	for _, x := range resultSlice {
		for j, _ := range x {
			x[j] = 1
		}
	}
	var sumOfTwoMl int
	for k, _ := range resultSlice {

		for i, _ := range sliceOfSlices2 {
			sumOfTwoMl = 0

			for j, _ := range sliceOfSlices1 {
				sumOfTwoMl += sliceOfSlices1[k][j] * sliceOfSlices2[j][i]

				fmt.Printf("sumOfTwoMl = sliceOfSlices1[k][i]*sliceOfSlices2[j][i]  "+
					"%d*%d = %d\n", sliceOfSlices1[k][j], sliceOfSlices2[j][i], sumOfTwoMl)

				resultSlice[k][i] = sumOfTwoMl
				fmt.Printf("resultSlice[%d][%d] = sumOfTwoMl:  %d = %d\n", k, i, resultSlice[k][i], sumOfTwoMl)

				fmt.Println("-----------------------")

			}

		}
	}

	fmt.Println("||||||||||||||||||||||||||||||||||||||||||")

	for _, x := range resultSlice {
		fmt.Println(x)
	}
}
