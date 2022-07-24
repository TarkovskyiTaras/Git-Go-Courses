package main

import (
	"fmt"
	"math"
)

//2.2 Проверить или является число простым

func main() {
	n := 37
	nSqRoot := int(math.Sqrt(float64(n)))
	var primeDet bool

	for x := 1; x <= nSqRoot; x++ {
		if x > 1 && n%x == 0 {
			primeDet = false
			break
		}

		if n%x != 0 {
			primeDet = true
		}

	}
	if primeDet == true {
		fmt.Println("prime")
	} else {
		fmt.Println("composite")
	}
}
