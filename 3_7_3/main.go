package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	inputedTime, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	inputedTime = inputedTime[:len(inputedTime)-2] // For Windows 10 symbols of end string "\r" and "\n"
	stringSplit := strings.Split(inputedTime, ",")
	firstTime := stringSplit[0]
	secondTime := stringSplit[1]

	t1, err := time.Parse("02.01.2006 15:04:05", firstTime)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	t2, err := time.Parse("02.01.2006 15:04:05", secondTime)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	if t1.After(t2) {
		fmt.Println(t1.Sub(t2).String())
	} else {
		fmt.Println(t2.Sub(t1).String())
	}
}

/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).

Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

Sample Input:

13.03.2018 14:00:15,12.03.2018 14:00:15
12.03.2018 14:00:15,13.03.2018 14:00:15
Sample Output:

24h0m0s
*/
