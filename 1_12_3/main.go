package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	fmt.Println(max(array))
}

func max(array [5]int) int {
	var maxNumber int
	for _, value := range array {
		if value > maxNumber {
			maxNumber = value
		}
	}
	return maxNumber
}

/*


На ввод подаются пять чисел, которые записываются в массив. Однако эта часть программы уже написана.

Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.

Sample Input:

2
3
56
45
21

Sample Output:

56

*/
