package main

import "fmt"

func main() {
	c := make(chan struct{})
	go func(c chan struct{}) {
		work()
		close(c)
	}(c)
	<-c
}

func work() {
	fmt.Println("Done")
}

/*
Внутри функции main , вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов
 ее выполнения.
*/
