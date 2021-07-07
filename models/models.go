package models

import (
	"fmt"
	"sort"
)

type Elevator struct {
	Persons []Person
	RoadMap []int
	Place   int
}

type Person struct {
	Begin int
	Dest  int
}

func (e *Elevator) Pickup(p Person) {
	e.Persons = append(e.Persons, p)
	e.Mapping(e.Persons)
}

func (e *Elevator) dropout(p Person) {
	for i, num := range e.Persons {
		if num == p {
			e.Persons = append(e.Persons[:i], e.Persons[i+1:]...)
			e.Mapping(e.Persons)
		}
	}
}

func (e Elevator) QuantityOfPeople() {
	fmt.Println("Кол-во людей в лифте =", len(e.Persons))
}

func (e *Elevator) moveDown(level int) {
	for i := e.Place; i >= level; i-- {
		//time.Sleep(1 * time.Second) //для красоты, реализма
		fmt.Println("Moving down\n level:", i)
		e.Place = i
	}
}

func (e *Elevator) moveUp(level int) {
	for i := e.Place; i <= level; i++ {
		//time.Sleep(1 * time.Second)
		fmt.Println("Moving up\n level:", i)
		e.Place = i
	}
}

func (e *Elevator) Move() {
	for _, level := range e.RoadMap {
		if (e.Place - level) < 0 {
			e.moveUp(level)
		} else {
			e.moveDown(level)
		}
		for _, num := range e.Persons {
			if level == num.Dest {
				e.dropout(num)
			}
		}
		fmt.Println("Приехал на ", e.Place)
		e.QuantityOfPeople()
		//time.Sleep(1 * time.Second)
	}
}

func (e *Elevator) Mapping(places []Person) {
	var NumbersForMapping []int
	var AfterNumbers []int
	var BeforeNumbers []int
	var EndNumber int
	for _, num := range places {
		if len(NumbersForMapping) == 0 {
			NumbersForMapping = append(NumbersForMapping, num.Begin)
			NumbersForMapping = append(NumbersForMapping, num.Dest)
			EndNumber = num.Dest
		} else {
			if (num.Begin >= NumbersForMapping[0] && num.Begin <= EndNumber) && (num.Begin < num.Dest) {
				NumbersForMapping = append(NumbersForMapping, num.Begin)
				NumbersForMapping = append(NumbersForMapping, num.Dest)
			} else if (num.Begin >= NumbersForMapping[0] && num.Begin <= EndNumber) && (num.Begin > num.Dest) {
				NumbersForMapping = append(NumbersForMapping, num.Begin)
				BeforeNumbers = append(BeforeNumbers, num.Dest)
			} else if (num.Begin > EndNumber) && (num.Begin < num.Dest) {
				NumbersForMapping = append(NumbersForMapping, num.Begin)
				NumbersForMapping = append(NumbersForMapping, num.Dest)
			} else if (num.Begin > EndNumber) && (num.Begin > num.Dest) {
				NumbersForMapping = append(NumbersForMapping, num.Begin)
				BeforeNumbers = append(BeforeNumbers, num.Dest)
			} else if (num.Begin < NumbersForMapping[0]) && (num.Begin > num.Dest) {
				AfterNumbers = append(AfterNumbers, num.Begin)
				AfterNumbers = append(AfterNumbers, num.Dest)
			} else if (num.Begin < NumbersForMapping[0]) && (num.Begin < num.Dest) {
				AfterNumbers = append(AfterNumbers, num.Begin)
				AfterNumbers = append(AfterNumbers, num.Dest)
			}
		}
	}
	sort.Ints(NumbersForMapping)
	sort.Ints(AfterNumbers)
	sort.Ints(BeforeNumbers)
	e.RoadMap = append(NumbersForMapping, BeforeNumbers...)
	e.RoadMap = append(e.RoadMap, AfterNumbers...)
}
