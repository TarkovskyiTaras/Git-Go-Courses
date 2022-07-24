package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//3.11*. Повернуть матрицу(Двухмерный массив). Матрица квадратная. Пользователь вводит угол кратный 90.

func main() {

	// input from the keyboard start
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the side of the matrix")
	inputSideDim, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter an angle you want it to to rotate on")
	inputAngle, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	inputSideDim = strings.TrimSpace(inputSideDim)
	inputAngle = strings.TrimSpace(inputAngle)

	inputSideDimInt, err := strconv.Atoi(inputSideDim)
	if err != nil {
		log.Fatal(err)
	}
	inputAngleInt, err := strconv.Atoi(inputAngle)
	if err != nil {
		log.Fatal(err)
	}
	// input from the keyboard end

	//random numb
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	//random numb ends

	sideOfSlice := inputSideDimInt
	lastIndex := sideOfSlice - 1

	sliceOfSlices := make([][]int, sideOfSlice)
	for i, _ := range sliceOfSlices {

		sliceOfSlices[i] = make([]int, sideOfSlice)
	}
	sliceOfSlices90 := make([][]int, sideOfSlice)
	for i, _ := range sliceOfSlices90 {
		sliceOfSlices90[i] = make([]int, sideOfSlice)
	}

	sliceOfSlices180 := make([][]int, sideOfSlice)
	for i, _ := range sliceOfSlices180 {
		sliceOfSlices180[i] = make([]int, sideOfSlice)
	}

	sliceOfSlices270 := make([][]int, sideOfSlice)
	for i, _ := range sliceOfSlices270 {
		sliceOfSlices270[i] = make([]int, sideOfSlice)
	}

	//assigning values to the elements
	for i, x := range sliceOfSlices {
		for j, _ := range x {

			//random number choosing
			randNumb := rand.Intn(90) + 10
			sliceOfSlices[i][j] = randNumb
		}
	}
	fmt.Println()
	fmt.Println("original matrix: ")
	for _, x := range sliceOfSlices {
		fmt.Println(x)
	}

	fmt.Println("----------------------")
	fmt.Println("rotated matrix:")
	// making 90
	if inputAngleInt == 90 {
		for i, x := range sliceOfSlices {
			for j, _ := range x {
				sliceOfSlices90[j][lastIndex-i] = sliceOfSlices[i][j]
			}
		}
		for _, x := range sliceOfSlices90 {
			fmt.Println(x)
		}
	}

	// making 180
	if inputAngleInt == 180 {
		for i, x := range sliceOfSlices {
			for j, _ := range x {
				sliceOfSlices180[lastIndex-i][lastIndex-j] = sliceOfSlices[i][j]
			}
		}
		for _, x := range sliceOfSlices180 {
			fmt.Println(x)
		}
	}

	//making 270
	if inputAngleInt == 270 {
		for i, x := range sliceOfSlices {
			for j, _ := range x {
				sliceOfSlices270[lastIndex-j][i] = sliceOfSlices[i][j]
			}
		}
		for _, x := range sliceOfSlices270 {
			fmt.Println(x)
		}
	}

	// 360
	if inputAngleInt == 360 {
		for _, x := range sliceOfSlices {
			fmt.Println(x)
		}
	}

}
