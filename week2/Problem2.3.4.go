package main

import "fmt"

func middleWay(array1 [3]int, array2 [3]int) [2]int {
	array3 := [2]int{array1[1], array2[1]}
	return array3
}

func main() {
	a := [3]int{22, 42, 53}
	b := [3]int{422, 532, 1024}
	fmt.Println(middleWay(a, b))
}
