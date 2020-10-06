package main

import "fmt"

func main() {
	fmt.Println(fibonacci(6))
}

func fibonacci(n int) int {
	var fin int
	if n == 1 {
		fin = 1
		return fin
	} else if n == 2 {
		fin = 1
		return fin
	} else if n > 1 {
		fi1, fi2 := 1, 1
		for i := 1; i < n-1; i++ {
			fin = fi1 + fi2
			fi1 = fi2
			fi2 = fin
		}

	}
	return fin
	//print your code here
}

/*
Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1. Начало ряда Фибоначчи выглядит
 следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ... Напишите функцию, которая по указанному натуральному n возвращает φn.

Входные данные

Вводится одно число n.

Выходные данные

Необходимо вывести  значение φn.

Sample Input:

4

Sample Output:

3

*/
