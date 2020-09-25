package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	if n >= 4 {
		var slice []int
		var tmp int
		for i := 0; i < n; i++ {
			fmt.Fscan(os.Stdin, &tmp)
			slice = append(slice, tmp)
		}
		fmt.Println(slice[3])
	}
}

/*


Напишите программу, принимающая на вход число N (N≥4), а затем N целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:

5
41 -231 24 49 6

Sample Output:

49

}*/
