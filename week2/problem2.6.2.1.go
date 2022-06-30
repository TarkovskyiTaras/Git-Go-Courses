package main

import "fmt"

//Return the number of even ints in the given array. Note: the % "mod" operator computes the remainder, e.g. 5 % 2 is 1.

func main() {
	givenArray := [10]int{1, 43, 3, 4, 5, 6, 7, 8, 9, 101}
	countEven := 0
	for i, _ := range givenArray {
		if givenArray[i]%2 == 0 {
			countEven++
		}
	}
	fmt.Println(countEven)
}
