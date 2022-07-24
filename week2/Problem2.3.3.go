package main

import "fmt"

func reverse3(x int, y int, z int) [3]int {
	array1 := [3]int{x, y, z}
	fmt.Println(array1)
	arrayReverse := [3]int{z, y, x}

	return arrayReverse
}

func main() {
	fmt.Println(reverse3(21, 256, 424))
}
