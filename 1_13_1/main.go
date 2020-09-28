package main

import (
	"fmt"
	"os"
)

func main() {
	var number, sum int
	fmt.Fscan(os.Stdin, &number)
	if number > 99 && number < 1000 {
		for number != 0 {
			sum += number % 10
			number /= 10
		}
	}
	fmt.Println(sum)

}

/*
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:

745

Sample Output:

16

*/
