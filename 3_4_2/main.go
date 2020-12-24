package main

import "fmt"

// Battery is global name
type Battery interface {
	String() string
}

//Li is type battery
type Li struct {
	countCharge string
}

func sortAndCheck(str string) string {
	strToRune := []rune(str)
	if len(strToRune) != 10 {
		panic("Count numbers != 10.")
	}
	for i := 0; i < len(strToRune); i++ {
		for j := i; j < len(strToRune); j++ {
			if strToRune[i] > strToRune[j] {
				tmp := strToRune[i]
				strToRune[i] = strToRune[j]
				strToRune[j] = tmp
			}
		}
	}
	resultString := make([]rune, 0)
	for _, val := range strToRune {
		if val == '0' {
			resultString = append(resultString, ' ')
		} else if val == '1' {
			resultString = append(resultString, 'X')
		} else {
			panic("Input 0 or 1")
		}
	}
	return string(resultString)
}

func (l *Li) String() string {
	return fmt.Sprintf("[%v]", l.countCharge)
}

func main() {
	fmt.Println("Input 10 numbers 0 or 1.")
	chargeString := ""
	_, err := fmt.Scan(&chargeString)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	chargeString = sortAndCheck(chargeString)
	var batteryForTest Battery = &Li{chargeString}
	result := batteryForTest.String()
	fmt.Println(result)

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
