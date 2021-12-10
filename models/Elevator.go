package models

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var wg2 sync.WaitGroup

func (e *Elevator) DoWork(ch chan Person, name int) {
	time.Sleep(1 * time.Second)
	go func() {
		for {
			select {
			case x := <-e.ServiceChanel:

				switch x {
				case "do":
					numPeople := 0
					numFall := 0
					{
						for i := 0; i < e.MaximumAmount; i++ {
							a := <-ch
							//fmt.Println(a)
							sort.Ints(e.AvailableFloors)
							if (e.MaximumAmount >= len(e.Persons)) && (a.Begin >= e.AvailableFloors[1] || (a.Begin == 1)) && (a.Dest <= e.AvailableFloors[len(e.AvailableFloors)-1]) {
								e.Chanel <- a
								numPeople++
							} else {
								//fmt.Printf("out of number: %d > %d\n", e.MaximumAmount, len(e.Persons))
								//fmt.Printf("begin %d >= available %d\n", a.Begin, e.AvailableFloors[1])
								//fmt.Printf("dest %d <= available %d\n\n", a.Dest, e.AvailableFloors[len(e.AvailableFloors)-1])
								numFall++
								i--
								ch <- a
								if numFall == e.MaximumAmount {
									break
								}
							}
						}
						wg1.Add(1)
						for j := 0; j < numPeople; j++ {
							wg2.Add(1)
							go e.Pickup(e.Chanel, &wg2)
						}
						wg2.Wait()
						go e.Move(name, &wg1)
						wg1.Wait()
					}
				case "exit":
					fmt.Println("exit")
					close(e.Chanel)
					close(e.ServiceChanel)
				}
			default:
				fmt.Println("default")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}
