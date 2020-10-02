package main

import (
	"fmt"
	"os"
)

func main() {
	var number, n, tmpNumber int
	var sliceTmpNumbers []int
	fmt.Fscan(os.Stdin, &number, &n)
	for number != 0 {
		tmpNumber = number % 10
		if tmpNumber != n {
			sliceTmpNumbers = append(sliceTmpNumbers, tmpNumber)
		}
		number /= 10
	}
	for i := len(sliceTmpNumbers) - 1; i >= 0; i-- {
		fmt.Print(sliceTmpNumbers[i])
	}
	fmt.Println()

}

/*
Из натурального числа удалить заданную цифру.

Входные данные

Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные

Вывести число без заданных цифр.

Sample Input:

38012732
3

Sample Output:

801272
*/
