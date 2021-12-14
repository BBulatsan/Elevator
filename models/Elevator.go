package models

import "context"

const httpCanceled = "httpCanceled"

func (e *Elevator) DoWork(ctx context.Context, ch chan Person, log chan Log, name int) {
	for {
		select {
		case <-ctx.Done():
			act := Log{
				Name:             name,
				Place:            e.Place,
				QuantityOfPeople: e.QuantityOfPeople(),
				Action:           httpCanceled,
			}
			log <- act
			return
		default:
			numFall := 0
			for i := 0; i < e.MaximumAmount; i++ {
				isAvailable := false
				a, isOk := <-ch
				if !isOk {
					break
				}
				for _, p := range e.AvailableFloors {
					if a.Begin == p {
						for _, j := range e.AvailableFloors {
							if a.Dest == j {
								isAvailable = true
								break
							}
						}
						break
					}
				}
				if (e.MaximumAmount >= len(e.Persons)) && isAvailable {
					e.Pickup(a, log, name)
				} else {
					ch <- a
					numFall++
					i--
					if numFall == e.MaximumAmount {
						break
					}

				}
			}
		}
		e.Move(log, name)
	}
}
