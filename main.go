package main

import (
	"time"

	"Elevator/models"
	"Elevator/modules/db"
	"Elevator/modules/log"
	"Elevator/modules/people"
)

func main() {
	qt := 100
	floes1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	floes2 := []int{1, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	ch := make(chan models.Person, qt)
	logs := make(chan models.Log, qt)

	con := db.Db()
	go log.Log(con, logs)

	var elev1 = models.NewElevator(8, floes1)
	var elev2 = models.NewElevator(8, floes2)
	for i := 0; i < 3; i++ {
		go func(id int) {
			elev1.DoWork(ch, logs, id)
		}(i)
	}
	for i := 3; i < 6; i++ {
		go func(id int) {
			elev2.DoWork(ch, logs, id)
		}(i)

	}

	for i := 0; i < qt; i++ {
		go people.GenEveningPeople(ch)
	}

	for {
		if len(ch) != 0 {
			time.Sleep(1 * time.Second)

		} else {
			break
		}
	}
}
