package main

import (
	"fmt"
	"math"
)

func main() {
	threeDigNumb := 122
	firstDig := math.Floor(float64(threeDigNumb) / 100)
	fmt.Println(firstDig)
	secondDig := math.Floor(float64(threeDigNumb/10) - firstDig*10)
	fmt.Println(secondDig)
	thirdDig := float64(threeDigNumb) - (math.Floor(float64(threeDigNumb/10)))*10
	fmt.Println(thirdDig)
	matchCount := 0
	slice1 := []float64{firstDig, secondDig, thirdDig}
	slice2 := []float64{firstDig, secondDig, thirdDig}
	for _, x := range slice1 {
		for _, y := range slice2 {
			if x == y {
				matchCount += 1
			}
		}
	}

	if matchCount == 9 {
		fmt.Println("All three digits of the number are the same.")
	} else if matchCount == 3 {
		fmt.Println("No digits match in the number")
	} else if matchCount == 5 {
		fmt.Println("Two digits have matched in the number")
	}

}
