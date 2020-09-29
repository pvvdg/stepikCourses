package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var number int
	inPow := math.Pow(10, 7)
	fmt.Fscan(os.Stdin, &number)
	if number <= int(inPow) && number > 0 {
		fmt.Println(inArray(number))
	}

}

func inArray(m int) int {
	var result, sum, tmp int
	for m != 0 {
		tmp = m % 10
		sum += tmp
		m /= 10
	}
	switch {
	case sum >= 10:
		sum = sum/10 + sum%10
		result = sum
	case sum < 10:
		result = sum
	}
	return result
}

/*
Цифровой корень

Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр,
 на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации.
 Этот процесс повторяется до тех пор, пока не будет получена одна цифра.

Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .

По данному числу определите его цифровой корень.

Входные данные

Вводится одно натуральное число n, не превышающее 10^7.

Выходные данные
Вывести цифровой корень числа n.

Sample Input:

3456

Sample Output:

9

*/
