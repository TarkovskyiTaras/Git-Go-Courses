package main

import "fmt"

func main() {
	array1 := [10]int{1, 222, 34, 5554, 123135, 46, 7555, 898568, 9, 10}
	var sum int

	for _, x := range array1 {
		sum += x
	}
	var average float64
	average = float64(sum) / float64(len(array1))
	fmt.Println(average)
}
