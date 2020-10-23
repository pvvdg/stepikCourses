package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	lastСharacterValue := utf8.RuneCountInString(text) - 3
	textInRune := []rune(text)
	dot := '.'
	if unicode.IsUpper(textInRune[0]) && textInRune[lastСharacterValue] == dot {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}

/*
На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной
 буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong

Sample Input:

Быть или не быть.
Sample Output:

Right
*/
