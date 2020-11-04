package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {

	var pass string
	fmt.Scan(&pass)
	fmt.Println(check(pass))

}

func check(text string) string {
	if utf8.RuneCountInString(text) >= 5 {
		count := 0
		passInRune := []rune(text)
		for _, val := range passInRune {
			if unicode.IsDigit(val) || unicode.Is(unicode.Latin, val) {
				count++
			}
		}
		if count == utf8.RuneCountInString(text) {
			return "Ok"
		}
	}
	return "Wrong password"
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
