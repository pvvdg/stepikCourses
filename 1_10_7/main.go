package main

import (
	"fmt"
	"os"
)

func main() {
	var x, p, y, year int
	fmt.Fscan(os.Stdin, &x, &p, &y)
	numberOfCentsInRuble := 100
	x = x * numberOfCentsInRuble
	y = y * numberOfCentsInRuble
	for {
		if x <= y {
			x += (x * p / numberOfCentsInRuble)
			year++
		} else if x >= y {
			break
		}

	}
	fmt.Println(year)
}

/*
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек
 отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не
 менее y рублей.
Входные данные
Программа получает на вход три натуральных числа: x, p, y.

Выходные данные
Программа должна вывести одно целое число.
Sample Input:
100
10
200
Sample Output:
8
*/
