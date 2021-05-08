package main

import "fmt"

func main() {
	c := make(chan int)
	go task(c, 2)
	/*for i := 0; i < 5; i++ {
		go func(c chan int) {
			c <- cap(c)
		}(c)
	}*/
	fmt.Print(<-c)
}

func task(c chan int, n int) int {
	c <- n + 1
	return <-c
}

/*
Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
Функция должна называться task().
*/
