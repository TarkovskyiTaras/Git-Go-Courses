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

func return2() (int, int) {

	reader := bufio.NewReader(os.Stdin)
	inputY, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputX, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputY = strings.TrimSpace(inputY)
	inputX = strings.TrimSpace(inputX)

	guessY, err := strconv.Atoi(inputY)
	if err != nil {
		log.Fatal(err)
	}
	guessX, err := strconv.Atoi(inputX)
	if err != nil {
		log.Fatal(err)
	}
	return guessY, guessX

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the side of the battlefield square:")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputNoSpace := strings.TrimSpace(input)
	inputInt, err := strconv.Atoi(inputNoSpace)
	if err != nil {
		log.Fatal(err)
	}

	width := inputInt
	height := inputInt

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(width) + 2
	if target > width-3 {
		if target%2 == 0 {
			target = width - (width/2 - 1)
		} else {
			target = width - (width/2 - 3)
		}
	}
	fmt.Printf("Guessed number: %d\n", target)
	sliceOfSlices := make([][]int, width)
	for i := 0; i < width; i++ {
		sliceOfSlices[i] = make([]int, width)
	}
	if seconds%2 == 0 {
		if target >= 2 && target <= width-8 {
			for j := target; j < target+6; j++ {
				sliceOfSlices[target][j] = 1
			}
		}
		if target <= width-3 && target > width-8 {
			for j := target; j > target-6; j-- {
				sliceOfSlices[target][j] = 1
			}
		}
	}
	if seconds%2 != 0 {
		if target >= 2 && target <= height-8 {
			for i := target; i < target+6; i++ {
				sliceOfSlices[i][target] = 1
			}
		}
		if target <= height-3 && target > height-8 {
			for i := target; i > target-6; i-- {
				sliceOfSlices[i][target] = 1
			}
		}
	}
	fmt.Println("Full vision battlefield")
	for i := 0; i < height; i++ {
		fmt.Println(sliceOfSlices[i])
	}
	fmt.Println("----------------------------")
	fmt.Println("Player's vision battlefield")

	sliceOfSlices2 := make([][]int, width)
	for i := 0; i < width; i++ {
		sliceOfSlices2[i] = make([]int, width)
	}
	for i := 0; i < height; i++ {
		fmt.Println(sliceOfSlices2[i])
	}

	fmt.Println("----------------------------")

	hitCount := 0
	missCount := 0

	for x := 1; true; x++ {

		guessY, guessX := return2()

		if sliceOfSlices[guessY][guessX] == 1 {
			hitCount += 1
			sliceOfSlices2[guessY][guessX] = 1

		} else {
			missCount += 1
			sliceOfSlices2[guessY][guessX] = 2
		}

		if hitCount == 6 {
			fmt.Println("Congratulations! You've sunk the enemy's warship.")
			fmt.Println("VICTORY!")
			break
		}
		fmt.Println("---------------------------")
		fmt.Println("Full vision battlefield")
		for i := 0; i < height; i++ {
			fmt.Println(sliceOfSlices[i])
		}
		fmt.Println("---------------------------")
		fmt.Println("Player's vision battlefield")
		for i := 0; i < height; i++ {
			fmt.Println(sliceOfSlices2[i])
		}
		fmt.Println("guessY:", guessY, "guessX:", guessX)
		fmt.Println("----------------------")
		fmt.Println("hitCount:", hitCount)
		fmt.Println("missCount:", missCount)

		if sliceOfSlices[guessY][guessX] == 1 {
			fmt.Println("----------------------")
			fmt.Println("You've hit!")
		} else {
			fmt.Println("----------------------")
			fmt.Println("You've missed!")
		}
		if guessY < target && sliceOfSlices[guessY][target] != 1 {
			fmt.Println("You might wanna aim lower!")
		}
		if guessY > target && sliceOfSlices[guessY][target] != 1 {
			fmt.Println("You might wanna aim higher!")
		}
		if guessX < target && sliceOfSlices[target][guessX] != 1 {
			fmt.Println("You might wanna aim right!")
		}
		if guessX > target && sliceOfSlices[target][guessX] != 1 {
			fmt.Println("You might wanna aim left!")
		}
	}

}
