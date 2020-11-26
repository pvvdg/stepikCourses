package main

import (
	"fmt"
	"strconv"
)

func fn(num uint) uint {
	if num > 0 {
		newStr := ""
		numToStr := strconv.Itoa(int(num))
		for _, val := range numToStr {
			intNum, _ := strconv.Atoi(string(val))
			if intNum%2 == 0 {
				newStr += strconv.Itoa(intNum)
			}
		}
		resInt, _ := strconv.Atoi(newStr)
		return uint(resInt)
	}
	return 100
}

func main() {
	var num uint
	fmt.Scan(&num)
	fmt.Println(fn(num))
}

/*
Используем анонимные функции на практике.

Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри функции
main в дальнейшем можно будет вызвать по имени fn.

Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные
цифры или цифра 0, если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот
 же порядок, что и в исходном числе.

Пакет main объявлять не нужно!
Вводить и выводить что-либо не нужно!
Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.
*/
