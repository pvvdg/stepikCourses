package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	s = s[:len(s)-2] // For Windows 10 symbols of end string "\r" and "\n"
	timeResult, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	hours := timeResult.Hour()
	minutes := timeResult.Minute()
	if hours > 13 || (minutes > 0 && hours >= 13) {
		timeResult = timeResult.AddDate(0, 0, 1)
		io.WriteString(os.Stdout, timeResult.Format("2006-01-02 15:04:05"))
	} else {
		io.WriteString(os.Stdout, s)
	}
}

/*
На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:

2020-05-15 08:00:00
Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.

Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день, а затем вывести на стандартный вывод
 в том же формате.

Sample Input:

2020-05-15 08:00:00
Sample Output:

2020-05-15 08:00:00
*/
