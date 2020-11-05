package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cache := make(map[int]int, 0)
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	strSlice := strings.Split(str, " ")
	strNew := strings.Join(strSlice, "")
	intSlice := make([]int, 0)
	fmt.Println(strNew)
	fmt.Println(strconv.Atoi(string(strNew[1])))
	/*for i := 0; i < 10; i++ {
		intSlice[i], _ = strconv.Atoi(string(strNew[i]))
	}*/
	fmt.Println(intSlice)
	fmt.Println(cache)
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
