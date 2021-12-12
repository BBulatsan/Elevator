package models

func NewElevator(maximumAmount int, availableFloors []int) *Elevator {
	sch := make(chan string)
	return &Elevator{
		MaximumAmount:   maximumAmount,
		AvailableFloors: availableFloors,
		ServiceChanel:   sch,
	}
}
