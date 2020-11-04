package main

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	var n string
	fmt.Scan(&n)
	strToRune := []rune(n)
	count := 0
	strLen := utf8.RuneCountInString(n)
	for _, val := range strToRune {
		if unicode.IsDigit(val) {
			count++
		}
	}
	if strLen == count {
		fmt.Println(toSqrt(n))
	}

}

func toSqrt(n string) string {
	slice := make([]int, 0)
	for _, val := range n {
		intValue, _ := strconv.Atoi(string(val))
		slice = append(slice, intValue)
	}
	for i := 0; i < len(slice); i++ {
		slice[i] *= slice[i]
	}
	newStr := ""
	for _, val := range slice {
		newStr += strconv.Itoa(val)
	}
	return newStr
}

/*


На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.

Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

Sample Input:

9119

Sample Output:

811181

*/
