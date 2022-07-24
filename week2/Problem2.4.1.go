package main

import (
	"fmt"
	"math"
)

func main() {
	array1 := [10]int{1234, -5, 161, 232, 2422, 2, 554, 121232, 4, 876}
	maxValue := math.Inf(-1)
	minValue := math.Inf(+1)
	for _, x := range array1 {
		if float64(x) > maxValue {
			maxValue = float64(x)
		}
		if float64(x) < minValue {
			minValue = float64(x)
		}
	}
	fmt.Println(maxValue)
	fmt.Println(minValue)
}
