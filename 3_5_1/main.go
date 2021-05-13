package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("task.data")
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comma = ';'
	records, err := r.Read()
	for i, v := range records {
		if v == "0" {
			fmt.Println(i + 1)
		}
	}
}

/*
В тестовом файле, который вы можете скачать из нашего репозитория на github.com, содержится длинный ряд чисел,
 разделенных символом ";". Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа.
  Требуется вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).

Например:  12;234;6;0;78 :
Правильный ответ будет порядковый номер числа: 4
*/
