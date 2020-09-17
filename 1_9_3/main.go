package main

import (
	"fmt"
	"os"
)

const stat int = 10000

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	if n <= stat && n > 0 {
		calc(n, stat)
	} else if n == 0 {
		fmt.Println("0")
	}
}

func calc(n, i int) {
	if n/i == 0 {
		calc(n, i/10)
	} else {
		fmt.Println(n / i)
	}
}

/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000. (10к включительно)

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.
*/
