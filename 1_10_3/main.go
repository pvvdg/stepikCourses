package main

import (
	"fmt"
	"os"
)

func main() {
	var numberOfSequence int
	fmt.Fscan(os.Stdin, &numberOfSequence)
	var sum int
	for i := 0; i < numberOfSequence; i++ {
		var tmp int
		fmt.Fscan(os.Stdin, &tmp)
		if tmp%8 == 0 && tmp >= 10 && tmp <= 99 {
			sum += tmp
		}

	}
	fmt.Println(sum)

}

/*


Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности,
во второй строке -- n чисел, входящих в данную последовательность.

Sample Input:

5
38 24 800 8 16

Sample Output:

40

*/
