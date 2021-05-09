package main

import (
	"fmt"
)

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	stroka := "112336445556886"

	go func() {
		for _, val := range stroka {
			inputStream <- string(val)
		}
		close(inputStream)
	}()

	go removeDuplicates(inputStream, outputStream)

	for x := range outputStream {
		fmt.Print(x)
	}
}

func removeDuplicates(inputStream, outputStream chan string) {
	bufferValues := ""
	for x := range inputStream {
		bufferValues += x
	}
	resultString := func(s string) string {
		sToRune := []rune(s)
		keys := make(map[rune]bool)
		list := []rune{}
		for _, entry := range sToRune {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				list = append(list, entry)
			}
		}
		return string(list)
	}(bufferValues)

	for _, v := range resultString {
		outputStream <- string(v)
	}

	close(outputStream)
}

/*
Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий
этап конвейера только если оно отличается от того, что пришло ранее.

Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки,
 во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения,
 которые не повторяются подряд. Не забудьте закрыть канал ;)

Функция должна называться removeDuplicates()
*/
