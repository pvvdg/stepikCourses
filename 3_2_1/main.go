package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str1, str2 := "", ""
	fmt.Scan(&str1, &str2)
	fmt.Println(adding(str1, str2))
}

func adding(str1 string, str2 string) int64 {
	return delNoise(str1) + delNoise(str2)
}

func delNoise(str string) int64 {
	strInRAnge := []rune(str)
	bufString := ""
	for _, val := range strInRAnge {
		if unicode.IsDigit(val) { //  Only positive numbers
			bufString += string(val)
		}
	}
	result, _ := strconv.Atoi(bufString)
	return int64(result)
}

/*


Представьте что вы работаете в большой компании где используется модульная архитектура. Ваш коллега написал модуль с какой-то логикой
(вы не знаете) и передает в вашу программу какие-то данные. Вы же пишите функцию которая считывает две переменных типа string ,
а возвращает число типа int64 которое нужно получить сложением этим строк.


Но было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того что гоферам платят больше. Поэтому он решил подшутить
 над вами и подсунул вам подвох. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.


Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и спец знаки.
 Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. Они уже импортированы, вам ничего
 импортировать не нужно!

Считывать и выводить ничего не нужно!

Ваша функция должна называться adding() !

Считайте что функция и пакет main уже объявлены!

Sample Input:

%^80 hhhhh20&&&&nd

Sample Output:

100


*/
