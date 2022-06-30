package main

import "fmt"

func splitArray(array [10]int, start int, end int) {
	var slice []int
	slice = make([]int, end-start)
	for i, _ := range array {
		if i+1 >= start && i < end-1 {
			for y := start - 1; y < end-1; y++ {
				if i == y {
					slice[y-start+1] = array[i]
				}
			}
		}
	}
	fmt.Println(slice)

}

func main() {
	array1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	splitArray(array1, 3, 7)
}
