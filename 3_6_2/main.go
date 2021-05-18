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
