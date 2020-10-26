package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")
	text = strings.Trim(text, "\r")
	fmt.Println(text)
	//lenText := utf8.RuneCountInString(text)
	textInRune := []rune(text)
	lenTextInRune := len(textInRune)
	midValueTextInRune := 0
	if lenTextInRune >= 3 {
		switch {
		case lenTextInRune == 3:
			midValueTextInRune = 2

		case lenTextInRune%2 == 0:
			midValueTextInRune = lenTextInRune / 2

		case lenTextInRune%2 != 0:
			midValueTextInRune = lenTextInRune/2 + 1

		}
	}
	partOfRuneText := textInRune[:midValueTextInRune]
	fmt.Println(partOfRuneText)
	countTrueRune := 0
	for i := len(partOfRuneText) - 1; i >= 0; i-- {
		for j := range partOfRuneText {
			if textInRune[i] == partOfRuneText[j] {
				countTrueRune++
				break
			}
		}
	}
	fmt.Println(countTrueRune)
	if countTrueRune == len(partOfRuneText) {
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
