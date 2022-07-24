package main

import "fmt"

func main() {

	array1 := [8]int{2, 3, 6, 77, 89, 24, 35, 23}
	x := len(array1) / 2
	array2 := [2]int{array1[x-1], array1[x]}
	fmt.Println(array1)
	fmt.Println(array2)
}
