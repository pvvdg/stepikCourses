package main

import (
	"fmt"
	"log"
	"time"
)

const formatParse = time.RFC3339
const formatUnixDate = time.UnixDate

func main() {
	s := ""
	fmt.Scan(&s)
	resTime, err := time.Parse(formatParse, s)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(resTime.Format(formatUnixDate))
}

/*
На стандартный ввод подается строковое представление даты и времени в следующем формате:

1986-04-16T05:20:00+06:00
Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:

Wed Apr 16 05:20:00 +0600 1986

Для более эффективной работы рекомендуется ознакомиться с документацией о содержащихся в модуле time константах.

Sample Input:

1986-04-16T05:20:00+06:00
Sample Output:

Wed Apr 16 05:20:00 +0600 1986
*/
