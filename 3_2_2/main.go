package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve()
	//solveCSV("base.csv")
}

func solve() {
	bufSliceFloat := make([]float64, 0)
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	strDelimit := strings.Split(str, ";")
	for _, val := range strDelimit {
		valWithoutSpaces := strings.Replace(strings.Replace(val, " ", "", -1), ",", ".", -1)
		valWithoutSpacesFloat, _ := strconv.ParseFloat(valWithoutSpaces, 64)
		bufSliceFloat = append(bufSliceFloat, valWithoutSpacesFloat)
	}
	fmt.Printf("%.4f \n", bufSliceFloat[0]/bufSliceFloat[1])
}

func solveCSV(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';' //delimiter (разделитель)
	records, _ := reader.Read()

	nums := make([]float64, 0)
	for _, val := range records {
		val = strings.Replace(strings.Replace(val, " ", "", -1), ",", ".", -1)
		numBuf, _ := strconv.ParseFloat(val, 64)
		nums = append(nums, numBuf)
	}
	result := nums[0] / nums[1]
	fmt.Printf("%.4f \n", result)
}

/*

Для решения данной задачи вам понадобится пакет strconv, возможно использовать пакеты strings или encoding/csv, или даже bufio - вы
не ограничены в выборе способа решения задачи. В решениях мы поделимся своими способами решения этой задачи, предлагаем вам сделать
 то же самое.

В привычных нам редакторах электронных таблиц присутствует удобное представление числа с разделителем разрядов в виде пробела, кроме
 того в России целая часть от дробной отделяется запятой. Набор таких чисел был экспортирован в формат CSV, где в качестве разделителя
  используется символ ";".

На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное от деления первого числа
 на второе с точностью до четырех знаков после "запятой" (на самом деле после точки, результат не требуется приводить к исходному
	 формату).

P.S. небольшое отступление, связанное с чтением из стандартного ввода. Кто-то захочет использовать для этого пакет bufio.Reader.
Это вполне допустимый вариант, но если вы резонно обрабатываете ошибку метода ReadString('\n'), то получаете ошибку EOF, а точнее
(io.EOF - End Of File). На самом деле это не ошибка, а состояние, означающее, что файл (а os.Stdin является файлом) прочитан до
конца. Чтобы ошибка была обработана правильно, вы можете поступить так:

if err != nil && err != io.EOF {
	...
}

Sample Input:

1 149,6088607594936;1 179,0666666666666

Sample Output:

0.9750


*/
