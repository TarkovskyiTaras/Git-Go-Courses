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
	array2 := [10]int{}
	sumOfEven1 := 0
	sumOfEven2 := 0
	for x := 0; x < 10; x++ {
		randNumber := rand.Intn(51) + 25
		array1[x] = randNumber
	}
	for x := 0; x < 10; x++ {
		randNumber := rand.Intn(51) + 25
		array2[x] = randNumber
	}

	for i, _ := range array1 {
		if array1[i]%2 == 0 {
			sumOfEven1 += 1
		}
	}
	for i, _ := range array2 {
		if array2[i]%2 == 0 {
			sumOfEven2 += 1
		}
	}
	fmt.Println(array1, array2)
	fmt.Println(sumOfEven1, sumOfEven2)
	if sumOfEven1 > sumOfEven2 {
		fmt.Println("first array contains more even elements")
	}
	if sumOfEven2 > sumOfEven1 {
		fmt.Println("second array contains more even elements")
	}
}
