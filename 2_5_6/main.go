package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {

	var pass string
	fmt.Scan(&pass)

}

func checkDigit(string text) bool {
	passInRune := []rune(text)
	countLatin := 0
	for _, val := range passInRune {
		if unicode.IsDigit(val) {
			countLatin++
		}
	}
	if countLatin == utf8.RuneCountInString(text) {
		return true
	}

}

func checkLatin(string val) bool {
	passInRune := []rune(text)
	countLatin := 0
	for _, val := range passInRune {
		if unicode.Is(unicode.Latin, val) {
			countLatin++
		}
	}
	if countLatin == utf8.RuneCountInString(text) {
		return true
	}
}

func checkCount(string val) bool {
	if utf8.RuneCountInString(pass) >= 5 {
		return true
	} else {
		false
	}
}

/*


Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля
должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита. На вход
подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:

fdsghdfgjsdDD1

Sample Output:

Ok


*/
