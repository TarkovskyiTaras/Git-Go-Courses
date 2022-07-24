package main

//3.6. Требуется найти самую длинную непрерывную цепочку нулей  и единиц в последовательности цифр.

import "fmt"

func main() {

	givenString := "1000300000012411fdsafsd2fds"
	zeroCounter := 0
	maxZeroInARow := 0
	zeroStartPoint := 0
	zeroEndPoint := 0
	oneCounter := 0
	maxOneInARow := 0
	oneStartPoint := 0
	oneEndPoint := 0

	for i, x := range givenString {

		if string(x) == "0" {
			zeroCounter++

		}
		if maxZeroInARow < zeroCounter {
			maxZeroInARow = zeroCounter
			zeroEndPoint = i
			zeroStartPoint = i - maxZeroInARow + 1
		}
		if string(x) != "0" {
			zeroCounter = 0
		}

		if string(x) == "1" {
			oneCounter++

		}
		if maxOneInARow < oneCounter {
			maxOneInARow = oneCounter
			oneEndPoint = i
			oneStartPoint = i - maxOneInARow + 1
		}
		if string(x) != "1" {
			oneCounter = 0
		}

	}
	fmt.Println("max zero in a row: ", maxZeroInARow)
	fmt.Printf("zeroStartPoint, zeroEndPoint: (%d ; %d)\n", zeroStartPoint, zeroEndPoint)
	fmt.Println("max one in a row: ", maxOneInARow)
	fmt.Printf("oneStartPoint, oneEndPoint: (%d ; %d)\n", oneStartPoint, oneEndPoint)
}
