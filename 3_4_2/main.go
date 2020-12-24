package main

import "fmt"

// Battery is
type Battery interface {
	String() string
}

//Li is
type Li struct {
	countCharge string
}

func bubbleSort(str string) string {
	strToRune := []rune(str)
	for i := 0; i < len(strToRune); i++ {
		for j := i; j < len(strToRune); j++ {
			if strToRune[i] > strToRune[j] {
				swap(strToRune, i, j)
			}
		}
	}
	return string(strToRune)
}

func swap(ar []rune, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}

/*
func sort(s string) string {
	resultSortString := make([]rune, 0)
	sToRune := []rune(s)
	firstValueInString := sToRune[0]
	for _, v := range sToRune {
		if v > firstValueInString {
			firstValueInString = v
			resultSortString = append(resultSortString, v)
		}
	}
	return string(resultSortString)
}


	resultString := ""
	for _, v := range l.countCharge {
		if v == '0' {
			resultString = append(resultString, ' ')
		} else if v == '1' {
			resultString = append(resultString, ' ')
		} else {

		}
	}
}*/

func (l *Li) String() string {
	return fmt.Sprintf("[%v]", l.countCharge)
}

func main() {
	chargeString := ""
	fmt.Scan(&chargeString)
	chargeString = bubbleSort(chargeString)
	fmt.Println(chargeString)
	// var batteryForTest Battery = &Li{chargeString}
	// batteryForTest.String()

}

/*

Давайте используем ваши знания структур, методов и интерфейсов на практике и реализуем объект, удовлетворяющий
 интерфейсу fmt.Stringer. Назовем его "Батарейка".

Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.

Ваш тип должен предусматривать, что на печати он будет выглядеть так: [      XXXX]: где пробелы - "опустошенная"
 емкость батареи, а X - "заряженная".

Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр: 0 или 1 (порядок 0/1 случайный).
 Ваша задача считать эту строку любым возможным способом и создать на основе этой строки объект объявленного вами
  на первом этапе типа: надеюсь, вы понимаете, что строка символизирует емкость батарейки: 0 - это "опустошенная"
   часть, а 1 - "заряженная".

В-третьих, созданный вами объект должен называться batteryForTest (использование этого имени обязательно).

В вашем распоряжении фактически весь файл, НО завершающая фигурная скобка функции main() вам не видна, но она
присутствует. Перед этой скобкой присутствует функция (которая вам тоже не видна), принимающая в качестве аргумента
 объект типа fmt.Stringer - batteryForTest, и направляющая его на стандартный вывод, поэтому вам не требуется выводить
  что-то на печать самостоятельно.

Удачи!

Sample Input:

1000010011
Sample Output:

[      XXXX]

ps. ох уж эти формулировки

Нужно:

1) Создать тип (например, Battery)

2) Расширить его методом, чтобы он удовлетворял Stringer

3) Считать инпут, например, в переменную test

4) Создать экземпляр вида  batteryForTest := Battery(text)

*/
