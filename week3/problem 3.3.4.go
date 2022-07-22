package main

import (
	"fmt"
	"math/rand"
	"time"
)

//3.4. Создать массив, который будет состоять из диагонали матрици

func main() {
	sliceLength := 10
	sliceHeight := 10

	timeNow := time.Now()
	seconds := timeNow.Unix()
	rand.Seed(seconds)

	givenSlice := make([][]int, sliceHeight)
	for j, _ := range givenSlice {
		givenSlice[j] = make([]int, sliceLength)

	}

	for _, x := range givenSlice {
		for i, _ := range x {
			randNumb := rand.Intn(90) + 10
			x[i] = randNumb
		}
		fmt.Println(x) //given slice in a matrix shape
	}

	sliceOfDiag := make([]int, 0)

	for i, x := range givenSlice {
		for j, y := range x {
			if i == j {
				sliceOfDiag = append(sliceOfDiag, y)
			}

		}
	}
	fmt.Println("_________________________")
	fmt.Println(sliceOfDiag)
}
