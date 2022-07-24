package main

import "fmt"

func main() {

	array1 := [10]int{1, 22, 3, -4, 5, 6, 7, 448, 9, 10}
	array2 := [10]int{1, 2, 3, 4}

	for i, _ := range array1 {
		for j, _ := range array2 {
			if i == j {
				array2[j] = array1[i]
			}
		}
	}
	fmt.Println(array1, array2)
}
