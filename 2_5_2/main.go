package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")  //delete '/n' from string
	text = strings.Trim(text, "\r")  //delete '/ r' from text in OS Windows 10 bash terminal
	text = strings.ToLower(text)     //lowercase conversion
	textInRune := []rune(text)       //convert string variable in rune
	lenTextInRune := len(textInRune) //count lenth in rune variable
	midValueTextInRune := 0          //creating variable which divides rune text by half

	if lenTextInRune >= 3 { // checking on half meaning
		switch {
		case lenTextInRune == 3:
			midValueTextInRune = 2

		case lenTextInRune%2 == 0:
			midValueTextInRune = lenTextInRune / 2

		case lenTextInRune%2 != 0:
			midValueTextInRune = lenTextInRune/2 + 1

		}
	}
	partOfRuneText := textInRune[:midValueTextInRune] //creating rune text by half

	var backPartOfRuneText string
	//coup rune values and adding in string variable backPartOfRuneText for checkig in text suffix
	for i := len(partOfRuneText) - 1; i >= 0; i-- {
		backPartOfRuneText += string(partOfRuneText[i])
	}
	// checking suffix in text variable
	if strings.HasSuffix(text, backPartOfRuneText) {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")
	}
}

/*
На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является паллиндромом -
 вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

Sample Input:

топот
Sample Output:

Палиндром
*/
