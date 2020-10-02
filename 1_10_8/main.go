package main

import (
	"fmt"
	"os"
)

func main() {
	var number1, number2 int
	fmt.Fscan(os.Stdin, &number1, &number2)
	if number1 >= 0 && number1 <= 10000 && number2 >= 0 && number2 <= 10000 {
		arrNum1, arrNum2 := inArray(number1), inArray(number2)
		searchingForTheSame(arrNum1, arrNum2)
	}

}

func inArray(m int) []int {
	var arr []int
	var tmp int
	for m != 0 {
		tmp = m % 10
		arr = append(arr, tmp)
		m /= 10
	}
	return arr
}

func searchingForTheSame(arr1, arr2 []int) {
	for i := len(arr1) - 1; i >= 0; i-- {
		for j := len(arr2) - 1; j >= 0; j-- {
			if arr1[i] == arr2[j] {
				fmt.Print(arr1[i], " ")
			}
		}
	}
}

/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

Входные данные

Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.

Выходные данные

Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения
 в первом числе.

Внимание! Если вам сложно решить эту задачу, пропустите и изучайте материал дальше, затем вернетесь к этой задаче позже.

Sample Input:

564 8954
Sample Output:

5 4
*/
