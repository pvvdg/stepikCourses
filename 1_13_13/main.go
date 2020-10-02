package main

import (
	"fmt"
	"os"
)

func main() {
	var a, nFib int
	count := 1
	var nfib1, nfib2 int = 0, 1
	fmt.Fscan(os.Stdin, &a)
	if a > 1 {
		for {
			nFib = nfib1 + nfib2
			nfib1 = nfib2
			nfib2 = nFib
			//fmt.Println(nfib1, " ", nfib2, " ", nFib)
			count++

			if a == nFib {
				break
			} else if nFib > a {
				count = -1
				break
			}
		}
		fmt.Println(count)

	}
}

/*

Номер числа Фибоначчи

Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n,
 что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу.

Sample Input:

8

Sample Output:

6

*/
