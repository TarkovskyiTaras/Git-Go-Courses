package main

import (
	"fmt"
	"strings"
)

//2.5 Заменить в строке все символы 'a' на '@' и вывести новую строку // "array is a data type" -> "@rr@y is @ d@t@ type"

func main() {

	givenString := "array is a data type"

	replacer := strings.NewReplacer("a", "@")

	resultString := replacer.Replace(givenString)

	fmt.Println(resultString)

}
