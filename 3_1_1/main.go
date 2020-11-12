package main

import "fmt"

func main() {
	n := 0 // from bufio
	i := 0 // start value in for

	bufNumbers := make([]int, 0)
	resWork := 0
	cache := make(map[int]int)

	for i != 10 { // read all numbers and append the'is in slice
		fmt.Scan(&n)
		if value, ok := cache[n]; ok {
			bufNumbers = append(bufNumbers, value)
		} else {
			resWork = work(n)
			cache[n] = resWork
			bufNumbers = append(bufNumbers, resWork)
		}
		i++
	}
	for _, valueFromBufNumbers := range bufNumbers {
		fmt.Print(valueFromBufNumbers, " ")
	}
}

func work(x int) int {
	x--
	return x
}

/*

Внутри функции main (объявлять функцию не нужно) необходимо написать программу:

На стандартный ввод подается 10 целых чисел, разделенных пробелами (числа могут повторяться). Для чтения из стандартного ввода импортирован пакет fmt.

Вам необходимо вычислить результат выполнения функции work для каждого из полученных чисел. Функция work имеет следующий вид:

func work(x int) int

Результаты вычислений , разделенные пробелами, должны быть напечатаны в строку.

Однако работа функции work занимает слишком много времени. Выделенного вам времени выполнения не хватит на последовательную обработку каждого числа, поэтому необходимо реализовать кэширование уже готовых результатов и использовать их в работе.

После завершения работы программы результат выполнения будет дополнен информацией о соблюдении установленного лимита времени выполнения.

Sample Input:

3 1 5 2 3 5 3 0 3 4

1 5 0 -1 5 0 7 0 6 7

Sample Output:

2 0 6 1 2 6 2 -1 2 5 time limit ok

*/
