package models

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var wg2 sync.WaitGroup

func DoWork(elev *Elevator, ch chan Person, name int) {
	//for {
	select {
	case x := <-elev.ServiceChanel:

		switch x {
		case "do":
			numPeople := 0
			numFall := 0
			fmt.Println("start")
			{
				for i := 0; i < elev.MaximumAmount; i++ {
					a := <-ch
					fmt.Println(a)
					sort.Ints(elev.AvailableFloors)
					if (elev.MaximumAmount > len(elev.Persons)) && (a.Begin >= elev.AvailableFloors[1] || (a.Begin == 1)) && (a.Dest <= elev.AvailableFloors[len(elev.AvailableFloors)-1]) {
						elev.Chanel <- a
						numPeople++
					} else {
						//fmt.Printf("out of number: %d > %d\n", elev.MaximumAmount, len(elev.Persons))
						//fmt.Printf("begin %d >= available %d\n", a.Begin, elev.AvailableFloors[1])
						//fmt.Printf("dest %d <= available %d\n\n", a.Dest, elev.AvailableFloors[len(elev.AvailableFloors)-1])
						numFall++
						i--
						ch <- a
						if numFall == elev.MaximumAmount {
							break
						}
					}
				}
				wg1.Add(1)
				for j := 0; j < numPeople; j++ {
					wg2.Add(1)
					go elev.Pickup(elev.Chanel, &wg2)
				}
				wg2.Wait()
				go elev.Move(name, &wg1)
				wg1.Wait()
			}
		case "exit":
			fmt.Println("exit")
			close(elev.Chanel)
			close(elev.ServiceChanel)
		}
	default:
		fmt.Println("default")
		time.Sleep(1 * time.Second)
	}
}

//}
