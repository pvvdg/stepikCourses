package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Fscan(os.Stdin, &number)
	if number > 99 && number < 1000 && number%10 != 0 {
		for number != 0 {
			fmt.Print(number % 10)
			number /= 10
		}
	}
}

/*


Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:

843

Sample Output:

348


*/
