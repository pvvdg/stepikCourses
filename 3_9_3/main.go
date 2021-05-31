package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	n := 100
	ch1 := make(chan int, n)
	ch2 := make(chan int, n)
	out := make(chan int, n)

	for i := 0; i < n; i++ {
		ch1 <- i
		ch2 <- i
	}

	merge2Channels1(fn, ch1, ch2, out, n)

	for i := 0; i < n; i++ {
		fmt.Println(<-out)
	}

}

func fn(n int) int {
	time.Sleep(time.Millisecond * 10)
	return n
}

func merge2Channels1(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		results := make([]int, n)
		wg := &sync.WaitGroup{}
		wg.Add(n)
		for i := 0; i < n; i++ {
			valueFromIn1 := <-in1
			valueFromIn2 := <-in2
			go func(wg *sync.WaitGroup, i int) {
				defer wg.Done()
				resultChan1 := make(chan int)
				resultChan2 := make(chan int)

				go func() {
					res := fn(valueFromIn1)
					resultChan1 <- res
				}()

				go func() {
					res := fn(valueFromIn2)
					resultChan2 <- res
				}()

				results[i] = <-resultChan1 + <-resultChan2

			}(wg, i)

		}
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- results[i]
		}
	}()
}

/*
Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

Описание ее работы:

n раз сделать следующее

прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
вычислить f(x1) + f(x2)
записать полученное значение в out
Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.



Формат ввода:

количество итераций передается через аргумент n.
целые числа подаются через аргументы-каналы in1 и in2.
функция для обработки чисел перед сложением передается через аргумент fn.
Формат вывода:

канал для вывода результатов передается через аргумент out.
*/
