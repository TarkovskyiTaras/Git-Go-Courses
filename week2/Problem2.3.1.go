package main

import "fmt"

func return1(x int, y int) bool {
	z := false
	array1 := [6]int{x, 6, 5, 14, 12, y}
	if array1[0] == 6 || array1[5] == 6 {
		z = true
	}
	return z
}

func main() {
	fmt.Println(return1(6, 12))
}
