package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b uint
	var lim uint = 100
	fmt.Fscan(os.Stdin, &a, &b)
	if a <= lim && b <= lim && a < b {
		var sum uint
		for i := a; i <= b; i++ {
			sum = sum + i
		}
		fmt.Println(sum)
	}
}
