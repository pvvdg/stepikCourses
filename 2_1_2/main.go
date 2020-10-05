package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin())
}

func findMin() int {
	n := 4
	number := 0
	slice := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		slice = append(slice, number)
	}
	min := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min
}

/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:

4 5 6 7
Sample Output:

4
*/
