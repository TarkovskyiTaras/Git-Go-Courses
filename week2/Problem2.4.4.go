package main

import (
	"fmt"
	"math"
)

//4) Посчитать сколько цифр(цифра задается пользователем) в массиве

func main() {
	givenArray := [10]int{1, 22, 243, 22144, 5, 6, 7, 8, 9, 10}
	numbOfElemToCountDig := 3
	valueElemToCountDig := givenArray[numbOfElemToCountDig]
	orderOfMagn := 1
	for l := 1; true; l++ {
		y := valueElemToCountDig / int(math.Pow(10, float64(orderOfMagn)))
		fmt.Println(y, " = ", valueElemToCountDig, " / ", int(math.Pow(10, float64(orderOfMagn))))
		orderOfMagn++
		fmt.Println(orderOfMagn)
		if y < 1 {
			orderOfMagn--
			fmt.Println(orderOfMagn)
			break
		}
	}
	fmt.Printf("The number of digits in the %d element is %d.", numbOfElemToCountDig, orderOfMagn)
}
