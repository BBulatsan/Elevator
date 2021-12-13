package models

import (
	"sort"
)

const (
	drop     = "drop"
	pick     = "pick"
	moveDown = "moveDown"
	moveUp   = "moveUp"
)

type Elevator struct {
	Persons         []Person
	RoadMap         []int
	MaximumAmount   int
	AvailableFloors []int
	Place           int
}

type Person struct {
	Begin int
	Dest  int
}

type Log struct {
	id               int
	Name             int
	Place            int
	QuantityOfPeople int
	Action           string
}

func (e *Elevator) Pickup(p Person, log chan Log, name int) {
	e.Persons = append(e.Persons, p)
	act := Log{
		Name:             name,
		Place:            e.Place,
		QuantityOfPeople: e.QuantityOfPeople(),
		Action:           pick,
	}
	log <- act
	e.Mapping(e.Persons)
}

func (e *Elevator) dropout(p Person, log chan Log, name int) {
	for i, num := range e.Persons {
		if num == p {
			e.Persons = append(e.Persons[:i], e.Persons[i+1:]...)
			e.Mapping(e.Persons)
			break
		}
	}
	act := Log{
		Name:             name,
		Place:            e.Place,
		QuantityOfPeople: e.QuantityOfPeople(),
		Action:           drop,
	}
	log <- act
}

func (e Elevator) QuantityOfPeople() int {
	return len(e.Persons)
}

func (e *Elevator) moveDown(level int, log chan Log, name int) {
	for i := e.Place; i >= level; i-- {
		e.Place = i
		act := Log{
			Name:             name,
			Place:            e.Place,
			QuantityOfPeople: e.QuantityOfPeople(),
			Action:           moveDown,
		}
		log <- act
	}
}

func (e *Elevator) moveUp(level int, log chan Log, name int) {
	for i := e.Place; i <= level; i++ {
		e.Place = i
		act := Log{
			Name:             name,
			Place:            e.Place,
			QuantityOfPeople: e.QuantityOfPeople(),
			Action:           moveUp,
		}
		log <- act
	}
}

func (e *Elevator) Move(log chan Log, name int) {
	for _, level := range e.RoadMap {
		if (e.Place - level) < 0 {
			e.moveUp(level, log, name)
		} else {
			e.moveDown(level, log, name)
		}
		for _, num := range e.Persons {
			if level == num.Dest {
				e.dropout(num, log, name)
			}
		}
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
