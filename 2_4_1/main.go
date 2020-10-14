package main

import "fmt"

type robotGame struct {
	On    bool
	Ammo  int
	Power int
}

func (r *robotGame) Shoot() bool {
	if r.On == true {
		if r.Ammo > 0 {
			r.Ammo--
			return true
		}
	}
	return false
}

func (r *robotGame) RideBike() bool {
	if r.On == true {
		if r.Power > 0 {
			r.Power--
			return true
		}
	}
	return false
}

func main() {
	testStruct := new(robotGame)
	*testStruct = robotGame{true, 22, 32}
	testStruct.Shoot()
	testStruct.RideBike()
	fmt.Println(*testStruct)
}

/*
Необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
 У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет,
 то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
*/
