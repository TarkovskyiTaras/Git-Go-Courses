package main

import "fmt"

func main() {
	status := false
	array1 := [10]int{3, 1, 0, 0, 1, 3, 0, 0, 1, 3}

	for i := 0; i <= 1; i++ {
		if array1[i] == 1 && array1[i+1] == 3 {
			status = true
		}
	}
	for i := len(array1) - 4; i < len(array1)-1; i++ {
		if array1[i] == 1 && array1[i+1] == 3 {
			status = true
		}
	}
	fmt.Println(array1)
	fmt.Println(status)
}
