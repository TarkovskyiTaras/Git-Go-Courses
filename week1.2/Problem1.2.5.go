package main

import "fmt"

func main() {

	var x float64
	var y float64
	var z float64

	x = 23
	y = 42
	z = 23
	fmt.Println(x, y, z)
	if x == y || x == z || y == z {
		fmt.Println("\nThere is a match")
	}

}
