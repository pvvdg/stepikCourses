package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	if 1 <= n && n <= 100 {
		var slice []int
		var inputValue int
		for i := 0; i < n; i++ {
			fmt.Fscan(os.Stdin, &inputValue)
			slice = append(slice, inputValue)
		}
		for i := 0; i < n; i++ {
			if i%2 == 0 || i == 0 {
				fmt.Printf("%d ", slice[i])
			}
		}

	}
}

/*


Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива,
индексы которых четны (0, 2, 4...).

Входные данные

Сначала задано число N — количество элементов в массиве (1≤N≤100). Далее через пробел записаны N чисел — элементы массива.
 Массив состоит из целых чисел.

Выходные данные

Необходимо вывести все элементы массива с чётными номерами.

Sample Input:

6
4 5 3 4 2 3

Sample Output:

4 3 2


*/
