package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt(numbers ...int) (sum int, count int) {
	count = len(numbers)
	for _, val := range numbers {
		sum += val
	}
	return count, sum
}

/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных
 функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/
