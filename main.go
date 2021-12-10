package main

import (
	"Elevator/models"
	"Elevator/modules/people"
)

func main() {
	qt := 100
	floes1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	floes2 := []int{1, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	ch := make(chan models.Person, qt)
	for i := 0; i < qt; i++ {
		go people.GenEveningPeople(ch)
	}
	var elev1 = models.NewElevator(8, floes1)
	var elev2 = models.NewElevator(8, floes2)

	go func() {
		for {
			if len(ch) != 0 {
				elev1.ServiceChanel <- "do"
				elev2.ServiceChanel <- "do"
			} else {
				elev1.ServiceChanel <- "exit"
				elev2.ServiceChanel <- "exit"
				break
			}
		}
	}()

	elev1.DoWork(ch, 1)
	elev2.DoWork(ch, 2)

	defer close(ch)
}
