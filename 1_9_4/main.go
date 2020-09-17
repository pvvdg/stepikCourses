package main

import (
	"fmt"
	"os"
)

const (
	limitFirst = 100000
	limitLast  = 999999
)

func main() {
	var happyNumber int
	var sumLeft, sumRight int
	var n1, n2, n3, n4, n5, n6 int
	fmt.Fscan(os.Stdin, &happyNumber)
	if happyNumber >= limitFirst && happyNumber <= limitLast {
		n1 = happyNumber / 100000
		n2 = happyNumber / 10000 % 10
		n3 = happyNumber / 1000 % 10
		n4 = happyNumber / 100 % 10
		n5 = happyNumber / 10 % 10
		n6 = happyNumber % 10
		sumLeft = n1 + n2 + n3
		sumRight = n4 + n5 + n6
		if sumLeft == sumRight {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		//fmt.Println(sumLeft, " ", sumRight)

	} else {
		fmt.Println("NO")
	}

}

/*Определите является ли билет счастливым. Счастливым считается билет,
в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
Формат входных данных

На вход подается номер билета - одно шестизначное  число.

Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".*/
