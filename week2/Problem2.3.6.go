package main

import "fmt"

func main() {
	array1 := [3]int{2, 3, 242}
	fmt.Println(array1)
	for i := 0; i < 2; i++ {
		if array1[i] == 2 || array1[i+1] == 3 {
			array1[i+1] = 0
		}
	}
	fmt.Println(array1)
}
