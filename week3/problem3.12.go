package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//3.12*. Строки. Ввод пароля.
//		Написать программу ввода пароля пользователя. Пользователь вводит пароль 2 раза. Пароль должен удовлетворять следующим требованиям:
//		- длинна пароля должна быть от 8 до 20 символов
//		- в пароле должны быть как маленькие так и большие буквы
//		- должны быть цифры
//		- не должен содержать слов 'password', 'pass', 'gfhjkm'
//		- генерировать случайный пароль используя символы: a-z A-Z 0-9

func main() {
	fmt.Println("Would you like to use a random password? (y/n)")
	keyboardInput := keyboardEnter()

	if keyboardInput == "y" {
		fmt.Println("Your password is:")
		fmt.Println(passwordRandomizer())
		return
	}

	if keyboardInput == "n" {
		fmt.Println("Please enter a password, 8-20 symbols a-z." +
			"\nIt also should include at lease 1 uppercase letter and 1 digit")
	}
	fmt.Println("first enter:")
	passInput1 := keyboardEnter()
	fmt.Println("second enter:")
	passInput2 := keyboardEnter()

	acceptableChar, upperCaseMatch, oneDigitMatch := passwordChecker(passInput1)

	if passInput1 == passInput2 {
		if acceptableChar != true {
			fmt.Println("Unacceptable symbol has been found")
		}
		if upperCaseMatch != true {
			fmt.Println("No uppercase letters have been found")
		}
		if oneDigitMatch != true {
			fmt.Println("No numerical symbols have been found")
		}
		if acceptableChar == true && upperCaseMatch == true && oneDigitMatch == true {
			fmt.Println("your password has been successfully created")
		}
	} else {
		fmt.Println("two inputs don't match")
	}
}

func keyboardEnter() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	return input
}

func passwordRandomizer() string {
	var sliceLowerCaseLet []string
	var sliceUpperCaseLet []string
	var sliceNumbs []string
	var sliceRandomPass []string

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	passwordLength := rand.Intn(13) + 8

	for x := 'a'; x <= 'z'; x++ {
		sliceLowerCaseLet = append(sliceLowerCaseLet, string(x))
	}

	for x := 'A'; x <= 'Z'; x++ {
		sliceUpperCaseLet = append(sliceUpperCaseLet, string(x))
	}

	for x := 0; x <= 9; x++ {
		y := fmt.Sprintf("%d", x)
		sliceNumbs = append(sliceNumbs, y)
	}

	count := 0

	for x := 0; x < passwordLength; x++ {
		if count < passwordLength {
			randomChar1 := rand.Intn(len(sliceLowerCaseLet) - 1)
			sliceRandomPass = append(sliceRandomPass, sliceLowerCaseLet[randomChar1])
			count++
		}
		if count < passwordLength {
			randomChar2 := rand.Intn(len(sliceUpperCaseLet) - 1)
			sliceRandomPass = append(sliceRandomPass, sliceUpperCaseLet[randomChar2])
			count++
		}
		if count < passwordLength {
			randomChar3 := rand.Intn(len(sliceNumbs) - 1)
			sliceRandomPass = append(sliceRandomPass, sliceNumbs[randomChar3])
			count++
		}

		if count == passwordLength {
			break
		}
	}
	var randPassString string
	for _, x := range sliceRandomPass {
		randPassString += x
	}
	return randPassString
}

func passwordChecker(usersPassword string) (acceptableChar bool, upperCaseMatch bool, oneDigitMatch bool) {
	acceptableChar = false
	upperCaseMatch = false
	oneDigitMatch = false
	var sliceAllChars []string

	for x := 'a'; x <= 'z'; x++ {
		sliceAllChars = append(sliceAllChars, string(x))
	}
	for x := 'A'; x <= 'Z'; x++ {
		sliceAllChars = append(sliceAllChars, string(x))
	}
	for x := 0; x <= 9; x++ {
		sliceAllChars = append(sliceAllChars, fmt.Sprintf("%d", x))
	}
	for _, x := range usersPassword {
		acceptableChar = false
		for _, y := range sliceAllChars {
			if string(x) == y {
				acceptableChar = true
				break
			}
		}
	}

	for x := 'A'; x <= 'Z'; x++ {
		for _, y := range usersPassword {
			if y == x {
				upperCaseMatch = true
				break
			}
		}
	}

	for x := 0; x <= 9; x++ {
		for _, y := range usersPassword {
			if string(y) == fmt.Sprintf("%d", x) {
				oneDigitMatch = true
				break
			}
		}
	}
	return acceptableChar, upperCaseMatch, oneDigitMatch
}
