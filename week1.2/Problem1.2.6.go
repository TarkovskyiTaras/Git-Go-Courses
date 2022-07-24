package main

import "fmt"

func main() {
	v := 100.0 //speed km/h
	s := 100.0 //distance km

	t := s / v //time
	fmt.Printf("\n%.2f hour(s) will take to reach the destination point.\n", t)
}
