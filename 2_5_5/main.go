package main

import (
	"fmt"
	"strings"
)

func main() {
	var text, newText string
	fmt.Scan(&text)
	for _, val := range text {
		if strings.Count(text, string(val)) > 1 {
			continue
		} else {
			newText += string(val)
		}
	}
	fmt.Println(newText)
}

/*


Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

Sample Input:

zaabcbd

Sample Output:

zcd


*/
