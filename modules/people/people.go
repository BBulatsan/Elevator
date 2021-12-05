package people

import (
	"math/rand"

	"Elevator/models"
)
// jhgj
func GenPeople(qt int) []models.Person {
	var result []models.Person
	for i := 0; i < qt; i++ {
		begin := 1 + rand.Intn(30-1+1)
		dest := 1 + rand.Intn(30-1+1)
		result = append(result, models.Person{Begin: begin, Dest: dest})
	}
	return result
}

func GenMorningPeople(qt int) []models.Person {
	var result []models.Person
	for i := 0; i < qt; i++ {
		dest := 1 + rand.Intn(30-1+1)
		result = append(result, models.Person{Begin: 1, Dest: dest})
	}
	return result
}

func GenEveningPeople(qt int) []models.Person {
	var result []models.Person
	for i := 0; i < qt; i++ {
		begin := 1 + rand.Intn(30-1+1)
		result = append(result, models.Person{Begin: begin, Dest: 1})
	}
	return result
}
