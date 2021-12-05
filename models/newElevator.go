package models

func NewElevator(maximumAmount int, availableFloors []int) *Elevator {
	ch := make(chan Person, maximumAmount)
	sch := make(chan string, 10)
	return &Elevator{MaximumAmount: maximumAmount, AvailableFloors: availableFloors, Chanel: ch, ServiceChanel: sch}
}
