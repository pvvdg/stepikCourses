package main

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	var str string
	fmt.Scan(&str)
	lenStr := utf8.RuneCountInString(str)
	count := 0
	for _, el := range str {
		if unicode.IsDigit(el) {
			count++
		}
	}
	if lenStr > 0 && lenStr <= 1000 && lenStr == count {
		fmt.Println(findMax(str))
	}
}

func findMax(str string) int {
	slice := make([]int, 0)
	for _, val := range str {
		intValue, _ := strconv.Atoi(string(val))
		slice = append(slice, intValue)
	}
	numMax := slice[0]
	for i := range slice {
		if slice[i] > numMax {
			numMax = slice[i]
		}
	}

	return numMax
}

/*


Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.

Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только десятичные цифры.

Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:

1112221112

Sample Output:

2

*/
