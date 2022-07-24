package main

import (
	"fmt"
	"strings"
)

//3.8 Напишите программу, которая позволяет пользователю ввести в консоли строку,
//переводит первый символ слов в верхний регистр и результат выводит в консоль

func main() {

	givenString := "nikita tarkovskyi is a dog"
	capitalizedString := strings.Title(givenString)
	fmt.Println(capitalizedString)
}
