package main

import "fmt"

//3.2. Создать логическую матрицу по следующим критериям, если і равно j, то элементом будет true.

func main() {
	sliceLength := 9 //horizontal dimension
	sliceHeight := 9 //vertical dimension

	resultSliceOfSlices := make([][]bool, sliceHeight)
	for i, _ := range resultSliceOfSlices {
		resultSliceOfSlices[i] = make([]bool, sliceLength)
	}
	for i, y := range resultSliceOfSlices {
		for j, _ := range y {
			if j == i {
				y[j] = true
			} else {
				y[j] = false
			}
		}
		fmt.Println(y)
	}

}
