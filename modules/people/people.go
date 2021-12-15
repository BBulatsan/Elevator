package people

import (
	"context"
	"math/rand"

	"Elevator/models"
)

func GenPeople(ctx context.Context, ch chan models.Person) {
	select {
	case <-ctx.Done():
		return
	default:
		begin := 1 + rand.Intn(30)
		dest := 1 + rand.Intn(30)
		ch <- models.Person{Begin: begin, Dest: dest}
	}
}

func GenMorningPeople(ctx context.Context, ch chan models.Person) {
	select {
	case <-ctx.Done():
		return
	default:
		dest := 2 + rand.Intn(29)
		ch <- models.Person{Begin: 1, Dest: dest}
	}
}

func GenEveningPeople(ctx context.Context, ch chan models.Person) {
	select {
	case <-ctx.Done():
		return
	default:
		begin := 2 + rand.Intn(29)
		ch <- models.Person{Begin: begin, Dest: 1}
	}
}
