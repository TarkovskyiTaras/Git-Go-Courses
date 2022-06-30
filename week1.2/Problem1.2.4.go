package main

import "fmt"

func main() {
	x := 72.0 //in km\h
	fmt.Printf("\nFirst speed is %f km/h\n", x)
	y := 20.0 //in m\s
	fmt.Printf("Second speed is %f m/s\n", y)

	if x*1000/(60*60) > y {
		fmt.Println("\nFirst speed > Second Speed")
	} else if x*1000/(60*60) < y {
		fmt.Println("\nSecond speed > First speed")
	} else {
		fmt.Println("\nFirst speed = Second speed")
	}
}
