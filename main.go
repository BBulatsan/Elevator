package main

import (
	"time"

	"Elevator/models"
	"Elevator/modules/people"
)

func main() {
	floes1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	Peoples := people.GenPeople(10)
	for i := 0; i < 3; i++ {
		var elev = models.Elevator{MaximumAmount: 8, AvailableFloors: floes1}
		for _, person := range Peoples {
			elev.Pickup(person)
		}
		go elev.Move(i)
	}
	time.Sleep(10 * time.Second)
}
