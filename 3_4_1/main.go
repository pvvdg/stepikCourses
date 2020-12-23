package main

import (
	"fmt"
	"os"
)

func readTask() (value1, value2, operation interface{}) {
	var v1, v2 float64
	var o string
	fmt.Scan(&v1, &v2, &o)
	return v1, v2, o
}

func printError(value1, value2 interface{}) {
	switch value1.(type) {
	case float64:
		break
	default:
		fmt.Printf("value=%v: %T\n", value1, value1)
		os.Exit(0)
	}
	switch value2.(type) {
	case float64:
		break
	default:
		fmt.Printf("value=%v: %T\n", value2, value2)
		os.Exit(0)
	}

}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции

	//fmt.Printf("%v:%T, %v:%T, %v:%T \n", value1, value1, value2, value2, operation, operation)

	switch operation.(string) {
	case "+":
		printError(value1, value2)
		res := value1.(float64) + value2.(float64)
		fmt.Printf("%.4f", res)
	case "-":
		printError(value1, value2)
		res := value1.(float64) - value2.(float64)
		fmt.Printf("%.4f", res)
	case "/":
		res := value1.(float64) / value2.(float64)
		fmt.Printf("%.4f", res)
	case "*":
		printError(value1, value2)
		res := value1.(float64) * value2.(float64)
		fmt.Printf("%.4f", res)
	default:
		fmt.Printf("value=%v: %T\n", value2, value2)
		os.Exit(0)
	}

}

/*
Пришло время для задач, где вы сможете применить полученные знания на практике.

Обязательные условия выполнения: данные со стандартного ввода читаются функцией readTask(), которая возвращает 3
значения типа пустой интерфейс. Эта функция использует пакеты encoding/json, fmt, и os - не удаляйте их из импорта.
 Скорее всего, вам понадобится пакет "fmt", но вы можете использовать любой другой пакет для записи в стандартный
  вывод не удаляя fmt.

И так, вы получаете 3 значения типа пустой интерфейс: если все удачно, то первые 2 значения будут float64. Третье
 значение в идеальном случае будет строкой, коточрая может иметь знаения: "+", "-", "*", "/" (определенная
	математическая операция). Но такие идеальные случаи будут не всегда, вы можете получить значения других типов,
	 а также строка в третьем значении может не относится к одной из указанных математических операций.

Результат выполнения программы должен быть таков:

в штатной ситуации вы должны напечатать в стандартный вывод результат выполнения математической операции с точностью
 до 4 цифры после запятой (fmt.Printf(%.4f));
если первое или второе значение не является типом float64, вы должны напечатать сообщение об ошибке вида:
 value=полученное_значение: тип_значения (например: value=true: bool)
если третье значение имеет неверный тип или передан знак, не относящийся к указанным выше математическим операциям,
 сообщение об ошибке должно иметь вид: неизвестная операция
Гарантируется, что ошибка в аргументах может быть только одна, поэтому если вы при проверке первого значения увидели,
 что оно содержит ошибку - печатайте сообщение об ошибке и завершайте работу программы, проверка остальных аргументов
  значения уже не имеет, а проверяющая система воспримет 2 сообщения об ошибке как нарушение условия выполнения задания.

Удачи!
*/
