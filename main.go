package main

import (
	"fmt"
	"sort"
	"time"
	//"Elevatorv2/interfaces"
	//"Elevatorv2/models"
)

type Elevatorer interface {
	Move
	Move([]int)
}
type Move interface {
	moveDown(int)
	moveUp(int)
}
type Elevator struct {
	persons []Person
	roadMap []int
	place   int
}
type Person struct {
	begin int
	dest  int
}

func (e Elevator) pickup(p Person) {
	e.persons = append(e.persons, p)
	e.Mapping(e.persons)
}

func (e *Elevator) dropout(p Person) {
	for i, num := range e.persons {
		if num == p {
			e.persons = append(e.persons[:i], e.persons[i+1:]...)
			e.Mapping(e.persons)
		}
	}
}

func (e Elevator) quantityOfPeople() {
	fmt.Println("Кол-во людей в лифте =", len(e.persons))
}

func (e *Elevator) moveDown(level int) {
	for i := e.place; i >= level; i-- {
		//time.Sleep(1 * time.Second) //для красоты, реализма
		fmt.Println("Moving down\n level:", i)
		e.place = i
	}
}

func (e *Elevator) moveUp(level int) {
	for i := e.place; i <= level; i++ {
		//time.Sleep(1 * time.Second)
		fmt.Println("Moving up\n level:", i)
		e.place = i
	}
}

func (e *Elevator) Move() {
	for _, level := range e.roadMap {
		if (e.place - level) < 0 {
			e.moveUp(level)
		} else {
			e.moveDown(level)
		}
		for _, num := range e.persons {
			if level == num.dest {
				e.dropout(num)
			}
		}
		fmt.Println("Приехал на ", e.place)
		e.quantityOfPeople()
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
			NumbersForMapping = append(NumbersForMapping, num.begin)
			NumbersForMapping = append(NumbersForMapping, num.dest)
			EndNumber = num.dest
		} else {
			if (num.begin >= NumbersForMapping[0] && num.begin <= EndNumber) && (num.begin < num.dest) {
				NumbersForMapping = append(NumbersForMapping, num.begin)
				NumbersForMapping = append(NumbersForMapping, num.dest)
			} else if (num.begin >= NumbersForMapping[0] && num.begin <= EndNumber) && (num.begin > num.dest) {
				NumbersForMapping = append(NumbersForMapping, num.begin)
				BeforeNumbers = append(BeforeNumbers, num.dest)
			} else if (num.begin > EndNumber) && (num.begin < num.dest) {
				NumbersForMapping = append(NumbersForMapping, num.begin)
				NumbersForMapping = append(NumbersForMapping, num.dest)
			} else if (num.begin > EndNumber) && (num.begin > num.dest) {
				NumbersForMapping = append(NumbersForMapping, num.begin)
				BeforeNumbers = append(BeforeNumbers, num.dest)
			} else if (num.begin < NumbersForMapping[0]) && (num.begin > num.dest) {
				AfterNumbers = append(AfterNumbers, num.begin)
				AfterNumbers = append(AfterNumbers, num.dest)
			} else if (num.begin < NumbersForMapping[0]) && (num.begin < num.dest) {
				AfterNumbers = append(AfterNumbers, num.begin)
				AfterNumbers = append(AfterNumbers, num.dest)
			}
		}
	}
	sort.Ints(NumbersForMapping)
	sort.Ints(AfterNumbers)
	sort.Ints(BeforeNumbers)
	e.roadMap = append(NumbersForMapping, BeforeNumbers...)
	e.roadMap = append(e.roadMap, AfterNumbers...)
}

func main() {
	//var el = Elevator{}
	//var ep = Elevator{}
	var p0 = Person{7, 13}
	var p1 = Person{5, 12}
	var p2 = Person{11, 1}
	var p3 = Person{15, 2}
	//el.pickup(p0)
	//el.pickup(p1)
	//el.pickup(p3)
	//el.pickup(p2)
	//ep.pickup(p0)
	//ep.pickup(p1)
	//ep.pickup(p3)
	//ep.pickup(p2)
	//el.Move()
	//ep.Move()
	for i := 0; i <= 2; i++ {
		var i = Elevator{}
		i.pickup(p0)
		i.pickup(p1)
		i.pickup(p2)
		i.pickup(p3)
		go i.Move()
		time.Sleep(1 * time.Second)
	}
}
