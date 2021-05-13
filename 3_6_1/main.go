package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//All struct
/*
type Students struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []uint8
}

type Group struct {
	ID       int64
	Number   string
	Year     int8
	Students []Students
}
*/
type Students struct {
	Rating []uint8
}

type Group struct {
	Students []Students
}

type Result struct {
	Average float64
}

func main() {
	var newGroup Group

	//For file
	file, err := os.Open("students.json")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()
	//For os.Stdin use reader and change input parameter in method ReadAll
	// reader := bufio.NewReader(os.Stdin)
	dataFromFile, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if err := json.Unmarshal(dataFromFile, &newGroup); err != nil {
		log.Fatalf("%s", err)
	}
	countStudents := 0
	countRating := 0
	for _, v := range newGroup.Students {
		countRating += len(v.Rating)
		countStudents++
	}
	generalResult := float64(countRating) / float64(countStudents) // общий результат
	dataMarshalIndent, err := json.MarshalIndent(Result{Average: generalResult}, "", "    ")
	if err != nil {
		log.Fatalf("%s", err)
	}
	_, err = os.Stdout.Write(dataMarshalIndent)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

/*В сведениях о каждом студенте содержится информация о полученных им оценках (Rating). Требуется прочитать
данные, и рассчитать среднее количество оценок, полученное студентами группы. Ответ на задачу требуется
записать на стандартный вывод в формате JSON в следующей форме:
Example result:
{
    "Average": 14.1
}
Как вы понимаете, для декодирования используется функция Unmarshal, а для кодирования MarshalIndent
(префикс - пустая строка, отступ - 4 пробела).*/
