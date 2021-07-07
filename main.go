package main

import (
	"time"

	"elevatorr/models"
)

func main() {
	//var el = Elevator{}
	//var ep = Elevator{}
	var p0 = models.Person{Begin: 7, Dest: 13}
	var p1 = models.Person{Begin: 5, Dest: 12}
	var p2 = models.Person{Begin: 11, Dest: 1}
	var p3 = models.Person{Begin: 15, Dest: 2}
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
		var i = models.Elevator{}
		i.Pickup(p0)
		i.Pickup(p1)
		i.Pickup(p2)
		i.Pickup(p3)
		go i.Move()
		time.Sleep(1 * time.Second)
	}
}
