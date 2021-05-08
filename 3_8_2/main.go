package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go task2(c, "Stroka")
	for v := range c {
		fmt.Println(v)
	}
}

func task2(c chan string, str string) string {
	for i := 0; i < 5; i++ {
		c <- str + " "
	}
	close(c)
	return <-c
}

/*
Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.

Функция должна называться task2().
*/
