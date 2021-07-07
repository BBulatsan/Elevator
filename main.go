package main

import (
	"time"

	"elevatorr/models"
)

func main() {
	//var el = models.Elevator{}
	var p0 = models.Person{Begin: 7, Dest: 13}
	var p1 = models.Person{Begin: 5, Dest: 12}
	var p2 = models.Person{Begin: 11, Dest: 1}
	var p3 = models.Person{Begin: 15, Dest: 2}
	//el.Pickup(p0)
	//el.Pickup(p1)
	//el.Pickup(p3)
	//el.Pickup(p2)
	//el.Move()

	for i := 0; i <= 2; i++ {
		var i = models.Elevator{}
		i.Pickup(p0)
		i.Pickup(p1)
		i.Pickup(p2)
		i.Pickup(p3)
		go i.Move()
		time.Sleep(5 * time.Second)
	}
}
