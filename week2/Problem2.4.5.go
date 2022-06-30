package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	array1 := [10]int{}
	valueEven := 0
	valueOdd := 0

	for x := 0; x < 10; x++ {

		for x := 1; true; x++ {
			value := rand.Intn(100) + 1
			if value%2 == 0 {
				valueEven = value
				break
			}

		}
		for x := 1; true; x++ {
			value := rand.Intn(100) + 1
			if value%2 != 0 {
				valueOdd = value
				break
			}

		}
		for i, _ := range array1 {
			if x == i && i%2 == 0 {
				array1[i] = valueEven
			}
			if x == i && i%2 != 0 {
				array1[i] = valueOdd
			}
		}

	}
	fmt.Println(array1)

}
