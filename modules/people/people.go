package people

import (
	"Elevator/models"
	"math/rand"
)

func GenPeople(ch chan models.Person) {
	begin := 1 + rand.Intn(30-1+1)
	dest := 1 + rand.Intn(30-1+1)
	ch <- models.Person{Begin: begin, Dest: dest}
}

func GenMorningPeople(ch chan models.Person) {
	dest := 2 + rand.Intn(29)
	ch <- models.Person{Begin: 1, Dest: dest}
}

func GenEveningPeople(ch chan models.Person) {
	begin := 2 + rand.Intn(29)
	ch <- models.Person{Begin: begin, Dest: 1}
}
