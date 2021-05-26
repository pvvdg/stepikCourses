package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan struct{}, 1)
	go func() {
		fmt.Println("In ", <-ch3)
		ch3 <- struct{}{}
		fmt.Println(<-ch3)
	}()
	defer close(ch3)
	fmt.Println(<-calculator(ch1, ch2, ch3))
	// stopChan := make(chan struct{}, 1)
	// go func() {
	// fmt.Println(<-stopChan)
	// }()
	// fmt.Println(<-calculator(ch1, ch2, stopChan))
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		select {
		case res := <-firstChan:
			ch <- res * res
		case res := <-secondChan:
			ch <- res * 3
		case <-stopChan:
			return
		}
	}()
	return ch
}
