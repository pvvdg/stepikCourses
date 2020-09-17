package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input integer value: ")
	var a int
	fmt.Fscan(os.Stdin, &a)
	a1 := a / 100
	a2 := a / 10 % 10
	a3 := a % 10

	switch {
	case a1 != a2 && a1 != a3 && a2 != a3:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}

/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".
*/
