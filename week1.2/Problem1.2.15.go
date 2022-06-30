package main

import "fmt"

func main() {
	squareArea := 243.0
	circleArea := 123.0
	if squareArea/circleArea == 2/3.14 {
		fmt.Println("The square is inscribed in the circle.")
	} else if squareArea/circleArea == 4/3.14 {
		fmt.Println("The circle is inscribed in the square.")
	} else {
		fmt.Println("no fit, try again.")
	}
}
