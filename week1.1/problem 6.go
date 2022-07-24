package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func return2() (float64, float64) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter 2 numbers:")
	input1, _ := reader.ReadString('\n')
	input2, _ := reader.ReadString('\n')
	//if err != nil {
	//log.Fatal(err)
	//}
	inputStripped1 := strings.TrimSpace(input1)
	inputStripped2 := strings.TrimSpace(input2)
	inputFloat1, _ := strconv.ParseFloat(inputStripped1, 64)
	//if err != nil {
	//log.Fatal(err)
	//}
	inputFloat2, _ := strconv.ParseFloat(inputStripped2, 64)
	//if err != nil {
	//log.Fatal(err)
	//}
	return inputFloat1, inputFloat2
}

func main() {
	for x := 1; true; x++ {
		numb1, numb2 := return2()
		if numb1+numb2 >= 11 && numb1+numb2 <= 19 {
			fmt.Println("The sum is:", numb1+numb2)
		} else {
			fmt.Println("The sum is out of range.")
		}

	}
}
