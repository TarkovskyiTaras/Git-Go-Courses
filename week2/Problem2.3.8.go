package main

import "fmt"

func main() {
	array1 := [7]int{23, 4, 25, 64, 1, 242, 21}
	fmt.Println(array1)
	x := len(array1) / 2
	array2 := [3]int{array1[x-1], array1[x], array1[x+1]}
	fmt.Println(array2)

}
