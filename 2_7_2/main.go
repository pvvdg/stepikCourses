package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	str := ""
	fmt.Scan(&str)
	lenStr := utf8.RuneCountInString(str)
	count := 0
	for _, el := range str {
		if unicode.Is(unicode.Latin, el) {
			count++
		}
	}
	if lenStr > 0 && lenStr <= 1000 && lenStr == count {
		fmt.Println(add(str))
	}

}

func add(str string) string {
	newStr := ""
	for _, el := range str {
		newStr += string(el) + "*"
	}
	lenNewStr := utf8.RuneCountInString(newStr)
	return newStr[:lenNewStr-1]
}

/*
Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой
	буквой и после последней символ ‘*’ добавлять не нужно).

Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.

Выходные данные

Вывести строку, которая получится после добавления символов '*'.

Sample Input:

LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO

Sample Output:

L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
*/
