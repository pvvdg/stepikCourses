package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, max int
	fmt.Fscan(os.Stdin, &a, &b)
	if a <= b {
		for i := a; i <= b; i++ {
			if i >= a && i%7 == 0 {
				max = i
			}
		}
		if max > 0 {
			fmt.Println(max)
		} else {
			fmt.Println("NO")
		}
	}
}

/*
Самое большое число, кратное 7

Найдите самое большее число на отрезке от a до b, кратное 7 .

Входные данные
Вводится два целых числа a и b (a≤b).

Выходные данные
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.

Sample Input:

100
500

Sample Output:

497

*/
