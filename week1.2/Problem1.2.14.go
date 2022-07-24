package main

import "fmt"

func main() {
	x := 360330
	lastDig := x % 10
	secondDig := x % 100 / 10
	thirdDig := x % 1000 / 100
	fourthDig := x % 10000 / 1000
	fifthDig := x % 100000 / 10000
	sixthDig := x % 1000000 / 100000
	if sixthDig+fifthDig+fourthDig == lastDig+secondDig+thirdDig {
		fmt.Println("You've got a lucky number!")
	} else {
		fmt.Println("Sorry, no luck this time around :(")
	}

}
