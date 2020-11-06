package main

import "fmt"

func main() {
	k := 10
	n := 0
	i := 0
	bufSlice := make([]int, 0)
	for i != k {
		fmt.Scan(&n)
		bufSlice = append(bufSlice, n)
		i++
	}
	resWork := 0
	cache := make(map[int]int, 0)
	cache[bufSlice[0]] = work(bufSlice[0])
	for keyCache := range cache {
		for _, valBuf := range bufSlice {
			if valBuf != keyCache {
				resWork = work(valBuf)
				cache[valBuf] = resWork
			}
		}
	}
	fmt.Println(cache)
	for _, valB := range bufSlice {
		for keyC, valC := range cache {
			if valB == keyC {
				fmt.Print(valC)
				break
			} else {
				fmt.Print(work(valB))
				break
			}

		}
	}
	//
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

Sample Output:

2 0 6 1 2 6 2 -1 2 5 time limit ok

*/
