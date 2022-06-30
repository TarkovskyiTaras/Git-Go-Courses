package main

import "fmt"

func no23(x int, y int) bool {
	boolZ := true
	array1 := [2]int{x, y}
	for _, z := range array1 {
		if z == 2 || z == 3 {
			boolZ = false
		}
	}

	return boolZ
}

func main() {
	fmt.Println(no23(2, 23))
}
