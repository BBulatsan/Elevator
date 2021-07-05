package models

type Elevator struct {
	persons []Person
	roadMap []int
	place   int
}

type Person struct {
	begin int
	dest  int
}
