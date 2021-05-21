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

const now = 1589570165

func main() {
	inputedTime, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatalf("%v", err)
		return
	}
	inputedTime = inputedTime[:len(inputedTime)-2]
	inputedTime = strings.Replace(inputedTime, "мин.", "m", 1)
	inputedTime = strings.Replace(inputedTime, "сек.", "s", 1)
	inputedTime = strings.Replace(inputedTime, " ", "", -1)
	fmt.Println(inputedTime)
	dur, err := time.ParseDuration(inputedTime)
	unixTime := time.Unix(now+int64(dur.Seconds()), 0)
	fmt.Println(unixTime.UTC())
}

/*
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате
 Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).

Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate)
дату и время, получившиеся при добавлении периода к стандартной дате.

Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

Sample Input:

12 мин. 13 сек.
126 мин. 15 сек.

Sample Output:

Fri May 15 19:28:18 UTC 2020
*/
