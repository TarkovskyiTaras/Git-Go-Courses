package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//3.5. а) Поменять наибольшую строку массива с наименьшей в двухмерном массиве. (Размер - сумма элементов в строке)

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

	max := math.Inf(-1)
	min := math.Inf(+1)
	numbOfRowMin := 0
	numbOfRowMax := 0
	sumOfElInRow := 0

	for i, x := range givenSlice {
		for j, _ := range x {
			randNumb := rand.Intn(90) + 10
			x[j] = randNumb
			sumOfElInRow += x[j]
			if j == sliceLength-1 {
				if float64(sumOfElInRow) > max {
					max = float64(sumOfElInRow)
					numbOfRowMax = i

				}
				if float64(sumOfElInRow) < min {
					min = float64(sumOfElInRow)
					numbOfRowMin = i

				}
				fmt.Println("sumofel", sumOfElInRow)
				fmt.Println("min", min)
				fmt.Println("numbofrowMIN", numbOfRowMin)
				fmt.Println("max", max)
				fmt.Println("numberofrowMAX", numbOfRowMax)
				sumOfElInRow = 0
			}
		}

	}
	for _, x := range givenSlice {
		fmt.Println(x)
	}
	fmt.Println("__________________________________")

	copyMinRowSlice := givenSlice[numbOfRowMin]
	givenSlice[numbOfRowMin] = givenSlice[numbOfRowMax]
	givenSlice[numbOfRowMax] = copyMinRowSlice

	for _, x := range givenSlice {
		fmt.Println(x)
	}
}
