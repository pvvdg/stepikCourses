package main

import "fmt"

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		sum := 0
		for {
			select {
			case val := <-arguments:
				sum += val
			case <-done:
				out <- sum
				return
			}
		}
	}()
	return out
}

func main() {
	arguments := make(chan int)
	done := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			arguments <- i
		}
		done <- struct{}{}
	}()

	/*for v := range calculator(arguments, done) {
		fmt.Println(v)
	}*/
	fmt.Println(<-calculator(arguments, done))
}

/*
 Вам необходимо написать функцию calculator следующего вида:

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.

Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу.
 Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.

Функция calculator должна быть неблокирующей, сразу возвращая управление.

Выходной канал должен быть закрыт после выполнения всех оговоренных условий, если вы этого не сделаете, то превысите предельное время работы.
*/
