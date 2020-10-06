package main

import "fmt"

func main() {
	fmt.Println(vote(0, 1, 1))
}

func vote(x int, y int, z int) int {
	var result int
	var countZero, countOne int
	var slice []int = []int{x, y, z}
	for _, value := range slice {
		if value == 0 {
			countZero++
		} else {
			countOne++
		}
	}
	if countOne > countZero {
		result = 1
		return result
	} else if countOne < countZero {
		return result
	}
	return result
}

/*
Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее
аргументов x, y, z встречается чаще.

Входные данные

Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).

Выходные данные

Необходимо вывести  значение функции от x, y и z.

Sample Input:

0 0 1
Sample Output:

0
*/
