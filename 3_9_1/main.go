package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan struct{}, 1)
	go func() {
		// ch3 <- struct{}{}
		ch2 <- 2
	}()
	fmt.Println(<-calculator(ch1, ch2, ch3))
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		select {
		case res := <-firstChan:
			ch <- res * res
		case res := <-secondChan:
			ch <- res * 3
		case <-stopChan:
			fmt.Println("cancel func")
			return
		}
	}()
	return ch
}
