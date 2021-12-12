package models

import (
	"fmt"
	"sort"
)

func (e *Elevator) DoWork(ch chan Person, name int) {
	sort.Ints(e.AvailableFloors)
	for {
		fmt.Printf("start elev%d\n", name)
		numFall := 0
		for i := 0; i < e.MaximumAmount; i++ {
			isAvailable := false
			a, isOk := <-ch
			if isOk {
				//isMaxPeople :=
				for _, p := range e.AvailableFloors {
					fmt.Println("Begin: ", p, a.Begin)
					if a.Begin == p {
						for _, j := range e.AvailableFloors {
							fmt.Println("Dest: ", j, a.Dest)
							if a.Dest == j {
								isAvailable = true
								break
							}
						}
						break
					}
				}
				if (e.MaximumAmount >= len(e.Persons)) && isAvailable {
					fmt.Printf("take elev%d %d %d \n", name, a.Begin, a.Dest)
					e.Pickup(a)
				} else {
					ch <- a
					numFall++
					i--
					if numFall == e.MaximumAmount {
						fmt.Printf("finish elev%d\n", name)
						break
					}

				}
			}
		}
		e.Move(name)
	}
}
