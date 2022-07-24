package main

import (
	"errors"
	"fmt"
	"log"
)

//2.4 Сдвинуть массив на заданное количество позиций // {1,2,3,4,5} -> 2 = {4,5,1,2,3}; {1,2,3,4,5} -> 3 = {3,4,5,1,2}
//Количество позиций и массив указывает пользователь

func main() {
	givenSlice := []int{1, 2, 3, 4, 5}
	shiftNumber := 2
	if shiftNumber > len(givenSlice) {
		var err error
		err = errors.New("incorrect value")
		log.Fatal(err)

	}

	for _, x := range givenSlice {
		givenSlice = append(givenSlice, x)
	}
	shiftNumber = len(givenSlice)/2 - shiftNumber

	resultSlice := givenSlice[shiftNumber : shiftNumber+len(givenSlice)/2]
	fmt.Println(resultSlice)
}
