package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ID struct {
	Id int `json:"global_id"`
}

func main() {
	file, err := os.Open("data-20190514T0100.json")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	arrayId := []ID{}
	dec.Decode(&arrayId)
	sum := 0
	for _, v := range arrayId {
		sum += v.Id
	}
	fmt.Printf("%v", sum)
}

/* Данная задача ориентирована на реальную работу с данными в формате JSON. Для решения мы будем использовать
 справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности), который был размещен на
 web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json, а сами данные,
 которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в нашем
 репозитории на github.com.
https://github.com/semyon-dev/stepik-go/tree/master/work_with_json

Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в качестве ответа
 записать сумму полей global_id всех элементов, закодированных в файле.
*/
