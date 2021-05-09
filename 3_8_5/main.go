package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1) // Увеличиваем счетчик горутин в группе
		go func(i int) {
			defer wg.Done()
			work(i)
		}(i)
	}

	wg.Wait() // ожидаем завершения всех горутин в группе
	fmt.Println("Горутины завершили выполнение")
}

func work(id int) {
	fmt.Printf("Горутина %d начала выполнение \n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Горутина %d завершила выполнение \n", id)
}

/*
Внутри функции main, вам необходимо в отдельных горутинах вызвать функцию work()
10 раз и дождаться результатов выполнения вызванных функций.
*/
