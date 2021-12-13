package models

func NewElevator(maximumAmount int, availableFloors []int) *Elevator {
	return &Elevator{
		MaximumAmount:   maximumAmount,
		AvailableFloors: availableFloors,
	}
}
