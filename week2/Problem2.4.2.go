package main

import (
	"fmt"
	"math"
)

func main() {
	array1 := [10]int{-1234, -5, 161, 232, 1000000, 2, 554, 121232, 4, 876}
	maxValue := math.Inf(-1)
	minValue := math.Inf(+1)
	iMaxValue := 0
	iMinValue := 0
	for i, _ := range array1 {
		if float64(array1[i]) > maxValue {
			maxValue = float64(array1[i])
			iMaxValue = i
		}
		if float64(array1[i]) < minValue {
			minValue = float64(array1[i])
			iMinValue = i
		}

	}
	array1[iMaxValue] = int(minValue)
	array1[iMinValue] = int(maxValue)
	fmt.Println(array1)
}
