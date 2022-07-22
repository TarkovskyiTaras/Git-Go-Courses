package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//b) Поменять наибольший и наименьший столбик в двухмерном массиве.

func main() {
	timeNow := time.Now()
	seconds := timeNow.Unix()
	rand.Seed(seconds)
	sliceLength := 10
	sliceHeight := 10
	givenSlice := make([][]int, sliceHeight)

	for i, _ := range givenSlice {
		givenSlice[i] = make([]int, sliceLength)
	}

	for _, x := range givenSlice {
		for j, _ := range x {
			randomNumb := rand.Intn(90) + 10
			x[j] = randomNumb
		}
	}
	minSumColumn := math.Inf(+1)
	maxXSumColumn := math.Inf(-1)
	minColumnNumb := 0
	maxColumnNumb := 0
	sumOfColEl := 0
	for z := 0; z < sliceHeight; z++ {
		for _, x := range givenSlice {
			for j, y := range x {
				if j == z {
					sumOfColEl += y
				}
			}
		}

		//---------------------------------------------------
		fmt.Println("sumofelcol: ", sumOfColEl)
		//---------------------------------------------------

		if float64(sumOfColEl) > maxXSumColumn {
			maxXSumColumn = float64(sumOfColEl)
			maxColumnNumb = z
		}
		if float64(sumOfColEl) < minSumColumn {
			minSumColumn = float64(sumOfColEl)
			minColumnNumb = z
		}
		sumOfColEl = 0
	}

	// --------------------------------------------------
	fmt.Println("max", maxXSumColumn)
	fmt.Println("min", minSumColumn)
	fmt.Println("maxcolumnnumb", maxColumnNumb)
	fmt.Println("mincolumnnumb", minColumnNumb)
	//---------------------------------------------------

	for _, x := range givenSlice {
		fmt.Println(x)
	}
	fmt.Println("__________________________")

	minColumnCopy := make([]int, sliceHeight)
	maxColumnCopy := make([]int, sliceHeight)

	for i, x := range givenSlice {
		for j, y := range x {
			if j == minColumnNumb {
				for l, _ := range minColumnCopy {
					if l == i {
						minColumnCopy[l] = x[j]
					}
				}
			}
			if j == maxColumnNumb {
				for l, _ := range maxColumnCopy {
					if l == i {
						maxColumnCopy[l] = y
					}
				}
			}
		}
	}

	for i, x := range givenSlice {
		for j, _ := range x {
			if j == maxColumnNumb {
				for l, z := range minColumnCopy {
					if l == i {
						x[j] = z
					}
				}

			}
			if j == minColumnNumb {
				for l, z := range maxColumnCopy {
					if l == i {
						x[j] = z
					}
				}

			}
		}
	}
	for _, x := range givenSlice {
		fmt.Println(x)
	}
}
