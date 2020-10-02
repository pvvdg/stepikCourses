package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	var value, count int
	var slice []int
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &value)
		slice = append(slice, value)
	}
	var min int = slice[0]
	for _, val := range slice {
		if val < min {
			min = val
		}
	}
	for _, val := range slice {
		if val == min {
			count++
		}
	}

	fmt.Printf("%d", count)
}

/*
Количество минимумов

Найдите количество минимальных элементов в последовательности.



Входные данные

Вводится натуральное число N, а затем N чисел.

Выходные данные

Выведите количество минимальных элементов.

Sample Input:

3
21 11 4


Sample Output:

1
*/
