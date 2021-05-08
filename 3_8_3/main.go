package main

import "fmt"

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	stroka := "11233445556886"

	go func() {
		for _, val := range stroka {
			inputStream <- string(val)
		}
		close(inputStream)
	}()

	go removeDuplicates(inputStream, outputStream)

	for x := range outputStream {
		fmt.Println(x)
	}
}

func removeDuplicates(inputStream, outputStream chan string) {
	bufferValue := <-inputStream
	for x := range inputStream {
		if x != bufferValue {
			outputStream <- bufferValue
			bufferValue = x
		}
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
