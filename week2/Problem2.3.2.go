package main

import "fmt"

func commonEnd(x1 int, y1 int, x2 int, y2 int) bool {
	z := false
	array1 := [4]int{x1, 2, 5, y1}
	array2 := [4]int{x2, 41, 56, y2}

	if array1[0] == array2[0] || array1[3] == array2[3] {
		z = true
	}
	return z
}

func main() {
	fmt.Println(commonEnd(22, 42, 24, 42))

}
